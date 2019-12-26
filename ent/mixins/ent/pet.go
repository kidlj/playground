// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
)

// Pet is the model entity for the Pet schema.
type Pet struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Age holds the value of the "age" field.
	Age int `json:"age,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Weight holds the value of the "weight" field.
	Weight float64 `json:"weight,omitempty"`
}

// FromRows scans the sql response data into Pet.
func (pe *Pet) FromRows(rows *sql.Rows) error {
	var scanpe struct {
		ID     int
		Age    sql.NullInt64
		Name   sql.NullString
		Weight sql.NullFloat64
	}
	// the order here should be the same as in the `pet.Columns`.
	if err := rows.Scan(
		&scanpe.ID,
		&scanpe.Age,
		&scanpe.Name,
		&scanpe.Weight,
	); err != nil {
		return err
	}
	pe.ID = scanpe.ID
	pe.Age = int(scanpe.Age.Int64)
	pe.Name = scanpe.Name.String
	pe.Weight = scanpe.Weight.Float64
	return nil
}

// Update returns a builder for updating this Pet.
// Note that, you need to call Pet.Unwrap() before calling this method, if this Pet
// was returned from a transaction, and the transaction was committed or rolled back.
func (pe *Pet) Update() *PetUpdateOne {
	return (&PetClient{pe.config}).UpdateOne(pe)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (pe *Pet) Unwrap() *Pet {
	tx, ok := pe.config.driver.(*txDriver)
	if !ok {
		panic("ent: Pet is not a transactional entity")
	}
	pe.config.driver = tx.drv
	return pe
}

// String implements the fmt.Stringer.
func (pe *Pet) String() string {
	var builder strings.Builder
	builder.WriteString("Pet(")
	builder.WriteString(fmt.Sprintf("id=%v", pe.ID))
	builder.WriteString(", age=")
	builder.WriteString(fmt.Sprintf("%v", pe.Age))
	builder.WriteString(", name=")
	builder.WriteString(pe.Name)
	builder.WriteString(", weight=")
	builder.WriteString(fmt.Sprintf("%v", pe.Weight))
	builder.WriteByte(')')
	return builder.String()
}

// Pets is a parsable slice of Pet.
type Pets []*Pet

// FromRows scans the sql response data into Pets.
func (pe *Pets) FromRows(rows *sql.Rows) error {
	for rows.Next() {
		scanpe := &Pet{}
		if err := scanpe.FromRows(rows); err != nil {
			return err
		}
		*pe = append(*pe, scanpe)
	}
	return nil
}

func (pe Pets) config(cfg config) {
	for _i := range pe {
		pe[_i].config = cfg
	}
}
