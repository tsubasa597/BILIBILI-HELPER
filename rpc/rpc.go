package rpc

import (
	"sync"

	"github.com/tsubasa597/BILIBILI-HELPER/api/comment"
	"github.com/tsubasa597/BILIBILI-HELPER/api/dynamic"
	"github.com/tsubasa597/BILIBILI-HELPER/info"
	"github.com/tsubasa597/BILIBILI-HELPER/rpc/service"
	"go.uber.org/zap"
)

// Comment rpc 服务
type Comment struct {
	Log *zap.Logger
}

// Dynamic rpc 服务
type Dynamic struct {
	Log *zap.Logger
}

var (
	_            service.CommentServer = (*Comment)(nil)
	_            service.DynamicServer = (*Dynamic)(nil)
	_commentPool *sync.Pool            = &sync.Pool{
		New: func() interface{} {
			return &service.CommentResponse{}
		},
	}
	_dynamicPool *sync.Pool = &sync.Pool{
		New: func() interface{} {
			return &service.DynamicResponse{}
		},
	}
)

// GetAll 获取评论区所有内容
func (c Comment) GetAll(req *service.AllCommentRequest, server service.Comment_GetAllServer) error {
	comms := comment.GetAllComments(
		info.DynamicType(req.BaseCommentRequest.Type),
		req.BaseCommentRequest.RID,
		req.Time,
	)
	for _, comm := range comms {
		resp := _commentPool.Get().(*service.CommentResponse)
		resp.DynamicUID = comm.DynamicUID
		resp.UID = comm.UID
		resp.RID = comm.RID
		resp.LikeNum = int32(comm.LikeNum)
		resp.Content = comm.Content
		resp.Time = comm.Time
		resp.Rpid = comm.Rpid
		resp.Name = comm.Name

		err := server.Send(resp)
		if err != nil {
			c.Log.Error(err.Error())
		}

		_commentPool.Put(resp)
	}

	return nil
}

// Get 获取评论区指定页数的内容
func (c Comment) Get(req *service.CommentRequest, server service.Comment_GetServer) error {
	comms, err := comment.GetComments(
		info.DynamicType(req.BaseCommentRequest.Type),
		info.Sort(req.Sort),
		req.BaseCommentRequest.RID,
		int(req.PageSum),
		int(req.PageNum),
	)
	if err != nil {
		return err
	}

	for _, comm := range comms {
		resp := _commentPool.Get().(*service.CommentResponse)
		resp.DynamicUID = comm.DynamicUID
		resp.UID = comm.UID
		resp.RID = comm.RID
		resp.LikeNum = int32(comm.LikeNum)
		resp.Content = comm.Content
		resp.Time = comm.Time
		resp.Rpid = comm.Rpid
		resp.Name = comm.Name

		err := server.Send(resp)
		if err != nil {
			c.Log.Error(err.Error())
		}

		_commentPool.Put(resp)
	}

	return nil
}

// Get 获取指定动态之后的一页动态
func (dy Dynamic) Get(req *service.DynamicRequest, server service.Dynamic_GetServer) error {
	dynamics, err := dynamic.GetDynamics(req.BaseCommentRequest.UID, req.Offect)
	if err != nil {
		return err
	}

	for _, d := range dynamics {
		resp := _dynamicPool.Get().(*service.DynamicResponse)
		resp.UID = d.UID
		resp.Content = d.Content
		resp.Card = d.Card
		resp.RID = d.RID
		resp.Offect = d.Offect
		resp.Type = int32(d.Type)
		resp.Name = d.Name
		resp.Time = d.Time

		err := server.Send(resp)
		if err != nil {
			dy.Log.Error(err.Error())
		}

		_dynamicPool.Put(resp)

	}

	return nil
}

// GetAll 获取指定时间之后的所有动态
func (dy Dynamic) GetAll(req *service.AllDynamicRequest, server service.Dynamic_GetAllServer) error {
	dynamics := dynamic.GetAllDynamics(req.BaseCommentRequest.UID, req.Time)
	for _, d := range dynamics {
		resp := _dynamicPool.Get().(*service.DynamicResponse)
		resp.UID = d.UID
		resp.Content = d.Content
		resp.Card = d.Card
		resp.RID = d.RID
		resp.Offect = d.Offect
		resp.Type = int32(d.Type)
		resp.Name = d.Name
		resp.Time = d.Time

		err := server.Send(resp)
		if err != nil {
			dy.Log.Error(err.Error())
		}

		_dynamicPool.Put(resp)
	}

	return nil
}
