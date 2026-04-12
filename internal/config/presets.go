package config

import (
	"archive/zip"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const (
	defaultPresetsURL  = "https://github.com/openshell-policy-manager/presets"
	defaultPresetsRepo = "https://github.com/openshell-policy-manager/presets"
	presetsIndexFile   = "presets.json"
	presetsCacheDir    = "presets"
)

var (
	PresetsCacheDir = filepath.Join(GlobalConfigDir, presetsCacheDir)
	PresetsIndex    = filepath.Join(PresetsCacheDir, presetsIndexFile)
)

type Preset struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Version     string   `json:"version"`
	Files       []string `json:"files"`
	PolicySHA   string   `json:"policy_sha"`
}

type PresetIndex struct {
	Updated   time.Time         `json:"updated"`
	Version   string            `json:"version"`
	Presets   map[string]Preset `json:"presets"`
	RemoteSHA string            `json:"remote_sha"`
	LocalSHA  string            `json:"local_sha"`
}

func GetPresetsURL() string {
	if url := os.Getenv("OSHELL_CONFIG_PRESETS_URL"); url != "" {
		return url
	}
	return defaultPresetsURL
}

func GetCustomPresetsRepo() string {
	return os.Getenv("OSHELL_CONFIG_PRESETS_REPO")
}

func GetPresetsSource() string {
	if repo := GetCustomPresetsRepo(); repo != "" {
		return repo
	}
	return defaultPresetsRepo
}

func GetOfflineZip() string {
	return os.Getenv("OSHELL_CONFIG_OFFLINE_ZIP")
}

func PresetsExist() bool {
	_, err := os.Stat(PresetsIndex)
	return err == nil
}

func GetRemoteSHA() (string, error) {
	url := GetPresetsURL() + "/releases/latest/download/index.json"

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Head(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		etag := resp.Header.Get("ETag")
		if etag != "" {
			return etag, nil
		}
		lastModified := resp.Header.Get("Last-Modified")
		return lastModified, nil
	}

	if resp.StatusCode == http.StatusNotFound {
		versionsURL := GetPresetsURL() + "/releases"
		req2, _ := http.NewRequest("GET", versionsURL, nil)
		resp2, err := client.Do(req2)
		if err != nil {
			return "", err
		}
		defer resp2.Body.Close()

		if resp2.Request != nil && resp2.Request.URL != nil {
			return resp2.Request.URL.String(), nil
		}
	}

	return "", fmt.Errorf("failed to get remote SHA: status %d", resp.StatusCode)
}

func FetchPresetsIndex() (*PresetIndex, error) {
	url := GetPresetsURL() + "/releases/latest/download/index.json"

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("download failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("download failed: status %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}

	var index PresetIndex
	if err := json.Unmarshal(data, &index); err != nil {
		return nil, fmt.Errorf("parse index: %w", err)
	}

	sha := sha256.Sum256(data)
	index.RemoteSHA = hex.EncodeToString(sha[:])

	return &index, nil
}

func LoadLocalIndex() (*PresetIndex, error) {
	data, err := os.ReadFile(PresetsIndex)
	if err != nil {
		return nil, err
	}

	var index PresetIndex
	if err := json.Unmarshal(data, &index); err != nil {
		return nil, err
	}

	return &index, nil
}

func SaveLocalIndex(index *PresetIndex) error {
	os.MkdirAll(PresetsCacheDir, 0755)

	data, err := json.MarshalIndent(index, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(PresetsIndex, data, 0644)
}

func DownloadPresets(offlineZip string) error {
	os.MkdirAll(PresetsCacheDir, 0755)

	url := GetPresetsURL() + "/releases/latest/download/presets.zip"

	client := &http.Client{
		Timeout: 60 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("download failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download failed: status %d", resp.StatusCode)
	}

	zipPath := filepath.Join(PresetsCacheDir, "presets.zip")
	out, err := os.Create(zipPath)
	if err != nil {
		return err
	}
	defer out.Close()

	if _, err := io.Copy(out, resp.Body); err != nil {
		return err
	}

	return extractZip(zipPath, PresetsCacheDir)
}

func extractZip(zipPath, destDir string) error {
	r, err := zip.OpenReader(zipPath)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		outPath := filepath.Join(destDir, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(outPath, 0755)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(outPath), 0755); err != nil {
			return err
		}

		outFile, err := os.Create(outPath)
		if err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			outFile.Close()
			return err
		}

		if _, err := io.Copy(outFile, rc); err != nil {
			rc.Close()
			outFile.Close()
			return err
		}

		rc.Close()
		outFile.Close()
	}

	return nil
}

func ExtractOfflineZip(zipPath string) error {
	if _, err := os.Stat(zipPath); err != nil {
		return fmt.Errorf("offline zip not found: %s", zipPath)
	}

	os.MkdirAll(PresetsCacheDir, 0755)

	return extractZip(zipPath, PresetsCacheDir)
}

func UpdatePresets(offlineZip string) (bool, error) {
	hasLocal := PresetsExist()

	var remoteSHA string
	var fetchErr error

	if offlineZip == "" {
		remoteSHA, fetchErr = GetRemoteSHA()
		if fetchErr != nil {
			fmt.Printf("⚠ Could not check remote: %v\n", fetchErr)
		}
	}

	if hasLocal {
		localIndex, err := LoadLocalIndex()
		if err != nil {
			fmt.Printf("⚠ Could not load local index: %v\n", err)
		} else if remoteSHA != "" && localIndex.RemoteSHA == remoteSHA {
			fmt.Println("✓ Presets are up to date")
			return false, nil
		}
		fmt.Println("→ Presets have changed, updating...")
	} else {
		fmt.Println("→ No local presets found, downloading...")
	}

	if offlineZip != "" || fetchErr != nil {
		if offlineZip != "" {
			fmt.Printf("→ Using offline ZIP: %s\n", offlineZip)
			if err := ExtractOfflineZip(offlineZip); err != nil {
				return false, fmt.Errorf("extract offline zip: %w", err)
			}
		} else if hasLocal {
			fmt.Println("⚠ Falling back to local cache")
			return false, nil
		} else {
			return false, fmt.Errorf("no presets available and offline zip not found")
		}
	} else {
		if err := DownloadPresets(""); err != nil {
			if hasLocal {
				fmt.Printf("⚠ Download failed: %v\n", err)
				fmt.Println("⚠ Falling back to local cache")
				return false, nil
			}
			return false, fmt.Errorf("download failed: %w", err)
		}
	}

	index, err := FetchPresetsIndex()
	if err != nil {
		fmt.Printf("⚠ Could not fetch index: %v\n", err)
		return true, nil
	}

	index.LocalSHA = remoteSHA
	if err := SaveLocalIndex(index); err != nil {
		fmt.Printf("⚠ Could not save index: %v\n", err)
	}

	return true, nil
}

func ListPresets() ([]Preset, error) {
	if !PresetsExist() {
		return nil, fmt.Errorf("no presets found. Run 'preset update' first")
	}

	index, err := LoadLocalIndex()
	if err != nil {
		return nil, err
	}

	presets := make([]Preset, 0, len(index.Presets))
	for _, p := range index.Presets {
		presets = append(presets, p)
	}

	return presets, nil
}

func GetPreset(name string) (*Preset, error) {
	if !PresetsExist() {
		return nil, fmt.Errorf("no presets found. Run 'preset update' first")
	}

	index, err := LoadLocalIndex()
	if err != nil {
		return nil, err
	}

	p, ok := index.Presets[name]
	if !ok {
		return nil, fmt.Errorf("preset not found: %s", name)
	}

	return &p, nil
}

func LoadBlueprintPresets() (map[string]ProfileAdditions, error) {
	bp, err := LoadBlueprint()
	if err != nil {
		return nil, fmt.Errorf("load blueprint: %w", err)
	}

	presets := make(map[string]ProfileAdditions)
	for _, profile := range bp.Profiles {
		addition, err := LoadProfileAddition(profile)
		if err != nil {
			continue
		}
		presets[profile] = *addition
	}

	return presets, nil
}

func GetBlueprintPreset(name string) (*ProfileAdditions, error) {
	presets, err := LoadBlueprintPresets()
	if err != nil {
		return nil, err
	}

	p, ok := presets[name]
	if !ok {
		return nil, fmt.Errorf("preset not found: %s", name)
	}

	return &p, nil
}

func ListBlueprintPresets() ([]string, error) {
	bp, err := LoadBlueprint()
	if err != nil {
		return nil, err
	}

	return bp.Profiles, nil
}
