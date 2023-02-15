// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/notif-manager/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/notif-manager/pkg/db/ent/smstemplate"
	"github.com/google/uuid"
)

// SMSTemplateUpdate is the builder for updating SMSTemplate entities.
type SMSTemplateUpdate struct {
	config
	hooks     []Hook
	mutation  *SMSTemplateMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the SMSTemplateUpdate builder.
func (stu *SMSTemplateUpdate) Where(ps ...predicate.SMSTemplate) *SMSTemplateUpdate {
	stu.mutation.Where(ps...)
	return stu
}

// SetCreatedAt sets the "created_at" field.
func (stu *SMSTemplateUpdate) SetCreatedAt(u uint32) *SMSTemplateUpdate {
	stu.mutation.ResetCreatedAt()
	stu.mutation.SetCreatedAt(u)
	return stu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (stu *SMSTemplateUpdate) SetNillableCreatedAt(u *uint32) *SMSTemplateUpdate {
	if u != nil {
		stu.SetCreatedAt(*u)
	}
	return stu
}

// AddCreatedAt adds u to the "created_at" field.
func (stu *SMSTemplateUpdate) AddCreatedAt(u int32) *SMSTemplateUpdate {
	stu.mutation.AddCreatedAt(u)
	return stu
}

// SetUpdatedAt sets the "updated_at" field.
func (stu *SMSTemplateUpdate) SetUpdatedAt(u uint32) *SMSTemplateUpdate {
	stu.mutation.ResetUpdatedAt()
	stu.mutation.SetUpdatedAt(u)
	return stu
}

// AddUpdatedAt adds u to the "updated_at" field.
func (stu *SMSTemplateUpdate) AddUpdatedAt(u int32) *SMSTemplateUpdate {
	stu.mutation.AddUpdatedAt(u)
	return stu
}

// SetDeletedAt sets the "deleted_at" field.
func (stu *SMSTemplateUpdate) SetDeletedAt(u uint32) *SMSTemplateUpdate {
	stu.mutation.ResetDeletedAt()
	stu.mutation.SetDeletedAt(u)
	return stu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (stu *SMSTemplateUpdate) SetNillableDeletedAt(u *uint32) *SMSTemplateUpdate {
	if u != nil {
		stu.SetDeletedAt(*u)
	}
	return stu
}

// AddDeletedAt adds u to the "deleted_at" field.
func (stu *SMSTemplateUpdate) AddDeletedAt(u int32) *SMSTemplateUpdate {
	stu.mutation.AddDeletedAt(u)
	return stu
}

// SetAppID sets the "app_id" field.
func (stu *SMSTemplateUpdate) SetAppID(u uuid.UUID) *SMSTemplateUpdate {
	stu.mutation.SetAppID(u)
	return stu
}

// SetLangID sets the "lang_id" field.
func (stu *SMSTemplateUpdate) SetLangID(u uuid.UUID) *SMSTemplateUpdate {
	stu.mutation.SetLangID(u)
	return stu
}

// SetUsedFor sets the "used_for" field.
func (stu *SMSTemplateUpdate) SetUsedFor(s string) *SMSTemplateUpdate {
	stu.mutation.SetUsedFor(s)
	return stu
}

// SetNillableUsedFor sets the "used_for" field if the given value is not nil.
func (stu *SMSTemplateUpdate) SetNillableUsedFor(s *string) *SMSTemplateUpdate {
	if s != nil {
		stu.SetUsedFor(*s)
	}
	return stu
}

// ClearUsedFor clears the value of the "used_for" field.
func (stu *SMSTemplateUpdate) ClearUsedFor() *SMSTemplateUpdate {
	stu.mutation.ClearUsedFor()
	return stu
}

// SetSubject sets the "subject" field.
func (stu *SMSTemplateUpdate) SetSubject(s string) *SMSTemplateUpdate {
	stu.mutation.SetSubject(s)
	return stu
}

// SetNillableSubject sets the "subject" field if the given value is not nil.
func (stu *SMSTemplateUpdate) SetNillableSubject(s *string) *SMSTemplateUpdate {
	if s != nil {
		stu.SetSubject(*s)
	}
	return stu
}

// ClearSubject clears the value of the "subject" field.
func (stu *SMSTemplateUpdate) ClearSubject() *SMSTemplateUpdate {
	stu.mutation.ClearSubject()
	return stu
}

// SetMessage sets the "message" field.
func (stu *SMSTemplateUpdate) SetMessage(s string) *SMSTemplateUpdate {
	stu.mutation.SetMessage(s)
	return stu
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (stu *SMSTemplateUpdate) SetNillableMessage(s *string) *SMSTemplateUpdate {
	if s != nil {
		stu.SetMessage(*s)
	}
	return stu
}

// ClearMessage clears the value of the "message" field.
func (stu *SMSTemplateUpdate) ClearMessage() *SMSTemplateUpdate {
	stu.mutation.ClearMessage()
	return stu
}

// Mutation returns the SMSTemplateMutation object of the builder.
func (stu *SMSTemplateUpdate) Mutation() *SMSTemplateMutation {
	return stu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (stu *SMSTemplateUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if err := stu.defaults(); err != nil {
		return 0, err
	}
	if len(stu.hooks) == 0 {
		affected, err = stu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SMSTemplateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			stu.mutation = mutation
			affected, err = stu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(stu.hooks) - 1; i >= 0; i-- {
			if stu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = stu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, stu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (stu *SMSTemplateUpdate) SaveX(ctx context.Context) int {
	affected, err := stu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (stu *SMSTemplateUpdate) Exec(ctx context.Context) error {
	_, err := stu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (stu *SMSTemplateUpdate) ExecX(ctx context.Context) {
	if err := stu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (stu *SMSTemplateUpdate) defaults() error {
	if _, ok := stu.mutation.UpdatedAt(); !ok {
		if smstemplate.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized smstemplate.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := smstemplate.UpdateDefaultUpdatedAt()
		stu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (stu *SMSTemplateUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SMSTemplateUpdate {
	stu.modifiers = append(stu.modifiers, modifiers...)
	return stu
}

func (stu *SMSTemplateUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   smstemplate.Table,
			Columns: smstemplate.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: smstemplate.FieldID,
			},
		},
	}
	if ps := stu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := stu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: smstemplate.FieldCreatedAt,
		})
	}
	if value, ok := stu.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: smstemplate.FieldCreatedAt,
		})
	}
	if value, ok := stu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: smstemplate.FieldUpdatedAt,
		})
	}
	if value, ok := stu.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: smstemplate.FieldUpdatedAt,
		})
	}
	if value, ok := stu.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: smstemplate.FieldDeletedAt,
		})
	}
	if value, ok := stu.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: smstemplate.FieldDeletedAt,
		})
	}
	if value, ok := stu.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: smstemplate.FieldAppID,
		})
	}
	if value, ok := stu.mutation.LangID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: smstemplate.FieldLangID,
		})
	}
	if value, ok := stu.mutation.UsedFor(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: smstemplate.FieldUsedFor,
		})
	}
	if stu.mutation.UsedForCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: smstemplate.FieldUsedFor,
		})
	}
	if value, ok := stu.mutation.Subject(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: smstemplate.FieldSubject,
		})
	}
	if stu.mutation.SubjectCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: smstemplate.FieldSubject,
		})
	}
	if value, ok := stu.mutation.Message(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: smstemplate.FieldMessage,
		})
	}
	if stu.mutation.MessageCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: smstemplate.FieldMessage,
		})
	}
	_spec.Modifiers = stu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, stu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{smstemplate.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// SMSTemplateUpdateOne is the builder for updating a single SMSTemplate entity.
type SMSTemplateUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *SMSTemplateMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetCreatedAt sets the "created_at" field.
func (stuo *SMSTemplateUpdateOne) SetCreatedAt(u uint32) *SMSTemplateUpdateOne {
	stuo.mutation.ResetCreatedAt()
	stuo.mutation.SetCreatedAt(u)
	return stuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (stuo *SMSTemplateUpdateOne) SetNillableCreatedAt(u *uint32) *SMSTemplateUpdateOne {
	if u != nil {
		stuo.SetCreatedAt(*u)
	}
	return stuo
}

// AddCreatedAt adds u to the "created_at" field.
func (stuo *SMSTemplateUpdateOne) AddCreatedAt(u int32) *SMSTemplateUpdateOne {
	stuo.mutation.AddCreatedAt(u)
	return stuo
}

// SetUpdatedAt sets the "updated_at" field.
func (stuo *SMSTemplateUpdateOne) SetUpdatedAt(u uint32) *SMSTemplateUpdateOne {
	stuo.mutation.ResetUpdatedAt()
	stuo.mutation.SetUpdatedAt(u)
	return stuo
}

// AddUpdatedAt adds u to the "updated_at" field.
func (stuo *SMSTemplateUpdateOne) AddUpdatedAt(u int32) *SMSTemplateUpdateOne {
	stuo.mutation.AddUpdatedAt(u)
	return stuo
}

// SetDeletedAt sets the "deleted_at" field.
func (stuo *SMSTemplateUpdateOne) SetDeletedAt(u uint32) *SMSTemplateUpdateOne {
	stuo.mutation.ResetDeletedAt()
	stuo.mutation.SetDeletedAt(u)
	return stuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (stuo *SMSTemplateUpdateOne) SetNillableDeletedAt(u *uint32) *SMSTemplateUpdateOne {
	if u != nil {
		stuo.SetDeletedAt(*u)
	}
	return stuo
}

// AddDeletedAt adds u to the "deleted_at" field.
func (stuo *SMSTemplateUpdateOne) AddDeletedAt(u int32) *SMSTemplateUpdateOne {
	stuo.mutation.AddDeletedAt(u)
	return stuo
}

// SetAppID sets the "app_id" field.
func (stuo *SMSTemplateUpdateOne) SetAppID(u uuid.UUID) *SMSTemplateUpdateOne {
	stuo.mutation.SetAppID(u)
	return stuo
}

// SetLangID sets the "lang_id" field.
func (stuo *SMSTemplateUpdateOne) SetLangID(u uuid.UUID) *SMSTemplateUpdateOne {
	stuo.mutation.SetLangID(u)
	return stuo
}

// SetUsedFor sets the "used_for" field.
func (stuo *SMSTemplateUpdateOne) SetUsedFor(s string) *SMSTemplateUpdateOne {
	stuo.mutation.SetUsedFor(s)
	return stuo
}

// SetNillableUsedFor sets the "used_for" field if the given value is not nil.
func (stuo *SMSTemplateUpdateOne) SetNillableUsedFor(s *string) *SMSTemplateUpdateOne {
	if s != nil {
		stuo.SetUsedFor(*s)
	}
	return stuo
}

// ClearUsedFor clears the value of the "used_for" field.
func (stuo *SMSTemplateUpdateOne) ClearUsedFor() *SMSTemplateUpdateOne {
	stuo.mutation.ClearUsedFor()
	return stuo
}

// SetSubject sets the "subject" field.
func (stuo *SMSTemplateUpdateOne) SetSubject(s string) *SMSTemplateUpdateOne {
	stuo.mutation.SetSubject(s)
	return stuo
}

// SetNillableSubject sets the "subject" field if the given value is not nil.
func (stuo *SMSTemplateUpdateOne) SetNillableSubject(s *string) *SMSTemplateUpdateOne {
	if s != nil {
		stuo.SetSubject(*s)
	}
	return stuo
}

// ClearSubject clears the value of the "subject" field.
func (stuo *SMSTemplateUpdateOne) ClearSubject() *SMSTemplateUpdateOne {
	stuo.mutation.ClearSubject()
	return stuo
}

// SetMessage sets the "message" field.
func (stuo *SMSTemplateUpdateOne) SetMessage(s string) *SMSTemplateUpdateOne {
	stuo.mutation.SetMessage(s)
	return stuo
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (stuo *SMSTemplateUpdateOne) SetNillableMessage(s *string) *SMSTemplateUpdateOne {
	if s != nil {
		stuo.SetMessage(*s)
	}
	return stuo
}

// ClearMessage clears the value of the "message" field.
func (stuo *SMSTemplateUpdateOne) ClearMessage() *SMSTemplateUpdateOne {
	stuo.mutation.ClearMessage()
	return stuo
}

// Mutation returns the SMSTemplateMutation object of the builder.
func (stuo *SMSTemplateUpdateOne) Mutation() *SMSTemplateMutation {
	return stuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (stuo *SMSTemplateUpdateOne) Select(field string, fields ...string) *SMSTemplateUpdateOne {
	stuo.fields = append([]string{field}, fields...)
	return stuo
}

// Save executes the query and returns the updated SMSTemplate entity.
func (stuo *SMSTemplateUpdateOne) Save(ctx context.Context) (*SMSTemplate, error) {
	var (
		err  error
		node *SMSTemplate
	)
	if err := stuo.defaults(); err != nil {
		return nil, err
	}
	if len(stuo.hooks) == 0 {
		node, err = stuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SMSTemplateMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			stuo.mutation = mutation
			node, err = stuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(stuo.hooks) - 1; i >= 0; i-- {
			if stuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = stuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, stuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*SMSTemplate)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from SMSTemplateMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (stuo *SMSTemplateUpdateOne) SaveX(ctx context.Context) *SMSTemplate {
	node, err := stuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (stuo *SMSTemplateUpdateOne) Exec(ctx context.Context) error {
	_, err := stuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (stuo *SMSTemplateUpdateOne) ExecX(ctx context.Context) {
	if err := stuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (stuo *SMSTemplateUpdateOne) defaults() error {
	if _, ok := stuo.mutation.UpdatedAt(); !ok {
		if smstemplate.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized smstemplate.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := smstemplate.UpdateDefaultUpdatedAt()
		stuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (stuo *SMSTemplateUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *SMSTemplateUpdateOne {
	stuo.modifiers = append(stuo.modifiers, modifiers...)
	return stuo
}

func (stuo *SMSTemplateUpdateOne) sqlSave(ctx context.Context) (_node *SMSTemplate, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   smstemplate.Table,
			Columns: smstemplate.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: smstemplate.FieldID,
			},
		},
	}
	id, ok := stuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SMSTemplate.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := stuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, smstemplate.FieldID)
		for _, f := range fields {
			if !smstemplate.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != smstemplate.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := stuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := stuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: smstemplate.FieldCreatedAt,
		})
	}
	if value, ok := stuo.mutation.AddedCreatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: smstemplate.FieldCreatedAt,
		})
	}
	if value, ok := stuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: smstemplate.FieldUpdatedAt,
		})
	}
	if value, ok := stuo.mutation.AddedUpdatedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: smstemplate.FieldUpdatedAt,
		})
	}
	if value, ok := stuo.mutation.DeletedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: smstemplate.FieldDeletedAt,
		})
	}
	if value, ok := stuo.mutation.AddedDeletedAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: smstemplate.FieldDeletedAt,
		})
	}
	if value, ok := stuo.mutation.AppID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: smstemplate.FieldAppID,
		})
	}
	if value, ok := stuo.mutation.LangID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: smstemplate.FieldLangID,
		})
	}
	if value, ok := stuo.mutation.UsedFor(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: smstemplate.FieldUsedFor,
		})
	}
	if stuo.mutation.UsedForCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: smstemplate.FieldUsedFor,
		})
	}
	if value, ok := stuo.mutation.Subject(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: smstemplate.FieldSubject,
		})
	}
	if stuo.mutation.SubjectCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: smstemplate.FieldSubject,
		})
	}
	if value, ok := stuo.mutation.Message(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: smstemplate.FieldMessage,
		})
	}
	if stuo.mutation.MessageCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: smstemplate.FieldMessage,
		})
	}
	_spec.Modifiers = stuo.modifiers
	_node = &SMSTemplate{config: stuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, stuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{smstemplate.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
