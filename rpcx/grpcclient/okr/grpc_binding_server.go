// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by grpc_generate_tools(grpcgengo) at
//

package okr

import (
	"context"
	"log"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/pengcainiao2/zero/rpcx/grpcbase"
	pb "github.com/pengcainiao2/zero/rpcx/protos"
	"google.golang.org/grpc"
)

type serverBinding struct {
	pb.UnimplementedOkrServer
	getOkr grpctransport.Handler
}

func (b *serverBinding) GetOkr(ctx context.Context, req *pb.GetOkrRequest) (*pb.Response, error) {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：GetOkr receive request context is nil，trace span将无法生效")
	}
	_, response, err := b.getOkr.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return response.(*pb.Response), nil
}

func (b *serverBinding) RegisterServer(srv *grpc.Server) error {
	pb.RegisterOkrServer(srv, b)
	return nil
}

func (b *serverBinding) GRPCHandler() map[string]grpctransport.Handler {
	return map[string]grpctransport.Handler{
		"getOkr": b.getOkr,
	}
}

func NewBinding(svc Repository) *serverBinding {
	return &serverBinding{
		getOkr: grpcbase.CreateGRPCServer(
			makeGetOkrEndpoint(svc),
			decodeGetOkrRequest,
			encodeGetOkrResponse,
		),
	}
}
