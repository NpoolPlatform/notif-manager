package announcement

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	valuedef "github.com/NpoolPlatform/message/npool"
	"github.com/NpoolPlatform/message/npool/notif/mgr/v1/channel"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	"bou.ke/monkey"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/NpoolPlatform/notif-manager/pkg/testinit"

	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/announcement"
	"github.com/stretchr/testify/assert"

	"github.com/google/uuid"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var data = &npool.Announcement{
	ID:       uuid.NewString(),
	AppID:    uuid.NewString(),
	Title:    uuid.NewString(),
	Content:  uuid.NewString(),
	Channels: []channel.NotifChannel{channel.NotifChannel_ChannelEmail, channel.NotifChannel_ChannelSMS},
	EndAt:    99999999,
}

var dataReq = &npool.AnnouncementReq{
	ID:       &data.ID,
	AppID:    &data.AppID,
	Title:    &data.Title,
	Content:  &data.Content,
	Channels: data.Channels,
	EndAt:    &data.EndAt,
}

func createAnnouncement(t *testing.T) {
	info, err := CreateAnnouncement(context.Background(), dataReq)
	if assert.Nil(t, err) {
		data.CreatedAt = info.CreatedAt
		data.UpdatedAt = info.UpdatedAt
		assert.Equal(t, data, info)
	}
}

func updateAnnouncement(t *testing.T) {
	info, err := UpdateAnnouncement(context.Background(), dataReq)
	if assert.Nil(t, err) {
		data.UpdatedAt = info.UpdatedAt
		assert.Equal(t, data, info)
	}
}

func createAnnouncements(t *testing.T) {
	datas := []npool.Announcement{
		{
			ID:       uuid.NewString(),
			AppID:    uuid.NewString(),
			Title:    uuid.NewString(),
			Content:  uuid.NewString(),
			Channels: []channel.NotifChannel{channel.NotifChannel_ChannelEmail, channel.NotifChannel_ChannelSMS},
			EndAt:    99999999,
		},
		{
			ID:       uuid.NewString(),
			AppID:    uuid.NewString(),
			Title:    uuid.NewString(),
			Content:  uuid.NewString(),
			Channels: []channel.NotifChannel{channel.NotifChannel_ChannelEmail, channel.NotifChannel_ChannelSMS},
			EndAt:    99999999,
		},
	}

	apps := []*npool.AnnouncementReq{}
	for key := range datas {
		apps = append(apps, &npool.AnnouncementReq{
			ID:       &datas[key].ID,
			AppID:    &datas[key].AppID,
			Title:    &datas[key].Title,
			Content:  &datas[key].Content,
			Channels: datas[key].Channels,
			EndAt:    &datas[key].EndAt,
		})
	}

	infos, err := CreateAnnouncements(context.Background(), apps)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func getAnnouncement(t *testing.T) {
	info, err := GetAnnouncement(context.Background(), data.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, data, info)
	}
}

func getAnnouncements(t *testing.T) {
	infos, total, err := GetAnnouncements(context.Background(), &npool.Conds{
		ID: &valuedef.StringVal{
			Op:    cruder.EQ,
			Value: data.ID,
		},
	}, 0, 1)
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, infos[0], data)
	}
}

func getAnnouncementOnly(t *testing.T) {
	info, err := GetAnnouncementOnly(context.Background(), &npool.Conds{
		ID: &valuedef.StringVal{
			Op:    cruder.EQ,
			Value: data.ID,
		},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, info, data)
	}
}

func existAppGood(t *testing.T) {
	exist, err := ExistAnnouncement(context.Background(), data.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existAppGoodConds(t *testing.T) {
	exist, err := ExistAnnouncementConds(context.Background(), &npool.Conds{
		ID: &valuedef.StringVal{
			Op:    cruder.EQ,
			Value: data.ID,
		},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func deleteAnnouncement(t *testing.T) {
	info, err := DeleteAnnouncement(context.Background(), data.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, data, info)
	}

	_, err = GetAnnouncement(context.Background(), info.ID)
	assert.NotNil(t, err)
}

func TestClient(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createAnnouncement", createAnnouncement)
	t.Run("createAnnouncements", createAnnouncements)
	t.Run("getAnnouncement", getAnnouncement)
	t.Run("getAnnouncements", getAnnouncements)
	t.Run("getAnnouncementOnly", getAnnouncementOnly)
	t.Run("updateAnnouncement", updateAnnouncement)
	t.Run("existAppGood", existAppGood)
	t.Run("existAppGoodConds", existAppGoodConds)
	t.Run("deleteAnnouncement", deleteAnnouncement)
}
