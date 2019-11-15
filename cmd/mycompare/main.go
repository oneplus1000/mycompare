package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/oneplus1000/mycompare/libmycompare"
)

/*
config.json file
{
    "src": {
        "connectionstring": ""
    },
    "dest": {
        "connectionstring": ""
    }
}
*/

func main() {

	configFilePathPtr := flag.String("config", "", " config file path")
	flag.Parse()
	if configFilePathPtr == nil || *configFilePathPtr == "" {
		flag.Usage()
		return
	}
	cfg, err := readConfigFile(*configFilePathPtr)
	if err != nil {
		log.Fatalf("%+v", err)
		return
	}

	var my = libmycompare.MyCompare{
		SrcConfig: libmycompare.ConnConfig{
			ConnString: cfg.Src.Connectionstring,
		},
		DestConfig: libmycompare.ConnConfig{
			ConnString: cfg.Dest.Connectionstring,
		},
	}
	err = my.Run()
	if err != nil {
		log.Fatalf("%+v", err)
	}
}

func readConfigFile(path string) (configFile, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return configFile{}, fmt.Errorf("%w", err)
	}
	var cfg configFile
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return configFile{}, fmt.Errorf("%w", err)
	}
	return cfg, nil
}

type configFile struct {
	Src  databaseConfigFile `json:"src"`
	Dest databaseConfigFile `json:"dest"`
}

type databaseConfigFile struct {
	Connectionstring string `json:"connectionstring"`
}
