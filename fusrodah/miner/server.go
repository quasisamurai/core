package miner

import (
	"crypto/ecdsa"
	"log"
	"time"

	"sync"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/whisper/whisperv2"
	"github.com/sonm-io/core/common"
	"github.com/sonm-io/core/fusrodah"
)

const defaultMinerPort = ":30342"

type HubInfo struct {
	Address   string
	PublicKey *ecdsa.PublicKey
}

type Server struct {
	PrivateKey *ecdsa.PrivateKey
	Frd        *fusrodah.Fusrodah
	Hub        *HubInfo
	knownHubs  HubStorage
}

type HubStorage interface {
	PutItem(info *HubInfo)
	GetItem(addr string) (*HubInfo, bool)
	GetAll() []*HubInfo
}

type inMemHubStorage struct {
	sync.RWMutex
	// data maps hub addr to it pub key
	data map[string]*HubInfo
}

func (s *inMemHubStorage) PutItem(info *HubInfo) {
	s.Lock()
	defer s.Unlock()

	s.data[info.Address] = info
}

func (s *inMemHubStorage) GetItem(addr string) (*HubInfo, bool) {
	s.Lock()
	defer s.Unlock()

	if hub, ok := s.data[addr]; ok {
		return hub, ok
	} else {
		return nil, false
	}
}

func (s *inMemHubStorage) GetAll() []*HubInfo {
	s.Lock()
	defer s.Unlock()

	hubs := make([]*HubInfo, 0, len(s.data))
	for _, hub := range s.data {
		hubs = append(hubs, hub)
	}

	return hubs
}

// TODO(sshaman1101): test storage impl
func newInMemHubStorage() HubStorage {
	return &inMemHubStorage{data: map[string]*HubInfo{}}
}

func NewServer(prv *ecdsa.PrivateKey, bootn ...string) (srv *Server, err error) {
	if prv == nil {
		prv, err = crypto.GenerateKey()
		if err != nil {
			return nil, err
		}
	}

	var bootnodes []string
	if len(bootn) == 0 {
		bootnodes = []string{common.BootNodeAddr, common.SecondBootNodeAddr}
	} else {
		bootnodes = bootn
	}

	frd, err := fusrodah.NewServer(prv, defaultMinerPort, bootnodes)
	if err != nil {
		return nil, err
	}

	srv = &Server{
		PrivateKey: prv,
		Frd:        frd,
		knownHubs:  newInMemHubStorage(),
	}

	return srv, nil
}

func (srv *Server) Start() (err error) {
	err = srv.Frd.Start()
	if err != nil {
		return err
	}
	return nil
}

func (srv *Server) Stop() (err error) {
	err = srv.Frd.Stop()
	if err != nil {
		return err
	}
	return nil
}

func (srv *Server) Serve() {
	srv.discovery()
}

func (srv *Server) discovery() {
	var filterID int

	done := make(chan struct{})

	filterID = srv.Frd.AddHandling(nil, nil, func(msg *whisperv2.Message) {
		if hubKey := msg.Recover(); hubKey != nil { // skip unauthenticated messages
			srv.Hub = &HubInfo{
				PublicKey: hubKey,
				Address:   string(msg.Payload),
			}
			srv.Frd.RemoveHandling(filterID)
			close(done)
		}
	}, common.TopicMinerDiscover)

	t := time.NewTicker(time.Second * 1)
	defer t.Stop()
	select {
	case <-done:
		return
	case <-t.C:
		srv.Frd.Send(srv.GetPubKeyString(), true, common.TopicHubDiscover)
	}
}

func (srv *Server) GetHub() *HubInfo {
	if srv.Hub == nil {
		srv.discovery()
	}
	return srv.Hub
}

func (srv *Server) GetPubKeyString() string {
	pkString := string(crypto.FromECDSAPub(&srv.PrivateKey.PublicKey))
	return pkString
}

func (srv *Server) Rediscovery(timer *time.Timer, allowAnon bool) []*HubInfo {
	hubChan := make(chan *HubInfo)
	handlerID := srv.Frd.AddHandling(nil, nil, func(msg *whisperv2.Message) {
		if hubKey := msg.Recover(); hubKey != nil || allowAnon {
			h := &HubInfo{
				PublicKey: hubKey,
				Address:   string(msg.Payload),
			}
			hubChan <- h
		} else {
			log.Printf("Drop anonymous response: %s\r\n", string(msg.Payload))
		}
	}, common.TopicMinerDiscover)

	t := time.NewTicker(5 * time.Second)
	defer t.Stop()

	for {
		select {
		case hub := <-hubChan:
			log.Printf("Read from chan, Saving hub")
			srv.knownHubs.PutItem(hub)
		case <-t.C:
			log.Printf("Broadcasing discovery topic")
			srv.Frd.Send(srv.GetPubKeyString(), true, common.TopicHubDiscover)
		case <-timer.C:
			log.Printf("Timer expires, finishing dicsovery")
			close(hubChan)
			srv.Frd.RemoveHandling(handlerID)
			return srv.knownHubs.GetAll()
		}
	}
}
