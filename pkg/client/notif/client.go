//nolint:dupl
package notif

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif"

	constant "github.com/NpoolPlatform/notif-manager/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.ManagerClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get notif connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewManagerClient(conn)

	return handler(_ctx, cli)
}

func CreateNotif(ctx context.Context, in *npool.NotifReq) (*npool.Notif, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateNotif(ctx, &npool.CreateNotifRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create notif: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create notif: %v", err)
	}
	return info.(*npool.Notif), nil
}

func CreateNotifs(ctx context.Context, in []*npool.NotifReq) ([]*npool.Notif, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateNotifs(ctx, &npool.CreateNotifsRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create notifs: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create notifs: %v", err)
	}
	return infos.([]*npool.Notif), nil
}

func UpdateNotif(ctx context.Context, in *npool.NotifReq) (*npool.Notif, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.UpdateNotif(ctx, &npool.UpdateNotifRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail add notif: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update notif: %v", err)
	}
	return info.(*npool.Notif), nil
}

func GetNotif(ctx context.Context, id string) (*npool.Notif, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetNotif(ctx, &npool.GetNotifRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get notif: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get notif: %v", err)
	}
	return info.(*npool.Notif), nil
}

func GetNotifOnly(ctx context.Context, conds *npool.Conds) (*npool.Notif, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetNotifOnly(ctx, &npool.GetNotifOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get notif: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get notif: %v", err)
	}
	return info.(*npool.Notif), nil
}

func GetNotifs(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.Notif, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetNotifs(ctx, &npool.GetNotifsRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get notifs: %v", err)
		}
		total = resp.GetTotal()
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get notifs: %v", err)
	}
	return infos.([]*npool.Notif), total, nil
}

func ExistNotif(ctx context.Context, id string) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistNotif(ctx, &npool.ExistNotifRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get notif: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get notif: %v", err)
	}
	return infos.(bool), nil
}

func ExistNotifConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistNotifConds(ctx, &npool.ExistNotifCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get notif: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get notif: %v", err)
	}
	return infos.(bool), nil
}

func CountNotifs(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CountNotifs(ctx, &npool.CountNotifsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail count notif: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return 0, fmt.Errorf("fail count notif: %v", err)
	}
	return infos.(uint32), nil
}

func DeleteNotif(ctx context.Context, id string) (*npool.Notif, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.DeleteNotif(ctx, &npool.DeleteNotifRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail delete notif: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail delete notif: %v", err)
	}
	return infos.(*npool.Notif), nil
}
