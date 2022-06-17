// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Author struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Books []*Book `json:"books"`
}

type Book struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	AuthorID int     `json:"authorId"`
	Author   *Author `json:"author"`
}
