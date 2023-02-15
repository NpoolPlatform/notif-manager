package sms

import (
	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/template/sms"
	"github.com/NpoolPlatform/notif-manager/pkg/db/ent"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
)

func Ent2Grpc(row *ent.SMSTemplate) *npool.SMSTemplate {
	if row == nil {
		return nil
	}

	return &npool.SMSTemplate{
		ID:      row.ID.String(),
		AppID:   row.AppID.String(),
		LangID:  row.LangID.String(),
		UsedFor: basetypes.UsedFor(basetypes.UsedFor_value[row.UsedFor]),
		Subject: row.Subject,
		Message: row.Message,
	}
}

func Ent2GrpcMany(rows []*ent.SMSTemplate) []*npool.SMSTemplate {
	infos := []*npool.SMSTemplate{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
