package models

type Student struct {
	ID       int    `json:"id,omitempty"`
	FirstName string `json:"first_name"`
	Age      int    `json:"age"`
	Enrolled bool   `json:"enrolled"`
}
