package frontend

import (
	"github.com/NpoolPlatform/message/npool/notif/mgr/v1/template/frontend"
	"google.golang.org/grpc"
)

type Server struct {
	frontend.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	frontend.RegisterManagerServer(server, &Server{})
}
