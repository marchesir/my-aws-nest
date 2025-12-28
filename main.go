package main

import (
	"fmt"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

type EksConfig struct {
	KubernetesVersion string `json:"kubernetesVersion,omitempty"`
	InstanceType      string `json:"instanceType,omitempty"`
	AutoScaling       bool   `json:"autoScaling,omitempty"`
	MinSize           int    `json:"minSize,omitempty"`
	MaxSize           int    `json:"maxSize,omitempty"`
	DesiredCapacity   int    `json:"desiredSize,omitempty"`
	VpcCidr           string `json:"vpcCidr,omitempty"`
	AzCount           int    `json:"azCount,omitempty"`
}

// Utility function to check for errors and print to STDOUT and exit.
//func checkErr(err error) {
//	if err != nil {
//		fmt.Printf("Error: %v\n", err)
//		os.Exit(1)
//	}
//}

// Generate a unique name for AWS resources.
//func generateResourceName(name string, region string, resourceType string) string {
// Convert region to a short code.
//shortRegion := strings.ReplaceAll(region, "-", "")
// Remove letters after the last dash if you prefer just letters+number
// For example, "eu-west-2" -> "euw2"
// Already handled by removing dashes

// Combine parts to form the resource name.
//return fmt.Sprintf("%s-%s-%s", name, shortRegion, resourceType)
//}

func parseEksConfig(cfg *config.Config) *EksConfig {
	kubernetesVersion := cfg.Get("kubernetesVersion")
	if kubernetesVersion == "" {
		kubernetesVersion = "1.33"
	}
	instanceType := cfg.Get("instanceType")
	if instanceType == "" {
		instanceType = "t3.medium"
	}
	autoScaling := cfg.GetBool("autoScaling")
	if !autoScaling {
		autoScaling = false
	}
	minSize := cfg.GetInt("minSize")
	if minSize == 0 {
		minSize = 1
	}
	maxSize := cfg.GetInt("maxSize")
	if maxSize == 0 {
		maxSize = 3
	}
	desiredCapacity := cfg.GetInt("desiredCapacity")
	if desiredCapacity == 0 {
		desiredCapacity = 2
	}
	vpcCidr := cfg.Get("vpcCidr")
	if vpcCidr == "" {
		vpcCidr = "10.0.0.0/20"
	}
	azCount := cfg.GetInt("azCount")
	if azCount == 0 {
		azCount = 3
	}

	return &EksConfig{
		KubernetesVersion: kubernetesVersion,
		VpcCidr:           vpcCidr,
		AzCount:           azCount,
		InstanceType:      instanceType,
		AutoScaling:       autoScaling,
		MinSize:           minSize,
		MaxSize:           maxSize,
		DesiredCapacity:   desiredCapacity,
	}
}

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Load root config.
		cfg := config.New(ctx, "")

		// Read common required values: type, region, name.
		stackType := cfg.Get("type")
		if stackType == "" {
			return fmt.Errorf("type must be specified in the config")
		}
		region := cfg.Get("aws:region")
		if region == "" {
			return fmt.Errorf("aws:region must be specified in the config")
		}
		name := cfg.Get("name")
		if name == "" {
			return fmt.Errorf("name must be specified in the config")
		}

		// Dispatch based on type.
		switch stackType {
		case "eks":
			// Parse EKS config.
			eksConfig := parseEksConfig(cfg)
			fmt.Printf("Creating EKS cluster '%v' \n", eksConfig)
		default:
			return fmt.Errorf("unknown type: %s", stackType)
		}
		return nil
	})
}
