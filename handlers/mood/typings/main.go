package typings

type Body struct {
	Icon        string `form:"icon" json:"icon" binding:"required"`
	Title       string `form:"title" json:"title" binding:"required"`
	Description string `form:"description" json:"description" binding:"required"`
}
