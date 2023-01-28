package notif

import (
	notif "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Server struct {
	notif.UnimplementedManagerServer
}

func Register(server grpc.ServiceRegistrar) {
	notif.RegisterManagerServer(server, &Server{})
}

func RegisterGateway(mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
	return nil
}
