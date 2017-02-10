package utils

import (
	"fmt"
	"io/ioutil"
)

// plugin interface
type IPlugin interface {
	DoWork() bool
}

// Core implements IPlugin
// there can be a lot of different implementations for IPlugin
// different plugin descriptions can use different implementations
type Core struct {
	Method string
}

func (c *Core) DoWork() bool {
	fmt.Println("doing work", c.Method)
	return true
}

func ResolvePlugins() map[string]IPlugin {
	pluginDir := "../plugins"
	files, err := ioutil.ReadDir(pluginDir)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	ret := make(map[string]IPlugin)
	for _, f := range files {
		c := ReadConfig(pluginDir + "/" + f.Name())
		ret[c["name"]] = &Core{c["generation_method"]}
	}
	return ret
}
