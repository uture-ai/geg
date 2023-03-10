// Code generated by ent, DO NOT EDIT.

package etcuser

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"myetc.lantron.ltd/m/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int64) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int64) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int64) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int64) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int64) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int64) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int64) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int64) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int64) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// EtcUsername applies equality check predicate on the "etc_username" field. It's identical to EtcUsernameEQ.
func EtcUsername(v int64) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEtcUsername), v))
	})
}

// EtcPassword applies equality check predicate on the "etc_password" field. It's identical to EtcPasswordEQ.
func EtcPassword(v string) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEtcPassword), v))
	})
}

// PointUsername applies equality check predicate on the "point_username" field. It's identical to PointUsernameEQ.
func PointUsername(v string) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPointUsername), v))
	})
}

// PointPassword applies equality check predicate on the "point_password" field. It's identical to PointPasswordEQ.
func PointPassword(v string) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPointPassword), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// EtcUsernameEQ applies the EQ predicate on the "etc_username" field.
func EtcUsernameEQ(v int64) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEtcUsername), v))
	})
}

// EtcUsernameNEQ applies the NEQ predicate on the "etc_username" field.
func EtcUsernameNEQ(v int64) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEtcUsername), v))
	})
}

// EtcUsernameIn applies the In predicate on the "etc_username" field.
func EtcUsernameIn(vs ...int64) predicate.ETCUser {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEtcUsername), v...))
	})
}

// EtcUsernameNotIn applies the NotIn predicate on the "etc_username" field.
func EtcUsernameNotIn(vs ...int64) predicate.ETCUser {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEtcUsername), v...))
	})
}

// EtcUsernameGT applies the GT predicate on the "etc_username" field.
func EtcUsernameGT(v int64) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEtcUsername), v))
	})
}

// EtcUsernameGTE applies the GTE predicate on the "etc_username" field.
func EtcUsernameGTE(v int64) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEtcUsername), v))
	})
}

// EtcUsernameLT applies the LT predicate on the "etc_username" field.
func EtcUsernameLT(v int64) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEtcUsername), v))
	})
}

// EtcUsernameLTE applies the LTE predicate on the "etc_username" field.
func EtcUsernameLTE(v int64) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEtcUsername), v))
	})
}

// EtcPasswordEQ applies the EQ predicate on the "etc_password" field.
func EtcPasswordEQ(v string) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEtcPassword), v))
	})
}

// EtcPasswordNEQ applies the NEQ predicate on the "etc_password" field.
func EtcPasswordNEQ(v string) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEtcPassword), v))
	})
}

// EtcPasswordIn applies the In predicate on the "etc_password" field.
func EtcPasswordIn(vs ...string) predicate.ETCUser {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldEtcPassword), v...))
	})
}

// EtcPasswordNotIn applies the NotIn predicate on the "etc_password" field.
func EtcPasswordNotIn(vs ...string) predicate.ETCUser {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldEtcPassword), v...))
	})
}

// EtcPasswordGT applies the GT predicate on the "etc_password" field.
func EtcPasswordGT(v string) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEtcPassword), v))
	})
}

// EtcPasswordGTE applies the GTE predicate on the "etc_password" field.
func EtcPasswordGTE(v string) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEtcPassword), v))
	})
}

// EtcPasswordLT applies the LT predicate on the "etc_password" field.
func EtcPasswordLT(v string) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEtcPassword), v))
	})
}

// EtcPasswordLTE applies the LTE predicate on the "etc_password" field.
func EtcPasswordLTE(v string) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEtcPassword), v))
	})
}

// EtcPasswordContains applies the Contains predicate on the "etc_password" field.
func EtcPasswordContains(v string) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldEtcPassword), v))
	})
}

// EtcPasswordHasPrefix applies the HasPrefix predicate on the "etc_password" field.
func EtcPasswordHasPrefix(v string) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldEtcPassword), v))
	})
}

// EtcPasswordHasSuffix applies the HasSuffix predicate on the "etc_password" field.
func EtcPasswordHasSuffix(v string) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldEtcPassword), v))
	})
}

// EtcPasswordEqualFold applies the EqualFold predicate on the "etc_password" field.
func EtcPasswordEqualFold(v string) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldEtcPassword), v))
	})
}

// EtcPasswordContainsFold applies the ContainsFold predicate on the "etc_password" field.
func EtcPasswordContainsFold(v string) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldEtcPassword), v))
	})
}

// PointUsernameEQ applies the EQ predicate on the "point_username" field.
func PointUsernameEQ(v string) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPointUsername), v))
	})
}

// PointUsernameNEQ applies the NEQ predicate on the "point_username" field.
func PointUsernameNEQ(v string) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPointUsername), v))
	})
}

// PointUsernameIn applies the In predicate on the "point_username" field.
func PointUsernameIn(vs ...string) predicate.ETCUser {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPointUsername), v...))
	})
}

// PointUsernameNotIn applies the NotIn predicate on the "point_username" field.
func PointUsernameNotIn(vs ...string) predicate.ETCUser {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPointUsername), v...))
	})
}

// PointUsernameGT applies the GT predicate on the "point_username" field.
func PointUsernameGT(v string) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPointUsername), v))
	})
}

// PointUsernameGTE applies the GTE predicate on the "point_username" field.
func PointUsernameGTE(v string) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPointUsername), v))
	})
}

// PointUsernameLT applies the LT predicate on the "point_username" field.
func PointUsernameLT(v string) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPointUsername), v))
	})
}

// PointUsernameLTE applies the LTE predicate on the "point_username" field.
func PointUsernameLTE(v string) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPointUsername), v))
	})
}

// PointUsernameContains applies the Contains predicate on the "point_username" field.
func PointUsernameContains(v string) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPointUsername), v))
	})
}

// PointUsernameHasPrefix applies the HasPrefix predicate on the "point_username" field.
func PointUsernameHasPrefix(v string) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPointUsername), v))
	})
}

// PointUsernameHasSuffix applies the HasSuffix predicate on the "point_username" field.
func PointUsernameHasSuffix(v string) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPointUsername), v))
	})
}

// PointUsernameEqualFold applies the EqualFold predicate on the "point_username" field.
func PointUsernameEqualFold(v string) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPointUsername), v))
	})
}

// PointUsernameContainsFold applies the ContainsFold predicate on the "point_username" field.
func PointUsernameContainsFold(v string) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPointUsername), v))
	})
}

// PointPasswordEQ applies the EQ predicate on the "point_password" field.
func PointPasswordEQ(v string) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPointPassword), v))
	})
}

// PointPasswordNEQ applies the NEQ predicate on the "point_password" field.
func PointPasswordNEQ(v string) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPointPassword), v))
	})
}

// PointPasswordIn applies the In predicate on the "point_password" field.
func PointPasswordIn(vs ...string) predicate.ETCUser {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldPointPassword), v...))
	})
}

// PointPasswordNotIn applies the NotIn predicate on the "point_password" field.
func PointPasswordNotIn(vs ...string) predicate.ETCUser {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldPointPassword), v...))
	})
}

// PointPasswordGT applies the GT predicate on the "point_password" field.
func PointPasswordGT(v string) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPointPassword), v))
	})
}

// PointPasswordGTE applies the GTE predicate on the "point_password" field.
func PointPasswordGTE(v string) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPointPassword), v))
	})
}

// PointPasswordLT applies the LT predicate on the "point_password" field.
func PointPasswordLT(v string) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPointPassword), v))
	})
}

// PointPasswordLTE applies the LTE predicate on the "point_password" field.
func PointPasswordLTE(v string) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPointPassword), v))
	})
}

// PointPasswordContains applies the Contains predicate on the "point_password" field.
func PointPasswordContains(v string) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPointPassword), v))
	})
}

// PointPasswordHasPrefix applies the HasPrefix predicate on the "point_password" field.
func PointPasswordHasPrefix(v string) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPointPassword), v))
	})
}

// PointPasswordHasSuffix applies the HasSuffix predicate on the "point_password" field.
func PointPasswordHasSuffix(v string) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPointPassword), v))
	})
}

// PointPasswordEqualFold applies the EqualFold predicate on the "point_password" field.
func PointPasswordEqualFold(v string) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPointPassword), v))
	})
}

// PointPasswordContainsFold applies the ContainsFold predicate on the "point_password" field.
func PointPasswordContainsFold(v string) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPointPassword), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.ETCUser {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.ETCUser {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.ETCUser {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.ETCUser {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ETCUser) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ETCUser) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ETCUser) predicate.ETCUser {
	return predicate.ETCUser(func(s *sql.Selector) {
		p(s.Not())
	})
}
