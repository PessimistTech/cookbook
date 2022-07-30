package models

type RecipeData struct {
	ID           string            `json:"id,omitempty" bson:"-"`
	Title        string            `json:"title"`
	Metadata     map[string]string `json:"metadata,omitempty"`
	Ingregients  []IngredientData  `json:"ingredients"`
	Instructions []InstructionData `json:"instructions"`
	Notes        []string          `json:"notes,omitempty"`
}

type IngredientData struct {
	Quantity   int    `json:"quantity"`
	Unit       string `json:"unit"`
	Ingredient string `json:"ingredient"`
}

type InstructionData struct {
	StepNum     int    `json:"stepNum"`
	Instruction string `json:"instruction"`
}
