package i18n

import (
	"os"
	"sync"
)

var (
	instance *Manager
	once     sync.Once
)

type Manager struct {
	locale string
	trans  map[string]string
	mu     sync.RWMutex
}

type TranslationData struct {
	Wizard map[string]string `yaml:"wizard"`
}

func Init(locale string) error {
	var err error
	once.Do(func() {
		err = loadTranslations(locale)
	})
	return err
}

func loadTranslations(locale string) error {
	if locale == "" {
		locale = os.Getenv("OSHELL_LOCALE")
	}
	if locale == "" {
		locale = "en"
	}

	instance = &Manager{
		locale: locale,
		trans:  make(map[string]string),
	}

	// Try to load from embedded or external locale file
	translations := getBuiltInTranslations(locale)
	instance.trans = translations

	return nil
}

func getBuiltInTranslations(locale string) map[string]string {
	// Default English translations
	translations := map[string]string{
		"select_template":     "Select a configuration template:",
		"available_sandboxes": "Available OpenShell sandboxes:",
		"sandbox_name":        "Enter sandbox name:",
		"gitlab_url":          "Enter GitLab URL (e.g., gitlab.com):",
		"gitlab_token":        "Enter GitLab Deploy Token:",
		"feature_config":      "Configure features:",
		"git_read":            "Enable Git read access?",
		"git_write":           "Enable Git write access (push)?",
		"websearch":           "Enable web search?",
		"apt_update":          "Enable APT update?",
		"apt_install":         "Enable APT install?",
		"npm_install":         "Enable npm install?",
		"npm_execute":         "Enable npm execute (npx)?",
		"policy_generated":    "Policy generated: %s",
		"defaults_saved":      "Defaults saved: %s",
		"apply_policy":        "Apply policy to sandbox now?",
	}

	if locale == "de" {
		translations = map[string]string{
			"select_template":     "Konfigurationsvorlage auswählen:",
			"available_sandboxes": "Verfügbare OpenShell-Sandboxes:",
			"sandbox_name":        "Sandbox-Name eingeben:",
			"gitlab_url":          "GitLab-URL eingeben (z.B. gitlab.com):",
			"gitlab_token":        "GitLab Deploy Token eingeben:",
			"feature_config":      "Features konfigurieren:",
			"git_read":            "Git-Lesezugriff aktivieren?",
			"git_write":           "Git-Schreibzugriff (push) aktivieren?",
			"websearch":           "Websuche aktivieren?",
			"apt_update":          "APT-Update aktivieren?",
			"apt_install":         "APT-Install aktivieren?",
			"npm_install":         "npm install aktivieren?",
			"npm_execute":         "npm execute (npx) aktivieren?",
			"policy_generated":    "Policy generiert: %s",
			"defaults_saved":      "Defaults gespeichert: %s",
			"apply_policy":        "Policy jetzt auf Sandbox anwenden?",
		}
	}

	return translations
}

func Get(key string) string {
	if instance == nil {
		// Return key if not initialized
		return key
	}
	instance.mu.RLock()
	defer instance.mu.RUnlock()
	if v, ok := instance.trans[key]; ok {
		return v
	}
	return key
}

func Locale() string {
	if instance == nil {
		return "en"
	}
	return instance.locale
}

func SetLocale(locale string) error {
	if instance == nil {
		return Init(locale)
	}
	instance.mu.Lock()
	defer instance.mu.Unlock()

	instance.locale = locale
	instance.trans = getBuiltInTranslations(locale)

	return nil
}

func GetModuleDir() string {
	// For relative path lookups - use current working directory as fallback
	wd, _ := os.Getwd()
	return wd
}
