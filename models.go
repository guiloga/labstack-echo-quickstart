package echo_quickstart

// User model
type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

// MockedUsersList is a fake fixed list of users
var MockedUsersList = &[]User{
	{Name: "Jon", Email: "jon@gmail.com"},
	{Name: "Shane", Email: "shane@gmail.com"},
	{Name: "Albert", Email: "albert@gmail.com"},
	{Name: "Pepe", Email: "pepe@gmail.com"},
}
