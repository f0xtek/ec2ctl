// Copyright © 2019 Luke Anderson <luke.anderson@cdl.co.uk>

package cmd

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
	"log"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start an EC2 instance given its 'Name' tag.",
	Long: `Starts the given EC2 instance.
For example:

	ec2ctl start --profile default-profile --name MyInstance`,
	Run: func(cmd *cobra.Command, args []string) {
		e := createEc2Client()
		instanceID := getInstanceID(e)

		params := &ec2.StartInstancesInput{
			InstanceIds: []*string{
				aws.String(instanceID),
			},
			// Set DryRun to true to simulate the start operation. This is useful to check whether we have permission
			// to to start the instance or not.
			DryRun: aws.Bool(true),
		}
		// Simulate the start operation
		_, err := e.StartInstances(params)
		awsErr, ok := err.(awserr.Error)

		if ok && awsErr.Code() == "DryRunOperation" {
			// Set dry run to false, to allow us to actually start the instance.
			log.Println("Dry run successful. Continuing...")
			params.DryRun = aws.Bool(false)
			res, err := e.StartInstances(params)
			if err != nil {
				log.Fatal(err)
			} else {
				log.Println("Success!", res.StartingInstances)
			}
		} else {
			// This could be a permissions error
			fmt.Println("Error", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
