package article

import (
	"errors"
	"time"
)

var (
	ErrAlreadyExists = errors.New(`Article already exists`)
)

type Article struct {
	ID        int
	Title     string
	Content   string
	URL       string
	CreatedAt time.Time
}
