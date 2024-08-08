// Code generated by paramgen. DO NOT EDIT.
// Source: github.com/ConduitIO/conduit-commons/tree/main/paramgen

package source

import (
	"github.com/conduitio/conduit-commons/config"
)

func (Config) Parameters() map[string]config.Parameter {
	return map[string]config.Parameter{
		"aws.accessKeyId": {
			Default:     "",
			Description: "amazon access key id",
			Type:        config.ParameterTypeString,
			Validations: []config.Validation{
				config.ValidationRequired{},
			},
		},
		"aws.queue": {
			Default:     "",
			Description: "amazon sqs queue name",
			Type:        config.ParameterTypeString,
			Validations: []config.Validation{
				config.ValidationRequired{},
			},
		},
		"aws.region": {
			Default:     "",
			Description: "amazon sqs region",
			Type:        config.ParameterTypeString,
			Validations: []config.Validation{
				config.ValidationRequired{},
			},
		},
		"aws.secretAccessKey": {
			Default:     "",
			Description: "amazon secret access key",
			Type:        config.ParameterTypeString,
			Validations: []config.Validation{
				config.ValidationRequired{},
			},
		},
		"aws.url": {
			Default:     "",
			Description: "aws.url is the URL for AWS (internal use only).",
			Type:        config.ParameterTypeString,
			Validations: []config.Validation{},
		},
		"aws.visibilityTimeout": {
			Default:     "0",
			Description: "visibility timeout",
			Type:        config.ParameterTypeInt,
			Validations: []config.Validation{},
		},
	}
}
