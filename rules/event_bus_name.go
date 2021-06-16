package rules

import (
	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/terraform/configs"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
)

// AwsInstanceExampleTypeRule checks whether ...
type AwsInstanceExampleTypeRule struct{}

// NewAwsInstanceExampleTypeRule returns a new rule
func NewAwsInstanceExampleTypeRule() *AwsInstanceExampleTypeRule {
	return &AwsInstanceExampleTypeRule{}
}

// Name returns the rule name
func (r *AwsInstanceExampleTypeRule) Name() string {
	return "awichmann_example"
}

// Enabled returns whether the rule is enabled by default
func (r *AwsInstanceExampleTypeRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *AwsInstanceExampleTypeRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *AwsInstanceExampleTypeRule) Link() string {
	return ""
}

// Checks whether the event bus topic name matches a topic in the event bus.
func (r *AwsInstanceExampleTypeRule) Check(runner tflint.Runner) error {
	return runner.WalkResources("aws_sns_topic", func(resource *configs.Resource) error {
		var body hcl.Attributes
		var topic_name string;

		body, _ = resource.Config.JustAttributes()
		runner.EvaluateExpr(body["name"].Expr, &topic_name, nil)

		if topic_name == "foobarbaz" {
			return runner.EmitIssue(r, "Balling", body["name"].NameRange)
		}

		return nil
	})
}
