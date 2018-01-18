package conf

import (
	"flag"
	"os"
	"log"
	"github.com/larspensjo/config"
	"fmt"
)


type _PeeConf struct {
	Address string
	App string
	SelectDB int
	Port int
}


var con_file = flag.String("con_file", "conf.ini", "General configuration file")

var PeeConf = make(map[string]string)


func init(){

	arg := os.Args

	if arg == nil || len(arg) < 2 {
		fmt.Println("please choose configuration platform dev|test|master")
		os.Exit(0)
	}

	flag.Parse()

	cfg, err := config.ReadDefault(*con_file)

	if err != nil {
		log.Fatalf("Fail to find", *con_file, err)
	}

	if cfg.HasSection(os.Args[1]) {
		section, err := cfg.SectionOptions(os.Args[1])

		if err == nil {
			for _, v := range section {
				options, err := cfg.String(os.Args[1], v)
				if err == nil {
					PeeConf[v] = options
				}
			}
		}
	}
}