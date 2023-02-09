//nolint:dupl
package user

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/announcement/user"

	constant "github.com/NpoolPlatform/notif-manager/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.ManagerClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get user connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewManagerClient(conn)

	return handler(_ctx, cli)
}

func CreateUser(ctx context.Context, in *npool.UserReq) (*npool.User, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateUser(ctx, &npool.CreateUserRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create user: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create user: %v", err)
	}
	return info.(*npool.User), nil
}

func CreateUsers(ctx context.Context, in []*npool.UserReq) ([]*npool.User, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateUsers(ctx, &npool.CreateUsersRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create users: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create users: %v", err)
	}
	return infos.([]*npool.User), nil
}

func UpdateUser(ctx context.Context, in *npool.UserReq) (*npool.User, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.UpdateUser(ctx, &npool.UpdateUserRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail add user: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update user: %v", err)
	}
	return info.(*npool.User), nil
}

func GetUser(ctx context.Context, id string) (*npool.User, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetUser(ctx, &npool.GetUserRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get user: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get user: %v", err)
	}
	return info.(*npool.User), nil
}

func GetUserOnly(ctx context.Context, conds *npool.Conds) (*npool.User, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetUserOnly(ctx, &npool.GetUserOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get user: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get user: %v", err)
	}
	return info.(*npool.User), nil
}

func GetUsers(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.User, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetUsers(ctx, &npool.GetUsersRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get users: %v", err)
		}
		total = resp.GetTotal()
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get users: %v", err)
	}
	return infos.([]*npool.User), total, nil
}

func ExistUser(ctx context.Context, id string) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistUser(ctx, &npool.ExistUserRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get user: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get user: %v", err)
	}
	return infos.(bool), nil
}

func ExistUserConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistUserConds(ctx, &npool.ExistUserCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get user: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get user: %v", err)
	}
	return infos.(bool), nil
}

func CountUsers(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CountUsers(ctx, &npool.CountUsersRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail count user: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return 0, fmt.Errorf("fail count user: %v", err)
	}
	return infos.(uint32), nil
}

func DeleteUser(ctx context.Context, id string) (*npool.User, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.DeleteUser(ctx, &npool.DeleteUserRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail delete user: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail delete user: %v", err)
	}
	return infos.(*npool.User), nil
}
