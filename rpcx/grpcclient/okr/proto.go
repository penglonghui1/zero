// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by grpc_generate_tools(grpcgengo) at
package okr

import (
	"context"
	"github.com/pengcainiao2/zero/rpcx/grpcbase"

	jsoniter "github.com/json-iterator/go"
)

var (
	JoinedTasksHandler func(ctx context.Context, req JoinedTasksRequest) grpcbase.Response
)

type Repository interface {
	JoinedTasks(ctx context.Context, request JoinedTasksRequest) grpcbase.Response
}

type Any struct {
	type_url string `json:"type_url,omitempty"`
	value    []byte `json:"value,omitempty"`
}

func (s Any) String() string {
	b, _ := jsoniter.Marshal(s)
	return string(b)
}

type JoinedTasksRequest struct {
	Keyword string       `json:"keyword,omitempty"`
	Paging  *Paging      `json:"paging,omitempty"`
	Context *UserContext `json:"context,omitempty"`
	Type    []string     `json:"type,omitempty"`
}

func (s JoinedTasksRequest) String() string {
	b, _ := jsoniter.Marshal(s)
	return string(b)
}

type JoinedTasksResponse struct {
	RefID []string `json:"ref_id,omitempty"`
}

func (s JoinedTasksResponse) String() string {
	b, _ := jsoniter.Marshal(s)
	return string(b)
}

type Paging struct {
	PageNumber int32 `json:"page_number,omitempty"`
	PageRecord int32 `json:"page_record,omitempty"`
}

func (s Paging) String() string {
	b, _ := jsoniter.Marshal(s)
	return string(b)
}

type Response struct {
	Message string `json:"message,omitempty"`
	Total   int32  `json:"total,omitempty"`
	Data    *Any   `json:"data,omitempty"`
}

func (s Response) String() string {
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

// NewService 新建datasyncer的grpc服务
func NewService() Repository {
	return service{}
}

func (s service) JoinedTasks(ctx context.Context, req JoinedTasksRequest) grpcbase.Response {
	if JoinedTasksHandler != nil {
		if req.Context == nil {
			panic("grpc context is nil and requestID must be set")
		}
		return grpcbase.RPCServerSideLogic("/pb.okr/JoinedTasks", req.Context.RequestID, req, func(ctx context.Context, req interface{}) grpcbase.Response {
			return JoinedTasksHandler(ctx, req.(JoinedTasksRequest))
		})
	}
	return grpcbase.NotImplErrorResponse
}