package user
type User struct {
	FirstName string
	LastName string
	Age int
	isAdmin bool
}

func (p User) FullName() string {
    return p.FirstName + " " + p.LastName
}

func (p *User) UpdateAge(age int) {
    (*p).Age = age
}

func NewUser(FirstName string, LastName string, Age int, isAdmin bool) *User {
    return &User{
        FirstName: FirstName,
        LastName: LastName,
        Age: Age,
				isAdmin: isAdmin,
		}
}