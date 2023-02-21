package channel

import (
	channel "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif/channel"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	channel.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	channel.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
