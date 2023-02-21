package frontend

import (
	"context"
	"fmt"

	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/notif-manager/pkg/db/ent"
	"github.com/NpoolPlatform/notif-manager/pkg/testinit"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	"github.com/NpoolPlatform/message/npool"
	"github.com/NpoolPlatform/message/npool/notif/mgr/v1/template/frontend"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var entFrontendTemplate = ent.FrontendTemplate{
	ID:      uuid.New(),
	AppID:   uuid.New(),
	LangID:  uuid.New(),
	UsedFor: basetypes.UsedFor_KYCApproved.String(),
	Title:   uuid.NewString(),
	Content: uuid.NewString(),
}

var (
	id               = entFrontendTemplate.ID.String()
	appID            = entFrontendTemplate.AppID.String()
	langID           = entFrontendTemplate.LangID.String()
	usedFor          = basetypes.UsedFor_KYCApproved
	FrontendTemplate = frontend.FrontendTemplateReq{
		ID:      &id,
		AppID:   &appID,
		LangID:  &langID,
		UsedFor: &usedFor,
		Title:   &entFrontendTemplate.Title,
		Content: &entFrontendTemplate.Content,
	}
)

var info *ent.FrontendTemplate

func rowToObject(row *ent.FrontendTemplate) *ent.FrontendTemplate {
	return &ent.FrontendTemplate{
		ID:        row.ID,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
		AppID:     row.AppID,
		LangID:    row.LangID,
		UsedFor:   row.UsedFor,
		Title:     row.Title,
		Content:   row.Content,
	}
}

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &FrontendTemplate)
	if assert.Nil(t, err) {
		if assert.NotEqual(t, info.ID, uuid.UUID{}.String()) {
			entFrontendTemplate.ID = info.ID
			entFrontendTemplate.CreatedAt = info.CreatedAt
		}
		entFrontendTemplate.UpdatedAt = info.UpdatedAt
		entFrontendTemplate.CreatedAt = info.CreatedAt
		assert.Equal(t, rowToObject(info), &entFrontendTemplate)
	}
}

func createBulk(t *testing.T) {
	entApp := []ent.FrontendTemplate{
		{
			ID:      uuid.New(),
			AppID:   uuid.New(),
			LangID:  uuid.New(),
			UsedFor: basetypes.UsedFor_KYCApproved.String(),
			Title:   uuid.NewString(),
			Content: uuid.NewString(),
		},
		{
			ID:      uuid.New(),
			AppID:   uuid.New(),
			LangID:  uuid.New(),
			UsedFor: basetypes.UsedFor_KYCApproved.String(),
			Title:   uuid.NewString(),
			Content: uuid.NewString(),
		},
	}

	apps := []*frontend.FrontendTemplateReq{}
	for key := range entApp {
		id := entApp[key].ID.String()
		appID = entApp[key].AppID.String()
		langID := entApp[key].LangID.String()
		usedFor = basetypes.UsedFor_KYCApproved
		apps = append(apps, &frontend.FrontendTemplateReq{
			ID:      &id,
			AppID:   &appID,
			LangID:  &langID,
			UsedFor: &usedFor,
			Title:   &entApp[key].Title,
			Content: &entApp[key].Content,
		})
	}
	infos, err := CreateBulk(context.Background(), apps)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
		assert.NotEqual(t, infos[0].ID, uuid.UUID{}.String())
		assert.NotEqual(t, infos[1].ID, uuid.UUID{}.String())
	}
}

func update(t *testing.T) {
	var err error
	info, err = Update(context.Background(), &FrontendTemplate)
	if assert.Nil(t, err) {
		entFrontendTemplate.UpdatedAt = info.UpdatedAt
		entFrontendTemplate.CreatedAt = info.CreatedAt
		assert.Equal(t, rowToObject(info), &entFrontendTemplate)
	}
}

func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), info.ID)
	if assert.Nil(t, err) {
		entFrontendTemplate.UpdatedAt = info.UpdatedAt
		entFrontendTemplate.CreatedAt = info.CreatedAt
		assert.Equal(t, rowToObject(info), &entFrontendTemplate)
	}
}

func rows(t *testing.T) {
	infos, total, err := Rows(context.Background(),
		&frontend.Conds{
			ID: &npool.StringVal{
				Value: info.ID.String(),
				Op:    cruder.EQ,
			},
		}, 0, 0)
	if assert.Nil(t, err) {
		assert.Equal(t, total, 1)
		entFrontendTemplate.UpdatedAt = infos[0].UpdatedAt
		entFrontendTemplate.CreatedAt = infos[0].CreatedAt
		assert.Equal(t, rowToObject(infos[0]), &entFrontendTemplate)
	}
}

func rowOnly(t *testing.T) {
	var err error
	info, err = RowOnly(context.Background(),
		&frontend.Conds{
			ID: &npool.StringVal{
				Value: info.ID.String(),
				Op:    cruder.EQ,
			},
		})
	if assert.Nil(t, err) {
		entFrontendTemplate.UpdatedAt = info.UpdatedAt
		entFrontendTemplate.CreatedAt = info.CreatedAt
		assert.Equal(t, rowToObject(info), &entFrontendTemplate)
	}
}

func count(t *testing.T) {
	count, err := Count(context.Background(),
		&frontend.Conds{
			ID: &npool.StringVal{
				Value: info.ID.String(),
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, count, count)
	}
}

func exist(t *testing.T) {
	exist, err := Exist(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existConds(t *testing.T) {
	exist, err := ExistConds(context.Background(),
		&frontend.Conds{
			ID: &npool.StringVal{
				Value: info.ID.String(),
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func deleteT(t *testing.T) {
	info, err := Delete(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entFrontendTemplate)
	}
}

func TestMainOrder(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	t.Run("create", create)
	t.Run("createBulk", createBulk)
	t.Run("row", row)
	t.Run("rows", rows)
	t.Run("rowOnly", rowOnly)
	t.Run("update", update)
	t.Run("exist", exist)
	t.Run("existConds", existConds)
	t.Run("delete", deleteT)
	t.Run("count", count)
}
