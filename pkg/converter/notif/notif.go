package notif

import (
	"github.com/NpoolPlatform/message/npool/notif/mgr/v1/channel"
	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif"
	"github.com/NpoolPlatform/notif-manager/pkg/db/ent"
)

func Ent2Grpc(row *ent.Notif) *npool.Notif {
	if row == nil {
		return nil
	}

	channels := []channel.NotifChannel{}
	for _, val := range row.Channels {
		channels = append(channels, channel.NotifChannel(channel.NotifChannel_value[val]))
	}
	return &npool.Notif{
		ID:          row.ID.String(),
		AppID:       row.AppID.String(),
		UserID:      row.UserID.String(),
		AlreadyRead: row.AlreadyRead,
		LangID:      row.LangID.String(),
		EventType:   npool.EventType(npool.EventType_value[row.EventType]),
		UseTemplate: row.UseTemplate,
		Title:       row.Title,
		Content:     row.Content,
		Channels:    channels,
		EmailSend:   row.EmailSend,
		CreatedAt:   row.CreatedAt,
		UpdatedAt:   row.UpdatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.Notif) []*npool.Notif {
	infos := []*npool.Notif{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
