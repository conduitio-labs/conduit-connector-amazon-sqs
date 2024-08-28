// Code generated by paramgen. DO NOT EDIT.
// Source: github.com/ConduitIO/conduit-commons/tree/main/paramgen

package source

import (
	"github.com/conduitio/conduit-commons/config"
)

const (
	ConfigAwsAccessKeyId       = "aws.accessKeyId"
	ConfigAwsQueue             = "aws.queue"
	ConfigAwsRegion            = "aws.region"
	ConfigAwsSecretAccessKey   = "aws.secretAccessKey"
	ConfigAwsUrl               = "aws.url"
	ConfigAwsVisibilityTimeout = "aws.visibilityTimeout"
	ConfigAwsWaitTimeSeconds   = "aws.waitTimeSeconds"
)

func (Config) Parameters() map[string]config.Parameter {
	return map[string]config.Parameter{
		ConfigAwsAccessKeyId: {
			Default:     "",
			Description: "AWSAccessKeyID is the amazon access key id",
			Type:        config.ParameterTypeString,
			Validations: []config.Validation{
				config.ValidationRequired{},
			},
		},
		ConfigAwsQueue: {
			Default:     "",
			Description: "QueueName is the sqs queue name",
			Type:        config.ParameterTypeString,
			Validations: []config.Validation{
				config.ValidationRequired{},
			},
		},
		ConfigAwsRegion: {
			Default:     "",
			Description: "AWSRegion is the amazon sqs region",
			Type:        config.ParameterTypeString,
			Validations: []config.Validation{
				config.ValidationRequired{},
			},
		},
		ConfigAwsSecretAccessKey: {
			Default:     "",
			Description: "AWSSecretAccessKey is the amazon secret access key",
			Type:        config.ParameterTypeString,
			Validations: []config.Validation{
				config.ValidationRequired{},
			},
		},
		ConfigAwsUrl: {
			Default:     "",
			Description: "AWSURL is the URL for AWS (internal use only).",
			Type:        config.ParameterTypeString,
			Validations: []config.Validation{},
		},
		ConfigAwsVisibilityTimeout: {
			Default:     "",
			Description: "VisibilityTimeout is the duration (in seconds) that the received messages\nare hidden from subsequent reads after being retrieved.",
			Type:        config.ParameterTypeInt,
			Validations: []config.Validation{},
		},
		ConfigAwsWaitTimeSeconds: {
			Default:     "10",
			Description: "WaitTimeSeconds is the duration (in seconds) for which the call waits for\na message to arrive in the queue before returning.",
			Type:        config.ParameterTypeInt,
			Validations: []config.Validation{},
		},
	}
}
