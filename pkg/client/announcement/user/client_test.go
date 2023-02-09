package user

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

	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/announcement/user"
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

var data = &npool.User{
	ID:             uuid.NewString(),
	AppID:          uuid.NewString(),
	UserID:         uuid.NewString(),
	AnnouncementID: uuid.NewString(),
}

var dataReq = &npool.UserReq{
	ID:             &data.ID,
	AppID:          &data.AppID,
	UserID:         &data.UserID,
	AnnouncementID: &data.AnnouncementID,
}

func createUser(t *testing.T) {
	info, err := CreateUser(context.Background(), dataReq)
	if assert.Nil(t, err) {
		data.CreatedAt = info.CreatedAt
		data.UpdatedAt = info.UpdatedAt
		assert.Equal(t, data, info)
	}
}

func updateUser(t *testing.T) {
	info, err := UpdateUser(context.Background(), dataReq)
	if assert.Nil(t, err) {
		data.UpdatedAt = info.UpdatedAt
		assert.Equal(t, data, info)
	}
}

func createUsers(t *testing.T) {
	datas := []npool.User{
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

	apps := []*npool.UserReq{}
	for key := range datas {
		apps = append(apps, &npool.UserReq{
			ID:             &datas[key].ID,
			AppID:          &datas[key].AppID,
			UserID:         &datas[key].UserID,
			AnnouncementID: &datas[key].AnnouncementID,
		})
	}

	infos, err := CreateUsers(context.Background(), apps)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func getUser(t *testing.T) {
	info, err := GetUser(context.Background(), data.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, data, info)
	}
}

func getUsers(t *testing.T) {
	infos, total, err := GetUsers(context.Background(), &npool.Conds{
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

func getUserOnly(t *testing.T) {
	info, err := GetUserOnly(context.Background(), &npool.Conds{
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
	exist, err := ExistUser(context.Background(), data.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existAppGoodConds(t *testing.T) {
	exist, err := ExistUserConds(context.Background(), &npool.Conds{
		ID: &valuedef.StringVal{
			Op:    cruder.EQ,
			Value: data.ID,
		},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func deleteUser(t *testing.T) {
	info, err := DeleteUser(context.Background(), data.ID)
	if assert.Nil(t, err) {
		data.UpdatedAt = info.UpdatedAt
		data.CreatedAt = info.CreatedAt
		assert.Equal(t, data, info)
	}

	_, err = GetUser(context.Background(), info.ID)
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

	t.Run("createUser", createUser)
	t.Run("createUsers", createUsers)
	t.Run("getUser", getUser)
	t.Run("getUsers", getUsers)
	t.Run("getUserOnly", getUserOnly)
	t.Run("updateUser", updateUser)
	t.Run("existAppGood", existAppGood)
	t.Run("existAppGoodConds", existAppGoodConds)
	t.Run("deleteUser", deleteUser)
}
