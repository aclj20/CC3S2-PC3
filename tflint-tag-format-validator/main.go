package main

import (
	"github.com/terraform-linters/tflint-plugin-sdk/plugin"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"tflint-ruleset/rules"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &tflint.BuiltinRuleSet{
			Name:    "tflint-ruleset",
			Version: "0.1.0",
			Rules: []tflint.Rule{
				rules.NewRequiredTagsFormatRule(),
			},
		},
	})
}
