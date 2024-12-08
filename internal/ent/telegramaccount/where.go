// Code generated by ent, DO NOT EDIT.

package telegramaccount

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/gotd/bot/internal/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldContainsFold(FieldID, id))
}

// Code applies equality check predicate on the "code" field. It's identical to CodeEQ.
func Code(v string) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldEQ(FieldCode, v))
}

// CodeAt applies equality check predicate on the "code_at" field. It's identical to CodeAtEQ.
func CodeAt(v time.Time) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldEQ(FieldCodeAt, v))
}

// Data applies equality check predicate on the "data" field. It's identical to DataEQ.
func Data(v []byte) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldEQ(FieldData, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldEQ(FieldStatus, v))
}

// CodeEQ applies the EQ predicate on the "code" field.
func CodeEQ(v string) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldEQ(FieldCode, v))
}

// CodeNEQ applies the NEQ predicate on the "code" field.
func CodeNEQ(v string) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldNEQ(FieldCode, v))
}

// CodeIn applies the In predicate on the "code" field.
func CodeIn(vs ...string) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldIn(FieldCode, vs...))
}

// CodeNotIn applies the NotIn predicate on the "code" field.
func CodeNotIn(vs ...string) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldNotIn(FieldCode, vs...))
}

// CodeGT applies the GT predicate on the "code" field.
func CodeGT(v string) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldGT(FieldCode, v))
}

// CodeGTE applies the GTE predicate on the "code" field.
func CodeGTE(v string) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldGTE(FieldCode, v))
}

// CodeLT applies the LT predicate on the "code" field.
func CodeLT(v string) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldLT(FieldCode, v))
}

// CodeLTE applies the LTE predicate on the "code" field.
func CodeLTE(v string) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldLTE(FieldCode, v))
}

// CodeContains applies the Contains predicate on the "code" field.
func CodeContains(v string) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldContains(FieldCode, v))
}

// CodeHasPrefix applies the HasPrefix predicate on the "code" field.
func CodeHasPrefix(v string) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldHasPrefix(FieldCode, v))
}

// CodeHasSuffix applies the HasSuffix predicate on the "code" field.
func CodeHasSuffix(v string) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldHasSuffix(FieldCode, v))
}

// CodeEqualFold applies the EqualFold predicate on the "code" field.
func CodeEqualFold(v string) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldEqualFold(FieldCode, v))
}

// CodeContainsFold applies the ContainsFold predicate on the "code" field.
func CodeContainsFold(v string) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldContainsFold(FieldCode, v))
}

// CodeAtEQ applies the EQ predicate on the "code_at" field.
func CodeAtEQ(v time.Time) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldEQ(FieldCodeAt, v))
}

// CodeAtNEQ applies the NEQ predicate on the "code_at" field.
func CodeAtNEQ(v time.Time) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldNEQ(FieldCodeAt, v))
}

// CodeAtIn applies the In predicate on the "code_at" field.
func CodeAtIn(vs ...time.Time) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldIn(FieldCodeAt, vs...))
}

// CodeAtNotIn applies the NotIn predicate on the "code_at" field.
func CodeAtNotIn(vs ...time.Time) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldNotIn(FieldCodeAt, vs...))
}

// CodeAtGT applies the GT predicate on the "code_at" field.
func CodeAtGT(v time.Time) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldGT(FieldCodeAt, v))
}

// CodeAtGTE applies the GTE predicate on the "code_at" field.
func CodeAtGTE(v time.Time) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldGTE(FieldCodeAt, v))
}

// CodeAtLT applies the LT predicate on the "code_at" field.
func CodeAtLT(v time.Time) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldLT(FieldCodeAt, v))
}

// CodeAtLTE applies the LTE predicate on the "code_at" field.
func CodeAtLTE(v time.Time) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldLTE(FieldCodeAt, v))
}

// DataEQ applies the EQ predicate on the "data" field.
func DataEQ(v []byte) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldEQ(FieldData, v))
}

// DataNEQ applies the NEQ predicate on the "data" field.
func DataNEQ(v []byte) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldNEQ(FieldData, v))
}

// DataIn applies the In predicate on the "data" field.
func DataIn(vs ...[]byte) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldIn(FieldData, vs...))
}

// DataNotIn applies the NotIn predicate on the "data" field.
func DataNotIn(vs ...[]byte) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldNotIn(FieldData, vs...))
}

// DataGT applies the GT predicate on the "data" field.
func DataGT(v []byte) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldGT(FieldData, v))
}

// DataGTE applies the GTE predicate on the "data" field.
func DataGTE(v []byte) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldGTE(FieldData, v))
}

// DataLT applies the LT predicate on the "data" field.
func DataLT(v []byte) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldLT(FieldData, v))
}

// DataLTE applies the LTE predicate on the "data" field.
func DataLTE(v []byte) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldLTE(FieldData, v))
}

// StateEQ applies the EQ predicate on the "state" field.
func StateEQ(v State) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldEQ(FieldState, v))
}

// StateNEQ applies the NEQ predicate on the "state" field.
func StateNEQ(v State) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldNEQ(FieldState, v))
}

// StateIn applies the In predicate on the "state" field.
func StateIn(vs ...State) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldIn(FieldState, vs...))
}

// StateNotIn applies the NotIn predicate on the "state" field.
func StateNotIn(vs ...State) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldNotIn(FieldState, vs...))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldLTE(FieldStatus, v))
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldContains(FieldStatus, v))
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldHasPrefix(FieldStatus, v))
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldHasSuffix(FieldStatus, v))
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldEqualFold(FieldStatus, v))
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.FieldContainsFold(FieldStatus, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TelegramAccount) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TelegramAccount) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.TelegramAccount) predicate.TelegramAccount {
	return predicate.TelegramAccount(sql.NotPredicates(p))
}