package json

import (
	"time"
)

// Commit commit のjson表現
type Commit struct {
	Hash         string    `json:"hash"`
	Author       string    `json:"author"`
	AuthorDate   time.Time `json:"authorDate"`
	Commiter     string    `json:"commiter"`
	CommiterDate time.Time `json:"commiterDate"`
	Message      string    `json:"message"`
}
