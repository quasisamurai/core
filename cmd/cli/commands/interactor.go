package commands

import (
	"time"

	frd "github.com/sonm-io/core/fusrodah/miner"
	pb "github.com/sonm-io/core/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type CliInteractor interface {
	HubPing(context.Context) (*pb.PingReply, error)
	HubStatus(context.Context) (*pb.HubStatusReply, error)
	HubFind(context.Context, time.Duration) ([]*frd.HubInfo, error)

	MinerList(context.Context) (*pb.ListReply, error)
	MinerStatus(minerID string, appCtx context.Context) (*pb.InfoReply, error)

	TaskList(appCtx context.Context, minerID string) (*pb.StatusMapReply, error)
	TaskLogs(appCtx context.Context, req *pb.TaskLogsRequest) (pb.Hub_TaskLogsClient, error)
	TaskStart(appCtx context.Context, req *pb.HubStartTaskRequest) (*pb.HubStartTaskReply, error)
	TaskStatus(appCtx context.Context, taskID string) (*pb.TaskStatusReply, error)
	TaskStop(appCtx context.Context, taskID string) (*pb.StopTaskReply, error)
}

type grpcInteractor struct {
	cc      *grpc.ClientConn
	timeout time.Duration
	hub     pb.HubClient
}

func (it *grpcInteractor) call(addr string) error {
	cc, err := grpc.Dial(addr, grpc.WithInsecure(),
		grpc.WithCompressor(grpc.NewGZIPCompressor()),
		grpc.WithDecompressor(grpc.NewGZIPDecompressor()))
	if err != nil {
		return err
	}

	it.cc = cc
	return nil
}

func (it *grpcInteractor) ctx(appCtx context.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(appCtx, it.timeout)
}

func (it *grpcInteractor) HubPing(appCtx context.Context) (*pb.PingReply, error) {
	ctx, cancel := it.ctx(appCtx)
	defer cancel()
	return pb.NewHubClient(it.cc).Ping(ctx, &pb.PingRequest{})
}

func (it *grpcInteractor) HubStatus(appCtx context.Context) (*pb.HubStatusReply, error) {
	ctx, cancel := it.ctx(appCtx)
	defer cancel()
	return pb.NewHubClient(it.cc).Status(ctx, &pb.HubStatusRequest{})
}

func (it *grpcInteractor) MinerList(appCtx context.Context) (*pb.ListReply, error) {
	ctx, cancel := it.ctx(appCtx)
	defer cancel()
	return pb.NewHubClient(it.cc).List(ctx, &pb.ListRequest{})
}

func (it *grpcInteractor) MinerStatus(minerID string, appCtx context.Context) (*pb.InfoReply, error) {
	ctx, cancel := it.ctx(appCtx)
	defer cancel()

	var req = pb.HubInfoRequest{Miner: minerID}
	return pb.NewHubClient(it.cc).Info(ctx, &req)
}

func (it *grpcInteractor) TaskList(appCtx context.Context, minerID string) (*pb.StatusMapReply, error) {
	ctx, cancel := it.ctx(appCtx)
	defer cancel()

	req := &pb.HubStatusMapRequest{Miner: minerID}
	return pb.NewHubClient(it.cc).MinerStatus(ctx, req)
}

func (it *grpcInteractor) TaskLogs(appCtx context.Context, req *pb.TaskLogsRequest) (pb.Hub_TaskLogsClient, error) {
	return pb.NewHubClient(it.cc).TaskLogs(appCtx, req)
}

func (it *grpcInteractor) TaskStart(appCtx context.Context, req *pb.HubStartTaskRequest) (*pb.HubStartTaskReply, error) {
	ctx, cancel := it.ctx(appCtx)
	defer cancel()
	return pb.NewHubClient(it.cc).StartTask(ctx, req)
}

func (it *grpcInteractor) TaskStatus(appCtx context.Context, taskID string) (*pb.TaskStatusReply, error) {
	ctx, cancel := it.ctx(appCtx)
	defer cancel()

	var req = &pb.TaskStatusRequest{Id: taskID}
	return pb.NewHubClient(it.cc).TaskStatus(ctx, req)
}

func (it *grpcInteractor) TaskStop(appCtx context.Context, taskID string) (*pb.StopTaskReply, error) {
	ctx, cancel := it.ctx(appCtx)
	defer cancel()

	var req = &pb.StopTaskRequest{Id: taskID}
	return pb.NewHubClient(it.cc).StopTask(ctx, req)
}

func (it *grpcInteractor) HubFind(ctx context.Context, to time.Duration) ([]*frd.HubInfo, error) {
	srv, err := frd.NewServer(nil, cfg.Bootnodes()...)
	if err != nil {
		return nil, err
	}

	err = srv.Start()
	if err != nil {
		return nil, err
	}

	srv.Serve()
	// set time to search for hubs into network
	hubs := srv.Rediscovery(time.NewTimer(to), true)
	// todo: try to connect to
	return hubs, nil
}

func NewGrpcInteractor(addr string, to time.Duration) (CliInteractor, error) {
	i := &grpcInteractor{timeout: to}
	if addr != "" {
		err := i.call(addr)
		if err != nil {
			return nil, err
		}
	}

	return i, nil
}
