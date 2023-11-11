package server

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Defines the needed configuration to run the service.
type PixParticipantsServer struct {

	// Struct with server configuration.
	config PixParticipantsServerConfig

	// A slice with the current Pix participants.
	pixParticipants []Participant

	// The logger object.
	logger *zap.SugaredLogger

	// The logger atomic level.
	atom zap.AtomicLevel
}

func init() {

}

// Initialize the server.
func (s *PixParticipantsServer) Start(config PixParticipantsServerConfig) {
	s.initLog(config)

	fileName := s.getCurrentCSVFileName(config.CSVFilePrefix)
	url := s.getCSVUrl(fileName)
	s.downloadFile(url, fileName)
}

// Saves the contents present at the url in the fileName.
func (s *PixParticipantsServer) downloadFile(url string, fileName string) {
	s.logger.Infof("Downloading file %s from url %s", fileName, url)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		s.logger.Error(err)
		return
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		s.logger.Error(err)
		return
	}

	if resp.StatusCode != http.StatusOK {
		s.logger.Errorf("Error: %s", resp.Status)
		return
	}

	f, err := os.Create(fileName)
	if err != nil {
		s.logger.Error(err)
		return
	}

	defer f.Close()

	io.Copy(f, resp.Body)

	s.logger.Infof("File %s downloaded from url %s successfully.", fileName, url)
}

// Get the Pix Participants csv file name for the current day in Brazil TimeZone.
func (s *PixParticipantsServer) getCurrentCSVFileName(csvFilePrefix string) string {
	loc, err := time.LoadLocation(s.config.BrazilTimeZone)
	if err != nil {
		s.logger.Error(err)
		panic(err)
	}

	now := time.Now().In(loc)
	today := now.Format("20060102")
	fileName := fmt.Sprintf("%s%s.csv", csvFilePrefix, today)
	return fileName
}

// Get the Pix participants csv file url.
func (s *PixParticipantsServer) getCSVUrl(csvFileName string) string {
	url := fmt.Sprintf("%s%s", s.config.CSVFileBaseUrl, csvFileName)
	return url
}

// Initializes the logger.
func (s *PixParticipantsServer) initLog(config PixParticipantsServerConfig) {
	s.config = config

	s.atom = zap.NewAtomicLevel()
	var encoderCfg zapcore.EncoderConfig
	encoderCfg = zap.NewProductionEncoderConfig()

	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)

	logger := zap.New(zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.Lock(os.Stdout),
		s.atom,
	))
	s.logger = logger.Sugar()

	s.setLogLevel(s.config.MinimumLogLevel)
	defer logger.Sync()
}

// Changes the minimum log level.
func (s *PixParticipantsServer) setLogLevel(logLevel string) {
	level, err := zapcore.ParseLevel(logLevel)

	if err == nil {
		s.atom.SetLevel(level)
	} else {
		s.logger.Error(err)
	}
}
