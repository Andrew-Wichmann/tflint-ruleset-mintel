package main

import (
	"github.com/terraform-linters/tflint-plugin-sdk/plugin"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/terraform-linters/tflint-ruleset-mintel/rules"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &tflint.BuiltinRuleSet{
			Name:    "mintel",
			Version: "0.1.0",
			Rules: []tflint.Rule{
				rules.NewAwsInstanceExampleTypeRule(),
				rules.NewAwsS3BucketExampleLifecycleRuleRule(),
				rules.NewLocalFileExampleProvisionerRule(),
				rules.NewTerraformBackendTypeRule(),
			},
		},
	})
}
