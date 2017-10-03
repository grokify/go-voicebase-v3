package voicebase

func (c *Configuration) SetAccessToken(accessToken string) {
	c.AccessToken = accessToken
	apiKeyIdentifier := "Authorization"
	c.APIKeyPrefix[apiKeyIdentifier] = "Bearer"
	c.APIKey[apiKeyIdentifier] = accessToken
}
