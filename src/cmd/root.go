package cmd

import "github.com/spf13/cobra"

var (
	rootCmd = &cobra.Command{
		Use:   "cobra-cli",
		Short: "A generator for Cobra based Applications",
		Long:  `Cobra is a CLI library for Go that empowers applications.`,
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {

	/*
	 * WEBSERVER SERVICE
	 */
	rootCmd.AddCommand(serverCmd)

	/**
	 * RABBITMQ SERVICE
	 */
	rootCmd.AddCommand(consumerMOCmd)
	rootCmd.AddCommand(consumerDRCmd)
	rootCmd.AddCommand(consumerRenewalCmd)
	rootCmd.AddCommand(consumerRetryCmd)

	rootCmd.AddCommand(publisherRenewalCmd)
	rootCmd.AddCommand(publisherRetryCmd)
	// rootCmd.AddCommand(publisherPurgeCmd)

}
