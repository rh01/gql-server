package graphql

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// Run ... 
// run server
func Run() error {
	cfgFile := "/report/config.yaml"
	cfgData, err := ioutil.ReadFile(cfgFile)

	var c Config
	switch {
	case os.IsNotExist(err):
	case err == nil:
		if err := yaml.Unmarshal(cfgData, &c); err != nil {
			return fmt.Errorf("Failed to parse config file %s: %v", cfgFile, err)
		}
	default:
		log.Println(err)
	}
	return nil
}
