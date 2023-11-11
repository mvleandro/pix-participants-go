package server

// Defines the needed configuration to run the service.
type PixParticipantsServerConfig struct {

	// Cron expression with the period in which the csv file will be downloaded.
	FileDownloadTimeExpression string

	// The number of csv files to keep saved localy.
	NumberOfFilesToKeepSaved uint

	// The csv base url to download the Pix participants csv file.
	CSVFileBaseUrl string

	// The csv file name prefix.
	CSVFilePrefix string

	// Defines the minimum log level to display.
	MinimumLogLevel string

	// Timezone in Brazil.
	BrazilTimeZone string
}
