package service

import (
	"context"
	pb "crudkit/server/api"
	"crudkit/server/dao"
	"crudkit/server/model"
)

type {{.Server}}Service struct {
	dao dao.Dao
}

func (s *{{.Server}}Service) Get{{.Model}}List(ctx context.Context, req *pb.Get{{.Model}}ListReq) (*pb.Get{{.Model}}ListResp, error) {
	datas, err := s.dao.Get{{.Model}}List(req.Current, req.PageSize)
	if err != nil {
		return nil, err
	}
	resp := new(pb.Get{{.Model}}ListResp)
	for _, data := range datas {
		resp.Result = append(resp.Result, &pb.{{.Model}}ListData{
			Id:         data.Id,
			Name:       data.Name,
			Age:        data.Age,
			CreateTime: data.CreateTime,
		})
	}
	return resp, nil
}

func (s *{{.Server}}Service) Create{{.Model}}(ctx context.Context, req *pb.Create{{.Model}}Req) (*pb.Create{{.Model}}Resp, error) {
	e := &model.{{.Model}}{
		Name: req.Name,
		Age:  req.Age,
	}
	_, err := s.dao.Create{{.Model}}(e)
	if err != nil {
		return nil, err
	}
	return &pb.Create{{.Model}}Resp{}, nil
}

func (s *{{.Server}}Service) Update{{.Model}}(ctx context.Context, req *pb.Update{{.Model}}Req) (*pb.Update{{.Model}}Resp, error) {
	e := &model.{{.Model}}{
		Id:   req.Id,
		Name: req.Name,
		Age:  req.Age,
	}
	err := s.dao.Update{{.Model}}(e)
	if err != nil {
		return nil, err
	}
	return &pb.Update{{.Model}}Resp{}, nil
}

func (s *{{.Server}}Service) Delete{{.Model}}(ctx context.Context, req *pb.Delete{{.Model}}Req) (*pb.Delete{{.Model}}Resp, error) {
	err := s.dao.Delete{{.Model}}ById(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Delete{{.Model}}Resp{}, nil
}
