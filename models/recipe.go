package models

import "cookbook/db"

type RecipeData struct {
	ID           string            `json:"id,omitempty" bson:"-"`
	Title        string            `json:"title"`
	Metadata     map[string]string `json:"metadata,omitempty"`
	Ingregients  []*db.Ingredient  `json:"ingredients"`
	Instructions []*db.Instruction `json:"instructions"`
	Notes        []*string         `json:"notes,omitempty"`
}
