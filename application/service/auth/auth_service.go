package auth

type AuthService interface {
	Login(loginUser LoginUser) (*SecurityAuthenticatedUser, error)
	AccessTokenByRefreshToken(refreshToken string) (*SecurityAuthenticatedUser, error)
}
