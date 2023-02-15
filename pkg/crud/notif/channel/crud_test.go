package channel

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	"github.com/NpoolPlatform/message/npool/notif/mgr/v1/channel"

	"github.com/NpoolPlatform/notif-manager/pkg/db/ent"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	valuedef "github.com/NpoolPlatform/message/npool"
	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif/channel"

	testinit "github.com/NpoolPlatform/notif-manager/pkg/testinit"
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

var (
	eventType = basetypes.UsedFor_KYCApproved
	channel1  = channel.NotifChannel_ChannelFrontend
	amt       = ent.NotifChannel{
		ID:        uuid.New(),
		AppID:     uuid.New(),
		EventType: eventType.String(),
		Channel:   channel1.String(),
	}
)

var (
	id    = amt.ID.String()
	appID = amt.AppID.String()
	req   = npool.NotifChannelReq{
		ID:        &id,
		AppID:     &appID,
		EventType: &eventType,
		Channel:   &channel1,
	}
)

var info *ent.NotifChannel

func create(t *testing.T) {
	var err error
	info, err = Create(context.Background(), &req)
	if assert.Nil(t, err) {
		amt.UpdatedAt = info.UpdatedAt
		amt.CreatedAt = info.CreatedAt
		assert.Equal(t, info.String(), amt.String())
	}
}

func createBulk(t *testing.T) {
	entities := []*ent.NotifChannel{
		{
			ID:        uuid.New(),
			AppID:     uuid.New(),
			EventType: eventType.String(),
			Channel:   channel1.String(),
		},
		{
			ID:        uuid.New(),
			AppID:     uuid.New(),
			EventType: eventType.String(),
			Channel:   channel1.String(),
		},
	}

	reqs := []*npool.NotifChannelReq{}
	for _, _amt := range entities {
		_id := _amt.ID.String()
		_appID := _amt.AppID.String()
		reqs = append(reqs, &npool.NotifChannelReq{
			ID:        &_id,
			AppID:     &_appID,
			EventType: &eventType,
			Channel:   &channel1,
		})
	}
	infos, err := CreateBulk(context.Background(), reqs)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func update(t *testing.T) {
	var err error
	info, err = Update(context.Background(), &req)
	if assert.Nil(t, err) {
		amt.UpdatedAt = info.UpdatedAt
		amt.CreatedAt = info.CreatedAt
		assert.Equal(t, info.String(), amt.String())
	}
}

func row(t *testing.T) {
	var err error
	info, err = Row(context.Background(), amt.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info.String(), amt.String())
	}
}

func rows(t *testing.T) {
	infos, total, err := Rows(context.Background(),
		&npool.Conds{
			ID: &valuedef.StringVal{
				Value: id,
				Op:    cruder.EQ,
			},
		}, 0, 0)
	if assert.Nil(t, err) {
		if assert.Equal(t, total, 1) {
			assert.Equal(t, infos[0].String(), amt.String())
		}
	}
}

func rowOnly(t *testing.T) {
	var err error
	info, err = RowOnly(context.Background(),
		&npool.Conds{
			ID: &valuedef.StringVal{
				Value: id,
				Op:    cruder.EQ,
			},
		})
	if assert.Nil(t, err) {
		assert.Equal(t, info.String(), amt.String())
	}
}

func count(t *testing.T) {
	count, err := Count(context.Background(),
		&npool.Conds{
			ID: &valuedef.StringVal{
				Value: id,
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, count, uint32(1))
	}
}

func exist(t *testing.T) {
	exist, err := Exist(context.Background(), amt.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existConds(t *testing.T) {
	exist, err := ExistConds(context.Background(),
		&npool.Conds{
			ID: &valuedef.StringVal{
				Value: id,
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func deleteA(t *testing.T) {
	info, err := Delete(context.Background(), amt.ID)
	if assert.Nil(t, err) {
		amt.DeletedAt = info.DeletedAt
		amt.UpdatedAt = info.UpdatedAt
		assert.Equal(t, info.String(), amt.String())
	}
}

func TestDetail(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	t.Run("create", create)
	t.Run("createBulk", createBulk)
	t.Run("update", update)
	t.Run("row", row)
	t.Run("rows", rows)
	t.Run("rowOnly", rowOnly)
	t.Run("exist", exist)
	t.Run("existConds", existConds)
	t.Run("count", count)
	t.Run("delete", deleteA)
}
