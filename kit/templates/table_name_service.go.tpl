package service

import (
	"context"
	pb "crudkit/crud/server/api"
	"crudkit/crud/server/dao"
	"crudkit/crud/server/model"
)

type {{.UCFirstServer}}Service struct {
	dao dao.Dao
}

func (s *{{.UCFirstServer}}Service) Get{{.Model}}List(ctx context.Context, req *pb.Get{{.Model}}ListReq) (*pb.Get{{.Model}}ListResp, error) {
	datas, err := s.dao.Get{{.Model}}List(req.CurrentPage, req.PageSize)
	if err != nil {
		return nil, err
	}
	resp := new(pb.Get{{.Model}}ListResp)
	for _, data := range datas {
		resp.Result = append(resp.Result, &pb.{{.Model}}ListData{
		    {{range .ServiceListFields}}{{.}}
            {{end}}
		})
	}
	total, err := s.dao.Count{{.Model}}()
	if err != nil {
        return nil, err
    }
	resp.Total = int64(total)
	return resp, nil
}

func (s *{{.UCFirstServer}}Service) Create{{.Model}}(ctx context.Context, req *pb.Create{{.Model}}Req) (*pb.Create{{.Model}}Resp, error) {
	e := &model.{{.Model}}{
		{{range .ServiceFields}}{{.}}
        {{end}}
	}
	_, err := s.dao.Create{{.Model}}(e)
	if err != nil {
		return nil, err
	}
	return &pb.Create{{.Model}}Resp{}, nil
}

func (s *{{.UCFirstServer}}Service) Update{{.Model}}(ctx context.Context, req *pb.Update{{.Model}}Req) (*pb.Update{{.Model}}Resp, error) {
	e := &model.{{.Model}}{
		{{range .ServiceFields}}{{.}}
        {{end}}
	}
	err := s.dao.Update{{.Model}}(e)
	if err != nil {
		return nil, err
	}
	return &pb.Update{{.Model}}Resp{}, nil
}

func (s *{{.UCFirstServer}}Service) Delete{{.Model}}(ctx context.Context, req *pb.Delete{{.Model}}Req) (*pb.Delete{{.Model}}Resp, error) {
	err := s.dao.Delete{{.Model}}ById(req.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Delete{{.Model}}Resp{}, nil
}
