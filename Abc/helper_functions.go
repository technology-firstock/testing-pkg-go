// Copyright (c) [2025] [abc]
// SPDX-License-Identifier: MIT
package Abc

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
)

func EncodePassword(pwd string) string {
	hash := sha256.Sum256([]byte(pwd))
	return hex.EncodeToString(hash[:])
}

func ReadJKeyFromConfig(configPath string, userId string) (string, error) {
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

func RemoveJKeyFromConfig(userId string) error {
	file, err := os.Open(config_file_path)
	if err != nil {
		return nil
	}
	defer file.Close()

	var config map[string]interface{}
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return fmt.Errorf("could not decode config JSON: %w", err)
	}

	if _, ok := config[userId]; !ok {
		return fmt.Errorf("userId %s not found in config", userId)
	}

	delete(config, userId)

	file, err = os.Create(config_file_path)
	if err != nil {
		return fmt.Errorf("could not create config file: %w", err)
	}
	defer file.Close()

	if err := json.NewEncoder(file).Encode(config); err != nil {
		return fmt.Errorf("could not write updated config JSON: %w", err)
	}

	return nil
}

func SaveJKeyToConfig(data LogoutRequest) error {
	// Extract userId and jkey from data
	userId := data.UserId
	jkey := data.JKey
	// Open or create file
	configFile, err := os.OpenFile(config_file_path, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer configFile.Close()

	// Load existing config
	config := map[string]map[string]string{}
	decoder := json.NewDecoder(configFile)
	_ = decoder.Decode(&config) // ignore error if file is empty

	// Add or update jKey for this user
	if config[userId] == nil {
		config[userId] = map[string]string{}
	}
	config[userId][j_key] = jkey

	// Truncate and seek to beginning before writing
	if err := configFile.Truncate(0); err != nil {
		return err
	}
	if _, err := configFile.Seek(0, 0); err != nil {
		return err
	}

	// Write updated config
	encoder := json.NewEncoder(configFile)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(config); err != nil {
		return err
	}
	return nil
}
