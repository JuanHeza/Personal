package models

var ()

//User structure for sesion data
type User struct {
	ID       string
	Password string
	Active   bool
	Admin    bool
}

//NewUser creates a new instance of the structure
func NewUser(id, password string) *User {
	user := &User{ID: id, Password: password}
	return user
}
