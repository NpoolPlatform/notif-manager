//nolint:nolintlint,dupl
package notif

import (
	"context"

	converter "github.com/NpoolPlatform/notif-manager/pkg/converter/notif"
	crud "github.com/NpoolPlatform/notif-manager/pkg/crud/notif"
	commontracer "github.com/NpoolPlatform/notif-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/notif-manager/pkg/tracer/notif"

	constant "github.com/NpoolPlatform/notif-manager/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif"

	"github.com/google/uuid"
)

func (s *Server) CreateNotif(ctx context.Context, in *npool.CreateNotifRequest) (*npool.CreateNotifResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateNotif")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	if err := validate(in.GetInfo()); err != nil {
		return &npool.CreateNotifResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "notif", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateNotif", "error", err)
		return &npool.CreateNotifResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateNotifResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateNotifs(ctx context.Context, in *npool.CreateNotifsRequest) (*npool.CreateNotifsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateNotifs")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		return &npool.CreateNotifsResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "notif", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorw("CreateNotifs", "error", err)
		return &npool.CreateNotifsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateNotifsResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateNotif(ctx context.Context, in *npool.UpdateNotifRequest) (*npool.UpdateNotifResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateNotif")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("UpdateNotif", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateNotifResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if in.GetInfo().Title != nil && in.GetInfo().GetTitle() == "" {
		logger.Sugar().Errorw("UpdateNotif", "Title", in.GetInfo().GetTitle())
		return &npool.UpdateNotifResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if in.GetInfo().Content != nil && in.GetInfo().GetContent() == "" {
		logger.Sugar().Errorw("UpdateNotif", "Content", in.GetInfo().GetContent())
		return &npool.UpdateNotifResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "notif", "crud", "Update")

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateNotif", "error", err)
		return &npool.UpdateNotifResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateNotifResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetNotif(ctx context.Context, in *npool.GetNotifRequest) (*npool.GetNotifResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetNotif")
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
		logger.Sugar().Errorw("GetNotif", "ID", in.GetID(), "error", err)
		return &npool.GetNotifResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "notif", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("GetNotif", "error", err)
		return &npool.GetNotifResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetNotifResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetNotifOnly(ctx context.Context, in *npool.GetNotifOnlyRequest) (*npool.GetNotifOnlyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetNotifOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := validateConds(in.GetConds()); err != nil {
		return &npool.GetNotifOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "notif", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetNotifOnly", "error", err)
		return &npool.GetNotifOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetNotifOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetNotifs(ctx context.Context, in *npool.GetNotifsRequest) (*npool.GetNotifsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetNotifs")
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
		return &npool.GetNotifsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "notif", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorw("GetNotifs", "error", err)
		return &npool.GetNotifsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetNotifsResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistNotif(ctx context.Context, in *npool.ExistNotifRequest) (*npool.ExistNotifResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistNotif")
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
		return &npool.ExistNotifResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "notif", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("ExistNotif", "error", err)
		return &npool.ExistNotifResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistNotifResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistNotifConds(ctx context.Context,
	in *npool.ExistNotifCondsRequest) (*npool.ExistNotifCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistNotifConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "notif", "crud", "ExistConds")

	if err := validateConds(in.GetConds()); err != nil {
		return &npool.ExistNotifCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("ExistNotifConds", "error", err)
		return &npool.ExistNotifCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistNotifCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountNotifs(ctx context.Context, in *npool.CountNotifsRequest) (*npool.CountNotifsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountNotifs")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "notif", "crud", "Count")

	if err := validateConds(in.GetConds()); err != nil {
		return &npool.CountNotifsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("CountNotifs", "error", err)
		return &npool.CountNotifsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountNotifsResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteNotif(ctx context.Context, in *npool.DeleteNotifRequest) (*npool.DeleteNotifResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteNotif")
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
		return &npool.DeleteNotifResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "notif", "crud", "Delete")

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("DeleteNotif", "error", err)
		return &npool.DeleteNotifResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteNotifResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
