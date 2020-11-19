// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/chat2542/app/ent/semester"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Semester is the model entity for the Semester schema.
type Semester struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Semester holds the value of the "Semester" field.
	Semester string `json:"Semester,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SemesterQuery when eager-loading is set.
	Edges SemesterEdges `json:"edges"`
}

// SemesterEdges holds the relations/edges for other nodes in the graph.
type SemesterEdges struct {
	// ScholarshipRequest holds the value of the ScholarshipRequest edge.
	ScholarshipRequest []*ScholarshipRequest
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ScholarshipRequestOrErr returns the ScholarshipRequest value or an error if the edge
// was not loaded in eager-loading.
func (e SemesterEdges) ScholarshipRequestOrErr() ([]*ScholarshipRequest, error) {
	if e.loadedTypes[0] {
		return e.ScholarshipRequest, nil
	}
	return nil, &NotLoadedError{edge: "ScholarshipRequest"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Semester) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // Semester
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Semester fields.
func (s *Semester) assignValues(values ...interface{}) error {
	if m, n := len(values), len(semester.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	s.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Semester", values[0])
	} else if value.Valid {
		s.Semester = value.String
	}
	return nil
}

// QueryScholarshipRequest queries the ScholarshipRequest edge of the Semester.
func (s *Semester) QueryScholarshipRequest() *ScholarshipRequestQuery {
	return (&SemesterClient{config: s.config}).QueryScholarshipRequest(s)
}

// Update returns a builder for updating this Semester.
// Note that, you need to call Semester.Unwrap() before calling this method, if this Semester
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Semester) Update() *SemesterUpdateOne {
	return (&SemesterClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (s *Semester) Unwrap() *Semester {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Semester is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Semester) String() string {
	var builder strings.Builder
	builder.WriteString("Semester(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", Semester=")
	builder.WriteString(s.Semester)
	builder.WriteByte(')')
	return builder.String()
}

// Semesters is a parsable slice of Semester.
type Semesters []*Semester

func (s Semesters) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
