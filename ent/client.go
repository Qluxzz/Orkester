// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"goreact/ent/migrate"

	"goreact/ent/album"
	"goreact/ent/albumimage"
	"goreact/ent/artist"
	"goreact/ent/likedtrack"
	"goreact/ent/track"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Album is the client for interacting with the Album builders.
	Album *AlbumClient
	// AlbumImage is the client for interacting with the AlbumImage builders.
	AlbumImage *AlbumImageClient
	// Artist is the client for interacting with the Artist builders.
	Artist *ArtistClient
	// LikedTrack is the client for interacting with the LikedTrack builders.
	LikedTrack *LikedTrackClient
	// Track is the client for interacting with the Track builders.
	Track *TrackClient
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
	c.Album = NewAlbumClient(c.config)
	c.AlbumImage = NewAlbumImageClient(c.config)
	c.Artist = NewArtistClient(c.config)
	c.LikedTrack = NewLikedTrackClient(c.config)
	c.Track = NewTrackClient(c.config)
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
		ctx:        ctx,
		config:     cfg,
		Album:      NewAlbumClient(cfg),
		AlbumImage: NewAlbumImageClient(cfg),
		Artist:     NewArtistClient(cfg),
		LikedTrack: NewLikedTrackClient(cfg),
		Track:      NewTrackClient(cfg),
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
		ctx:        ctx,
		config:     cfg,
		Album:      NewAlbumClient(cfg),
		AlbumImage: NewAlbumImageClient(cfg),
		Artist:     NewArtistClient(cfg),
		LikedTrack: NewLikedTrackClient(cfg),
		Track:      NewTrackClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Album.
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
	c.Album.Use(hooks...)
	c.AlbumImage.Use(hooks...)
	c.Artist.Use(hooks...)
	c.LikedTrack.Use(hooks...)
	c.Track.Use(hooks...)
}

// AlbumClient is a client for the Album schema.
type AlbumClient struct {
	config
}

// NewAlbumClient returns a client for the Album from the given config.
func NewAlbumClient(c config) *AlbumClient {
	return &AlbumClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `album.Hooks(f(g(h())))`.
func (c *AlbumClient) Use(hooks ...Hook) {
	c.hooks.Album = append(c.hooks.Album, hooks...)
}

// Create returns a create builder for Album.
func (c *AlbumClient) Create() *AlbumCreate {
	mutation := newAlbumMutation(c.config, OpCreate)
	return &AlbumCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Album entities.
func (c *AlbumClient) CreateBulk(builders ...*AlbumCreate) *AlbumCreateBulk {
	return &AlbumCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Album.
func (c *AlbumClient) Update() *AlbumUpdate {
	mutation := newAlbumMutation(c.config, OpUpdate)
	return &AlbumUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AlbumClient) UpdateOne(a *Album) *AlbumUpdateOne {
	mutation := newAlbumMutation(c.config, OpUpdateOne, withAlbum(a))
	return &AlbumUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AlbumClient) UpdateOneID(id int) *AlbumUpdateOne {
	mutation := newAlbumMutation(c.config, OpUpdateOne, withAlbumID(id))
	return &AlbumUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Album.
func (c *AlbumClient) Delete() *AlbumDelete {
	mutation := newAlbumMutation(c.config, OpDelete)
	return &AlbumDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *AlbumClient) DeleteOne(a *Album) *AlbumDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *AlbumClient) DeleteOneID(id int) *AlbumDeleteOne {
	builder := c.Delete().Where(album.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AlbumDeleteOne{builder}
}

// Query returns a query builder for Album.
func (c *AlbumClient) Query() *AlbumQuery {
	return &AlbumQuery{
		config: c.config,
	}
}

// Get returns a Album entity by its id.
func (c *AlbumClient) Get(ctx context.Context, id int) (*Album, error) {
	return c.Query().Where(album.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AlbumClient) GetX(ctx context.Context, id int) *Album {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryArtist queries the artist edge of a Album.
func (c *AlbumClient) QueryArtist(a *Album) *ArtistQuery {
	query := &ArtistQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(album.Table, album.FieldID, id),
			sqlgraph.To(artist.Table, artist.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, album.ArtistTable, album.ArtistColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryTracks queries the tracks edge of a Album.
func (c *AlbumClient) QueryTracks(a *Album) *TrackQuery {
	query := &TrackQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(album.Table, album.FieldID, id),
			sqlgraph.To(track.Table, track.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, album.TracksTable, album.TracksColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCover queries the cover edge of a Album.
func (c *AlbumClient) QueryCover(a *Album) *AlbumImageQuery {
	query := &AlbumImageQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(album.Table, album.FieldID, id),
			sqlgraph.To(albumimage.Table, albumimage.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, album.CoverTable, album.CoverColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *AlbumClient) Hooks() []Hook {
	return c.hooks.Album
}

// AlbumImageClient is a client for the AlbumImage schema.
type AlbumImageClient struct {
	config
}

// NewAlbumImageClient returns a client for the AlbumImage from the given config.
func NewAlbumImageClient(c config) *AlbumImageClient {
	return &AlbumImageClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `albumimage.Hooks(f(g(h())))`.
func (c *AlbumImageClient) Use(hooks ...Hook) {
	c.hooks.AlbumImage = append(c.hooks.AlbumImage, hooks...)
}

// Create returns a create builder for AlbumImage.
func (c *AlbumImageClient) Create() *AlbumImageCreate {
	mutation := newAlbumImageMutation(c.config, OpCreate)
	return &AlbumImageCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of AlbumImage entities.
func (c *AlbumImageClient) CreateBulk(builders ...*AlbumImageCreate) *AlbumImageCreateBulk {
	return &AlbumImageCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for AlbumImage.
func (c *AlbumImageClient) Update() *AlbumImageUpdate {
	mutation := newAlbumImageMutation(c.config, OpUpdate)
	return &AlbumImageUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AlbumImageClient) UpdateOne(ai *AlbumImage) *AlbumImageUpdateOne {
	mutation := newAlbumImageMutation(c.config, OpUpdateOne, withAlbumImage(ai))
	return &AlbumImageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AlbumImageClient) UpdateOneID(id int) *AlbumImageUpdateOne {
	mutation := newAlbumImageMutation(c.config, OpUpdateOne, withAlbumImageID(id))
	return &AlbumImageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for AlbumImage.
func (c *AlbumImageClient) Delete() *AlbumImageDelete {
	mutation := newAlbumImageMutation(c.config, OpDelete)
	return &AlbumImageDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *AlbumImageClient) DeleteOne(ai *AlbumImage) *AlbumImageDeleteOne {
	return c.DeleteOneID(ai.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *AlbumImageClient) DeleteOneID(id int) *AlbumImageDeleteOne {
	builder := c.Delete().Where(albumimage.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AlbumImageDeleteOne{builder}
}

// Query returns a query builder for AlbumImage.
func (c *AlbumImageClient) Query() *AlbumImageQuery {
	return &AlbumImageQuery{
		config: c.config,
	}
}

// Get returns a AlbumImage entity by its id.
func (c *AlbumImageClient) Get(ctx context.Context, id int) (*AlbumImage, error) {
	return c.Query().Where(albumimage.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AlbumImageClient) GetX(ctx context.Context, id int) *AlbumImage {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *AlbumImageClient) Hooks() []Hook {
	return c.hooks.AlbumImage
}

// ArtistClient is a client for the Artist schema.
type ArtistClient struct {
	config
}

// NewArtistClient returns a client for the Artist from the given config.
func NewArtistClient(c config) *ArtistClient {
	return &ArtistClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `artist.Hooks(f(g(h())))`.
func (c *ArtistClient) Use(hooks ...Hook) {
	c.hooks.Artist = append(c.hooks.Artist, hooks...)
}

// Create returns a create builder for Artist.
func (c *ArtistClient) Create() *ArtistCreate {
	mutation := newArtistMutation(c.config, OpCreate)
	return &ArtistCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Artist entities.
func (c *ArtistClient) CreateBulk(builders ...*ArtistCreate) *ArtistCreateBulk {
	return &ArtistCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Artist.
func (c *ArtistClient) Update() *ArtistUpdate {
	mutation := newArtistMutation(c.config, OpUpdate)
	return &ArtistUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ArtistClient) UpdateOne(a *Artist) *ArtistUpdateOne {
	mutation := newArtistMutation(c.config, OpUpdateOne, withArtist(a))
	return &ArtistUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ArtistClient) UpdateOneID(id int) *ArtistUpdateOne {
	mutation := newArtistMutation(c.config, OpUpdateOne, withArtistID(id))
	return &ArtistUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Artist.
func (c *ArtistClient) Delete() *ArtistDelete {
	mutation := newArtistMutation(c.config, OpDelete)
	return &ArtistDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ArtistClient) DeleteOne(a *Artist) *ArtistDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ArtistClient) DeleteOneID(id int) *ArtistDeleteOne {
	builder := c.Delete().Where(artist.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ArtistDeleteOne{builder}
}

// Query returns a query builder for Artist.
func (c *ArtistClient) Query() *ArtistQuery {
	return &ArtistQuery{
		config: c.config,
	}
}

// Get returns a Artist entity by its id.
func (c *ArtistClient) Get(ctx context.Context, id int) (*Artist, error) {
	return c.Query().Where(artist.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ArtistClient) GetX(ctx context.Context, id int) *Artist {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryAlbums queries the albums edge of a Artist.
func (c *ArtistClient) QueryAlbums(a *Artist) *AlbumQuery {
	query := &AlbumQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(artist.Table, artist.FieldID, id),
			sqlgraph.To(album.Table, album.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, artist.AlbumsTable, artist.AlbumsColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryTracks queries the tracks edge of a Artist.
func (c *ArtistClient) QueryTracks(a *Artist) *TrackQuery {
	query := &TrackQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(artist.Table, artist.FieldID, id),
			sqlgraph.To(track.Table, track.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, artist.TracksTable, artist.TracksPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ArtistClient) Hooks() []Hook {
	return c.hooks.Artist
}

// LikedTrackClient is a client for the LikedTrack schema.
type LikedTrackClient struct {
	config
}

// NewLikedTrackClient returns a client for the LikedTrack from the given config.
func NewLikedTrackClient(c config) *LikedTrackClient {
	return &LikedTrackClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `likedtrack.Hooks(f(g(h())))`.
func (c *LikedTrackClient) Use(hooks ...Hook) {
	c.hooks.LikedTrack = append(c.hooks.LikedTrack, hooks...)
}

// Create returns a create builder for LikedTrack.
func (c *LikedTrackClient) Create() *LikedTrackCreate {
	mutation := newLikedTrackMutation(c.config, OpCreate)
	return &LikedTrackCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of LikedTrack entities.
func (c *LikedTrackClient) CreateBulk(builders ...*LikedTrackCreate) *LikedTrackCreateBulk {
	return &LikedTrackCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for LikedTrack.
func (c *LikedTrackClient) Update() *LikedTrackUpdate {
	mutation := newLikedTrackMutation(c.config, OpUpdate)
	return &LikedTrackUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *LikedTrackClient) UpdateOne(lt *LikedTrack) *LikedTrackUpdateOne {
	mutation := newLikedTrackMutation(c.config, OpUpdateOne, withLikedTrack(lt))
	return &LikedTrackUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *LikedTrackClient) UpdateOneID(id int) *LikedTrackUpdateOne {
	mutation := newLikedTrackMutation(c.config, OpUpdateOne, withLikedTrackID(id))
	return &LikedTrackUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for LikedTrack.
func (c *LikedTrackClient) Delete() *LikedTrackDelete {
	mutation := newLikedTrackMutation(c.config, OpDelete)
	return &LikedTrackDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *LikedTrackClient) DeleteOne(lt *LikedTrack) *LikedTrackDeleteOne {
	return c.DeleteOneID(lt.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *LikedTrackClient) DeleteOneID(id int) *LikedTrackDeleteOne {
	builder := c.Delete().Where(likedtrack.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &LikedTrackDeleteOne{builder}
}

// Query returns a query builder for LikedTrack.
func (c *LikedTrackClient) Query() *LikedTrackQuery {
	return &LikedTrackQuery{
		config: c.config,
	}
}

// Get returns a LikedTrack entity by its id.
func (c *LikedTrackClient) Get(ctx context.Context, id int) (*LikedTrack, error) {
	return c.Query().Where(likedtrack.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *LikedTrackClient) GetX(ctx context.Context, id int) *LikedTrack {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryTrack queries the track edge of a LikedTrack.
func (c *LikedTrackClient) QueryTrack(lt *LikedTrack) *TrackQuery {
	query := &TrackQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := lt.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(likedtrack.Table, likedtrack.FieldID, id),
			sqlgraph.To(track.Table, track.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, likedtrack.TrackTable, likedtrack.TrackColumn),
		)
		fromV = sqlgraph.Neighbors(lt.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *LikedTrackClient) Hooks() []Hook {
	return c.hooks.LikedTrack
}

// TrackClient is a client for the Track schema.
type TrackClient struct {
	config
}

// NewTrackClient returns a client for the Track from the given config.
func NewTrackClient(c config) *TrackClient {
	return &TrackClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `track.Hooks(f(g(h())))`.
func (c *TrackClient) Use(hooks ...Hook) {
	c.hooks.Track = append(c.hooks.Track, hooks...)
}

// Create returns a create builder for Track.
func (c *TrackClient) Create() *TrackCreate {
	mutation := newTrackMutation(c.config, OpCreate)
	return &TrackCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Track entities.
func (c *TrackClient) CreateBulk(builders ...*TrackCreate) *TrackCreateBulk {
	return &TrackCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Track.
func (c *TrackClient) Update() *TrackUpdate {
	mutation := newTrackMutation(c.config, OpUpdate)
	return &TrackUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TrackClient) UpdateOne(t *Track) *TrackUpdateOne {
	mutation := newTrackMutation(c.config, OpUpdateOne, withTrack(t))
	return &TrackUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TrackClient) UpdateOneID(id int) *TrackUpdateOne {
	mutation := newTrackMutation(c.config, OpUpdateOne, withTrackID(id))
	return &TrackUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Track.
func (c *TrackClient) Delete() *TrackDelete {
	mutation := newTrackMutation(c.config, OpDelete)
	return &TrackDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *TrackClient) DeleteOne(t *Track) *TrackDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *TrackClient) DeleteOneID(id int) *TrackDeleteOne {
	builder := c.Delete().Where(track.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TrackDeleteOne{builder}
}

// Query returns a query builder for Track.
func (c *TrackClient) Query() *TrackQuery {
	return &TrackQuery{
		config: c.config,
	}
}

// Get returns a Track entity by its id.
func (c *TrackClient) Get(ctx context.Context, id int) (*Track, error) {
	return c.Query().Where(track.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TrackClient) GetX(ctx context.Context, id int) *Track {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryArtists queries the artists edge of a Track.
func (c *TrackClient) QueryArtists(t *Track) *ArtistQuery {
	query := &ArtistQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(track.Table, track.FieldID, id),
			sqlgraph.To(artist.Table, artist.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, track.ArtistsTable, track.ArtistsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryAlbum queries the album edge of a Track.
func (c *TrackClient) QueryAlbum(t *Track) *AlbumQuery {
	query := &AlbumQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(track.Table, track.FieldID, id),
			sqlgraph.To(album.Table, album.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, track.AlbumTable, track.AlbumColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryLiked queries the liked edge of a Track.
func (c *TrackClient) QueryLiked(t *Track) *LikedTrackQuery {
	query := &LikedTrackQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(track.Table, track.FieldID, id),
			sqlgraph.To(likedtrack.Table, likedtrack.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, track.LikedTable, track.LikedColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TrackClient) Hooks() []Hook {
	return c.hooks.Track
}
