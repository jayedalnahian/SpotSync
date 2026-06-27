package dto



type CreateRequest struct{
	Name     string `json:"name" validate:"required,min=2,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
	Role     string `json:"role" validate:"omitempty,oneof=driver admin"`
}


type LoginRequest struct{
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}