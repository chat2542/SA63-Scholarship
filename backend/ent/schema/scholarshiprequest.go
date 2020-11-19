package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
)

// ScholarshipRequest holds the schema definition for the ScholarshipRequest entity.
type ScholarshipRequest struct {
	ent.Schema
}

// Fields of the ScholarshipRequest.
func (ScholarshipRequest) Fields() []ent.Field {
	return nil
}

// Edges of the ScholarshipRequest.
func (ScholarshipRequest) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("User", User.Type).
			Ref("ScholarshipRequest").
			Unique(),
		edge.From("Scholarshiptype", Scholarshiptype.Type).
			Ref("ScholarshipRequest").
			Unique(),
		edge.From("Scholarshipinformation", Scholarshipinformation.Type).
			Ref("ScholarshipRequest").
			Unique(),
		edge.From("Semester", Semester.Type).
			Ref("ScholarshipRequest").
			Unique(),
	}
}
