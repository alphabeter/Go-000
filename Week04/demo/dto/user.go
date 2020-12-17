package dto

// UserDTO struct.
type UserDTO struct {
	ID   uint   `json:"id,string,omitempty"`
	Name string `json:"name"`
	Age  uint   `json:"age,string"`
}
