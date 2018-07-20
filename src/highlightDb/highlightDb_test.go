package highlightDb

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	highlightTexts := []string{"New"}
	InsertHighlights(highlightTexts)
}

func TestMoveHighlight(t *testing.T) {
	highlightTexts := []string{"New"}
	InsertHighlights(highlightTexts)
	SetHighlightAsPosted("New")
	highlights := GetUnpostedHighlights()
	RemoveHighlightFromDatabase("New")
	if len(highlights) <= 0 {

	} else {
		t.Errorf("should have deleted New highlight, size: %d", len(highlights))
	}
}

func TestGetHighlights(t *testing.T) {
	highlightTexts := []string{"New"}
	InsertHighlights(highlightTexts)
	highlights := GetUnpostedHighlights()
	if len(highlights) <= 0 {
		t.Errorf("GetHighlights did not return any highlights!")
	} else {
		for _, highlight := range highlights {
			SetHighlightAsPosted(highlight.Text)
			RemoveHighlightFromDatabase(highlight.Text)
		}
	}
}
