package awssns

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/arn"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/pkg/errors"

	"github.com/batchcorp/plumber-schemas/build/go/protos/opts"

	"github.com/batchcorp/plumber/dynamic"
)

func (a *AWSSNS) Dynamic(ctx context.Context, dynamicOpts *opts.DynamicOptions) error {
	if err := validateDynamicOptions(dynamicOpts); err != nil {
		return errors.Wrap(err, "unable to validate dynamic options")
	}

	// Start up dynamic connection
	grpc, err := dynamic.New(dynamicOpts, "AWS SNS")
	if err != nil {
		return errors.Wrap(err, "could not establish connection to Batch")
	}

	go grpc.Start()

	topic := dynamicOpts.Awssns.Args.Topic

	// Continually loop looking for messages on the channel.
	for {
		select {
		case outbound := <-grpc.OutboundMessageCh:
			_, err := a.Service.Publish(&sns.PublishInput{
				Message:  aws.String(string(outbound.Blob)),
				TopicArn: aws.String(topic),
			})
			if err != nil {
				a.log.Errorf("Unable to replay message: %s", err)
				break
			}

			a.log.Debugf("Replayed message to AWSSNS topic '%s' for replay '%s'", topic, outbound.ReplayId)
		}
	}

	return nil
}

func validateDynamicOptions(dynamicOpts *opts.DynamicOptions) error {
	if dynamicOpts == nil {
		return errors.New("write options cannot be nil")
	}

	if dynamicOpts.Awssns == nil {
		return errors.New("backend group options cannot be nil")
	}

	if dynamicOpts.Awssns.Args == nil {
		return errors.New("backend arg options cannot be nil")
	}

	topic := dynamicOpts.Awssns.Args.Topic

	if topic == "" {
		return ErrMissingTopicARN
	}

	if arn.IsARN(topic) == false {
		return fmt.Errorf("'%s' is not a valid ARN", topic)
	}
	return nil
}