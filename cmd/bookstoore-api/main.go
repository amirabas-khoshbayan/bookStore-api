package main

import (
	"github.com/amiremperor/bookStore-api/internal/config"
	"log"
)
var cfg  = &config.Config{}
//قبل از اجرای برنامه
func init()  {

	if err := config.Parse("build/config/config.yaml",cfg); err != nil{
		log.Fatalln(err)
	}
	if err := config.ReadEnv(cfg); err != nil {
		log.Fatalln(err)
	}
    config.SetConfig(cfg)
}

func main()  {
	
}
