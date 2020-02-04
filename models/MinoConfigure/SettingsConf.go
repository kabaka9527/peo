package MinoConfigure

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
)

func GetConf() config.Configer {
	conf, err := config.NewConfig("ini", "conf/settings.conf")
	if err != nil {
		panic("cant get settings.conf: " + err.Error())
		return nil
	}
	return conf
}

func ConfGetHostName() string {
	conf := GetConf()
	secure, err := conf.Bool("WebSecure")
	if err != nil {
		panic(err)
	}
	if secure {
		return "https://" + conf.String("WebHostName")
	}
	return "http://" + conf.String("WebHostName")
}

func ConfGetWebName() string {
	conf := GetConf()
	return conf.String("WebApplicationName")
}

func ConfGetSMTPEnabled() bool {
	conf := GetConf()
	enabled, err := conf.Bool("SMTPEnabled")
	if err != nil {
		beego.Error(err)
		return false
	}
	return enabled
}