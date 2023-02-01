package sendstate

import (
	sendstate "github.com/NpoolPlatform/message/npool/notif/mgr/v1/announcement/sendstate"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	sendstate.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	sendstate.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
