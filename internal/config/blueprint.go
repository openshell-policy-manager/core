package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Blueprint struct {
	Version             string     `yaml:"version"`
	MinOpenShellVersion string     `yaml:"min_openshell_version"`
	Digest              string     `yaml:"digest"`
	Profiles            []string   `yaml:"profiles"`
	Components          Components `yaml:"components"`
}

type Components struct {
	Sandbox SandboxComponent `yaml:"sandbox"`
	Policy  PolicyComponent  `yaml:"policy"`
}

type SandboxComponent struct {
	Name string `yaml:"name"`
}

type PolicyComponent struct {
	Base      string            `yaml:"base"`
	Additions map[string]string `yaml:"additions"`
}

type ProfileAdditions struct {
	Name        string           `yaml:"name"`
	Description string           `yaml:"description"`
	Version     string           `yaml:"version"`
	Network     NetworkAdditions `yaml:"network,omitempty"`
	Binaries    []string         `yaml:"binaries,omitempty"`
	Files       []string         `yaml:"files,omitempty"`
}

type NetworkAdditions struct {
	Allowed []Endpoint `yaml:"allowed,omitempty"`
	Denied  []Endpoint `yaml:"denied,omitempty"`
}

var (
	BlueprintDir        = filepath.Join(PolicyDir, "nemoclaw-blueprint")
	BlueprintFile       = filepath.Join(BlueprintDir, "blueprint.yaml")
	PresetsBlueprintDir = filepath.Join(BlueprintDir, "policies", "presets")
)

func LoadBlueprint() (*Blueprint, error) {
	data, err := os.ReadFile(BlueprintFile)
	if err != nil {
		return nil, fmt.Errorf("read blueprint: %w", err)
	}

	var bp Blueprint
	if err := yaml.Unmarshal(data, &bp); err != nil {
		return nil, fmt.Errorf("parse blueprint: %w", err)
	}

	return &bp, nil
}

func GetProfiles(bp *Blueprint) []string {
	return bp.Profiles
}

func GetProfileAddition(bp *Blueprint, profile string) (string, bool) {
	addition, ok := bp.Components.Policy.Additions[profile]
	return addition, ok
}

func LoadProfileAddition(profile string) (*ProfileAdditions, error) {
	bp, err := LoadBlueprint()
	if err != nil {
		return nil, err
	}

	additionPath, ok := GetProfileAddition(bp, profile)
	if !ok {
		return nil, fmt.Errorf("no policy addition for profile: %s", profile)
	}

	fullPath := filepath.Join(BlueprintDir, additionPath)
	data, err := os.ReadFile(fullPath)
	if err != nil {
		return nil, fmt.Errorf("read profile addition: %w", err)
	}

	var addition ProfileAdditions
	if err := yaml.Unmarshal(data, &addition); err != nil {
		return nil, fmt.Errorf("parse profile addition: %w", err)
	}

	return &addition, nil
}
