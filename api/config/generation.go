package config

// Default
type Default struct {
	Phone Phone  `yaml:"phone"`
	Host  string `yaml:"host"`
	Api   Api    `yaml:"api"`
	Log   Log    `yaml:"log"`
	Auth  Auth   `yaml:"auth"`
	Mysql Mysql  `yaml:"mysql"`
}

// Phone
type Phone struct {
	AccessKeyId     string `yaml:"accessKeyId"`
	AccessKeySecret string `yaml:"accessKeySecret"`
	RegionId        string `yaml:"regionId"`
	ExpirationTime  string `yaml:"expirationTime"`
	SendInterval    string `yaml:"sendInterval"`
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
	MaxAge     int    `yaml:"maxAge"`
	MaxBackups int    `yaml:"maxBackups"`
	MaxSize    int    `yaml:"maxSize"`
	Mode       string `yaml:"mode"`
	Level      string `yaml:"level"`
	Filename   string `yaml:"filename"`
}

// Auth
type Auth struct {
	IsUserExpireTime    string `yaml:"isUserExpireTime"`
	RoleCacheExpireTime string `yaml:"roleCacheExpireTime"`
	RateLimitInterval   string `yaml:"rateLimitInterval"`
	JwtSecret           string `yaml:"jwtSecret"`
	PsdErrorLimit       int    `yaml:"psdErrorLimit"`
	PsdErrorLockTime    string `yaml:"psdErrorLockTime"`
	RateLimitCap        int    `yaml:"rateLimitCap"`
	JwtExpireTime       string `yaml:"jwtExpireTime"`
	Issuer              string `yaml:"issuer"`
}

// Mysql
type Mysql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Dbname   string `yaml:"dbname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
}
