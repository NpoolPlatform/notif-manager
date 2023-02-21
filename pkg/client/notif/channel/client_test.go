package channel

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	"github.com/NpoolPlatform/message/npool/notif/mgr/v1/channel"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	valuedef "github.com/NpoolPlatform/message/npool"

	"bou.ke/monkey"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/NpoolPlatform/notif-manager/pkg/testinit"

	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif/channel"
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
	eventType = basetypes.UsedFor_KYCApproved
	channel1  = channel.NotifChannel_ChannelFrontend
	data      = &npool.Channel{
		ID:        uuid.NewString(),
		AppID:     uuid.NewString(),
		EventType: eventType,
		Channel:   channel1,
	}
)

var dataReq = &npool.ChannelReq{
	ID:        &data.ID,
	AppID:     &data.AppID,
	EventType: &data.EventType,
	Channel:   &data.Channel,
}

func createChannel(t *testing.T) {
	info, err := CreateChannel(context.Background(), dataReq)
	if assert.Nil(t, err) {
		data.CreatedAt = info.CreatedAt
		data.UpdatedAt = info.UpdatedAt
		assert.Equal(t, data, info)
	}
}

func createChannels(t *testing.T) {
	datas := []npool.Channel{
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

	apps := []*npool.ChannelReq{}
	for key := range datas {
		apps = append(apps, &npool.ChannelReq{
			ID:        &datas[key].ID,
			AppID:     &datas[key].AppID,
			EventType: &datas[key].EventType,
			Channel:   &datas[key].Channel,
		})
	}

	infos, err := CreateChannels(context.Background(), apps)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func getChannel(t *testing.T) {
	info, err := GetChannel(context.Background(), data.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, data, info)
	}
}

func getChannels(t *testing.T) {
	infos, total, err := GetChannels(context.Background(), &npool.Conds{
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

func getChannelOnly(t *testing.T) {
	info, err := GetChannelOnly(context.Background(), &npool.Conds{
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
	exist, err := ExistChannel(context.Background(), data.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existAppGoodConds(t *testing.T) {
	exist, err := ExistChannelConds(context.Background(), &npool.Conds{
		ID: &valuedef.StringVal{
			Op:    cruder.EQ,
			Value: data.ID,
		},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func deleteChannel(t *testing.T) {
	info, err := DeleteChannel(context.Background(), data.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, data, info)
	}

	_, err = GetChannel(context.Background(), info.ID)
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

	t.Run("createChannel", createChannel)
	t.Run("createChannels", createChannels)
	t.Run("getChannel", getChannel)
	t.Run("getChannels", getChannels)
	t.Run("getChannelOnly", getChannelOnly)
	t.Run("existAppGood", existAppGood)
	t.Run("existAppGoodConds", existAppGoodConds)
	t.Run("deleteChannel", deleteChannel)
}
