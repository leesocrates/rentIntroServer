package config

type DbServer struct {
  IP       string
  Port     int
  User     string
  Pwd      string
  Database string
}

type Redis struct {
  IP   string
  Port int
  Pwd  string
  Db   int
}

type Influx struct {
  IP       string
  Port     int
  UserName string
  Pwd      string
  Db       string
}

type ApiService struct {
  Addr   string
  Port   int
  Domain string
}

type GeneralSettings struct {
  IP          string
  Port        int
  Debug       int
  DisableAuth int
}

type Notifications struct {
  VerifyCode string
}

type AuthSettings struct {
  Key string
}

type WeChatAuthSetting struct {
  Token string
}

type SmsAuthSetting struct {
  AccessKeyID     string
  SecretAccessKey string
  SignName        string
}

type LdapServer struct {
  Host   string
  Port   int
  Domain string
}

type FileUploadSetting struct {
  MaxSize  int64
  DistPath string
}

var General GeneralSettings
var RedisSetting Redis
var Auth AuthSettings
var DbSetting = map[string]DbServer{}
var Settings GeneralSettings
var SmsAuth SmsAuthSetting
var WeChat WeChatAuthSetting
var SysNotify Notifications
var UploadConfig FileUploadSetting
