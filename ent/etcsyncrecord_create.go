// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"myetc.lantron.ltd/m/ent/etcsyncrecord"
)

// ETCSyncRecordCreate is the builder for creating a ETCSyncRecord entity.
type ETCSyncRecordCreate struct {
	config
	mutation *ETCSyncRecordMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetUsername sets the "username" field.
func (esrc *ETCSyncRecordCreate) SetUsername(s string) *ETCSyncRecordCreate {
	esrc.mutation.SetUsername(s)
	return esrc
}

// SetSyncData sets the "sync_data" field.
func (esrc *ETCSyncRecordCreate) SetSyncData(s string) *ETCSyncRecordCreate {
	esrc.mutation.SetSyncData(s)
	return esrc
}

// SetCreatedAt sets the "created_at" field.
func (esrc *ETCSyncRecordCreate) SetCreatedAt(t time.Time) *ETCSyncRecordCreate {
	esrc.mutation.SetCreatedAt(t)
	return esrc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (esrc *ETCSyncRecordCreate) SetNillableCreatedAt(t *time.Time) *ETCSyncRecordCreate {
	if t != nil {
		esrc.SetCreatedAt(*t)
	}
	return esrc
}

// SetUpdatedAt sets the "updated_at" field.
func (esrc *ETCSyncRecordCreate) SetUpdatedAt(t time.Time) *ETCSyncRecordCreate {
	esrc.mutation.SetUpdatedAt(t)
	return esrc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (esrc *ETCSyncRecordCreate) SetNillableUpdatedAt(t *time.Time) *ETCSyncRecordCreate {
	if t != nil {
		esrc.SetUpdatedAt(*t)
	}
	return esrc
}

// SetID sets the "id" field.
func (esrc *ETCSyncRecordCreate) SetID(i int64) *ETCSyncRecordCreate {
	esrc.mutation.SetID(i)
	return esrc
}

// Mutation returns the ETCSyncRecordMutation object of the builder.
func (esrc *ETCSyncRecordCreate) Mutation() *ETCSyncRecordMutation {
	return esrc.mutation
}

// Save creates the ETCSyncRecord in the database.
func (esrc *ETCSyncRecordCreate) Save(ctx context.Context) (*ETCSyncRecord, error) {
	var (
		err  error
		node *ETCSyncRecord
	)
	esrc.defaults()
	if len(esrc.hooks) == 0 {
		if err = esrc.check(); err != nil {
			return nil, err
		}
		node, err = esrc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ETCSyncRecordMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = esrc.check(); err != nil {
				return nil, err
			}
			esrc.mutation = mutation
			if node, err = esrc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(esrc.hooks) - 1; i >= 0; i-- {
			if esrc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = esrc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, esrc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*ETCSyncRecord)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ETCSyncRecordMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (esrc *ETCSyncRecordCreate) SaveX(ctx context.Context) *ETCSyncRecord {
	v, err := esrc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (esrc *ETCSyncRecordCreate) Exec(ctx context.Context) error {
	_, err := esrc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (esrc *ETCSyncRecordCreate) ExecX(ctx context.Context) {
	if err := esrc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (esrc *ETCSyncRecordCreate) defaults() {
	if _, ok := esrc.mutation.CreatedAt(); !ok {
		v := etcsyncrecord.DefaultCreatedAt
		esrc.mutation.SetCreatedAt(v)
	}
	if _, ok := esrc.mutation.UpdatedAt(); !ok {
		v := etcsyncrecord.DefaultUpdatedAt
		esrc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (esrc *ETCSyncRecordCreate) check() error {
	if _, ok := esrc.mutation.Username(); !ok {
		return &ValidationError{Name: "username", err: errors.New(`ent: missing required field "ETCSyncRecord.username"`)}
	}
	if _, ok := esrc.mutation.SyncData(); !ok {
		return &ValidationError{Name: "sync_data", err: errors.New(`ent: missing required field "ETCSyncRecord.sync_data"`)}
	}
	if _, ok := esrc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "ETCSyncRecord.created_at"`)}
	}
	if _, ok := esrc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "ETCSyncRecord.updated_at"`)}
	}
	return nil
}

func (esrc *ETCSyncRecordCreate) sqlSave(ctx context.Context) (*ETCSyncRecord, error) {
	_node, _spec := esrc.createSpec()
	if err := sqlgraph.CreateNode(ctx, esrc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	return _node, nil
}

func (esrc *ETCSyncRecordCreate) createSpec() (*ETCSyncRecord, *sqlgraph.CreateSpec) {
	var (
		_node = &ETCSyncRecord{config: esrc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: etcsyncrecord.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: etcsyncrecord.FieldID,
			},
		}
	)
	_spec.OnConflict = esrc.conflict
	if id, ok := esrc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := esrc.mutation.Username(); ok {
		_spec.SetField(etcsyncrecord.FieldUsername, field.TypeString, value)
		_node.Username = value
	}
	if value, ok := esrc.mutation.SyncData(); ok {
		_spec.SetField(etcsyncrecord.FieldSyncData, field.TypeString, value)
		_node.SyncData = value
	}
	if value, ok := esrc.mutation.CreatedAt(); ok {
		_spec.SetField(etcsyncrecord.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := esrc.mutation.UpdatedAt(); ok {
		_spec.SetField(etcsyncrecord.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ETCSyncRecord.Create().
//		SetUsername(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ETCSyncRecordUpsert) {
//			SetUsername(v+v).
//		}).
//		Exec(ctx)
func (esrc *ETCSyncRecordCreate) OnConflict(opts ...sql.ConflictOption) *ETCSyncRecordUpsertOne {
	esrc.conflict = opts
	return &ETCSyncRecordUpsertOne{
		create: esrc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ETCSyncRecord.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (esrc *ETCSyncRecordCreate) OnConflictColumns(columns ...string) *ETCSyncRecordUpsertOne {
	esrc.conflict = append(esrc.conflict, sql.ConflictColumns(columns...))
	return &ETCSyncRecordUpsertOne{
		create: esrc,
	}
}

type (
	// ETCSyncRecordUpsertOne is the builder for "upsert"-ing
	//  one ETCSyncRecord node.
	ETCSyncRecordUpsertOne struct {
		create *ETCSyncRecordCreate
	}

	// ETCSyncRecordUpsert is the "OnConflict" setter.
	ETCSyncRecordUpsert struct {
		*sql.UpdateSet
	}
)

// SetUsername sets the "username" field.
func (u *ETCSyncRecordUpsert) SetUsername(v string) *ETCSyncRecordUpsert {
	u.Set(etcsyncrecord.FieldUsername, v)
	return u
}

// UpdateUsername sets the "username" field to the value that was provided on create.
func (u *ETCSyncRecordUpsert) UpdateUsername() *ETCSyncRecordUpsert {
	u.SetExcluded(etcsyncrecord.FieldUsername)
	return u
}

// SetSyncData sets the "sync_data" field.
func (u *ETCSyncRecordUpsert) SetSyncData(v string) *ETCSyncRecordUpsert {
	u.Set(etcsyncrecord.FieldSyncData, v)
	return u
}

// UpdateSyncData sets the "sync_data" field to the value that was provided on create.
func (u *ETCSyncRecordUpsert) UpdateSyncData() *ETCSyncRecordUpsert {
	u.SetExcluded(etcsyncrecord.FieldSyncData)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *ETCSyncRecordUpsert) SetCreatedAt(v time.Time) *ETCSyncRecordUpsert {
	u.Set(etcsyncrecord.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ETCSyncRecordUpsert) UpdateCreatedAt() *ETCSyncRecordUpsert {
	u.SetExcluded(etcsyncrecord.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ETCSyncRecordUpsert) SetUpdatedAt(v time.Time) *ETCSyncRecordUpsert {
	u.Set(etcsyncrecord.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ETCSyncRecordUpsert) UpdateUpdatedAt() *ETCSyncRecordUpsert {
	u.SetExcluded(etcsyncrecord.FieldUpdatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.ETCSyncRecord.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(etcsyncrecord.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ETCSyncRecordUpsertOne) UpdateNewValues() *ETCSyncRecordUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(etcsyncrecord.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ETCSyncRecord.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ETCSyncRecordUpsertOne) Ignore() *ETCSyncRecordUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ETCSyncRecordUpsertOne) DoNothing() *ETCSyncRecordUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ETCSyncRecordCreate.OnConflict
// documentation for more info.
func (u *ETCSyncRecordUpsertOne) Update(set func(*ETCSyncRecordUpsert)) *ETCSyncRecordUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ETCSyncRecordUpsert{UpdateSet: update})
	}))
	return u
}

// SetUsername sets the "username" field.
func (u *ETCSyncRecordUpsertOne) SetUsername(v string) *ETCSyncRecordUpsertOne {
	return u.Update(func(s *ETCSyncRecordUpsert) {
		s.SetUsername(v)
	})
}

// UpdateUsername sets the "username" field to the value that was provided on create.
func (u *ETCSyncRecordUpsertOne) UpdateUsername() *ETCSyncRecordUpsertOne {
	return u.Update(func(s *ETCSyncRecordUpsert) {
		s.UpdateUsername()
	})
}

// SetSyncData sets the "sync_data" field.
func (u *ETCSyncRecordUpsertOne) SetSyncData(v string) *ETCSyncRecordUpsertOne {
	return u.Update(func(s *ETCSyncRecordUpsert) {
		s.SetSyncData(v)
	})
}

// UpdateSyncData sets the "sync_data" field to the value that was provided on create.
func (u *ETCSyncRecordUpsertOne) UpdateSyncData() *ETCSyncRecordUpsertOne {
	return u.Update(func(s *ETCSyncRecordUpsert) {
		s.UpdateSyncData()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *ETCSyncRecordUpsertOne) SetCreatedAt(v time.Time) *ETCSyncRecordUpsertOne {
	return u.Update(func(s *ETCSyncRecordUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ETCSyncRecordUpsertOne) UpdateCreatedAt() *ETCSyncRecordUpsertOne {
	return u.Update(func(s *ETCSyncRecordUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ETCSyncRecordUpsertOne) SetUpdatedAt(v time.Time) *ETCSyncRecordUpsertOne {
	return u.Update(func(s *ETCSyncRecordUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ETCSyncRecordUpsertOne) UpdateUpdatedAt() *ETCSyncRecordUpsertOne {
	return u.Update(func(s *ETCSyncRecordUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *ETCSyncRecordUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ETCSyncRecordCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ETCSyncRecordUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ETCSyncRecordUpsertOne) ID(ctx context.Context) (id int64, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ETCSyncRecordUpsertOne) IDX(ctx context.Context) int64 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ETCSyncRecordCreateBulk is the builder for creating many ETCSyncRecord entities in bulk.
type ETCSyncRecordCreateBulk struct {
	config
	builders []*ETCSyncRecordCreate
	conflict []sql.ConflictOption
}

// Save creates the ETCSyncRecord entities in the database.
func (esrcb *ETCSyncRecordCreateBulk) Save(ctx context.Context) ([]*ETCSyncRecord, error) {
	specs := make([]*sqlgraph.CreateSpec, len(esrcb.builders))
	nodes := make([]*ETCSyncRecord, len(esrcb.builders))
	mutators := make([]Mutator, len(esrcb.builders))
	for i := range esrcb.builders {
		func(i int, root context.Context) {
			builder := esrcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ETCSyncRecordMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, esrcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = esrcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, esrcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, esrcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (esrcb *ETCSyncRecordCreateBulk) SaveX(ctx context.Context) []*ETCSyncRecord {
	v, err := esrcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (esrcb *ETCSyncRecordCreateBulk) Exec(ctx context.Context) error {
	_, err := esrcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (esrcb *ETCSyncRecordCreateBulk) ExecX(ctx context.Context) {
	if err := esrcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ETCSyncRecord.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ETCSyncRecordUpsert) {
//			SetUsername(v+v).
//		}).
//		Exec(ctx)
func (esrcb *ETCSyncRecordCreateBulk) OnConflict(opts ...sql.ConflictOption) *ETCSyncRecordUpsertBulk {
	esrcb.conflict = opts
	return &ETCSyncRecordUpsertBulk{
		create: esrcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ETCSyncRecord.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (esrcb *ETCSyncRecordCreateBulk) OnConflictColumns(columns ...string) *ETCSyncRecordUpsertBulk {
	esrcb.conflict = append(esrcb.conflict, sql.ConflictColumns(columns...))
	return &ETCSyncRecordUpsertBulk{
		create: esrcb,
	}
}

// ETCSyncRecordUpsertBulk is the builder for "upsert"-ing
// a bulk of ETCSyncRecord nodes.
type ETCSyncRecordUpsertBulk struct {
	create *ETCSyncRecordCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.ETCSyncRecord.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(etcsyncrecord.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ETCSyncRecordUpsertBulk) UpdateNewValues() *ETCSyncRecordUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(etcsyncrecord.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ETCSyncRecord.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ETCSyncRecordUpsertBulk) Ignore() *ETCSyncRecordUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ETCSyncRecordUpsertBulk) DoNothing() *ETCSyncRecordUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ETCSyncRecordCreateBulk.OnConflict
// documentation for more info.
func (u *ETCSyncRecordUpsertBulk) Update(set func(*ETCSyncRecordUpsert)) *ETCSyncRecordUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ETCSyncRecordUpsert{UpdateSet: update})
	}))
	return u
}

// SetUsername sets the "username" field.
func (u *ETCSyncRecordUpsertBulk) SetUsername(v string) *ETCSyncRecordUpsertBulk {
	return u.Update(func(s *ETCSyncRecordUpsert) {
		s.SetUsername(v)
	})
}

// UpdateUsername sets the "username" field to the value that was provided on create.
func (u *ETCSyncRecordUpsertBulk) UpdateUsername() *ETCSyncRecordUpsertBulk {
	return u.Update(func(s *ETCSyncRecordUpsert) {
		s.UpdateUsername()
	})
}

// SetSyncData sets the "sync_data" field.
func (u *ETCSyncRecordUpsertBulk) SetSyncData(v string) *ETCSyncRecordUpsertBulk {
	return u.Update(func(s *ETCSyncRecordUpsert) {
		s.SetSyncData(v)
	})
}

// UpdateSyncData sets the "sync_data" field to the value that was provided on create.
func (u *ETCSyncRecordUpsertBulk) UpdateSyncData() *ETCSyncRecordUpsertBulk {
	return u.Update(func(s *ETCSyncRecordUpsert) {
		s.UpdateSyncData()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *ETCSyncRecordUpsertBulk) SetCreatedAt(v time.Time) *ETCSyncRecordUpsertBulk {
	return u.Update(func(s *ETCSyncRecordUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *ETCSyncRecordUpsertBulk) UpdateCreatedAt() *ETCSyncRecordUpsertBulk {
	return u.Update(func(s *ETCSyncRecordUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *ETCSyncRecordUpsertBulk) SetUpdatedAt(v time.Time) *ETCSyncRecordUpsertBulk {
	return u.Update(func(s *ETCSyncRecordUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *ETCSyncRecordUpsertBulk) UpdateUpdatedAt() *ETCSyncRecordUpsertBulk {
	return u.Update(func(s *ETCSyncRecordUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *ETCSyncRecordUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ETCSyncRecordCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ETCSyncRecordCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ETCSyncRecordUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}