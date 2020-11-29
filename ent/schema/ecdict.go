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
		field.String("word"),
		field.String("sw"),
		field.String("phonetic"),
		field.String("definition"),
		field.String("translation"),
		field.String("pos"),
		field.Int("collins"),
		field.Int("oxford"),
		field.String("tag"),
		field.Int("bnc"),
		field.Int("frq"),
		field.String("exchange"),
		field.String("detail"),
		field.String("audio"),
	}
}

// Edges of the Ecdict.
func (Ecdict) Edges() []ent.Edge {
	return nil
}
