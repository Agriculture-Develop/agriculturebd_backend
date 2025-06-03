package config

// Default
type Default struct {
	Host string `yaml:"host"`
	Api  Api    `yaml:"api"`
	Log  Log    `yaml:"log"`
	Auth Auth   `yaml:"Auth"`
}

// Api
type Api struct {
	BaseUrl    string `yaml:"baseUrl"`
	Host       string `yaml:"host"`
	Post       int    `yaml:"post"`
	StaticPath string `yaml:"staticPath"`
}

// Log
type Log struct {
	MaxSize    int    `yaml:"maxSize"`
	Mode       string `yaml:"mode"`
	Level      string `yaml:"level"`
	Filename   string `yaml:"filename"`
	MaxAge     int    `yaml:"maxAge"`
	MaxBackups int    `yaml:"maxBackups"`
}

// Auth
type Auth struct {
	AccessJwtSecret      string `yaml:"accessJwtSecret"`
	RefreshJwtSecret     string `yaml:"refreshJwtSecret"`
	Issuer               string `yaml:"issuer"`
	RoleCacheExpireTime  string `yaml:"roleCacheExpireTime"`
	PsdErrorLimit        int    `yaml:"psdErrorLimit"`
	RateLimitInterval    string `yaml:"rateLimitInterval"`
	Model                string `yaml:"model"`
	RefreshJwtExpireTime string `yaml:"refreshJwtExpireTime"`
	IsUserExpireTime     string `yaml:"isUserExpireTime"`
	PsdErrorLockTime     string `yaml:"psdErrorLockTime"`
	RateLimitCap         int    `yaml:"rateLimitCap"`
	Policy               string `yaml:"policy"`
	AccessJwtExpireTime  string `yaml:"accessJwtExpireTime"`
}
