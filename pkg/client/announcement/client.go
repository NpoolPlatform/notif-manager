//nolint:dupl
package announcement

import (
	"context"
	"fmt"
	"time"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/announcement"

	constant "github.com/NpoolPlatform/notif-manager/pkg/message/const"
)

var timeout = 10 * time.Second

type handler func(context.Context, npool.ManagerClient) (cruder.Any, error)

func withCRUD(ctx context.Context, handler handler) (cruder.Any, error) {
	_ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get announcement connection: %v", err)
	}

	defer conn.Close()

	cli := npool.NewManagerClient(conn)

	return handler(_ctx, cli)
}

func CreateAnnouncement(ctx context.Context, in *npool.AnnouncementReq) (*npool.Announcement, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateAnnouncement(ctx, &npool.CreateAnnouncementRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create announcement: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create announcement: %v", err)
	}
	return info.(*npool.Announcement), nil
}

func CreateAnnouncements(ctx context.Context, in []*npool.AnnouncementReq) ([]*npool.Announcement, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CreateAnnouncements(ctx, &npool.CreateAnnouncementsRequest{
			Infos: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail create announcements: %v", err)
		}
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail create announcements: %v", err)
	}
	return infos.([]*npool.Announcement), nil
}

func UpdateAnnouncement(ctx context.Context, in *npool.AnnouncementReq) (*npool.Announcement, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.UpdateAnnouncement(ctx, &npool.UpdateAnnouncementRequest{
			Info: in,
		})
		if err != nil {
			return nil, fmt.Errorf("fail add announcement: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail update announcement: %v", err)
	}
	return info.(*npool.Announcement), nil
}

func GetAnnouncement(ctx context.Context, id string) (*npool.Announcement, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetAnnouncement(ctx, &npool.GetAnnouncementRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get announcement: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get announcement: %v", err)
	}
	return info.(*npool.Announcement), nil
}

func GetAnnouncementOnly(ctx context.Context, conds *npool.Conds) (*npool.Announcement, error) {
	info, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetAnnouncementOnly(ctx, &npool.GetAnnouncementOnlyRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get announcement: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get announcement: %v", err)
	}
	return info.(*npool.Announcement), nil
}

func GetAnnouncements(ctx context.Context, conds *npool.Conds, offset, limit int32) ([]*npool.Announcement, uint32, error) {
	var total uint32
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.GetAnnouncements(ctx, &npool.GetAnnouncementsRequest{
			Conds:  conds,
			Limit:  limit,
			Offset: offset,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get announcements: %v", err)
		}
		total = resp.GetTotal()
		return resp.GetInfos(), nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get announcements: %v", err)
	}
	return infos.([]*npool.Announcement), total, nil
}

func ExistAnnouncement(ctx context.Context, id string) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistAnnouncement(ctx, &npool.ExistAnnouncementRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get announcement: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get announcement: %v", err)
	}
	return infos.(bool), nil
}

func ExistAnnouncementConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.ExistAnnouncementConds(ctx, &npool.ExistAnnouncementCondsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail get announcement: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return false, fmt.Errorf("fail get announcement: %v", err)
	}
	return infos.(bool), nil
}

func CountAnnouncements(ctx context.Context, conds *npool.Conds) (uint32, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.CountAnnouncements(ctx, &npool.CountAnnouncementsRequest{
			Conds: conds,
		})
		if err != nil {
			return nil, fmt.Errorf("fail count announcement: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return 0, fmt.Errorf("fail count announcement: %v", err)
	}
	return infos.(uint32), nil
}

func DeleteAnnouncement(ctx context.Context, id string) (*npool.Announcement, error) {
	infos, err := withCRUD(ctx, func(_ctx context.Context, cli npool.ManagerClient) (cruder.Any, error) {
		resp, err := cli.DeleteAnnouncement(ctx, &npool.DeleteAnnouncementRequest{
			ID: id,
		})
		if err != nil {
			return nil, fmt.Errorf("fail delete announcement: %v", err)
		}
		return resp.GetInfo(), nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail delete announcement: %v", err)
	}
	return infos.(*npool.Announcement), nil
}
