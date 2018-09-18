package config

import (
	"github.com/spf13/viper"
	"github.com/wonktnodi/go-utils/log"
)

func InitConfig(filename string) (inst *viper.Viper, err error) {
	inst = viper.New()
	if nil == inst {
		return
	}

	inst.AddConfigPath("./conf")
	inst.SetConfigName(filename)

	if err = inst.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	return
}

func GetSettingsByKey(inst *viper.Viper, key string, rawInterface interface{}) (err error) {
	if err := inst.UnmarshalKey(key, rawInterface); err != nil {
		log.Errorf("Error reading config file, %s", err)
		return err
	}
	log.Tracef("get settings[%s]: %v", key, rawInterface)
	return
}