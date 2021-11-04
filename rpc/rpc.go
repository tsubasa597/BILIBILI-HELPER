package rpc

import (
	"github.com/sirupsen/logrus"
	"github.com/tsubasa597/BILIBILI-HELPER/api"
	"github.com/tsubasa597/BILIBILI-HELPER/info"
	"github.com/tsubasa597/BILIBILI-HELPER/rpc/service"
)

// Comment rpc 服务
type Comment struct {
	Log *logrus.Entry
}

// Dynamic rpc 服务
type Dynamic struct {
	Log *logrus.Entry
}

var (
	_ service.CommentServer = (*Comment)(nil)
	_ service.DynamicServer = (*Dynamic)(nil)
)

// GetAll 获取评论区所有内容
func (c Comment) GetAll(req *service.AllCommentRequest, server service.Comment_GetAllServer) error {
	comms := api.GetAllComments(info.Type(req.BaseCommentRequest.Type),
		req.BaseCommentRequest.RID, req.Time)
	for _, comm := range comms {
		err := server.Send(&service.CommentResponse{
			UserID:  comm.UserID,
			UID:     comm.UID,
			RID:     comm.RID,
			Like:    int32(comm.LikeNum),
			Content: comm.Content,
			Time:    comm.Time,
			Rpid:    comm.Rpid,
			Name:    comm.Name,
		})
		if err != nil {
			c.Log.Error(err)
		}
	}

	return nil
}

// Get 获取评论区指定页数的内容
func (c Comment) Get(req *service.CommentRequest, server service.Comment_GetServer) error {
	comms, err := api.GetComments(info.Type(req.BaseCommentRequest.Type),
		info.Sort(req.Sort), req.BaseCommentRequest.RID, int(req.PageSum), int(req.PageNum))
	if err != nil {
		return err
	}

	for _, reply := range comms.Replies {
		err := server.Send(&service.CommentResponse{
			Time:    reply.Ctime,
			UserID:  comms.Upper.Mid,
			UID:     reply.Mid,
			Rpid:    reply.Rpid,
			Like:    reply.Like,
			Content: reply.Content.Message,
			RID:     req.BaseCommentRequest.RID,
			Name:    reply.Member.Uname,
		})
		if err != nil {
			c.Log.Error(err)
		}
	}

	return nil
}

// Get 获取指定动态之后的一页动态
func (dy Dynamic) Get(req *service.DynamicRequest, server service.Dynamic_GetServer) error {
	cards, err := api.GetDynamics(req.BaseCommentRequest.UID, req.Offect)
	if err != nil {
		return err
	}

	for _, card := range cards {
		if d, err := api.GetOriginCard(card); err == nil {
			err := server.Send(&service.DynamicResponse{
				UID:     d.UID,
				Content: d.Content,
				Card:    d.Card,
				RID:     d.RID,
				Offect:  d.Offect,
				Type:    int32(d.Type),
				Name:    d.Name,
				Time:    d.Time,
			})
			if err != nil {
				dy.Log.Error(err)
			}
		}
	}

	return nil
}

// GetAll 获取指定时间之后的所有动态
func (dy Dynamic) GetAll(req *service.AllDynamicRequest, server service.Dynamic_GetAllServer) error {
	dynamics := api.GetAllDynamics(req.BaseCommentRequest.UID, req.Time)
	for _, dynamic := range dynamics {
		err := server.Send(&service.DynamicResponse{
			UID:     dynamic.UID,
			Content: dynamic.Content,
			Card:    dynamic.Card,
			RID:     dynamic.RID,
			Offect:  dynamic.Offect,
			Type:    int32(dynamic.Type),
			Name:    dynamic.Name,
			Time:    dynamic.Time,
		})
		if err != nil {
			dy.Log.Error(err)
		}
	}

	return nil
}
