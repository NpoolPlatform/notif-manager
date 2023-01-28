package notif

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/notif-manager/pkg/db/ent"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	valuedef "github.com/NpoolPlatform/message/npool"
	"github.com/NpoolPlatform/message/npool/notif/mgr/v1/channel"
	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif"

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

var amt = ent.Notif{
	ID:          uuid.New(),
	AppID:       uuid.New(),
	UserID:      uuid.New(),
	AlreadyRead: true,
	LangID:      uuid.New(),
	EventType:   "DefaultEventType",
	UseTemplate: true,
	Title:       uuid.NewString(),
	Content:     uuid.NewString(),
	Channels:    []string{channel.NotifChannel_ChannelEmail.String(), channel.NotifChannel_ChannelSMS.String()},
	EmailSend:   true,
}

var (
	id        = amt.ID.String()
	appID     = amt.AppID.String()
	userID    = amt.UserID.String()
	langID    = amt.LangID.String()
	eventType = npool.EventType(npool.EventType_value[amt.EventType])
	channel1  = []channel.NotifChannel{channel.NotifChannel_ChannelEmail, channel.NotifChannel_ChannelSMS}
	req       = npool.NotifReq{
		ID:          &id,
		AppID:       &appID,
		UserID:      &userID,
		AlreadyRead: &amt.AlreadyRead,
		LangID:      &langID,
		EventType:   &eventType,
		UseTemplate: &amt.UseTemplate,
		Title:       &amt.Title,
		Content:     &amt.Content,
		Channels:    channel1,
		EmailSend:   &amt.EmailSend,
	}
)

var info *ent.Notif

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
	entities := []*ent.Notif{
		{
			ID:          uuid.New(),
			AppID:       uuid.New(),
			UserID:      uuid.New(),
			AlreadyRead: true,
			LangID:      uuid.New(),
			EventType:   "DefaultEventType",
			UseTemplate: true,
			Title:       uuid.NewString(),
			Content:     uuid.NewString(),
			Channels:    []string{channel.NotifChannel_ChannelEmail.String(), channel.NotifChannel_ChannelSMS.String()},
			EmailSend:   true,
		},
		{
			ID:          uuid.New(),
			AppID:       uuid.New(),
			UserID:      uuid.New(),
			AlreadyRead: true,
			LangID:      uuid.New(),
			EventType:   "DefaultEventType",
			UseTemplate: true,
			Title:       uuid.NewString(),
			Content:     uuid.NewString(),
			Channels:    []string{channel.NotifChannel_ChannelEmail.String(), channel.NotifChannel_ChannelSMS.String()},
			EmailSend:   true,
		},
	}

	reqs := []*npool.NotifReq{}
	for _, _amt := range entities {
		_id := _amt.ID.String()
		_appID := _amt.AppID.String()
		_userID := _amt.UserID.String()
		_langID := _amt.LangID.String()
		_eventType := npool.EventType(npool.EventType_value[_amt.EventType])
		_channel1 := []channel.NotifChannel{channel.NotifChannel_ChannelEmail, channel.NotifChannel_ChannelSMS}
		reqs = append(reqs, &npool.NotifReq{
			ID:          &_id,
			AppID:       &_appID,
			UserID:      &_userID,
			AlreadyRead: &_amt.AlreadyRead,
			LangID:      &_langID,
			EventType:   &_eventType,
			UseTemplate: &_amt.UseTemplate,
			Title:       &_amt.Title,
			Content:     &_amt.Content,
			Channels:    _channel1,
			EmailSend:   &_amt.EmailSend,
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
