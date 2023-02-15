package tx

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

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif/tx"
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
	txState = npool.TxState_WaitNotified
	txType  = basetypes.TxType_TxWithdraw
	data    = &npool.Tx{
		ID:         uuid.NewString(),
		TxID:       uuid.NewString(),
		NotifState: txState,
		TxType:     txType,
	}
)

var dataReq = &npool.TxReq{
	ID:         &data.ID,
	TxID:       &data.TxID,
	NotifState: &txState,
	TxType:     &txType,
}

func createTx(t *testing.T) {
	info, err := CreateTx(context.Background(), dataReq)
	if assert.Nil(t, err) {
		data.CreatedAt = info.CreatedAt
		data.UpdatedAt = info.UpdatedAt
		assert.Equal(t, data, info)
	}
}

func updateTx(t *testing.T) {
	info, err := UpdateTx(context.Background(), dataReq)
	if assert.Nil(t, err) {
		data.UpdatedAt = info.UpdatedAt
		assert.Equal(t, data, info)
	}
}

func createTxs(t *testing.T) {
	datas := []npool.Tx{
		{
			ID:         uuid.NewString(),
			TxID:       uuid.NewString(),
			NotifState: txState,
			TxType:     txType,
		},
		{
			ID:         uuid.NewString(),
			TxID:       uuid.NewString(),
			NotifState: txState,
			TxType:     txType,
		},
	}

	apps := []*npool.TxReq{}
	for key := range datas {
		apps = append(apps, &npool.TxReq{
			ID:         &datas[key].ID,
			TxID:       &datas[key].TxID,
			NotifState: &datas[key].NotifState,
			TxType:     &datas[key].TxType,
		})
	}

	infos, err := CreateTxs(context.Background(), apps)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func getTx(t *testing.T) {
	info, err := GetTx(context.Background(), data.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, data, info)
	}
}

func getTxs(t *testing.T) {
	infos, total, err := GetTxs(context.Background(), &npool.Conds{
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

func getTxOnly(t *testing.T) {
	info, err := GetTxOnly(context.Background(), &npool.Conds{
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
	exist, err := ExistTx(context.Background(), data.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existAppGoodConds(t *testing.T) {
	exist, err := ExistTxConds(context.Background(), &npool.Conds{
		ID: &valuedef.StringVal{
			Op:    cruder.EQ,
			Value: data.ID,
		},
	})
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func deleteTx(t *testing.T) {
	info, err := DeleteTx(context.Background(), data.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, data, info)
	}

	_, err = GetTx(context.Background(), info.ID)
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

	t.Run("createTx", createTx)
	t.Run("createTxs", createTxs)
	t.Run("getTx", getTx)
	t.Run("getTxs", getTxs)
	t.Run("getTxOnly", getTxOnly)
	t.Run("updateTx", updateTx)
	t.Run("existAppGood", existAppGood)
	t.Run("existAppGoodConds", existAppGoodConds)
	t.Run("deleteTx", deleteTx)
}
