package json

import (
	"time"
)

// Commit commit のjson表現
type Commit struct {
	Hash         string
	Author       string
	AuthorDate   time.Time
	Commiter     string
	CommiterDate time.Time
	Message      string
}
