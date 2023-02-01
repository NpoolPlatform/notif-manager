package sendstate

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/message/npool/notif/mgr/v1/channel"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	valuedef "github.com/NpoolPlatform/message/npool"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	"bou.ke/monkey"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/NpoolPlatform/notif-manager/pkg/testinit"

	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/announcement/sendstate"
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

var data = &npool.SendState{
	ID:             uuid.NewString(),
	AppID:          uuid.NewString(),
	UserID:         uuid.NewString(),
	AnnouncementID: uuid.NewString(),
	Channel:        channel.NotifChannel_ChannelEmail,
}

var dataReq = &npool.SendStateReq{
	ID:             &data.ID,
	AppID:          &data.AppID,
	UserID:         &data.UserID,
	AnnouncementID: &data.AnnouncementID,
	Channel:        &data.Channel,
}

func createSendState(t *testing.T) {
	info, err := CreateSendState(context.Background(), dataReq)
	if assert.Nil(t, err) {
		data.CreatedAt = info.CreatedAt
		data.UpdatedAt = info.UpdatedAt
		assert.Equal(t, data, info)
	}
}

func updateSendState(t *testing.T) {
	info, err := UpdateSendState(context.Background(), dataReq)
	if assert.Nil(t, err) {
		data.UpdatedAt = info.UpdatedAt
		assert.Equal(t, data, info)
	}
}

func createSendStates(t *testing.T) {
	datas := []npool.SendState{
		{
			ID:             uuid.NewString(),
			AppID:          uuid.NewString(),
			UserID:         uuid.NewString(),
			AnnouncementID: uuid.NewString(),
			Channel:        channel.NotifChannel_ChannelEmail,
		},
		{
			ID:             uuid.NewString(),
			AppID:          uuid.NewString(),
			UserID:         uuid.NewString(),
			AnnouncementID: uuid.NewString(),
			Channel:        channel.NotifChannel_ChannelEmail,
		},
	}

	apps := []*npool.SendStateReq{}
	for key := range datas {
		apps = append(apps, &npool.SendStateReq{
			ID:             &datas[key].ID,
			AppID:          &datas[key].AppID,
			UserID:         &datas[key].UserID,
			AnnouncementID: &datas[key].AnnouncementID,
			Channel:        &datas[key].Channel,
		})
	}

	infos, err := CreateSendStates(context.Background(), apps)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func getSendState(t *testing.T) {
	info, err := GetSendState(context.Background(), data.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, data, info)
	}
}

func getSendStates(t *testing.T) {
	infos, total, err := GetSendStates(context.Background(), &npool.Conds{
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

func getSendStateOnly(t *testing.T) {
	info, err := GetSendStateOnly(context.Background(), &npool.Conds{
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
	exist, err := ExistSendState(context.Background(), data.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existAppGoodConds(t *testing.T) {
	exist, err := ExistSendStateConds(context.Background(), &npool.Conds{
		ID: &valuedef.StringVal{
			Op:    cruder.EQ,
			Value: data.ID,
		},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func deleteSendState(t *testing.T) {
	info, err := DeleteSendState(context.Background(), data.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, data, info)
	}

	_, err = GetSendState(context.Background(), info.ID)
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

	t.Run("createSendState", createSendState)
	t.Run("createSendStates", createSendStates)
	t.Run("getSendState", getSendState)
	t.Run("getSendStates", getSendStates)
	t.Run("getSendStateOnly", getSendStateOnly)
	t.Run("updateSendState", updateSendState)
	t.Run("existAppGood", existAppGood)
	t.Run("existAppGoodConds", existAppGoodConds)
	t.Run("deleteSendState", deleteSendState)
}
