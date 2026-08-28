// Package yo is the Go client for yo, an embedded multi-model database that
// lives in one .yo file.
//
// This version is a name reservation and not a release. Nothing in it works.
// The engine is at M0 and the record plane it would sit on top of is M1, so
// there is no binding to ship yet. See https://github.com/tamnd/yo.
//
// The package clause is yo rather than yogo, so the import needs no alias:
//
//	import yo "github.com/tamnd/yo-go"
//
// works, and so does a bare import.
package yo

import "errors"

// Version is the version of this module. Every tier 1 and tier 2 SDK shares one
// version number with the engine, so this is the engine's number and not a line
// of its own.
const Version = "0.0.1"

// NotYet is the message every placeholder in every yo ecosystem carries. It is
// one sentence, identical across all of them, so that somebody who meets it in
// one language and searches for it finds one answer rather than six.
//
// It is built from Version rather than written out, because the first draft
// wrote the version into the string and the string then outlived the version it
// named.
const NotYet = "yo is not usable yet. This is a reserved placeholder at " + Version + "; see https://github.com/tamnd/yo"

// Db is the database handle. It has no fields, because there is nothing yet to
// put in them.
type Db struct{}

// Open opens a database at path. It always fails: this version is a reserved
// placeholder.
//
// It returns an error rather than panicking. Every other language's placeholder
// raises, because raising is what their callers already handle at that point in
// the code; in Go the thing a caller already handles is a non-nil error, and a
// package whose first call panics is one nobody can evaluate in a scratch main.
func Open(path string) (*Db, error) {
	return nil, errors.New(NotYet)
}

// Close closes the database. It is here so that the deferred close in the
// documented snippet compiles, and it does nothing.
func (db *Db) Close() error {
	return nil
}
