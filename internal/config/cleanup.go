package config

import (
	"os"
	"path/filepath"
)

type CleanupPath struct {
	Path        string
	Description string
}

// CleanupPaths returns all paths created by ospm
func CleanupPaths() []CleanupPath {
	home := os.Getenv("HOME")
	configHome := os.Getenv("XDG_CONFIG_HOME")

	// Fallback zu ~/.config wenn XDG_CONFIG_HOME nicht gesetzt ist
	if configHome == "" {
		configHome = filepath.Join(home, ".config")
	}

	paths := []CleanupPath{
		{
			Path:        filepath.Join(configHome, "ospm"),
			Description: "Configuration",
		},
		{
			Path:        filepath.Join(home, ".oshell-config"),
			Description: "Policies",
		},
		{
			Path:        filepath.Join(home, ".config", "opencode", "credentials", "gitlab-deploy-token"),
			Description: "GitLab Token",
		},
	}

	return paths
}

// VerifyPathsExist prüft welche Pfade existieren
func VerifyPathsExist() []CleanupPath {
	existing := []CleanupPath{}
	for _, p := range CleanupPaths() {
		if _, err := os.Stat(p.Path); err == nil {
			existing = append(existing, p)
		}
	}
	return existing
}

// RemovePath entfernt ein Verzeichnis sicher
func RemovePath(path string) error {
	return os.RemoveAll(path)
}

// RemovePaths entfernt mehrere Verzeichnisse
func RemovePaths(paths []CleanupPath) []error {
	errors := []error{}
	for _, p := range paths {
		if err := RemovePath(p.Path); err != nil {
			errors = append(errors, err)
		}
	}
	return errors
}

// ExpandPath ersetzt ~ mit dem tatsächlichen Home-Verzeichnis
func ExpandPath(path string) string {
	home := os.Getenv("HOME")
	if path == "" {
		return path
	}
	// Ersetze nur wenn mit ~ beginnt
	if len(path) > 0 && path[0] == '~' {
		return filepath.Join(home, path[1:])
	}
	return path
}
