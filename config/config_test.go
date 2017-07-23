package config

import (
	"testing"
	"os"
)

func TestT(t *testing.T) {

	path, err := os.Getwd()
	if err != nil {
		t.Error(err)
		return
	}

	confPath := path +"/invalid.ts"
	conf := NewConfig()
	err = conf.Config(confPath)
	if err == nil {
		t.Error(err)
		return
	}

	confPath = path + "/errconfigType.type"
	err = conf.Config(confPath)
	if err == nil {
		t.Error(err)
		return
	}

	confPath = path + "/testconfig.conf"
	err = conf.Config(confPath)
	if err != nil {
		t.Error(err)
		return
	}

	conf.Print()

	if len(conf.Map) != 3 {
		t.Error(len(conf.Map))
		return
	}

	confPath = path + "/config.conf"
	conf = NewConfig()
	err = conf.Config(confPath)
	if err != nil {
		t.Error(err)
		return
	}

	conf.Print()

	if len(conf.Map) != 6 {
		t.Error(len(conf.Map))
		return
	}

	if "1" != conf.String("int") {
		t.Error(conf.String("int"))
		return
	}
}
