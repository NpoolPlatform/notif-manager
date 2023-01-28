package readstate

import (
	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/announcement/readstate"
	"github.com/NpoolPlatform/notif-manager/pkg/db/ent"
)

func Ent2Grpc(row *ent.ReadAnnouncement) *npool.ReadState {
	if row == nil {
		return nil
	}
	return &npool.ReadState{
		ID:             row.ID.String(),
		AppID:          row.AppID.String(),
		UserID:         row.UserID.String(),
		AnnouncementID: row.AnnouncementID.String(),
		CreatedAt:      row.CreatedAt,
		UpdatedAt:      row.UpdatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.ReadAnnouncement) []*npool.ReadState {
	infos := []*npool.ReadState{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
