package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"creates_at:"`
}

func (note Note) Display() {
	fmt.Printf("Title: %v\n%v\n", note.Title, note.Content)
	// time := time.Now()
	// fmt.Println(time)

}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(note)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func New(ntitle, ncontent string) (*Note, error) {

	if ntitle == "" || ncontent == "" {
		return nil, errors.New("enter the text")
	}

	return &Note{
		Title:     ntitle,
		Content:   ncontent,
		CreatedAt: time.Now(),
	}, nil

}
