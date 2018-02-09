package bootstrap

import (
	"github.com/devfeel/dotweb"
)

var App *dotweb.DotWeb

func init() {
	//init DotApp
	App = dotweb.New()
	//set log path
	App.SetLogPath("../log")

}