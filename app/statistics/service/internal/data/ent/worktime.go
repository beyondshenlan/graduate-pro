// Code generated by entc, DO NOT EDIT.

package ent

import (
	"clock-in/app/worktime/service/internal/data/ent/worktime"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Worktime is the model entity for the Worktime schema.
type Worktime struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// User holds the value of the "user" field.
	User int64 `json:"user,omitempty"`
	// Day holds the value of the "day" field.
	Day int64 `json:"day,omitempty"`
	// Minute holds the value of the "minute" field.
	Minute int64 `json:"minute,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Worktime) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case worktime.FieldID, worktime.FieldUser, worktime.FieldDay, worktime.FieldMinute:
			values[i] = new(sql.NullInt64)
		case worktime.FieldCreatedAt, worktime.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Worktime", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Worktime fields.
func (w *Worktime) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case worktime.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			w.ID = int(value.Int64)
		case worktime.FieldUser:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user", values[i])
			} else if value.Valid {
				w.User = value.Int64
			}
		case worktime.FieldDay:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field day", values[i])
			} else if value.Valid {
				w.Day = value.Int64
			}
		case worktime.FieldMinute:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field minute", values[i])
			} else if value.Valid {
				w.Minute = value.Int64
			}
		case worktime.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				w.CreatedAt = value.Time
			}
		case worktime.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				w.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Worktime.
// Note that you need to call Worktime.Unwrap() before calling this method if this Worktime
// was returned from a transaction, and the transaction was committed or rolled back.
func (w *Worktime) Update() *WorktimeUpdateOne {
	return (&WorktimeClient{config: w.config}).UpdateOne(w)
}

// Unwrap unwraps the Worktime entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (w *Worktime) Unwrap() *Worktime {
	tx, ok := w.config.driver.(*txDriver)
	if !ok {
		panic("ent: Worktime is not a transactional entity")
	}
	w.config.driver = tx.drv
	return w
}

// String implements the fmt.Stringer.
func (w *Worktime) String() string {
	var builder strings.Builder
	builder.WriteString("Worktime(")
	builder.WriteString(fmt.Sprintf("id=%v", w.ID))
	builder.WriteString(", user=")
	builder.WriteString(fmt.Sprintf("%v", w.User))
	builder.WriteString(", day=")
	builder.WriteString(fmt.Sprintf("%v", w.Day))
	builder.WriteString(", minute=")
	builder.WriteString(fmt.Sprintf("%v", w.Minute))
	builder.WriteString(", created_at=")
	builder.WriteString(w.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(w.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Worktimes is a parsable slice of Worktime.
type Worktimes []*Worktime

func (w Worktimes) config(cfg config) {
	for _i := range w {
		w[_i].config = cfg
	}
}
