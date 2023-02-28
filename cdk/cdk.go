package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	awslambdago "github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type CdkStackProps struct {
	awscdk.StackProps
}

// 23.9 ldflags added to strip
// 15.57 no binary executed at all
// 79.88 Go 1.0 runtime, no stripped binary
// 65.08 Provided AL2
// 49ms Switched to Graviton
// 44ms Removing the RPC code
// 40ms Removing the reflection

func NewCdkStack(scope constructs.Construct, id string, props *CdkStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	bundlingOptions := &awslambdago.BundlingOptions{
		GoBuildFlags: &[]*string{jsii.String(`-ldflags "-s -w" -tags lambda.norpc`)},
	}
	f := awslambdago.NewGoFunction(stack, jsii.String("handler"), &awslambdago.GoFunctionProps{
		Runtime:      awslambda.Runtime_PROVIDED_AL2(),
		Architecture: awslambda.Architecture_ARM_64(),
		Entry:        jsii.String("../lambda"),
		Bundling:     bundlingOptions,
		MemorySize:   jsii.Number(1024),
		Timeout:      awscdk.Duration_Millis(jsii.Number(15000)),
		Tracing:      awslambda.Tracing_ACTIVE,
		Environment: &map[string]*string{
			"AWS_XRAY_CONTEXT_MISSING": jsii.String("IGNORE_ERROR"),
		},
	})
	fu := f.AddFunctionUrl(&awslambda.FunctionUrlOptions{
		AuthType: awslambda.FunctionUrlAuthType_NONE,
	})
	awscdk.NewCfnOutput(stack, jsii.String("apigatewayV2ExampleUrl"), &awscdk.CfnOutputProps{
		ExportName: jsii.String("apigatewayV2ExampleUrl"),
		Value:      fu.Url(),
	})

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewCdkStack(app, "ColdStartTest", &CdkStackProps{
		awscdk.StackProps{},
	})

	app.Synth(nil)
}
