// Code generated by entc, DO NOT EDIT.

package ent

import (
	"ECDICT-GO/ent/ecdict"
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
)

// Ecdict is the model entity for the Ecdict schema.
type Ecdict struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Word holds the value of the "word" field.
	Word string `json:"word,omitempty"`
	// Sw holds the value of the "sw" field.
	Sw string `json:"sw,omitempty"`
	// Phonetic holds the value of the "phonetic" field.
	Phonetic string `json:"phonetic,omitempty"`
	// Definition holds the value of the "definition" field.
	Definition string `json:"definition,omitempty"`
	// Translation holds the value of the "translation" field.
	Translation string `json:"translation,omitempty"`
	// Pos holds the value of the "pos" field.
	Pos string `json:"pos,omitempty"`
	// Collins holds the value of the "collins" field.
	Collins int `json:"collins,omitempty"`
	// Oxford holds the value of the "oxford" field.
	Oxford int `json:"oxford,omitempty"`
	// Tag holds the value of the "tag" field.
	Tag string `json:"tag,omitempty"`
	// Bnc holds the value of the "bnc" field.
	Bnc int `json:"bnc,omitempty"`
	// Frq holds the value of the "frq" field.
	Frq int `json:"frq,omitempty"`
	// Exchange holds the value of the "exchange" field.
	Exchange string `json:"exchange,omitempty"`
	// Detail holds the value of the "detail" field.
	Detail string `json:"detail,omitempty"`
	// Audio holds the value of the "audio" field.
	Audio string `json:"audio,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Ecdict) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // word
		&sql.NullString{}, // sw
		&sql.NullString{}, // phonetic
		&sql.NullString{}, // definition
		&sql.NullString{}, // translation
		&sql.NullString{}, // pos
		&sql.NullInt64{},  // collins
		&sql.NullInt64{},  // oxford
		&sql.NullString{}, // tag
		&sql.NullInt64{},  // bnc
		&sql.NullInt64{},  // frq
		&sql.NullString{}, // exchange
		&sql.NullString{}, // detail
		&sql.NullString{}, // audio
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Ecdict fields.
func (e *Ecdict) assignValues(values ...interface{}) error {
	if m, n := len(values), len(ecdict.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	e.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field word", values[0])
	} else if value.Valid {
		e.Word = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field sw", values[1])
	} else if value.Valid {
		e.Sw = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field phonetic", values[2])
	} else if value.Valid {
		e.Phonetic = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field definition", values[3])
	} else if value.Valid {
		e.Definition = value.String
	}
	if value, ok := values[4].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field translation", values[4])
	} else if value.Valid {
		e.Translation = value.String
	}
	if value, ok := values[5].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field pos", values[5])
	} else if value.Valid {
		e.Pos = value.String
	}
	if value, ok := values[6].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field collins", values[6])
	} else if value.Valid {
		e.Collins = int(value.Int64)
	}
	if value, ok := values[7].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field oxford", values[7])
	} else if value.Valid {
		e.Oxford = int(value.Int64)
	}
	if value, ok := values[8].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field tag", values[8])
	} else if value.Valid {
		e.Tag = value.String
	}
	if value, ok := values[9].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field bnc", values[9])
	} else if value.Valid {
		e.Bnc = int(value.Int64)
	}
	if value, ok := values[10].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field frq", values[10])
	} else if value.Valid {
		e.Frq = int(value.Int64)
	}
	if value, ok := values[11].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field exchange", values[11])
	} else if value.Valid {
		e.Exchange = value.String
	}
	if value, ok := values[12].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field detail", values[12])
	} else if value.Valid {
		e.Detail = value.String
	}
	if value, ok := values[13].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field audio", values[13])
	} else if value.Valid {
		e.Audio = value.String
	}
	return nil
}

// Update returns a builder for updating this Ecdict.
// Note that, you need to call Ecdict.Unwrap() before calling this method, if this Ecdict
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Ecdict) Update() *EcdictUpdateOne {
	return (&EcdictClient{config: e.config}).UpdateOne(e)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (e *Ecdict) Unwrap() *Ecdict {
	tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("ent: Ecdict is not a transactional entity")
	}
	e.config.driver = tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Ecdict) String() string {
	var builder strings.Builder
	builder.WriteString("Ecdict(")
	builder.WriteString(fmt.Sprintf("id=%v", e.ID))
	builder.WriteString(", word=")
	builder.WriteString(e.Word)
	builder.WriteString(", sw=")
	builder.WriteString(e.Sw)
	builder.WriteString(", phonetic=")
	builder.WriteString(e.Phonetic)
	builder.WriteString(", definition=")
	builder.WriteString(e.Definition)
	builder.WriteString(", translation=")
	builder.WriteString(e.Translation)
	builder.WriteString(", pos=")
	builder.WriteString(e.Pos)
	builder.WriteString(", collins=")
	builder.WriteString(fmt.Sprintf("%v", e.Collins))
	builder.WriteString(", oxford=")
	builder.WriteString(fmt.Sprintf("%v", e.Oxford))
	builder.WriteString(", tag=")
	builder.WriteString(e.Tag)
	builder.WriteString(", bnc=")
	builder.WriteString(fmt.Sprintf("%v", e.Bnc))
	builder.WriteString(", frq=")
	builder.WriteString(fmt.Sprintf("%v", e.Frq))
	builder.WriteString(", exchange=")
	builder.WriteString(e.Exchange)
	builder.WriteString(", detail=")
	builder.WriteString(e.Detail)
	builder.WriteString(", audio=")
	builder.WriteString(e.Audio)
	builder.WriteByte(')')
	return builder.String()
}

// Ecdicts is a parsable slice of Ecdict.
type Ecdicts []*Ecdict

func (e Ecdicts) config(cfg config) {
	for _i := range e {
		e[_i].config = cfg
	}
}