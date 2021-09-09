package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env/v6"
)

//LoadConfigFormToml 从toml中漏加配置文件，并初始化全局对象
func LoadConfigFormToml(filePath string) error {
	cfg := newConfig()
	if _, err := toml.DecodeFile(filePath, cfg); err != nil {
		return err
	}
	//加载全局配组单例
	//global =cfg
	return nil
}

//从环境变量中加载配置
func LoadConfigFormEnv() error {
	cfg := newConfig()
	if err := env.Parse(cfg); err != nil {
		return err
	}
	//加載全局配置單例
	//global =cfg
	return nil
}
