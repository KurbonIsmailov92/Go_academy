package models

type Employee struct {
	ID         int
	FirstName  string
	LastName   string
	Age        int
	Department string
	Position   string
	Salary     int
	Email      string
	IsDeleted  bool
}
