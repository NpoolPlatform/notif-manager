package notif

import (
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif"

	channel "github.com/NpoolPlatform/message/npool/notif/mgr/v1/channel"
	usedfor "github.com/NpoolPlatform/message/npool/third/mgr/v1/usedfor"
	"github.com/google/uuid"
)

func validate(in *npool.NotifReq) error { //nolint
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
	case usedfor.UsedFor_WithdrawalRequest:
	case usedfor.UsedFor_WithdrawalCompleted:
	case usedfor.UsedFor_DepositReceived:
	case usedfor.UsedFor_KYCApproved:
	case usedfor.UsedFor_KYCRejected:
	case usedfor.UsedFor_Announcement:
	default:
		return fmt.Errorf("eventtype is invalid")
	}

	switch in.GetChannel() {
	case channel.NotifChannel_ChannelFrontend:
	case channel.NotifChannel_ChannelEmail:
	case channel.NotifChannel_ChannelSMS:
	default:
		logger.Sugar().Errorw("validate", "Channel", in.GetChannel(), "error", "invalid channel")
		return fmt.Errorf("channel is invalid")
	}

	if in.GetTitle() == "" {
		logger.Sugar().Errorw("validate", "Title", in.GetTitle())
		return fmt.Errorf("title is invalid")
	}
	if in.GetContent() == "" {
		logger.Sugar().Errorw("validate", "Content", in.GetContent())
		return fmt.Errorf("title is invalid")
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
		case uint32(usedfor.UsedFor_WithdrawalRequest):
		case uint32(usedfor.UsedFor_WithdrawalCompleted):
		case uint32(usedfor.UsedFor_DepositReceived):
		case uint32(usedfor.UsedFor_KYCApproved):
		case uint32(usedfor.UsedFor_KYCRejected):
		case uint32(usedfor.UsedFor_Announcement):
		default:
			return fmt.Errorf("eventtype is invalid")
		}
	}
	if in.Channel != nil {
		switch in.GetChannel().GetValue() {
		case uint32(channel.NotifChannel_ChannelFrontend):
		case uint32(channel.NotifChannel_ChannelEmail):
		case uint32(channel.NotifChannel_ChannelSMS):
		default:
			logger.Sugar().Errorw("validate", "Channel", in.GetChannel(), "error", "invalid channel")
			return fmt.Errorf("channel is invalid")
		}
	}
	return nil
}
