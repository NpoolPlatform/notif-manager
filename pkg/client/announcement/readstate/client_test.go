package readstate

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	valuedef "github.com/NpoolPlatform/message/npool"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	"bou.ke/monkey"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/NpoolPlatform/notif-manager/pkg/testinit"

	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/announcement/readstate"
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

var data = &npool.ReadState{
	ID:             uuid.NewString(),
	AppID:          uuid.NewString(),
	UserID:         uuid.NewString(),
	AnnouncementID: uuid.NewString(),
}

var dataReq = &npool.ReadStateReq{
	ID:             &data.ID,
	AppID:          &data.AppID,
	UserID:         &data.UserID,
	AnnouncementID: &data.AnnouncementID,
}

func createReadState(t *testing.T) {
	info, err := CreateReadState(context.Background(), dataReq)
	if assert.Nil(t, err) {
		data.CreatedAt = info.CreatedAt
		data.UpdatedAt = info.UpdatedAt
		assert.Equal(t, data, info)
	}
}

func updateReadState(t *testing.T) {
	info, err := UpdateReadState(context.Background(), dataReq)
	if assert.Nil(t, err) {
		data.UpdatedAt = info.UpdatedAt
		assert.Equal(t, data, info)
	}
}

func createReadStates(t *testing.T) {
	datas := []npool.ReadState{
		{
			ID:             uuid.NewString(),
			AppID:          uuid.NewString(),
			UserID:         uuid.NewString(),
			AnnouncementID: uuid.NewString(),
		},
		{
			ID:             uuid.NewString(),
			AppID:          uuid.NewString(),
			UserID:         uuid.NewString(),
			AnnouncementID: uuid.NewString(),
		},
	}

	apps := []*npool.ReadStateReq{}
	for key := range datas {
		apps = append(apps, &npool.ReadStateReq{
			ID:             &datas[key].ID,
			AppID:          &datas[key].AppID,
			UserID:         &datas[key].UserID,
			AnnouncementID: &datas[key].AnnouncementID,
		})
	}

	infos, err := CreateReadStates(context.Background(), apps)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func getReadState(t *testing.T) {
	info, err := GetReadState(context.Background(), data.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, data, info)
	}
}

func getReadStates(t *testing.T) {
	infos, total, err := GetReadStates(context.Background(), &npool.Conds{
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

func getReadStateOnly(t *testing.T) {
	info, err := GetReadStateOnly(context.Background(), &npool.Conds{
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
	exist, err := ExistReadState(context.Background(), data.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existAppGoodConds(t *testing.T) {
	exist, err := ExistReadStateConds(context.Background(), &npool.Conds{
		ID: &valuedef.StringVal{
			Op:    cruder.EQ,
			Value: data.ID,
		},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func deleteReadState(t *testing.T) {
	info, err := DeleteReadState(context.Background(), data.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, data, info)
	}

	_, err = GetReadState(context.Background(), info.ID)
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

	t.Run("createReadState", createReadState)
	t.Run("createReadStates", createReadStates)
	t.Run("getReadState", getReadState)
	t.Run("getReadStates", getReadStates)
	t.Run("getReadStateOnly", getReadStateOnly)
	t.Run("updateReadState", updateReadState)
	t.Run("existAppGood", existAppGood)
	t.Run("existAppGoodConds", existAppGoodConds)
	t.Run("deleteReadState", deleteReadState)
}
