package server

import (
	"blog_gin/tag-service/pkg/bapi"
	pb "blog_gin/tag-service/proto"
	"context"
	"encoding/json"
	"errors"
)




type TagServer struct {}

func NewTagServer() *TagServer {
	return &TagServer{}
}

func (t *TagServer) GetTagList(ctx context.Context, r *pb.GetTagListRequest) (*pb.GetTagListReply, error) {
	api := bapi.NewAPI("http://127.0.0.1:8080")
	body, err := api.GetTagList(ctx, r.GetName())
	if err != nil {
		return nil, err
	}

	tagList := pb.GetTagListReply{}
	err = json.Unmarshal(body, &tagList)
	if err != nil {
		return nil,errors.New("RPC ERROR")
	}

	return &tagList, nil
}