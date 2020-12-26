package entities

import (
	"time"

	"home.com/lists/backend/utils"
)

type ListItem struct {
	ID          int64
	ListID      int64
	Information string
	Assignees   []int
	CreatedAt   time.Time
}

func (listItem *ListItem) AddAssignee(user_id int) error {
	if utils.Contains(listItem.Assignees, user_id) {
		return ErrAssigneeAlreadyAssigned
	}

	listItem.Assignees = append(listItem.Assignees, user_id)

	return nil
}

func (listItem *ListItem) Validate() error {
	if len(listItem.Information) <= 10 {
		return ErrInvalidEntity
	}
	return nil
}
