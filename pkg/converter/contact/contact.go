package contact

import (
	"github.com/NpoolPlatform/message/npool/appuser/mgr/v2/signmethod"
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
		UsedFor:     usedfor.UsedFor(usedfor.UsedFor_value[row.UsedFor]),
		Account:     row.Account,
		AccountType: signmethod.SignMethodType(signmethod.SignMethodType_value[row.AccountType]),
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
