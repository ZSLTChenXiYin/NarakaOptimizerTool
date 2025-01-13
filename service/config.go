package service

import (
	"encoding/json"
	"io"
	"os"
)

const (
	config_file_path = "NarakaOptimizerToolConfig.json"
)

var (
	config_map = make(map[string]any)
)

func configInit() {
	file, err := os.Open(config_file_path)
	if err != nil {
		ErrorLogger.Fatalf(NewLog("Failed to open config file: %v", err))
	}
	defer file.Close()

	json_config, err := io.ReadAll(file)
	if err != nil {
		ErrorLogger.Fatalf(NewLog("Failed to read config file: %v", err))
	}

	err = json.Unmarshal(json_config, &config_map)
	if err != nil {
		ErrorLogger.Fatalf(NewLog("Failed to unmarshal config file: %v", err))
	}
}

func saveConfig() {
	file, err := os.OpenFile(config_file_path, os.O_WRONLY, 0666)
	if err != nil {
		ErrorLogger.Fatalf(NewLog("Failed to open config file: %v", err))
	}
	defer file.Close()

	json_config, err := json.MarshalIndent(config_map, "", "    ")
	if err != nil {
		ErrorLogger.Fatalf(NewLog("Failed to marshal config file: %v", err))
	}

	err = file.Truncate(0)
	if err != nil {
		ErrorLogger.Fatalf(NewLog("Failed to truncate config file: %v", err))
	}

	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		ErrorLogger.Fatalf(NewLog("Failed to seek config file: %v", err))
	}

	_, err = file.Write(json_config)
	if err != nil {
		ErrorLogger.Fatalf(NewLog("Failed to write config file: %v", err))
	}
}

func developerMode() bool {
	v, ok := config_map["DeveloperMode"]
	if !ok {
		return false
	}

	return v.(bool)
}

func narakaInstallPath() string {
	v, ok := config_map["NarakaInstallPath"]
	if !ok {
		return ""
	}

	return v.(string)
}

func setDeveloperMode(mode bool) {
	config_map["DeveloperMode"] = mode
}

func setNarakaInstallPath(path string) {
	config_map["NarakaInstallPath"] = path
}
