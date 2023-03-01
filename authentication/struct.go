package authentication

type Users struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type Organization struct {
	Name    string
	Code    string
	Address string
	Phone   string
}
