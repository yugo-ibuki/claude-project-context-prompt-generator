package typescript

import (
	"fmt"
	"os"
	"strings"
)

func OutputResults(answers TSAnswers) {
	content := "use these technologies below.\n\n"

	content += "Language:\n- TypeScript\n\n"

	if answers.Testing != "" {
		content += "Testing:\n- " + answers.Testing + "\n\n"
	}

	if answers.Framework != "" {
		content += "Framework:\n- " + answers.Framework + "\n\n"
	}

	if answers.Hosting != "" {
		content += "Hosting:\n- " + answers.Hosting + "\n\n"
	}

	if len(answers.API) > 0 {
		content += "API:\n- " + answers.API + "\n\n"
	}

	if answers.Others != "" {
		content += "Others:\n- " + answers.Others + "\n\n"
	}

	if answers.Auth != "" {
		content += "Auth:\n- " + answers.Auth + "\n\n"
	}

	if len(answers.Linter) > 0 {
		content += "Linter:\n- " + strings.Join(answers.Linter, "\n- ") + "\n\n"
	}

	if answers.PackageManager != "" {
		content += "PackageManager:\n- " + answers.PackageManager + "\n"
	}

	// ファイルに書き込む
	err := os.WriteFile("project_context.txt", []byte(content), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Results have been written to project_context.txt")
}
