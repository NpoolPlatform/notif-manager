package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/notif-manager/pkg/db/mixin"
	"github.com/google/uuid"

	"github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif"
)

// Notif holds the schema definition for the Notif entity.
type Notif struct {
	ent.Schema
}

func (Notif) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the Notif.
func (Notif) Fields() []ent.Field {
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
			Bool("already_read").
			Optional().
			Default(false),
		field.
			UUID("lang_id", uuid.UUID{}).
			Optional().
			Default(uuid.New),
		field.
			String("event_type").
			Optional().
			Default(notif.EventType_DefaultEventType.String()),
		field.
			Bool("use_template").
			Optional().
			Default(false),
		field.
			String("title").
			Optional().
			Default(""),
		field.
			Text("content").
			Optional().
			Default(""),
		field.
			JSON("channels", []string{}).
			Optional().
			Default([]string{}),
		field.
			Bool("email_send").
			Optional().
			Default(true),
		field.
			Text("extra").
			Optional().
			Default(""),
	}
}

// Edges of the Notif.
func (Notif) Edges() []ent.Edge {
	return nil
}
