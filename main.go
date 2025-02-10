package main

import (
	"fmt"
	"os"

	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// AWS_REGION_REGEX_STRICT = r"^(us|eu|ap|ca|sa)-(east|west|north|south|central|southeast|northeast)-\d{1}$"
// ADJECTIVES = ["Cool","Fast","Bold","Calm","Sharp","Quick","Keen","Brave","Happy","Proud"]
// NOUNS = ["Panda","Tiger","Wolf","Bear","Eagle","Fox","Deer","Owl","Lion","Hawk"]

// Utility function to check for errors and print to STDOUT and exit.
func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

// Function to create a VPC
func createVpc(ctx *pulumi.Context) *ec2.Vpc {
	vpc, err := ec2.NewVpc(ctx, "vpc", nil)
	checkErr(err)
	return vpc
}

// Main entry point for the Pulumi program.
func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		vpcConfig, _ := ctx.GetConfig("my-aws-nest:vpc")
		if vpcConfig != "" {
			createVpc(ctx)
		}
		return nil
	})
}
