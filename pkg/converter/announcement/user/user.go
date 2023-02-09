package user

import (
	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/announcement/user"
	"github.com/NpoolPlatform/notif-manager/pkg/db/ent"
)

func Ent2Grpc(row *ent.ReadAnnouncement) *npool.User {
	if row == nil {
		return nil
	}
	return &npool.User{
		ID:             row.ID.String(),
		AppID:          row.AppID.String(),
		UserID:         row.UserID.String(),
		AnnouncementID: row.AnnouncementID.String(),
		CreatedAt:      row.CreatedAt,
		UpdatedAt:      row.UpdatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.ReadAnnouncement) []*npool.User {
	infos := []*npool.User{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
