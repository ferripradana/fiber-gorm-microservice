package auth

type AuthService interface {
	Login(loginUser LoginUser) (*SecurityAuthenticatedUser, error)
}
