package goquestions

import (
	"fmt"
	"os"
	"strings"
)

func OutputResults(answers GoAnswers) {
	content := "use these technologies below.\n\n"

	content += "Language:\n- Go\n\n"

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
		content += "API:\n- " + strings.Join(answers.API, "\n- ") + "\n\n"
	}

	if answers.Others != "" {
		content += "Others:\n- " + answers.Others + "\n\n"
	}

	if answers.Usage != "" {
		content += "Usage:\n- " + answers.Usage + "\n"
	}

	// ファイルに書き込む
	err := os.WriteFile("project_context.txt", []byte(content), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Results have been written to project_context.txt")
}
