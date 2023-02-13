//nolint:nolintlint,dupl
package notifchannel

import (
	"context"

	converter "github.com/NpoolPlatform/notif-manager/pkg/converter/notif/notifchannel"
	crud "github.com/NpoolPlatform/notif-manager/pkg/crud/notif/notifchannel"
	commontracer "github.com/NpoolPlatform/notif-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/notif-manager/pkg/tracer/notif/notifchannel"

	constant "github.com/NpoolPlatform/notif-manager/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif/notifchannel"

	"github.com/google/uuid"
)

func (s *Server) CreateNotifChannel(
	ctx context.Context,
	in *npool.CreateNotifChannelRequest,
) (
	*npool.CreateNotifChannelResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateNotifChannel")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	if err := validate(in.GetInfo()); err != nil {
		return &npool.CreateNotifChannelResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "notifchannel", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateNotifChannel", "error", err)
		return &npool.CreateNotifChannelResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateNotifChannelResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateNotifChannels(
	ctx context.Context,
	in *npool.CreateNotifChannelsRequest,
) (
	*npool.CreateNotifChannelsResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateNotifChannels")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		return &npool.CreateNotifChannelsResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "notifchannel", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorw("CreateNotifChannels", "error", err)
		return &npool.CreateNotifChannelsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateNotifChannelsResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) GetNotifChannel(ctx context.Context, in *npool.GetNotifChannelRequest) (*npool.GetNotifChannelResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetNotifChannel")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, in.GetID())

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorw("GetNotifChannel", "ID", in.GetID(), "error", err)
		return &npool.GetNotifChannelResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "notifchannel", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("GetNotifChannel", "error", err)
		return &npool.GetNotifChannelResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetNotifChannelResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetNotifChannelOnly(
	ctx context.Context,
	in *npool.GetNotifChannelOnlyRequest,
) (
	*npool.GetNotifChannelOnlyResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetNotifChannelOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := validateConds(in.GetConds()); err != nil {
		return &npool.GetNotifChannelOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "notifchannel", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetNotifChannelOnly", "error", err)
		return &npool.GetNotifChannelOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetNotifChannelOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetNotifChannels(
	ctx context.Context,
	in *npool.GetNotifChannelsRequest,
) (
	*npool.GetNotifChannelsResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetNotifChannels")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))

	if err := validateConds(in.GetConds()); err != nil {
		return &npool.GetNotifChannelsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "notifchannel", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorw("GetNotifChannels", "error", err)
		return &npool.GetNotifChannelsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetNotifChannelsResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistNotifChannel(
	ctx context.Context,
	in *npool.ExistNotifChannelRequest,
) (
	*npool.ExistNotifChannelResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistNotifChannel")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, in.GetID())

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.ExistNotifChannelResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "notifchannel", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("ExistNotifChannel", "error", err)
		return &npool.ExistNotifChannelResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistNotifChannelResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistNotifChannelConds(
	ctx context.Context,
	in *npool.ExistNotifChannelCondsRequest,
) (*npool.ExistNotifChannelCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistNotifChannelConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "notifchannel", "crud", "ExistConds")

	if err := validateConds(in.GetConds()); err != nil {
		return &npool.ExistNotifChannelCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("ExistNotifChannelConds", "error", err)
		return &npool.ExistNotifChannelCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistNotifChannelCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountNotifChannels(ctx context.Context, in *npool.CountNotifChannelsRequest) (*npool.CountNotifChannelsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountNotifChannels")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "notifchannel", "crud", "Count")

	if err := validateConds(in.GetConds()); err != nil {
		return &npool.CountNotifChannelsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("CountNotifChannels", "error", err)
		return &npool.CountNotifChannelsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountNotifChannelsResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteNotifChannel(ctx context.Context, in *npool.DeleteNotifChannelRequest) (*npool.DeleteNotifChannelResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteNotifChannel")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, in.GetID())

	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.DeleteNotifChannelResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "notifchannel", "crud", "Delete")

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("DeleteNotifChannel", "error", err)
		return &npool.DeleteNotifChannelResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteNotifChannelResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
