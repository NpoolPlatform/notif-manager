package readstate

import (
	readstate "github.com/NpoolPlatform/message/npool/notif/mgr/v1/announcement/readstate"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	readstate.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	readstate.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
