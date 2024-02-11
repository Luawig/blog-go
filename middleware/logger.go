package middleware

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

func Logger() gin.HandlerFunc {
	logger := logrus.New()

	logger.Formatter = &logrus.JSONFormatter{}
	logger.SetLevel(logrus.DebugLevel)

	logPath := "log"
	logFileName := "blog-go"
	fileName := path.Join(logPath, logFileName)

	_ = os.MkdirAll(logPath, 0755)
	writer, err := rotatelogs.New(
		fileName+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(fileName),
		rotatelogs.WithMaxAge(time.Duration(168)*time.Hour),
		rotatelogs.WithRotationTime(time.Duration(24)*time.Hour),
	)
	if err != nil {
		logger.Errorf("config local file system logger error. %+v", err)
	}
	logger.AddHook(lfshook.NewHook(lfshook.WriterMap{
		logrus.InfoLevel:  writer,
		logrus.FatalLevel: writer,
		logrus.DebugLevel: writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.PanicLevel: writer,
	}, &logrus.JSONFormatter{}))

	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		latencyTime := time.Since(startTime).Milliseconds()
		statusCode := c.Writer.Status()
		fields := logger.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": fmt.Sprintf("%dms", latencyTime),
			"client_ip":    c.ClientIP(),
			"req_method":   c.Request.Method,
			"req_uri":      c.Request.RequestURI,
		})
		if len(c.Errors) > 0 {
			fields.Error(c.Errors.String())
			return
		}
		switch {
		case statusCode >= 500:
			fields.Error("Error Log")
		case statusCode >= 400:
			fields.Warn("Warn Log")
		default:
			fields.Info("Request Log")
		}
	}
}
