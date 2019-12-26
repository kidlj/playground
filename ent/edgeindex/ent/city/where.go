// Code generated by entc, DO NOT EDIT.

package city

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/kidlj/demo/ent/edgeindex/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.City {
	return predicate.City(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldID), id))
		},
	)
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.City {
	return predicate.City(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldID), id))
		},
	)
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.City {
	return predicate.City(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldID), id))
		},
	)
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.City {
	return predicate.City(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(ids) == 0 {
				s.Where(sql.False())
				return
			}
			v := make([]interface{}, len(ids))
			for i := range v {
				v[i] = ids[i]
			}
			s.Where(sql.In(s.C(FieldID), v...))
		},
	)
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.City {
	return predicate.City(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(ids) == 0 {
				s.Where(sql.False())
				return
			}
			v := make([]interface{}, len(ids))
			for i := range v {
				v[i] = ids[i]
			}
			s.Where(sql.NotIn(s.C(FieldID), v...))
		},
	)
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.City {
	return predicate.City(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldID), id))
		},
	)
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.City {
	return predicate.City(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldID), id))
		},
	)
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.City {
	return predicate.City(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldID), id))
		},
	)
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.City {
	return predicate.City(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldID), id))
		},
	)
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.City {
	return predicate.City(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldName), v))
		},
	)
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.City {
	return predicate.City(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldName), v))
		},
	)
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.City {
	return predicate.City(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldName), v))
		},
	)
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.City {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.City(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldName), v...))
		},
	)
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.City {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.City(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldName), v...))
		},
	)
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.City {
	return predicate.City(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldName), v))
		},
	)
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.City {
	return predicate.City(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldName), v))
		},
	)
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.City {
	return predicate.City(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldName), v))
		},
	)
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.City {
	return predicate.City(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldName), v))
		},
	)
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.City {
	return predicate.City(
		func(s *sql.Selector) {
			s.Where(sql.Contains(s.C(FieldName), v))
		},
	)
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.City {
	return predicate.City(
		func(s *sql.Selector) {
			s.Where(sql.HasPrefix(s.C(FieldName), v))
		},
	)
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.City {
	return predicate.City(
		func(s *sql.Selector) {
			s.Where(sql.HasSuffix(s.C(FieldName), v))
		},
	)
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.City {
	return predicate.City(
		func(s *sql.Selector) {
			s.Where(sql.EqualFold(s.C(FieldName), v))
		},
	)
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.City {
	return predicate.City(
		func(s *sql.Selector) {
			s.Where(sql.ContainsFold(s.C(FieldName), v))
		},
	)
}

// HasStreets applies the HasEdge predicate on the "streets" edge.
func HasStreets() predicate.City {
	return predicate.City(
		func(s *sql.Selector) {
			step := sql.NewStep(
				sql.From(Table, FieldID),
				sql.To(StreetsTable, FieldID),
				sql.Edge(sql.O2M, false, StreetsTable, StreetsColumn),
			)
			sql.HasNeighbors(s, step)
		},
	)
}

// HasStreetsWith applies the HasEdge predicate on the "streets" edge with a given conditions (other predicates).
func HasStreetsWith(preds ...predicate.Street) predicate.City {
	return predicate.City(
		func(s *sql.Selector) {
			step := sql.NewStep(
				sql.From(Table, FieldID),
				sql.To(StreetsInverseTable, FieldID),
				sql.Edge(sql.O2M, false, StreetsTable, StreetsColumn),
			)
			sql.HasNeighborsWith(s, step, func(s *sql.Selector) {
				for _, p := range preds {
					p(s)
				}
			})
		},
	)
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.City) predicate.City {
	return predicate.City(
		func(s *sql.Selector) {
			s1 := s.Clone().SetP(nil)
			for _, p := range predicates {
				p(s1)
			}
			s.Where(s1.P())
		},
	)
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.City) predicate.City {
	return predicate.City(
		func(s *sql.Selector) {
			s1 := s.Clone().SetP(nil)
			for i, p := range predicates {
				if i > 0 {
					s1.Or()
				}
				p(s1)
			}
			s.Where(s1.P())
		},
	)
}

// Not applies the not operator on the given predicate.
func Not(p predicate.City) predicate.City {
	return predicate.City(
		func(s *sql.Selector) {
			p(s.Not())
		},
	)
}
