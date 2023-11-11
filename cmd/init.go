/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"flag"

	"github.com/mvleandro/pix-participants/server"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var config server.PixParticipantsServerConfig

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Starts the service",
	Long:  `Starts the service by downloading the Pix participants CSV file from Brazilian Central Bank official endpoint and makes it available for consultation`,
	Run: func(cmd *cobra.Command, args []string) {
		var server server.PixParticipantsServer
		server.Start(config)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	var fileDownloadTimeExpression string
	var csvFileBaseUrl string
	var csvFilePrefix string
	var numberOfFilesToKeepSaved uint
	var minimumLogLevel string
	var brazilTimeZone string

	flag.StringVar(&fileDownloadTimeExpression, "fileDownloadTimeExpression", "", "Cron expression with the period in which the csv file will be downloaded.")
	flag.UintVar(&numberOfFilesToKeepSaved, "numberOfFilesToKeepSaved", 3, "The number of csv files to keep saved localy.")
	flag.StringVar(&csvFileBaseUrl, "csvFileBaseUrl", "https://www.bcb.gov.br/content/estabilidadefinanceira/spi/", "The csv base url to download the Pix participants csv file.")
	flag.StringVar(&csvFilePrefix, "csvFilePrefix", "participantes-spi-", "The csv file name prefix.")
	flag.StringVar(&minimumLogLevel, "minimumLogLevel", "info", "Defines the minimum log level to display.")
	flag.StringVar(&brazilTimeZone, "brazilTimeZone", "America/Sao_Paulo", "Timezone in Brazil.")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	config = server.PixParticipantsServerConfig{
		FileDownloadTimeExpression: fileDownloadTimeExpression,
		NumberOfFilesToKeepSaved:   numberOfFilesToKeepSaved,
		CSVFileBaseUrl:             csvFileBaseUrl,
		CSVFilePrefix:              csvFilePrefix,
		MinimumLogLevel:            minimumLogLevel,
		BrazilTimeZone:             brazilTimeZone,
	}

}
