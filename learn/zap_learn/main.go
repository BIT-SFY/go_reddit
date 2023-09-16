/*
 * @Author: BIT-SFY
 * @Date: 2023-09-14 09:31:49
 * @LastEditTime: 2023-09-14 13:53:29
 * @Description: zap的学习以及zap与gin框架的结合
 */
package main

import (
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.Logger
var sugarLogger *zap.SugaredLogger

func InitLogger() {
	writeSyncer := getLogWriter() //WriterSyncer ：指定日志将写到哪里去。
	encoder := getEncoder()       //Encoder:编码器(如何写入日志)
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)
	//使用zap.New(…)方法来手动传递所有配置，而不是使用像zap.NewProduction()这样的预置方法来创建logger。
	//AddCaller()添加将调用函数信息记录到日志中的功能
	logger = zap.New(core, zap.AddCaller())
	sugarLogger = logger.Sugar() //调用此方法，实现Sugared Logger
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder   //时间非人类可读，所以要修改时间编码器
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder //在日志文件中使用大写字母记录日志级别
	return zapcore.NewConsoleEncoder(encoderConfig)
	// return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()) //使用预先设置的ProductionEncoderConfig()
	// return zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig()) //将编码器从JSON Encoder更改为普通Encoder
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./test.log", //日志文件的位置
		MaxSize:    1,            //在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: 5,            //保留旧文件的最大个数,超出这个文件数就会删除最新产生的文件
		MaxAge:     30,           //保留旧文件的最大天数
		Compress:   false,        //是否压缩/归档旧文件
	}
	return zapcore.AddSync(lumberJackLogger)
}

func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		sugarLogger.Error(
			"Error fetching url..",
			zap.String("url", url),
			zap.Error(err))
	} else {
		sugarLogger.Info(
			"Success..",
			zap.String("statusCode", resp.Status),
			zap.String("url", url))
		resp.Body.Close()
	}
}

func main() {
	InitLogger()
	defer sugarLogger.Sync()
	// simpleHttpGet("996.icu")
	// simpleHttpGet("http://www.baidu.com")
	r := gin.New()
	r.Use(GinLogger(logger), GinRecovery(logger, true)) //使用我们自己的中间件
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello reddit!")
	})
	r.Run()
}

// GinLogger 接收gin框架默认的日志
func GinLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()             //开始时间
		path := c.Request.URL.Path      //请求路径
		query := c.Request.URL.RawQuery //请求参数
		c.Next()                        //去执行后面的中间件

		cost := time.Since(start)
		logger.Info(path,
			zap.Int("status", c.Writer.Status()),   //状态码
			zap.String("method", c.Request.Method), //方法
			zap.String("path", path),               //路径
			zap.String("query", query),             //参数
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost), //计算该请求所消耗的时间
		)
	}
}

// GinRecovery recover掉项目可能出现的panic
func GinRecovery(logger *zap.Logger, stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					logger.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
