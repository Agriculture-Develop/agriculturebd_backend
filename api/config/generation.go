package config

// Auth
type Auth struct {
	IsUserExpireTime    string `yaml:"isUserExpireTime"`
	RoleCacheExpireTime string `yaml:"roleCacheExpireTime"`
	PsdErrorLimit       int    `yaml:"psdErrorLimit"`
	RateLimitInterval   string `yaml:"rateLimitInterval"`
	JwtExpireTime       string `yaml:"jwtExpireTime"`
	PsdErrorLockTime    string `yaml:"psdErrorLockTime"`
	RateLimitCap        int    `yaml:"rateLimitCap"`
	JwtSecret           string `yaml:"jwtSecret"`
	Issuer              string `yaml:"issuer"`
}

// Mysql
type Mysql struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Dbname   string `yaml:"dbname"`
}

// Default
type Default struct {
	Host  string `yaml:"host"`
	Api   Api    `yaml:"api"`
	Log   Log    `yaml:"log"`
	Auth  Auth   `yaml:"auth"`
	Mysql Mysql  `yaml:"mysql"`
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
	MaxBackups int    `yaml:"maxBackups"`
	MaxSize    int    `yaml:"maxSize"`
	Mode       string `yaml:"mode"`
	Level      string `yaml:"level"`
	Filename   string `yaml:"filename"`
	MaxAge     int    `yaml:"maxAge"`
}
