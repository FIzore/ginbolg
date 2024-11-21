package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	retalog "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"math"
	"os"
	"time"
)

func Logger() gin.HandlerFunc {
	filepath := "log/ginblog.log"
	linkName := "latest_log.log"
	src, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("err:", err)
	}
	logger := logrus.New()
	logger.Out = src
	logger.SetLevel(logrus.DebugLevel)
	logWriter, _ := retalog.New(
		filepath+"%Y%m%d.log",
		retalog.WithMaxAge(7*24*time.Hour),
		retalog.WithRotationTime(24*time.Hour),
		retalog.WithLinkName(linkName),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	Hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logger.AddHook(Hook)
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		stopTime := time.Since(startTime)
		spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds())/1000000.0)))
		hostName, err := os.Hostname() //获取主机名
		if err != nil {
			hostName = "unknown"
		}
		StatusCode := c.Writer.Status()    //获取状态码
		clientIp := c.ClientIP()           //获取客户端ip
		userAgent := c.Request.UserAgent() //获取客户端信息
		dataSize := c.Writer.Size()        //获取数据大小
		if dataSize <= 0 {
			dataSize = 0
		}
		method := c.Request.Method //获取请求方法
		path := c.Request.URL.Path //获取请求路径
		entry := logger.WithFields(logrus.Fields{
			"HoseName":  hostName,
			"Status":    StatusCode,
			"SpendTime": spendTime,
			"Ip":        clientIp,
			"Method":    method,
			"Path":      path,
			"Agent":     userAgent,
			"DataSize":  dataSize,
		})
		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String()) //记录错误信息
		}
		if StatusCode >= 500 {
			entry.Error()
		} else if StatusCode >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}
	}
}
