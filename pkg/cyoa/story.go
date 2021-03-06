package cyoa

import (
	"encoding/json"
	"io"
)

// Story is Choose Your Own Adventure book
type Story map[string]Chapter

// Chapter of a book
type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

// Option of a Chaper
type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

// JSONStory reads Story from reader
func JSONStory(r io.Reader) (Story, error) {
	var book Story
	if err := json.NewDecoder(r).Decode(&book); err != nil {
		return nil, err
	}
	return book, nil
}
