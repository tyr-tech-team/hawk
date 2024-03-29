package iris

import (
	"fmt"
	"runtime"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/sirupsen/logrus"
	"github.com/tyr-tech-team/hawk/response"
	"github.com/tyr-tech-team/hawk/status"
	"go.uber.org/zap"
)

// Recover -
func Recover() context.Handler {
	return func(ctx iris.Context) {
		defer func() {
			if err := recover(); err != nil {
				if ctx.IsStopped() {
					return
				}

				var stacktrace string
				for i := 1; ; i++ {
					_, f, l, got := runtime.Caller(i)
					if !got {
						break
					}

					stacktrace += fmt.Sprintf("%s:%d\n", f, l)
				}

				// when stack finishes
				logMessage := fmt.Sprintf("Recovered from a route's Handler('%s')\n", ctx.HandlerName())
				logMessage += fmt.Sprintf("Trace: %s\n", err)
				logMessage += fmt.Sprintf("\n%s", stacktrace)
				ctx.Application().Logger().Warn(logMessage)

				ctx.StatusCode(500)
				ctx.StopExecution()
				ctx.JSON(response.Error(ctx.Request().Context(), status.UnKnownError.Err()))
			}
		}()

		ctx.Next()
	}
}

// RecoverByLogrus -
func RecoverByLogrus(log *logrus.Entry) context.Handler {
	return func(ctx iris.Context) {
		defer func() {
			if err := recover(); err != nil {
				if ctx.IsStopped() {
					return
				}

				var stacktrace string
				for i := 1; ; i++ {
					_, f, l, got := runtime.Caller(i)
					if !got {
						break
					}

					stacktrace += fmt.Sprintf("%s:%d\n", f, l)
				}

				// when stack finishes
				logMessage := fmt.Sprintf("Recovered from a route's Handler('%s')\n", ctx.HandlerName())
				logMessage += fmt.Sprintf("Trace: %s\n", err)
				logMessage += fmt.Sprintf("\n%s", stacktrace)
				log.Errorln(logMessage)

				ctx.StatusCode(500)
				ctx.StopExecution()
				ctx.JSON(response.Error(ctx.Request().Context(), status.UnKnownError.Err()))
			}
		}()

		ctx.Next()
	}
}

// RecoverByZap -
func RecoverByZap(log *zap.Logger) context.Handler {
	return func(ctx iris.Context) {
		defer func() {
			if err := recover(); err != nil {
				if ctx.IsStopped() {
					return
				}

				var stacktrace string
				for i := 1; ; i++ {
					_, f, l, got := runtime.Caller(i)
					if !got {
						break
					}

					stacktrace += fmt.Sprintf("%s:%d\n", f, l)
				}

				// when stack finishes
				logMessage := fmt.Sprintf("Recovered from a route's Handler('%s')\n", ctx.HandlerName())
				logMessage += fmt.Sprintf("Trace: %s\n", err)
				logMessage += fmt.Sprintf("\n%s", stacktrace)
				log.Error(logMessage)

				ctx.StatusCode(500)
				ctx.StopExecution()
				ctx.JSON(response.Error(ctx.Request().Context(), status.UnKnownError.Err()))
			}
		}()

		ctx.Next()
	}
}
