// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"myetc.lantron.ltd/m/ent/etcrecord"
)

// ETCRecord is the model entity for the ETCRecord schema.
type ETCRecord struct {
	config `json:"-"`
	// ID of the ent.
	ID int64 `json:"id,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// Entry holds the value of the "entry" field.
	Entry string `json:"entry,omitempty"`
	// EntryYear holds the value of the "entry_year" field.
	EntryYear int32 `json:"entry_year,omitempty"`
	// EntryMonth holds the value of the "entry_month" field.
	EntryMonth int32 `json:"entry_month,omitempty"`
	// EntryDay holds the value of the "entry_day" field.
	EntryDay int32 `json:"entry_day,omitempty"`
	// Exit holds the value of the "exit" field.
	Exit string `json:"exit,omitempty"`
	// ExitDate holds the value of the "exit_date" field.
	ExitDate string `json:"exit_date,omitempty"`
	// ExitTime holds the value of the "exit_time" field.
	ExitTime string `json:"exit_time,omitempty"`
	// TotalPrice holds the value of the "total_price" field.
	TotalPrice int32 `json:"total_price,omitempty"`
	// DiscountPrice holds the value of the "discount_price" field.
	DiscountPrice int32 `json:"discount_price,omitempty"`
	// PaidPrice holds the value of the "paid_price" field.
	PaidPrice int32 `json:"paid_price,omitempty"`
	// CarType holds the value of the "car_type" field.
	CarType int8 `json:"car_type,omitempty"`
	// CarNumber holds the value of the "car_number" field.
	CarNumber string `json:"car_number,omitempty"`
	// CardNumber holds the value of the "card_number" field.
	CardNumber string `json:"card_number,omitempty"`
	// Status holds the value of the "status" field.
	Status string `json:"status,omitempty"`
	// Comment holds the value of the "comment" field.
	Comment *string `json:"comment,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ETCRecord) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case etcrecord.FieldID, etcrecord.FieldEntryYear, etcrecord.FieldEntryMonth, etcrecord.FieldEntryDay, etcrecord.FieldTotalPrice, etcrecord.FieldDiscountPrice, etcrecord.FieldPaidPrice, etcrecord.FieldCarType:
			values[i] = new(sql.NullInt64)
		case etcrecord.FieldUsername, etcrecord.FieldEntry, etcrecord.FieldExit, etcrecord.FieldExitDate, etcrecord.FieldExitTime, etcrecord.FieldCarNumber, etcrecord.FieldCardNumber, etcrecord.FieldStatus, etcrecord.FieldComment:
			values[i] = new(sql.NullString)
		case etcrecord.FieldCreatedAt, etcrecord.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ETCRecord", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ETCRecord fields.
func (er *ETCRecord) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case etcrecord.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			er.ID = int64(value.Int64)
		case etcrecord.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				er.Username = value.String
			}
		case etcrecord.FieldEntry:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field entry", values[i])
			} else if value.Valid {
				er.Entry = value.String
			}
		case etcrecord.FieldEntryYear:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field entry_year", values[i])
			} else if value.Valid {
				er.EntryYear = int32(value.Int64)
			}
		case etcrecord.FieldEntryMonth:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field entry_month", values[i])
			} else if value.Valid {
				er.EntryMonth = int32(value.Int64)
			}
		case etcrecord.FieldEntryDay:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field entry_day", values[i])
			} else if value.Valid {
				er.EntryDay = int32(value.Int64)
			}
		case etcrecord.FieldExit:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field exit", values[i])
			} else if value.Valid {
				er.Exit = value.String
			}
		case etcrecord.FieldExitDate:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field exit_date", values[i])
			} else if value.Valid {
				er.ExitDate = value.String
			}
		case etcrecord.FieldExitTime:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field exit_time", values[i])
			} else if value.Valid {
				er.ExitTime = value.String
			}
		case etcrecord.FieldTotalPrice:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field total_price", values[i])
			} else if value.Valid {
				er.TotalPrice = int32(value.Int64)
			}
		case etcrecord.FieldDiscountPrice:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field discount_price", values[i])
			} else if value.Valid {
				er.DiscountPrice = int32(value.Int64)
			}
		case etcrecord.FieldPaidPrice:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field paid_price", values[i])
			} else if value.Valid {
				er.PaidPrice = int32(value.Int64)
			}
		case etcrecord.FieldCarType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field car_type", values[i])
			} else if value.Valid {
				er.CarType = int8(value.Int64)
			}
		case etcrecord.FieldCarNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field car_number", values[i])
			} else if value.Valid {
				er.CarNumber = value.String
			}
		case etcrecord.FieldCardNumber:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field card_number", values[i])
			} else if value.Valid {
				er.CardNumber = value.String
			}
		case etcrecord.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				er.Status = value.String
			}
		case etcrecord.FieldComment:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field comment", values[i])
			} else if value.Valid {
				er.Comment = new(string)
				*er.Comment = value.String
			}
		case etcrecord.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				er.CreatedAt = value.Time
			}
		case etcrecord.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				er.UpdatedAt = value.Time
			}
		}
	}
	return nil
}

// Update returns a builder for updating this ETCRecord.
// Note that you need to call ETCRecord.Unwrap() before calling this method if this ETCRecord
// was returned from a transaction, and the transaction was committed or rolled back.
func (er *ETCRecord) Update() *ETCRecordUpdateOne {
	return (&ETCRecordClient{config: er.config}).UpdateOne(er)
}

// Unwrap unwraps the ETCRecord entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (er *ETCRecord) Unwrap() *ETCRecord {
	_tx, ok := er.config.driver.(*txDriver)
	if !ok {
		panic("ent: ETCRecord is not a transactional entity")
	}
	er.config.driver = _tx.drv
	return er
}

// String implements the fmt.Stringer.
func (er *ETCRecord) String() string {
	var builder strings.Builder
	builder.WriteString("ETCRecord(")
	builder.WriteString(fmt.Sprintf("id=%v, ", er.ID))
	builder.WriteString("username=")
	builder.WriteString(er.Username)
	builder.WriteString(", ")
	builder.WriteString("entry=")
	builder.WriteString(er.Entry)
	builder.WriteString(", ")
	builder.WriteString("entry_year=")
	builder.WriteString(fmt.Sprintf("%v", er.EntryYear))
	builder.WriteString(", ")
	builder.WriteString("entry_month=")
	builder.WriteString(fmt.Sprintf("%v", er.EntryMonth))
	builder.WriteString(", ")
	builder.WriteString("entry_day=")
	builder.WriteString(fmt.Sprintf("%v", er.EntryDay))
	builder.WriteString(", ")
	builder.WriteString("exit=")
	builder.WriteString(er.Exit)
	builder.WriteString(", ")
	builder.WriteString("exit_date=")
	builder.WriteString(er.ExitDate)
	builder.WriteString(", ")
	builder.WriteString("exit_time=")
	builder.WriteString(er.ExitTime)
	builder.WriteString(", ")
	builder.WriteString("total_price=")
	builder.WriteString(fmt.Sprintf("%v", er.TotalPrice))
	builder.WriteString(", ")
	builder.WriteString("discount_price=")
	builder.WriteString(fmt.Sprintf("%v", er.DiscountPrice))
	builder.WriteString(", ")
	builder.WriteString("paid_price=")
	builder.WriteString(fmt.Sprintf("%v", er.PaidPrice))
	builder.WriteString(", ")
	builder.WriteString("car_type=")
	builder.WriteString(fmt.Sprintf("%v", er.CarType))
	builder.WriteString(", ")
	builder.WriteString("car_number=")
	builder.WriteString(er.CarNumber)
	builder.WriteString(", ")
	builder.WriteString("card_number=")
	builder.WriteString(er.CardNumber)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(er.Status)
	builder.WriteString(", ")
	if v := er.Comment; v != nil {
		builder.WriteString("comment=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(er.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(er.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// ETCRecords is a parsable slice of ETCRecord.
type ETCRecords []*ETCRecord

func (er ETCRecords) config(cfg config) {
	for _i := range er {
		er[_i].config = cfg
	}
}
