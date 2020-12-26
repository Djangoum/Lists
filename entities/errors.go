package entities

import "errors"

//ErrNotFound not found
var ErrNotFound = errors.New("Not found")

//ErrInvalidEntity invalid entity
var ErrInvalidEntity = errors.New("Invalid entity")

//ErrCannotBeDeleted cannot be deleted
var ErrCannotBeDeleted = errors.New("Cannot Be Deleted")

//ErrNotEnoughBooks cannot borrow
var ErrNotEnoughBooks = errors.New("Not enough books")

//ErrAssigneeAlreadyAssigned cannot be assigned
var ErrAssigneeAlreadyAssigned = errors.New("Book already borrowed")

//ErrBookNotBorrowed cannot return
var ErrBookNotBorrowed = errors.New("Book not borrowed")

//ErrEmailAlreadyExists user email already exists
var ErrEmailAlreadyExists = errors.New("Email already exists")
