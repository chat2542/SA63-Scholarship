package schema

import (
    "github.com/facebookincubator/ent"
    "github.com/facebookincubator/ent/schema/edge"
    "github.com/facebookincubator/ent/schema/field"
)

// Scholarshipinformation holds the schema definition for the Scholarshipinformation entity.
type Scholarshipinformation struct {
    ent.Schema
}

// Fields of the Scholarshipinformation.
func (Scholarshipinformation) Fields() []ent.Field {
    return []ent.Field{
        field.String("Scholarship_name").
            NotEmpty().
            Unique(),
    }
}

// Edges of the Scholarshipinformation.
func (Scholarshipinformation) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("Scholarshiptype", Scholarshiptype.Type).
			Ref("Scholarshipinformation").
			Unique(),
        edge.To("ScholarshipRequest", ScholarshipRequest.Type),
    }
}