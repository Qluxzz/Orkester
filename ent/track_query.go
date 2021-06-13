// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"goreact/ent/album"
	"goreact/ent/artist"
	"goreact/ent/predicate"
	"goreact/ent/track"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TrackQuery is the builder for querying Track entities.
type TrackQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Track
	// eager-loading edges.
	withArtists *ArtistQuery
	withAlbum   *AlbumQuery
	withFKs     bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TrackQuery builder.
func (tq *TrackQuery) Where(ps ...predicate.Track) *TrackQuery {
	tq.predicates = append(tq.predicates, ps...)
	return tq
}

// Limit adds a limit step to the query.
func (tq *TrackQuery) Limit(limit int) *TrackQuery {
	tq.limit = &limit
	return tq
}

// Offset adds an offset step to the query.
func (tq *TrackQuery) Offset(offset int) *TrackQuery {
	tq.offset = &offset
	return tq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tq *TrackQuery) Unique(unique bool) *TrackQuery {
	tq.unique = &unique
	return tq
}

// Order adds an order step to the query.
func (tq *TrackQuery) Order(o ...OrderFunc) *TrackQuery {
	tq.order = append(tq.order, o...)
	return tq
}

// QueryArtists chains the current query on the "artists" edge.
func (tq *TrackQuery) QueryArtists() *ArtistQuery {
	query := &ArtistQuery{config: tq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(track.Table, track.FieldID, selector),
			sqlgraph.To(artist.Table, artist.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, track.ArtistsTable, track.ArtistsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryAlbum chains the current query on the "album" edge.
func (tq *TrackQuery) QueryAlbum() *AlbumQuery {
	query := &AlbumQuery{config: tq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(track.Table, track.FieldID, selector),
			sqlgraph.To(album.Table, album.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, track.AlbumTable, track.AlbumColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Track entity from the query.
// Returns a *NotFoundError when no Track was found.
func (tq *TrackQuery) First(ctx context.Context) (*Track, error) {
	nodes, err := tq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{track.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tq *TrackQuery) FirstX(ctx context.Context) *Track {
	node, err := tq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Track ID from the query.
// Returns a *NotFoundError when no Track ID was found.
func (tq *TrackQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{track.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tq *TrackQuery) FirstIDX(ctx context.Context) int {
	id, err := tq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Track entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one Track entity is not found.
// Returns a *NotFoundError when no Track entities are found.
func (tq *TrackQuery) Only(ctx context.Context) (*Track, error) {
	nodes, err := tq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{track.Label}
	default:
		return nil, &NotSingularError{track.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tq *TrackQuery) OnlyX(ctx context.Context) *Track {
	node, err := tq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Track ID in the query.
// Returns a *NotSingularError when exactly one Track ID is not found.
// Returns a *NotFoundError when no entities are found.
func (tq *TrackQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{track.Label}
	default:
		err = &NotSingularError{track.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tq *TrackQuery) OnlyIDX(ctx context.Context) int {
	id, err := tq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Tracks.
func (tq *TrackQuery) All(ctx context.Context) ([]*Track, error) {
	if err := tq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return tq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (tq *TrackQuery) AllX(ctx context.Context) []*Track {
	nodes, err := tq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Track IDs.
func (tq *TrackQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := tq.Select(track.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tq *TrackQuery) IDsX(ctx context.Context) []int {
	ids, err := tq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tq *TrackQuery) Count(ctx context.Context) (int, error) {
	if err := tq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return tq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (tq *TrackQuery) CountX(ctx context.Context) int {
	count, err := tq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tq *TrackQuery) Exist(ctx context.Context) (bool, error) {
	if err := tq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return tq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (tq *TrackQuery) ExistX(ctx context.Context) bool {
	exist, err := tq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TrackQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tq *TrackQuery) Clone() *TrackQuery {
	if tq == nil {
		return nil
	}
	return &TrackQuery{
		config:      tq.config,
		limit:       tq.limit,
		offset:      tq.offset,
		order:       append([]OrderFunc{}, tq.order...),
		predicates:  append([]predicate.Track{}, tq.predicates...),
		withArtists: tq.withArtists.Clone(),
		withAlbum:   tq.withAlbum.Clone(),
		// clone intermediate query.
		sql:  tq.sql.Clone(),
		path: tq.path,
	}
}

// WithArtists tells the query-builder to eager-load the nodes that are connected to
// the "artists" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TrackQuery) WithArtists(opts ...func(*ArtistQuery)) *TrackQuery {
	query := &ArtistQuery{config: tq.config}
	for _, opt := range opts {
		opt(query)
	}
	tq.withArtists = query
	return tq
}

// WithAlbum tells the query-builder to eager-load the nodes that are connected to
// the "album" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TrackQuery) WithAlbum(opts ...func(*AlbumQuery)) *TrackQuery {
	query := &AlbumQuery{config: tq.config}
	for _, opt := range opts {
		opt(query)
	}
	tq.withAlbum = query
	return tq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Track.Query().
//		GroupBy(track.FieldTitle).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (tq *TrackQuery) GroupBy(field string, fields ...string) *TrackGroupBy {
	group := &TrackGroupBy{config: tq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return tq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Title string `json:"title,omitempty"`
//	}
//
//	client.Track.Query().
//		Select(track.FieldTitle).
//		Scan(ctx, &v)
//
func (tq *TrackQuery) Select(field string, fields ...string) *TrackSelect {
	tq.fields = append([]string{field}, fields...)
	return &TrackSelect{TrackQuery: tq}
}

func (tq *TrackQuery) prepareQuery(ctx context.Context) error {
	for _, f := range tq.fields {
		if !track.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tq.path != nil {
		prev, err := tq.path(ctx)
		if err != nil {
			return err
		}
		tq.sql = prev
	}
	return nil
}

func (tq *TrackQuery) sqlAll(ctx context.Context) ([]*Track, error) {
	var (
		nodes       = []*Track{}
		withFKs     = tq.withFKs
		_spec       = tq.querySpec()
		loadedTypes = [2]bool{
			tq.withArtists != nil,
			tq.withAlbum != nil,
		}
	)
	if tq.withAlbum != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, track.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &Track{config: tq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, tq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := tq.withArtists; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[int]*Track, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.Artists = []*Artist{}
		}
		var (
			edgeids []int
			edges   = make(map[int][]*Track)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: false,
				Table:   track.ArtistsTable,
				Columns: track.ArtistsPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(track.ArtistsPrimaryKey[0], fks...))
			},
			ScanValues: func() [2]interface{} {
				return [2]interface{}{new(sql.NullInt64), new(sql.NullInt64)}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullInt64)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := int(eout.Int64)
				inValue := int(ein.Int64)
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				if _, ok := edges[inValue]; !ok {
					edgeids = append(edgeids, inValue)
				}
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, tq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "artists": %w`, err)
		}
		query.Where(artist.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "artists" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Artists = append(nodes[i].Edges.Artists, n)
			}
		}
	}

	if query := tq.withAlbum; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Track)
		for i := range nodes {
			if nodes[i].album_tracks == nil {
				continue
			}
			fk := *nodes[i].album_tracks
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(album.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "album_tracks" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Album = n
			}
		}
	}

	return nodes, nil
}

func (tq *TrackQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tq.querySpec()
	return sqlgraph.CountNodes(ctx, tq.driver, _spec)
}

func (tq *TrackQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := tq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (tq *TrackQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   track.Table,
			Columns: track.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: track.FieldID,
			},
		},
		From:   tq.sql,
		Unique: true,
	}
	if unique := tq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := tq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, track.FieldID)
		for i := range fields {
			if fields[i] != track.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tq *TrackQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tq.driver.Dialect())
	t1 := builder.Table(track.Table)
	columns := tq.fields
	if len(columns) == 0 {
		columns = track.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tq.sql != nil {
		selector = tq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range tq.predicates {
		p(selector)
	}
	for _, p := range tq.order {
		p(selector)
	}
	if offset := tq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TrackGroupBy is the group-by builder for Track entities.
type TrackGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tgb *TrackGroupBy) Aggregate(fns ...AggregateFunc) *TrackGroupBy {
	tgb.fns = append(tgb.fns, fns...)
	return tgb
}

// Scan applies the group-by query and scans the result into the given value.
func (tgb *TrackGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := tgb.path(ctx)
	if err != nil {
		return err
	}
	tgb.sql = query
	return tgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (tgb *TrackGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := tgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (tgb *TrackGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(tgb.fields) > 1 {
		return nil, errors.New("ent: TrackGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := tgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (tgb *TrackGroupBy) StringsX(ctx context.Context) []string {
	v, err := tgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tgb *TrackGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = tgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{track.Label}
	default:
		err = fmt.Errorf("ent: TrackGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (tgb *TrackGroupBy) StringX(ctx context.Context) string {
	v, err := tgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (tgb *TrackGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(tgb.fields) > 1 {
		return nil, errors.New("ent: TrackGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := tgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (tgb *TrackGroupBy) IntsX(ctx context.Context) []int {
	v, err := tgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tgb *TrackGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = tgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{track.Label}
	default:
		err = fmt.Errorf("ent: TrackGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (tgb *TrackGroupBy) IntX(ctx context.Context) int {
	v, err := tgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (tgb *TrackGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(tgb.fields) > 1 {
		return nil, errors.New("ent: TrackGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := tgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (tgb *TrackGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := tgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tgb *TrackGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = tgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{track.Label}
	default:
		err = fmt.Errorf("ent: TrackGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (tgb *TrackGroupBy) Float64X(ctx context.Context) float64 {
	v, err := tgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (tgb *TrackGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(tgb.fields) > 1 {
		return nil, errors.New("ent: TrackGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := tgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (tgb *TrackGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := tgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (tgb *TrackGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = tgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{track.Label}
	default:
		err = fmt.Errorf("ent: TrackGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (tgb *TrackGroupBy) BoolX(ctx context.Context) bool {
	v, err := tgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (tgb *TrackGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range tgb.fields {
		if !track.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := tgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (tgb *TrackGroupBy) sqlQuery() *sql.Selector {
	selector := tgb.sql.Select()
	aggregation := make([]string, 0, len(tgb.fns))
	for _, fn := range tgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(tgb.fields)+len(tgb.fns))
		for _, f := range tgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(tgb.fields...)...)
}

// TrackSelect is the builder for selecting fields of Track entities.
type TrackSelect struct {
	*TrackQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ts *TrackSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ts.prepareQuery(ctx); err != nil {
		return err
	}
	ts.sql = ts.TrackQuery.sqlQuery(ctx)
	return ts.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ts *TrackSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ts.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ts *TrackSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ts.fields) > 1 {
		return nil, errors.New("ent: TrackSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ts *TrackSelect) StringsX(ctx context.Context) []string {
	v, err := ts.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ts *TrackSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ts.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{track.Label}
	default:
		err = fmt.Errorf("ent: TrackSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ts *TrackSelect) StringX(ctx context.Context) string {
	v, err := ts.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ts *TrackSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ts.fields) > 1 {
		return nil, errors.New("ent: TrackSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ts *TrackSelect) IntsX(ctx context.Context) []int {
	v, err := ts.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ts *TrackSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ts.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{track.Label}
	default:
		err = fmt.Errorf("ent: TrackSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ts *TrackSelect) IntX(ctx context.Context) int {
	v, err := ts.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ts *TrackSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ts.fields) > 1 {
		return nil, errors.New("ent: TrackSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ts *TrackSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ts.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ts *TrackSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ts.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{track.Label}
	default:
		err = fmt.Errorf("ent: TrackSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ts *TrackSelect) Float64X(ctx context.Context) float64 {
	v, err := ts.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ts *TrackSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ts.fields) > 1 {
		return nil, errors.New("ent: TrackSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ts.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ts *TrackSelect) BoolsX(ctx context.Context) []bool {
	v, err := ts.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ts *TrackSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ts.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{track.Label}
	default:
		err = fmt.Errorf("ent: TrackSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ts *TrackSelect) BoolX(ctx context.Context) bool {
	v, err := ts.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ts *TrackSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ts.sql.Query()
	if err := ts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
