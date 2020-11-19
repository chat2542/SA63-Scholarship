package schema

import (
    "github.com/facebookincubator/ent"
    "github.com/facebookincubator/ent/schema/edge"
    "github.com/facebookincubator/ent/schema/field"
)

// Scholarshiptype holds the schema definition for the Scholarshiptype entity.
type Scholarshiptype struct {
    ent.Schema
}

// Fields of the Scholarshiptype.
func (Scholarshiptype) Fields() []ent.Field {
    return []ent.Field{
        field.String("Type_name").
            NotEmpty().
            Unique(),
    }
}

// Edges of the Scholarshiptype.
func (Scholarshiptype) Edges() []ent.Edge {
    return []ent.Edge{
        edge.To("Scholarshipinformation", Scholarshipinformation.Type),
        edge.To("ScholarshipRequest", ScholarshipRequest.Type),
    }
}
