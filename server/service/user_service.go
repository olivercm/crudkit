package service

import (
	"context"
	pb "crudkit/server/api"
	"crudkit/server/dao"
	"crudkit/server/model"
)

type CustomerService struct {
	dao dao.Dao
}

func (s *CustomerService) GetUserList(ctx context.Context, req *pb.GetUserListReq) (*pb.GetUserListResp, error) {
	datas, err := s.dao.GetUserList(req.Current, req.PageSize)
	if err != nil {
		return nil, err
	}
	resp := new(pb.GetUserListResp)
	for _, data := range datas {
		resp.Result = append(resp.Result, &pb.UserListData{
		    Id: data.Id,
            Name: data.Name,
            Age: data.Age,
            CreateTime: data.CreateTime,
            UpdateTime: data.UpdateTime,
            DeleteTime: data.DeleteTime,
            
		})
	}
	return resp, nil
}

func (s *CustomerService) CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.CreateUserResp, error) {
	e := &model.User{
		Id: req.Id,
        Name: req.Name,
        Age: req.Age,
        CreateTime: req.CreateTime,
        UpdateTime: req.UpdateTime,
        DeleteTime: req.DeleteTime,
        
	}
	_, err := s.dao.CreateUser(e)
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserResp{}, nil
}

func (s *CustomerService) UpdateUser(ctx context.Context, req *pb.UpdateUserReq) (*pb.UpdateUserResp, error) {
	e := &model.User{
		Id: req.Id,
        Name: req.Name,
        Age: req.Age,
        CreateTime: req.CreateTime,
        UpdateTime: req.UpdateTime,
        DeleteTime: req.DeleteTime,
        
	}
	err := s.dao.UpdateUser(e)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateUserResp{}, nil
}

func (s *CustomerService) DeleteUser(ctx context.Context, req *pb.DeleteUserReq) (*pb.DeleteUserResp, error) {
	err := s.dao.DeleteUserById(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteUserResp{}, nil
}
