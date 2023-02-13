package notifchannel

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/message/npool/notif/mgr/v1/channel"
	"github.com/NpoolPlatform/message/npool/third/mgr/v1/usedfor"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	valuedef "github.com/NpoolPlatform/message/npool"

	"bou.ke/monkey"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/NpoolPlatform/notif-manager/pkg/testinit"

	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif/notifchannel"
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

var (
	eventType = usedfor.UsedFor_KYCApproved
	channel1  = channel.NotifChannel_ChannelFrontend
	data      = &npool.NotifChannel{
		ID:        uuid.NewString(),
		AppID:     uuid.NewString(),
		EventType: eventType,
		Channel:   channel1,
	}
)

var dataReq = &npool.NotifChannelReq{
	ID:        &data.ID,
	AppID:     &data.AppID,
	EventType: &data.EventType,
	Channel:   &data.Channel,
}

func createNotifChannel(t *testing.T) {
	info, err := CreateNotifChannel(context.Background(), dataReq)
	if assert.Nil(t, err) {
		data.CreatedAt = info.CreatedAt
		data.UpdatedAt = info.UpdatedAt
		assert.Equal(t, data, info)
	}
}

func createNotifChannels(t *testing.T) {
	datas := []npool.NotifChannel{
		{
			ID:        uuid.NewString(),
			AppID:     uuid.NewString(),
			EventType: eventType,
			Channel:   channel1,
		},
		{
			ID:        uuid.NewString(),
			AppID:     uuid.NewString(),
			EventType: eventType,
			Channel:   channel1,
		},
	}

	apps := []*npool.NotifChannelReq{}
	for key := range datas {
		apps = append(apps, &npool.NotifChannelReq{
			ID:        &datas[key].ID,
			AppID:     &datas[key].AppID,
			EventType: &datas[key].EventType,
			Channel:   &datas[key].Channel,
		})
	}

	infos, err := CreateNotifChannels(context.Background(), apps)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func getNotifChannel(t *testing.T) {
	info, err := GetNotifChannel(context.Background(), data.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, data, info)
	}
}

func getNotifChannels(t *testing.T) {
	infos, total, err := GetNotifChannels(context.Background(), &npool.Conds{
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

func getNotifChannelOnly(t *testing.T) {
	info, err := GetNotifChannelOnly(context.Background(), &npool.Conds{
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
	exist, err := ExistNotifChannel(context.Background(), data.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existAppGoodConds(t *testing.T) {
	exist, err := ExistNotifChannelConds(context.Background(), &npool.Conds{
		ID: &valuedef.StringVal{
			Op:    cruder.EQ,
			Value: data.ID,
		},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func deleteNotifChannel(t *testing.T) {
	info, err := DeleteNotifChannel(context.Background(), data.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, data, info)
	}

	_, err = GetNotifChannel(context.Background(), info.ID)
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

	t.Run("createNotifChannel", createNotifChannel)
	t.Run("createNotifChannels", createNotifChannels)
	t.Run("getNotifChannel", getNotifChannel)
	t.Run("getNotifChannels", getNotifChannels)
	t.Run("getNotifChannelOnly", getNotifChannelOnly)
	t.Run("existAppGood", existAppGood)
	t.Run("existAppGoodConds", existAppGoodConds)
	t.Run("deleteNotifChannel", deleteNotifChannel)
}
