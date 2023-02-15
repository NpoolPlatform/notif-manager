//nolint:dupl
package channel

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif/channel"

	constant "github.com/NpoolPlatform/notif-manager/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.ManagerClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get channel connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewManagerClient(conn)

	return handler(_ctx, cli)
}

func CreateChannel(ctx context.Context, in *npool.ChannelReq) (*npool.Channel, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateChannel(ctx, &npool.CreateChannelRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create channel: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create channel: %v", err)
	}
	return info.(*npool.Channel), nil
}

func CreateChannels(ctx context.Context, in []*npool.ChannelReq) ([]*npool.Channel, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateChannels(ctx, &npool.CreateChannelsRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create channels: %v", err)
		}
		return resp.Infos, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create channels: %v", err)
	}
	return infos.([]*npool.Channel), nil
}

func GetChannel(ctx context.Context, id string) (*npool.Channel, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetChannel(ctx, &npool.GetChannelRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get channel: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get channel: %v", err)
	}
	return info.(*npool.Channel), nil
}

func GetChannelOnly(ctx context.Context, conds *npool.Conds) (*npool.Channel, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetChannelOnly(ctx, &npool.GetChannelOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get channel: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get channel: %v", err)
	}
	return info.(*npool.Channel), nil
}

func GetChannels(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.Channel, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetChannels(ctx, &npool.GetChannelsRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get channels: %v", err)
		}
		total = resp.GetTotal()
		return resp.Infos, nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get channels: %v", err)
	}
	return infos.([]*npool.Channel), total, nil
}

func ExistChannel(ctx context.Context, id string) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistChannel(ctx, &npool.ExistChannelRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get channel: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get channel: %v", err)
	}
	return infos.(bool), nil
}

func ExistChannelConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistChannelConds(ctx, &npool.ExistChannelCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get channel: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get channel: %v", err)
	}
	return infos.(bool), nil
}

func CountChannels(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CountChannels(ctx, &npool.CountChannelsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail count channel: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return 0, fmt.Errorf("fail count channel: %v", err)
	}
	return infos.(uint32), nil
}

func DeleteChannel(ctx context.Context, id string) (*npool.Channel, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.DeleteChannel(ctx, &npool.DeleteChannelRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail delete channel: %v", err)
		}
		return resp.Info, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail delete channel: %v", err)
	}
	return infos.(*npool.Channel), nil
}
