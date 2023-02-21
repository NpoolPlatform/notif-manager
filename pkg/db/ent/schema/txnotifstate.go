package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/notif-manager/pkg/db/mixin"
	"github.com/google/uuid"

	basetypes "github.com/NpoolPlatform/message/npool/basetypes/v1"
	npool "github.com/NpoolPlatform/message/npool/notif/mgr/v1/notif/tx"
)

// TxNotifState holds the schema definition for the TxNotifState entity.
type TxNotifState struct {
	ent.Schema
}

func (TxNotifState) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the TxNotifState.
func (TxNotifState) Fields() []ent.Field {
	return []ent.Field{
		field.
			UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.
			UUID("tx_id", uuid.UUID{}).
			Optional().
			Default(uuid.New),
		field.
			String("notif_state").
			Optional().
			Default(npool.TxState_DefaultState.String()),
		field.
			String("tx_type").
			Optional().
			Default(basetypes.TxType_DefaultTxType.String()),
	}
}

// Edges of the TxNotifState.
func (TxNotifState) Edges() []ent.Edge {
	return nil
}
