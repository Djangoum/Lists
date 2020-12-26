package entities

import (
	"time"
)

type List struct {
	ID         int64
	Title      string
	CreatedAt  time.Time
	ModifiedAt time.Time
}

func NewList(title string, asignees []int) (*List, error) {
	newList := &List{
		Title:     title,
		CreatedAt: time.Now().UTC(),
	}

	err := newList.Validate()
	if err != nil {
		return nil, ErrInvalidEntity
	}
	return newList, nil
}

func (list *List) AddListItem(information string) (*ListItem, error) {
	newListItem := &ListItem{
		ListID:      list.ID,
		Information: information,
		CreatedAt:   time.Now().UTC(),
	}

	err := newListItem.Validate()

	if err != nil {
		return nil, ErrInvalidEntity
	}

	return newListItem, nil
}

func (list *List) Validate() error {
	if list.Title == "" {
		return ErrInvalidEntity
	}
	return nil
}
