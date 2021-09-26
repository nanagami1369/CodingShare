// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/nanagami1369/CodingShare/ent/migrate"

	"github.com/nanagami1369/CodingShare/ent/session"
	"github.com/nanagami1369/CodingShare/ent/user"
	"github.com/nanagami1369/CodingShare/ent/video"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Session is the client for interacting with the Session builders.
	Session *SessionClient
	// User is the client for interacting with the User builders.
	User *UserClient
	// Video is the client for interacting with the Video builders.
	Video *VideoClient
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
	c.Session = NewSessionClient(c.config)
	c.User = NewUserClient(c.config)
	c.Video = NewVideoClient(c.config)
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
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:     ctx,
		config:  cfg,
		Session: NewSessionClient(cfg),
		User:    NewUserClient(cfg),
		Video:   NewVideoClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
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
		config:  cfg,
		Session: NewSessionClient(cfg),
		User:    NewUserClient(cfg),
		Video:   NewVideoClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Session.
//		Query().
//		Count(ctx)
//
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
	c.Session.Use(hooks...)
	c.User.Use(hooks...)
	c.Video.Use(hooks...)
}

// SessionClient is a client for the Session schema.
type SessionClient struct {
	config
}

// NewSessionClient returns a client for the Session from the given config.
func NewSessionClient(c config) *SessionClient {
	return &SessionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `session.Hooks(f(g(h())))`.
func (c *SessionClient) Use(hooks ...Hook) {
	c.hooks.Session = append(c.hooks.Session, hooks...)
}

// Create returns a create builder for Session.
func (c *SessionClient) Create() *SessionCreate {
	mutation := newSessionMutation(c.config, OpCreate)
	return &SessionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Session entities.
func (c *SessionClient) CreateBulk(builders ...*SessionCreate) *SessionCreateBulk {
	return &SessionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Session.
func (c *SessionClient) Update() *SessionUpdate {
	mutation := newSessionMutation(c.config, OpUpdate)
	return &SessionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SessionClient) UpdateOne(s *Session) *SessionUpdateOne {
	mutation := newSessionMutation(c.config, OpUpdateOne, withSession(s))
	return &SessionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SessionClient) UpdateOneID(id uuid.UUID) *SessionUpdateOne {
	mutation := newSessionMutation(c.config, OpUpdateOne, withSessionID(id))
	return &SessionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Session.
func (c *SessionClient) Delete() *SessionDelete {
	mutation := newSessionMutation(c.config, OpDelete)
	return &SessionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *SessionClient) DeleteOne(s *Session) *SessionDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *SessionClient) DeleteOneID(id uuid.UUID) *SessionDeleteOne {
	builder := c.Delete().Where(session.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SessionDeleteOne{builder}
}

// Query returns a query builder for Session.
func (c *SessionClient) Query() *SessionQuery {
	return &SessionQuery{
		config: c.config,
	}
}

// Get returns a Session entity by its id.
func (c *SessionClient) Get(ctx context.Context, id uuid.UUID) (*Session, error) {
	return c.Query().Where(session.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SessionClient) GetX(ctx context.Context, id uuid.UUID) *Session {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a Session.
func (c *SessionClient) QueryUser(s *Session) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(session.Table, session.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, session.UserTable, session.UserColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *SessionClient) Hooks() []Hook {
	return c.hooks.Session
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QuerySessions queries the sessions edge of a User.
func (c *UserClient) QuerySessions(u *User) *SessionQuery {
	query := &SessionQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(session.Table, session.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, user.SessionsTable, user.SessionsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// VideoClient is a client for the Video schema.
type VideoClient struct {
	config
}

// NewVideoClient returns a client for the Video from the given config.
func NewVideoClient(c config) *VideoClient {
	return &VideoClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `video.Hooks(f(g(h())))`.
func (c *VideoClient) Use(hooks ...Hook) {
	c.hooks.Video = append(c.hooks.Video, hooks...)
}

// Create returns a create builder for Video.
func (c *VideoClient) Create() *VideoCreate {
	mutation := newVideoMutation(c.config, OpCreate)
	return &VideoCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Video entities.
func (c *VideoClient) CreateBulk(builders ...*VideoCreate) *VideoCreateBulk {
	return &VideoCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Video.
func (c *VideoClient) Update() *VideoUpdate {
	mutation := newVideoMutation(c.config, OpUpdate)
	return &VideoUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *VideoClient) UpdateOne(v *Video) *VideoUpdateOne {
	mutation := newVideoMutation(c.config, OpUpdateOne, withVideo(v))
	return &VideoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *VideoClient) UpdateOneID(id int) *VideoUpdateOne {
	mutation := newVideoMutation(c.config, OpUpdateOne, withVideoID(id))
	return &VideoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Video.
func (c *VideoClient) Delete() *VideoDelete {
	mutation := newVideoMutation(c.config, OpDelete)
	return &VideoDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *VideoClient) DeleteOne(v *Video) *VideoDeleteOne {
	return c.DeleteOneID(v.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *VideoClient) DeleteOneID(id int) *VideoDeleteOne {
	builder := c.Delete().Where(video.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &VideoDeleteOne{builder}
}

// Query returns a query builder for Video.
func (c *VideoClient) Query() *VideoQuery {
	return &VideoQuery{
		config: c.config,
	}
}

// Get returns a Video entity by its id.
func (c *VideoClient) Get(ctx context.Context, id int) (*Video, error) {
	return c.Query().Where(video.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *VideoClient) GetX(ctx context.Context, id int) *Video {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a Video.
func (c *VideoClient) QueryUser(v *Video) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := v.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(video.Table, video.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, video.UserTable, video.UserColumn),
		)
		fromV = sqlgraph.Neighbors(v.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *VideoClient) Hooks() []Hook {
	return c.hooks.Video
}
