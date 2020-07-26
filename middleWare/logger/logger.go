package logger

import (
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"time"
)

func LogWrite2File()gin.HandlerFunc{
	fileName:="D:/www/blog_go/logs/web"
	logger := log.New()
	logger.SetLevel(log.DebugLevel)
	rotateLogs, err := rotatelogs.New(
		fileName+".%Y%m%d%H.log",
		rotatelogs.WithLinkName(fileName),
		rotatelogs.WithRotationCount(3),
		rotatelogs.WithRotationTime(time.Hour),
	)
	if err != nil {
		log.Fatal("rotatelogs config failed ; err: ",err)
	}
	writerMap := lfshook.WriterMap{
		log.DebugLevel: rotateLogs,
		log.InfoLevel:  rotateLogs,
		log.WarnLevel:  rotateLogs,
		log.ErrorLevel: rotateLogs,
		log.FatalLevel: rotateLogs,
		log.PanicLevel: rotateLogs,
	}
	loghook := lfshook.NewHook(writerMap, log.Formatter(
		&log.JSONFormatter{
			TimestampFormat: "2006.01.02/03:04:05",
		}),)
	logger.AddHook(loghook)
	return func(ctx *gin.Context) {
		start:=time.Now()
		ctx.Next()
		end:=time.Now()
		logger.WithFields(log.Fields {
			"request_method"    : ctx.Request.Method,
			"request_uri"       : ctx.Request.RequestURI,
			"request_proto"     : ctx.Request.Proto,
			"request_useragent" : ctx.Request.UserAgent(),
			"request_referer"   : ctx.Request.Referer(),
			"request_post_data" : ctx.Request.PostForm.Encode(),
			"request_client_ip" : ctx.ClientIP(),
			"response_status_code" : ctx.Writer.Status(),
			"cost_time" : end.Sub(start),
		}).Info()
	}

}
