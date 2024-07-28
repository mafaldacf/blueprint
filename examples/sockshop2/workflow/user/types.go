package user

// A user with an account.  Accounts are optional for ordering.
type User struct {
    FirstName string 
    LastName  string
    Email     string
    Username  string
    Password  string
    Addresses Address
    Cards     Card
    UserID    string
    Salt      string
}

// A street address
type Address struct {
    Street   string
    Number   string
    Country  string
    City     string
    PostCode string
    ID       string
}

// A credit card
type Card struct {
    LongNum string
    Expires string
    CCV     string
    ID      string
}
