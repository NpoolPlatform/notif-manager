//nolint:dupl
package notifchannel

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif/notifchannel"

	constant "github.com/NpoolPlatform/notif-manager/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.ManagerClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get notifchannel connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewManagerClient(conn)

	return handler(_ctx, cli)
}

func CreateNotifChannel(ctx context.Context, in *npool.NotifChannelReq) (*npool.NotifChannel, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateNotifChannel(ctx, &npool.CreateNotifChannelRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create notifchannel: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create notifchannel: %v", err)
	}
	return info.(*npool.NotifChannel), nil
}

func CreateNotifChannels(ctx context.Context, in []*npool.NotifChannelReq) ([]*npool.NotifChannel, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateNotifChannels(ctx, &npool.CreateNotifChannelsRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create notifchannels: %v", err)
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create notifchannels: %v", err)
	}
	return infos.([]*npool.NotifChannel), nil
}

func GetNotifChannel(ctx context.Context, id string) (*npool.NotifChannel, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetNotifChannel(ctx, &npool.GetNotifChannelRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get notifchannel: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get notifchannel: %v", err)
	}
	return info.(*npool.NotifChannel), nil
}

func GetNotifChannelOnly(ctx context.Context, conds *npool.Conds) (*npool.NotifChannel, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetNotifChannelOnly(ctx, &npool.GetNotifChannelOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get notifchannel: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get notifchannel: %v", err)
	}
	return info.(*npool.NotifChannel), nil
}

func GetNotifChannels(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.NotifChannel, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetNotifChannels(ctx, &npool.GetNotifChannelsRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get notifchannels: %v", err)
		}
		total = resp.GetTotal()
		return resp.Infos, nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get notifchannels: %v", err)
	}
	return infos.([]*npool.NotifChannel), total, nil
}

func ExistNotifChannel(ctx context.Context, id string) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistNotifChannel(ctx, &npool.ExistNotifChannelRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get notifchannel: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get notifchannel: %v", err)
	}
	return infos.(bool), nil
}

func ExistNotifChannelConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistNotifChannelConds(ctx, &npool.ExistNotifChannelCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get notifchannel: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get notifchannel: %v", err)
	}
	return infos.(bool), nil
}

func CountNotifChannels(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CountNotifChannels(ctx, &npool.CountNotifChannelsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail count notifchannel: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return 0, fmt.Errorf("fail count notifchannel: %v", err)
	}
	return infos.(uint32), nil
}

func DeleteNotifChannel(ctx context.Context, id string) (*npool.NotifChannel, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.DeleteNotifChannel(ctx, &npool.DeleteNotifChannelRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail delete notifchannel: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail delete notifchannel: %v", err)
	}
	return infos.(*npool.NotifChannel), nil
}
