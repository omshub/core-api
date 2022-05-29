package newrelicshipper

import (
	"bytes"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

const USANewRelicEndpoint = "https://log-api.newrelic.com/log/v1"

type LogFormat struct {
	Timestamp int64  `json:"timestamp"`
	Message   string `json:"message"`
	Severity  string `json:"severity"`
}

type LogShipper struct {
	newrelicApiKey string
	httpClient     *http.Client
	log            *logrus.Logger
}

func NewLogShipHook(apiKey string) (*LogShipper, error) {
	client := http.Client{Timeout: 5 * time.Second}
	// Custom logger for the log shipper so if NewRelic is unavailable we don't try to ship the log ship errors
	// to New Relic.
	logger := logrus.New()
	logger.Out = os.Stdout
	logger.SetLevel(logrus.InfoLevel)
	return &LogShipper{
		newrelicApiKey: apiKey,
		httpClient:     &client,
		log:            logger,
	}, nil
}

func (c *LogShipper) ShipLog(severity string, logString string) error {
	msg := LogFormat{
		Message:   logString,
		Severity:  severity,
		Timestamp: time.Now().UnixMilli(),
	}
	payload, err := json.Marshal(msg)
	if err != nil {
		c.log.Errorf("error marshaling log: %s", err)
		return err
	}
	req, err := http.NewRequest(http.MethodPost, USANewRelicEndpoint, bytes.NewBuffer(payload))
	if err != nil {
		c.log.Errorf("error shipping logs: %s", err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Api-Key", c.newrelicApiKey)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		c.log.Errorf("bad response from new relic: %s, status: %d", err, resp.StatusCode)
		return err
	}
	return nil
}

// Implement Logrus Hook

func (c *LogShipper) Fire(entry *logrus.Entry) error {
	logMsg, err := entry.String()
	if err != nil {
		c.log.Errorf("bad log entry: %s", err)
		return err
	}
	return c.ShipLog(entry.Level.String(), logMsg)
}

func (c *LogShipper) Levels() []logrus.Level {
	// Send everything except Debug and Trace
	return []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
		logrus.WarnLevel,
		logrus.InfoLevel,
	}
}
