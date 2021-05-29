package server

import (
	"blog_gin/tag-service/pkg/bapi"
	pb "blog_gin/tag-service/proto"
	"context"
	"encoding/json"
	"github.com/JoyZF/blog_gin/pkg/errcode"
)




type TagServer struct {}

func NewTagServer() *TagServer {
	return &TagServer{}
}

func (t *TagServer) GetTagList(ctx context.Context, r *pb.GetTagListRequest) (*pb.GetTagListReply, error) {
	api := bapi.NewAPI("http://127.0.0.1:8000")
	body, err := api.GetTagList(ctx, r.GetName())
	if err != nil {
		return nil, err
	}

	tagList := pb.GetTagListReply{}
	err = json.Unmarshal(body, &tagList)
	if err != nil {
		return nil,errcode.TogRPCError(errcode.Fail)
	}

	return &tagList, nil
}