package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
)

func getConf() config.Configer {
	conf, err := config.NewConfig("yaml", "settings.yml")
	if err != nil {
		panic("cant get settings.yml: " + err.Error())
		return nil
	}
	return conf
}

func confGetParams() ParamsData {
	conf := getConf()
	sec, err := conf.Bool("conn::Serversecure")
	if err != nil {
		beego.Error(err.Error())
	}
	data := ParamsData{
		Serverhostname: conf.String("conn::Serverhostname"),
		Serversecure:   sec,
		Serverpassword: conf.String("conn::Serverpassword"),
	}
	return data
}
