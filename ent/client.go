// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"myetc.lantron.ltd/m/ent/migrate"

	"myetc.lantron.ltd/m/ent/etcrecord"
	"myetc.lantron.ltd/m/ent/etcsyncrecord"
	"myetc.lantron.ltd/m/ent/etcuser"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// ETCRecord is the client for interacting with the ETCRecord builders.
	ETCRecord *ETCRecordClient
	// ETCSyncRecord is the client for interacting with the ETCSyncRecord builders.
	ETCSyncRecord *ETCSyncRecordClient
	// ETCUser is the client for interacting with the ETCUser builders.
	ETCUser *ETCUserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.ETCRecord = NewETCRecordClient(c.config)
	c.ETCSyncRecord = NewETCSyncRecordClient(c.config)
	c.ETCUser = NewETCUserClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:           ctx,
		config:        cfg,
		ETCRecord:     NewETCRecordClient(cfg),
		ETCSyncRecord: NewETCSyncRecordClient(cfg),
		ETCUser:       NewETCUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:           ctx,
		config:        cfg,
		ETCRecord:     NewETCRecordClient(cfg),
		ETCSyncRecord: NewETCSyncRecordClient(cfg),
		ETCUser:       NewETCUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		ETCRecord.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.ETCRecord.Use(hooks...)
	c.ETCSyncRecord.Use(hooks...)
	c.ETCUser.Use(hooks...)
}

// ETCRecordClient is a client for the ETCRecord schema.
type ETCRecordClient struct {
	config
}

// NewETCRecordClient returns a client for the ETCRecord from the given config.
func NewETCRecordClient(c config) *ETCRecordClient {
	return &ETCRecordClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `etcrecord.Hooks(f(g(h())))`.
func (c *ETCRecordClient) Use(hooks ...Hook) {
	c.hooks.ETCRecord = append(c.hooks.ETCRecord, hooks...)
}

// Create returns a builder for creating a ETCRecord entity.
func (c *ETCRecordClient) Create() *ETCRecordCreate {
	mutation := newETCRecordMutation(c.config, OpCreate)
	return &ETCRecordCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ETCRecord entities.
func (c *ETCRecordClient) CreateBulk(builders ...*ETCRecordCreate) *ETCRecordCreateBulk {
	return &ETCRecordCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ETCRecord.
func (c *ETCRecordClient) Update() *ETCRecordUpdate {
	mutation := newETCRecordMutation(c.config, OpUpdate)
	return &ETCRecordUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ETCRecordClient) UpdateOne(er *ETCRecord) *ETCRecordUpdateOne {
	mutation := newETCRecordMutation(c.config, OpUpdateOne, withETCRecord(er))
	return &ETCRecordUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ETCRecordClient) UpdateOneID(id int64) *ETCRecordUpdateOne {
	mutation := newETCRecordMutation(c.config, OpUpdateOne, withETCRecordID(id))
	return &ETCRecordUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ETCRecord.
func (c *ETCRecordClient) Delete() *ETCRecordDelete {
	mutation := newETCRecordMutation(c.config, OpDelete)
	return &ETCRecordDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ETCRecordClient) DeleteOne(er *ETCRecord) *ETCRecordDeleteOne {
	return c.DeleteOneID(er.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ETCRecordClient) DeleteOneID(id int64) *ETCRecordDeleteOne {
	builder := c.Delete().Where(etcrecord.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ETCRecordDeleteOne{builder}
}

// Query returns a query builder for ETCRecord.
func (c *ETCRecordClient) Query() *ETCRecordQuery {
	return &ETCRecordQuery{
		config: c.config,
	}
}

// Get returns a ETCRecord entity by its id.
func (c *ETCRecordClient) Get(ctx context.Context, id int64) (*ETCRecord, error) {
	return c.Query().Where(etcrecord.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ETCRecordClient) GetX(ctx context.Context, id int64) *ETCRecord {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ETCRecordClient) Hooks() []Hook {
	return c.hooks.ETCRecord
}

// ETCSyncRecordClient is a client for the ETCSyncRecord schema.
type ETCSyncRecordClient struct {
	config
}

// NewETCSyncRecordClient returns a client for the ETCSyncRecord from the given config.
func NewETCSyncRecordClient(c config) *ETCSyncRecordClient {
	return &ETCSyncRecordClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `etcsyncrecord.Hooks(f(g(h())))`.
func (c *ETCSyncRecordClient) Use(hooks ...Hook) {
	c.hooks.ETCSyncRecord = append(c.hooks.ETCSyncRecord, hooks...)
}

// Create returns a builder for creating a ETCSyncRecord entity.
func (c *ETCSyncRecordClient) Create() *ETCSyncRecordCreate {
	mutation := newETCSyncRecordMutation(c.config, OpCreate)
	return &ETCSyncRecordCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ETCSyncRecord entities.
func (c *ETCSyncRecordClient) CreateBulk(builders ...*ETCSyncRecordCreate) *ETCSyncRecordCreateBulk {
	return &ETCSyncRecordCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ETCSyncRecord.
func (c *ETCSyncRecordClient) Update() *ETCSyncRecordUpdate {
	mutation := newETCSyncRecordMutation(c.config, OpUpdate)
	return &ETCSyncRecordUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ETCSyncRecordClient) UpdateOne(esr *ETCSyncRecord) *ETCSyncRecordUpdateOne {
	mutation := newETCSyncRecordMutation(c.config, OpUpdateOne, withETCSyncRecord(esr))
	return &ETCSyncRecordUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ETCSyncRecordClient) UpdateOneID(id int64) *ETCSyncRecordUpdateOne {
	mutation := newETCSyncRecordMutation(c.config, OpUpdateOne, withETCSyncRecordID(id))
	return &ETCSyncRecordUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ETCSyncRecord.
func (c *ETCSyncRecordClient) Delete() *ETCSyncRecordDelete {
	mutation := newETCSyncRecordMutation(c.config, OpDelete)
	return &ETCSyncRecordDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ETCSyncRecordClient) DeleteOne(esr *ETCSyncRecord) *ETCSyncRecordDeleteOne {
	return c.DeleteOneID(esr.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ETCSyncRecordClient) DeleteOneID(id int64) *ETCSyncRecordDeleteOne {
	builder := c.Delete().Where(etcsyncrecord.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ETCSyncRecordDeleteOne{builder}
}

// Query returns a query builder for ETCSyncRecord.
func (c *ETCSyncRecordClient) Query() *ETCSyncRecordQuery {
	return &ETCSyncRecordQuery{
		config: c.config,
	}
}

// Get returns a ETCSyncRecord entity by its id.
func (c *ETCSyncRecordClient) Get(ctx context.Context, id int64) (*ETCSyncRecord, error) {
	return c.Query().Where(etcsyncrecord.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ETCSyncRecordClient) GetX(ctx context.Context, id int64) *ETCSyncRecord {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ETCSyncRecordClient) Hooks() []Hook {
	return c.hooks.ETCSyncRecord
}

// ETCUserClient is a client for the ETCUser schema.
type ETCUserClient struct {
	config
}

// NewETCUserClient returns a client for the ETCUser from the given config.
func NewETCUserClient(c config) *ETCUserClient {
	return &ETCUserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `etcuser.Hooks(f(g(h())))`.
func (c *ETCUserClient) Use(hooks ...Hook) {
	c.hooks.ETCUser = append(c.hooks.ETCUser, hooks...)
}

// Create returns a builder for creating a ETCUser entity.
func (c *ETCUserClient) Create() *ETCUserCreate {
	mutation := newETCUserMutation(c.config, OpCreate)
	return &ETCUserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ETCUser entities.
func (c *ETCUserClient) CreateBulk(builders ...*ETCUserCreate) *ETCUserCreateBulk {
	return &ETCUserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ETCUser.
func (c *ETCUserClient) Update() *ETCUserUpdate {
	mutation := newETCUserMutation(c.config, OpUpdate)
	return &ETCUserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ETCUserClient) UpdateOne(eu *ETCUser) *ETCUserUpdateOne {
	mutation := newETCUserMutation(c.config, OpUpdateOne, withETCUser(eu))
	return &ETCUserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ETCUserClient) UpdateOneID(id int64) *ETCUserUpdateOne {
	mutation := newETCUserMutation(c.config, OpUpdateOne, withETCUserID(id))
	return &ETCUserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ETCUser.
func (c *ETCUserClient) Delete() *ETCUserDelete {
	mutation := newETCUserMutation(c.config, OpDelete)
	return &ETCUserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ETCUserClient) DeleteOne(eu *ETCUser) *ETCUserDeleteOne {
	return c.DeleteOneID(eu.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ETCUserClient) DeleteOneID(id int64) *ETCUserDeleteOne {
	builder := c.Delete().Where(etcuser.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ETCUserDeleteOne{builder}
}

// Query returns a query builder for ETCUser.
func (c *ETCUserClient) Query() *ETCUserQuery {
	return &ETCUserQuery{
		config: c.config,
	}
}

// Get returns a ETCUser entity by its id.
func (c *ETCUserClient) Get(ctx context.Context, id int64) (*ETCUser, error) {
	return c.Query().Where(etcuser.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ETCUserClient) GetX(ctx context.Context, id int64) *ETCUser {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ETCUserClient) Hooks() []Hook {
	return c.hooks.ETCUser
}
