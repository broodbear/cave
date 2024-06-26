//go:build linux

package config

import "os"

func GetConfigPath() string {
	home, _ := os.UserHomeDir()

	return home + "/.config/cave"
}
