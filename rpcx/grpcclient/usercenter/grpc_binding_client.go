// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by grpc_generate_tools(grpcgengo) at
//

package usercenter

import (
	"context"
	"log"

	"github.com/go-kit/kit/endpoint"
	"github.com/pengcainiao2/zero/rpcx/grpcbase"
	"google.golang.org/grpc"
)

func init() {
	grpcbase.RegisterClients(grpcbase.ServerAddr(grpcbase.UserCenterSVC), &clientBinding{})
}

type clientBinding struct {
	getUser endpoint.Endpoint
}

func (c *clientBinding) GetUser(ctx context.Context, params GetUserRequest) grpcbase.Response {
	if ctx == nil {
		ctx = context.Background()
		log.Println("GRPC：GetUser request context is nil，trace span将无法生效")
	}
	response, err := c.getUser(ctx, params)
	if err != nil {
		return grpcbase.Response{
			Message: err.Error(),
		}
	}
	r := response.(grpcbase.Response)
	return r
}

func (c *clientBinding) GRPCClient(cc *grpc.ClientConn) interface{} {
	c.newClient(cc)
	return c
}

func (c *clientBinding) newClient(cc *grpc.ClientConn) {

	c.getUser = grpcbase.CreateGRPCClientEndpoint(cc, "pb.UserCenter",
		"GetUser",
		encodeGetUserRequest,
		decodeGetUserResponse)
}
