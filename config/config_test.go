package config

import (
	"testing"
	"os"
	"fmt"
)

func TestT(t *testing.T) {

	path, err := os.Getwd()
	if err != nil {
		t.Error(err)
		return
	}

	confPath := path + "/invalid.ts"
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

	if len(conf.Map) != 4 {
		t.Error(len(conf.Map))
		return
	}

	fmt.Println("----------------------")
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

	if "halo" != conf.String("string") {
		t.Error(conf.String("key"))
		return
	}
	if 1 != conf.Int("int") {
		t.Error(conf.Int("int"))
		return
	}
	if 2 != conf.Int32("int32") {
		t.Error(conf.Int32("int32"))
		return
	}
	if 3 != conf.Int64("int64") {
		t.Error(conf.Int64("int64"))
		return
	}
	if 1.1 != conf.Float32("float32") {
		t.Error(conf.Float32("float32"))
		return
	}
	fmt.Println("----------")
	if 2.2 != conf.Float64("float64") {
		t.Error(conf.Float64("float64"))
		return
	}

}
