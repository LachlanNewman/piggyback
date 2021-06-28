package user

type User struct {
	firstName string
	lastName string
}

func NewUser(firstName string, lastName string) *User{
    return &User{firstName,lastName}
}

func (user *User) GetFirstName() string {
	return user.firstName
}

func (user *User) GetLastName() string {
	return user.lastName
}
