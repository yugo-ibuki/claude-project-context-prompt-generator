package goquestions

import (
	"errors"
	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/core"
)

type GoAnswers struct {
	Testing   string
	Framework string
	Hosting   string
	Others    string
	Usage     string
	API       []string
}

func AskQuestions() GoAnswers {
	answers := GoAnswers{}

	mandatoryQuestions := []*survey.Question{
		{
			Name: "Testing",
			Prompt: &survey.Select{
				Message: "Choose a Testing framework:",
				Options: []string{"GoMock", "GoTest"},
				Default: "GoMock",
			},
			Validate: survey.Required,
		},
		{
			Name: "Framework",
			Prompt: &survey.Select{
				Message: "Choose a Framework:",
				Options: []string{"Gin", "Echo"},
				Default: "Gin",
			},
			Validate: survey.Required,
		},
		{
			Name: "Hosting",
			Prompt: &survey.Select{
				Message: "Choose a Hosting platform:",
				Options: []string{"Cloud Run"},
				Default: "Cloud Run",
			},
			Validate: survey.Required,
		},
		{
			Name: "Others",
			Prompt: &survey.Select{
				Message: "Choose other services:",
				Options: []string{"Cobra", "Google Cloud"},
				Default: "Cobra",
			},
			Validate: survey.Required,
		},
		{
			Name: "Usage",
			Prompt: &survey.Select{
				Message: "Choose the usage:",
				Options: []string{"CLI", "Web"},
				Default: "CLI",
			},
			Validate: survey.Required,
		},
	}

	err := survey.Ask(mandatoryQuestions, &answers)
	if err != nil {
		panic(err)
	}

	apiQuestion := &survey.Question{
		Name: "API",
		Prompt: &survey.MultiSelect{
			Message: "Choose an API style (optional, select up to 1):",
			Options: []string{"GraphQL", "REST"},
		},
		Validate: func(ans interface{}) error {
			choices, ok := ans.([]core.OptionAnswer)
			if !ok {
				return errors.New("unexpected answer type")
			}
			if len(choices) > 1 {
				return errors.New("please select at most 1 API style")
			}
			return nil
		},
	}

	err = survey.Ask([]*survey.Question{apiQuestion}, &answers)
	if err != nil {
		panic(err)
	}

	return answers
}
