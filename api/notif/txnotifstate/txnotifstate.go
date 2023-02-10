//nolint:nolintlint,dupl
package txnotifstate

import (
	"context"

	converter "github.com/NpoolPlatform/notif-manager/pkg/converter/notif/txnotifstate"
	crud "github.com/NpoolPlatform/notif-manager/pkg/crud/notif/txnotifstate"
	commontracer "github.com/NpoolPlatform/notif-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/notif-manager/pkg/tracer/notif/txnotifstate"

	constant "github.com/NpoolPlatform/notif-manager/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif/txnotifstate"

	"github.com/google/uuid"
)

func (s *Server) CreateTxNotifState(
	ctx context.Context,
	in *npool.CreateTxNotifStateRequest,
) (
	*npool.CreateTxNotifStateResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateTxNotifState")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	if err := validate(in.GetInfo()); err != nil {
		return &npool.CreateTxNotifStateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "txnotifstate", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateTxNotifState", "error", err)
		return &npool.CreateTxNotifStateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateTxNotifStateResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateTxNotifStates(
	ctx context.Context,
	in *npool.CreateTxNotifStatesRequest,
) (
	*npool.CreateTxNotifStatesResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateTxNotifStates")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		return &npool.CreateTxNotifStatesResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "txnotifstate", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorw("CreateTxNotifStates", "error", err)
		return &npool.CreateTxNotifStatesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateTxNotifStatesResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateTxNotifState(
	ctx context.Context,
	in *npool.UpdateTxNotifStateRequest,
) (
	*npool.UpdateTxNotifStateResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateTxNotifState")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("UpdateTxNotifState", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateTxNotifStateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if in.Info == nil {
		logger.Sugar().Errorw("UpdateTxNotifState", "error", err)
		return &npool.UpdateTxNotifStateResponse{}, status.Error(codes.InvalidArgument, "Info is empty")
	}

	if in.GetInfo().NotifState != nil {
		switch in.GetInfo().GetNotifState() {
		case npool.TxState_WaitTxSuccess:
		case npool.TxState_WaitSend:
		case npool.TxState_AlreadySend:
		default:
			logger.Sugar().Errorw("UpdateTxNotifState", "ID", in.GetInfo().GetID(), "error", err)
			return &npool.UpdateTxNotifStateResponse{}, status.Error(codes.InvalidArgument, "NotifState is invalid")
		}
	}

	if in.GetInfo().NotifType != nil {
		switch in.GetInfo().GetNotifType() {
		case npool.TxType_Withdraw:
		default:
			logger.Sugar().Errorw("UpdateTxNotifType", "ID", in.GetInfo().GetID(), "error", err)
			return &npool.UpdateTxNotifStateResponse{}, status.Error(codes.InvalidArgument, "NotifType is invalid")
		}
	}

	span = commontracer.TraceInvoker(span, "txnotifstate", "crud", "Update")

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateTxNotifState", "error", err)
		return &npool.UpdateTxNotifStateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateTxNotifStateResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetTxNotifState(ctx context.Context, in *npool.GetTxNotifStateRequest) (*npool.GetTxNotifStateResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetTxNotifState")
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
		logger.Sugar().Errorw("GetTxNotifState", "ID", in.GetID(), "error", err)
		return &npool.GetTxNotifStateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "txnotifstate", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("GetTxNotifState", "error", err)
		return &npool.GetTxNotifStateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetTxNotifStateResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetTxNotifStateOnly(
	ctx context.Context,
	in *npool.GetTxNotifStateOnlyRequest,
) (
	*npool.GetTxNotifStateOnlyResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetTxNotifStateOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := validateConds(in.GetConds()); err != nil {
		return &npool.GetTxNotifStateOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "txnotifstate", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetTxNotifStateOnly", "error", err)
		return &npool.GetTxNotifStateOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetTxNotifStateOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetTxNotifStates(
	ctx context.Context,
	in *npool.GetTxNotifStatesRequest,
) (
	*npool.GetTxNotifStatesResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetTxNotifStates")
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
		return &npool.GetTxNotifStatesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "txnotifstate", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorw("GetTxNotifStates", "error", err)
		return &npool.GetTxNotifStatesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetTxNotifStatesResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistTxNotifState(
	ctx context.Context,
	in *npool.ExistTxNotifStateRequest,
) (
	*npool.ExistTxNotifStateResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistTxNotifState")
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
		return &npool.ExistTxNotifStateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "txnotifstate", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("ExistTxNotifState", "error", err)
		return &npool.ExistTxNotifStateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistTxNotifStateResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistTxNotifStateConds(
	ctx context.Context,
	in *npool.ExistTxNotifStateCondsRequest,
) (*npool.ExistTxNotifStateCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistTxNotifStateConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "txnotifstate", "crud", "ExistConds")

	if err := validateConds(in.GetConds()); err != nil {
		return &npool.ExistTxNotifStateCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("ExistTxNotifStateConds", "error", err)
		return &npool.ExistTxNotifStateCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistTxNotifStateCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountTxNotifStates(ctx context.Context, in *npool.CountTxNotifStatesRequest) (*npool.CountTxNotifStatesResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountTxNotifStates")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "txnotifstate", "crud", "Count")

	if err := validateConds(in.GetConds()); err != nil {
		return &npool.CountTxNotifStatesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("CountTxNotifStates", "error", err)
		return &npool.CountTxNotifStatesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountTxNotifStatesResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteTxNotifState(ctx context.Context, in *npool.DeleteTxNotifStateRequest) (*npool.DeleteTxNotifStateResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteTxNotifState")
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
		return &npool.DeleteTxNotifStateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "txnotifstate", "crud", "Delete")

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("DeleteTxNotifState", "error", err)
		return &npool.DeleteTxNotifStateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteTxNotifStateResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
