package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

type NodePoolConfig struct {
	Name         string `json:"name"`
	InstanceType string `json:"instanceType"`
	MinSize      int    `json:"minSize"`
	MaxSize      int    `json:"maxSize"`
}

type VpcConfig struct {
	Cidr    string `json:"cidr"`
	AzCount int    `json:"azCount"`
}

type EksConfig struct {
	Vpc               *VpcConfig       `json:"vpc,omitempty"`
	KubernetesVersion string           `json:"kubernetesVersion,omitempty"`
	NodePools         []NodePoolConfig `json:"nodePools,omitempty"`
}

// Utility function to check for errors and print to STDOUT and exit.
func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

// Generate a unique name for AWS resources.
func generateResourceName(name string, region string, resourceType string) string {
	// Convert region to a short code.
	shortRegion := strings.ReplaceAll(region, "-", "")
	// Remove letters after the last dash if you prefer just letters+number
	// For example, "eu-west-2" -> "euw2"
	// Already handled by removing dashes

	// Combine parts to form the resource name.
	return fmt.Sprintf("%s-%s-%s", name, shortRegion, resourceType)
}

func parseEksConfig(cfg *config.Config) (*EksConfig, error) {
	// TODO: Parse EKS-specific configuration.
	return &EksConfig{
		KubernetesVersion: "1.33",
		Vpc:               nil,
		NodePools:         nil,
	}, nil
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Load root config.
		cfg := config.New(ctx, "")
		// Read common required values: type, region, name.
		stackType := cfg.Get("type")
		if stackType == "" {
			// TODO: error out if no type is set.
		}
		region := cfg.Get("aws:region")
		if region == "" {
			// TODO: error out if no region is set.
		}
		name := cfg.Get("name")
		if name == "" {
			// TODO: error out if no name is set.
		}

		// Dispatch based on type.
		switch stackType {
		case "eks":
			// Parse EKS config.
			eksConfig, err := parseEksConfig(cfg)
			checkErr(err)
			fmt.Printf("EKS Config: %+v\n", eksConfig)
		default:
			return fmt.Errorf("unknown type: %s", stackType)
		}
		return nil
	})
}
