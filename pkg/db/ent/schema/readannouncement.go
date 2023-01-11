package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/notif-manager/pkg/db/mixin"
	"github.com/google/uuid"
)

// ReadAnnouncement holds the schema definition for the ReadAnnouncement entity.
type ReadAnnouncement struct {
	ent.Schema
}

func (ReadAnnouncement) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the ReadAnnouncement.
func (ReadAnnouncement) Fields() []ent.Field {
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
	}
}

// Edges of the ReadAnnouncement.
func (ReadAnnouncement) Edges() []ent.Edge {
	return nil
}
