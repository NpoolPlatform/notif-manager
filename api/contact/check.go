package contact

import (
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/contact"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func validate(info *npool.ContactReq) error {
	if info.AppID == nil {
		logger.Sugar().Errorw("validate", "AppID", info.AppID)
		return status.Error(codes.InvalidArgument, "AppID is empty")
	}

	if _, err := uuid.Parse(info.GetAppID()); err != nil {
		logger.Sugar().Errorw("validate", "AppID", info.GetAppID(), "error", err)
		return status.Error(codes.InvalidArgument, "AppID is invalid")
	}

	if info.UsedFor == nil {
		logger.Sugar().Errorw("validate", "UsedFor", info.UsedFor, "GetUsedFor", info.GetUsedFor())
		return status.Error(codes.InvalidArgument, "UsedFor is empty")
	}

	if info.Account == nil || info.GetAccount() == "" {
		logger.Sugar().Errorw("validate", "Account", info.Account, "GetAccount", info.GetAccount())
		return status.Error(codes.InvalidArgument, "Account is empty")
	}

	if info.AccountType == nil {
		logger.Sugar().Errorw("validate", "AccountType", info.AccountType, "GetAccountType", info.GetAccountType())
		return status.Error(codes.InvalidArgument, "AccountType is empty")
	}

	switch info.GetAccountType() {
	case basetypes.SignMethod_Email:
	case basetypes.SignMethod_Mobile:
	default:
		logger.Sugar().Errorw("validate", "AccountType", info.GetAccountType(), "GetAccountType", info.GetAccountType())
		return status.Error(codes.InvalidArgument, "AccountType is invalid")
	}

	if info.Sender == nil || info.GetSender() == "" {
		logger.Sugar().Errorw("validate", "Sender", info.Sender, "GetSender", info.GetSender())
		return status.Error(codes.InvalidArgument, "Sender is empty")
	}

	return nil
}

func Validate(info *npool.ContactReq) error {
	return validate(info)
}
