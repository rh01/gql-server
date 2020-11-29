package models

import (
	"errors"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

// Function for converting a time-object to an RFC3339-String with GraphQL.
// Returns the corresponding marshaller to perform this task.
func MarshalTimestamp(t time.Time) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		//_, _ = io.WriteString(w, t.Format(strconv.Quote(time.RFC3339)))
		_, _ = io.WriteString(w, t.Format(strconv.Quote("2006-01-02 15:04:05")))
		//_, _ = io.WriteString(w, t.Format(strconv.Quote(time.RFC3339Nano)))

	})
}

// Function for converting a RFC3339 Time-String into an time-object. Used by GraphQL.
// Returns a Time-Object representing the Time-String.
func UnmarshalTimestamp(v interface{}) (time.Time, error) {
	if timeString, ok := v.(string); ok {
		parsedTime, err := time.Parse(time.RFC3339, timeString)
		if err != nil {
			return time.Time{}, err
		}
		return parsedTime, nil
	}
	return time.Time{}, errors.New("time should be a string")
}
