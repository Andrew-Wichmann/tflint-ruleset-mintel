package rules

import (
	"testing"

	hcl "github.com/hashicorp/hcl/v2"
	"github.com/terraform-linters/tflint-plugin-sdk/helper"
)

func Test_AwsInstanceExampleType(t *testing.T) {
	cases := []struct {
		Name     string
		Content  string
		Expected helper.Issues
	}{
		{
			Name: "invalid event bus topic",
			Content: `
resource "aws_sns_topic" "topic1" {
    name = "not-a-event-bus-topic"
	tags = {
		EventBus: "true"
	}
}`,
			Expected: helper.Issues{
				{
					Rule:    NewEventBusTopicNameRule(),
					Message: "Event bus topic name invalid: not-a-event-bus-topic",
					Range: hcl.Range{
						Filename: "resource.tf",
						Start:    hcl.Pos{Line: 3, Column: 5},
						End:      hcl.Pos{Line: 3, Column: 9},
					},
				},
			},
		},
		{
			Name: "do not check non event bus topic",
			Content: `
resource "aws_sns_topic" "topic1" {
    name = "not-a-event-bus-topic"
}`,
			Expected: helper.Issues{},
		},
		{
			Name: "valid event bus topic name",
			Content: `
resource "aws_sns_topic" "topic1" {
    name = "example"
	tags = {
		EventBus: "true"
	}
}`,
			Expected: helper.Issues{},
		},
	}

	rule := NewEventBusTopicNameRule()

	for _, tc := range cases {
		runner := helper.TestRunner(t, map[string]string{"resource.tf": tc.Content})

		if err := rule.Check(runner); err != nil {
			t.Fatalf("Unexpected error occurred: %s", err)
		}

		helper.AssertIssues(t, tc.Expected, runner.Issues)
	}
}
