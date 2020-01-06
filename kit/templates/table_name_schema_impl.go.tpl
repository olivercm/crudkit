package schemas

import (
	"context"
	"crudkit/framework/rpcs"
	"crudkit/framework/types"
	"crudkit/framework/utils"
	api{{.UCFirstServer}} "crudkit/crud/server/api"
	"github.com/graphql-go/graphql"
	"time"
)

//rpc {{.Server}}服务
var {{.Server}}Conn = rpcs.MustDial("{{.Server}}")

func {{.Server}}Get{{.Model}}ListField() *graphql.Field {
	return &graphql.Field{
		Type: types.NewGraphQLListTypeFromRPCType(api{{.UCFirstServer}}.{{.Model}}ListData{}),
		Args: types.NewGraphQLArgsFromRPCType(api{{.UCFirstServer}}.Get{{.Model}}ListReq{}),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			req := &api{{.UCFirstServer}}.Get{{.Model}}ListReq{}
			err := utils.BindType(p.Args, req)
			if err != nil {
				return utils.BaseTypeErrors(err)
			}
			client := api{{.UCFirstServer}}.New{{.UCFirstServer}}ServiceClient({{.Server}}Conn)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()
			r, err := client.Get{{.Model}}List(ctx, req)
			if err != nil {
				return utils.BaseTypeErrors(err)
			}
			return utils.BaseTypesManyOK(r.Result, int64(len(r.Result)), req.CurrentPage)
		},
	}
}

func {{.Server}}Create{{.Model}}Field() *graphql.Field {
	return &graphql.Field{
		Type: types.NewGraphQLTypeFromRPCType(api{{.UCFirstServer}}.Create{{.Model}}Resp{}),
		Args: types.NewGraphQLArgsFromRPCType(api{{.UCFirstServer}}.Create{{.Model}}Req{}),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			req := &api{{.UCFirstServer}}.Create{{.Model}}Req{}
			err := utils.BindType(p.Args, req)
			if err != nil {
				return utils.BaseTypeError(err)
			}
			client := api{{.UCFirstServer}}.New{{.UCFirstServer}}ServiceClient({{.Server}}Conn)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()
			r, err := client.Create{{.Model}}(ctx, req)
			if err != nil {
				return utils.BaseTypeError(err)
			}
			return utils.BaseTypeOK(r)
		},
	}
}

func {{.Server}}Update{{.Model}}Field() *graphql.Field {
	return &graphql.Field{
		Type: types.NewGraphQLTypeFromRPCType(api{{.UCFirstServer}}.Update{{.Model}}Resp{}),
		Args: types.NewGraphQLArgsFromRPCType(api{{.UCFirstServer}}.Update{{.Model}}Req{}),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			req := &api{{.UCFirstServer}}.Update{{.Model}}Req{}
			err := utils.BindType(p.Args, req)
			if err != nil {
				return utils.BaseTypeError(err)
			}
			client := api{{.UCFirstServer}}.New{{.UCFirstServer}}ServiceClient({{.Server}}Conn)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()
			r, err := client.Update{{.Model}}(ctx, req)
			if err != nil {
				return utils.BaseTypeError(err)
			}
			return utils.BaseTypeOK(r)
		},
	}
}

func {{.Server}}Delete{{.Model}}Field() *graphql.Field {
	return &graphql.Field{
		Type: types.NewGraphQLTypeFromRPCType(api{{.UCFirstServer}}.Delete{{.Model}}Resp{}),
		Args: types.NewGraphQLArgsFromRPCType(api{{.UCFirstServer}}.Delete{{.Model}}Req{}),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			req := &api{{.UCFirstServer}}.Delete{{.Model}}Req{}
			err := utils.BindType(p.Args, req)
			if err != nil {
				return utils.BaseTypeError(err)
			}
			client := api{{.UCFirstServer}}.New{{.UCFirstServer}}ServiceClient({{.Server}}Conn)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()
			r, err := client.Delete{{.Model}}(ctx, req)
			if err != nil {
				return utils.BaseTypeError(err)
			}
			return utils.BaseTypeOK(r)
		},
	}
}
