package wizard

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"ospm/internal/config"
	"ospm/internal/i18n"
)

type Answers struct {
	Template    string
	Sandbox     string
	GitlabURL   string
	GitlabToken string
	GitRead     bool
	GitWrite    bool
	Websearch   bool
	APTUpdate   bool
	APTInstall  bool
	NPMInstall  bool
	NPMExecute  bool
}

func getOrDefault(val, def string) string {
	if val == "" {
		return def
	}
	return val
}

func Run() error {
	defaults, _ := config.LoadDefaults()
	answers := &Answers{}

	// Ensure defaults is not nil
	if defaults == nil {
		defaults = &config.Config{}
	}

	templates := []string{"minimal", "dev", "full"}
	survey.AskOne(&survey.Select{
		Message: i18n.Get("select_template"),
		Options: templates,
		Default: "dev",
	}, &answers.Template)

	if out, err := exec.Command("openshell", "sandbox", "list").Output(); err == nil {
		fmt.Println(i18n.Get("available_sandboxes"))
		fmt.Println(string(out))
	}

	survey.AskOne(&survey.Input{
		Message: i18n.Get("sandbox_name"),
		Default: getOrDefault(defaults.Sandbox, "default"),
	}, &answers.Sandbox)

	// Ensure sandbox name is never empty
	if answers.Sandbox == "" {
		answers.Sandbox = "default"
	}

	survey.AskOne(&survey.Input{
		Message: i18n.Get("gitlab_url"),
		Default: getOrDefault(defaults.GitlabURL, "gitlab.com"),
	}, &answers.GitlabURL)

	// Ensure GitLab URL is never empty
	if answers.GitlabURL == "" {
		answers.GitlabURL = "gitlab.com"
	}

	survey.AskOne(&survey.Input{
		Message: i18n.Get("gitlab_token"),
	}, &answers.GitlabToken)

	fmt.Println("\n" + i18n.Get("feature_config"))

	survey.AskOne(&survey.Confirm{Message: i18n.Get("git_read"), Default: true}, &answers.GitRead)
	survey.AskOne(&survey.Confirm{Message: i18n.Get("git_write"), Default: false}, &answers.GitWrite)
	survey.AskOne(&survey.Confirm{Message: i18n.Get("websearch"), Default: true}, &answers.Websearch)
	survey.AskOne(&survey.Confirm{Message: i18n.Get("apt_update"), Default: false}, &answers.APTUpdate)
	survey.AskOne(&survey.Confirm{Message: i18n.Get("apt_install"), Default: false}, &answers.APTInstall)
	survey.AskOne(&survey.Confirm{Message: i18n.Get("npm_install"), Default: true}, &answers.NPMInstall)
	survey.AskOne(&survey.Confirm{Message: i18n.Get("npm_execute"), Default: true}, &answers.NPMExecute)

	cfg := &config.Config{
		Template:        answers.Template,
		Sandbox:         answers.Sandbox,
		GitlabURL:       answers.GitlabURL,
		TokenConfigured: answers.GitlabToken != "",
		Git: config.GitConfig{
			Enable:   answers.GitRead || answers.GitWrite,
			Read:     answers.GitRead,
			Write:    answers.GitWrite,
			Endpoint: answers.GitlabURL,
		},
		Websearch: answers.Websearch,
		APT: config.APTConfig{
			Enable:  answers.APTUpdate || answers.APTInstall,
			Update:  answers.APTUpdate,
			Install: answers.APTInstall,
		},
		NPM: config.NPMConfig{
			Enable:  answers.NPMInstall || answers.NPMExecute,
			Install: answers.NPMInstall,
			Execute: answers.NPMExecute,
		},
	}

	if answers.GitlabToken != "" {
		tokenPath := filepath.Join(os.Getenv("HOME"), ".config/opencode/credentials/gitlab-deploy-token")
		os.MkdirAll(filepath.Dir(tokenPath), 0700)
		os.WriteFile(tokenPath, []byte(answers.GitlabToken), 0600)
	}

	policy := config.GeneratePolicy(cfg)
	policyPath := filepath.Join(config.PolicyDir, answers.Sandbox+"-policy.yaml")
	if err := policy.Save(policyPath); err != nil {
		return fmt.Errorf("save policy: %w", err)
	}

	cfg.GitlabToken = ""
	if err := config.SaveDefaults(cfg); err != nil {
		return fmt.Errorf("save defaults: %w", err)
	}

	fmt.Printf("\n"+i18n.Get("policy_generated")+"\n", policyPath)
	fmt.Printf(i18n.Get("defaults_saved")+"\n", config.DefaultsFile)

	apply := false
	survey.AskOne(&survey.Confirm{Message: i18n.Get("apply_policy"), Default: false}, &apply)
	if apply {
		cmd := exec.Command("openshell", "policy", "set", answers.Sandbox, "--policy", policyPath, "--wait")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("apply policy: %w", err)
		}
	}

	return nil
}
