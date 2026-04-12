package main

import (
	"fmt"
	"os"

	"ospm/internal/config"
	"ospm/internal/i18n"
	"ospm/internal/wizard"

	"github.com/urfave/cli/v2"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	i18n.Init("en")

	app := &cli.App{
		Name:  "ospm",
		Usage: "OpenShell Policy Manager",
		Commands: []*cli.Command{
			{
				Name:  "init",
				Usage: "Initialize a new configuration",
				Action: func(c *cli.Context) error {
					return wizard.Run()
				},
			},
			{
				Name:  "generate",
				Usage: "Generate policy YAML",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "sandbox", Aliases: []string{"s"}, Value: "default"},
					&cli.StringFlag{Name: "template", Aliases: []string{"t"}, Value: "dev"},
					&cli.BoolFlag{Name: "gitlab", Aliases: []string{"g"}, Usage: "Enable GitLab access"},
					&cli.BoolFlag{Name: "websearch", Aliases: []string{"w"}, Usage: "Enable web search"},
					&cli.StringFlag{Name: "output", Aliases: []string{"o"}, Usage: "Output file path"},
				},
				Action: func(c *cli.Context) error {
					cfg := &config.Config{
						Template: c.String("template"),
						Sandbox:  c.String("sandbox"),
						Git: config.GitConfig{
							Enable: c.Bool("gitlab"),
							Read:   c.Bool("gitlab"),
							Write:  c.Bool("gitlab"),
						},
						Websearch: c.Bool("websearch"),
					}
					policy := config.GeneratePolicy(cfg)
					outputPath := c.String("output")
					if outputPath == "" {
						outputPath = "-" + cfg.Sandbox + "-policy.yaml"
					}
					return policy.Save(outputPath)
				},
			},
			{
				Name:  "status",
				Usage: "Show current configuration status",
				Action: func(c *cli.Context) error {
					cfg, err := config.LoadDefaults()
					if err != nil {
						fmt.Println("No configuration found. Run 'ospm init' first.")
						return nil
					}
					fmt.Printf("Template: %s\n", cfg.Template)
					fmt.Printf("Sandbox: %s\n", cfg.Sandbox)
					fmt.Printf("GitLab: %v\n", cfg.Git.Enable)
					fmt.Printf("Websearch: %v\n", cfg.Websearch)
					return nil
				},
			},
		},
	}

	return app.Run(os.Args)
}
