// Code generated by Kitex v0.4.4. DO NOT EDIT.

package chatservice

import (
	"context"
	chat "github.com/808-not-found/tik_duck/kitex_gen/chat"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return chatServiceServiceInfo
}

var chatServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "ChatService"
	handlerType := (*chat.ChatService)(nil)
	methods := map[string]kitex.MethodInfo{
		"GetMsg":  kitex.NewMethodInfo(getMsgHandler, newChatServiceGetMsgArgs, newChatServiceGetMsgResult, false),
		"PostMsg": kitex.NewMethodInfo(postMsgHandler, newChatServicePostMsgArgs, newChatServicePostMsgResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "chat",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func getMsgHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*chat.ChatServiceGetMsgArgs)
	realResult := result.(*chat.ChatServiceGetMsgResult)
	success, err := handler.(chat.ChatService).GetMsg(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newChatServiceGetMsgArgs() interface{} {
	return chat.NewChatServiceGetMsgArgs()
}

func newChatServiceGetMsgResult() interface{} {
	return chat.NewChatServiceGetMsgResult()
}

func postMsgHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*chat.ChatServicePostMsgArgs)
	realResult := result.(*chat.ChatServicePostMsgResult)
	success, err := handler.(chat.ChatService).PostMsg(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newChatServicePostMsgArgs() interface{} {
	return chat.NewChatServicePostMsgArgs()
}

func newChatServicePostMsgResult() interface{} {
	return chat.NewChatServicePostMsgResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) GetMsg(ctx context.Context, req *chat.MessageChatRequest) (r *chat.MessageChatResponse, err error) {
	var _args chat.ChatServiceGetMsgArgs
	_args.Req = req
	var _result chat.ChatServiceGetMsgResult
	if err = p.c.Call(ctx, "GetMsg", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) PostMsg(ctx context.Context, req *chat.RelationActionRequest) (r *chat.RelationActionResponse, err error) {
	var _args chat.ChatServicePostMsgArgs
	_args.Req = req
	var _result chat.ChatServicePostMsgResult
	if err = p.c.Call(ctx, "PostMsg", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}