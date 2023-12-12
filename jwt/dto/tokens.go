package dto

type TokensDto struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokenDto struct {
	RefreshToken string `json:"refreshToken"`
}
