//nolint:nolintlint,dupl
package sms

import (
	"context"

	converter "github.com/NpoolPlatform/notif-manager/pkg/converter/template/sms"
	crud "github.com/NpoolPlatform/notif-manager/pkg/crud/template/sms"
	constant "github.com/NpoolPlatform/notif-manager/pkg/message/const"
	tracer "github.com/NpoolPlatform/notif-manager/pkg/tracer/template/sms"

	commontracer "github.com/NpoolPlatform/notif-manager/pkg/tracer"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/template/sms"

	"github.com/google/uuid"
)

func (s *Server) CreateSMSTemplate(ctx context.Context, in *npool.CreateSMSTemplateRequest) (*npool.CreateSMSTemplateResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateSMSTemplate")
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
		logger.Sugar().Errorw("CreateSMSTemplate", "error", err)
		return &npool.CreateSMSTemplateResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "app", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateSMSTemplate", "error", err)
		return &npool.CreateSMSTemplateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateSMSTemplateResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateSMSTemplates(ctx context.Context, in *npool.CreateSMSTemplatesRequest) (*npool.CreateSMSTemplatesResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateSMSTemplates")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		logger.Sugar().Errorw("CreateSMSTemplates", "error", "Infos is empty")
		return &npool.CreateSMSTemplatesResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "app", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorw("CreateSMSTemplates", "error", err)
		return &npool.CreateSMSTemplatesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateSMSTemplatesResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateSMSTemplate(ctx context.Context, in *npool.UpdateSMSTemplateRequest) (*npool.UpdateSMSTemplateResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateSMSTemplate")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("UpdateSMSTemplate", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateSMSTemplateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}
	if _, err := uuid.Parse(in.GetInfo().GetLangID()); err != nil {
		logger.Sugar().Errorw("UpdateSMSTemplate", "LangID", in.GetInfo().GetLangID(), "error", err)
		return &npool.UpdateSMSTemplateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "app", "crud", "Update")

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateSMSTemplate", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateSMSTemplateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateSMSTemplateResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetSMSTemplate(ctx context.Context, in *npool.GetSMSTemplateRequest) (*npool.GetSMSTemplateResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetSMSTemplate")
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
		logger.Sugar().Errorw("GetSMSTemplate", "ID", in.GetID(), "error", err)
		return &npool.GetSMSTemplateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "app", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("GetSMSTemplate", "ID", in.GetID(), "error", err)
		return &npool.GetSMSTemplateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetSMSTemplateResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetSMSTemplateOnly(ctx context.Context, in *npool.GetSMSTemplateOnlyRequest) (*npool.GetSMSTemplateOnlyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetSMSTemplateOnly")
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
		logger.Sugar().Errorw("GetSMSTemplateOnly", "error", err)
		return &npool.GetSMSTemplateOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetSMSTemplateOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetSMSTemplates(ctx context.Context, in *npool.GetSMSTemplatesRequest) (*npool.GetSMSTemplatesResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetSMSTemplates")
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
		logger.Sugar().Errorw("GetSMSTemplates", "error", err)
		return &npool.GetSMSTemplatesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetSMSTemplatesResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistSMSTemplate(ctx context.Context, in *npool.ExistSMSTemplateRequest) (*npool.ExistSMSTemplateResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistSMSTemplate")
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
		logger.Sugar().Errorw("ExistSMSTemplate", "ID", in.GetID(), "error", err)
		return &npool.ExistSMSTemplateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "app", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("ExistSMSTemplate", "ID", in.GetID(), "error", err)
		return &npool.ExistSMSTemplateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistSMSTemplateResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistSMSTemplateConds(
	ctx context.Context,
	in *npool.ExistSMSTemplateCondsRequest,
) (
	*npool.ExistSMSTemplateCondsResponse,
	error,
) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistSMSTemplateConds")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "app", "crud", "ExistSMSTemplateConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("ExistSMSTemplateConds", "error", err)
		return &npool.ExistSMSTemplateCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistSMSTemplateCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) DeleteSMSTemplate(ctx context.Context, in *npool.DeleteSMSTemplateRequest) (*npool.DeleteSMSTemplateResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteSMSTemplate")
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
		logger.Sugar().Errorw("DeleteSMSTemplate", "ID", in.GetID(), "error", err)
		return &npool.DeleteSMSTemplateResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "app", "crud", "Delete")

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("DeleteSMSTemplate", "ID", in.GetID(), "error", err)
		return &npool.DeleteSMSTemplateResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteSMSTemplateResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
