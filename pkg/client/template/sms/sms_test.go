package sms

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"google.golang.org/grpc/credentials/insecure"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	"bou.ke/monkey"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	val "github.com/NpoolPlatform/message/npool"

	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/template/sms"
	testinit "github.com/NpoolPlatform/notif-manager/pkg/testinit"
	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var appDate = npool.SMSTemplate{
	ID:      uuid.NewString(),
	AppID:   uuid.NewString(),
	LangID:  uuid.NewString(),
	UsedFor: basetypes.UsedFor_Signin,
	Subject: uuid.NewString(),
	Message: uuid.NewString(),
}

var (
	appInfo = npool.SMSTemplateReq{
		ID:      &appDate.ID,
		AppID:   &appDate.AppID,
		LangID:  &appDate.LangID,
		UsedFor: &appDate.UsedFor,
		Subject: &appDate.Subject,
		Message: &appDate.Message,
	}
)

var info *npool.SMSTemplate

func createSMSTemplate(t *testing.T) {
	var err error
	info, err = CreateSMSTemplate(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appDate)
	}
}

func createSMSTemplates(t *testing.T) {
	appDates := []npool.SMSTemplate{
		{
			ID:      uuid.NewString(),
			AppID:   uuid.NewString(),
			LangID:  uuid.NewString(),
			UsedFor: basetypes.UsedFor_Signin,
			Subject: uuid.NewString(),
			Message: uuid.NewString(),
		},
		{
			ID:      uuid.NewString(),
			AppID:   uuid.NewString(),
			LangID:  uuid.NewString(),
			UsedFor: basetypes.UsedFor_Signin,
			Subject: uuid.NewString(),
			Message: uuid.NewString(),
		},
	}

	apps := []*npool.SMSTemplateReq{}
	for key := range appDates {
		apps = append(apps, &npool.SMSTemplateReq{
			ID:      &appDates[key].ID,
			AppID:   &appDates[key].AppID,
			LangID:  &appDates[key].LangID,
			UsedFor: &appDates[key].UsedFor,
			Subject: &appDates[key].Subject,
			Message: &appDates[key].Message,
		})
	}

	infos, err := CreateSMSTemplates(context.Background(), apps)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func updateSMSTemplate(t *testing.T) {
	var err error
	info, err = UpdateSMSTemplate(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appDate)
	}
}

func getSMSTemplate(t *testing.T) {
	var err error
	info, err = GetSMSTemplate(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appDate)
	}
}

func getSMSTemplates(t *testing.T) {
	infos, total, err := GetSMSTemplates(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		}, 0, 1)
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, infos[0], &appDate)
	}
}

func getSMSTemplateOnly(t *testing.T) {
	var err error
	info, err = GetSMSTemplateOnly(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		})
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appDate)
	}
}

func existSMSTemplate(t *testing.T) {
	exist, err := ExistSMSTemplate(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existSMSTemplateConds(t *testing.T) {
	exist, err := ExistSMSTemplateConds(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func deleteSMSTemplate(t *testing.T) {
	info, err := DeleteSMSTemplate(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appDate)
	}
}

func TestMainOrder(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createSMSTemplate", createSMSTemplate)
	t.Run("createSMSTemplates", createSMSTemplates)
	t.Run("getSMSTemplate", getSMSTemplate)
	t.Run("getSMSTemplates", getSMSTemplates)
	t.Run("getSMSTemplateOnly", getSMSTemplateOnly)
	t.Run("updateSMSTemplate", updateSMSTemplate)
	t.Run("existSMSTemplate", existSMSTemplate)
	t.Run("existSMSTemplateConds", existSMSTemplateConds)
	t.Run("delete", deleteSMSTemplate)
}
