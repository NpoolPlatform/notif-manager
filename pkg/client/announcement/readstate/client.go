//nolint:dupl
package readstate

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/announcement/readstate"

	constant "github.com/NpoolPlatform/notif-manager/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.ManagerClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get readstate connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewManagerClient(conn)

	return handler(_ctx, cli)
}

func CreateReadState(ctx context.Context, in *npool.ReadStateReq) (*npool.ReadState, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateReadState(ctx, &npool.CreateReadStateRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create readstate: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create readstate: %v", err)
	}
	return info.(*npool.ReadState), nil
}

func CreateReadStates(ctx context.Context, in []*npool.ReadStateReq) ([]*npool.ReadState, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateReadStates(ctx, &npool.CreateReadStatesRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create readstates: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create readstates: %v", err)
	}
	return infos.([]*npool.ReadState), nil
}

func UpdateReadState(ctx context.Context, in *npool.ReadStateReq) (*npool.ReadState, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.UpdateReadState(ctx, &npool.UpdateReadStateRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail add readstate: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update readstate: %v", err)
	}
	return info.(*npool.ReadState), nil
}

func GetReadState(ctx context.Context, id string) (*npool.ReadState, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetReadState(ctx, &npool.GetReadStateRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get readstate: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get readstate: %v", err)
	}
	return info.(*npool.ReadState), nil
}

func GetReadStateOnly(ctx context.Context, conds *npool.Conds) (*npool.ReadState, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetReadStateOnly(ctx, &npool.GetReadStateOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get readstate: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get readstate: %v", err)
	}
	return info.(*npool.ReadState), nil
}

func GetReadStates(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.ReadState, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetReadStates(ctx, &npool.GetReadStatesRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get readstates: %v", err)
		}
		total = resp.GetTotal()
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get readstates: %v", err)
	}
	return infos.([]*npool.ReadState), total, nil
}

func ExistReadState(ctx context.Context, id string) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistReadState(ctx, &npool.ExistReadStateRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get readstate: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get readstate: %v", err)
	}
	return infos.(bool), nil
}

func ExistReadStateConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistReadStateConds(ctx, &npool.ExistReadStateCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get readstate: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get readstate: %v", err)
	}
	return infos.(bool), nil
}

func CountReadStates(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CountReadStates(ctx, &npool.CountReadStatesRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail count readstate: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return 0, fmt.Errorf("fail count readstate: %v", err)
	}
	return infos.(uint32), nil
}

func DeleteReadState(ctx context.Context, id string) (*npool.ReadState, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.DeleteReadState(ctx, &npool.DeleteReadStateRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail delete readstate: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail delete readstate: %v", err)
	}
	return infos.(*npool.ReadState), nil
}
