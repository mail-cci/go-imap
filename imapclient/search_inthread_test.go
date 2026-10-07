package imapclient

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/emersion/go-imap/v2"
	"github.com/emersion/go-imap/v2/internal/imapwire"
)

func TestWriteSearchKeyInThread(t *testing.T) {
	var buf bytes.Buffer
	w := bufio.NewWriter(&buf)
	enc := imapwire.NewEncoder(w, imapwire.ConnSideClient)
	var uids imap.UIDSet
	uids.AddNum(83763)
	criteria := &imap.SearchCriteria{
		InThread: []imap.SearchCriteriaInThread{{
			Algorithm: imap.ThreadRefs,
			Criteria:  imap.SearchCriteria{UID: []imap.UIDSet{uids}},
		}},
	}
	writeSearchKey(enc, criteria)
	if err := enc.CRLF(); err != nil {
		t.Fatal(err)
	}
	if err := w.Flush(); err != nil {
		t.Fatal(err)
	}
	want := "INTHREAD REFS (UID 83763)\r\n"
	if got := buf.String(); got != want {
		t.Fatalf("got %q, want %q", got, want)
	}
}
