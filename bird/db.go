package bird

import (
	"fmt"
	"strings"
)

// DB denotes a bird DB
type DB map[string]*Bird

// NewDB returns a new DB instance
func NewDB() *DB {
	return &DB{}
}

// Add adds a bird to the DB
func (db *DB) Add(bird *Bird) error {
	if _, ok := (*db)[bird.GenericName]; ok {
		return fmt.Errorf("Duplicate entry for %s", bird.GenericName)
	}
	(*db)[bird.GenericName] = bird

	return nil
}

// GetByGenericName returns a bird from db based on genericName
func (db *DB) GetByGenericName(genericName string) (*Bird, error) {
	bird, ok := (*db)[genericName]
	if !ok {
		return nil, fmt.Errorf("No entry for %s", genericName)
	}
	return bird, nil
}

// GetAll returns a list of all birds in db
func (db *DB) GetAll() []Bird {
	birds := []Bird{}
	for _, bird := range *db {
		birds = append(birds, *bird)
	}
	return birds
}

func (db *DB) String() string {
	sb := strings.Builder{}
	for _, bird := range *db {
		sb.WriteString((*bird).String())
		sb.WriteString("\n")
	}

	return sb.String()
}
