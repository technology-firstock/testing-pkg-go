// Copyright (c) [2025] [abc]
// SPDX-License-Identifier: MIT
package Abc

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

func EncodePassword(pwd string) string {
	hash := sha256.Sum256([]byte(pwd))
	return hex.EncodeToString(hash[:])
}

func ReadJKeyFromConfig(userId string) (string, error) {
	configPath, err := getConfigPath()

	if err != nil {
		return "", fmt.Errorf("could not open config file: %w", err)
	}
	file, err := os.Open(configPath)
	if err != nil {
		return "", fmt.Errorf("could not open config file: %w", err)
	}
	defer file.Close()

	var config map[string]interface{}
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return "", fmt.Errorf("could not decode config JSON: %w", err)
	}

	userConfigRaw, ok := config[userId]
	if !ok {
		return "", fmt.Errorf("userId %s not found in config", userId)
	}

	userConfig, ok := userConfigRaw.(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("invalid user config format for userId %s", userId)
	}

	jkeyRaw, ok := userConfig[j_key]
	if !ok {
		return "", fmt.Errorf("jkey not found for userId %s", userId)
	}

	jkey, ok := jkeyRaw.(string)
	if !ok {
		return "", fmt.Errorf("jkey is not a string for userId %s", userId)
	}

	return jkey, nil
}

var configMu sync.Mutex

func SaveJKeyToConfig(data LogoutRequest) error {
	userId := data.UserId
	jkey := data.JKey

	const configFile = "config.json"

	configMu.Lock()
	defer configMu.Unlock()

	config := map[string]map[string]string{}

	if _, err := os.Stat(configFile); err == nil {
		bytes, err := os.ReadFile(configFile)
		if err != nil {
			return err
		}
		if err := json.Unmarshal(bytes, &config); err != nil {
			return err
		}
	}

	if _, ok := config[userId]; !ok {
		config[userId] = map[string]string{}
	}
	config[userId]["jkey"] = jkey

	jsonBytes, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(configFile, jsonBytes, 0644)
}

func RemoveJKeyFromConfig(userId string) error {
	const configFile = "config.json"

	configMu.Lock()
	defer configMu.Unlock()

	// Check if config file exists
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return fmt.Errorf("config file does not exist")
	}

	// Read and unmarshal config
	bytes, err := os.ReadFile(configFile)
	if err != nil {
		return err
	}
	config := map[string]map[string]string{}
	if err := json.Unmarshal(bytes, &config); err != nil {
		return err
	}

	// Remove jkey if present
	if userConfig, ok := config[userId]; ok {
		if _, exists := userConfig["jkey"]; exists {
			delete(userConfig, "jkey")
			// If userConfig is now empty, remove the userId entry
			if len(userConfig) == 0 {
				delete(config, userId)
			}
		}
	}

	// Write updated config back to file
	jsonBytes, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(configFile, jsonBytes, 0644)
}

func getConfigPath() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Join(cwd, "config.json"), nil
}
