package channel

import (
	"fmt"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif/channel"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	"github.com/NpoolPlatform/message/npool/notif/mgr/v1/channel"

	"github.com/google/uuid"
)

func validate(in *npool.ChannelReq) error {
	if in.ID != nil {
		if _, err := uuid.Parse(in.GetID()); err != nil {
			logger.Sugar().Errorw("validate", "ID", in.GetID(), "error", err)
			return err
		}
	}
	if in.AppID != nil {
		if _, err := uuid.Parse(in.GetAppID()); err != nil {
			logger.Sugar().Errorw("validate", "AppID", in.GetAppID(), "error", err)
			return err
		}
	}
	switch in.GetEventType() {
	case basetypes.UsedFor_WithdrawalRequest:
	case basetypes.UsedFor_WithdrawalCompleted:
	case basetypes.UsedFor_DepositReceived:
	case basetypes.UsedFor_KYCApproved:
	case basetypes.UsedFor_KYCRejected:
	case basetypes.UsedFor_Announcement:
	default:
		return fmt.Errorf("EventType is invalid")
	}
	if in.Channel != nil {
		switch in.GetChannel() {
		case channel.NotifChannel_ChannelFrontend:
		case channel.NotifChannel_ChannelEmail:
		case channel.NotifChannel_ChannelSMS:
		default:
			logger.Sugar().Errorw("validate", "Channel", in.GetChannel(), "error", "invalid channel")
			return fmt.Errorf("invalid channel")
		}
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
	if in.EventType != nil {
		switch in.GetEventType().GetValue() {
		case uint32(basetypes.UsedFor_WithdrawalRequest):
		case uint32(basetypes.UsedFor_WithdrawalCompleted):
		case uint32(basetypes.UsedFor_DepositReceived):
		case uint32(basetypes.UsedFor_KYCApproved):
		case uint32(basetypes.UsedFor_KYCRejected):
		case uint32(basetypes.UsedFor_Announcement):
		default:
			return fmt.Errorf("EventType is invalid")
		}
	}
	for _, typ := range in.GetEventTypes().GetValue() {
		switch typ {
		case uint32(basetypes.UsedFor_WithdrawalRequest):
		case uint32(basetypes.UsedFor_WithdrawalCompleted):
		case uint32(basetypes.UsedFor_DepositReceived):
		case uint32(basetypes.UsedFor_KYCApproved):
		case uint32(basetypes.UsedFor_KYCRejected):
		case uint32(basetypes.UsedFor_Announcement):
		default:
			return fmt.Errorf("EventType is invalid")
		}
	}
	if in.Channel != nil {
		switch in.GetChannel().GetValue() {
		case uint32(channel.NotifChannel_ChannelFrontend):
		case uint32(channel.NotifChannel_ChannelEmail):
		case uint32(channel.NotifChannel_ChannelSMS):
		default:
			logger.Sugar().Errorw("validate", "channel", in.GetChannel(), "error", "invalid notif channel")
			return fmt.Errorf("invalid channel")
		}
	}
	for _, ch := range in.GetChannels().GetValue() {
		switch ch {
		case uint32(channel.NotifChannel_ChannelFrontend):
		case uint32(channel.NotifChannel_ChannelEmail):
		case uint32(channel.NotifChannel_ChannelSMS):
		default:
			logger.Sugar().Errorw("validate", "channel", in.GetChannel(), "error", "invalid notif channel")
			return fmt.Errorf("invalid channel")
		}
	}
	return nil
}
