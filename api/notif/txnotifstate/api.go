package txnotifstate

import (
	txnotifstate "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif/txnotifstate"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	txnotifstate.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	txnotifstate.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
