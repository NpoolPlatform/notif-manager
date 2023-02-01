package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/message/npool/notif/mgr/v1/channel"
	"github.com/NpoolPlatform/notif-manager/pkg/db/mixin"
	"github.com/google/uuid"
)

// SendAnnouncement holds the schema definition for the SendAnnouncement entity.
type SendAnnouncement struct {
	ent.Schema
}

func (SendAnnouncement) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the SendAnnouncement.
func (SendAnnouncement) Fields() []ent.Field {
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
			UUID("user_id", uuid.UUID{}).
			Optional().
			Default(uuid.New),
		field.
			UUID("announcement_id", uuid.UUID{}).
			Optional().
			Default(uuid.New),
		field.
			String("channel").
			Optional().
			Default(channel.NotifChannel_DefaultChannel.String()),
	}
}

// Edges of the SendAnnouncement.
func (SendAnnouncement) Edges() []ent.Edge {
	return nil
}
