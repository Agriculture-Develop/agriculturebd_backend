package config

// File
type File struct {
	Path    string `yaml:"path"`
	MaxSize int    `yaml:"maxSize"`
}

// Api
type Api struct {
	BaseUrl    string `yaml:"baseUrl"`
	Host       string `yaml:"host"`
	Post       int    `yaml:"post"`
	StaticPath string `yaml:"staticPath"`
}

// Default
type Default struct {
	Log   Log    `yaml:"log"`
	Auth  Auth   `yaml:"auth"`
	Mysql Mysql  `yaml:"mysql"`
	Phone Phone  `yaml:"phone"`
	File  File   `yaml:"file"`
	Host  string `yaml:"host"`
	Api   Api    `yaml:"api"`
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
	Issuer              string `yaml:"issuer"`
	RoleCacheExpireTime string `yaml:"roleCacheExpireTime"`
	PsdErrorLimit       int    `yaml:"psdErrorLimit"`
	PsdErrorLockTime    string `yaml:"psdErrorLockTime"`
	RateLimitInterval   string `yaml:"rateLimitInterval"`
	RateLimitCap        int    `yaml:"rateLimitCap"`
	JwtExpireTime       string `yaml:"jwtExpireTime"`
	JwtSecret           string `yaml:"jwtSecret"`
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

// Phone
type Phone struct {
	ExpirationTime  string `yaml:"expirationTime"`
	SendInterval    string `yaml:"sendInterval"`
	AccessKeyId     string `yaml:"accessKeyId"`
	AccessKeySecret string `yaml:"accessKeySecret"`
	TemplateCode    string `yaml:"templateCode"`
	Endpoint        string `yaml:"endpoint"`
	SignName        string `yaml:"signName"`
}
