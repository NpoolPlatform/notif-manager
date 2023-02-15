//nolint:nolintlint,dupl
package frontend

import (
	"context"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	converter "github.com/NpoolPlatform/notif-manager/pkg/converter/template/frontend"
	crud "github.com/NpoolPlatform/notif-manager/pkg/crud/template/frontend"
	constant "github.com/NpoolPlatform/notif-manager/pkg/message/const"
	tracer "github.com/NpoolPlatform/notif-manager/pkg/tracer/template/frontend"

	commontracer "github.com/NpoolPlatform/notif-manager/pkg/tracer"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/template/frontend"

	"github.com/google/uuid"
)

func (s *Server) CreateFrontendTemplate(
	ctx context.Context,
	in *npool.CreateFrontendTemplateRequest,
) (
	*npool.CreateFrontendTemplateResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateFrontendTemplate")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	err = validate(in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateFrontendTemplate", "error", err)
		return &npool.CreateFrontendTemplateResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "app", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateFrontendTemplate", "error", err)
		return &npool.CreateFrontendTemplateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateFrontendTemplateResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateFrontendTemplates(
	ctx context.Context,
	in *npool.CreateFrontendTemplatesRequest,
) (
	*npool.CreateFrontendTemplatesResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateFrontendTemplates")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		logger.Sugar().Errorw("CreateFrontendTemplates", "error", "Infos is empty")
		return &npool.CreateFrontendTemplatesResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "app", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorw("CreateFrontendTemplates", "error", err)
		return &npool.CreateFrontendTemplatesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateFrontendTemplatesResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

//nolint:gocyclo
func (s *Server) UpdateFrontendTemplate(
	ctx context.Context,
	in *npool.UpdateFrontendTemplateRequest,
) (
	*npool.UpdateFrontendTemplateResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateFrontendTemplate")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("UpdateFrontendTemplate", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateFrontendTemplateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if _, err := uuid.Parse(in.GetInfo().GetLangID()); err != nil && in.Info.LangID != nil {
		logger.Sugar().Errorw("UpdateSMSTemplate", "LangID", in.GetInfo().GetLangID(), "error", err)
		return &npool.UpdateFrontendTemplateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if in.GetInfo().UsedFor != nil {
		switch in.GetInfo().GetUsedFor() {
		case basetypes.UsedFor_WithdrawalRequest:
		case basetypes.UsedFor_WithdrawalCompleted:
		case basetypes.UsedFor_DepositReceived:
		case basetypes.UsedFor_KYCApproved:
		case basetypes.UsedFor_KYCRejected:
		case basetypes.UsedFor_Announcement:
		default:
			logger.Sugar().Errorw("UpdateFrontendTemplate", "ID", in.GetInfo().GetID(), "error", "UsedFor is invalid")
			return &npool.UpdateFrontendTemplateResponse{}, status.Error(codes.Internal, "UsedFor is invalid")
		}
	}

	if in.GetInfo().Title != nil && in.GetInfo().GetTitle() == "" {
		logger.Sugar().Errorw("UpdateFrontendTemplate", "ID", in.GetInfo().GetID(), "error", "Title is empty")
		return &npool.UpdateFrontendTemplateResponse{}, status.Error(codes.Internal, "Title is empty")
	}

	if in.GetInfo().Content != nil && in.GetInfo().GetContent() == "" {
		logger.Sugar().Errorw("UpdateFrontendTemplate", "ID", in.GetInfo().GetID(), "error", "Content is empty")
		return &npool.UpdateFrontendTemplateResponse{}, status.Error(codes.Internal, "Content is empty")
	}

	span = commontracer.TraceInvoker(span, "app", "crud", "Update")

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateFrontendTemplate", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateFrontendTemplateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateFrontendTemplateResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetFrontendTemplate(
	ctx context.Context,
	in *npool.GetFrontendTemplateRequest,
) (
	*npool.GetFrontendTemplateResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetFrontendTemplate")
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
		logger.Sugar().Errorw("GetFrontendTemplate", "ID", in.GetID(), "error", err)
		return &npool.GetFrontendTemplateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "app", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("GetFrontendTemplate", "ID", in.GetID(), "error", err)
		return &npool.GetFrontendTemplateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetFrontendTemplateResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetFrontendTemplateOnly(
	ctx context.Context,
	in *npool.GetFrontendTemplateOnlyRequest,
) (
	*npool.GetFrontendTemplateOnlyResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetFrontendTemplateOnly")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "app", "crud", "RowOnly")

	err = validateConds(in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetFrontendTemplateOnly", "error", err)
		return &npool.GetFrontendTemplateOnlyResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetFrontendTemplateOnly", "error", err)
		return &npool.GetFrontendTemplateOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetFrontendTemplateOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetFrontendTemplates(
	ctx context.Context,
	in *npool.GetFrontendTemplatesRequest,
) (
	*npool.GetFrontendTemplatesResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetFrontendTemplates")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceOffsetLimit(span, int(in.GetOffset()), int(in.GetLimit()))
	span = commontracer.TraceInvoker(span, "app", "crud", "Rows")

	err = validateConds(in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetFrontendTemplateOnly", "error", err)
		return &npool.GetFrontendTemplatesResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorw("GetFrontendTemplates", "error", err)
		return &npool.GetFrontendTemplatesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetFrontendTemplatesResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistFrontendTemplate(
	ctx context.Context,
	in *npool.ExistFrontendTemplateRequest,
) (
	*npool.ExistFrontendTemplateResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistFrontendTemplate")
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
		logger.Sugar().Errorw("ExistFrontendTemplate", "ID", in.GetID(), "error", err)
		return &npool.ExistFrontendTemplateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "app", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("ExistFrontendTemplate", "ID", in.GetID(), "error", err)
		return &npool.ExistFrontendTemplateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistFrontendTemplateResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistFrontendTemplateConds(
	ctx context.Context,
	in *npool.ExistFrontendTemplateCondsRequest,
) (
	*npool.ExistFrontendTemplateCondsResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistFrontendTemplateConds")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "app", "crud", "ExistFrontendTemplateConds")

	err = validateConds(in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetFrontendTemplateOnly", "error", err)
		return &npool.ExistFrontendTemplateCondsResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("ExistFrontendTemplateConds", "error", err)
		return &npool.ExistFrontendTemplateCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistFrontendTemplateCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) DeleteFrontendTemplate(
	ctx context.Context,
	in *npool.DeleteFrontendTemplateRequest,
) (
	*npool.DeleteFrontendTemplateResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteFrontendTemplate")
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
		logger.Sugar().Errorw("DeleteFrontendTemplate", "ID", in.GetID(), "error", err)
		return &npool.DeleteFrontendTemplateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "app", "crud", "Delete")

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("DeleteFrontendTemplate", "ID", in.GetID(), "error", err)
		return &npool.DeleteFrontendTemplateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteFrontendTemplateResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
