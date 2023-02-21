//nolint:nolintlint,dupl
package tx

import (
	"context"

	converter "github.com/NpoolPlatform/notif-manager/pkg/converter/notif/tx"
	crud "github.com/NpoolPlatform/notif-manager/pkg/crud/notif/tx"
	commontracer "github.com/NpoolPlatform/notif-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/notif-manager/pkg/tracer/notif/tx"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	constant "github.com/NpoolPlatform/notif-manager/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif/tx"

	"github.com/google/uuid"
)

func (s *Server) CreateTx(
	ctx context.Context,
	in *npool.CreateTxRequest,
) (
	*npool.CreateTxResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateTx")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	if err := validate(in.GetInfo()); err != nil {
		return &npool.CreateTxResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "tx", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateTx", "error", err)
		return &npool.CreateTxResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateTxResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateTxs(
	ctx context.Context,
	in *npool.CreateTxsRequest,
) (
	*npool.CreateTxsResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateTxs")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		return &npool.CreateTxsResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "tx", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorw("CreateTxs", "error", err)
		return &npool.CreateTxsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateTxsResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateTx(
	ctx context.Context,
	in *npool.UpdateTxRequest,
) (
	*npool.UpdateTxResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateTx")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("UpdateTx", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateTxResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if in.Info == nil {
		logger.Sugar().Errorw("UpdateTx", "error", err)
		return &npool.UpdateTxResponse{}, status.Error(codes.InvalidArgument, "Info is empty")
	}

	if in.GetInfo().NotifState != nil {
		switch in.GetInfo().GetNotifState() {
		case npool.TxState_WaitSuccess:
		case npool.TxState_WaitNotified:
		case npool.TxState_Notified:
		default:
			logger.Sugar().Errorw("UpdateTx", "ID", in.GetInfo().GetID(), "error", err)
			return &npool.UpdateTxResponse{}, status.Error(codes.InvalidArgument, "NotifState is invalid")
		}
	}

	if in.GetInfo().TxType != nil {
		switch in.GetInfo().GetTxType() {
		case basetypes.TxType_TxWithdraw:
		default:
			logger.Sugar().Errorw("UpdateTxTxType", "ID", in.GetInfo().GetID(), "error", err)
			return &npool.UpdateTxResponse{}, status.Error(codes.InvalidArgument, "TxType is invalid")
		}
	}

	span = commontracer.TraceInvoker(span, "tx", "crud", "Update")

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateTx", "error", err)
		return &npool.UpdateTxResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateTxResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetTx(ctx context.Context, in *npool.GetTxRequest) (*npool.GetTxResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetTx")
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
		logger.Sugar().Errorw("GetTx", "ID", in.GetID(), "error", err)
		return &npool.GetTxResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "tx", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("GetTx", "error", err)
		return &npool.GetTxResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetTxResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetTxOnly(
	ctx context.Context,
	in *npool.GetTxOnlyRequest,
) (
	*npool.GetTxOnlyResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetTxOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := validateConds(in.GetConds()); err != nil {
		return &npool.GetTxOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "tx", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetTxOnly", "error", err)
		return &npool.GetTxOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetTxOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetTxs(
	ctx context.Context,
	in *npool.GetTxsRequest,
) (
	*npool.GetTxsResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetTxs")
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
		return &npool.GetTxsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "tx", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorw("GetTxs", "error", err)
		return &npool.GetTxsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetTxsResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistTx(
	ctx context.Context,
	in *npool.ExistTxRequest,
) (
	*npool.ExistTxResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistTx")
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
		return &npool.ExistTxResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "tx", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("ExistTx", "error", err)
		return &npool.ExistTxResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistTxResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistTxConds(
	ctx context.Context,
	in *npool.ExistTxCondsRequest,
) (*npool.ExistTxCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistTxConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "tx", "crud", "ExistConds")

	if err := validateConds(in.GetConds()); err != nil {
		return &npool.ExistTxCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("ExistTxConds", "error", err)
		return &npool.ExistTxCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistTxCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountTxs(ctx context.Context, in *npool.CountTxsRequest) (*npool.CountTxsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountTxs")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "tx", "crud", "Count")

	if err := validateConds(in.GetConds()); err != nil {
		return &npool.CountTxsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("CountTxs", "error", err)
		return &npool.CountTxsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountTxsResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteTx(ctx context.Context, in *npool.DeleteTxRequest) (*npool.DeleteTxResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteTx")
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
		return &npool.DeleteTxResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "tx", "crud", "Delete")

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("DeleteTx", "error", err)
		return &npool.DeleteTxResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteTxResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
