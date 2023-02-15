//nolint:nolintlint,dupl
package tx

import (
	"context"

	converter "github.com/NpoolPlatform/notif-manager/pkg/converter/notif/tx"
	crud "github.com/NpoolPlatform/notif-manager/pkg/crud/notif/tx"
	commontracer "github.com/NpoolPlatform/notif-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/notif-manager/pkg/tracer/notif/tx"

	constant "github.com/NpoolPlatform/notif-manager/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif/tx"

	"github.com/google/uuid"
)

func (s *Server) CreateTran(
	ctx context.Context,
	in *npool.CreateTranRequest,
) (
	*npool.CreateTranResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateTran")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	if err := validate(in.GetInfo()); err != nil {
		return &npool.CreateTranResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "tx", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateTran", "error", err)
		return &npool.CreateTranResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateTranResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateTrans(
	ctx context.Context,
	in *npool.CreateTransRequest,
) (
	*npool.CreateTransResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateTrans")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		return &npool.CreateTransResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "tx", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorw("CreateTrans", "error", err)
		return &npool.CreateTransResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateTransResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateTran(
	ctx context.Context,
	in *npool.UpdateTranRequest,
) (
	*npool.UpdateTranResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateTran")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("UpdateTran", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateTranResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if in.Info == nil {
		logger.Sugar().Errorw("UpdateTran", "error", err)
		return &npool.UpdateTranResponse{}, status.Error(codes.InvalidArgument, "Info is empty")
	}

	if in.GetInfo().NotifState != nil {
		switch in.GetInfo().GetNotifState() {
		case npool.TxState_WaitTxSuccess:
		case npool.TxState_WaitSend:
		case npool.TxState_AlreadySend:
		default:
			logger.Sugar().Errorw("UpdateTran", "ID", in.GetInfo().GetID(), "error", err)
			return &npool.UpdateTranResponse{}, status.Error(codes.InvalidArgument, "NotifState is invalid")
		}
	}

	if in.GetInfo().NotifType != nil {
		switch in.GetInfo().GetNotifType() {
		case npool.TxType_Withdraw:
		default:
			logger.Sugar().Errorw("UpdateTxNotifType", "ID", in.GetInfo().GetID(), "error", err)
			return &npool.UpdateTranResponse{}, status.Error(codes.InvalidArgument, "NotifType is invalid")
		}
	}

	span = commontracer.TraceInvoker(span, "tx", "crud", "Update")

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateTran", "error", err)
		return &npool.UpdateTranResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateTranResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetTran(ctx context.Context, in *npool.GetTranRequest) (*npool.GetTranResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetTran")
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
		logger.Sugar().Errorw("GetTran", "ID", in.GetID(), "error", err)
		return &npool.GetTranResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "tx", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("GetTran", "error", err)
		return &npool.GetTranResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetTranResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetTranOnly(
	ctx context.Context,
	in *npool.GetTranOnlyRequest,
) (
	*npool.GetTranOnlyResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetTranOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := validateConds(in.GetConds()); err != nil {
		return &npool.GetTranOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "tx", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetTranOnly", "error", err)
		return &npool.GetTranOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetTranOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetTrans(
	ctx context.Context,
	in *npool.GetTransRequest,
) (
	*npool.GetTransResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetTrans")
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
		return &npool.GetTransResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "tx", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorw("GetTrans", "error", err)
		return &npool.GetTransResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetTransResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistTran(
	ctx context.Context,
	in *npool.ExistTranRequest,
) (
	*npool.ExistTranResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistTran")
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
		return &npool.ExistTranResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "tx", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("ExistTran", "error", err)
		return &npool.ExistTranResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistTranResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistTranConds(
	ctx context.Context,
	in *npool.ExistTranCondsRequest,
) (*npool.ExistTranCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistTranConds")
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
		return &npool.ExistTranCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("ExistTranConds", "error", err)
		return &npool.ExistTranCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistTranCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountTrans(ctx context.Context, in *npool.CountTransRequest) (*npool.CountTransResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountTrans")
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
		return &npool.CountTransResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("CountTrans", "error", err)
		return &npool.CountTransResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountTransResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteTran(ctx context.Context, in *npool.DeleteTranRequest) (*npool.DeleteTranResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteTran")
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
		return &npool.DeleteTranResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "tx", "crud", "Delete")

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("DeleteTran", "error", err)
		return &npool.DeleteTranResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteTranResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
