// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets information about a specified configuration for DNS query logging. For more
// information about DNS query logs, see CreateQueryLoggingConfig
// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_CreateQueryLoggingConfig.html)
// and Logging DNS Queries
// (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/query-logs.html).
func (c *Client) GetQueryLoggingConfig(ctx context.Context, params *GetQueryLoggingConfigInput, optFns ...func(*Options)) (*GetQueryLoggingConfigOutput, error) {
	if params == nil {
		params = &GetQueryLoggingConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetQueryLoggingConfig", params, optFns, addOperationGetQueryLoggingConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetQueryLoggingConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetQueryLoggingConfigInput struct {

	// The ID of the configuration for DNS query logging that you want to get
	// information about.
	//
	// This member is required.
	Id *string
}

type GetQueryLoggingConfigOutput struct {

	// A complex type that contains information about the query logging configuration
	// that you specified in a GetQueryLoggingConfig
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_GetQueryLoggingConfig.html)
	// request.
	//
	// This member is required.
	QueryLoggingConfig *types.QueryLoggingConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetQueryLoggingConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetQueryLoggingConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetQueryLoggingConfig{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetQueryLoggingConfigValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetQueryLoggingConfig(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetQueryLoggingConfig(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "GetQueryLoggingConfig",
	}
}
