package pq

import (
	"database/sql"

	"github.com/gopsql/standard"
	"github.com/lib/pq"
)

var Array = pq.Array

type (
	BoolArray    = pq.BoolArray
	ByteaArray   = pq.ByteaArray
	Float64Array = pq.Float64Array
	Float32Array = pq.Float32Array
	GenericArray = pq.GenericArray
	Int64Array   = pq.Int64Array
	Int32Array   = pq.Int32Array
	StringArray  = pq.StringArray
)

// MustOpen is like Open but panics if connect operation fails.
func MustOpen(conn string) *standard.DB {
	c, err := Open(conn)
	if err != nil {
		panic(err)
	}
	return c
}

// Open creates and establishes one connection to database.
func Open(conn string) (*standard.DB, error) {
	c, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}
	if err := c.Ping(); err != nil {
		return nil, err
	}
	return standard.NewDB("postgres", c), nil
}
