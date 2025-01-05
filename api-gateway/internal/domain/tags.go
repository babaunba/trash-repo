package domain

import (
	"go.temporal.io/sdk/workflow"

	tags "github.com/babaunba/project-management/api-gateway/gen/tags/v1"
)

// GenerateTagsWF is not implemented
// TODO: Implement once ML service is finished
func (domain *Domain) GenerateTagsWF(workflow.Context, *tags.GenerateTagsRequest) (resp *tags.GenerateTagsResponse, err error) {
	resp = &tags.GenerateTagsResponse{Tags: []string{"bug", "shit"}}
	return
}
