package email

import (
	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/template/email"
	"github.com/NpoolPlatform/notif-manager/pkg/db/ent"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
)

func Ent2Grpc(row *ent.EmailTemplate) *npool.EmailTemplate {
	if row == nil {
		return nil
	}

	return &npool.EmailTemplate{
		ID:                row.ID.String(),
		AppID:             row.AppID.String(),
		LangID:            row.LangID.String(),
		UsedFor:           basetypes.UsedFor(basetypes.UsedFor_value[row.UsedFor]),
		Sender:            row.Sender,
		ReplyTos:          row.ReplyTos,
		CCTos:             row.CcTos,
		Subject:           row.Subject,
		Body:              row.Body,
		DefaultToUsername: row.DefaultToUsername,
	}
}

func Ent2GrpcMany(rows []*ent.EmailTemplate) []*npool.EmailTemplate {
	infos := []*npool.EmailTemplate{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
