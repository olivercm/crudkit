package schemas

import (
	"context"
	"crudkit/framework/rpcs"
	"crudkit/framework/types"
	"crudkit/framework/utils"
	apiCustomer "crudkit/server/api"
	"github.com/graphql-go/graphql"
	"time"
)

//rpc customer服务
var customerConn = rpcs.MustDial("customer")

func customerGetUserListField() *graphql.Field {
	return &graphql.Field{
		Type: types.NewGraphQLListTypeFromRPCType(apiCustomer.UserListData{}),
		Args: types.NewGraphQLArgsFromRPCType(apiCustomer.GetUserListReq{}),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			req := &apiCustomer.GetUserListReq{}
			err := utils.BindType(p.Args, req)
			if err != nil {
				return utils.BaseTypeErrors(err)
			}
			client := apiCustomer.NewCustomerServiceClient(customerConn)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()
			r, err := client.GetUserList(ctx, req)
			if err != nil {
				return utils.BaseTypeErrors(err)
			}
			return utils.BaseTypesManyOK(r.Result, int64(len(r.Result)), 1)
		},
	}
}

func customerCreateUserField() *graphql.Field {
	return &graphql.Field{
		Type: types.NewGraphQLTypeFromRPCType(apiCustomer.CreateUserResp{}),
		Args: types.NewGraphQLArgsFromRPCType(apiCustomer.CreateUserReq{}),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			req := &apiCustomer.CreateUserReq{}
			err := utils.BindType(p.Args, req)
			if err != nil {
				return utils.BaseTypeError(err)
			}
			client := apiCustomer.NewCustomerServiceClient(customerConn)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()
			r, err := client.CreateUser(ctx, req)
			if err != nil {
				return utils.BaseTypeError(err)
			}
			return utils.BaseTypeOK(r)
		},
	}
}

func customerUpdateUserField() *graphql.Field {
	return &graphql.Field{
		Type: types.NewGraphQLTypeFromRPCType(apiCustomer.UpdateUserResp{}),
		Args: types.NewGraphQLArgsFromRPCType(apiCustomer.UpdateUserReq{}),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			req := &apiCustomer.UpdateUserReq{}
			err := utils.BindType(p.Args, req)
			if err != nil {
				return utils.BaseTypeError(err)
			}
			client := apiCustomer.NewCustomerServiceClient(customerConn)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()
			r, err := client.UpdateUser(ctx, req)
			if err != nil {
				return utils.BaseTypeError(err)
			}
			return utils.BaseTypeOK(r)
		},
	}
}

func customerDeleteUserField() *graphql.Field {
	return &graphql.Field{
		Type: types.NewGraphQLTypeFromRPCType(apiCustomer.DeleteUserResp{}),
		Args: types.NewGraphQLArgsFromRPCType(apiCustomer.DeleteUserReq{}),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			req := &apiCustomer.DeleteUserReq{}
			err := utils.BindType(p.Args, req)
			if err != nil {
				return utils.BaseTypeError(err)
			}
			client := apiCustomer.NewCustomerServiceClient(customerConn)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()
			r, err := client.DeleteUser(ctx, req)
			if err != nil {
				return utils.BaseTypeError(err)
			}
			return utils.BaseTypeOK(r)
		},
	}
}
