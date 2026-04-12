package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Template struct {
	Name        string     `yaml:"name"`
	Description string     `yaml:"description"`
	Questions   []Question `yaml:"questions"`
}

type Question struct {
	ID         string          `yaml:"id"`
	Prompt     string          `yaml:"prompt"`
	Type       string          `yaml:"type"` // yesno, text, choice
	Default    string          `yaml:"default"`
	Optional   bool            `yaml:"optional"`
	Options    []string        `yaml:"options"` // for choice type
	PolicyMods PolicyModifiers `yaml:"adds_policy"`
}

type PolicyModifiers struct {
	ReadWrite []string   `yaml:"read_write"`
	ReadOnly  []string   `yaml:"read_only"`
	Binaries  []Binary   `yaml:"binaries"`
	Endpoints []Endpoint `yaml:"endpoints"`
}

type Binary struct {
	Path string `yaml:"path"`
}

type Endpoint struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	TLS  string `yaml:"tls"`
}

func ListTemplates() ([]Template, error) {
	entries, err := os.ReadDir(TemplatesDir)
	if err != nil {
		return nil, err
	}

	var templates []Template
	for _, entry := range entries {
		if filepath.Ext(entry.Name()) != ".yaml" {
			continue
		}

		data, err := os.ReadFile(filepath.Join(TemplatesDir, entry.Name()))
		if err != nil {
			continue
		}

		var t Template
		if err := yaml.Unmarshal(data, &t); err != nil {
			continue
		}
		templates = append(templates, t)
	}

	return templates, nil
}

func LoadTemplate(name string) (*Template, error) {
	data, err := os.ReadFile(filepath.Join(TemplatesDir, name+".yaml"))
	if err != nil {
		return nil, err
	}

	var t Template
	if err := yaml.Unmarshal(data, &t); err != nil {
		return nil, err
	}

	return &t, nil
}
