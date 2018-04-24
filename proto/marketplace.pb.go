// Code generated by protoc-gen-go. DO NOT EDIT.
// source: marketplace.proto

package sonm

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// grpccmd imports
import (
	"io"

	"github.com/spf13/cobra"
	"github.com/sshaman1101/grpccmd"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type OrderType int32

const (
	OrderType_ANY OrderType = 0
	OrderType_BID OrderType = 1
	OrderType_ASK OrderType = 2
)

var OrderType_name = map[int32]string{
	0: "ANY",
	1: "BID",
	2: "ASK",
}
var OrderType_value = map[string]int32{
	"ANY": 0,
	"BID": 1,
	"ASK": 2,
}

func (x OrderType) String() string {
	return proto.EnumName(OrderType_name, int32(x))
}
func (OrderType) EnumDescriptor() ([]byte, []int) { return fileDescriptor10, []int{0} }

type OrderStatus int32

const (
	OrderStatus_ORDER_INACTIVE OrderStatus = 0
	OrderStatus_ORDER_ACTIVE   OrderStatus = 1
)

var OrderStatus_name = map[int32]string{
	0: "ORDER_INACTIVE",
	1: "ORDER_ACTIVE",
}
var OrderStatus_value = map[string]int32{
	"ORDER_INACTIVE": 0,
	"ORDER_ACTIVE":   1,
}

func (x OrderStatus) String() string {
	return proto.EnumName(OrderStatus_name, int32(x))
}
func (OrderStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor10, []int{1} }

type IdentityLevel int32

const (
	IdentityLevel_ANONIMOUS    IdentityLevel = 0
	IdentityLevel_PSEUDONYMOUS IdentityLevel = 1
	IdentityLevel_IDENTIFIED   IdentityLevel = 2
)

var IdentityLevel_name = map[int32]string{
	0: "ANONIMOUS",
	1: "PSEUDONYMOUS",
	2: "IDENTIFIED",
}
var IdentityLevel_value = map[string]int32{
	"ANONIMOUS":    0,
	"PSEUDONYMOUS": 1,
	"IDENTIFIED":   2,
}

func (x IdentityLevel) String() string {
	return proto.EnumName(IdentityLevel_name, int32(x))
}
func (IdentityLevel) EnumDescriptor() ([]byte, []int) { return fileDescriptor10, []int{2} }

type DealStatus int32

const (
	DealStatus_DEAL_UNKNOWN  DealStatus = 0
	DealStatus_DEAL_ACCEPTED DealStatus = 1
	DealStatus_DEAL_CLOSED   DealStatus = 2
)

var DealStatus_name = map[int32]string{
	0: "DEAL_UNKNOWN",
	1: "DEAL_ACCEPTED",
	2: "DEAL_CLOSED",
}
var DealStatus_value = map[string]int32{
	"DEAL_UNKNOWN":  0,
	"DEAL_ACCEPTED": 1,
	"DEAL_CLOSED":   2,
}

func (x DealStatus) String() string {
	return proto.EnumName(DealStatus_name, int32(x))
}
func (DealStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor10, []int{3} }

type ChangeRequestStatus int32

const (
	ChangeRequestStatus_REQUEST_UNKNOWN  ChangeRequestStatus = 0
	ChangeRequestStatus_REQUEST_CREATED  ChangeRequestStatus = 1
	ChangeRequestStatus_REQUEST_CANCELED ChangeRequestStatus = 2
	ChangeRequestStatus_REQUEST_REJECTED ChangeRequestStatus = 3
	ChangeRequestStatus_REQUEST_ACCEPTED ChangeRequestStatus = 4
)

var ChangeRequestStatus_name = map[int32]string{
	0: "REQUEST_UNKNOWN",
	1: "REQUEST_CREATED",
	2: "REQUEST_CANCELED",
	3: "REQUEST_REJECTED",
	4: "REQUEST_ACCEPTED",
}
var ChangeRequestStatus_value = map[string]int32{
	"REQUEST_UNKNOWN":  0,
	"REQUEST_CREATED":  1,
	"REQUEST_CANCELED": 2,
	"REQUEST_REJECTED": 3,
	"REQUEST_ACCEPTED": 4,
}

func (x ChangeRequestStatus) String() string {
	return proto.EnumName(ChangeRequestStatus_name, int32(x))
}
func (ChangeRequestStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor10, []int{4} }

type GetOrdersReply struct {
	Orders []*Order `protobuf:"bytes,1,rep,name=orders" json:"orders,omitempty"`
}

func (m *GetOrdersReply) Reset()                    { *m = GetOrdersReply{} }
func (m *GetOrdersReply) String() string            { return proto.CompactTextString(m) }
func (*GetOrdersReply) ProtoMessage()               {}
func (*GetOrdersReply) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{0} }

func (m *GetOrdersReply) GetOrders() []*Order {
	if m != nil {
		return m.Orders
	}
	return nil
}

type Benchmarks struct {
	CPUSysbenchMulti uint64 `protobuf:"varint,1,opt,name=CPUSysbenchMulti" json:"CPUSysbenchMulti,omitempty"`
	CPUSysbenchOne   uint64 `protobuf:"varint,2,opt,name=CPUSysbenchOne" json:"CPUSysbenchOne,omitempty"`
	CPUCores         uint64 `protobuf:"varint,3,opt,name=CPUCores" json:"CPUCores,omitempty"`
	RAMSize          uint64 `protobuf:"varint,4,opt,name=RAMSize" json:"RAMSize,omitempty"`
	StorageSize      uint64 `protobuf:"varint,5,opt,name=StorageSize" json:"StorageSize,omitempty"`
	NetTrafficIn     uint64 `protobuf:"varint,6,opt,name=NetTrafficIn" json:"NetTrafficIn,omitempty"`
	NetTrafficOut    uint64 `protobuf:"varint,7,opt,name=NetTrafficOut" json:"NetTrafficOut,omitempty"`
	GPUCount         uint64 `protobuf:"varint,8,opt,name=GPUCount" json:"GPUCount,omitempty"`
	GPUMem           uint64 `protobuf:"varint,9,opt,name=GPUMem" json:"GPUMem,omitempty"`
	GPUEthHashrate   uint64 `protobuf:"varint,10,opt,name=GPUEthHashrate" json:"GPUEthHashrate,omitempty"`
	GPUCashHashrate  uint64 `protobuf:"varint,11,opt,name=GPUCashHashrate" json:"GPUCashHashrate,omitempty"`
	GPURedshift      uint64 `protobuf:"varint,12,opt,name=GPURedshift" json:"GPURedshift,omitempty"`
}

func (m *Benchmarks) Reset()                    { *m = Benchmarks{} }
func (m *Benchmarks) String() string            { return proto.CompactTextString(m) }
func (*Benchmarks) ProtoMessage()               {}
func (*Benchmarks) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{1} }

func (m *Benchmarks) GetCPUSysbenchMulti() uint64 {
	if m != nil {
		return m.CPUSysbenchMulti
	}
	return 0
}

func (m *Benchmarks) GetCPUSysbenchOne() uint64 {
	if m != nil {
		return m.CPUSysbenchOne
	}
	return 0
}

func (m *Benchmarks) GetCPUCores() uint64 {
	if m != nil {
		return m.CPUCores
	}
	return 0
}

func (m *Benchmarks) GetRAMSize() uint64 {
	if m != nil {
		return m.RAMSize
	}
	return 0
}

func (m *Benchmarks) GetStorageSize() uint64 {
	if m != nil {
		return m.StorageSize
	}
	return 0
}

func (m *Benchmarks) GetNetTrafficIn() uint64 {
	if m != nil {
		return m.NetTrafficIn
	}
	return 0
}

func (m *Benchmarks) GetNetTrafficOut() uint64 {
	if m != nil {
		return m.NetTrafficOut
	}
	return 0
}

func (m *Benchmarks) GetGPUCount() uint64 {
	if m != nil {
		return m.GPUCount
	}
	return 0
}

func (m *Benchmarks) GetGPUMem() uint64 {
	if m != nil {
		return m.GPUMem
	}
	return 0
}

func (m *Benchmarks) GetGPUEthHashrate() uint64 {
	if m != nil {
		return m.GPUEthHashrate
	}
	return 0
}

func (m *Benchmarks) GetGPUCashHashrate() uint64 {
	if m != nil {
		return m.GPUCashHashrate
	}
	return 0
}

func (m *Benchmarks) GetGPURedshift() uint64 {
	if m != nil {
		return m.GPURedshift
	}
	return 0
}

type Deal struct {
	Id             string      `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Benchmarks     *Benchmarks `protobuf:"bytes,2,opt,name=benchmarks" json:"benchmarks,omitempty"`
	SupplierID     string      `protobuf:"bytes,3,opt,name=supplierID" json:"supplierID,omitempty"`
	ConsumerID     string      `protobuf:"bytes,4,opt,name=consumerID" json:"consumerID,omitempty"`
	MasterID       string      `protobuf:"bytes,5,opt,name=masterID" json:"masterID,omitempty"`
	AskID          string      `protobuf:"bytes,6,opt,name=askID" json:"askID,omitempty"`
	BidID          string      `protobuf:"bytes,7,opt,name=bidID" json:"bidID,omitempty"`
	Duration       uint64      `protobuf:"varint,8,opt,name=duration" json:"duration,omitempty"`
	Price          *BigInt     `protobuf:"bytes,9,opt,name=price" json:"price,omitempty"`
	StartTime      *Timestamp  `protobuf:"bytes,10,opt,name=startTime" json:"startTime,omitempty"`
	EndTime        *Timestamp  `protobuf:"bytes,11,opt,name=endTime" json:"endTime,omitempty"`
	Status         DealStatus  `protobuf:"varint,12,opt,name=status,enum=sonm.DealStatus" json:"status,omitempty"`
	BlockedBalance *BigInt     `protobuf:"bytes,13,opt,name=blockedBalance" json:"blockedBalance,omitempty"`
	TotalPayout    *BigInt     `protobuf:"bytes,14,opt,name=totalPayout" json:"totalPayout,omitempty"`
	LastBillTS     *Timestamp  `protobuf:"bytes,15,opt,name=lastBillTS" json:"lastBillTS,omitempty"`
}

func (m *Deal) Reset()                    { *m = Deal{} }
func (m *Deal) String() string            { return proto.CompactTextString(m) }
func (*Deal) ProtoMessage()               {}
func (*Deal) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{2} }

func (m *Deal) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Deal) GetBenchmarks() *Benchmarks {
	if m != nil {
		return m.Benchmarks
	}
	return nil
}

func (m *Deal) GetSupplierID() string {
	if m != nil {
		return m.SupplierID
	}
	return ""
}

func (m *Deal) GetConsumerID() string {
	if m != nil {
		return m.ConsumerID
	}
	return ""
}

func (m *Deal) GetMasterID() string {
	if m != nil {
		return m.MasterID
	}
	return ""
}

func (m *Deal) GetAskID() string {
	if m != nil {
		return m.AskID
	}
	return ""
}

func (m *Deal) GetBidID() string {
	if m != nil {
		return m.BidID
	}
	return ""
}

func (m *Deal) GetDuration() uint64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Deal) GetPrice() *BigInt {
	if m != nil {
		return m.Price
	}
	return nil
}

func (m *Deal) GetStartTime() *Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *Deal) GetEndTime() *Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *Deal) GetStatus() DealStatus {
	if m != nil {
		return m.Status
	}
	return DealStatus_DEAL_UNKNOWN
}

func (m *Deal) GetBlockedBalance() *BigInt {
	if m != nil {
		return m.BlockedBalance
	}
	return nil
}

func (m *Deal) GetTotalPayout() *BigInt {
	if m != nil {
		return m.TotalPayout
	}
	return nil
}

func (m *Deal) GetLastBillTS() *Timestamp {
	if m != nil {
		return m.LastBillTS
	}
	return nil
}

type Order struct {
	Id             string        `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	DealID         string        `protobuf:"bytes,2,opt,name=dealID" json:"dealID,omitempty"`
	OrderType      OrderType     `protobuf:"varint,3,opt,name=orderType,enum=sonm.OrderType" json:"orderType,omitempty"`
	OrderStatus    OrderStatus   `protobuf:"varint,4,opt,name=orderStatus,enum=sonm.OrderStatus" json:"orderStatus,omitempty"`
	AuthorID       string        `protobuf:"bytes,5,opt,name=authorID" json:"authorID,omitempty"`
	CounterpartyID string        `protobuf:"bytes,6,opt,name=counterpartyID" json:"counterpartyID,omitempty"`
	Duration       uint64        `protobuf:"varint,7,opt,name=duration" json:"duration,omitempty"`
	Price          *BigInt       `protobuf:"bytes,8,opt,name=price" json:"price,omitempty"`
	Netflags       uint64        `protobuf:"varint,9,opt,name=netflags" json:"netflags,omitempty"`
	IdentityLevel  IdentityLevel `protobuf:"varint,10,opt,name=identityLevel,enum=sonm.IdentityLevel" json:"identityLevel,omitempty"`
	Blacklist      string        `protobuf:"bytes,11,opt,name=blacklist" json:"blacklist,omitempty"`
	Tag            []byte        `protobuf:"bytes,12,opt,name=tag,proto3" json:"tag,omitempty"`
	Benchmarks     *Benchmarks   `protobuf:"bytes,13,opt,name=benchmarks" json:"benchmarks,omitempty"`
	FrozenSum      *BigInt       `protobuf:"bytes,14,opt,name=frozenSum" json:"frozenSum,omitempty"`
}

func (m *Order) Reset()                    { *m = Order{} }
func (m *Order) String() string            { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()               {}
func (*Order) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{3} }

func (m *Order) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Order) GetDealID() string {
	if m != nil {
		return m.DealID
	}
	return ""
}

func (m *Order) GetOrderType() OrderType {
	if m != nil {
		return m.OrderType
	}
	return OrderType_ANY
}

func (m *Order) GetOrderStatus() OrderStatus {
	if m != nil {
		return m.OrderStatus
	}
	return OrderStatus_ORDER_INACTIVE
}

func (m *Order) GetAuthorID() string {
	if m != nil {
		return m.AuthorID
	}
	return ""
}

func (m *Order) GetCounterpartyID() string {
	if m != nil {
		return m.CounterpartyID
	}
	return ""
}

func (m *Order) GetDuration() uint64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Order) GetPrice() *BigInt {
	if m != nil {
		return m.Price
	}
	return nil
}

func (m *Order) GetNetflags() uint64 {
	if m != nil {
		return m.Netflags
	}
	return 0
}

func (m *Order) GetIdentityLevel() IdentityLevel {
	if m != nil {
		return m.IdentityLevel
	}
	return IdentityLevel_ANONIMOUS
}

func (m *Order) GetBlacklist() string {
	if m != nil {
		return m.Blacklist
	}
	return ""
}

func (m *Order) GetTag() []byte {
	if m != nil {
		return m.Tag
	}
	return nil
}

func (m *Order) GetBenchmarks() *Benchmarks {
	if m != nil {
		return m.Benchmarks
	}
	return nil
}

func (m *Order) GetFrozenSum() *BigInt {
	if m != nil {
		return m.FrozenSum
	}
	return nil
}

func init() {
	proto.RegisterType((*GetOrdersReply)(nil), "sonm.GetOrdersReply")
	proto.RegisterType((*Benchmarks)(nil), "sonm.Benchmarks")
	proto.RegisterType((*Deal)(nil), "sonm.Deal")
	proto.RegisterType((*Order)(nil), "sonm.Order")
	proto.RegisterEnum("sonm.OrderType", OrderType_name, OrderType_value)
	proto.RegisterEnum("sonm.OrderStatus", OrderStatus_name, OrderStatus_value)
	proto.RegisterEnum("sonm.IdentityLevel", IdentityLevel_name, IdentityLevel_value)
	proto.RegisterEnum("sonm.DealStatus", DealStatus_name, DealStatus_value)
	proto.RegisterEnum("sonm.ChangeRequestStatus", ChangeRequestStatus_name, ChangeRequestStatus_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Market service

type MarketClient interface {
	// GetOrders returns orders by given filter parameters.
	// Note that set of filters may be changed in the closest future.
	GetOrders(ctx context.Context, in *Count, opts ...grpc.CallOption) (*GetOrdersReply, error)
	// CreateOrder places new order on the Marketplace.
	// Note that current impl of Node API prevents you from
	// creating ASKs orders.
	CreateOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Order, error)
	// GetOrderByID returns order by given ID.
	// If order save an `inactive` status returns error instead.
	GetOrderByID(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Order, error)
	// CancelOrder removes active order from the Marketplace.
	CancelOrder(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Empty, error)
}

type marketClient struct {
	cc *grpc.ClientConn
}

func NewMarketClient(cc *grpc.ClientConn) MarketClient {
	return &marketClient{cc}
}

func (c *marketClient) GetOrders(ctx context.Context, in *Count, opts ...grpc.CallOption) (*GetOrdersReply, error) {
	out := new(GetOrdersReply)
	err := grpc.Invoke(ctx, "/sonm.Market/GetOrders", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) CreateOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := grpc.Invoke(ctx, "/sonm.Market/CreateOrder", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) GetOrderByID(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Order, error) {
	out := new(Order)
	err := grpc.Invoke(ctx, "/sonm.Market/GetOrderByID", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketClient) CancelOrder(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/sonm.Market/CancelOrder", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Market service

type MarketServer interface {
	// GetOrders returns orders by given filter parameters.
	// Note that set of filters may be changed in the closest future.
	GetOrders(context.Context, *Count) (*GetOrdersReply, error)
	// CreateOrder places new order on the Marketplace.
	// Note that current impl of Node API prevents you from
	// creating ASKs orders.
	CreateOrder(context.Context, *Order) (*Order, error)
	// GetOrderByID returns order by given ID.
	// If order save an `inactive` status returns error instead.
	GetOrderByID(context.Context, *ID) (*Order, error)
	// CancelOrder removes active order from the Marketplace.
	CancelOrder(context.Context, *ID) (*Empty, error)
}

func RegisterMarketServer(s *grpc.Server, srv MarketServer) {
	s.RegisterService(&_Market_serviceDesc, srv)
}

func _Market_GetOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Count)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).GetOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Market/GetOrders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).GetOrders(ctx, req.(*Count))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Order)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Market/CreateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).CreateOrder(ctx, req.(*Order))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_GetOrderByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).GetOrderByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Market/GetOrderByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).GetOrderByID(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Market_CancelOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketServer).CancelOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sonm.Market/CancelOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketServer).CancelOrder(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

var _Market_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sonm.Market",
	HandlerType: (*MarketServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOrders",
			Handler:    _Market_GetOrders_Handler,
		},
		{
			MethodName: "CreateOrder",
			Handler:    _Market_CreateOrder_Handler,
		},
		{
			MethodName: "GetOrderByID",
			Handler:    _Market_GetOrderByID_Handler,
		},
		{
			MethodName: "CancelOrder",
			Handler:    _Market_CancelOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "marketplace.proto",
}

// Begin grpccmd
var _ = grpccmd.RunE

// Market
var _MarketCmd = &cobra.Command{
	Use:   "market [method]",
	Short: "Subcommand for the Market service.",
}

var _Market_GetOrdersCmd = &cobra.Command{
	Use:   "getOrders",
	Short: "Make the GetOrders method call, input-type: sonm.Count output-type: sonm.GetOrdersReply",
	RunE: grpccmd.RunE(
		"GetOrders",
		"sonm.Count",
		func(c io.Closer) interface{} {
			cc := c.(*grpc.ClientConn)
			return NewMarketClient(cc)
		},
	),
}

var _Market_GetOrdersCmd_gen = &cobra.Command{
	Use:   "getOrders-gen",
	Short: "Generate JSON for method call of GetOrders (input-type: sonm.Count)",
	RunE:  grpccmd.TypeToJson("sonm.Count"),
}

var _Market_CreateOrderCmd = &cobra.Command{
	Use:   "createOrder",
	Short: "Make the CreateOrder method call, input-type: sonm.Order output-type: sonm.Order",
	RunE: grpccmd.RunE(
		"CreateOrder",
		"sonm.Order",
		func(c io.Closer) interface{} {
			cc := c.(*grpc.ClientConn)
			return NewMarketClient(cc)
		},
	),
}

var _Market_CreateOrderCmd_gen = &cobra.Command{
	Use:   "createOrder-gen",
	Short: "Generate JSON for method call of CreateOrder (input-type: sonm.Order)",
	RunE:  grpccmd.TypeToJson("sonm.Order"),
}

var _Market_GetOrderByIDCmd = &cobra.Command{
	Use:   "getOrderByID",
	Short: "Make the GetOrderByID method call, input-type: sonm.ID output-type: sonm.Order",
	RunE: grpccmd.RunE(
		"GetOrderByID",
		"sonm.ID",
		func(c io.Closer) interface{} {
			cc := c.(*grpc.ClientConn)
			return NewMarketClient(cc)
		},
	),
}

var _Market_GetOrderByIDCmd_gen = &cobra.Command{
	Use:   "getOrderByID-gen",
	Short: "Generate JSON for method call of GetOrderByID (input-type: sonm.ID)",
	RunE:  grpccmd.TypeToJson("sonm.ID"),
}

var _Market_CancelOrderCmd = &cobra.Command{
	Use:   "cancelOrder",
	Short: "Make the CancelOrder method call, input-type: sonm.ID output-type: sonm.Empty",
	RunE: grpccmd.RunE(
		"CancelOrder",
		"sonm.ID",
		func(c io.Closer) interface{} {
			cc := c.(*grpc.ClientConn)
			return NewMarketClient(cc)
		},
	),
}

var _Market_CancelOrderCmd_gen = &cobra.Command{
	Use:   "cancelOrder-gen",
	Short: "Generate JSON for method call of CancelOrder (input-type: sonm.ID)",
	RunE:  grpccmd.TypeToJson("sonm.ID"),
}

// Register commands with the root command and service command
func init() {
	grpccmd.RegisterServiceCmd(_MarketCmd)
	_MarketCmd.AddCommand(
		_Market_GetOrdersCmd,
		_Market_GetOrdersCmd_gen,
		_Market_CreateOrderCmd,
		_Market_CreateOrderCmd_gen,
		_Market_GetOrderByIDCmd,
		_Market_GetOrderByIDCmd_gen,
		_Market_CancelOrderCmd,
		_Market_CancelOrderCmd_gen,
	)
}

// End grpccmd

func init() { proto.RegisterFile("marketplace.proto", fileDescriptor10) }

var fileDescriptor10 = []byte{
	// 1018 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x96, 0xcf, 0x6f, 0xe2, 0x46,
	0x14, 0xc7, 0x21, 0x10, 0x12, 0x3f, 0x83, 0xf1, 0x4e, 0xa2, 0x95, 0x85, 0xaa, 0x2a, 0xa2, 0x55,
	0x9a, 0x45, 0x6a, 0xba, 0x22, 0xed, 0xa1, 0xb7, 0x82, 0xed, 0xa6, 0xee, 0x26, 0x86, 0x8e, 0xa1,
	0xd5, 0x9e, 0x56, 0x03, 0x4c, 0x60, 0x14, 0x63, 0xbb, 0xf6, 0xb8, 0x12, 0x7b, 0xeb, 0xbf, 0xd3,
	0xde, 0xfb, 0xd7, 0xf5, 0x50, 0xcd, 0xd8, 0x06, 0x43, 0x58, 0xa9, 0x37, 0xde, 0xe7, 0x7d, 0xe7,
	0xe5, 0xcd, 0xfb, 0xe1, 0x09, 0xbc, 0x5a, 0x93, 0xf8, 0x99, 0xf2, 0xc8, 0x27, 0x73, 0x7a, 0x1b,
	0xc5, 0x21, 0x0f, 0x51, 0x3d, 0x09, 0x83, 0x75, 0xa7, 0x39, 0x63, 0x4b, 0x16, 0xf0, 0x8c, 0x75,
	0xda, 0x2c, 0x10, 0x34, 0x60, 0xa4, 0x00, 0x9c, 0xad, 0x69, 0xc2, 0xc9, 0x3a, 0xca, 0x40, 0xf7,
	0x3b, 0xd0, 0xee, 0x29, 0x1f, 0xc5, 0x0b, 0x1a, 0x27, 0x98, 0x46, 0xfe, 0x06, 0x7d, 0x01, 0x8d,
	0x50, 0x9a, 0x46, 0xf5, 0xaa, 0x76, 0xa3, 0xf6, 0xd5, 0x5b, 0x11, 0xe2, 0x56, 0x4a, 0x70, 0xee,
	0xea, 0xfe, 0x55, 0x03, 0x18, 0xd2, 0x60, 0xbe, 0x12, 0x79, 0x24, 0xa8, 0x07, 0xba, 0x39, 0x9e,
	0x7a, 0x9b, 0x64, 0x26, 0xd8, 0x63, 0xea, 0x73, 0x66, 0x54, 0xaf, 0xaa, 0x37, 0x75, 0xfc, 0x82,
	0xa3, 0x6b, 0xd0, 0x4a, 0x6c, 0x14, 0x50, 0xe3, 0x44, 0x2a, 0x0f, 0x28, 0xea, 0xc0, 0xb9, 0x39,
	0x9e, 0x9a, 0x61, 0x4c, 0x13, 0xa3, 0x26, 0x15, 0x5b, 0x1b, 0x19, 0x70, 0x86, 0x07, 0x8f, 0x1e,
	0xfb, 0x48, 0x8d, 0xba, 0x74, 0x15, 0x26, 0xba, 0x02, 0xd5, 0xe3, 0x61, 0x4c, 0x96, 0x54, 0x7a,
	0x4f, 0xa5, 0xb7, 0x8c, 0x50, 0x17, 0x9a, 0x2e, 0xe5, 0x93, 0x98, 0x3c, 0x3d, 0xb1, 0xb9, 0x13,
	0x18, 0x0d, 0x29, 0xd9, 0x63, 0xe8, 0x4b, 0x68, 0xed, 0xec, 0x51, 0xca, 0x8d, 0x33, 0x29, 0xda,
	0x87, 0x22, 0xc3, 0x7b, 0x91, 0x51, 0x1a, 0x70, 0xe3, 0x3c, 0xcb, 0xb0, 0xb0, 0xd1, 0x6b, 0x68,
	0xdc, 0x8f, 0xa7, 0x8f, 0x74, 0x6d, 0x28, 0xd2, 0x93, 0x5b, 0xe2, 0xf6, 0xf7, 0xe3, 0xa9, 0xcd,
	0x57, 0x3f, 0x91, 0x64, 0x15, 0x13, 0x4e, 0x0d, 0xc8, 0x6e, 0xbf, 0x4f, 0xd1, 0x0d, 0xb4, 0x45,
	0x2c, 0x92, 0xec, 0x84, 0xaa, 0x14, 0x1e, 0x62, 0x71, 0xe3, 0xfb, 0xf1, 0x14, 0xd3, 0x45, 0xb2,
	0x62, 0x4f, 0xdc, 0x68, 0x66, 0x37, 0x2e, 0xa1, 0xee, 0xdf, 0x75, 0xa8, 0x5b, 0x94, 0xf8, 0x48,
	0x83, 0x13, 0xb6, 0x90, 0x8d, 0x51, 0xf0, 0x09, 0x5b, 0xa0, 0xb7, 0x00, 0xb3, 0x6d, 0x13, 0x65,
	0x1b, 0xd4, 0xbe, 0x9e, 0xb5, 0x7b, 0xd7, 0x5c, 0x5c, 0xd2, 0xa0, 0xcf, 0x01, 0x92, 0x34, 0x8a,
	0x7c, 0x46, 0x63, 0xc7, 0x92, 0x6d, 0x51, 0x70, 0x89, 0x08, 0xff, 0x3c, 0x0c, 0x92, 0x74, 0x2d,
	0xfd, 0xf5, 0xcc, 0xbf, 0x23, 0xa2, 0x64, 0x6b, 0x92, 0x70, 0xe9, 0x3d, 0x95, 0xde, 0xad, 0x8d,
	0x2e, 0xe1, 0x94, 0x24, 0xcf, 0x8e, 0x25, 0x3b, 0xa2, 0xe0, 0xcc, 0x10, 0x74, 0xc6, 0x16, 0x8e,
	0x25, 0x5b, 0xa0, 0xe0, 0xcc, 0x10, 0x71, 0x16, 0x69, 0x4c, 0x38, 0x0b, 0x83, 0xa2, 0xf4, 0x85,
	0x8d, 0xba, 0x70, 0x1a, 0xc5, 0x6c, 0x4e, 0x65, 0xe5, 0xd5, 0x7e, 0x33, 0xbf, 0x10, 0x5b, 0x3a,
	0x01, 0xc7, 0x99, 0x0b, 0x7d, 0x0d, 0x4a, 0xc2, 0x49, 0xcc, 0x27, 0x6c, 0x9d, 0x75, 0x40, 0xed,
	0xb7, 0x33, 0xdd, 0xa4, 0x58, 0x10, 0xbc, 0x53, 0xa0, 0x37, 0x70, 0x46, 0x83, 0x85, 0x14, 0xab,
	0xc7, 0xc5, 0x85, 0x1f, 0xdd, 0x40, 0x23, 0xe1, 0x84, 0xa7, 0x89, 0xec, 0x84, 0x56, 0xd4, 0x53,
	0xd4, 0xdf, 0x93, 0x1c, 0xe7, 0x7e, 0xf4, 0x2d, 0x68, 0x33, 0x3f, 0x9c, 0x3f, 0xd3, 0xc5, 0x90,
	0xf8, 0x24, 0x98, 0x53, 0xa3, 0x75, 0x24, 0xe1, 0x03, 0x0d, 0xba, 0x05, 0x95, 0x87, 0x9c, 0xf8,
	0x63, 0xb2, 0x09, 0x53, 0x6e, 0x68, 0x47, 0x8e, 0x94, 0x05, 0xe8, 0x1b, 0x00, 0x9f, 0x24, 0x7c,
	0xc8, 0x7c, 0x7f, 0xe2, 0x19, 0xed, 0xe3, 0xd9, 0x97, 0x24, 0xdd, 0x7f, 0x6b, 0x70, 0x2a, 0x97,
	0xfd, 0xc5, 0xb8, 0xbc, 0x86, 0xc6, 0x82, 0x12, 0xdf, 0xb1, 0xe4, 0xa8, 0x28, 0x38, 0xb7, 0x44,
	0x31, 0xe5, 0x67, 0x61, 0xb2, 0x89, 0xa8, 0x9c, 0x09, 0xad, 0xf8, 0x0b, 0xa3, 0x02, 0xe3, 0x9d,
	0x02, 0xdd, 0x81, 0x2a, 0x8d, 0xac, 0x1c, 0x72, 0x48, 0xb4, 0xfe, 0xab, 0xd2, 0x81, 0xbc, 0x4e,
	0x65, 0x95, 0x68, 0x38, 0x49, 0xf9, 0x2a, 0x2c, 0x0d, 0x4e, 0x61, 0x8b, 0x9d, 0x9a, 0x8b, 0xa5,
	0xa3, 0x71, 0x44, 0x62, 0xbe, 0xd9, 0x4e, 0xd0, 0x01, 0xdd, 0x1b, 0x9a, 0xb3, 0x4f, 0x0d, 0xcd,
	0xf9, 0xa7, 0x87, 0xa6, 0x03, 0xe7, 0x01, 0xe5, 0x4f, 0x3e, 0x59, 0x26, 0xf9, 0x56, 0x6f, 0x6d,
	0xf4, 0x3d, 0xb4, 0xd8, 0x82, 0x06, 0x9c, 0xf1, 0xcd, 0x03, 0xfd, 0x83, 0xfa, 0x72, 0xa8, 0xb4,
	0xfe, 0x45, 0x16, 0xc7, 0x29, 0xbb, 0xf0, 0xbe, 0x12, 0x7d, 0x06, 0xca, 0xcc, 0x27, 0xf3, 0x67,
	0x9f, 0x25, 0x5c, 0x8e, 0x97, 0x82, 0x77, 0x00, 0xe9, 0x50, 0xe3, 0x64, 0x29, 0x87, 0xa9, 0x89,
	0xc5, 0xcf, 0x83, 0xad, 0x6d, 0xfd, 0x8f, 0xad, 0xed, 0x81, 0xf2, 0x14, 0x87, 0x1f, 0x69, 0xe0,
	0xa5, 0xeb, 0xa3, 0x13, 0xb3, 0x73, 0xf7, 0xae, 0x41, 0xd9, 0x76, 0x0d, 0x9d, 0x41, 0x6d, 0xe0,
	0xbe, 0xd7, 0x2b, 0xe2, 0xc7, 0xd0, 0xb1, 0xf4, 0xaa, 0x24, 0xde, 0x3b, 0xfd, 0xa4, 0x77, 0x07,
	0x6a, 0xa9, 0x59, 0x08, 0x81, 0x36, 0xc2, 0x96, 0x8d, 0x3f, 0x38, 0xee, 0xc0, 0x9c, 0x38, 0xbf,
	0xda, 0x7a, 0x05, 0xe9, 0xd0, 0xcc, 0x58, 0x4e, 0xaa, 0xbd, 0x1f, 0xa0, 0xb5, 0x57, 0x0a, 0xd4,
	0x02, 0x65, 0xe0, 0x8e, 0x5c, 0xe7, 0x71, 0x34, 0xf5, 0xb2, 0x13, 0x63, 0xcf, 0x9e, 0x5a, 0x23,
	0xf7, 0xbd, 0x24, 0x55, 0xa4, 0x01, 0x38, 0x96, 0xed, 0x4e, 0x9c, 0x1f, 0x1d, 0xdb, 0xd2, 0x4f,
	0x7a, 0x43, 0x80, 0xdd, 0x2a, 0x09, 0xbd, 0x65, 0x0f, 0x1e, 0x3e, 0x4c, 0xdd, 0x77, 0xee, 0xe8,
	0x37, 0x57, 0xaf, 0xa0, 0x57, 0xd0, 0x92, 0x64, 0x60, 0x9a, 0xf6, 0x78, 0x62, 0x8b, 0x94, 0xdb,
	0xa0, 0x4a, 0x64, 0x3e, 0x8c, 0x3c, 0x19, 0xe3, 0xcf, 0x2a, 0x5c, 0x98, 0x2b, 0x12, 0x2c, 0x29,
	0xa6, 0xbf, 0xa7, 0x34, 0xe1, 0x79, 0xb4, 0x0b, 0x68, 0x63, 0xfb, 0x97, 0xa9, 0xed, 0x4d, 0x4a,
	0x01, 0x4b, 0xd0, 0xc4, 0xf6, 0x20, 0x0b, 0x79, 0x09, 0xfa, 0x16, 0x0e, 0x5c, 0xd3, 0x7e, 0x10,
	0x71, 0xcb, 0x14, 0xdb, 0x3f, 0xdb, 0xa6, 0xd0, 0xd6, 0xca, 0x74, 0x9b, 0x54, 0xbd, 0xff, 0x4f,
	0x15, 0x1a, 0x8f, 0xf2, 0x0d, 0x47, 0x6f, 0x41, 0xd9, 0x3e, 0xc1, 0x28, 0x7f, 0x6d, 0xe5, 0x23,
	0xd2, 0xb9, 0xcc, 0x8c, 0xfd, 0x07, 0xba, 0x5b, 0x41, 0x6f, 0x40, 0x35, 0x63, 0x4a, 0x38, 0xcd,
	0xf6, 0xb4, 0xfc, 0x42, 0x77, 0xca, 0x46, 0xb7, 0x82, 0xbe, 0x82, 0x66, 0x71, 0x7c, 0x28, 0x76,
	0xe0, 0x3c, 0x1f, 0x48, 0xeb, 0x50, 0x78, 0x0d, 0xaa, 0x29, 0x3e, 0x30, 0x7e, 0x16, 0xf3, 0x85,
	0xce, 0x5e, 0x47, 0x7c, 0xd3, 0xad, 0xcc, 0x1a, 0xf2, 0xff, 0x86, 0xbb, 0xff, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x0f, 0xde, 0xda, 0xad, 0x82, 0x08, 0x00, 0x00,
}
