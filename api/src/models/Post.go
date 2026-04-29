package models

import (
	"errors"
	"html"
	"strings"
	"time"
)

type Post struct {
	ID         uint64    `json:"id,omitempty"`
	Title      string    `json:"title,omitempty"`
	Content    string    `json:"content,omitempty"`
	AuthorID   uint64    `json:"authorId,omitempty"`
	AuthorNick string    `json:"authorNick,omitempty"`
	Likes      uint64    `json:"likes"`
	CreatedAt  time.Time `json:"createdAt,omitempty"`
}

func (post *Post) Prepare() error {
	if err := post.validate(); err != nil {
		return err
	}
	post.format()
	return nil
}

func (post *Post) validate() error {
	if post.Title == "" {
		return errors.New("the title is required and cannot be empty")
	}
	if post.Content == "" {
		return errors.New("the content is required and cannot be empty")
	}
	if post.AuthorID == 0 {
		return errors.New("the author is required and cannot be empty")
	}
	return nil
}

func (post *Post) format() {
	post.Title = html.EscapeString(strings.TrimSpace(post.Title))
	post.Content = html.EscapeString(strings.TrimSpace(post.Content))
}
