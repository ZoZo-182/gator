// this package is responsible for reading and writing the JSON file.
package config

import(
    "os"
    "errors"
    "fmt"
    "encoding/json"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
    JSONFile, err := getConfigFilePath()
    if err != nil {
        return Config{}, fmt.Errorf("Error getting JSON file: %w", err)
    }
    file, err := os.Open(JSONFile)
    if err != nil {
        return Config{}, fmt.Errorf("error opening file: %w", err)
    }


    var cfg Config
    decoder := json.NewDecoder(file)
    err = decoder.Decode(&cfg)
    if err != nil {
        return Config{}, err
    }

    return cfg, nil

}

func (cfg Config) SetUser(username string) error {
    cfg.CurrentUserName = username
    err := write(cfg)
    if err != nil {
        return err
    }
    return nil
}

// my helpers 
func getConfigFilePath() (string, error) {
    homeDirName, err := os.UserHomeDir()
    if err != nil {
        return "", errors.New("Error getting home directory")
    }

    configFilePath := homeDirName + configFileName 

    return configFilePath, nil
}

func write(cfg Config) error {
    JSONFile, err := getConfigFilePath()
    data, err := json.Marshal(cfg)
    if err != nil {
        return err
    }

    err = os.WriteFile(JSONFile, data, 0666)
    if err != nil {
        return fmt.Errorf("Error writing jsondata to file %w", err)
    }

    return nil
}
