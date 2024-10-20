package rpc

import "context"

type UserService struct {
	GetById func(ctx context.Context, req *GetByIdReq) (*GetByIdResp, error)
}

func (s *UserService) Name() string {
	return "user-service"
}

type GetByIdReq struct {
	Id int64 `json:"id"`
}

type GetByIdResp struct {
}
