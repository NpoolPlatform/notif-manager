package contact

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	"google.golang.org/grpc/credentials/insecure"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"

	"bou.ke/monkey"
	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"
	"google.golang.org/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	val "github.com/NpoolPlatform/message/npool"

	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/contact"
	testinit "github.com/NpoolPlatform/notif-manager/pkg/testinit"
	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

var appDate = npool.Contact{
	ID:          uuid.NewString(),
	AppID:       uuid.NewString(),
	UsedFor:     basetypes.UsedFor_Signin,
	Account:     uuid.NewString(),
	AccountType: basetypes.SignMethod_Email,
	Sender:      uuid.NewString(),
}

var (
	appInfo = npool.ContactReq{
		ID:          &appDate.ID,
		AppID:       &appDate.AppID,
		UsedFor:     &appDate.UsedFor,
		Account:     &appDate.Account,
		AccountType: &appDate.AccountType,
		Sender:      &appDate.Sender,
	}
)

var info *npool.Contact

func createContact(t *testing.T) {
	var err error
	info, err = CreateContact(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appDate)
	}
}

func createContacts(t *testing.T) {
	appDates := []npool.Contact{
		{
			ID:          uuid.NewString(),
			AppID:       uuid.NewString(),
			UsedFor:     basetypes.UsedFor_Signin,
			Account:     uuid.NewString(),
			AccountType: basetypes.SignMethod_Email,
			Sender:      uuid.NewString(),
		},
		{
			ID:          uuid.NewString(),
			AppID:       uuid.NewString(),
			UsedFor:     basetypes.UsedFor_Signin,
			Account:     uuid.NewString(),
			AccountType: basetypes.SignMethod_Email,
			Sender:      uuid.NewString(),
		},
	}

	apps := []*npool.ContactReq{}
	for key := range appDates {
		apps = append(apps, &npool.ContactReq{
			ID:          &appDates[key].ID,
			AppID:       &appDates[key].AppID,
			UsedFor:     &appDates[key].UsedFor,
			Account:     &appDates[key].Account,
			AccountType: &appDates[key].AccountType,
			Sender:      &appDates[key].Sender,
		})
	}

	infos, err := CreateContacts(context.Background(), apps)
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
	}
}

func updateContact(t *testing.T) {
	var err error
	info, err = UpdateContact(context.Background(), &appInfo)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appDate)
	}
}

func getContact(t *testing.T) {
	var err error
	info, err = GetContact(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appDate)
	}
}

func getContacts(t *testing.T) {
	infos, total, err := GetContacts(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		}, 0, 1)
	if assert.Nil(t, err) {
		assert.Equal(t, total, uint32(1))
		assert.Equal(t, infos[0], &appDate)
	}
}

func getContactOnly(t *testing.T) {
	var err error
	info, err = GetContactOnly(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		})
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appDate)
	}
}

func existContact(t *testing.T) {
	exist, err := ExistContact(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func existContactConds(t *testing.T) {
	exist, err := ExistContactConds(context.Background(),
		&npool.Conds{
			ID: &val.StringVal{
				Value: info.ID,
				Op:    cruder.EQ,
			},
		},
	)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}
}

func deleteContact(t *testing.T) {
	info, err := DeleteContact(context.Background(), info.ID)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &appDate)
	}
}

func TestMainOrder(t *testing.T) {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}

	gport := config.GetIntValueWithNameSpace("", config.KeyGRPCPort)

	monkey.Patch(grpc2.GetGRPCConn, func(service string, tags ...string) (*grpc.ClientConn, error) {
		return grpc.Dial(fmt.Sprintf("localhost:%v", gport), grpc.WithTransportCredentials(insecure.NewCredentials()))
	})

	t.Run("createContact", createContact)
	t.Run("createContacts", createContacts)
	t.Run("getContact", getContact)
	t.Run("getContacts", getContacts)
	t.Run("getContactOnly", getContactOnly)
	t.Run("updateContact", updateContact)
	t.Run("existContact", existContact)
	t.Run("existContactConds", existContactConds)
	t.Run("delete", deleteContact)
}
