package sms

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
	"github.com/NpoolPlatform/message/npool/notif/mgr/v1/template/sms"

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

var entSMSTemplate = ent.SMSTemplate{
	ID:      uuid.New(),
	AppID:   uuid.New(),
	LangID:  uuid.New(),
	UsedFor: basetypes.UsedFor_Signin.String(),
	Subject: uuid.NewString(),
	Message: uuid.NewString(),
}

var (
	id             = entSMSTemplate.ID.String()
	appID          = entSMSTemplate.AppID.String()
	langID         = entSMSTemplate.LangID.String()
	usedFor        = basetypes.UsedFor_Signin
	appSMSTemplate = sms.SMSTemplateReq{
		ID:      &id,
		AppID:   &appID,
		LangID:  &langID,
		UsedFor: &usedFor,
		Subject: &entSMSTemplate.Subject,
		Message: &entSMSTemplate.Message,
	}
)

var info *ent.SMSTemplate

func rowToObject(row *ent.SMSTemplate) *ent.SMSTemplate {
	return &ent.SMSTemplate{
		ID:        row.ID,
		CreatedAt: row.CreatedAt,
		AppID:     row.AppID,
		LangID:    row.LangID,
		UsedFor:   row.UsedFor,
		Subject:   row.Subject,
		Message:   row.Message,
	}
}

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &appSMSTemplate)
	if assert.Nil(t, err) {
		if assert.NotEqual(t, info.ID, uuid.UUID{}.String()) {
			entSMSTemplate.ID = info.ID
			entSMSTemplate.CreatedAt = info.CreatedAt
		}
		assert.Equal(t, rowToObject(info), &entSMSTemplate)
	}
}

func createBulk(t *testing.T) {
	entApp := []ent.SMSTemplate{
		{
			ID:      uuid.New(),
			AppID:   uuid.New(),
			LangID:  uuid.New(),
			UsedFor: basetypes.UsedFor_Signin.String(),
			Subject: uuid.NewString(),
			Message: uuid.NewString(),
		},
		{
			ID:      uuid.New(),
			AppID:   uuid.New(),
			LangID:  uuid.New(),
			UsedFor: basetypes.UsedFor_Signin.String(),
			Subject: uuid.NewString(),
			Message: uuid.NewString(),
		},
	}

	apps := []*sms.SMSTemplateReq{}
	for key := range entApp {
		id := entApp[key].ID.String()
		appID = entApp[key].AppID.String()
		langID := entApp[key].LangID.String()
		usedFor = basetypes.UsedFor_Signin
		apps = append(apps, &sms.SMSTemplateReq{
			ID:      &id,
			AppID:   &appID,
			LangID:  &langID,
			UsedFor: &usedFor,
			Subject: &entApp[key].Subject,
			Message: &entApp[key].Message,
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
	info, err = Update(context.Background(), &appSMSTemplate)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entSMSTemplate)
	}
}

func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entSMSTemplate)
	}
}

func rows(t *testing.T) {
	infos, total, err := Rows(context.Background(),
		&sms.Conds{
			ID: &npool.StringVal{
				Value: info.ID.String(),
				Op:    cruder.EQ,
			},
		}, 0, 0)
	if assert.Nil(t, err) {
		assert.Equal(t, total, 1)
		assert.Equal(t, rowToObject(infos[0]), &entSMSTemplate)
	}
}

func rowOnly(t *testing.T) {
	var err error
	info, err = RowOnly(context.Background(),
		&sms.Conds{
			ID: &npool.StringVal{
				Value: info.ID.String(),
				Op:    cruder.EQ,
			},
		})
	if assert.Nil(t, err) {
		assert.Equal(t, rowToObject(info), &entSMSTemplate)
	}
}

func count(t *testing.T) {
	count, err := Count(context.Background(),
		&sms.Conds{
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
		&sms.Conds{
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
		assert.Equal(t, rowToObject(info), &entSMSTemplate)
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
