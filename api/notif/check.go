package notif

import (
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif"

	"github.com/google/uuid"
)

func validate(in *npool.NotifReq) error {
	if in.ID != nil {
		if _, err := uuid.Parse(in.GetID()); err != nil {
			logger.Sugar().Errorw("validate", "ID", in.GetID(), "error", err)
			return err
		}
	}
	if _, err := uuid.Parse(in.GetAppID()); err != nil {
		logger.Sugar().Errorw("validate", "AppID", in.GetAppID(), "error", err)
		return err
	}
	if _, err := uuid.Parse(in.GetUserID()); err != nil {
		logger.Sugar().Errorw("validate", "UserID", in.GetUserID(), "error", err)
		return err
	}
	if _, err := uuid.Parse(in.GetLangID()); err != nil {
		logger.Sugar().Errorw("validate", "LangID", in.GetLangID(), "error", err)
		return err
	}

	switch in.GetEventType() {
	case npool.EventType_KycReviewApproved:
	case npool.EventType_KycReviewRejected:
	case npool.EventType_WithdrawReviewApproved:
	case npool.EventType_WithdrawReviewRejected:
	case npool.EventType_WithdrawAddressReviewApproved:
	case npool.EventType_WithdrawAddressReviewRejected:
	default:
		return fmt.Errorf("EventType is invalid")
	}

	if in.GetTitle() == "" {
		logger.Sugar().Errorw("validate", "Title", in.GetTitle())
		return fmt.Errorf("title is invalid")
	}
	if in.GetContent() == "" {
		logger.Sugar().Errorw("validate", "Content", in.GetContent())
		return fmt.Errorf("title is invalid")
	}
	if len(in.GetChannels()) == 0 {
		logger.Sugar().Errorw("validate", "Channels", in.GetChannels())
		return fmt.Errorf("channels is invalid")
	}
	return nil
}

//nolint:gocyclo
func validateConds(in *npool.Conds) error {
	if in.ID != nil {
		if _, err := uuid.Parse(in.GetID().GetValue()); err != nil {
			logger.Sugar().Errorw("validateConds", "ID", in.GetID().GetValue(), "error", err)
			return err
		}
	}
	if in.AppID != nil {
		if _, err := uuid.Parse(in.GetAppID().GetValue()); err != nil {
			logger.Sugar().Errorw("validateConds", "AppID", in.GetAppID().GetValue(), "error", err)
			return err
		}
	}
	if in.UserID != nil {
		if _, err := uuid.Parse(in.GetUserID().GetValue()); err != nil {
			logger.Sugar().Errorw("validateConds", "UserID", in.GetUserID().GetValue(), "error", err)
			return err
		}
	}
	if in.LangID != nil {
		if _, err := uuid.Parse(in.GetLangID().GetValue()); err != nil {
			logger.Sugar().Errorw("validateConds", "LangID", in.GetLangID().GetValue(), "error", err)
			return err
		}
	}
	if in.EventType != nil {
		switch in.GetEventType().GetValue() {
		case uint32(npool.EventType_KycReviewApproved):
		case uint32(npool.EventType_KycReviewRejected):
		case uint32(npool.EventType_WithdrawReviewApproved):
		case uint32(npool.EventType_WithdrawReviewRejected):
		case uint32(npool.EventType_WithdrawAddressReviewApproved):
		case uint32(npool.EventType_WithdrawAddressReviewRejected):
		default:
			return fmt.Errorf("EventType is invalid")
		}
	}
	if in.Channels != nil {
		if len(in.GetChannels().GetValue()) == 0 {
			return fmt.Errorf("channels is invalid")
		}
	}
	return nil
}
