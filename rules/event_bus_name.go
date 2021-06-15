package rules

import (
	"strings"

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

// Check checks whether ...
func (r *AwsInstanceExampleTypeRule) Check(runner tflint.Runner) error {
	return runner.WalkModuleCalls(func(call *configs.ModuleCall) error {
		if strings.HasPrefix(call.SourceAddr, "/mintel/satoshi/infrastructure/aws-infrastructure-modules//modules/services/event-bus-topic") {
			return runner.EmitIssue(r, "We're ballin'", call.SourceAddrRange)
		}

		return nil
	})
}
