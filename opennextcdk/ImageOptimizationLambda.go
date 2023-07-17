// Deploy a NextJS app using OpenNext packaging to serverless AWS using CDK
package opennextcdk

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/datasprayio/open-next-cdk/opennextcdk/jsii"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudwatch"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/datasprayio/open-next-cdk/opennextcdk/internal"
)

// This lambda handles image optimization.
type ImageOptimizationLambda interface {
	awslambda.Function
	// The architecture of this Lambda Function (this is an optional attribute and defaults to X86_64).
	Architecture() awslambda.Architecture
	Bucket() awss3.IBucket
	SetBucket(val awss3.IBucket)
	// Whether the addPermission() call adds any permissions.
	//
	// True for new Lambdas, false for version $LATEST and imported Lambdas
	// from different accounts.
	CanCreatePermissions() *bool
	// Access the Connections object.
	//
	// Will fail if not a VPC-enabled Lambda Function.
	Connections() awsec2.Connections
	// Returns a `lambda.Version` which represents the current version of this Lambda function. A new version will be created every time the function's configuration changes.
	//
	// You can specify options for this version using the `currentVersionOptions`
	// prop when initializing the `lambda.Function`.
	CurrentVersion() awslambda.Version
	// The DLQ (as queue) associated with this Lambda Function (this is an optional attribute).
	DeadLetterQueue() awssqs.IQueue
	// The DLQ (as topic) associated with this Lambda Function (this is an optional attribute).
	DeadLetterTopic() awssns.ITopic
	// The environment this resource belongs to.
	//
	// For resources that are created and managed by the CDK
	// (generally, those created by creating new class instances like Role, Bucket, etc.),
	// this is always the same as the environment of the stack they belong to;
	// however, for imported resources
	// (those obtained from static methods like fromRoleArn, fromBucketName, etc.),
	// that might be different than the stack they were imported into.
	Env() *awscdk.ResourceEnvironment
	// ARN of this function.
	FunctionArn() *string
	// Name of this function.
	FunctionName() *string
	// The principal this Lambda Function is running as.
	GrantPrincipal() awsiam.IPrincipal
	// Whether or not this Lambda function was bound to a VPC.
	//
	// If this is is `false`, trying to access the `connections` object will fail.
	IsBoundToVpc() *bool
	// The `$LATEST` version of this function.
	//
	// Note that this is reference to a non-specific AWS Lambda version, which
	// means the function this version refers to can return different results in
	// different invocations.
	//
	// To obtain a reference to an explicit version which references the current
	// function configuration, use `lambdaFunction.currentVersion` instead.
	LatestVersion() awslambda.IVersion
	// The LogGroup where the Lambda function's logs are made available.
	//
	// If either `logRetention` is set or this property is called, a CloudFormation custom resource is added to the stack that
	// pre-creates the log group as part of the stack deployment, if it already doesn't exist, and sets the correct log retention
	// period (never expire, by default).
	//
	// Further, if the log group already exists and the `logRetention` is not set, the custom resource will reset the log retention
	// to never expire even if it was configured with a different value.
	LogGroup() awslogs.ILogGroup
	// The tree node.
	Node() constructs.Node
	// The construct node where permissions are attached.
	PermissionsNode() constructs.Node
	// Returns a string-encoded token that resolves to the physical name that should be passed to the CloudFormation resource.
	//
	// This value will resolve to one of the following:
	// - a concrete value (e.g. `"my-awesome-bucket"`)
	// - `undefined`, when a name should be generated by CloudFormation
	// - a concrete name generated automatically during synthesis, in
	//   cross-environment scenarios.
	PhysicalName() *string
	// The ARN(s) to put into the resource field of the generated IAM policy for grantInvoke().
	ResourceArnsForGrantInvoke() *[]*string
	// Execution role associated with this function.
	Role() awsiam.IRole
	// The runtime configured for this lambda.
	Runtime() awslambda.Runtime
	// The stack in which this resource is defined.
	Stack() awscdk.Stack
	// The timeout configured for this lambda.
	Timeout() awscdk.Duration
	// Defines an alias for this function.
	//
	// The alias will automatically be updated to point to the latest version of
	// the function as it is being updated during a deployment.
	//
	// ```ts
	// declare const fn: lambda.Function;
	//
	// fn.addAlias('Live');
	//
	// // Is equivalent to
	//
	// new lambda.Alias(this, 'AliasLive', {
	//   aliasName: 'Live',
	//   version: fn.currentVersion,
	// });
	// ```.
	AddAlias(aliasName *string, options *awslambda.AliasOptions) awslambda.Alias
	// Adds an environment variable to this Lambda function.
	//
	// If this is a ref to a Lambda function, this operation results in a no-op.
	AddEnvironment(key *string, value *string, options *awslambda.EnvironmentOptions) awslambda.Function
	// Adds an event source to this function.
	//
	// Event sources are implemented in the @aws-cdk/aws-lambda-event-sources module.
	//
	// The following example adds an SQS Queue as an event source:
	// ```
	// import { SqsEventSource } from '@aws-cdk/aws-lambda-event-sources';
	// myFunction.addEventSource(new SqsEventSource(myQueue));
	// ```.
	AddEventSource(source awslambda.IEventSource)
	// Adds an event source that maps to this AWS Lambda function.
	AddEventSourceMapping(id *string, options *awslambda.EventSourceMappingOptions) awslambda.EventSourceMapping
	// Adds a url to this lambda function.
	AddFunctionUrl(options *awslambda.FunctionUrlOptions) awslambda.FunctionUrl
	// Adds one or more Lambda Layers to this Lambda function.
	AddLayers(layers ...awslambda.ILayerVersion)
	// Adds a permission to the Lambda resource policy.
	// See: Permission for details.
	//
	AddPermission(id *string, permission *awslambda.Permission)
	// Adds a statement to the IAM role assumed by the instance.
	AddToRolePolicy(statement awsiam.PolicyStatement)
	// Apply the given removal policy to this resource.
	//
	// The Removal Policy controls what happens to this resource when it stops
	// being managed by CloudFormation, either because you've removed it from the
	// CDK application or because you've made a change that requires the resource
	// to be replaced.
	//
	// The resource can be deleted (`RemovalPolicy.DESTROY`), or left in your AWS
	// account for data recovery and cleanup later (`RemovalPolicy.RETAIN`).
	ApplyRemovalPolicy(policy awscdk.RemovalPolicy)
	// Configures options for asynchronous invocation.
	ConfigureAsyncInvoke(options *awslambda.EventInvokeConfigOptions)
	// A warning will be added to functions under the following conditions: - permissions that include `lambda:InvokeFunction` are added to the unqualified function.
	//
	// - function.currentVersion is invoked before or after the permission is created.
	//
	// This applies only to permissions on Lambda functions, not versions or aliases.
	// This function is overridden as a noOp for QualifiedFunctionBase.
	ConsiderWarningOnInvokeFunctionPermissions(scope constructs.Construct, action *string)
	GeneratePhysicalName() *string
	// Returns an environment-sensitive token that should be used for the resource's "ARN" attribute (e.g. `bucket.bucketArn`).
	//
	// Normally, this token will resolve to `arnAttr`, but if the resource is
	// referenced across environments, `arnComponents` will be used to synthesize
	// a concrete ARN with the resource's physical name. Make sure to reference
	// `this.physicalName` in `arnComponents`.
	GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string
	// Returns an environment-sensitive token that should be used for the resource's "name" attribute (e.g. `bucket.bucketName`).
	//
	// Normally, this token will resolve to `nameAttr`, but if the resource is
	// referenced across environments, it will be resolved to `this.physicalName`,
	// which will be a concrete name.
	GetResourceNameAttribute(nameAttr *string) *string
	// Grant the given identity permissions to invoke this Lambda.
	GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant
	// Grant the given identity permissions to invoke this Lambda Function URL.
	GrantInvokeUrl(grantee awsiam.IGrantable) awsiam.Grant
	// Mix additional information into the hash of the Version object.
	//
	// The Lambda Function construct does its best to automatically create a new
	// Version when anything about the Function changes (its code, its layers,
	// any of the other properties).
	//
	// However, you can sometimes source information from places that the CDK cannot
	// look into, like the deploy-time values of SSM parameters. In those cases,
	// the CDK would not force the creation of a new Version object when it actually
	// should.
	//
	// This method can be used to invalidate the current Version object. Pass in
	// any string into this method, and make sure the string changes when you know
	// a new Version needs to be created.
	//
	// This method may be called more than once.
	InvalidateVersionBasedOn(x *string)
	// Return the given named metric for this Function.
	Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// How long execution of this Lambda takes.
	//
	// Average over 5 minutes.
	MetricDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// How many invocations of this Lambda fail.
	//
	// Sum over 5 minutes.
	MetricErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// How often this Lambda is invoked.
	//
	// Sum over 5 minutes.
	MetricInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// How often this Lambda is throttled.
	//
	// Sum over 5 minutes.
	MetricThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric
	// Returns a string representation of this construct.
	ToString() *string
	WarnInvokeFunctionPermissions(scope constructs.Construct)
}

// The jsii proxy struct for ImageOptimizationLambda
type jsiiProxy_ImageOptimizationLambda struct {
	internal.Type__awslambdaFunction
}

func (j *jsiiProxy_ImageOptimizationLambda) Architecture() awslambda.Architecture {
	var returns awslambda.Architecture
	_jsii_.Get(
		j,
		"architecture",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageOptimizationLambda) Bucket() awss3.IBucket {
	var returns awss3.IBucket
	_jsii_.Get(
		j,
		"bucket",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageOptimizationLambda) CanCreatePermissions() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"canCreatePermissions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageOptimizationLambda) Connections() awsec2.Connections {
	var returns awsec2.Connections
	_jsii_.Get(
		j,
		"connections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageOptimizationLambda) CurrentVersion() awslambda.Version {
	var returns awslambda.Version
	_jsii_.Get(
		j,
		"currentVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageOptimizationLambda) DeadLetterQueue() awssqs.IQueue {
	var returns awssqs.IQueue
	_jsii_.Get(
		j,
		"deadLetterQueue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageOptimizationLambda) DeadLetterTopic() awssns.ITopic {
	var returns awssns.ITopic
	_jsii_.Get(
		j,
		"deadLetterTopic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageOptimizationLambda) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageOptimizationLambda) FunctionArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageOptimizationLambda) FunctionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"functionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageOptimizationLambda) GrantPrincipal() awsiam.IPrincipal {
	var returns awsiam.IPrincipal
	_jsii_.Get(
		j,
		"grantPrincipal",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageOptimizationLambda) IsBoundToVpc() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"isBoundToVpc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageOptimizationLambda) LatestVersion() awslambda.IVersion {
	var returns awslambda.IVersion
	_jsii_.Get(
		j,
		"latestVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageOptimizationLambda) LogGroup() awslogs.ILogGroup {
	var returns awslogs.ILogGroup
	_jsii_.Get(
		j,
		"logGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageOptimizationLambda) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageOptimizationLambda) PermissionsNode() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"permissionsNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageOptimizationLambda) PhysicalName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"physicalName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageOptimizationLambda) ResourceArnsForGrantInvoke() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceArnsForGrantInvoke",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageOptimizationLambda) Role() awsiam.IRole {
	var returns awsiam.IRole
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageOptimizationLambda) Runtime() awslambda.Runtime {
	var returns awslambda.Runtime
	_jsii_.Get(
		j,
		"runtime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageOptimizationLambda) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ImageOptimizationLambda) Timeout() awscdk.Duration {
	var returns awscdk.Duration
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}


func NewImageOptimizationLambda(scope constructs.Construct, id *string, props *ImageOptimizationProps) ImageOptimizationLambda {
	_init_.Initialize()

	if err := validateNewImageOptimizationLambdaParameters(scope, id, props); err != nil {
		panic(err)
	}
	j := jsiiProxy_ImageOptimizationLambda{}

	_jsii_.Create(
		"open-next-cdk.ImageOptimizationLambda",
		[]interface{}{scope, id, props},
		&j,
	)

	return &j
}

func NewImageOptimizationLambda_Override(i ImageOptimizationLambda, scope constructs.Construct, id *string, props *ImageOptimizationProps) {
	_init_.Initialize()

	_jsii_.Create(
		"open-next-cdk.ImageOptimizationLambda",
		[]interface{}{scope, id, props},
		i,
	)
}

func (j *jsiiProxy_ImageOptimizationLambda)SetBucket(val awss3.IBucket) {
	if err := j.validateSetBucketParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bucket",
		val,
	)
}

// Record whether specific properties in the `AWS::Lambda::Function` resource should also be associated to the Version resource.
//
// See 'currentVersion' section in the module README for more details.
func ImageOptimizationLambda_ClassifyVersionProperty(propertyName *string, locked *bool) {
	_init_.Initialize()

	if err := validateImageOptimizationLambda_ClassifyVersionPropertyParameters(propertyName, locked); err != nil {
		panic(err)
	}
	_jsii_.StaticInvokeVoid(
		"open-next-cdk.ImageOptimizationLambda",
		"classifyVersionProperty",
		[]interface{}{propertyName, locked},
	)
}

// Import a lambda function into the CDK using its ARN.
func ImageOptimizationLambda_FromFunctionArn(scope constructs.Construct, id *string, functionArn *string) awslambda.IFunction {
	_init_.Initialize()

	if err := validateImageOptimizationLambda_FromFunctionArnParameters(scope, id, functionArn); err != nil {
		panic(err)
	}
	var returns awslambda.IFunction

	_jsii_.StaticInvoke(
		"open-next-cdk.ImageOptimizationLambda",
		"fromFunctionArn",
		[]interface{}{scope, id, functionArn},
		&returns,
	)

	return returns
}

// Creates a Lambda function object which represents a function not defined within this stack.
func ImageOptimizationLambda_FromFunctionAttributes(scope constructs.Construct, id *string, attrs *awslambda.FunctionAttributes) awslambda.IFunction {
	_init_.Initialize()

	if err := validateImageOptimizationLambda_FromFunctionAttributesParameters(scope, id, attrs); err != nil {
		panic(err)
	}
	var returns awslambda.IFunction

	_jsii_.StaticInvoke(
		"open-next-cdk.ImageOptimizationLambda",
		"fromFunctionAttributes",
		[]interface{}{scope, id, attrs},
		&returns,
	)

	return returns
}

// Import a lambda function into the CDK using its name.
func ImageOptimizationLambda_FromFunctionName(scope constructs.Construct, id *string, functionName *string) awslambda.IFunction {
	_init_.Initialize()

	if err := validateImageOptimizationLambda_FromFunctionNameParameters(scope, id, functionName); err != nil {
		panic(err)
	}
	var returns awslambda.IFunction

	_jsii_.StaticInvoke(
		"open-next-cdk.ImageOptimizationLambda",
		"fromFunctionName",
		[]interface{}{scope, id, functionName},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
// Deprecated: use `x instanceof Construct` instead.
func ImageOptimizationLambda_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateImageOptimizationLambda_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"open-next-cdk.ImageOptimizationLambda",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Returns true if the construct was created by CDK, and false otherwise.
func ImageOptimizationLambda_IsOwnedResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateImageOptimizationLambda_IsOwnedResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"open-next-cdk.ImageOptimizationLambda",
		"isOwnedResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Check whether the given construct is a Resource.
func ImageOptimizationLambda_IsResource(construct constructs.IConstruct) *bool {
	_init_.Initialize()

	if err := validateImageOptimizationLambda_IsResourceParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"open-next-cdk.ImageOptimizationLambda",
		"isResource",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

// Return the given named metric for this Lambda.
func ImageOptimizationLambda_MetricAll(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	if err := validateImageOptimizationLambda_MetricAllParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"open-next-cdk.ImageOptimizationLambda",
		"metricAll",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

// Metric for the number of concurrent executions across all Lambdas.
func ImageOptimizationLambda_MetricAllConcurrentExecutions(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	if err := validateImageOptimizationLambda_MetricAllConcurrentExecutionsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"open-next-cdk.ImageOptimizationLambda",
		"metricAllConcurrentExecutions",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the Duration executing all Lambdas.
func ImageOptimizationLambda_MetricAllDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	if err := validateImageOptimizationLambda_MetricAllDurationParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"open-next-cdk.ImageOptimizationLambda",
		"metricAllDuration",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of Errors executing all Lambdas.
func ImageOptimizationLambda_MetricAllErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	if err := validateImageOptimizationLambda_MetricAllErrorsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"open-next-cdk.ImageOptimizationLambda",
		"metricAllErrors",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of invocations of all Lambdas.
func ImageOptimizationLambda_MetricAllInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	if err := validateImageOptimizationLambda_MetricAllInvocationsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"open-next-cdk.ImageOptimizationLambda",
		"metricAllInvocations",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of throttled invocations of all Lambdas.
func ImageOptimizationLambda_MetricAllThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	if err := validateImageOptimizationLambda_MetricAllThrottlesParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"open-next-cdk.ImageOptimizationLambda",
		"metricAllThrottles",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Metric for the number of unreserved concurrent executions across all Lambdas.
func ImageOptimizationLambda_MetricAllUnreservedConcurrentExecutions(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	_init_.Initialize()

	if err := validateImageOptimizationLambda_MetricAllUnreservedConcurrentExecutionsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.StaticInvoke(
		"open-next-cdk.ImageOptimizationLambda",
		"metricAllUnreservedConcurrentExecutions",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageOptimizationLambda) AddAlias(aliasName *string, options *awslambda.AliasOptions) awslambda.Alias {
	if err := i.validateAddAliasParameters(aliasName, options); err != nil {
		panic(err)
	}
	var returns awslambda.Alias

	_jsii_.Invoke(
		i,
		"addAlias",
		[]interface{}{aliasName, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageOptimizationLambda) AddEnvironment(key *string, value *string, options *awslambda.EnvironmentOptions) awslambda.Function {
	if err := i.validateAddEnvironmentParameters(key, value, options); err != nil {
		panic(err)
	}
	var returns awslambda.Function

	_jsii_.Invoke(
		i,
		"addEnvironment",
		[]interface{}{key, value, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageOptimizationLambda) AddEventSource(source awslambda.IEventSource) {
	if err := i.validateAddEventSourceParameters(source); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addEventSource",
		[]interface{}{source},
	)
}

func (i *jsiiProxy_ImageOptimizationLambda) AddEventSourceMapping(id *string, options *awslambda.EventSourceMappingOptions) awslambda.EventSourceMapping {
	if err := i.validateAddEventSourceMappingParameters(id, options); err != nil {
		panic(err)
	}
	var returns awslambda.EventSourceMapping

	_jsii_.Invoke(
		i,
		"addEventSourceMapping",
		[]interface{}{id, options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageOptimizationLambda) AddFunctionUrl(options *awslambda.FunctionUrlOptions) awslambda.FunctionUrl {
	if err := i.validateAddFunctionUrlParameters(options); err != nil {
		panic(err)
	}
	var returns awslambda.FunctionUrl

	_jsii_.Invoke(
		i,
		"addFunctionUrl",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageOptimizationLambda) AddLayers(layers ...awslambda.ILayerVersion) {
	args := []interface{}{}
	for _, a := range layers {
		args = append(args, a)
	}

	_jsii_.InvokeVoid(
		i,
		"addLayers",
		args,
	)
}

func (i *jsiiProxy_ImageOptimizationLambda) AddPermission(id *string, permission *awslambda.Permission) {
	if err := i.validateAddPermissionParameters(id, permission); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addPermission",
		[]interface{}{id, permission},
	)
}

func (i *jsiiProxy_ImageOptimizationLambda) AddToRolePolicy(statement awsiam.PolicyStatement) {
	if err := i.validateAddToRolePolicyParameters(statement); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addToRolePolicy",
		[]interface{}{statement},
	)
}

func (i *jsiiProxy_ImageOptimizationLambda) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_ImageOptimizationLambda) ConfigureAsyncInvoke(options *awslambda.EventInvokeConfigOptions) {
	if err := i.validateConfigureAsyncInvokeParameters(options); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"configureAsyncInvoke",
		[]interface{}{options},
	)
}

func (i *jsiiProxy_ImageOptimizationLambda) ConsiderWarningOnInvokeFunctionPermissions(scope constructs.Construct, action *string) {
	if err := i.validateConsiderWarningOnInvokeFunctionPermissionsParameters(scope, action); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"considerWarningOnInvokeFunctionPermissions",
		[]interface{}{scope, action},
	)
}

func (i *jsiiProxy_ImageOptimizationLambda) GeneratePhysicalName() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"generatePhysicalName",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageOptimizationLambda) GetResourceArnAttribute(arnAttr *string, arnComponents *awscdk.ArnComponents) *string {
	if err := i.validateGetResourceArnAttributeParameters(arnAttr, arnComponents); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getResourceArnAttribute",
		[]interface{}{arnAttr, arnComponents},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageOptimizationLambda) GetResourceNameAttribute(nameAttr *string) *string {
	if err := i.validateGetResourceNameAttributeParameters(nameAttr); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getResourceNameAttribute",
		[]interface{}{nameAttr},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageOptimizationLambda) GrantInvoke(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantInvokeParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantInvoke",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageOptimizationLambda) GrantInvokeUrl(grantee awsiam.IGrantable) awsiam.Grant {
	if err := i.validateGrantInvokeUrlParameters(grantee); err != nil {
		panic(err)
	}
	var returns awsiam.Grant

	_jsii_.Invoke(
		i,
		"grantInvokeUrl",
		[]interface{}{grantee},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageOptimizationLambda) InvalidateVersionBasedOn(x *string) {
	if err := i.validateInvalidateVersionBasedOnParameters(x); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"invalidateVersionBasedOn",
		[]interface{}{x},
	)
}

func (i *jsiiProxy_ImageOptimizationLambda) Metric(metricName *string, props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricParameters(metricName, props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metric",
		[]interface{}{metricName, props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageOptimizationLambda) MetricDuration(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricDurationParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricDuration",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageOptimizationLambda) MetricErrors(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricErrorsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricErrors",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageOptimizationLambda) MetricInvocations(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricInvocationsParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricInvocations",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageOptimizationLambda) MetricThrottles(props *awscloudwatch.MetricOptions) awscloudwatch.Metric {
	if err := i.validateMetricThrottlesParameters(props); err != nil {
		panic(err)
	}
	var returns awscloudwatch.Metric

	_jsii_.Invoke(
		i,
		"metricThrottles",
		[]interface{}{props},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageOptimizationLambda) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_ImageOptimizationLambda) WarnInvokeFunctionPermissions(scope constructs.Construct) {
	if err := i.validateWarnInvokeFunctionPermissionsParameters(scope); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"warnInvokeFunctionPermissions",
		[]interface{}{scope},
	)
}

