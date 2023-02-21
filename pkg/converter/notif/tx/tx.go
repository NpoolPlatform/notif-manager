package tx

import (
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif/tx"
	"github.com/NpoolPlatform/notif-manager/pkg/db/ent"
)

func Ent2Grpc(row *ent.TxNotifState) *npool.Tx {
	if row == nil {
		return nil
	}
	return &npool.Tx{
		ID:         row.ID.String(),
		TxID:       row.TxID.String(),
		NotifState: npool.TxState(npool.TxState_value[row.NotifState]),
		TxType:     basetypes.TxType(basetypes.TxType_value[row.TxType]),
		CreatedAt:  row.CreatedAt,
		UpdatedAt:  row.UpdatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.TxNotifState) []*npool.Tx {
	var infos []*npool.Tx
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
