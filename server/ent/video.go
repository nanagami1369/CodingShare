// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/nanagami1369/CodingShare/ent/user"
	"github.com/nanagami1369/CodingShare/ent/video"
	"github.com/nanagami1369/CodingShare/model"
)

// Video is the model entity for the Video schema.
type Video struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// LanguageTag holds the value of the "language_tag" field.
	LanguageTag *model.Language `json:"language_tag,omitempty"`
	// UploadTime holds the value of the "upload_time" field.
	UploadTime time.Time `json:"upload_time,omitempty"`
	// RecordingTime holds the value of the "recording_time" field.
	RecordingTime int `json:"recording_time,omitempty"`
	// CodingSequence holds the value of the "coding_sequence" field.
	CodingSequence *[]model.CodingSequence `json:"coding_sequence,omitempty"`
	// Comment holds the value of the "comment" field.
	Comment string `json:"comment,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the VideoQuery when eager-loading is set.
	Edges      VideoEdges `json:"edges"`
	video_user *int
}

// VideoEdges holds the relations/edges for other nodes in the graph.
type VideoEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e VideoEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Video) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case video.FieldLanguageTag, video.FieldCodingSequence:
			values[i] = new([]byte)
		case video.FieldID, video.FieldRecordingTime:
			values[i] = new(sql.NullInt64)
		case video.FieldTitle, video.FieldComment:
			values[i] = new(sql.NullString)
		case video.FieldUploadTime:
			values[i] = new(sql.NullTime)
		case video.ForeignKeys[0]: // video_user
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Video", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Video fields.
func (v *Video) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case video.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			v.ID = int(value.Int64)
		case video.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				v.Title = value.String
			}
		case video.FieldLanguageTag:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field language_tag", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &v.LanguageTag); err != nil {
					return fmt.Errorf("unmarshal field language_tag: %w", err)
				}
			}
		case video.FieldUploadTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field upload_time", values[i])
			} else if value.Valid {
				v.UploadTime = value.Time
			}
		case video.FieldRecordingTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field recording_time", values[i])
			} else if value.Valid {
				v.RecordingTime = int(value.Int64)
			}
		case video.FieldCodingSequence:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field coding_sequence", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &v.CodingSequence); err != nil {
					return fmt.Errorf("unmarshal field coding_sequence: %w", err)
				}
			}
		case video.FieldComment:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field comment", values[i])
			} else if value.Valid {
				v.Comment = value.String
			}
		case video.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field video_user", value)
			} else if value.Valid {
				v.video_user = new(int)
				*v.video_user = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the Video entity.
func (v *Video) QueryUser() *UserQuery {
	return (&VideoClient{config: v.config}).QueryUser(v)
}

// Update returns a builder for updating this Video.
// Note that you need to call Video.Unwrap() before calling this method if this Video
// was returned from a transaction, and the transaction was committed or rolled back.
func (v *Video) Update() *VideoUpdateOne {
	return (&VideoClient{config: v.config}).UpdateOne(v)
}

// Unwrap unwraps the Video entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (v *Video) Unwrap() *Video {
	tx, ok := v.config.driver.(*txDriver)
	if !ok {
		panic("ent: Video is not a transactional entity")
	}
	v.config.driver = tx.drv
	return v
}

// String implements the fmt.Stringer.
func (v *Video) String() string {
	var builder strings.Builder
	builder.WriteString("Video(")
	builder.WriteString(fmt.Sprintf("id=%v", v.ID))
	builder.WriteString(", title=")
	builder.WriteString(v.Title)
	builder.WriteString(", language_tag=")
	builder.WriteString(fmt.Sprintf("%v", v.LanguageTag))
	builder.WriteString(", upload_time=")
	builder.WriteString(v.UploadTime.Format(time.ANSIC))
	builder.WriteString(", recording_time=")
	builder.WriteString(fmt.Sprintf("%v", v.RecordingTime))
	builder.WriteString(", coding_sequence=")
	builder.WriteString(fmt.Sprintf("%v", v.CodingSequence))
	builder.WriteString(", comment=")
	builder.WriteString(v.Comment)
	builder.WriteByte(')')
	return builder.String()
}

// Videos is a parsable slice of Video.
type Videos []*Video

func (v Videos) config(cfg config) {
	for _i := range v {
		v[_i].config = cfg
	}
}