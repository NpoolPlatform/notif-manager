package notif

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"

	"time"

	"github.com/NpoolPlatform/notif-manager/pkg/db/ent/notif"
	tracer "github.com/NpoolPlatform/notif-manager/pkg/tracer/notif"

	constant "github.com/NpoolPlatform/notif-manager/pkg/message/const"
	commontracer "github.com/NpoolPlatform/notif-manager/pkg/tracer"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif"
	"github.com/NpoolPlatform/notif-manager/pkg/db"
	"github.com/NpoolPlatform/notif-manager/pkg/db/ent"

	"github.com/google/uuid"
)

func CreateSet(c *ent.NotifCreate, in *npool.NotifReq) (*ent.NotifCreate, error) {
	if in.ID != nil {
		c.SetID(uuid.MustParse(in.GetID()))
	}
	if in.AppID != nil {
		c.SetAppID(uuid.MustParse(in.GetAppID()))
	}
	if in.UserID != nil {
		c.SetUserID(uuid.MustParse(in.GetUserID()))
	}

	c.SetAlreadyRead(false)

	if in.LangID != nil {
		c.SetLangID(uuid.MustParse(in.GetLangID()))
	}
	if in.EventType != nil {
		c.SetEventType(in.GetEventType().String())
	}
	if in.UseTemplate != nil {
		c.SetUseTemplate(in.GetUseTemplate())
	}
	if in.Title != nil {
		c.SetTitle(in.GetTitle())
	}
	if in.Content != nil {
		c.SetContent(in.GetContent())
	}
	if in.Channels != nil {
		channels := []string{}
		for _, m := range in.GetChannels() {
			channels = append(channels, m.String())
		}
		c.SetChannels(channels)
	}
	c.SetEmailSend(false)

	return c, nil
}

func Create(ctx context.Context, in *npool.NotifReq) (*ent.Notif, error) {
	var info *ent.Notif
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
		c := cli.Notif.Create()
		stm, err := CreateSet(c, in)
		if err != nil {
			return err
		}
		info, err = stm.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func CreateBulk(ctx context.Context, in []*npool.NotifReq) ([]*ent.Notif, error) {
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "CreateBulk")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.TraceMany(span, in)

	rows := []*ent.Notif{}
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.NotifCreate, len(in))
		for i, info := range in {
			bulk[i] = tx.Notif.Create()
			bulk[i], err = CreateSet(bulk[i], info)
			if err != nil {
				return err
			}
		}
		rows, err = tx.Notif.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func UpdateSet(u *ent.NotifUpdateOne, in *npool.NotifReq) (*ent.NotifUpdateOne, error) {
	if in.Title != nil {
		u.SetTitle(in.GetTitle())
	}
	if in.Content != nil {
		u.SetContent(in.GetContent())
	}
	if in.Channels != nil {
		channels := []string{}
		for _, m := range in.GetChannels() {
			channels = append(channels, m.String())
		}
		u.SetChannels(channels)
	}
	if in.GetEmailSend() {
		u.SetEmailSend(in.GetEmailSend())
	}
	if in.AlreadyRead != nil {
		u.SetAlreadyRead(in.GetAlreadyRead())
	}
	if in.UseTemplate != nil {
		u.SetUseTemplate(in.GetUseTemplate())
	}
	return u, nil
}

func Update(ctx context.Context, in *npool.NotifReq) (*ent.Notif, error) {
	var info *ent.Notif
	var err error

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Update")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = tracer.Trace(span, in)

	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		info, err = tx.Notif.Query().Where(notif.ID(uuid.MustParse(in.GetID()))).ForUpdate().Only(_ctx)
		if err != nil {
			return err
		}

		stm, err := UpdateSet(info.Update(), in)
		if err != nil {
			return err
		}

		info, err = stm.Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func Row(ctx context.Context, id uuid.UUID) (*ent.Notif, error) {
	var info *ent.Notif
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
		info, err = cli.Notif.Query().Where(notif.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

//nolint:funlen,gocyclo
func SetQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.NotifQuery, error) {
	stm := cli.Notif.Query()
	if conds == nil {
		return stm, nil
	}
	if len(conds.GetChannels().GetValue()) > 0 {
		stm.Where(func(selector *sql.Selector) {
			channels := conds.GetChannels().GetValue()
			for i := 0; i < len(channels); i++ {
				if i == 0 {
					selector.Where(sqljson.ValueContains(notif.FieldChannels, channels[i]))
				} else {
					selector.Or().Where(sqljson.ValueContains(notif.FieldChannels, channels[i]))
				}
			}
		})
	}
	if conds.ID != nil {
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(notif.ID(uuid.MustParse(conds.GetID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid notif field")
		}
	}
	if conds.IDs != nil {
		switch conds.GetIDs().GetOp() {
		case cruder.IN:
			ids := []uuid.UUID{}
			for _, val := range conds.GetIDs().GetValue() {
				id, err := uuid.Parse(val)
				if err != nil {
					return nil, err
				}
				ids = append(ids, id)
			}
			stm.Where(notif.IDIn(ids...))
		default:
			return nil, fmt.Errorf("invalid notif field")
		}
	}
	if conds.AppID != nil {
		switch conds.GetAppID().GetOp() {
		case cruder.EQ:
			stm.Where(notif.AppID(uuid.MustParse(conds.GetAppID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid notif field")
		}
	}
	if conds.UserID != nil {
		switch conds.GetUserID().GetOp() {
		case cruder.EQ:
			stm.Where(notif.UserID(uuid.MustParse(conds.GetUserID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid notif field")
		}
	}
	if conds.AlreadyRead != nil {
		switch conds.GetAlreadyRead().GetOp() {
		case cruder.EQ:
			stm.Where(notif.AlreadyRead(conds.GetAlreadyRead().GetValue()))
		default:
			return nil, fmt.Errorf("invalid notif field")
		}
	}
	if conds.LangID != nil {
		switch conds.GetLangID().GetOp() {
		case cruder.EQ:
			stm.Where(notif.LangID(uuid.MustParse(conds.GetLangID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid notif field")
		}
	}
	if conds.EventType != nil {
		switch conds.GetEventType().GetOp() {
		case cruder.EQ:
			stm.Where(notif.EventType(npool.EventType(conds.GetEventType().GetValue()).String()))
		default:
			return nil, fmt.Errorf("invalid notif field")
		}
	}
	if conds.UseTemplate != nil {
		switch conds.GetUseTemplate().GetOp() {
		case cruder.EQ:
			stm.Where(notif.UseTemplate(conds.GetUseTemplate().GetValue()))
		default:
			return nil, fmt.Errorf("invalid notif field")
		}
	}

	if conds.EmailSend != nil {
		switch conds.GetEmailSend().GetOp() {
		case cruder.EQ:
			stm.Where(notif.EmailSend(conds.GetEmailSend().GetValue()))
		default:
			return nil, fmt.Errorf("invalid notif field")
		}
	}
	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.Notif, int, error) {
	var err error

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

	rows := []*ent.Notif{}
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

		rows, err = stm.
			Offset(offset).
			Order(ent.Desc(notif.FieldCreatedAt)).
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

func RowOnly(ctx context.Context, conds *npool.Conds) (*ent.Notif, error) {
	var info *ent.Notif
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
	var total int

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Count")
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
	exist := false

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "Exist")
	defer span.End()

	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, "db operation fail")
			span.RecordError(err)
		}
	}()

	span = commontracer.TraceID(span, id.String())

	err = db.WithClient(ctx, func(_ctx context.Context, cli *ent.Client) error {
		exist, err = cli.Notif.Query().Where(notif.ID(id)).Exist(_ctx)
		return err
	})
	if err != nil {
		return false, err
	}

	return exist, nil
}

func ExistConds(ctx context.Context, conds *npool.Conds) (bool, error) {
	var err error
	exist := false

	_, span := otel.Tracer(constant.ServiceName).Start(ctx, "ExistConds")
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

func Delete(ctx context.Context, id uuid.UUID) (*ent.Notif, error) {
	var info *ent.Notif
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
		info, err = cli.Notif.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
