package rules

import (
	hcl "github.com/hashicorp/hcl/v2"
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

// Check checks whether ...
func (r *AwsInstanceExampleTypeRule) Check(runner tflint.Runner) error {
	return runner.WalkResourceAttributes("aws_instance", "instance_type", func(attribute *hcl.Attribute) error {
		return runner.EmitIssueOnExpr(
			r,
			"Just because",
			attribute.Expr,
		)
	});
}
