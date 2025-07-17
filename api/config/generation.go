package config

// Default
type Default struct {
	Host  string `yaml:"host"`
	Api   Api    `yaml:"api"`
	Log   Log    `yaml:"log"`
	Auth  Auth   `yaml:"auth"`
	Mysql Mysql  `yaml:"mysql"`
	Phone Phone  `yaml:"phone"`
	File  File   `yaml:"file"`
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
	PsdErrorLockTime    string `yaml:"psdErrorLockTime"`
	RateLimitCap        int    `yaml:"rateLimitCap"`
	JwtExpireTime       string `yaml:"jwtExpireTime"`
	JwtSecret           string `yaml:"jwtSecret"`
	RoleCacheExpireTime string `yaml:"roleCacheExpireTime"`
	PsdErrorLimit       int    `yaml:"psdErrorLimit"`
	RateLimitInterval   string `yaml:"rateLimitInterval"`
	Issuer              string `yaml:"issuer"`
	IsUserExpireTime    string `yaml:"isUserExpireTime"`
}

// Mysql
type Mysql struct {
	Charset  string `yaml:"charset"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Dbname   string `yaml:"dbname"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

// Phone
type Phone struct {
	Endpoint        string `yaml:"endpoint"`
	SignName        string `yaml:"signName"`
	ExpirationTime  string `yaml:"expirationTime"`
	SendInterval    string `yaml:"sendInterval"`
	AccessKeyId     string `yaml:"accessKeyId"`
	AccessKeySecret string `yaml:"accessKeySecret"`
	TemplateCode    string `yaml:"templateCode"`
}

// File
type File struct {
	Path    string `yaml:"path"`
	MaxSize int64  `yaml:"maxSize"`
}
