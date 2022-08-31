package config

import (
	"encoding/json"
	"os"
	"path"
	"runtime"
)

type ServerCfg struct {
	Port string `json:"port"`
}

func Config() *ServerCfg {
	_, b, _, _ := runtime.Caller(0)
	fpath := path.Join(path.Dir(b), "config.json")

	file, err := os.ReadFile(fpath)
	if err != nil {
		panic(err)
	}

	var cfg ServerCfg
	err = json.Unmarshal(file, &cfg)
	if err != nil {
		panic(err)
	}

	return &cfg
}
