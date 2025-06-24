package config

// Api
type Api struct {
	Post       int    `yaml:"post"`
	StaticPath string `yaml:"staticPath"`
	BaseUrl    string `yaml:"baseUrl"`
	Host       string `yaml:"host"`
}

// Log
type Log struct {
	Level      string `yaml:"level"`
	Filename   string `yaml:"filename"`
	MaxAge     int    `yaml:"maxAge"`
	MaxBackups int    `yaml:"maxBackups"`
	MaxSize    int    `yaml:"maxSize"`
	Mode       string `yaml:"mode"`
}

// Auth
type Auth struct {
	JwtSecret           string `yaml:"jwtSecret"`
	Issuer              string `yaml:"issuer"`
	IsUserExpireTime    string `yaml:"isUserExpireTime"`
	RoleCacheExpireTime string `yaml:"roleCacheExpireTime"`
	PsdErrorLockTime    string `yaml:"psdErrorLockTime"`
	JwtExpireTime       string `yaml:"jwtExpireTime"`
	PsdErrorLimit       int    `yaml:"psdErrorLimit"`
	RateLimitInterval   string `yaml:"rateLimitInterval"`
	RateLimitCap        int    `yaml:"rateLimitCap"`
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

// Phone
type Phone struct {
	TemplateCode    string `yaml:"templateCode"`
	Endpoint        string `yaml:"endpoint"`
	SignName        string `yaml:"signName"`
	ExpirationTime  string `yaml:"expirationTime"`
	SendInterval    string `yaml:"sendInterval"`
	AccessKeyId     string `yaml:"accessKeyId"`
	AccessKeySecret string `yaml:"accessKeySecret"`
}

// Default
type Default struct {
	Api   Api    `yaml:"api"`
	Log   Log    `yaml:"log"`
	Auth  Auth   `yaml:"auth"`
	Mysql Mysql  `yaml:"mysql"`
	Phone Phone  `yaml:"phone"`
	Host  string `yaml:"host"`
}
