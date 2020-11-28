package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

// Ecdict holds the schema definition for the Ecdict entity.
type Ecdict struct {
	ent.Schema
}

// Fields of the Ecdict.
func (Ecdict) Fields() []ent.Field {
	return []ent.Field{
		field.String("word").Positive(),
		field.String(""),
		field.String(""),
		field.String(""),
		field.String(""),
		field.String(""),
		field.String(""),
		field.String(""),
		field.String(""),
		field.String(""),
		field.String(""),
		field.String(""),
		field.String(""),
	}
}

// Edges of the Ecdict.
func (Ecdict) Edges() []ent.Edge {
	return nil
}
