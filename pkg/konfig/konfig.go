package konfig

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"strings"

	env "github.com/Netflix/go-env"
	"gopkg.in/yaml.v2"
)

// LoadCfg - loads configuration from the given source in order of increasing preference
func LoadCfg(config interface{}, sources ...string) error {
	for _, src := range sources {
		parts := strings.Split(src, ".")
		var cfgBytes []byte
		var err error
		if len(parts) > 1 {
			// its a file
			cfgBytes, err = ioutil.ReadFile(src)
			if err != nil {
				return err
			}
		}
		switch parts[len(parts)-1:][0] {
		case "json":
			if jsonErr := json.Unmarshal(cfgBytes, config); jsonErr != nil {
				continue
			}
		case "yaml":
			if yamlErr := yaml.Unmarshal(cfgBytes, config); yamlErr != nil {
				continue
			}
		case "xml":
			if xmlErr := xml.Unmarshal(cfgBytes, config); xmlErr != nil {
				continue
			}
		case "env":
			_, err := env.UnmarshalFromEnviron(config)
			if err != nil {
				continue
			}
		default:
			continue

		}
	}
	return nil
}
