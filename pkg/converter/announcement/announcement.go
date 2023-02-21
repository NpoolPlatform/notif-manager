package announcement

import (
	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/announcement"
	"github.com/NpoolPlatform/message/npool/notif/mgr/v1/channel"
	"github.com/NpoolPlatform/notif-manager/pkg/db/ent"
)

func Ent2Grpc(row *ent.Announcement) *npool.Announcement {
	if row == nil {
		return nil
	}

	return &npool.Announcement{
		ID:               row.ID.String(),
		AppID:            row.AppID.String(),
		LangID:           row.LangID.String(),
		Title:            row.Title,
		Content:          row.Content,
		Channel:          channel.NotifChannel(channel.NotifChannel_value[row.Channel]),
		EndAt:            row.EndAt,
		AnnouncementType: npool.AnnouncementType(npool.AnnouncementType_value[row.Type]),
		CreatedAt:        row.CreatedAt,
		UpdatedAt:        row.UpdatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.Announcement) []*npool.Announcement {
	infos := []*npool.Announcement{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
