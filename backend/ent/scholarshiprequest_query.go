// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/chat2542/app/ent/predicate"
	"github.com/chat2542/app/ent/scholarshipinformation"
	"github.com/chat2542/app/ent/scholarshiprequest"
	"github.com/chat2542/app/ent/scholarshiptype"
	"github.com/chat2542/app/ent/semester"
	"github.com/chat2542/app/ent/user"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// ScholarshipRequestQuery is the builder for querying ScholarshipRequest entities.
type ScholarshipRequestQuery struct {
	config
	limit      *int
	offset     *int
	order      []OrderFunc
	unique     []string
	predicates []predicate.ScholarshipRequest
	// eager-loading edges.
	withUser                   *UserQuery
	withScholarshiptype        *ScholarshiptypeQuery
	withScholarshipinformation *ScholarshipinformationQuery
	withSemester               *SemesterQuery
	withFKs                    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the builder.
func (srq *ScholarshipRequestQuery) Where(ps ...predicate.ScholarshipRequest) *ScholarshipRequestQuery {
	srq.predicates = append(srq.predicates, ps...)
	return srq
}

// Limit adds a limit step to the query.
func (srq *ScholarshipRequestQuery) Limit(limit int) *ScholarshipRequestQuery {
	srq.limit = &limit
	return srq
}

// Offset adds an offset step to the query.
func (srq *ScholarshipRequestQuery) Offset(offset int) *ScholarshipRequestQuery {
	srq.offset = &offset
	return srq
}

// Order adds an order step to the query.
func (srq *ScholarshipRequestQuery) Order(o ...OrderFunc) *ScholarshipRequestQuery {
	srq.order = append(srq.order, o...)
	return srq
}

// QueryUser chains the current query on the User edge.
func (srq *ScholarshipRequestQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: srq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := srq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(scholarshiprequest.Table, scholarshiprequest.FieldID, srq.sqlQuery()),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, scholarshiprequest.UserTable, scholarshiprequest.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(srq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryScholarshiptype chains the current query on the Scholarshiptype edge.
func (srq *ScholarshipRequestQuery) QueryScholarshiptype() *ScholarshiptypeQuery {
	query := &ScholarshiptypeQuery{config: srq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := srq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(scholarshiprequest.Table, scholarshiprequest.FieldID, srq.sqlQuery()),
			sqlgraph.To(scholarshiptype.Table, scholarshiptype.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, scholarshiprequest.ScholarshiptypeTable, scholarshiprequest.ScholarshiptypeColumn),
		)
		fromU = sqlgraph.SetNeighbors(srq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryScholarshipinformation chains the current query on the Scholarshipinformation edge.
func (srq *ScholarshipRequestQuery) QueryScholarshipinformation() *ScholarshipinformationQuery {
	query := &ScholarshipinformationQuery{config: srq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := srq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(scholarshiprequest.Table, scholarshiprequest.FieldID, srq.sqlQuery()),
			sqlgraph.To(scholarshipinformation.Table, scholarshipinformation.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, scholarshiprequest.ScholarshipinformationTable, scholarshiprequest.ScholarshipinformationColumn),
		)
		fromU = sqlgraph.SetNeighbors(srq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySemester chains the current query on the Semester edge.
func (srq *ScholarshipRequestQuery) QuerySemester() *SemesterQuery {
	query := &SemesterQuery{config: srq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := srq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(scholarshiprequest.Table, scholarshiprequest.FieldID, srq.sqlQuery()),
			sqlgraph.To(semester.Table, semester.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, scholarshiprequest.SemesterTable, scholarshiprequest.SemesterColumn),
		)
		fromU = sqlgraph.SetNeighbors(srq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ScholarshipRequest entity in the query. Returns *NotFoundError when no scholarshiprequest was found.
func (srq *ScholarshipRequestQuery) First(ctx context.Context) (*ScholarshipRequest, error) {
	srs, err := srq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(srs) == 0 {
		return nil, &NotFoundError{scholarshiprequest.Label}
	}
	return srs[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (srq *ScholarshipRequestQuery) FirstX(ctx context.Context) *ScholarshipRequest {
	sr, err := srq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return sr
}

// FirstID returns the first ScholarshipRequest id in the query. Returns *NotFoundError when no id was found.
func (srq *ScholarshipRequestQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = srq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{scholarshiprequest.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (srq *ScholarshipRequestQuery) FirstXID(ctx context.Context) int {
	id, err := srq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only ScholarshipRequest entity in the query, returns an error if not exactly one entity was returned.
func (srq *ScholarshipRequestQuery) Only(ctx context.Context) (*ScholarshipRequest, error) {
	srs, err := srq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(srs) {
	case 1:
		return srs[0], nil
	case 0:
		return nil, &NotFoundError{scholarshiprequest.Label}
	default:
		return nil, &NotSingularError{scholarshiprequest.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (srq *ScholarshipRequestQuery) OnlyX(ctx context.Context) *ScholarshipRequest {
	sr, err := srq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return sr
}

// OnlyID returns the only ScholarshipRequest id in the query, returns an error if not exactly one id was returned.
func (srq *ScholarshipRequestQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = srq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{scholarshiprequest.Label}
	default:
		err = &NotSingularError{scholarshiprequest.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (srq *ScholarshipRequestQuery) OnlyIDX(ctx context.Context) int {
	id, err := srq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ScholarshipRequests.
func (srq *ScholarshipRequestQuery) All(ctx context.Context) ([]*ScholarshipRequest, error) {
	if err := srq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return srq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (srq *ScholarshipRequestQuery) AllX(ctx context.Context) []*ScholarshipRequest {
	srs, err := srq.All(ctx)
	if err != nil {
		panic(err)
	}
	return srs
}

// IDs executes the query and returns a list of ScholarshipRequest ids.
func (srq *ScholarshipRequestQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := srq.Select(scholarshiprequest.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (srq *ScholarshipRequestQuery) IDsX(ctx context.Context) []int {
	ids, err := srq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (srq *ScholarshipRequestQuery) Count(ctx context.Context) (int, error) {
	if err := srq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return srq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (srq *ScholarshipRequestQuery) CountX(ctx context.Context) int {
	count, err := srq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (srq *ScholarshipRequestQuery) Exist(ctx context.Context) (bool, error) {
	if err := srq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return srq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (srq *ScholarshipRequestQuery) ExistX(ctx context.Context) bool {
	exist, err := srq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (srq *ScholarshipRequestQuery) Clone() *ScholarshipRequestQuery {
	return &ScholarshipRequestQuery{
		config:     srq.config,
		limit:      srq.limit,
		offset:     srq.offset,
		order:      append([]OrderFunc{}, srq.order...),
		unique:     append([]string{}, srq.unique...),
		predicates: append([]predicate.ScholarshipRequest{}, srq.predicates...),
		// clone intermediate query.
		sql:  srq.sql.Clone(),
		path: srq.path,
	}
}

//  WithUser tells the query-builder to eager-loads the nodes that are connected to
// the "User" edge. The optional arguments used to configure the query builder of the edge.
func (srq *ScholarshipRequestQuery) WithUser(opts ...func(*UserQuery)) *ScholarshipRequestQuery {
	query := &UserQuery{config: srq.config}
	for _, opt := range opts {
		opt(query)
	}
	srq.withUser = query
	return srq
}

//  WithScholarshiptype tells the query-builder to eager-loads the nodes that are connected to
// the "Scholarshiptype" edge. The optional arguments used to configure the query builder of the edge.
func (srq *ScholarshipRequestQuery) WithScholarshiptype(opts ...func(*ScholarshiptypeQuery)) *ScholarshipRequestQuery {
	query := &ScholarshiptypeQuery{config: srq.config}
	for _, opt := range opts {
		opt(query)
	}
	srq.withScholarshiptype = query
	return srq
}

//  WithScholarshipinformation tells the query-builder to eager-loads the nodes that are connected to
// the "Scholarshipinformation" edge. The optional arguments used to configure the query builder of the edge.
func (srq *ScholarshipRequestQuery) WithScholarshipinformation(opts ...func(*ScholarshipinformationQuery)) *ScholarshipRequestQuery {
	query := &ScholarshipinformationQuery{config: srq.config}
	for _, opt := range opts {
		opt(query)
	}
	srq.withScholarshipinformation = query
	return srq
}

//  WithSemester tells the query-builder to eager-loads the nodes that are connected to
// the "Semester" edge. The optional arguments used to configure the query builder of the edge.
func (srq *ScholarshipRequestQuery) WithSemester(opts ...func(*SemesterQuery)) *ScholarshipRequestQuery {
	query := &SemesterQuery{config: srq.config}
	for _, opt := range opts {
		opt(query)
	}
	srq.withSemester = query
	return srq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (srq *ScholarshipRequestQuery) GroupBy(field string, fields ...string) *ScholarshipRequestGroupBy {
	group := &ScholarshipRequestGroupBy{config: srq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := srq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return srq.sqlQuery(), nil
	}
	return group
}

// Select one or more fields from the given query.
func (srq *ScholarshipRequestQuery) Select(field string, fields ...string) *ScholarshipRequestSelect {
	selector := &ScholarshipRequestSelect{config: srq.config}
	selector.fields = append([]string{field}, fields...)
	selector.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := srq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return srq.sqlQuery(), nil
	}
	return selector
}

func (srq *ScholarshipRequestQuery) prepareQuery(ctx context.Context) error {
	if srq.path != nil {
		prev, err := srq.path(ctx)
		if err != nil {
			return err
		}
		srq.sql = prev
	}
	return nil
}

func (srq *ScholarshipRequestQuery) sqlAll(ctx context.Context) ([]*ScholarshipRequest, error) {
	var (
		nodes       = []*ScholarshipRequest{}
		withFKs     = srq.withFKs
		_spec       = srq.querySpec()
		loadedTypes = [4]bool{
			srq.withUser != nil,
			srq.withScholarshiptype != nil,
			srq.withScholarshipinformation != nil,
			srq.withSemester != nil,
		}
	)
	if srq.withUser != nil || srq.withScholarshiptype != nil || srq.withScholarshipinformation != nil || srq.withSemester != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, scholarshiprequest.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &ScholarshipRequest{config: srq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, srq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := srq.withUser; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*ScholarshipRequest)
		for i := range nodes {
			if fk := nodes[i].user_scholarship_request; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "user_scholarship_request" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.User = n
			}
		}
	}

	if query := srq.withScholarshiptype; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*ScholarshipRequest)
		for i := range nodes {
			if fk := nodes[i].scholarshiptype_scholarship_request; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(scholarshiptype.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "scholarshiptype_scholarship_request" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Scholarshiptype = n
			}
		}
	}

	if query := srq.withScholarshipinformation; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*ScholarshipRequest)
		for i := range nodes {
			if fk := nodes[i].scholarshipinformation_scholarship_request; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(scholarshipinformation.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "scholarshipinformation_scholarship_request" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Scholarshipinformation = n
			}
		}
	}

	if query := srq.withSemester; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*ScholarshipRequest)
		for i := range nodes {
			if fk := nodes[i].semester_scholarship_request; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(semester.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "semester_scholarship_request" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Semester = n
			}
		}
	}

	return nodes, nil
}

func (srq *ScholarshipRequestQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := srq.querySpec()
	return sqlgraph.CountNodes(ctx, srq.driver, _spec)
}

func (srq *ScholarshipRequestQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := srq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (srq *ScholarshipRequestQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   scholarshiprequest.Table,
			Columns: scholarshiprequest.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: scholarshiprequest.FieldID,
			},
		},
		From:   srq.sql,
		Unique: true,
	}
	if ps := srq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := srq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := srq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := srq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (srq *ScholarshipRequestQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(srq.driver.Dialect())
	t1 := builder.Table(scholarshiprequest.Table)
	selector := builder.Select(t1.Columns(scholarshiprequest.Columns...)...).From(t1)
	if srq.sql != nil {
		selector = srq.sql
		selector.Select(selector.Columns(scholarshiprequest.Columns...)...)
	}
	for _, p := range srq.predicates {
		p(selector)
	}
	for _, p := range srq.order {
		p(selector)
	}
	if offset := srq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := srq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ScholarshipRequestGroupBy is the builder for group-by ScholarshipRequest entities.
type ScholarshipRequestGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (srgb *ScholarshipRequestGroupBy) Aggregate(fns ...AggregateFunc) *ScholarshipRequestGroupBy {
	srgb.fns = append(srgb.fns, fns...)
	return srgb
}

// Scan applies the group-by query and scan the result into the given value.
func (srgb *ScholarshipRequestGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := srgb.path(ctx)
	if err != nil {
		return err
	}
	srgb.sql = query
	return srgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (srgb *ScholarshipRequestGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := srgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (srgb *ScholarshipRequestGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(srgb.fields) > 1 {
		return nil, errors.New("ent: ScholarshipRequestGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := srgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (srgb *ScholarshipRequestGroupBy) StringsX(ctx context.Context) []string {
	v, err := srgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from group-by. It is only allowed when querying group-by with one field.
func (srgb *ScholarshipRequestGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = srgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{scholarshiprequest.Label}
	default:
		err = fmt.Errorf("ent: ScholarshipRequestGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (srgb *ScholarshipRequestGroupBy) StringX(ctx context.Context) string {
	v, err := srgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (srgb *ScholarshipRequestGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(srgb.fields) > 1 {
		return nil, errors.New("ent: ScholarshipRequestGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := srgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (srgb *ScholarshipRequestGroupBy) IntsX(ctx context.Context) []int {
	v, err := srgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from group-by. It is only allowed when querying group-by with one field.
func (srgb *ScholarshipRequestGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = srgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{scholarshiprequest.Label}
	default:
		err = fmt.Errorf("ent: ScholarshipRequestGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (srgb *ScholarshipRequestGroupBy) IntX(ctx context.Context) int {
	v, err := srgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (srgb *ScholarshipRequestGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(srgb.fields) > 1 {
		return nil, errors.New("ent: ScholarshipRequestGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := srgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (srgb *ScholarshipRequestGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := srgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from group-by. It is only allowed when querying group-by with one field.
func (srgb *ScholarshipRequestGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = srgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{scholarshiprequest.Label}
	default:
		err = fmt.Errorf("ent: ScholarshipRequestGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (srgb *ScholarshipRequestGroupBy) Float64X(ctx context.Context) float64 {
	v, err := srgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (srgb *ScholarshipRequestGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(srgb.fields) > 1 {
		return nil, errors.New("ent: ScholarshipRequestGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := srgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (srgb *ScholarshipRequestGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := srgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from group-by. It is only allowed when querying group-by with one field.
func (srgb *ScholarshipRequestGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = srgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{scholarshiprequest.Label}
	default:
		err = fmt.Errorf("ent: ScholarshipRequestGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (srgb *ScholarshipRequestGroupBy) BoolX(ctx context.Context) bool {
	v, err := srgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (srgb *ScholarshipRequestGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := srgb.sqlQuery().Query()
	if err := srgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (srgb *ScholarshipRequestGroupBy) sqlQuery() *sql.Selector {
	selector := srgb.sql
	columns := make([]string, 0, len(srgb.fields)+len(srgb.fns))
	columns = append(columns, srgb.fields...)
	for _, fn := range srgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(srgb.fields...)
}

// ScholarshipRequestSelect is the builder for select fields of ScholarshipRequest entities.
type ScholarshipRequestSelect struct {
	config
	fields []string
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Scan applies the selector query and scan the result into the given value.
func (srs *ScholarshipRequestSelect) Scan(ctx context.Context, v interface{}) error {
	query, err := srs.path(ctx)
	if err != nil {
		return err
	}
	srs.sql = query
	return srs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (srs *ScholarshipRequestSelect) ScanX(ctx context.Context, v interface{}) {
	if err := srs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (srs *ScholarshipRequestSelect) Strings(ctx context.Context) ([]string, error) {
	if len(srs.fields) > 1 {
		return nil, errors.New("ent: ScholarshipRequestSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := srs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (srs *ScholarshipRequestSelect) StringsX(ctx context.Context) []string {
	v, err := srs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from selector. It is only allowed when selecting one field.
func (srs *ScholarshipRequestSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = srs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{scholarshiprequest.Label}
	default:
		err = fmt.Errorf("ent: ScholarshipRequestSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (srs *ScholarshipRequestSelect) StringX(ctx context.Context) string {
	v, err := srs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (srs *ScholarshipRequestSelect) Ints(ctx context.Context) ([]int, error) {
	if len(srs.fields) > 1 {
		return nil, errors.New("ent: ScholarshipRequestSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := srs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (srs *ScholarshipRequestSelect) IntsX(ctx context.Context) []int {
	v, err := srs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from selector. It is only allowed when selecting one field.
func (srs *ScholarshipRequestSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = srs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{scholarshiprequest.Label}
	default:
		err = fmt.Errorf("ent: ScholarshipRequestSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (srs *ScholarshipRequestSelect) IntX(ctx context.Context) int {
	v, err := srs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (srs *ScholarshipRequestSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(srs.fields) > 1 {
		return nil, errors.New("ent: ScholarshipRequestSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := srs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (srs *ScholarshipRequestSelect) Float64sX(ctx context.Context) []float64 {
	v, err := srs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from selector. It is only allowed when selecting one field.
func (srs *ScholarshipRequestSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = srs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{scholarshiprequest.Label}
	default:
		err = fmt.Errorf("ent: ScholarshipRequestSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (srs *ScholarshipRequestSelect) Float64X(ctx context.Context) float64 {
	v, err := srs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (srs *ScholarshipRequestSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(srs.fields) > 1 {
		return nil, errors.New("ent: ScholarshipRequestSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := srs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (srs *ScholarshipRequestSelect) BoolsX(ctx context.Context) []bool {
	v, err := srs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from selector. It is only allowed when selecting one field.
func (srs *ScholarshipRequestSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = srs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{scholarshiprequest.Label}
	default:
		err = fmt.Errorf("ent: ScholarshipRequestSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (srs *ScholarshipRequestSelect) BoolX(ctx context.Context) bool {
	v, err := srs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (srs *ScholarshipRequestSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := srs.sqlQuery().Query()
	if err := srs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (srs *ScholarshipRequestSelect) sqlQuery() sql.Querier {
	selector := srs.sql
	selector.Select(selector.Columns(srs.fields...)...)
	return selector
}
