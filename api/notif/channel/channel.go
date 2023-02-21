//nolint:nolintlint,dupl
package channel

import (
	"context"

	converter "github.com/NpoolPlatform/notif-manager/pkg/converter/notif/channel"
	crud "github.com/NpoolPlatform/notif-manager/pkg/crud/notif/channel"
	commontracer "github.com/NpoolPlatform/notif-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/notif-manager/pkg/tracer/notif/channel"

	constant "github.com/NpoolPlatform/notif-manager/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif/channel"

	"github.com/google/uuid"
)

func (s *Server) CreateChannel(
	ctx context.Context,
	in *npool.CreateChannelRequest,
) (
	*npool.CreateChannelResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateChannel")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	if err := validate(in.GetInfo()); err != nil {
		return &npool.CreateChannelResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "channel", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateChannel", "error", err)
		return &npool.CreateChannelResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateChannelResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateChannels(
	ctx context.Context,
	in *npool.CreateChannelsRequest,
) (
	*npool.CreateChannelsResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateChannels")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		return &npool.CreateChannelsResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "channel", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorw("CreateChannels", "error", err)
		return &npool.CreateChannelsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateChannelsResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) GetChannel(ctx context.Context, in *npool.GetChannelRequest) (*npool.GetChannelResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetChannel")
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
		logger.Sugar().Errorw("GetChannel", "ID", in.GetID(), "error", err)
		return &npool.GetChannelResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "channel", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("GetChannel", "error", err)
		return &npool.GetChannelResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetChannelResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetChannelOnly(
	ctx context.Context,
	in *npool.GetChannelOnlyRequest,
) (
	*npool.GetChannelOnlyResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetChannelOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := validateConds(in.GetConds()); err != nil {
		return &npool.GetChannelOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "channel", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetChannelOnly", "error", err)
		return &npool.GetChannelOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetChannelOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetChannels(
	ctx context.Context,
	in *npool.GetChannelsRequest,
) (
	*npool.GetChannelsResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetChannels")
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
		return &npool.GetChannelsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "channel", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorw("GetChannels", "error", err)
		return &npool.GetChannelsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetChannelsResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistChannel(
	ctx context.Context,
	in *npool.ExistChannelRequest,
) (
	*npool.ExistChannelResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistChannel")
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
		return &npool.ExistChannelResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "channel", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("ExistChannel", "error", err)
		return &npool.ExistChannelResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistChannelResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistChannelConds(
	ctx context.Context,
	in *npool.ExistChannelCondsRequest,
) (*npool.ExistChannelCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistChannelConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "channel", "crud", "ExistConds")

	if err := validateConds(in.GetConds()); err != nil {
		return &npool.ExistChannelCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("ExistChannelConds", "error", err)
		return &npool.ExistChannelCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistChannelCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountChannels(ctx context.Context, in *npool.CountChannelsRequest) (*npool.CountChannelsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountChannels")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "channel", "crud", "Count")

	if err := validateConds(in.GetConds()); err != nil {
		return &npool.CountChannelsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("CountChannels", "error", err)
		return &npool.CountChannelsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountChannelsResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteChannel(ctx context.Context, in *npool.DeleteChannelRequest) (*npool.DeleteChannelResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteChannel")
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
		return &npool.DeleteChannelResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "channel", "crud", "Delete")

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("DeleteChannel", "error", err)
		return &npool.DeleteChannelResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteChannelResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
