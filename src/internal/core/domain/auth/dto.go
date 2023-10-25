package authdom

type TokenResponse struct {
	AccessToken     string `json:"access_token"`
	AccessExpiresIn uint   `json:"expires_in"`
	TokenType       string `json:"token_type"`
}

// Client's credentials fetched from database.
type CredRecord struct {
	ID           uint
	Slug         string
	HashedSecret string
	IsBlocked    bool
}
