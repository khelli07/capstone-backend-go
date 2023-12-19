package payload

type RegisterRequest struct {
	Username string  `form:"username" binding:"required"`
	Email    string  `form:"email" binding:"required"`
	Password string  `form:"password" binding:"required"`
	Lat      float64 `form:"lat"`
	Long     float64 `form:"long"`
	// Example: sports,adventure
	PreferenceCategories string `form:"preference_categories" binding:"required"`
}

type LoginRequest struct {
	Email    string `form:"email" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type UpdateUserRequest struct {
	Username string  `form:"username" binding:"required"`
	Lat      float64 `form:"lat"`
	Long     float64 `form:"long"`
	// Example: sports,adventure
	PreferenceCategories string `form:"preference_categories" binding:"required"`
}
