package notif

import (
	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif/txnotifstate"
	"github.com/NpoolPlatform/notif-manager/pkg/db/ent"
)

func Ent2Grpc(row *ent.TxNotifState) *npool.TxNotifState {
	if row == nil {
		return nil
	}
	return &npool.TxNotifState{
		ID:         row.ID.String(),
		TxID:       row.TxID.String(),
		NotifState: npool.TxState(npool.TxState_value[row.NotifState]),
		NotifType:  npool.TxType(npool.TxType_value[row.NotifType]),
		CreatedAt:  row.CreatedAt,
		UpdatedAt:  row.UpdatedAt,
	}
}

func Ent2GrpcMany(rows []*ent.TxNotifState) []*npool.TxNotifState {
	var infos []*npool.TxNotifState
	for _, row := range rows {
		infos = append(infos, Ent2Grpc(row))
	}
	return infos
}
