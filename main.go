package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"gopkg.in/yaml.v2"
)

const (
	AWS_REGION_REGEX = `^(us|eu|ap|ca|sa)-(east|west|north|south|central|southeast|northeast)-\d{1}$`
	ADJECTIVES       = `["Cool", "Fast", "Bold", "Calm", "Sharp", "Quick", "Keen", "Brave", "Happy", "Proud"]`
	NOUNS            = `["Panda", "Tiger", "Wolf", "Bear", "Eagle", "Fox", "Deer", "Owl", "Lion", "Hawk"]`
)

type subnet struct {
	CidrBlock string `json:"cidrBlock"`
	Type      string `json:"type"` // "private" or "public"
}

type vpc struct {
	CidrBlock string   `json:"cidrBlock"`
	Subnets   []subnet `json:"subnets"`
}

type vpcConfig struct {
	VPCs []vpc `json:"vpcs" yaml:"vpcs"`
}

// Utility function to check for errors and print to STDOUT and exit.
func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

// Generate a unique name for AWS resources
func generateResourceName(resourceType, appName string) string {
	// Seed the random generator.
	rand.Seed(time.Now().UnixNano())

	// Pick a random adjective and noun.
	adjective := ADJECTIVES[rand.Intn(len(ADJECTIVES))]
	noun := NOUNS[rand.Intn(len(NOUNS))]

	// Generate a random suffix for uniqueness
	randomSuffix := rand.Intn(10000)

	// Combine parts to form the resource name.
	return fmt.Sprintf("%s-%s%s-%s-%04d", resourceType, adjective, noun, randomSuffix)
}

// Parse the VPC configuration from the Pulumi.vpc stack YAML.
func parseVPCConfig(ctx *pulumi.Context) (*vpcConfig, error) {
	// Retrieve the raw configuration as a string.
	rawConfig, _ := ctx.GetConfig("my-aws-nest:vpc")
	if rawConfig == "" {
		return nil, fmt.Errorf("Configuration not found for key: my-aws-nest:vpc")
	}

	// Unmarshal the YAML into the vpcConfig struct.
	var config vpcConfig
	err := yaml.Unmarshal([]byte(rawConfig), &config)
	checkErr(err)

	return &config, nil
}

func createVPCs(ctx *pulumi.Context, config *vpcConfig) error {
	for _, v := range config.VPCs {
		// Create the VPC.
		vpc, err := ec2.NewVpc(
			ctx,
			"vpc-TODO",
			&ec2.VpcArgs{
				CidrBlock: pulumi.String(v.CidrBlock),
				Tags: pulumi.StringMap{
					"Name": pulumi.String("vpc-TODO"),
				},
			},
		)
		checkErr(err)

		// Create subnets for the VPC.
		for _, s := range v.Subnets {
			_, err := ec2.NewSubnet(
				ctx,
				"subnet-TODO",
				&ec2.SubnetArgs{
					CidrBlock:           pulumi.String(s.CidrBlock),
					VpcId:               vpc.ID(),
					MapPublicIpOnLaunch: pulumi.Bool(s.Type == "public"),
					Tags: pulumi.StringMap{
						"Name": pulumi.String("subnet-TODO"),
					},
				},
			)
			checkErr(err)
		}
	}
	return nil
}

// Main entry point for the Pulumi program.
func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Parse the VPC configuration.
		config, err := parseVPCConfig(ctx)
		checkErr(err)

		// Create the VPCs and any linked resources.
		err = createVPCs(ctx, config)
		checkErr(err)

		return nil
	})
}
