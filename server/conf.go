package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type ClusterConf struct {
	Nodes []string `json:"nodes"`
}

func GetConfig(path string) ClusterConf {
	var conf ClusterConf

	if path == "" {
		return conf
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(data, &conf); err != nil {
		log.Fatal(err)
	}

	return conf
}
