// Code generated by entc, DO NOT EDIT.

package notes

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/ArifulProtik/BlackPen/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
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
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
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
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// Body applies equality check predicate on the "body" field. It's identical to BodyEQ.
func Body(v string) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBody), v))
	})
}

// Slug applies equality check predicate on the "slug" field. It's identical to SlugEQ.
func Slug(v string) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSlug), v))
	})
}

// FImage applies equality check predicate on the "f_image" field. It's identical to FImageEQ.
func FImage(v string) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFImage), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTitle), v))
	})
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Notes {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Notes(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTitle), v...))
	})
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Notes {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Notes(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTitle), v...))
	})
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTitle), v))
	})
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTitle), v))
	})
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTitle), v))
	})
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTitle), v))
	})
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTitle), v))
	})
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTitle), v))
	})
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTitle), v))
	})
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTitle), v))
	})
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTitle), v))
	})
}

// BodyEQ applies the EQ predicate on the "body" field.
func BodyEQ(v string) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBody), v))
	})
}

// BodyNEQ applies the NEQ predicate on the "body" field.
func BodyNEQ(v string) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBody), v))
	})
}

// BodyIn applies the In predicate on the "body" field.
func BodyIn(vs ...string) predicate.Notes {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Notes(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBody), v...))
	})
}

// BodyNotIn applies the NotIn predicate on the "body" field.
func BodyNotIn(vs ...string) predicate.Notes {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Notes(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBody), v...))
	})
}

// BodyGT applies the GT predicate on the "body" field.
func BodyGT(v string) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBody), v))
	})
}

// BodyGTE applies the GTE predicate on the "body" field.
func BodyGTE(v string) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBody), v))
	})
}

// BodyLT applies the LT predicate on the "body" field.
func BodyLT(v string) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBody), v))
	})
}

// BodyLTE applies the LTE predicate on the "body" field.
func BodyLTE(v string) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBody), v))
	})
}

// BodyContains applies the Contains predicate on the "body" field.
func BodyContains(v string) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBody), v))
	})
}

// BodyHasPrefix applies the HasPrefix predicate on the "body" field.
func BodyHasPrefix(v string) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBody), v))
	})
}

// BodyHasSuffix applies the HasSuffix predicate on the "body" field.
func BodyHasSuffix(v string) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBody), v))
	})
}

// BodyEqualFold applies the EqualFold predicate on the "body" field.
func BodyEqualFold(v string) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBody), v))
	})
}

// BodyContainsFold applies the ContainsFold predicate on the "body" field.
func BodyContainsFold(v string) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBody), v))
	})
}

// SlugEQ applies the EQ predicate on the "slug" field.
func SlugEQ(v string) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSlug), v))
	})
}

// SlugNEQ applies the NEQ predicate on the "slug" field.
func SlugNEQ(v string) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSlug), v))
	})
}

// SlugIn applies the In predicate on the "slug" field.
func SlugIn(vs ...string) predicate.Notes {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Notes(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSlug), v...))
	})
}

// SlugNotIn applies the NotIn predicate on the "slug" field.
func SlugNotIn(vs ...string) predicate.Notes {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Notes(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSlug), v...))
	})
}

// SlugGT applies the GT predicate on the "slug" field.
func SlugGT(v string) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSlug), v))
	})
}

// SlugGTE applies the GTE predicate on the "slug" field.
func SlugGTE(v string) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSlug), v))
	})
}

// SlugLT applies the LT predicate on the "slug" field.
func SlugLT(v string) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSlug), v))
	})
}

// SlugLTE applies the LTE predicate on the "slug" field.
func SlugLTE(v string) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSlug), v))
	})
}

// SlugContains applies the Contains predicate on the "slug" field.
func SlugContains(v string) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSlug), v))
	})
}

// SlugHasPrefix applies the HasPrefix predicate on the "slug" field.
func SlugHasPrefix(v string) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSlug), v))
	})
}

// SlugHasSuffix applies the HasSuffix predicate on the "slug" field.
func SlugHasSuffix(v string) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSlug), v))
	})
}

// SlugEqualFold applies the EqualFold predicate on the "slug" field.
func SlugEqualFold(v string) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSlug), v))
	})
}

// SlugContainsFold applies the ContainsFold predicate on the "slug" field.
func SlugContainsFold(v string) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSlug), v))
	})
}

// FImageEQ applies the EQ predicate on the "f_image" field.
func FImageEQ(v string) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFImage), v))
	})
}

// FImageNEQ applies the NEQ predicate on the "f_image" field.
func FImageNEQ(v string) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFImage), v))
	})
}

// FImageIn applies the In predicate on the "f_image" field.
func FImageIn(vs ...string) predicate.Notes {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Notes(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldFImage), v...))
	})
}

// FImageNotIn applies the NotIn predicate on the "f_image" field.
func FImageNotIn(vs ...string) predicate.Notes {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Notes(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldFImage), v...))
	})
}

// FImageGT applies the GT predicate on the "f_image" field.
func FImageGT(v string) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFImage), v))
	})
}

// FImageGTE applies the GTE predicate on the "f_image" field.
func FImageGTE(v string) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFImage), v))
	})
}

// FImageLT applies the LT predicate on the "f_image" field.
func FImageLT(v string) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFImage), v))
	})
}

// FImageLTE applies the LTE predicate on the "f_image" field.
func FImageLTE(v string) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFImage), v))
	})
}

// FImageContains applies the Contains predicate on the "f_image" field.
func FImageContains(v string) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldFImage), v))
	})
}

// FImageHasPrefix applies the HasPrefix predicate on the "f_image" field.
func FImageHasPrefix(v string) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldFImage), v))
	})
}

// FImageHasSuffix applies the HasSuffix predicate on the "f_image" field.
func FImageHasSuffix(v string) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldFImage), v))
	})
}

// FImageEqualFold applies the EqualFold predicate on the "f_image" field.
func FImageEqualFold(v string) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldFImage), v))
	})
}

// FImageContainsFold applies the ContainsFold predicate on the "f_image" field.
func FImageContainsFold(v string) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldFImage), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Notes {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Notes(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Notes {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Notes(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Notes {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Notes(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Notes {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Notes(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUpdatedAt)))
	})
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUpdatedAt)))
	})
}

// HasAuthor applies the HasEdge predicate on the "author" edge.
func HasAuthor() predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AuthorTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AuthorTable, AuthorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAuthorWith applies the HasEdge predicate on the "author" edge with a given conditions (other predicates).
func HasAuthorWith(preds ...predicate.User) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AuthorInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AuthorTable, AuthorColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Notes) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Notes) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
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
func Not(p predicate.Notes) predicate.Notes {
	return predicate.Notes(func(s *sql.Selector) {
		p(s.Not())
	})
}