package announcement

import (
	announcement "github.com/NpoolPlatform/message/npool/notif/mgr/v1/announcement"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	announcement.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	announcement.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
