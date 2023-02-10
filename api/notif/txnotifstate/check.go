package txnotifstate

import (
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif/txnotifstate"

	"github.com/google/uuid"
)

func validate(in *npool.TxNotifStateReq) error {
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
		case npool.TxState_WaitSend:
		case npool.TxState_WaitTxSuccess:
		case npool.TxState_AlreadySend:
		default:
			logger.Sugar().Errorw("validate", "TxID", in.GetTxID(), "error", "invalid notif statu")
			return fmt.Errorf("invalid notif statu")
		}
	}
	if in.NotifType != nil {
		switch in.GetNotifType() {
		case npool.TxType_Withdraw:
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
		case uint32(npool.TxState_WaitTxSuccess):
		case uint32(npool.TxState_WaitSend):
		case uint32(npool.TxState_AlreadySend):
		default:
			logger.Sugar().Errorw("validate", "TxID", in.GetTxID(), "error", "invalid notif statu")
			return fmt.Errorf("invalid notif statu")
		}
	}
	if in.NotifType != nil {
		switch in.GetNotifType().GetValue() {
		case uint32(npool.TxType_Withdraw):
		default:
			logger.Sugar().Errorw("validate", "TxID", in.GetTxID(), "error", "invalid notif statu")
			return fmt.Errorf("invalid notif statu")
		}
	}
	return nil
}
