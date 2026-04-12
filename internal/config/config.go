package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var (
	// XDG_CONFIG_HOME defaults to ~/.config if not set
	configHome = os.Getenv("XDG_CONFIG_HOME")
	homeDir    = os.Getenv("HOME")

	// TokenDir for credentials
	TokenDir string
	// PolicyDir stays at ~/.ospm/ (OpenShell specific)
	PolicyDir string

	// GlobalConfigDir follows XDG spec: ~/.config/ospm/
	GlobalConfigDir string
	// DefaultsFile in global config directory
	DefaultsFile string
	// PolicyJSON in policy directory
	PolicyJSON string
	// TemplatesDir in global config directory
	TemplatesDir string
)

func init() {
	// Set homeDir to HOME if not set
	if homeDir == "" {
		homeDir = os.Getenv("USERPROFILE")
	}

	// Set default if XDG_CONFIG_HOME is not set
	if configHome == "" {
		configHome = filepath.Join(homeDir, ".config")
	}

	// Initialize all paths
	GlobalConfigDir = filepath.Join(configHome, "ospm")
	DefaultsFile = filepath.Join(GlobalConfigDir, "defaults.yaml")
	TemplatesDir = filepath.Join(GlobalConfigDir, "templates")

	TokenDir = filepath.Join(homeDir, ".config", "opencode", "credentials")
	PolicyDir = filepath.Join(homeDir, ".ospm")
	PolicyJSON = filepath.Join(PolicyDir, "policy.json")
}

// LocalConfigDir returns the local config path in project root
func LocalConfigDir(projectRoot string) string {
	return filepath.Join(projectRoot, ".ospm")
}

// LocalDefaultsFile returns the local defaults file path
func LocalDefaultsFile(projectRoot string) string {
	return filepath.Join(LocalConfigDir(projectRoot), "defaults.yaml")
}

// LoadConfigWithHierarchy loads config with precedence: project > global > defaults
// Priority: LocalDefaultsFile > GlobalConfigDir > hardcoded defaults
func LoadConfigWithHierarchy(projectRoot string) (*Config, error) {
	cfg := &Config{
		Template:  "minimal",
		Sandbox:   "default",
		Websearch: false,
		Git:       GitConfig{Enable: false},
		NPM:       NPMConfig{Enable: false},
		APT:       APTConfig{Enable: false},
	}

	// 1. Load global defaults first
	if globalCfg, err := LoadDefaults(); err == nil {
		cfg = globalCfg
	}

	// 2. Override with local project config if exists
	localFile := LocalDefaultsFile(projectRoot)
	if data, err := os.ReadFile(localFile); err == nil {
		var localCfg Config
		if err := yaml.Unmarshal(data, &localCfg); err == nil {
			// Merge: local overrides global
			if localCfg.Template != "" {
				cfg.Template = localCfg.Template
			}
			if localCfg.Sandbox != "" {
				cfg.Sandbox = localCfg.Sandbox
			}
			if localCfg.GitlabURL != "" {
				cfg.GitlabURL = localCfg.GitlabURL
			}
			cfg.TokenConfigured = localCfg.TokenConfigured
			cfg.Websearch = localCfg.Websearch
			cfg.Git = localCfg.Git
			cfg.NPM = localCfg.NPM
			cfg.APT = localCfg.APT
		}
	}

	return cfg, nil
}

type Config struct {
	Template        string `yaml:"template"`
	Sandbox         string `yaml:"sandbox"`
	GitlabURL       string `yaml:"gitlab_url"`
	GitlabToken     string `yaml:"-"` // Never persist
	TokenConfigured bool   `yaml:"token_configured"`

	Git       GitConfig `yaml:"git"`
	Websearch bool      `yaml:"allow_websearch"`
	NPM       NPMConfig `yaml:"npm"`
	APT       APTConfig `yaml:"apt"`
}

type GitConfig struct {
	Enable   bool   `yaml:"enable"`
	Read     bool   `yaml:"read"`
	Write    bool   `yaml:"write"`
	Endpoint string `yaml:"endpoint"`
}

type NPMConfig struct {
	Enable  bool `yaml:"enable"`
	Install bool `yaml:"install"`
	Execute bool `yaml:"execute"`
}

type APTConfig struct {
	Enable  bool `yaml:"enable"`
	Update  bool `yaml:"update"`
	Install bool `yaml:"install"`
}

func LoadDefaults() (*Config, error) {
	data, err := os.ReadFile(DefaultsFile)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func SaveDefaults(cfg *Config) error {
	os.MkdirAll(GlobalConfigDir, 0755)

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	return os.WriteFile(DefaultsFile, data, 0644)
}
