package main

import (
	"fmt"
	"time"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create an AWS resource (S3 Bucket).

		// Append current time in secs to make bucket name unique.
		uniqueBucket := fmt.Sprintf("my-bucket-%d", time.Now().Unix())

		// Invoke pulumi S3 API to create the bucket.
		bucket, err := s3.NewBucket(ctx, uniqueBucket, nil)
		if err != nil {
			return err
		}

		// Export the name of the bucket
		ctx.Export("bucketName", bucket.ID())
		return nil
	})
}
