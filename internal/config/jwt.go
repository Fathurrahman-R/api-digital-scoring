package config

type JWTConfig struct {
	AccessSecret         string `mapstructure:"access_secret"`
	RefreshSecret        string `mapstructure:"refresh_secret"`
	AccessTokenLifetime  string `mapstructure:"access_token_lifetime"`
	RefreshTokenLifetime string `mapstructure:"refresh_token_lifetime"`
}
