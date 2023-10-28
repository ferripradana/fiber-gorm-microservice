package auth

import "time"

type LoginUser struct {
	Email    string
	Password string
}

type RegisterUser struct {
	UserName  string `example:"UserName"`
	Email     string `example:"some@mail.com"`
	FirstName string `example:"John"`
	LastName  string `example:"Doe"`
	Password  string `example:"SomeHashPass"`
}

type Auth struct {
	AccessToken               string
	RefreshToken              string
	ExpirationAccessDateTime  time.Time
	ExpirationRefreshDateTime time.Time
}

type DataUserAuthenticated struct {
	ID        int    `json:"id"`
	UserName  string `json:"user_name" gorm:"unique"`
	Email     string `json:"email" gorm:"unique"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Status    bool   `json:"status"`
}

type DataSecurityAuthenticated struct {
	JWTAccessToken            string    `json:"jwtAccessToken"`
	JWTRefreshToken           string    `json:"jwtRefreshToken"`
	ExpirationAccessDateTime  time.Time `json:"expirationAccessDateTime"`
	ExpirationRefreshDateTime time.Time `json:"expirationRefreshDateTime"`
}

type SecurityAuthenticatedUser struct {
	Data     DataUserAuthenticated     `json:"data"`
	Security DataSecurityAuthenticated `json:"security"`
}
