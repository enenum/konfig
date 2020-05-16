package main

import (
	"fmt"
	"os"

	"github.com/enenum/konfig/pkg/konfig"
)

func main() {
	type myConfig struct {
		AString  string `json:"aString" yaml:"aString" xml:"aString"`
		AnInt    int    `json:"anInt" yaml:"anInt" xml:"anInt" env:"ANINT"`
		ABool    bool   `json:"aBool" yaml:"aBool" xml:"aBool"`
		Jsonlang string `json:"jsonlang" yaml:"jsonlang" xml:"jsonlang" env:"JSONLANG"`
		Yamllang string `json:"yamllang" yaml:"yamllang" xml:"yamllang" env:"YAMLLANG"`
		Xmllang  string `json:"xmllang" yaml:"xmllang" xml:"xmllang" env:"XMLLANG"`
		GoPath   string `env:"GOPATH"`
	}
	cfg := &myConfig{}
	os.Setenv("ANINT", "100")
	defer os.Unsetenv("ANINT")
	konfig.LoadCfg(cfg, "config.json", "config.yaml", "config.xml", "env")
	fmt.Printf("%+v\n", cfg)
}
