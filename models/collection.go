package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
	"github.com/gofrs/uuid"
)

// Collection is used by pop to map your collections database table to your go code.
type Collection struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	Owner       string    `json:"owner" db:"owner"`
	SharedWith  string    `json:"shared_with" db:"shared_with"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (c Collection) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Collections is not required by pop and may be deleted
type Collections []Collection

// String is not required by pop and may be deleted
func (c Collections) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (c *Collection) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: c.Name, Name: "Name"},
		&validators.StringIsPresent{Field: c.Description, Name: "Description"},
		&validators.StringIsPresent{Field: c.Owner, Name: "Owner"},
		&validators.StringIsPresent{Field: c.SharedWith, Name: "SharedWith"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (c *Collection) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (c *Collection) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
