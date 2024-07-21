package main

import (
	"fmt"
	"github.com/yugo-ibuki/claude-project-context-prompt-generator/goquestions"
	"github.com/yugo-ibuki/claude-project-context-prompt-generator/typescript"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "prompt-generate-cli",
		Short: "prompt-generate-cli is a CLI tool that generates a prompt.",
		Run:   run,
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

type Answers struct {
	Language string
	TS       typescript.TSAnswers
	Go       goquestions.GoAnswers
	Rust     struct {
		message string
	}
}

func run(cmd *cobra.Command, args []string) {
	answers := Answers{}

	// 言語の選択
	prompt := &survey.Select{
		Message: "Choose a language:",
		Options: []string{"TS", "Go", "Rust"},
	}

	err := survey.AskOne(prompt, &answers.Language)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// 言語に基づいて処理を分岐
	switch answers.Language {
	case "TS":
		answers.TS = typescript.AskQuestions()
	case "Go":
		answers.Go = goquestions.AskQuestions()
	case "Rust":
		answers.Rust.message = "Not defined yet."
	}

	// 結果を表示
	printResults(answers)
}

func printResults(answers Answers) {
	fmt.Printf("Selected Language: %s\n", answers.Language)
	switch answers.Language {
	case "TS":
		fmt.Printf("Package Manager: %s\n", answers.TS.PackageManager)
		fmt.Printf("Linter: %s\n", answers.TS.Linter)
		fmt.Printf("Testing: %s\n", answers.TS.Testing)
		fmt.Printf("Framework: %s\n", answers.TS.Framework)
		fmt.Printf("Hosting: %s\n", answers.TS.Hosting)
		fmt.Printf("Auth: %s\n", answers.TS.Auth)
		fmt.Printf("API: %s\n", answers.TS.API)
		fmt.Printf("Others: %s\n", answers.TS.Others)
		typescript.OutputResults(answers.TS)
	case "Go":
		fmt.Printf("Testing: %s\n", answers.Go.Testing)
		fmt.Printf("Framework: %s\n", answers.Go.Framework)
		fmt.Printf("Hosting: %s\n", answers.Go.Hosting)
		fmt.Printf("Others: %s\n", answers.Go.Others)
		fmt.Printf("Usage: %s\n", answers.Go.Usage)
		if len(answers.Go.API) > 0 {
			fmt.Printf("API: %s\n", answers.Go.API[0])
		}
		goquestions.OutputResults(answers.Go)
	case "Rust":
		fmt.Printf("%s\n", answers.Rust.message)
	}
}
