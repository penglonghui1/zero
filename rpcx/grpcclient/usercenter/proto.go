// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by grpc_generate_tools(grpcgengo) at
package usercenter

import (
	"context"

	jsoniter "github.com/json-iterator/go"
	"github.com/pengcainiao2/zero/rpcx/grpcbase"
)

var (
	GetUserHandler func(ctx context.Context, req GetUserRequest) grpcbase.Response
)

type Repository interface {
	GetUser(ctx context.Context, request GetUserRequest) grpcbase.Response
}

type GetUserRequest struct {
	Keyword string       `json:"keyword,omitempty"`
	Context *UserContext `json:"context,omitempty"`
}

func (s GetUserRequest) String() string {
	b, _ := jsoniter.Marshal(s)
	return string(b)
}

type GetUserResponse struct {
	Name string `json:"name,omitempty"`
}

func (s GetUserResponse) String() string {
	b, _ := jsoniter.Marshal(s)
	return string(b)
}

type UserContext struct {
	UserID        string `json:"user_id,omitempty"`
	Platform      string `json:"platform,omitempty"`
	ClientVersion string `json:"client_version,omitempty"`
	Token         string `json:"token,omitempty"`
	ClientIP      string `json:"client_ip,omitempty"`
	RequestID     string `json:"request_id,omitempty"`
}

func (s UserContext) String() string {
	b, _ := jsoniter.Marshal(s)
	return string(b)
}

type service struct{}

// RpcContextFromHeader 从 httprouter.Context转换为grpc中所需的用户上下文
func RpcContextFromHeader(header string) *UserContext {
	var ctx *UserContext
	_ = jsoniter.UnmarshalFromString(header, ctx)
	return ctx
}

// NewService 新建usercenter的grpc服务
func NewService() Repository {
	return service{}
}

func (s service) GetUser(ctx context.Context, req GetUserRequest) grpcbase.Response {
	if GetUserHandler != nil {
		if req.Context == nil {
			panic("grpc context is nil and requestID must be set")
		}
		return grpcbase.RPCServerSideLogic("/pb.usercenter/GetUser", req.Context.RequestID, req, func(ctx context.Context, req interface{}) grpcbase.Response {
			return GetUserHandler(ctx, req.(GetUserRequest))
		})
	}
	return grpcbase.NotImplErrorResponse
}