// Copyright © 2019 Luke Anderson <luke.anderson@cdl.co.uk>

package cmd

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/ec2"
	"log"

	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop an EC2 instance given its 'Name' tag.",
	Long: `Stops the given EC2 instance.
For example:

	ec2ctl stop --profile cdl-general-dev --name MyInstance`,
	Run: func(cmd *cobra.Command, args []string) {
		e := createEc2Client()
		instanceID := getInstanceID(e)

		params := &ec2.StopInstancesInput{
			InstanceIds: []*string{
				aws.String(instanceID),
			},
			// Set DryRun to true to simulate the stop operation. This is useful to check whether we have permission
			// to to start the instance or not.
			DryRun: aws.Bool(true),
		}
		// Simulate the stop operation
		_, err := e.StopInstances(params)
		awsErr, ok := err.(awserr.Error)

		if ok && awsErr.Code() == "DryRunOperation" {
			// Set dry run to false, to allow us to actually stop the instance.
			log.Println("Dry run successful. Continuing...")
			params.DryRun = aws.Bool(false)
			res, err := e.StopInstances(params)
			if err != nil {
				log.Fatal(err)
			} else {
				log.Println("Success", res.StoppingInstances)
			}
		} else {
			// This could be a permissions error
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
