package models

import (
	bson "gopkg.in/mgo.v2/bson"
	"io"
	"strconv"

	graphql "github.com/99designs/gqlgen/graphql"

)

// MarshalID redefine the base ID type to use an id from an external library
func MarshalID(id bson.ObjectId) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(id.Hex()))
	})
}

// UnmarshalID returns primitive.ObjectID and error
func UnmarshalID(v interface{}) (bson.ObjectId, error) {
	str, ok := v.(string)
	if !ok {
		return bson.ObjectIdHex(""), nil
	}
	return bson.ObjectIdHex(str), nil
}
