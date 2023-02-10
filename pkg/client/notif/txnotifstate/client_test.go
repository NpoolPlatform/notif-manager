package txnotifstate

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	valuedef "github.com/NpoolPlatform/message/npool"

	"bou.ke/monkey"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/NpoolPlatform/notif-manager/pkg/testinit"

	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif/txnotifstate"
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
	txState = npool.TxState_WaitSend
	txType  = npool.TxType_Withdraw
	data    = &npool.TxNotifState{
		ID:         uuid.NewString(),
		TxID:       uuid.NewString(),
		NotifState: txState,
		NotifType:  txType,
	}
)

var dataReq = &npool.TxNotifStateReq{
	ID:         &data.ID,
	TxID:       &data.TxID,
	NotifState: &txState,
	NotifType:  &txType,
}

func createTxNotifState(t *testing.T) {
	info, err := CreateTxNotifState(context.Background(), dataReq)
	if assert.Nil(t, err) {
		data.CreatedAt = info.CreatedAt
		data.UpdatedAt = info.UpdatedAt
		assert.Equal(t, data, info)
	}
}

func updateTxNotifState(t *testing.T) {
	info, err := UpdateTxNotifState(context.Background(), dataReq)
	if assert.Nil(t, err) {
		data.UpdatedAt = info.UpdatedAt
		assert.Equal(t, data, info)
	}
}

func createTxNotifStates(t *testing.T) {
	datas := []npool.TxNotifState{
		{
			ID:         uuid.NewString(),
			TxID:       uuid.NewString(),
			NotifState: txState,
			NotifType:  txType,
		},
		{
			ID:         uuid.NewString(),
			TxID:       uuid.NewString(),
			NotifState: txState,
			NotifType:  txType,
		},
	}

	apps := []*npool.TxNotifStateReq{}
	for key := range datas {
		apps = append(apps, &npool.TxNotifStateReq{
			ID:         &datas[key].ID,
			TxID:       &datas[key].TxID,
			NotifState: &datas[key].NotifState,
			NotifType:  &datas[key].NotifType,
		})
	}

	infos, err := CreateTxNotifStates(context.Background(), apps)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func getTxNotifState(t *testing.T) {
	info, err := GetTxNotifState(context.Background(), data.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, data, info)
	}
}

func getTxNotifStates(t *testing.T) {
	infos, total, err := GetTxNotifStates(context.Background(), &npool.Conds{
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

func getTxNotifStateOnly(t *testing.T) {
	info, err := GetTxNotifStateOnly(context.Background(), &npool.Conds{
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
	exist, err := ExistTxNotifState(context.Background(), data.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existAppGoodConds(t *testing.T) {
	exist, err := ExistTxNotifStateConds(context.Background(), &npool.Conds{
		ID: &valuedef.StringVal{
			Op:    cruder.EQ,
			Value: data.ID,
		},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func deleteTxNotifState(t *testing.T) {
	info, err := DeleteTxNotifState(context.Background(), data.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, data, info)
	}

	_, err = GetTxNotifState(context.Background(), info.ID)
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

	t.Run("createTxNotifState", createTxNotifState)
	t.Run("createTxNotifStates", createTxNotifStates)
	t.Run("getTxNotifState", getTxNotifState)
	t.Run("getTxNotifStates", getTxNotifStates)
	t.Run("getTxNotifStateOnly", getTxNotifStateOnly)
	t.Run("updateTxNotifState", updateTxNotifState)
	t.Run("existAppGood", existAppGood)
	t.Run("existAppGoodConds", existAppGoodConds)
	t.Run("deleteTxNotifState", deleteTxNotifState)
}
