//nolint:dupl
package tx

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif/tx"

	constant "github.com/NpoolPlatform/notif-manager/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.ManagerClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get tx connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewManagerClient(conn)

	return handler(_ctx, cli)
}

func CreateTx(ctx context.Context, in *npool.TxReq) (*npool.Tx, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateTx(ctx, &npool.CreateTxRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create tx: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create tx: %v", err)
	}
	return info.(*npool.Tx), nil
}

func CreateTxs(ctx context.Context, in []*npool.TxReq) ([]*npool.Tx, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateTxs(ctx, &npool.CreateTxsRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create txs: %v", err)
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create txs: %v", err)
	}
	return infos.([]*npool.Tx), nil
}

func UpdateTx(ctx context.Context, in *npool.TxReq) (*npool.Tx, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.UpdateTx(ctx, &npool.UpdateTxRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail add tx: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update tx: %v", err)
	}
	return info.(*npool.Tx), nil
}

func GetTx(ctx context.Context, id string) (*npool.Tx, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetTx(ctx, &npool.GetTxRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get tx: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get tx: %v", err)
	}
	return info.(*npool.Tx), nil
}

func GetTxOnly(ctx context.Context, conds *npool.Conds) (*npool.Tx, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetTxOnly(ctx, &npool.GetTxOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get tx: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get tx: %v", err)
	}
	return info.(*npool.Tx), nil
}

func GetTxs(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.Tx, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetTxs(ctx, &npool.GetTxsRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get txs: %v", err)
		}
		total = resp.GetTotal()
		return resp.Infos, nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get txs: %v", err)
	}
	return infos.([]*npool.Tx), total, nil
}

func ExistTx(ctx context.Context, id string) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistTx(ctx, &npool.ExistTxRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get tx: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get tx: %v", err)
	}
	return infos.(bool), nil
}

func ExistTxConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistTxConds(ctx, &npool.ExistTxCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get tx: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get tx: %v", err)
	}
	return infos.(bool), nil
}

func CountTxs(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CountTxs(ctx, &npool.CountTxsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail count tx: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return 0, fmt.Errorf("fail count tx: %v", err)
	}
	return infos.(uint32), nil
}

func DeleteTx(ctx context.Context, id string) (*npool.Tx, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.DeleteTx(ctx, &npool.DeleteTxRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail delete tx: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail delete tx: %v", err)
	}
	return infos.(*npool.Tx), nil
}
