package frontend

import (
	"context"
	"fmt"
	"time"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"

	tracer "github.com/NpoolPlatform/notif-manager/pkg/tracer/template/frontend"

	constant "github.com/NpoolPlatform/notif-manager/pkg/message/const"
	commontracer "github.com/NpoolPlatform/notif-manager/pkg/tracer"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/template/frontend"
	"github.com/NpoolPlatform/notif-manager/pkg/db"
	"github.com/NpoolPlatform/notif-manager/pkg/db/ent"
	"github.com/NpoolPlatform/notif-manager/pkg/db/ent/frontendtemplate"
	"github.com/google/uuid"
)

func CreateSet(c *ent.FrontendTemplateCreate, info *npool.FrontendTemplateReq) *ent.FrontendTemplateCreate {
	if info.ID != nil {
		c.SetID(uuid.MustParse(info.GetID()))
	}
	if info.AppID != nil {
		c.SetAppID(uuid.MustParse(info.GetAppID()))
	}
	if info.LangID != nil {
		c.SetLangID(uuid.MustParse(info.GetLangID()))
	}
	if info.UsedFor != nil {
		c.SetUsedFor(info.GetUsedFor().String())
	}
	if info.Title != nil {
		c.SetTitle(info.GetTitle())
	}
	if info.Content != nil {
		c.SetContent(info.GetContent())
	}
	if info.Sender != nil {
		c.SetSender(info.GetSender())
	}

	return c
}
func Create(ctx context.Context, in *npool.FrontendTemplateReq) (*ent.FrontendTemplate, error) {
	var info *ent.FrontendTemplate
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Create")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		c := CreateSet(cli.FrontendTemplate.Create(), in)
		info, err = c.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateBulk(ctx context.Context, in []*npool.FrontendTemplateReq) ([]*ent.FrontendTemplate, error) {
	var err error
	rows := []*ent.FrontendTemplate{}

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateBulk")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceMany(span, in)

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.FrontendTemplateCreate, len(in))
		for i, info := range in {
			bulk[i] = CreateSet(tx.FrontendTemplate.Create(), info)
		}
		rows, err = tx.FrontendTemplate.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func Update(ctx context.Context, in *npool.FrontendTemplateReq) (*ent.FrontendTemplate, error) {
	var err error
	var info *ent.FrontendTemplate

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Update")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		u := UpdateSet(cli.FrontendTemplate.UpdateOneID(uuid.MustParse(in.GetID())), in)
		info, err = u.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func UpdateSet(u *ent.FrontendTemplateUpdateOne, in *npool.FrontendTemplateReq) *ent.FrontendTemplateUpdateOne {
	if in.LangID != nil {
		u.SetLangID(uuid.MustParse(in.GetLangID()))
	}
	if in.UsedFor != nil {
		u.SetUsedFor(in.GetUsedFor().String())
	}
	if in.Title != nil {
		u.SetTitle(in.GetTitle())
	}
	if in.Content != nil {
		u.SetContent(in.GetContent())
	}
	if in.Sender != nil {
		u.SetSender(in.GetSender())
	}
	return u
}

func Row(ctx context.Context, id uuid.UUID) (*ent.FrontendTemplate, error) {
	var info *ent.FrontendTemplate
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Row")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id.String())

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.FrontendTemplate.Query().Where(frontendtemplate.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

//nolint:funlen,gocyclo
func SetQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.FrontendTemplateQuery, error) {
	stm := cli.FrontendTemplate.Query()

	if conds == nil {
		return stm, nil
	}

	if conds.ID != nil {
		id, err := uuid.Parse(conds.GetID().GetValue())
		if err != nil {
			return nil, err
		}

		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(frontendtemplate.ID(id))
		default:
			return nil, fmt.Errorf("invalid frontend field")
		}
	}

	if conds.AppID != nil {
		id, err := uuid.Parse(conds.GetAppID().GetValue())
		if err != nil {
			return nil, err
		}

		switch conds.GetAppID().GetOp() {
		case cruder.EQ:
			stm.Where(frontendtemplate.AppID(id))
		default:
			return nil, fmt.Errorf("invalid frontend field")
		}
	}

	if conds.LangID != nil {
		id, err := uuid.Parse(conds.GetLangID().GetValue())
		if err != nil {
			return nil, err
		}

		switch conds.GetLangID().GetOp() {
		case cruder.EQ:
			stm.Where(frontendtemplate.LangID(id))
		default:
			return nil, fmt.Errorf("invalid frontend field")
		}
	}

	if conds.UsedFor != nil {
		switch conds.GetUsedFor().GetOp() {
		case cruder.EQ:
			stm.Where(frontendtemplate.UsedFor(basetypes.UsedFor(conds.GetUsedFor().GetValue()).String()))
		default:
			return nil, fmt.Errorf("invalid frontend field")
		}
	}

	if conds.AppIDs != nil {
		switch conds.GetAppIDs().GetOp() {
		case cruder.IN:
			ids := []uuid.UUID{}
			for _, val := range conds.GetAppIDs().GetValue() {
				id, err := uuid.Parse(val)
				if err != nil {
					return nil, err
				}
				ids = append(ids, id)
			}
			stm.Where(frontendtemplate.AppIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid frontend field")
		}
	}
	if conds.LangIDs != nil {
		switch conds.GetLangIDs().GetOp() {
		case cruder.IN:
			ids := []uuid.UUID{}
			for _, val := range conds.GetLangIDs().GetValue() {
				id, err := uuid.Parse(val)
				if err != nil {
					return nil, err
				}
				ids = append(ids, id)
			}
			stm.Where(frontendtemplate.LangIDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid frontend field")
		}
	}
	if conds.UsedFors != nil {
		switch conds.GetUsedFors().GetOp() {
		case cruder.IN:
			stm.Where(frontendtemplate.UsedForIn(conds.GetUsedFors().GetValue()...))
		default:
			return nil, fmt.Errorf("invalid frontend field")
		}
	}
	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.FrontendTemplate, int, error) {
	var err error
	rows := []*ent.FrontendTemplate{}
	var total int

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Rows")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)
	span = commontracer.TraceOffsetLimit(span, offset, limit)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := SetQueryConds(conds, cli)
		if err != nil {
			return err
		}
		total, err = stm.Count(_ctx)
		if err != nil {
			return err
		}

		rows, err = stm.
			Offset(offset).
			Order(ent.Desc(frontendtemplate.FieldUpdatedAt)).
			Limit(limit).
			All(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, 0, err
	}

	return rows, total, nil
}

func RowOnly(ctx context.Context, conds *npool.Conds) (*ent.FrontendTemplate, error) {
	var info *ent.FrontendTemplate
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "RowOnly")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := SetQueryConds(conds, cli)
		if err != nil {
			return err
		}

		info, err = stm.Only(_ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return nil
			}
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func Count(ctx context.Context, conds *npool.Conds) (uint32, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Count")
	defer span.End()
	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)
	var total int

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := SetQueryConds(conds, cli)
		if err != nil {
			return err
		}

		total, err = stm.Count(_ctx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return 0, err
	}

	return uint32(total), nil
}

func Exist(ctx context.Context, id uuid.UUID) (bool, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Exist")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id.String())

	exist := false

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.FrontendTemplate.Query().Where(frontendtemplate.ID(id)).Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func ExistConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistConds")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceConds(span, conds)
	exist := false

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		stm, err := SetQueryConds(conds, cli)
		if err != nil {
			return err
		}

		exist, err = stm.Exist(_ctx)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func Delete(ctx context.Context, id uuid.UUID) (*ent.FrontendTemplate, error) {
	var info *ent.FrontendTemplate
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Delete")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id.String())

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		info, err = cli.FrontendTemplate.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
