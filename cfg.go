package xap

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path/filepath"
)

func (ctx *Ctx) InitSysConfig() {
	ctx.LoadConfig(&ctx.Config)
}

func (ctx *Ctx) LoadConfig(confOut interface{}) {
	configPath := ctx.SysPath.ConfigDir

	configBase, err := ioutil.ReadFile(filepath.Join(configPath, "config.yml"))
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(configBase, confOut)
	if err != nil {
		panic(err)
	}

	name := fmt.Sprintf("config-%s.yml", ctx.Config.Profiles)
	configFile := filepath.Join(configPath, name)

	_, err = os.Stat(configPath)
	if err != nil {
		panic(err)
	}

	configEnv, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(configEnv, confOut)
	if err != nil {
		panic(err)
	}
}
