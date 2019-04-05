// Copyright © 2019 Luke Anderson <luke.anderson@cdl.co.uk>

package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ec2ctl",
	Short: "Manage the power state for an EC2 instance",
	Long: `Manages the power state for an EC2 instance, given its 'Name' tag.
For example:

	ec2ctl start --name MyInstance`,
}

// Name is the value of the EC2 instance 'Name' tag
var Name string

// Profile is the AWS CLI profile used for authentication to AWS
var Profile string

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&Name, "name", "n", "", "Specify the value for the EC2 'Name' tag")
	rootCmd.PersistentFlags().StringVarP(&Profile, "profile", "p", "default", "Specify the AWS CLI profile to use for authentication.")
}

// initConfig reads in ENV variables if set.
func initConfig() {
	viper.AutomaticEnv() // read in environment variables that match
	viper.SetEnvPrefix("aws")
}

// createEc2Client creates & returns a session to the AWS EC2 SDK for use in other functions
func createEc2Client() *ec2.EC2 {
	sess, err := session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Profile:           Profile,
	})
	if err != nil {
		log.Fatal("Error creating AWS session: ", err)
	}

	ec2Svc := ec2.New(sess)
	return ec2Svc
}

// getInstanceID return the instance ID for the EC2 instance, specified by the 'Name' tag
func getInstanceID(e *ec2.EC2) string {
	params := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("tag:Name"),
				Values: []*string{
					aws.String(Name),
				},
			},
		},
	}

	res, err := e.DescribeInstances(params)
	if err != nil {
		log.Fatal(err)
	}

	var id string
	for _, i := range res.Reservations[0].Instances {
		id = *i.InstanceId
	}

	return id
}
