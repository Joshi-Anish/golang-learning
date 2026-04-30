package note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

func (note Note) Display() {
	fmt.Printf("Title: %v\n%v", note.title, note.content)
	// fmt.Print("", note.content)
}

func New(ntitle, ncontent string) (*Note, error) {

	if ntitle == "" || ncontent == "" {
		return nil, errors.New("enter the text")
	}

	return &Note{
		title:     ntitle,
		content:   ncontent,
		createdAt: time.Now(),
	}, nil

}
