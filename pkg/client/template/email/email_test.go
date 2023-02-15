package email

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

	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/template/email"
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

var appDate = npool.EmailTemplate{
	ID:                uuid.NewString(),
	AppID:             uuid.NewString(),
	LangID:            uuid.NewString(),
	UsedFor:           basetypes.UsedFor_Signin,
	Sender:            uuid.NewString(),
	ReplyTos:          []string{uuid.NewString()},
	CCTos:             []string{uuid.NewString()},
	Subject:           uuid.NewString(),
	Body:              uuid.NewString(),
	DefaultToUsername: uuid.NewString(),
}

var (
	appInfo = npool.EmailTemplateReq{
		ID:                &appDate.ID,
		AppID:             &appDate.AppID,
		LangID:            &appDate.LangID,
		UsedFor:           &appDate.UsedFor,
		Sender:            &appDate.Sender,
		ReplyTos:          appDate.ReplyTos,
		CCTos:             appDate.CCTos,
		Subject:           &appDate.Subject,
		Body:              &appDate.Body,
		DefaultToUsername: &appDate.DefaultToUsername,
	}
)

var info *npool.EmailTemplate

func createEmailTemplate(t *testing.T) {
	var err error
	info, err = CreateEmailTemplate(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appDate)
	}
}

func createEmailTemplates(t *testing.T) {
	appDates := []npool.EmailTemplate{
		{
			ID:                uuid.NewString(),
			AppID:             uuid.NewString(),
			LangID:            uuid.NewString(),
			UsedFor:           basetypes.UsedFor_Signin,
			Sender:            uuid.NewString(),
			ReplyTos:          []string{uuid.NewString()},
			CCTos:             []string{uuid.NewString()},
			Subject:           uuid.NewString(),
			Body:              uuid.NewString(),
			DefaultToUsername: uuid.NewString(),
		},
		{
			ID:                uuid.NewString(),
			AppID:             uuid.NewString(),
			LangID:            uuid.NewString(),
			UsedFor:           basetypes.UsedFor_Signin,
			Sender:            uuid.NewString(),
			ReplyTos:          []string{uuid.NewString()},
			CCTos:             []string{uuid.NewString()},
			Subject:           uuid.NewString(),
			Body:              uuid.NewString(),
			DefaultToUsername: uuid.NewString(),
		},
	}

	apps := []*npool.EmailTemplateReq{}
	for key := range appDates {
		apps = append(apps, &npool.EmailTemplateReq{
			ID:                &appDates[key].ID,
			AppID:             &appDates[key].AppID,
			LangID:            &appDates[key].LangID,
			UsedFor:           &appDates[key].UsedFor,
			Sender:            &appDates[key].Sender,
			ReplyTos:          appDates[key].ReplyTos,
			CCTos:             appDates[key].CCTos,
			Subject:           &appDates[key].Subject,
			Body:              &appDates[key].Body,
			DefaultToUsername: &appDates[key].DefaultToUsername,
		})
	}

	infos, err := CreateEmailTemplates(context.Background(), apps)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func updateEmailTemplate(t *testing.T) {
	var err error
	info, err = UpdateEmailTemplate(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appDate)
	}
}

func getEmailTemplate(t *testing.T) {
	var err error
	info, err = GetEmailTemplate(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appDate)
	}
}

func getEmailTemplates(t *testing.T) {
	infos, total, err := GetEmailTemplates(context.Background(),
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

func getEmailTemplateOnly(t *testing.T) {
	var err error
	info, err = GetEmailTemplateOnly(context.Background(),
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

func existEmailTemplate(t *testing.T) {
	exist, err := ExistEmailTemplate(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existEmailTemplateConds(t *testing.T) {
	exist, err := ExistEmailTemplateConds(context.Background(),
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

func deleteEmailTemplate(t *testing.T) {
	info, err := DeleteEmailTemplate(context.Background(), info.ID)
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

	t.Run("createEmailTemplate", createEmailTemplate)
	t.Run("createEmailTemplates", createEmailTemplates)
	t.Run("getEmailTemplate", getEmailTemplate)
	t.Run("getEmailTemplates", getEmailTemplates)
	t.Run("getEmailTemplateOnly", getEmailTemplateOnly)
	t.Run("updateEmailTemplate", updateEmailTemplate)
	t.Run("existEmailTemplate", existEmailTemplate)
	t.Run("existEmailTemplateConds", existEmailTemplateConds)
	t.Run("delete", deleteEmailTemplate)
}
