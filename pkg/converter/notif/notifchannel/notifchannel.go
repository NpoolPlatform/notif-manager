package notifchannel

import (
	"github.com/NpoolPlatform/message/npool/notif/mgr/v1/channel"
	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif/notifchannel"
	"github.com/NpoolPlatform/message/npool/third/mgr/v1/usedfor"
	"github.com/NpoolPlatform/notif-manager/pkg/db/ent"
)

func Ent2Grpc(row *ent.NotifChannel) *npool.NotifChannel {
	if row == nil {
		return nil
	}
	return &npool.NotifChannel{
		ID:        row.ID.String(),
		AppID:     row.AppID.String(),
		EventType: usedfor.UsedFor(usedfor.UsedFor_value[row.EventType]),
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
