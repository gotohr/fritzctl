package cmd

import (
	"github.com/gotohr/fritzctl/logger"
	"github.com/spf13/cobra"
)

var pingCmd = &cobra.Command{
	Use:     "ping",
	Short:   "Check if the FRITZ!Box responds",
	Long:    "Attempt to contact the FRITZ!Box by trying to solve the login challenge.",
	Example: "fritzctl ping",
	RunE:    ping,
}

func init() {
	RootCmd.AddCommand(pingCmd)
}

func ping(_ *cobra.Command, _ []string) error {
	clientLogin()
	logger.Success("Success! FRITZ!Box seems to be alive!")
	return nil
}
