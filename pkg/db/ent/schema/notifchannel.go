package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/NpoolPlatform/notif-manager/pkg/db/mixin"
	"github.com/google/uuid"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	"github.com/NpoolPlatform/message/npool/notif/mgr/v1/channel"
)

// NotifChannel holds the schema definition for the NotifChannel entity.
type NotifChannel struct {
	ent.Schema
}

func (NotifChannel) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the NotifChannel.
func (NotifChannel) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.
			UUID("app_id", uuid.UUID{}).
			Optional().
			Default(uuid.New),
		field.
			String("event_type").
			Optional().
			Default(basetypes.UsedFor_DefaultUsedFor.String()),
		field.
			String("channel").
			Optional().
			Default(channel.NotifChannel_DefaultChannel.String()),
	}
}

// Edges of the NotifChannel.
func (NotifChannel) Edges() []ent.Edge {
	return nil
}
