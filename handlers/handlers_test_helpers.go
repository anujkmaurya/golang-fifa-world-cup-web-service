package handlers

import (
	"path"
	"path/filepath"

	"github.com/anujkmaurya/golang-fifa-world-cup-web-service/data"
)

// reloads JSON into memory to ensure
// proper winner count during tests.
func setup() {
	p, _ := filepath.Abs("./../data/")
	fullpath := path.Join(p, "winners.json")
	data.LoadFromJSON(fullpath)
}
