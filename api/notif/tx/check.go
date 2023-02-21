package tx

import (
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif/tx"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/google/uuid"
)

func validate(in *npool.TxReq) error {
	if in.ID != nil {
		if _, err := uuid.Parse(in.GetID()); err != nil {
			logger.Sugar().Errorw("validate", "ID", in.GetID(), "error", err)
			return err
		}
	}
	if in.TxID != nil {
		if _, err := uuid.Parse(in.GetTxID()); err != nil {
			logger.Sugar().Errorw("validate", "TxID", in.GetTxID(), "error", err)
			return err
		}
	}
	if in.NotifState != nil {
		switch in.GetNotifState() {
		case npool.TxState_WaitNotified:
		case npool.TxState_WaitSuccess:
		case npool.TxState_Notified:
		default:
			logger.Sugar().Errorw("validate", "TxID", in.GetTxID(), "error", "invalid notif statu")
			return fmt.Errorf("invalid notif statu")
		}
	}
	if in.TxType != nil {
		switch in.GetTxType() {
		case basetypes.TxType_TxWithdraw:
		default:
			logger.Sugar().Errorw("validate", "TxID", in.GetTxID(), "error", "invalid notif statu")
			return fmt.Errorf("invalid notif statu")
		}
	}
	return nil
}

func validateConds(in *npool.Conds) error {
	if in.ID != nil {
		if _, err := uuid.Parse(in.GetID().GetValue()); err != nil {
			logger.Sugar().Errorw("validateConds", "ID", in.GetID().GetValue(), "error", err)
			return err
		}
	}
	if in.TxID != nil {
		if _, err := uuid.Parse(in.GetTxID().GetValue()); err != nil {
			logger.Sugar().Errorw("validateConds", "TxID", in.GetTxID().GetValue(), "error", err)
			return err
		}
	}
	if in.NotifState != nil {
		switch in.GetNotifState().GetValue() {
		case uint32(npool.TxState_WaitSuccess):
		case uint32(npool.TxState_WaitNotified):
		case uint32(npool.TxState_Notified):
		default:
			logger.Sugar().Errorw("validate", "TxID", in.GetTxID(), "error", "invalid notif statu")
			return fmt.Errorf("invalid notif statu")
		}
	}
	if in.TxType != nil {
		switch in.GetTxType().GetValue() {
		case uint32(basetypes.TxType_TxWithdraw):
		default:
			logger.Sugar().Errorw("validate", "TxID", in.GetTxID(), "error", "invalid notif statu")
			return fmt.Errorf("invalid notif statu")
		}
	}
	return nil
}
