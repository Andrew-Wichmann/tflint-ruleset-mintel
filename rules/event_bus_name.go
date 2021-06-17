package rules

import (
	"encoding/gob"
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/terraform/configs"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"gitlab.com/mintel/everest/event-bus/events/mintel-events-go/topics"
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
		gob.Register(map[string]interface{}{})
		var body hcl.Attributes
		var resource_topic_name string
		resource_tags := make(map[string]string)

		body, _ = resource.Config.JustAttributes()

		tags, ok := body["tags"]
		if !ok {
			return nil
		}
		err := runner.EvaluateExpr(tags.Expr, &resource_tags, nil)
		err = runner.EnsureNoError(err, func() error {
			return nil
		})
		if err != nil {
			return err
		}
		if val, ok := resource_tags["EventBus"]; ok && val == "true" {
			err := runner.EvaluateExpr(body["name"].Expr, &resource_topic_name, nil)
			err = runner.EnsureNoError(err, func() error {
				return nil
			})
			if err != nil {
				return err
			}

			for _, topic_name := range topics.TOPICS {
				if resource_topic_name == topic_name {
					return nil
				}
			}
			return runner.EmitIssue(r, fmt.Sprintf("Event bus topic name invalid: %s", resource_topic_name), body["name"].NameRange)
		}
		return nil
	})
}
