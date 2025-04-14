package config

import (
	"akai-ito/utils"
	"os"
	"path/filepath"

	"github.com/kirsle/configdir"
	"gopkg.in/yaml.v2"
)

type Colors struct {
	Base00 string
	Base03 string
	Base05 string
	Base0D string
}

var DefaultColors = Colors{
	Base00: "#282828",
	Base03: "#585858",
	Base05: "#EBDBB2",
	Base0D: "#83A598",
}

type ConfigS struct {
	Colors Colors
}

var Config ConfigS

func Init() {
	configPath := configdir.LocalConfig("ontake", "akai-ito")
	err := configdir.MakePath(configPath) // Ensure it exists.
	utils.CheckError(err)

	configFile := filepath.Join(configPath, "config.yml")

	// Does the file not exist?
	if _, err = os.Stat(configFile); os.IsNotExist(err) {
		// Create the new config file.
		fh, err := os.Create(configFile)
		utils.CheckError(err)
		defer fh.Close()

		encoder := yaml.NewEncoder(fh)
		encoder.Encode(&ConfigS{Colors: DefaultColors})
		Config = ConfigS{Colors: DefaultColors}
	} else {
		// Load the existing file.
		fh, err := os.Open(configFile)
		utils.CheckError(err)
		defer fh.Close()

		decoder := yaml.NewDecoder(fh)
		decoder.Decode(&Config)

	}
}
