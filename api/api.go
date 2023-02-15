package api

import (
	"context"

	"github.com/NpoolPlatform/notif-manager/api/announcement"
	"github.com/NpoolPlatform/notif-manager/api/announcement/readstate"
	"github.com/NpoolPlatform/notif-manager/api/announcement/sendstate"
	"github.com/NpoolPlatform/notif-manager/api/announcement/user"
	"github.com/NpoolPlatform/notif-manager/api/contact"
	"github.com/NpoolPlatform/notif-manager/api/notif"
	"github.com/NpoolPlatform/notif-manager/api/notif/channel"
	"github.com/NpoolPlatform/notif-manager/api/notif/tx"
	"github.com/NpoolPlatform/notif-manager/api/template/email"
	"github.com/NpoolPlatform/notif-manager/api/template/frontend"
	"github.com/NpoolPlatform/notif-manager/api/template/sms"

	v1 "github.com/NpoolPlatform/message/npool/notif/mgr/v1"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	v1.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	v1.RegisterManagerServer(server, &Server{})
	notif.Register(server)
	announcement.Register(server)
	readstate.Register(server)
	sendstate.Register(server)
	user.Register(server)
	tx.Register(server)
	contact.Register(server)
	email.Register(server)
	frontend.Register(server)
	sms.Register(server)
	channel.Register(server)
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	if err := v1.RegisterManagerHandlerFromEndpoint(context.Background(), mux, endpoint, opts); err != nil {
		return err
	}
	return nil
}
