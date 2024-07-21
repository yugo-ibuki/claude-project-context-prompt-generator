package typescript

import (
	"github.com/AlecAivazis/survey/v2"
)

type TSAnswers struct {
	PackageManager string
	Linter         []string
	Testing        string
	Framework      string
	Hosting        string
	Auth           string
	API            string
	Others         string
}

func AskQuestions() TSAnswers {
	answers := TSAnswers{}

	questions := []*survey.Question{
		{
			Name: "PackageManager",
			Prompt: &survey.Select{
				Message: "Choose a Package Manager:",
				Options: []string{"yarn", "pnpm"},
				Default: "yarn",
			},
		},
		{
			Name: "Linter",
			Prompt: &survey.MultiSelect{
				Message: "Choose a Linter:",
				Options: []string{"ESLint", "Prettier", "Biome"},
				Default: []string{"ESLint", "Prettier"},
			},
		},
		{
			Name: "Testing",
			Prompt: &survey.Select{
				Message: "Choose a Testing framework:",
				Options: []string{"Vitest", "Playwright", "Jest"},
				Default: "Vitest",
			},
		},
		{
			Name: "Framework",
			Prompt: &survey.Select{
				Message: "Choose a Framework:",
				Options: []string{"Next.js", "Hono", "Remix", "Express"},
				Default: "Next.js",
			},
		},
		{
			Name: "Hosting",
			Prompt: &survey.Select{
				Message: "Choose a Hosting platform:",
				Options: []string{"Vercel", "Cloud Run", "Firebase"},
				Default: "Vercel",
			},
		},
		{
			Name: "Auth",
			Prompt: &survey.Select{
				Message: "Choose an Auth provider:",
				Options: []string{"FireAuth", "Auth0", "NextAuth"},
				Default: "FireAuth",
			},
		},
		{
			Name: "API",
			Prompt: &survey.Select{
				Message: "Choose an API style:",
				Options: []string{"GraphQL", "REST"},
				Default: "GraphQL",
			},
		},
		{
			Name: "Others",
			Prompt: &survey.Select{
				Message: "Choose other services:",
				Options: []string{"Firebase", "Google Cloud", "AWS"},
				Default: "Firebase",
			},
		},
	}

	survey.Ask(questions, &answers)
	return answers
}
