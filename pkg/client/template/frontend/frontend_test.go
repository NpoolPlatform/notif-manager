package frontend

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

	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/template/frontend"
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

var appDate = npool.FrontendTemplate{
	ID:      uuid.NewString(),
	AppID:   uuid.NewString(),
	LangID:  uuid.NewString(),
	UsedFor: basetypes.UsedFor_KYCApproved,
	Title:   uuid.NewString(),
	Content: uuid.NewString(),
}

var (
	appInfo = npool.FrontendTemplateReq{
		ID:      &appDate.ID,
		AppID:   &appDate.AppID,
		LangID:  &appDate.LangID,
		UsedFor: &appDate.UsedFor,
		Title:   &appDate.Title,
		Content: &appDate.Content,
	}
)

var info *npool.FrontendTemplate

func createFrontendTemplate(t *testing.T) {
	var err error
	info, err = CreateFrontendTemplate(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		appDate.UpdatedAt = info.UpdatedAt
		appDate.CreatedAt = info.CreatedAt
		assert.Equal(t, info, &appDate)
	}
}

func createFrontendTemplates(t *testing.T) {
	appDates := []npool.FrontendTemplate{
		{
			ID:      uuid.NewString(),
			AppID:   uuid.NewString(),
			LangID:  uuid.NewString(),
			UsedFor: basetypes.UsedFor_KYCApproved,
			Title:   uuid.NewString(),
			Content: uuid.NewString(),
		},
		{
			ID:      uuid.NewString(),
			AppID:   uuid.NewString(),
			LangID:  uuid.NewString(),
			UsedFor: basetypes.UsedFor_KYCApproved,
			Title:   uuid.NewString(),
			Content: uuid.NewString(),
		},
	}

	apps := []*npool.FrontendTemplateReq{}
	for key := range appDates {
		apps = append(apps, &npool.FrontendTemplateReq{
			ID:      &appDates[key].ID,
			AppID:   &appDates[key].AppID,
			LangID:  &appDates[key].LangID,
			UsedFor: &appDates[key].UsedFor,
			Title:   &appDates[key].Title,
			Content: &appDates[key].Content,
		})
	}

	infos, err := CreateFrontendTemplates(context.Background(), apps)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func updateFrontendTemplate(t *testing.T) {
	var err error
	info, err = UpdateFrontendTemplate(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		appDate.UpdatedAt = info.UpdatedAt
		appDate.CreatedAt = info.CreatedAt
		assert.Equal(t, info, &appDate)
	}
}

func getFrontendTemplate(t *testing.T) {
	var err error
	info, err = GetFrontendTemplate(context.Background(), info.ID)
	if assert.Nil(t, err) {
		appDate.UpdatedAt = info.UpdatedAt
		appDate.CreatedAt = info.CreatedAt
		assert.Equal(t, info, &appDate)
	}
}

func getFrontendTemplates(t *testing.T) {
	infos, total, err := GetFrontendTemplates(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		}, 0, 1)
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(1))
		appDate.UpdatedAt = infos[0].UpdatedAt
		appDate.CreatedAt = infos[0].CreatedAt
		assert.Equal(t, infos[0], &appDate)
	}
}

func getFrontendTemplateOnly(t *testing.T) {
	var err error
	info, err = GetFrontendTemplateOnly(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		})
	if assert.Nil(t, err) {
		appDate.UpdatedAt = info.UpdatedAt
		appDate.CreatedAt = info.CreatedAt
		assert.Equal(t, info, &appDate)
	}
}

func existFrontendTemplate(t *testing.T) {
	exist, err := ExistFrontendTemplate(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existFrontendTemplateConds(t *testing.T) {
	exist, err := ExistFrontendTemplateConds(context.Background(),
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

func deleteFrontendTemplate(t *testing.T) {
	info, err := DeleteFrontendTemplate(context.Background(), info.ID)
	if assert.Nil(t, err) {
		appDate.UpdatedAt = info.UpdatedAt
		appDate.CreatedAt = info.CreatedAt
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

	t.Run("createFrontendTemplate", createFrontendTemplate)
	t.Run("createFrontendTemplates", createFrontendTemplates)
	t.Run("getFrontendTemplate", getFrontendTemplate)
	t.Run("getFrontendTemplates", getFrontendTemplates)
	t.Run("getFrontendTemplateOnly", getFrontendTemplateOnly)
	t.Run("updateFrontendTemplate", updateFrontendTemplate)
	t.Run("existFrontendTemplate", existFrontendTemplate)
	t.Run("existFrontendTemplateConds", existFrontendTemplateConds)
	t.Run("delete", deleteFrontendTemplate)
}
