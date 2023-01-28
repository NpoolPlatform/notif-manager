//nolint:nolintlint,dupl
package readstate

import (
	"context"

	converter "github.com/NpoolPlatform/notif-manager/pkg/converter/announcement/readstate"
	crud "github.com/NpoolPlatform/notif-manager/pkg/crud/announcement/readstate"
	commontracer "github.com/NpoolPlatform/notif-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/notif-manager/pkg/tracer/announcement/readstate"

	constant "github.com/NpoolPlatform/notif-manager/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/announcement/readstate"

	"github.com/google/uuid"
)

func (s *Server) CreateReadState(ctx context.Context, in *npool.CreateReadStateRequest) (*npool.CreateReadStateResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateReadState")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	if err := validate(in.GetInfo()); err != nil {
		return &npool.CreateReadStateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "announcement/readstate", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateReadState", "error", err)
		return &npool.CreateReadStateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateReadStateResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateReadStates(ctx context.Context, in *npool.CreateReadStatesRequest) (*npool.CreateReadStatesResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateReadStates")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		return &npool.CreateReadStatesResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "announcement/readstate", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorw("CreateReadStates", "error", err)
		return &npool.CreateReadStatesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateReadStatesResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateReadState(ctx context.Context, in *npool.UpdateReadStateRequest) (*npool.UpdateReadStateResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateReadState")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("UpdateReadState", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateReadStateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "announcement/readstate", "crud", "Update")

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateReadState", "error", err)
		return &npool.UpdateReadStateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateReadStateResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetReadState(ctx context.Context, in *npool.GetReadStateRequest) (*npool.GetReadStateResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetReadState")
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
		logger.Sugar().Errorw("GetReadState", "ID", in.GetID(), "error", err)
		return &npool.GetReadStateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "announcement/readstate", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("GetReadState", "error", err)
		return &npool.GetReadStateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetReadStateResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetReadStateOnly(ctx context.Context, in *npool.GetReadStateOnlyRequest) (*npool.GetReadStateOnlyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetReadStateOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := validateConds(in.GetConds()); err != nil {
		return &npool.GetReadStateOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "announcement/readstate", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetReadStateOnly", "error", err)
		return &npool.GetReadStateOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetReadStateOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetReadStates(ctx context.Context, in *npool.GetReadStatesRequest) (*npool.GetReadStatesResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetReadStates")
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
		return &npool.GetReadStatesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "announcement/readstate", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorw("GetReadStates", "error", err)
		return &npool.GetReadStatesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetReadStatesResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistReadState(ctx context.Context, in *npool.ExistReadStateRequest) (*npool.ExistReadStateResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistReadState")
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
		return &npool.ExistReadStateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "announcement/readstate", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("ExistReadState", "error", err)
		return &npool.ExistReadStateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistReadStateResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistReadStateConds(ctx context.Context,
	in *npool.ExistReadStateCondsRequest) (*npool.ExistReadStateCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistReadStateConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "announcement/readstate", "crud", "ExistConds")

	if err := validateConds(in.GetConds()); err != nil {
		return &npool.ExistReadStateCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("ExistReadStateConds", "error", err)
		return &npool.ExistReadStateCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistReadStateCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountReadStates(ctx context.Context, in *npool.CountReadStatesRequest) (*npool.CountReadStatesResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountReadStates")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "announcement/readstate", "crud", "Count")

	if err := validateConds(in.GetConds()); err != nil {
		return &npool.CountReadStatesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("CountReadStates", "error", err)
		return &npool.CountReadStatesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountReadStatesResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteReadState(ctx context.Context, in *npool.DeleteReadStateRequest) (*npool.DeleteReadStateResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteReadState")
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
		return &npool.DeleteReadStateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "announcement/readstate", "crud", "Delete")

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("DeleteReadState", "error", err)
		return &npool.DeleteReadStateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteReadStateResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
