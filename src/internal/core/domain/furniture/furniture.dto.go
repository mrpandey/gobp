package fdom

type FurnitureRecord struct {
	ID   uint          `json:"id"`
	Type FurnitureType `json:"type"`
	Name string        `json:"name"`
}

type FurnitureID struct {
	ID uint `json:"id" validate:"required,gt=0"`
}

type AddFurnitureRequest struct {
	Type FurnitureType `json:"type" validate:"required,valid_furniture_type"`
	Name string        `json:"name" validate:"required,min=3,max=200"`
}
