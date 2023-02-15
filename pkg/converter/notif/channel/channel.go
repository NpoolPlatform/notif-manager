package channel

import (
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	"github.com/NpoolPlatform/message/npool/notif/mgr/v1/channel"
	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif/channel"
	"github.com/NpoolPlatform/notif-manager/pkg/db/ent"
)

func Ent2Grpc(row *ent.NotifChannel) *npool.NotifChannel {
	if row == nil {
		return nil
	}
	return &npool.NotifChannel{
		ID:        row.ID.String(),
		AppID:     row.AppID.String(),
		EventType: basetypes.UsedFor(basetypes.UsedFor_value[row.EventType]),
		Channel:   channel.NotifChannel(channel.NotifChannel_value[row.Channel]),
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.NotifChannel) []*npool.NotifChannel {
	var infos []*npool.NotifChannel
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
