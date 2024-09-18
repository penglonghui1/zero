package test

//
//import (
//	"context"
//	"log"
//	"testing"
//	"time"
//
//	"github.com/pengcainiao2/zero/rpcx/grpcclient/operationsystem"
//
//	"github.com/pengcainiao2/zero/rpcx/grpcbase"
//	"github.com/pengcainiao2/zero/rpcx/grpcclient/clouddisk"
//	"github.com/pengcainiao2/zero/rpcx/grpcclient/notice"
//	"github.com/pengcainiao2/zero/rpcx/grpcclient/record"
//	"github.com/pengcainiao2/zero/rpcx/grpcclient/task"
//	"github.com/pengcainiao2/zero/rpcx/grpcclient/usercenter"
//)
//
//type RpcServer struct {
//	impl interface{}
//}
//
//func TestRpc(t *testing.T) {
//	var m = []int{1, 2, 3, 4, 5, 6, 7}
//	var x = make([]int, 0)
//	x = append(x, m[:3]...)
//	x = append(x, m[4:]...)
//	printResponse(UserCenter().HelloRpcBalance(context.Background(), usercenter.HelloRpcBalanceRequest{Context: &usercenter.UserContext{}}))
//	//printResponse(OperationSystem().SendOperationForLoginUser(context.Background(), operationsystem.SendOperationRequest{Context: &operationsystem.UserContext{}}))
//
//	for i := 0; i < 2; i++ { //
//		go printResponse(Task().GetTask(context.Background(), task.GetTaskRequest{Context: &task.UserContext{}}))
//		go printResponse(UserCenter().HelloRpcBalance(context.Background(), usercenter.HelloRpcBalanceRequest{Context: &usercenter.UserContext{}}))
//		//go printResponse(Task().GetTask(context.Background(), task.GetTaskRequest{Context: &task.UserContext{}}))
//		//go printResponse(UserCenter().HelloRpcBalance(context.Background(), usercenter.HelloRpcBalanceRequest{Context: &usercenter.UserContext{}}))
//		//go printResponse(Task().GetTask(context.Background(), task.GetTaskRequest{Context: &task.UserContext{}}))
//		//go printResponse(UserCenter().HelloRpcBalance(context.Background(), usercenter.HelloRpcBalanceRequest{Context: &usercenter.UserContext{}}))
//		//go printResponse(Task().GetTask(context.Background(), task.GetTaskRequest{Context: &task.UserContext{}}))
//		//go printResponse(UserCenter().HelloRpcBalance(context.Background(), usercenter.HelloRpcBalanceRequest{Context: &usercenter.UserContext{}}))
//		//go printResponse(Task().GetTask(context.Background(), task.GetTaskRequest{Context: &task.UserContext{}}))
//		//go printResponse(UserCenter().HelloRpcBalance(context.Background(), usercenter.HelloRpcBalanceRequest{Context: &usercenter.UserContext{}}))
//		//go printResponse(Task().GetTask(context.Background(), task.GetTaskRequest{Context: &task.UserContext{}}))
//		//go printResponse(UserCenter().HelloRpcBalance(context.Background(), usercenter.HelloRpcBalanceRequest{Context: &usercenter.UserContext{}}))
//		//go printResponse(Task().GetTask(context.Background(), task.GetTaskRequest{Context: &task.UserContext{}}))
//		//go printResponse(UserCenter().HelloRpcBalance(context.Background(), usercenter.HelloRpcBalanceRequest{Context: &usercenter.UserContext{}}))
//	}
//	time.Sleep(time.Second * 20)
//	printResponse(UserCenter().HelloRpcBalance(context.Background(), usercenter.HelloRpcBalanceRequest{Context: &usercenter.UserContext{}}))
//	select {}
//}
//
//func printResponse(response grpcbase.Response) {
//	println(response.Message, "  ", response.Data)
//}
//func newRpcServer(serviceName string) interface{} {
//	c, err := grpcbase.DialClient(grpcbase.ServerAddr(serviceName))
//	if err != nil {
//		log.Fatal(err)
//	}
//	return c
//}
//
//func UserCenter() usercenter.Repository {
//	r := newRpcServer(grpcbase.UserCenterSVC)
//	return r.(usercenter.Repository)
//}
//func Task() task.Repository {
//	r := newRpcServer(grpcbase.TaskSVC)
//	return r.(task.Repository)
//}
//func Notice() notice.Repository {
//	r := newRpcServer(grpcbase.NoticeSVC)
//	return r.(notice.Repository)
//}
//func Record() record.Repository {
//	r := newRpcServer(grpcbase.RecordSVC)
//	return r.(record.Repository)
//}
//
//func Cloudisk() clouddisk.Repository {
//	r := newRpcServer(grpcbase.CloudDiskSVC)
//	return r.(clouddisk.Repository)
//}
//
//func OperationSystem() operationsystem.Repository {
//	r := newRpcServer(grpcbase.OperationSystemSVC)
//	return r.(operationsystem.Repository)
//}
