//nolint:nolintlint,dupl
package email

import (
	"context"

	converter "github.com/NpoolPlatform/notif-manager/pkg/converter/template/email"
	crud "github.com/NpoolPlatform/notif-manager/pkg/crud/template/email"
	constant "github.com/NpoolPlatform/notif-manager/pkg/message/const"
	tracer "github.com/NpoolPlatform/notif-manager/pkg/tracer/template/email"

	commontracer "github.com/NpoolPlatform/notif-manager/pkg/tracer"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/template/email"

	"github.com/google/uuid"
)

func (s *Server) CreateEmailTemplate(
	ctx context.Context,
	in *npool.CreateEmailTemplateRequest,
) (
	*npool.CreateEmailTemplateResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateEmailTemplate")
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
		logger.Sugar().Errorw("CreateEmailTemplate", "error", err)
		return &npool.CreateEmailTemplateResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "app", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateEmailTemplate", "error", err)
		return &npool.CreateEmailTemplateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateEmailTemplateResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateEmailTemplates(
	ctx context.Context,
	in *npool.CreateEmailTemplatesRequest,
) (
	*npool.CreateEmailTemplatesResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateEmailTemplates")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		logger.Sugar().Errorw("CreateEmailTemplates", "error", "Infos is empty")
		return &npool.CreateEmailTemplatesResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "app", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorw("CreateEmailTemplates", "error", err)
		return &npool.CreateEmailTemplatesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateEmailTemplatesResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateEmailTemplate(
	ctx context.Context,
	in *npool.UpdateEmailTemplateRequest,
) (
	*npool.UpdateEmailTemplateResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateEmailTemplate")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("UpdateEmailTemplate", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateEmailTemplateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	if _, err := uuid.Parse(in.GetInfo().GetLangID()); err != nil {
		logger.Sugar().Errorw("UpdateSMSTemplate", "LangID", in.GetInfo().GetLangID(), "error", err)
		return &npool.UpdateEmailTemplateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "app", "crud", "Update")

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateEmailTemplate", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateEmailTemplateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateEmailTemplateResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetEmailTemplate(
	ctx context.Context,
	in *npool.GetEmailTemplateRequest,
) (
	*npool.GetEmailTemplateResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetEmailTemplate")
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
		logger.Sugar().Errorw("GetEmailTemplate", "ID", in.GetID(), "error", err)
		return &npool.GetEmailTemplateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "app", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("GetEmailTemplate", "ID", in.GetID(), "error", err)
		return &npool.GetEmailTemplateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetEmailTemplateResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetEmailTemplateOnly(
	ctx context.Context,
	in *npool.GetEmailTemplateOnlyRequest,
) (
	*npool.GetEmailTemplateOnlyResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetEmailTemplateOnly")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "app", "crud", "RowOnly")

	info, err := crud.RowOnly(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("GetEmailTemplateOnly", "error", err)
		return &npool.GetEmailTemplateOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetEmailTemplateOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetEmailTemplates(
	ctx context.Context,
	in *npool.GetEmailTemplatesRequest,
) (
	*npool.GetEmailTemplatesResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetEmailTemplates")
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

	rows, total, err := crud.Rows(ctx, in.GetConds(), int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorw("GetEmailTemplates", "error", err)
		return &npool.GetEmailTemplatesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetEmailTemplatesResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistEmailTemplate(
	ctx context.Context,
	in *npool.ExistEmailTemplateRequest,
) (
	*npool.ExistEmailTemplateResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistEmailTemplate")
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
		logger.Sugar().Errorw("ExistEmailTemplate", "ID", in.GetID(), "error", err)
		return &npool.ExistEmailTemplateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "app", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("ExistEmailTemplate", "ID", in.GetID(), "error", err)
		return &npool.ExistEmailTemplateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistEmailTemplateResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistEmailTemplateConds(
	ctx context.Context,
	in *npool.ExistEmailTemplateCondsRequest,
) (
	*npool.ExistEmailTemplateCondsResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistEmailTemplateConds")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "app", "crud", "ExistEmailTemplateConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("ExistEmailTemplateConds", "error", err)
		return &npool.ExistEmailTemplateCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistEmailTemplateCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) DeleteEmailTemplate(
	ctx context.Context,
	in *npool.DeleteEmailTemplateRequest,
) (
	*npool.DeleteEmailTemplateResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteEmailTemplate")
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
		logger.Sugar().Errorw("DeleteEmailTemplate", "ID", in.GetID(), "error", err)
		return &npool.DeleteEmailTemplateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "app", "crud", "Delete")

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("DeleteEmailTemplate", "ID", in.GetID(), "error", err)
		return &npool.DeleteEmailTemplateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteEmailTemplateResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
