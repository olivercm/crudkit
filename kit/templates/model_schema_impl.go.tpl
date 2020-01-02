package schemas

import (
	"context"
	"crudkit/framework/rpcs"
	"crudkit/framework/types"
	"crudkit/framework/utils"
	api{{.Server}} "crudkit/server/api"
	"github.com/graphql-go/graphql"
	"time"
)

//rpc {{.LCFirstServer}}服务
var {{.LCFirstServer}}Conn = rpcs.MustDial("{{.LCFirstServer}}")

func {{.LCFirstServer}}Get{{.Model}}ListField() *graphql.Field {
	return &graphql.Field{
		Type: types.NewGraphQLListTypeFromRPCType(api{{.Server}}.{{.Model}}ListData{}),
		Args: types.NewGraphQLArgsFromRPCType(api{{.Server}}.Get{{.Model}}ListReq{}),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			req := &api{{.Server}}.Get{{.Model}}ListReq{}
			err := utils.BindType(p.Args, req)
			if err != nil {
				return utils.BaseTypeErrors(err)
			}
			client := api{{.Server}}.New{{.Server}}ServiceClient({{.LCFirstServer}}Conn)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()
			r, err := client.Get{{.Model}}List(ctx, req)
			if err != nil {
				return utils.BaseTypeErrors(err)
			}
			return utils.BaseTypesManyOK(r.Result, int64(len(r.Result)), 1)
		},
	}
}

func {{.LCFirstServer}}Create{{.Model}}Field() *graphql.Field {
	return &graphql.Field{
		Type: types.NewGraphQLTypeFromRPCType(api{{.Server}}.Create{{.Model}}Resp{}),
		Args: types.NewGraphQLArgsFromRPCType(api{{.Server}}.Create{{.Model}}Req{}),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			req := &api{{.Server}}.Create{{.Model}}Req{}
			err := utils.BindType(p.Args, req)
			if err != nil {
				return utils.BaseTypeError(err)
			}
			client := api{{.Server}}.New{{.Server}}ServiceClient({{.LCFirstServer}}Conn)
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

func {{.LCFirstServer}}Update{{.Model}}Field() *graphql.Field {
	return &graphql.Field{
		Type: types.NewGraphQLTypeFromRPCType(api{{.Server}}.Update{{.Model}}Resp{}),
		Args: types.NewGraphQLArgsFromRPCType(api{{.Server}}.Update{{.Model}}Req{}),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			req := &api{{.Server}}.Update{{.Model}}Req{}
			err := utils.BindType(p.Args, req)
			if err != nil {
				return utils.BaseTypeError(err)
			}
			client := api{{.Server}}.New{{.Server}}ServiceClient({{.LCFirstServer}}Conn)
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

func {{.LCFirstServer}}Delete{{.Model}}Field() *graphql.Field {
	return &graphql.Field{
		Type: types.NewGraphQLTypeFromRPCType(api{{.Server}}.Delete{{.Model}}Resp{}),
		Args: types.NewGraphQLArgsFromRPCType(api{{.Server}}.Delete{{.Model}}Req{}),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			req := &api{{.Server}}.Delete{{.Model}}Req{}
			err := utils.BindType(p.Args, req)
			if err != nil {
				return utils.BaseTypeError(err)
			}
			client := api{{.Server}}.New{{.Server}}ServiceClient({{.LCFirstServer}}Conn)
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
