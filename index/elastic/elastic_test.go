package elastic

import (
	"fmt"
	"testing"
	"time"

	"github.com/anmil/quicknote/test"
)

var indexName = "qnote-test"
var indexHost = "http://127.0.0.1:9200"
var index *Index

func TestIndexNoteElasticSearchIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping TestIndexNoteElasticSearchIntegration in short mode")
	}

	var err error
	index, err = NewIndex(indexHost, indexName)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("elasticsearch-index-note", testIndexNote)
	t.Run("elasticsearch-index-notes", testIndexNotes)
	t.Run("elasticsearch-search-note", testSearchNote)
	t.Run("elasticsearch-search-phrase-note", testSearchNotePhrase)
	t.Run("elasticsearch-delete-note", testDeleteNote)
	t.Run("elasticsearch-delete-book", testDeleteBook)

	if err := index.DeleteIndex(); err != nil {
		t.Fatal(err)
	}
}

func testIndexNote(t *testing.T) {
	n := test.GetTestNotes()[0]
	if err := index.IndexNote(n); err != nil {
		t.Fatal(err)
	}
}

func testIndexNotes(t *testing.T) {
	notes := test.GetTestNotes()
	if err := index.IndexNotes(notes); err != nil {
		t.Fatal(err)
	}
}

func testSearchNote(t *testing.T) {
	n := test.GetTestNotes()[0]
	if err := index.IndexNote(n); err != nil {
		t.Fatal(err)
	}

	// Have to wait at least 1 second for the index to complete
	// I hate this, but there is no API to block till completion
	// https://github.com/elastic/elasticsearch/issues/1063
	time.Sleep(time.Millisecond * 1000)

	query := fmt.Sprintf("id:%d", n.ID)
	if ids, total, err := index.SearchNote(query, 10, 0); err != nil {
		t.Fatal(err)
	} else if total != 1 {
		t.Fatalf("Expected 1 results, got %d", total)
	} else if len(ids) != 1 {
		t.Fatalf("Expected 1 ID, got %d", len(ids))
	} else if ids[0] != n.ID {
		t.Fatalf("Expected ID %d, got %d", n.ID, ids[0])
	}
}

func testSearchNotePhrase(t *testing.T) {
	n := test.GetTestNotes()[0]
	if err := index.IndexNote(n); err != nil {
		t.Fatal(err)
	}

	time.Sleep(time.Millisecond * 1000)

	query := "This is test 1 of the basic par"
	if ids, total, err := index.SearchNotePhrase(query, nil, "asc", 10, 0); err != nil {
		t.Fatal(err)
	} else if total != 1 {
		t.Fatalf("Expected 1 results, got %d", total)
	} else if len(ids) != 1 {
		t.Fatalf("Expected 1 ID, got %d", len(ids))
	} else if ids[0] != n.ID {
		t.Fatalf("Expected ID %d, got %d", n.ID, ids[0])
	}
}

func testDeleteNote(t *testing.T) {
	n := test.GetTestNotes()[0]
	if err := index.IndexNote(n); err != nil {
		t.Fatal(err)
	}

	time.Sleep(time.Millisecond * 1000)

	query := fmt.Sprintf("id:%d", n.ID)
	if ids, total, err := index.SearchNote(query, 10, 0); err != nil {
		t.Fatal(err)
	} else if total != 1 {
		t.Fatalf("Expected 1 results, got %d", total)
	} else if len(ids) != 1 {
		t.Fatalf("Expected 1 ID, got %d", len(ids))
	} else if ids[0] != n.ID {
		t.Fatalf("Expected ID %d, got %d", n.ID, ids[0])
	}

	if err := index.DeleteNote(n); err != nil {
		t.Fatal(err)
	}

	time.Sleep(time.Millisecond * 1000)

	query = fmt.Sprintf("id:%d", n.ID)
	if _, total, err := index.SearchNote(query, 10, 0); err != nil {
		t.Fatal(err)
	} else if total != 0 {
		t.Fatalf("Expected 0 results, got %d", total)
	}
}

func testDeleteBook(t *testing.T) {
	notes := test.GetTestNotes()
	n := notes[0]

	if err := index.IndexNotes(notes); err != nil {
		t.Fatal(err)
	}

	time.Sleep(time.Millisecond * 1000)

	query := fmt.Sprintf("book:%s", n.Book.Name)
	if ids, total, err := index.SearchNote(query, 10, 0); err != nil {
		t.Fatal(err)
	} else if int(total) != len(notes) {
		t.Fatalf("Expected %d results, got %d", len(notes), total)
	} else if len(ids) != len(notes) {
		t.Fatalf("Expected %d ID, got %d", len(notes), len(ids))
	}

	if err := index.DeleteBook(n.Book); err != nil {
		t.Fatal(err)
	}

	time.Sleep(time.Millisecond * 1000)

	query = fmt.Sprintf("book:%s", n.Book.Name)
	if _, total, err := index.SearchNote(query, 10, 0); err != nil {
		t.Fatal(err)
	} else if total != 0 {
		t.Fatalf("Expected 0 results, got %d", total)
	}
}
