//nolint:nolintlint,dupl
package sendstate

import (
	"context"

	converter "github.com/NpoolPlatform/notif-manager/pkg/converter/announcement/sendstate"
	crud "github.com/NpoolPlatform/notif-manager/pkg/crud/announcement/sendstate"
	commontracer "github.com/NpoolPlatform/notif-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/notif-manager/pkg/tracer/announcement/sendstate"

	constant "github.com/NpoolPlatform/notif-manager/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/announcement/sendstate"

	"github.com/google/uuid"
)

func (s *Server) CreateSendState(ctx context.Context, in *npool.CreateSendStateRequest) (*npool.CreateSendStateResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateSendState")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	if err := validate(in.GetInfo()); err != nil {
		return &npool.CreateSendStateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "announcement/sendstate", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateSendState", "error", err)
		return &npool.CreateSendStateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateSendStateResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateSendStates(ctx context.Context, in *npool.CreateSendStatesRequest) (*npool.CreateSendStatesResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateSendStates")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		return &npool.CreateSendStatesResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "announcement/sendstate", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorw("CreateSendStates", "error", err)
		return &npool.CreateSendStatesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateSendStatesResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateSendState(ctx context.Context, in *npool.UpdateSendStateRequest) (*npool.UpdateSendStateResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateSendState")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("UpdateSendState", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateSendStateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "announcement/sendstate", "crud", "Update")

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateSendState", "error", err)
		return &npool.UpdateSendStateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateSendStateResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetSendState(ctx context.Context, in *npool.GetSendStateRequest) (*npool.GetSendStateResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetSendState")
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
		logger.Sugar().Errorw("GetSendState", "ID", in.GetID(), "error", err)
		return &npool.GetSendStateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "announcement/sendstate", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("GetSendState", "error", err)
		return &npool.GetSendStateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetSendStateResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetSendStateOnly(ctx context.Context, in *npool.GetSendStateOnlyRequest) (*npool.GetSendStateOnlyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetSendStateOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := validateConds(in.GetConds()); err != nil {
		return &npool.GetSendStateOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "announcement/sendstate", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetSendStateOnly", "error", err)
		return &npool.GetSendStateOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetSendStateOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetSendStates(ctx context.Context, in *npool.GetSendStatesRequest) (*npool.GetSendStatesResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetSendStates")
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
		return &npool.GetSendStatesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "announcement/sendstate", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorw("GetSendStates", "error", err)
		return &npool.GetSendStatesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetSendStatesResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistSendState(ctx context.Context, in *npool.ExistSendStateRequest) (*npool.ExistSendStateResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistSendState")
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
		return &npool.ExistSendStateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "announcement/sendstate", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("ExistSendState", "error", err)
		return &npool.ExistSendStateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistSendStateResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistSendStateConds(ctx context.Context,
	in *npool.ExistSendStateCondsRequest) (*npool.ExistSendStateCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistSendStateConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "announcement/sendstate", "crud", "ExistConds")

	if err := validateConds(in.GetConds()); err != nil {
		return &npool.ExistSendStateCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("ExistSendStateConds", "error", err)
		return &npool.ExistSendStateCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistSendStateCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountSendStates(ctx context.Context, in *npool.CountSendStatesRequest) (*npool.CountSendStatesResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountSendStates")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "announcement/sendstate", "crud", "Count")

	if err := validateConds(in.GetConds()); err != nil {
		return &npool.CountSendStatesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("CountSendStates", "error", err)
		return &npool.CountSendStatesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountSendStatesResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteSendState(ctx context.Context, in *npool.DeleteSendStateRequest) (*npool.DeleteSendStateResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteSendState")
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
		return &npool.DeleteSendStateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "announcement/sendstate", "crud", "Delete")

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("DeleteSendState", "error", err)
		return &npool.DeleteSendStateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteSendStateResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
