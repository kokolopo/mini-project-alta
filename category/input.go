package category

type InputNewCategory struct {
	Name string `json:"name" binding:"required"`
}
