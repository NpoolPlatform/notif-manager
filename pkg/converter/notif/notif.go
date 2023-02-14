package notif

import (
	"github.com/NpoolPlatform/message/npool/notif/mgr/v1/channel"
	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif"
	"github.com/NpoolPlatform/message/npool/third/mgr/v1/usedfor"
	"github.com/NpoolPlatform/notif-manager/pkg/db/ent"
)

func Ent2Grpc(row *ent.Notif) *npool.Notif {
	if row == nil {
		return nil
	}

	return &npool.Notif{
		ID:          row.ID.String(),
		AppID:       row.AppID.String(),
		UserID:      row.UserID.String(),
		Notified:    row.Notified,
		LangID:      row.LangID.String(),
		EventType:   usedfor.UsedFor(usedfor.UsedFor_value[row.EventType]),
		UseTemplate: row.UseTemplate,
		Title:       row.Title,
		Content:     row.Content,
		Channel:     channel.NotifChannel(channel.NotifChannel_value[row.Channel]),
		Extra:       row.Extra,
		CreatedAt:   row.CreatedAt,
		UpdatedAt:   row.UpdatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.Notif) []*npool.Notif {
	var infos []*npool.Notif
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
