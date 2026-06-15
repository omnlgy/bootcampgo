package auth

type Authenticator interface {
	ValidateToken(token string) bool
}

type MerchantAuth struct {
	tokenKey string
}

func NewMerchantAuth(tokenKey string) *MerchantAuth {
	return &MerchantAuth{
		tokenKey: tokenKey,
	}
}

func (m *MerchantAuth) ValidateToken(token string) bool {
	return token == m.tokenKey
}
