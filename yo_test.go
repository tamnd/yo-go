package yo

import (
	"strings"
	"testing"
)

// The point of this test is not that Open fails. It is that Open fails with the
// exact sentence the other bindings use, character for character. That sentence
// is the one thing a placeholder owes its ecosystem, and it is also the thing
// most likely to drift when somebody edits a doc comment nearby.
func TestOpenReturnsTheSharedMessage(t *testing.T) {
	db, err := Open("app.yo")
	if db != nil {
		t.Fatalf("Open returned a non-nil handle: %#v", db)
	}
	if err == nil {
		t.Fatal("Open returned a nil error")
	}
	const want = "yo is not usable yet. This is a reserved placeholder at 0.0.0; see https://github.com/tamnd/yo"
	if err.Error() != want {
		t.Errorf("message drifted\n got: %s\nwant: %s", err.Error(), want)
	}
	if !strings.Contains(NotYet, "0.0.0") {
		t.Error("NotYet no longer names the version it is reserving")
	}
}
