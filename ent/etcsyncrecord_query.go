// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"myetc.lantron.ltd/m/ent/etcsyncrecord"
	"myetc.lantron.ltd/m/ent/predicate"
)

// ETCSyncRecordQuery is the builder for querying ETCSyncRecord entities.
type ETCSyncRecordQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ETCSyncRecord
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ETCSyncRecordQuery builder.
func (esrq *ETCSyncRecordQuery) Where(ps ...predicate.ETCSyncRecord) *ETCSyncRecordQuery {
	esrq.predicates = append(esrq.predicates, ps...)
	return esrq
}

// Limit adds a limit step to the query.
func (esrq *ETCSyncRecordQuery) Limit(limit int) *ETCSyncRecordQuery {
	esrq.limit = &limit
	return esrq
}

// Offset adds an offset step to the query.
func (esrq *ETCSyncRecordQuery) Offset(offset int) *ETCSyncRecordQuery {
	esrq.offset = &offset
	return esrq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (esrq *ETCSyncRecordQuery) Unique(unique bool) *ETCSyncRecordQuery {
	esrq.unique = &unique
	return esrq
}

// Order adds an order step to the query.
func (esrq *ETCSyncRecordQuery) Order(o ...OrderFunc) *ETCSyncRecordQuery {
	esrq.order = append(esrq.order, o...)
	return esrq
}

// First returns the first ETCSyncRecord entity from the query.
// Returns a *NotFoundError when no ETCSyncRecord was found.
func (esrq *ETCSyncRecordQuery) First(ctx context.Context) (*ETCSyncRecord, error) {
	nodes, err := esrq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{etcsyncrecord.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (esrq *ETCSyncRecordQuery) FirstX(ctx context.Context) *ETCSyncRecord {
	node, err := esrq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ETCSyncRecord ID from the query.
// Returns a *NotFoundError when no ETCSyncRecord ID was found.
func (esrq *ETCSyncRecordQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = esrq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{etcsyncrecord.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (esrq *ETCSyncRecordQuery) FirstIDX(ctx context.Context) int64 {
	id, err := esrq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ETCSyncRecord entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ETCSyncRecord entity is found.
// Returns a *NotFoundError when no ETCSyncRecord entities are found.
func (esrq *ETCSyncRecordQuery) Only(ctx context.Context) (*ETCSyncRecord, error) {
	nodes, err := esrq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{etcsyncrecord.Label}
	default:
		return nil, &NotSingularError{etcsyncrecord.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (esrq *ETCSyncRecordQuery) OnlyX(ctx context.Context) *ETCSyncRecord {
	node, err := esrq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ETCSyncRecord ID in the query.
// Returns a *NotSingularError when more than one ETCSyncRecord ID is found.
// Returns a *NotFoundError when no entities are found.
func (esrq *ETCSyncRecordQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = esrq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{etcsyncrecord.Label}
	default:
		err = &NotSingularError{etcsyncrecord.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (esrq *ETCSyncRecordQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := esrq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ETCSyncRecords.
func (esrq *ETCSyncRecordQuery) All(ctx context.Context) ([]*ETCSyncRecord, error) {
	if err := esrq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return esrq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (esrq *ETCSyncRecordQuery) AllX(ctx context.Context) []*ETCSyncRecord {
	nodes, err := esrq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ETCSyncRecord IDs.
func (esrq *ETCSyncRecordQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := esrq.Select(etcsyncrecord.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (esrq *ETCSyncRecordQuery) IDsX(ctx context.Context) []int64 {
	ids, err := esrq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (esrq *ETCSyncRecordQuery) Count(ctx context.Context) (int, error) {
	if err := esrq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return esrq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (esrq *ETCSyncRecordQuery) CountX(ctx context.Context) int {
	count, err := esrq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (esrq *ETCSyncRecordQuery) Exist(ctx context.Context) (bool, error) {
	if err := esrq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return esrq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (esrq *ETCSyncRecordQuery) ExistX(ctx context.Context) bool {
	exist, err := esrq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ETCSyncRecordQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (esrq *ETCSyncRecordQuery) Clone() *ETCSyncRecordQuery {
	if esrq == nil {
		return nil
	}
	return &ETCSyncRecordQuery{
		config:     esrq.config,
		limit:      esrq.limit,
		offset:     esrq.offset,
		order:      append([]OrderFunc{}, esrq.order...),
		predicates: append([]predicate.ETCSyncRecord{}, esrq.predicates...),
		// clone intermediate query.
		sql:    esrq.sql.Clone(),
		path:   esrq.path,
		unique: esrq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Username string `json:"username,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ETCSyncRecord.Query().
//		GroupBy(etcsyncrecord.FieldUsername).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (esrq *ETCSyncRecordQuery) GroupBy(field string, fields ...string) *ETCSyncRecordGroupBy {
	grbuild := &ETCSyncRecordGroupBy{config: esrq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := esrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return esrq.sqlQuery(ctx), nil
	}
	grbuild.label = etcsyncrecord.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Username string `json:"username,omitempty"`
//	}
//
//	client.ETCSyncRecord.Query().
//		Select(etcsyncrecord.FieldUsername).
//		Scan(ctx, &v)
func (esrq *ETCSyncRecordQuery) Select(fields ...string) *ETCSyncRecordSelect {
	esrq.fields = append(esrq.fields, fields...)
	selbuild := &ETCSyncRecordSelect{ETCSyncRecordQuery: esrq}
	selbuild.label = etcsyncrecord.Label
	selbuild.flds, selbuild.scan = &esrq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a ETCSyncRecordSelect configured with the given aggregations.
func (esrq *ETCSyncRecordQuery) Aggregate(fns ...AggregateFunc) *ETCSyncRecordSelect {
	return esrq.Select().Aggregate(fns...)
}

func (esrq *ETCSyncRecordQuery) prepareQuery(ctx context.Context) error {
	for _, f := range esrq.fields {
		if !etcsyncrecord.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if esrq.path != nil {
		prev, err := esrq.path(ctx)
		if err != nil {
			return err
		}
		esrq.sql = prev
	}
	return nil
}

func (esrq *ETCSyncRecordQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ETCSyncRecord, error) {
	var (
		nodes = []*ETCSyncRecord{}
		_spec = esrq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ETCSyncRecord).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ETCSyncRecord{config: esrq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(esrq.modifiers) > 0 {
		_spec.Modifiers = esrq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, esrq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (esrq *ETCSyncRecordQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := esrq.querySpec()
	if len(esrq.modifiers) > 0 {
		_spec.Modifiers = esrq.modifiers
	}
	_spec.Node.Columns = esrq.fields
	if len(esrq.fields) > 0 {
		_spec.Unique = esrq.unique != nil && *esrq.unique
	}
	return sqlgraph.CountNodes(ctx, esrq.driver, _spec)
}

func (esrq *ETCSyncRecordQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := esrq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (esrq *ETCSyncRecordQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   etcsyncrecord.Table,
			Columns: etcsyncrecord.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: etcsyncrecord.FieldID,
			},
		},
		From:   esrq.sql,
		Unique: true,
	}
	if unique := esrq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := esrq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, etcsyncrecord.FieldID)
		for i := range fields {
			if fields[i] != etcsyncrecord.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := esrq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := esrq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := esrq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := esrq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (esrq *ETCSyncRecordQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(esrq.driver.Dialect())
	t1 := builder.Table(etcsyncrecord.Table)
	columns := esrq.fields
	if len(columns) == 0 {
		columns = etcsyncrecord.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if esrq.sql != nil {
		selector = esrq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if esrq.unique != nil && *esrq.unique {
		selector.Distinct()
	}
	for _, m := range esrq.modifiers {
		m(selector)
	}
	for _, p := range esrq.predicates {
		p(selector)
	}
	for _, p := range esrq.order {
		p(selector)
	}
	if offset := esrq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := esrq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (esrq *ETCSyncRecordQuery) Modify(modifiers ...func(s *sql.Selector)) *ETCSyncRecordSelect {
	esrq.modifiers = append(esrq.modifiers, modifiers...)
	return esrq.Select()
}

// ETCSyncRecordGroupBy is the group-by builder for ETCSyncRecord entities.
type ETCSyncRecordGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (esrgb *ETCSyncRecordGroupBy) Aggregate(fns ...AggregateFunc) *ETCSyncRecordGroupBy {
	esrgb.fns = append(esrgb.fns, fns...)
	return esrgb
}

// Scan applies the group-by query and scans the result into the given value.
func (esrgb *ETCSyncRecordGroupBy) Scan(ctx context.Context, v any) error {
	query, err := esrgb.path(ctx)
	if err != nil {
		return err
	}
	esrgb.sql = query
	return esrgb.sqlScan(ctx, v)
}

func (esrgb *ETCSyncRecordGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range esrgb.fields {
		if !etcsyncrecord.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := esrgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := esrgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (esrgb *ETCSyncRecordGroupBy) sqlQuery() *sql.Selector {
	selector := esrgb.sql.Select()
	aggregation := make([]string, 0, len(esrgb.fns))
	for _, fn := range esrgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(esrgb.fields)+len(esrgb.fns))
		for _, f := range esrgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(esrgb.fields...)...)
}

// ETCSyncRecordSelect is the builder for selecting fields of ETCSyncRecord entities.
type ETCSyncRecordSelect struct {
	*ETCSyncRecordQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (esrs *ETCSyncRecordSelect) Aggregate(fns ...AggregateFunc) *ETCSyncRecordSelect {
	esrs.fns = append(esrs.fns, fns...)
	return esrs
}

// Scan applies the selector query and scans the result into the given value.
func (esrs *ETCSyncRecordSelect) Scan(ctx context.Context, v any) error {
	if err := esrs.prepareQuery(ctx); err != nil {
		return err
	}
	esrs.sql = esrs.ETCSyncRecordQuery.sqlQuery(ctx)
	return esrs.sqlScan(ctx, v)
}

func (esrs *ETCSyncRecordSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(esrs.fns))
	for _, fn := range esrs.fns {
		aggregation = append(aggregation, fn(esrs.sql))
	}
	switch n := len(*esrs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		esrs.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		esrs.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := esrs.sql.Query()
	if err := esrs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (esrs *ETCSyncRecordSelect) Modify(modifiers ...func(s *sql.Selector)) *ETCSyncRecordSelect {
	esrs.modifiers = append(esrs.modifiers, modifiers...)
	return esrs
}
