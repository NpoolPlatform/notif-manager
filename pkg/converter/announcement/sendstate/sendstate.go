package sendstate

import (
	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/announcement/sendstate"
	"github.com/NpoolPlatform/message/npool/notif/mgr/v1/channel"
	"github.com/NpoolPlatform/notif-manager/pkg/db/ent"
)

func Ent2Grpc(row *ent.SendAnnouncement) *npool.SendState {
	if row == nil {
		return nil
	}
	return &npool.SendState{
		ID:             row.ID.String(),
		AppID:          row.AppID.String(),
		UserID:         row.UserID.String(),
		AnnouncementID: row.AnnouncementID.String(),
		Channel:        channel.NotifChannel(channel.NotifChannel_value[row.Channel]),
		CreatedAt:      row.CreatedAt,
		UpdatedAt:      row.UpdatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.SendAnnouncement) []*npool.SendState {
	infos := []*npool.SendState{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
