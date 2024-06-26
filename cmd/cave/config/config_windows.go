//go:build windows

package config

func GetConfigPath() string {
	return ".config/cave"
}
