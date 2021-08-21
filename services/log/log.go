package log

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"time"
)

var (
	logFilePath = "./"
	logFileName = "log"
	Log = logrus.New()
)

func InitLog() {
	fileName := path.Join(logFilePath, logFileName)
	var src *os.File
	var err error
	if _, err:=os.Stat(fileName); os.IsNotExist(err){
		src,err=os.Create(fileName)
	} else {
		src, err=os.OpenFile(fileName, os.O_APPEND | os.O_WRONLY, os.ModeAppend)
	}
	if err != nil {
		fmt.Println("init log err !")
		panic(err)
	}

	Log.SetLevel(logrus.TraceLevel)

	mw:=io.MultiWriter(src,os.Stdout)

	Log.Out = mw

	logWriterL1, err := rotatelogs.New(
		// 分割后的文件名称
		fileName + ".%Y-%m-%d.L1.log",

		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileName),

		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	logWriterL2, err := rotatelogs.New(
		// 分割后的文件名称
		fileName + ".%Y-%m-%d.L2.log",

		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileName),

		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	logWriterL3, err := rotatelogs.New(
		// 分割后的文件名称
		fileName + ".%Y-%m-%d.L3.log",

		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileName),

		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriterL1,
		logrus.DebugLevel: logWriterL1,
		logrus.WarnLevel:  logWriterL2,
		logrus.FatalLevel: logWriterL3,
		logrus.ErrorLevel: logWriterL3,
		logrus.PanicLevel: logWriterL3,
	}

	Log.AddHook(lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}))

	Log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat:   "2006-01-02 15:04:05",
	})
}
