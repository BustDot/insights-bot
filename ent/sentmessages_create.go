// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/nekomeowww/insights-bot/ent/sentmessages"
)

// SentMessagesCreate is the builder for creating a SentMessages entity.
type SentMessagesCreate struct {
	config
	mutation *SentMessagesMutation
	hooks    []Hook
}

// SetChatID sets the "chat_id" field.
func (smc *SentMessagesCreate) SetChatID(i int64) *SentMessagesCreate {
	smc.mutation.SetChatID(i)
	return smc
}

// SetNillableChatID sets the "chat_id" field if the given value is not nil.
func (smc *SentMessagesCreate) SetNillableChatID(i *int64) *SentMessagesCreate {
	if i != nil {
		smc.SetChatID(*i)
	}
	return smc
}

// SetMessageID sets the "message_id" field.
func (smc *SentMessagesCreate) SetMessageID(i int) *SentMessagesCreate {
	smc.mutation.SetMessageID(i)
	return smc
}

// SetNillableMessageID sets the "message_id" field if the given value is not nil.
func (smc *SentMessagesCreate) SetNillableMessageID(i *int) *SentMessagesCreate {
	if i != nil {
		smc.SetMessageID(*i)
	}
	return smc
}

// SetText sets the "text" field.
func (smc *SentMessagesCreate) SetText(s string) *SentMessagesCreate {
	smc.mutation.SetText(s)
	return smc
}

// SetNillableText sets the "text" field if the given value is not nil.
func (smc *SentMessagesCreate) SetNillableText(s *string) *SentMessagesCreate {
	if s != nil {
		smc.SetText(*s)
	}
	return smc
}

// SetIsPinned sets the "is_pinned" field.
func (smc *SentMessagesCreate) SetIsPinned(b bool) *SentMessagesCreate {
	smc.mutation.SetIsPinned(b)
	return smc
}

// SetNillableIsPinned sets the "is_pinned" field if the given value is not nil.
func (smc *SentMessagesCreate) SetNillableIsPinned(b *bool) *SentMessagesCreate {
	if b != nil {
		smc.SetIsPinned(*b)
	}
	return smc
}

// SetFromPlatform sets the "from_platform" field.
func (smc *SentMessagesCreate) SetFromPlatform(i int) *SentMessagesCreate {
	smc.mutation.SetFromPlatform(i)
	return smc
}

// SetNillableFromPlatform sets the "from_platform" field if the given value is not nil.
func (smc *SentMessagesCreate) SetNillableFromPlatform(i *int) *SentMessagesCreate {
	if i != nil {
		smc.SetFromPlatform(*i)
	}
	return smc
}

// SetMessageType sets the "message_type" field.
func (smc *SentMessagesCreate) SetMessageType(i int) *SentMessagesCreate {
	smc.mutation.SetMessageType(i)
	return smc
}

// SetNillableMessageType sets the "message_type" field if the given value is not nil.
func (smc *SentMessagesCreate) SetNillableMessageType(i *int) *SentMessagesCreate {
	if i != nil {
		smc.SetMessageType(*i)
	}
	return smc
}

// SetCreatedAt sets the "created_at" field.
func (smc *SentMessagesCreate) SetCreatedAt(i int64) *SentMessagesCreate {
	smc.mutation.SetCreatedAt(i)
	return smc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (smc *SentMessagesCreate) SetNillableCreatedAt(i *int64) *SentMessagesCreate {
	if i != nil {
		smc.SetCreatedAt(*i)
	}
	return smc
}

// SetUpdatedAt sets the "updated_at" field.
func (smc *SentMessagesCreate) SetUpdatedAt(i int64) *SentMessagesCreate {
	smc.mutation.SetUpdatedAt(i)
	return smc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (smc *SentMessagesCreate) SetNillableUpdatedAt(i *int64) *SentMessagesCreate {
	if i != nil {
		smc.SetUpdatedAt(*i)
	}
	return smc
}

// SetID sets the "id" field.
func (smc *SentMessagesCreate) SetID(u uuid.UUID) *SentMessagesCreate {
	smc.mutation.SetID(u)
	return smc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (smc *SentMessagesCreate) SetNillableID(u *uuid.UUID) *SentMessagesCreate {
	if u != nil {
		smc.SetID(*u)
	}
	return smc
}

// Mutation returns the SentMessagesMutation object of the builder.
func (smc *SentMessagesCreate) Mutation() *SentMessagesMutation {
	return smc.mutation
}

// Save creates the SentMessages in the database.
func (smc *SentMessagesCreate) Save(ctx context.Context) (*SentMessages, error) {
	smc.defaults()
	return withHooks(ctx, smc.sqlSave, smc.mutation, smc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (smc *SentMessagesCreate) SaveX(ctx context.Context) *SentMessages {
	v, err := smc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (smc *SentMessagesCreate) Exec(ctx context.Context) error {
	_, err := smc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (smc *SentMessagesCreate) ExecX(ctx context.Context) {
	if err := smc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (smc *SentMessagesCreate) defaults() {
	if _, ok := smc.mutation.ChatID(); !ok {
		v := sentmessages.DefaultChatID
		smc.mutation.SetChatID(v)
	}
	if _, ok := smc.mutation.MessageID(); !ok {
		v := sentmessages.DefaultMessageID
		smc.mutation.SetMessageID(v)
	}
	if _, ok := smc.mutation.Text(); !ok {
		v := sentmessages.DefaultText
		smc.mutation.SetText(v)
	}
	if _, ok := smc.mutation.IsPinned(); !ok {
		v := sentmessages.DefaultIsPinned
		smc.mutation.SetIsPinned(v)
	}
	if _, ok := smc.mutation.FromPlatform(); !ok {
		v := sentmessages.DefaultFromPlatform
		smc.mutation.SetFromPlatform(v)
	}
	if _, ok := smc.mutation.MessageType(); !ok {
		v := sentmessages.DefaultMessageType
		smc.mutation.SetMessageType(v)
	}
	if _, ok := smc.mutation.CreatedAt(); !ok {
		v := sentmessages.DefaultCreatedAt()
		smc.mutation.SetCreatedAt(v)
	}
	if _, ok := smc.mutation.UpdatedAt(); !ok {
		v := sentmessages.DefaultUpdatedAt()
		smc.mutation.SetUpdatedAt(v)
	}
	if _, ok := smc.mutation.ID(); !ok {
		v := sentmessages.DefaultID()
		smc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (smc *SentMessagesCreate) check() error {
	if _, ok := smc.mutation.ChatID(); !ok {
		return &ValidationError{Name: "chat_id", err: errors.New(`ent: missing required field "SentMessages.chat_id"`)}
	}
	if _, ok := smc.mutation.MessageID(); !ok {
		return &ValidationError{Name: "message_id", err: errors.New(`ent: missing required field "SentMessages.message_id"`)}
	}
	if _, ok := smc.mutation.Text(); !ok {
		return &ValidationError{Name: "text", err: errors.New(`ent: missing required field "SentMessages.text"`)}
	}
	if _, ok := smc.mutation.IsPinned(); !ok {
		return &ValidationError{Name: "is_pinned", err: errors.New(`ent: missing required field "SentMessages.is_pinned"`)}
	}
	if _, ok := smc.mutation.FromPlatform(); !ok {
		return &ValidationError{Name: "from_platform", err: errors.New(`ent: missing required field "SentMessages.from_platform"`)}
	}
	if _, ok := smc.mutation.MessageType(); !ok {
		return &ValidationError{Name: "message_type", err: errors.New(`ent: missing required field "SentMessages.message_type"`)}
	}
	if _, ok := smc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "SentMessages.created_at"`)}
	}
	if _, ok := smc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "SentMessages.updated_at"`)}
	}
	return nil
}

func (smc *SentMessagesCreate) sqlSave(ctx context.Context) (*SentMessages, error) {
	if err := smc.check(); err != nil {
		return nil, err
	}
	_node, _spec := smc.createSpec()
	if err := sqlgraph.CreateNode(ctx, smc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	smc.mutation.id = &_node.ID
	smc.mutation.done = true
	return _node, nil
}

func (smc *SentMessagesCreate) createSpec() (*SentMessages, *sqlgraph.CreateSpec) {
	var (
		_node = &SentMessages{config: smc.config}
		_spec = sqlgraph.NewCreateSpec(sentmessages.Table, sqlgraph.NewFieldSpec(sentmessages.FieldID, field.TypeUUID))
	)
	_spec.Schema = smc.schemaConfig.SentMessages
	if id, ok := smc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := smc.mutation.ChatID(); ok {
		_spec.SetField(sentmessages.FieldChatID, field.TypeInt64, value)
		_node.ChatID = value
	}
	if value, ok := smc.mutation.MessageID(); ok {
		_spec.SetField(sentmessages.FieldMessageID, field.TypeInt, value)
		_node.MessageID = value
	}
	if value, ok := smc.mutation.Text(); ok {
		_spec.SetField(sentmessages.FieldText, field.TypeString, value)
		_node.Text = value
	}
	if value, ok := smc.mutation.IsPinned(); ok {
		_spec.SetField(sentmessages.FieldIsPinned, field.TypeBool, value)
		_node.IsPinned = value
	}
	if value, ok := smc.mutation.FromPlatform(); ok {
		_spec.SetField(sentmessages.FieldFromPlatform, field.TypeInt, value)
		_node.FromPlatform = value
	}
	if value, ok := smc.mutation.MessageType(); ok {
		_spec.SetField(sentmessages.FieldMessageType, field.TypeInt, value)
		_node.MessageType = value
	}
	if value, ok := smc.mutation.CreatedAt(); ok {
		_spec.SetField(sentmessages.FieldCreatedAt, field.TypeInt64, value)
		_node.CreatedAt = value
	}
	if value, ok := smc.mutation.UpdatedAt(); ok {
		_spec.SetField(sentmessages.FieldUpdatedAt, field.TypeInt64, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// SentMessagesCreateBulk is the builder for creating many SentMessages entities in bulk.
type SentMessagesCreateBulk struct {
	config
	err      error
	builders []*SentMessagesCreate
}

// Save creates the SentMessages entities in the database.
func (smcb *SentMessagesCreateBulk) Save(ctx context.Context) ([]*SentMessages, error) {
	if smcb.err != nil {
		return nil, smcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(smcb.builders))
	nodes := make([]*SentMessages, len(smcb.builders))
	mutators := make([]Mutator, len(smcb.builders))
	for i := range smcb.builders {
		func(i int, root context.Context) {
			builder := smcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SentMessagesMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, smcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, smcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, smcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (smcb *SentMessagesCreateBulk) SaveX(ctx context.Context) []*SentMessages {
	v, err := smcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (smcb *SentMessagesCreateBulk) Exec(ctx context.Context) error {
	_, err := smcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (smcb *SentMessagesCreateBulk) ExecX(ctx context.Context) {
	if err := smcb.Exec(ctx); err != nil {
		panic(err)
	}
}
