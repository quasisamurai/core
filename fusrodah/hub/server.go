package hub

import (
	"crypto/ecdsa"

	"fmt"
	"github.com/ethereum/go-ethereum/whisper/whisperv2"
	"github.com/sonm-io/core/common"
	"github.com/sonm-io/core/fusrodah"
	"log"
)

const defaultHubPort = ":30343"

type Server struct {
	PrivateKey *ecdsa.PrivateKey
	Frd        *fusrodah.Fusrodah

	HubIp string
}

func NewServer(prv *ecdsa.PrivateKey, hubIp string) (srv *Server, err error) {
	bootnodes := []string{common.BootNodeAddr, common.SecondBootNodeAddr}
	frd, err := fusrodah.NewServer(prv, defaultHubPort, bootnodes)
	if err != nil {
		return nil, err
	}

	srv = &Server{
		PrivateKey: prv,
		HubIp:      hubIp,
		Frd:        frd,
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
	srv.Frd.AddHandling(nil, nil, func(msg *whisperv2.Message) {
		log.Print("Handling discovery message")
		body := fmt.Sprintf("%s,v1", srv.HubIp)
		srv.Frd.Send(body, true, common.TopicMinerDiscover)
	}, common.TopicHubDiscover)
}
