package channel

import (
	"context"
	"fmt"

	"github.com/NpoolPlatform/message/npool/notif/mgr/v1/channel"

	"time"

	"github.com/NpoolPlatform/notif-manager/pkg/db/ent/notifchannel"
	tracer "github.com/NpoolPlatform/notif-manager/pkg/tracer/notif/channel"

	constant "github.com/NpoolPlatform/notif-manager/pkg/message/const"
	commontracer "github.com/NpoolPlatform/notif-manager/pkg/tracer"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif/channel"
	"github.com/NpoolPlatform/notif-manager/pkg/db"
	"github.com/NpoolPlatform/notif-manager/pkg/db/ent"

	"github.com/google/uuid"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
)

func CreateSet(c *ent.NotifChannelCreate, in *npool.ChannelReq) (*ent.NotifChannelCreate, error) {
	if in.ID != nil {
		c.SetID(uuid.MustParse(in.GetID()))
	}
	if in.AppID != nil {
		c.SetAppID(uuid.MustParse(in.GetAppID()))
	}
	if in.EventType != nil {
		c.SetEventType(in.GetEventType().String())
	}
	if in.Channel != nil {
		c.SetChannel(in.GetChannel().String())
	}

	return c, nil
}

func Create(ctx context.Context, in *npool.ChannelReq) (*ent.NotifChannel, error) {
	var info *ent.NotifChannel
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
		c := cli.NotifChannel.Create()
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

func CreateBulk(ctx context.Context, in []*npool.ChannelReq) ([]*ent.NotifChannel, error) {
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

	rows := []*ent.NotifChannel{}
	err = db.WithTx(ctx, func(_ctx context.Context, tx *ent.Tx) error {
		bulk := make([]*ent.NotifChannelCreate, len(in))
		for i, info := range in {
			bulk[i] = tx.NotifChannel.Create()
			bulk[i], err = CreateSet(bulk[i], info)
			if err != nil {
				return err
			}
		}
		rows, err = tx.NotifChannel.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func UpdateSet(u *ent.NotifChannelUpdateOne, in *npool.ChannelReq) (*ent.NotifChannelUpdateOne, error) {
	return u, nil
}

func Update(ctx context.Context, in *npool.ChannelReq) (*ent.NotifChannel, error) {
	var info *ent.NotifChannel
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
		info, err = tx.NotifChannel.Query().Where(notifchannel.ID(uuid.MustParse(in.GetID()))).ForUpdate().Only(_ctx)
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

func Row(ctx context.Context, id uuid.UUID) (*ent.NotifChannel, error) {
	var info *ent.NotifChannel
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
		info, err = cli.NotifChannel.Query().Where(notifchannel.ID(id)).Only(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}

func SetQueryConds(conds *npool.Conds, cli *ent.Client) (*ent.NotifChannelQuery, error) { //nolint
	stm := cli.NotifChannel.Query()
	if conds == nil {
		return stm, nil
	}
	if conds.ID != nil {
		switch conds.GetID().GetOp() {
		case cruder.EQ:
			stm.Where(notifchannel.ID(uuid.MustParse(conds.GetID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid notifchannel field")
		}
	}

	if conds.AppID != nil {
		switch conds.GetAppID().GetOp() {
		case cruder.EQ:
			stm.Where(notifchannel.AppID(uuid.MustParse(conds.GetAppID().GetValue())))
		default:
			return nil, fmt.Errorf("invalid notifchannel field")
		}
	}

	if conds.EventType != nil {
		switch conds.GetEventType().GetOp() {
		case cruder.EQ:
			stm.Where(notifchannel.EventType(basetypes.UsedFor(conds.GetEventType().GetValue()).String()))
		default:
			return nil, fmt.Errorf("invalid notifchannel field")
		}
	}

	if len(conds.GetEventTypes().GetValue()) > 0 {
		switch conds.GetEventTypes().GetOp() {
		case cruder.IN:
			types := []string{}
			for _, typ := range conds.GetEventTypes().GetValue() {
				types = append(types, basetypes.UsedFor(typ).String())
			}
			stm.Where(notifchannel.EventTypeIn(types...))
		default:
			return nil, fmt.Errorf("invalid notifchannel field")
		}
	}

	if conds.Channel != nil {
		switch conds.GetChannel().GetOp() {
		case cruder.EQ:
			stm.Where(notifchannel.Channel(channel.NotifChannel(conds.GetChannel().GetValue()).String()))
		default:
			return nil, fmt.Errorf("invalid notifchannel field")
		}
	}

	if len(conds.GetChannels().GetValue()) > 0 {
		switch conds.GetChannels().GetOp() {
		case cruder.IN:
			chans := []string{}
			for _, ch := range conds.GetChannels().GetValue() {
				chans = append(chans, channel.NotifChannel(ch).String())
			}
			stm.Where(notifchannel.ChannelIn(chans...))
		default:
			return nil, fmt.Errorf("invalid notifchannel field")
		}
	}

	return stm, nil
}

func Rows(ctx context.Context, conds *npool.Conds, offset, limit int) ([]*ent.NotifChannel, int, error) {
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

	rows := []*ent.NotifChannel{}
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
			Order(ent.Desc(notifchannel.FieldCreatedAt)).
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

func RowOnly(ctx context.Context, conds *npool.Conds) (*ent.NotifChannel, error) {
	var info *ent.NotifChannel
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
		exist, err = cli.NotifChannel.Query().Where(notifchannel.ID(id)).Exist(_ctx)
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

func Delete(ctx context.Context, id uuid.UUID) (*ent.NotifChannel, error) {
	var info *ent.NotifChannel
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
		info, err = cli.NotifChannel.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, err
	}

	return info, nil
}
