package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/notif-manager/pkg/db/mixin"
	"github.com/google/uuid"
)

// Announcement holds the schema definition for the Announcement entity.
type Announcement struct {
	ent.Schema
}

func (Announcement) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the Announcement.
func (Announcement) Fields() []ent.Field {
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
			String("title").
			Optional().
			Default(""),
		field.
			String("content").
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
	}
}

// Edges of the Announcement.
func (Announcement) Edges() []ent.Edge {
	return nil
}
