// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/chat2542/app/ent/scholarshipinformation"
	"github.com/chat2542/app/ent/scholarshiprequest"
	"github.com/chat2542/app/ent/scholarshiptype"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// ScholarshipinformationCreate is the builder for creating a Scholarshipinformation entity.
type ScholarshipinformationCreate struct {
	config
	mutation *ScholarshipinformationMutation
	hooks    []Hook
}

// SetScholarshipName sets the Scholarship_name field.
func (sc *ScholarshipinformationCreate) SetScholarshipName(s string) *ScholarshipinformationCreate {
	sc.mutation.SetScholarshipName(s)
	return sc
}

// SetScholarshiptypeID sets the Scholarshiptype edge to Scholarshiptype by id.
func (sc *ScholarshipinformationCreate) SetScholarshiptypeID(id int) *ScholarshipinformationCreate {
	sc.mutation.SetScholarshiptypeID(id)
	return sc
}

// SetNillableScholarshiptypeID sets the Scholarshiptype edge to Scholarshiptype by id if the given value is not nil.
func (sc *ScholarshipinformationCreate) SetNillableScholarshiptypeID(id *int) *ScholarshipinformationCreate {
	if id != nil {
		sc = sc.SetScholarshiptypeID(*id)
	}
	return sc
}

// SetScholarshiptype sets the Scholarshiptype edge to Scholarshiptype.
func (sc *ScholarshipinformationCreate) SetScholarshiptype(s *Scholarshiptype) *ScholarshipinformationCreate {
	return sc.SetScholarshiptypeID(s.ID)
}

// AddScholarshipRequestIDs adds the ScholarshipRequest edge to ScholarshipRequest by ids.
func (sc *ScholarshipinformationCreate) AddScholarshipRequestIDs(ids ...int) *ScholarshipinformationCreate {
	sc.mutation.AddScholarshipRequestIDs(ids...)
	return sc
}

// AddScholarshipRequest adds the ScholarshipRequest edges to ScholarshipRequest.
func (sc *ScholarshipinformationCreate) AddScholarshipRequest(s ...*ScholarshipRequest) *ScholarshipinformationCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sc.AddScholarshipRequestIDs(ids...)
}

// Mutation returns the ScholarshipinformationMutation object of the builder.
func (sc *ScholarshipinformationCreate) Mutation() *ScholarshipinformationMutation {
	return sc.mutation
}

// Save creates the Scholarshipinformation in the database.
func (sc *ScholarshipinformationCreate) Save(ctx context.Context) (*Scholarshipinformation, error) {
	if _, ok := sc.mutation.ScholarshipName(); !ok {
		return nil, &ValidationError{Name: "Scholarship_name", err: errors.New("ent: missing required field \"Scholarship_name\"")}
	}
	if v, ok := sc.mutation.ScholarshipName(); ok {
		if err := scholarshipinformation.ScholarshipNameValidator(v); err != nil {
			return nil, &ValidationError{Name: "Scholarship_name", err: fmt.Errorf("ent: validator failed for field \"Scholarship_name\": %w", err)}
		}
	}
	var (
		err  error
		node *Scholarshipinformation
	)
	if len(sc.hooks) == 0 {
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ScholarshipinformationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sc.mutation = mutation
			node, err = sc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			mut = sc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *ScholarshipinformationCreate) SaveX(ctx context.Context) *Scholarshipinformation {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sc *ScholarshipinformationCreate) sqlSave(ctx context.Context) (*Scholarshipinformation, error) {
	s, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	s.ID = int(id)
	return s, nil
}

func (sc *ScholarshipinformationCreate) createSpec() (*Scholarshipinformation, *sqlgraph.CreateSpec) {
	var (
		s     = &Scholarshipinformation{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: scholarshipinformation.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: scholarshipinformation.FieldID,
			},
		}
	)
	if value, ok := sc.mutation.ScholarshipName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: scholarshipinformation.FieldScholarshipName,
		})
		s.ScholarshipName = value
	}
	if nodes := sc.mutation.ScholarshiptypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   scholarshipinformation.ScholarshiptypeTable,
			Columns: []string{scholarshipinformation.ScholarshiptypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: scholarshiptype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.ScholarshipRequestIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   scholarshipinformation.ScholarshipRequestTable,
			Columns: []string{scholarshipinformation.ScholarshipRequestColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: scholarshiprequest.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return s, _spec
}
