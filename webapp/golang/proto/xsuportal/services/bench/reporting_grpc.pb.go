// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package bench

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// BenchmarkReportClient is the client API for BenchmarkReport service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BenchmarkReportClient interface {
	ReportBenchmarkResult(ctx context.Context, opts ...grpc.CallOption) (BenchmarkReport_ReportBenchmarkResultClient, error)
}

type benchmarkReportClient struct {
	cc grpc.ClientConnInterface
}

func NewBenchmarkReportClient(cc grpc.ClientConnInterface) BenchmarkReportClient {
	return &benchmarkReportClient{cc}
}

var benchmarkReportReportBenchmarkResultStreamDesc = &grpc.StreamDesc{
	StreamName:    "ReportBenchmarkResult",
	ServerStreams: true,
	ClientStreams: true,
}

func (c *benchmarkReportClient) ReportBenchmarkResult(ctx context.Context, opts ...grpc.CallOption) (BenchmarkReport_ReportBenchmarkResultClient, error) {
	stream, err := c.cc.NewStream(ctx, benchmarkReportReportBenchmarkResultStreamDesc, "/xsuportal.proto.services.bench.BenchmarkReport/ReportBenchmarkResult", opts...)
	if err != nil {
		return nil, err
	}
	x := &benchmarkReportReportBenchmarkResultClient{stream}
	return x, nil
}

type BenchmarkReport_ReportBenchmarkResultClient interface {
	Send(*ReportBenchmarkResultRequest) error
	Recv() (*ReportBenchmarkResultResponse, error)
	grpc.ClientStream
}

type benchmarkReportReportBenchmarkResultClient struct {
	grpc.ClientStream
}

func (x *benchmarkReportReportBenchmarkResultClient) Send(m *ReportBenchmarkResultRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *benchmarkReportReportBenchmarkResultClient) Recv() (*ReportBenchmarkResultResponse, error) {
	m := new(ReportBenchmarkResultResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BenchmarkReportService is the service API for BenchmarkReport service.
// Fields should be assigned to their respective handler implementations only before
// RegisterBenchmarkReportService is called.  Any unassigned fields will result in the
// handler for that method returning an Unimplemented error.
type BenchmarkReportService struct {
	ReportBenchmarkResult func(BenchmarkReport_ReportBenchmarkResultServer) error
}

func (s *BenchmarkReportService) reportBenchmarkResult(_ interface{}, stream grpc.ServerStream) error {
	return s.ReportBenchmarkResult(&benchmarkReportReportBenchmarkResultServer{stream})
}

type BenchmarkReport_ReportBenchmarkResultServer interface {
	Send(*ReportBenchmarkResultResponse) error
	Recv() (*ReportBenchmarkResultRequest, error)
	grpc.ServerStream
}

type benchmarkReportReportBenchmarkResultServer struct {
	grpc.ServerStream
}

func (x *benchmarkReportReportBenchmarkResultServer) Send(m *ReportBenchmarkResultResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *benchmarkReportReportBenchmarkResultServer) Recv() (*ReportBenchmarkResultRequest, error) {
	m := new(ReportBenchmarkResultRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RegisterBenchmarkReportService registers a service implementation with a gRPC server.
func RegisterBenchmarkReportService(s grpc.ServiceRegistrar, srv *BenchmarkReportService) {
	srvCopy := *srv
	if srvCopy.ReportBenchmarkResult == nil {
		srvCopy.ReportBenchmarkResult = func(BenchmarkReport_ReportBenchmarkResultServer) error {
			return status.Errorf(codes.Unimplemented, "method ReportBenchmarkResult not implemented")
		}
	}
	sd := grpc.ServiceDesc{
		ServiceName: "xsuportal.proto.services.bench.BenchmarkReport",
		Methods:     []grpc.MethodDesc{},
		Streams: []grpc.StreamDesc{
			{
				StreamName:    "ReportBenchmarkResult",
				Handler:       srvCopy.reportBenchmarkResult,
				ServerStreams: true,
				ClientStreams: true,
			},
		},
		Metadata: "xsuportal/services/bench/reporting.proto",
	}

	s.RegisterService(&sd, nil)
}
