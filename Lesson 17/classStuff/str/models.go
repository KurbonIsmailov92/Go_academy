package str

type User struct {
	FirstName string
	LastName  string
	Age       int
	tin       int
}

func (u *User) SetTin(tin int) {
	u.tin = tin
}

func (u *User) GetTin() int {
	return u.tin
}
