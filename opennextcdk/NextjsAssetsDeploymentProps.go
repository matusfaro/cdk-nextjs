// Deploy a NextJS app using OpenNEXT packaging to serverless AWS using CDK
package opennextcdk

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

type NextjsAssetsDeploymentProps struct {
	// Relative path to the directory where the NextJS project is located.
	//
	// Can be the root of your project (`.`) or a subdirectory (`packages/web`).
	NextjsPath *string `field:"required" json:"nextjsPath" yaml:"nextjsPath"`
	// Optional value used to install NextJS node dependencies.
	//
	// It defaults to 'npx --yes open-next@latest build'.
	BuildCommand *string `field:"optional" json:"buildCommand" yaml:"buildCommand"`
	// The directory to execute `npm run build` from.
	//
	// By default, it is `nextjsPath`.
	// Can be overridden, particularly useful for monorepos where `build` is expected to run
	// at the root of the project.
	BuildPath *string `field:"optional" json:"buildPath" yaml:"buildPath"`
	// 0 - no compression, fastest 9 - maximum compression, slowest.
	CompressionLevel *float64 `field:"optional" json:"compressionLevel" yaml:"compressionLevel"`
	// Custom environment variables to pass to the NextJS build and runtime.
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
	// Skip building app and deploy a placeholder.
	//
	// Useful when using `next dev` for local development.
	IsPlaceholder *bool `field:"optional" json:"isPlaceholder" yaml:"isPlaceholder"`
	// Optional value for NODE_ENV during build and runtime.
	NodeEnv *string `field:"optional" json:"nodeEnv" yaml:"nodeEnv"`
	// Root of your project, if different from `nextjsPath`.
	//
	// Defaults to current working directory.
	ProjectRoot *string `field:"optional" json:"projectRoot" yaml:"projectRoot"`
	// Less build output.
	Quiet *bool `field:"optional" json:"quiet" yaml:"quiet"`
	// Optional arn for the sharp lambda layer.
	//
	// If omitted, the layer will be created.
	SharpLayerArn *string `field:"optional" json:"sharpLayerArn" yaml:"sharpLayerArn"`
	// Directory to store temporary build files in.
	//
	// Defaults to os.tmpdir().
	TempBuildDir *string `field:"optional" json:"tempBuildDir" yaml:"tempBuildDir"`
	// Properties for the S3 bucket containing the NextJS assets.
	Bucket awss3.IBucket `field:"required" json:"bucket" yaml:"bucket"`
	// The `NextjsBuild` instance representing the built Nextjs application.
	NextBuild NextjsBuild `field:"required" json:"nextBuild" yaml:"nextBuild"`
	// Override the default S3 cache policies created internally.
	CachePolicies *NextjsAssetsCachePolicyProps `field:"optional" json:"cachePolicies" yaml:"cachePolicies"`
	// Distribution to invalidate when assets change.
	Distribution awscloudfront.IDistribution `field:"optional" json:"distribution" yaml:"distribution"`
	// ephemeralStorageSize for lambda function which been run by BucketDeployment.
	EphemeralStorageSize awscdk.Size `field:"optional" json:"ephemeralStorageSize" yaml:"ephemeralStorageSize"`
	// memoryLimit for lambda function which been run by BucketDeployment.
	MemoryLimit *float64 `field:"optional" json:"memoryLimit" yaml:"memoryLimit"`
	// Set to true to delete old assets (defaults to false).
	//
	// Recommended to only set to true if you don't need the ability to roll back deployments.
	Prune *bool `field:"optional" json:"prune" yaml:"prune"`
	// In case of useEfs, vpc is required.
	UseEfs *bool `field:"optional" json:"useEfs" yaml:"useEfs"`
	// In case of useEfs, vpc is required.
	Vpc awsec2.IVpc `field:"optional" json:"vpc" yaml:"vpc"`
}

