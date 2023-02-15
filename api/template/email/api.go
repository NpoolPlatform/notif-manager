package email

import (
	"github.com/NpoolPlatform/message/npool/notif/mgr/v1/template/email"
	"google.golang.org/grpc"
)

type Server struct {
	email.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	email.RegisterManagerServer(server, &Server{})
}
