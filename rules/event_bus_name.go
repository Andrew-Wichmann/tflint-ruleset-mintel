package rules

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/terraform/configs"
	"github.com/terraform-linters/tflint-plugin-sdk/tflint"
	"github.com/zclconf/go-cty/cty"
	"gitlab.com/mintel/everest/event-bus/events/mintel-events-go/topics"
)

// EventBusTopicNameRule checks whether ...
type EventBusTopicNameRule struct{}

// NewEventBusTopicNameRule returns a new rule
func NewEventBusTopicNameRule() *EventBusTopicNameRule {
	return &EventBusTopicNameRule{}
}

// Name returns the rule name
func (r *EventBusTopicNameRule) Name() string {
	return "event_bus_topic_naming"
}

// Enabled returns whether the rule is enabled by default
func (r *EventBusTopicNameRule) Enabled() bool {
	return true
}

// Severity returns the rule severity
func (r *EventBusTopicNameRule) Severity() string {
	return tflint.ERROR
}

// Link returns the rule reference link
func (r *EventBusTopicNameRule) Link() string {
	return ""
}

type eventBusTag struct {
	EventBus string   `cty:"EventBus"`
}

// Checks whether the event bus topic name matches a topic in the event bus.
func (r *EventBusTopicNameRule) Check(runner tflint.Runner) error {
	return runner.WalkResources("aws_sns_topic", func(resource *configs.Resource) error {
		var body hcl.Attributes
		var resource_topic_name string
		wantType := cty.Object(map[string]cty.Type{
			"EventBus":                 cty.String,
		})
		var resource_tags eventBusTag
		body, _ = resource.Config.JustAttributes()

		tags, ok := body["tags"]
		if !ok {
			return nil
		}
		err := runner.EvaluateExpr(tags.Expr, &resource_tags, &wantType)
		err = runner.EnsureNoError(err, func() error {
			return nil
		})
		if err != nil {
			return err
		}
		if resource_tags.EventBus == "true" {
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
