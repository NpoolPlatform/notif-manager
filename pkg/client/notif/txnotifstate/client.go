//nolint:dupl
package txnotifstate

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif/txnotifstate"

	constant "github.com/NpoolPlatform/notif-manager/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.ManagerClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get txnotifstate connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewManagerClient(conn)

	return handler(_ctx, cli)
}

func CreateTxNotifState(ctx context.Context, in *npool.TxNotifStateReq) (*npool.TxNotifState, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateTxNotifState(ctx, &npool.CreateTxNotifStateRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create txnotifstate: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create txnotifstate: %v", err)
	}
	return info.(*npool.TxNotifState), nil
}

func CreateTxNotifStates(ctx context.Context, in []*npool.TxNotifStateReq) ([]*npool.TxNotifState, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateTxNotifStates(ctx, &npool.CreateTxNotifStatesRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create txnotifstates: %v", err)
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create txnotifstates: %v", err)
	}
	return infos.([]*npool.TxNotifState), nil
}

func UpdateTxNotifState(ctx context.Context, in *npool.TxNotifStateReq) (*npool.TxNotifState, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.UpdateTxNotifState(ctx, &npool.UpdateTxNotifStateRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail add txnotifstate: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update txnotifstate: %v", err)
	}
	return info.(*npool.TxNotifState), nil
}

func GetTxNotifState(ctx context.Context, id string) (*npool.TxNotifState, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetTxNotifState(ctx, &npool.GetTxNotifStateRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get txnotifstate: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get txnotifstate: %v", err)
	}
	return info.(*npool.TxNotifState), nil
}

func GetTxNotifStateOnly(ctx context.Context, conds *npool.Conds) (*npool.TxNotifState, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetTxNotifStateOnly(ctx, &npool.GetTxNotifStateOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get txnotifstate: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get txnotifstate: %v", err)
	}
	return info.(*npool.TxNotifState), nil
}

func GetTxNotifStates(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.TxNotifState, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetTxNotifStates(ctx, &npool.GetTxNotifStatesRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get txnotifstates: %v", err)
		}
		total = resp.GetTotal()
		return resp.Infos, nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get txnotifstates: %v", err)
	}
	return infos.([]*npool.TxNotifState), total, nil
}

func ExistTxNotifState(ctx context.Context, id string) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistTxNotifState(ctx, &npool.ExistTxNotifStateRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get txnotifstate: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get txnotifstate: %v", err)
	}
	return infos.(bool), nil
}

func ExistTxNotifStateConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistTxNotifStateConds(ctx, &npool.ExistTxNotifStateCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get txnotifstate: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get txnotifstate: %v", err)
	}
	return infos.(bool), nil
}

func CountTxNotifStates(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CountTxNotifStates(ctx, &npool.CountTxNotifStatesRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail count txnotifstate: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return 0, fmt.Errorf("fail count txnotifstate: %v", err)
	}
	return infos.(uint32), nil
}

func DeleteTxNotifState(ctx context.Context, id string) (*npool.TxNotifState, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.DeleteTxNotifState(ctx, &npool.DeleteTxNotifStateRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail delete txnotifstate: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail delete txnotifstate: %v", err)
	}
	return infos.(*npool.TxNotifState), nil
}
