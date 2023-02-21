//nolint:nolintlint,dupl
package contact

import (
	"context"

	converter "github.com/NpoolPlatform/notif-manager/pkg/converter/contact"
	crud "github.com/NpoolPlatform/notif-manager/pkg/crud/contact"
	constant "github.com/NpoolPlatform/notif-manager/pkg/message/const"
	tracer "github.com/NpoolPlatform/notif-manager/pkg/tracer/contact"

	commontracer "github.com/NpoolPlatform/notif-manager/pkg/tracer"

	"go.opentelemetry.io/otel"
	scodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/contact"

	"github.com/google/uuid"
)

func (s *Server) CreateContact(ctx context.Context, in *npool.CreateContactRequest) (*npool.CreateContactResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateContact")
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
		logger.Sugar().Errorw("CreateContact", "error", err)
		return &npool.CreateContactResponse{}, err
	}

	span = commontracer.TraceInvoker(span, "app", "crud", "Create")

	info, err := crud.Create(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("CreateContact", "error", err)
		return &npool.CreateContactResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateContactResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) CreateContacts(ctx context.Context, in *npool.CreateContactsRequest) (*npool.CreateContactsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateContacts")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	if len(in.GetInfos()) == 0 {
		logger.Sugar().Errorw("CreateContacts", "error", "Infos is empty")
		return &npool.CreateContactsResponse{}, status.Error(codes.InvalidArgument, "Infos is empty")
	}

	span = tracer.TraceMany(span, in.GetInfos())
	span = commontracer.TraceInvoker(span, "app", "crud", "CreateBulk")

	rows, err := crud.CreateBulk(ctx, in.GetInfos())
	if err != nil {
		logger.Sugar().Errorw("CreateContacts", "error", err)
		return &npool.CreateContactsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.CreateContactsResponse{
		Infos: converter.Ent2GrpcMany(rows),
	}, nil
}

func (s *Server) UpdateContact(ctx context.Context, in *npool.UpdateContactRequest) (*npool.UpdateContactResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "UpdateContact")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in.GetInfo())

	if _, err := uuid.Parse(in.GetInfo().GetID()); err != nil {
		logger.Sugar().Errorw("UpdateContact", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateContactResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "app", "crud", "Update")

	info, err := crud.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorw("UpdateContact", "ID", in.GetInfo().GetID(), "error", err)
		return &npool.UpdateContactResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateContactResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetContact(ctx context.Context, in *npool.GetContactRequest) (*npool.GetContactResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetContact")
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
		logger.Sugar().Errorw("GetContact", "ID", in.GetID(), "error", err)
		return &npool.GetContactResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "app", "crud", "Row")

	info, err := crud.Row(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("GetContact", "ID", in.GetID(), "error", err)
		return &npool.GetContactResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetContactResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetContactOnly(ctx context.Context, in *npool.GetContactOnlyRequest) (*npool.GetContactOnlyResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetContactOnly")
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
		logger.Sugar().Errorw("GetContactOnly", "error", err)
		return &npool.GetContactOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetContactOnlyResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}

func (s *Server) GetContacts(ctx context.Context, in *npool.GetContactsRequest) (*npool.GetContactsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "GetContacts")
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
		logger.Sugar().Errorw("GetContacts", "error", err)
		return &npool.GetContactsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetContactsResponse{
		Infos: converter.Ent2GrpcMany(rows),
		Total: uint32(total),
	}, nil
}

func (s *Server) ExistContact(ctx context.Context, in *npool.ExistContactRequest) (*npool.ExistContactResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistContact")
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
		logger.Sugar().Errorw("ExistContact", "ID", in.GetID(), "error", err)
		return &npool.ExistContactResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "app", "crud", "Exist")

	exist, err := crud.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("ExistContact", "ID", in.GetID(), "error", err)
		return &npool.ExistContactResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistContactResponse{
		Info: exist,
	}, nil
}

func (s *Server) ExistContactConds(ctx context.Context, in *npool.ExistContactCondsRequest) (*npool.ExistContactCondsResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistContactConds")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(scodes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, in.GetConds())
	span = commontracer.TraceInvoker(span, "app", "crud", "ExistContactConds")

	exist, err := crud.ExistConds(ctx, in.GetConds())
	if err != nil {
		logger.Sugar().Errorw("ExistContactConds", "error", err)
		return &npool.ExistContactCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistContactCondsResponse{
		Info: exist,
	}, nil
}

func (s *Server) DeleteContact(ctx context.Context, in *npool.DeleteContactRequest) (*npool.DeleteContactResponse, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "DeleteContact")
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
		logger.Sugar().Errorw("DeleteContact", "ID", in.GetID(), "error", err)
		return &npool.DeleteContactResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	span = commontracer.TraceInvoker(span, "app", "crud", "Delete")

	info, err := crud.Delete(ctx, id)
	if err != nil {
		logger.Sugar().Errorw("DeleteContact", "ID", in.GetID(), "error", err)
		return &npool.DeleteContactResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteContactResponse{
		Info: converter.Ent2Grpc(info),
	}, nil
}
