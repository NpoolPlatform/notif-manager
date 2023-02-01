//nolint:dupl
package sendstate

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/announcement/sendstate"

	constant "github.com/NpoolPlatform/notif-manager/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.ManagerClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get sendstate connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewManagerClient(conn)

	return handler(_ctx, cli)
}

func CreateSendState(ctx context.Context, in *npool.SendStateReq) (*npool.SendState, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateSendState(ctx, &npool.CreateSendStateRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create sendstate: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create sendstate: %v", err)
	}
	return info.(*npool.SendState), nil
}

func CreateSendStates(ctx context.Context, in []*npool.SendStateReq) ([]*npool.SendState, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateSendStates(ctx, &npool.CreateSendStatesRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create sendstates: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create sendstates: %v", err)
	}
	return infos.([]*npool.SendState), nil
}

func UpdateSendState(ctx context.Context, in *npool.SendStateReq) (*npool.SendState, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.UpdateSendState(ctx, &npool.UpdateSendStateRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail add sendstate: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update sendstate: %v", err)
	}
	return info.(*npool.SendState), nil
}

func GetSendState(ctx context.Context, id string) (*npool.SendState, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetSendState(ctx, &npool.GetSendStateRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get sendstate: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get sendstate: %v", err)
	}
	return info.(*npool.SendState), nil
}

func GetSendStateOnly(ctx context.Context, conds *npool.Conds) (*npool.SendState, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetSendStateOnly(ctx, &npool.GetSendStateOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get sendstate: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get sendstate: %v", err)
	}
	return info.(*npool.SendState), nil
}

func GetSendStates(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.SendState, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetSendStates(ctx, &npool.GetSendStatesRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get sendstates: %v", err)
		}
		total = resp.GetTotal()
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get sendstates: %v", err)
	}
	return infos.([]*npool.SendState), total, nil
}

func ExistSendState(ctx context.Context, id string) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistSendState(ctx, &npool.ExistSendStateRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get sendstate: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get sendstate: %v", err)
	}
	return infos.(bool), nil
}

func ExistSendStateConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistSendStateConds(ctx, &npool.ExistSendStateCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get sendstate: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get sendstate: %v", err)
	}
	return infos.(bool), nil
}

func CountSendStates(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CountSendStates(ctx, &npool.CountSendStatesRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail count sendstate: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return 0, fmt.Errorf("fail count sendstate: %v", err)
	}
	return infos.(uint32), nil
}

func DeleteSendState(ctx context.Context, id string) (*npool.SendState, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.DeleteSendState(ctx, &npool.DeleteSendStateRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail delete sendstate: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail delete sendstate: %v", err)
	}
	return infos.(*npool.SendState), nil
}
