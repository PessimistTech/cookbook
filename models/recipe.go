package models

type Recipe struct {
	ID           ObjectID          `json:"id,omitempty" bson:"_id,omitempty"`
	Title        string            `json:"title"`
	Metadata     map[string]string `json:"metadata"`
	Ingregients  []Ingredient      `json:"ingredients"`
	Instructions []Instruction     `json:"instructions"`
	Notes        []string          `json:"notes,omitempty"`
}

type ObjectID string

type Ingredient struct {
	Quantity   int    `json:"quantity"`
	Unit       string `json:"unit"`
	Ingredient string `json:"ingredient"`
}

type Instruction struct {
	StepNum     int    `json:"stepNum"`
	Instruction string `json:"instruction"`
}
