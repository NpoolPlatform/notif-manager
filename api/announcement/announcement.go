//nolint:nolintlint,dupl
package announcement

import (
	"context"

	converter "github.com/NpoolPlatform/notif-manager/pkg/converter/announcement"
	crud "github.com/NpoolPlatform/notif-manager/pkg/crud/announcement"
	commontracer "github.com/NpoolPlatform/notif-manager/pkg/tracer"
	tracer "github.com/NpoolPlatform/notif-manager/pkg/tracer/announcement"

	constant "github.com/NpoolPlatform/notif-manager/pkg/message/const"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/announcement"

	"github.com/google/uuid"
)

func (s *Server) CreateAnnouncement(
	ctx context.Context,
	in *npool.CreateAnnouncementRequest,
) (
	*npool.CreateAnnouncementResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateAnnouncement")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	if err := validate(in.GetInfo()); err != nil {
		return &npool.CreateAnnouncementResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "announcement", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateAnnouncement", "error", err)
		return &npool.CreateAnnouncementResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAnnouncementResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateAnnouncements(
	ctx context.Context,
	in *npool.CreateAnnouncementsRequest,
) (
	*npool.CreateAnnouncementsResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateAnnouncements")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		return &npool.CreateAnnouncementsResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "announcement", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorw("CreateAnnouncements", "error", err)
		return &npool.CreateAnnouncementsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateAnnouncementsResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateAnnouncement(
	ctx context.Context,
	in *npool.UpdateAnnouncementRequest,
) (
	*npool.UpdateAnnouncementResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateAnnouncement")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("UpdateAnnouncement", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateAnnouncementResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if in.GetInfo().Title != nil && in.GetInfo().GetTitle() == "" {
		logger.Sugar().Errorw("UpdateAnnouncement", "Title", in.GetInfo().GetTitle())
		return &npool.UpdateAnnouncementResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if in.GetInfo().Content != nil && in.GetInfo().GetContent() == "" {
		logger.Sugar().Errorw("UpdateAnnouncement", "Content", in.GetInfo().GetContent())
		return &npool.UpdateAnnouncementResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "announcement", "crud", "Update")

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateAnnouncement", "error", err)
		return &npool.UpdateAnnouncementResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateAnnouncementResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetAnnouncement(
	ctx context.Context,
	in *npool.GetAnnouncementRequest,
) (
	*npool.GetAnnouncementResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAnnouncement")
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
		logger.Sugar().Errorw("GetAnnouncement", "ID", in.GetID(), "error", err)
		return &npool.GetAnnouncementResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "announcement", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("GetAnnouncement", "error", err)
		return &npool.GetAnnouncementResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAnnouncementResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetAnnouncementOnly(
	ctx context.Context,
	in *npool.GetAnnouncementOnlyRequest,
) (
	*npool.GetAnnouncementOnlyResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAnnouncementOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if err := validateConds(in.GetConds()); err != nil {
		return &npool.GetAnnouncementOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "announcement", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetAnnouncementOnly", "error", err)
		return &npool.GetAnnouncementOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAnnouncementOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetAnnouncements(
	ctx context.Context,
	in *npool.GetAnnouncementsRequest,
) (
	*npool.GetAnnouncementsResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetAnnouncements")
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
		return &npool.GetAnnouncementsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "announcement", "crud", "Rows")

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorw("GetAnnouncements", "error", err)
		return &npool.GetAnnouncementsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetAnnouncementsResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistAnnouncement(
	ctx context.Context,
	in *npool.ExistAnnouncementRequest,
) (
	*npool.ExistAnnouncementResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistAnnouncement")
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
		return &npool.ExistAnnouncementResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "announcement", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("ExistAnnouncement", "error", err)
		return &npool.ExistAnnouncementResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAnnouncementResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistAnnouncementConds(
	ctx context.Context,
	in *npool.ExistAnnouncementCondsRequest,
) (
	*npool.ExistAnnouncementCondsResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistAnnouncementConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "announcement", "crud", "ExistConds")

	if err := validateConds(in.GetConds()); err != nil {
		return &npool.ExistAnnouncementCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("ExistAnnouncementConds", "error", err)
		return &npool.ExistAnnouncementCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistAnnouncementCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) CountAnnouncements(
	ctx context.Context,
	in *npool.CountAnnouncementsRequest,
) (
	*npool.CountAnnouncementsResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CountAnnouncements")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "announcement", "crud", "Count")

	if err := validateConds(in.GetConds()); err != nil {
		return &npool.CountAnnouncementsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	total, err := crud.Count(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("CountAnnouncements", "error", err)
		return &npool.CountAnnouncementsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CountAnnouncementsResponse{
		Info: total,
	}, nil
}

func (s *Server) DeleteAnnouncement(
	ctx context.Context,
	in *npool.DeleteAnnouncementRequest,
) (
	*npool.DeleteAnnouncementResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteAnnouncement")
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
		return &npool.DeleteAnnouncementResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "announcement", "crud", "Delete")

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("DeleteAnnouncement", "error", err)
		return &npool.DeleteAnnouncementResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteAnnouncementResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
