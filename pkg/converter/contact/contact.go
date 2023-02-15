package contact

import (
	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/contact"
	"github.com/NpoolPlatform/notif-manager/pkg/db/ent"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
)

func Ent2Grpc(row *ent.Contact) *npool.Contact {
	if row == nil {
		return nil
	}

	return &npool.Contact{
		ID:          row.ID.String(),
		AppID:       row.AppID.String(),
		UsedFor:     basetypes.UsedFor(basetypes.UsedFor_value[row.UsedFor]),
		Account:     row.Account,
		AccountType: basetypes.SignMethod(basetypes.SignMethod_value[row.AccountType]),
		Sender:      row.Sender,
	}
}

func Ent2GrpcMany(rows []*ent.Contact) []*npool.Contact {
	infos := []*npool.Contact{}
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
