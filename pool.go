package gopool

import (
	"context"

	"log"
	"sync"
)

// StdLogger is used to log error messages.
type Logger interface {
	Debug(v ...interface{})
	Info(v ...interface{})
	Crit(v ...interface{})
	Fatal(v ...interface{})
}

var PoolLogger GoPoolLogger

type GoPoolLogger struct {
}

func (gpl *GoPoolLogger) Debug(v ...interface{}) {
	log.Println(v)
}
func (gpl *GoPoolLogger) Info(v ...interface{}) {
	log.Println(v)
}
func (gpl *GoPoolLogger) Crit(v ...interface{}) {
	log.Println(v)
}
func (gpl *GoPoolLogger) Fatal(v ...interface{}) {
	log.Println(v)
}

type goPool struct {
	wg        sync.WaitGroup
	cancelctx context.Context
	canclFunc context.CancelFunc
}

//this is a function which takes a context and exits when the context is done.
type GoPoolFunc func(ctx context.Context) error

func NewGoPool(ctx context.Context) *goPool {
	cancelctx, canclFunc := context.WithCancel(ctx)
	return &goPool{
		cancelctx: cancelctx,
		canclFunc: canclFunc,
	}
}

func (gp *goPool) Context() context.Context {
	return gp.cancelctx
}

var (
	PanicHandler = func(err interface{}, method string) {

		switch err.(type) {

		case error:
			PoolLogger.Info("RECOVERED FROM PANIC", "error is ", err.(error).Error(), "details", method)
		case string:
			PoolLogger.Info("RECOVERED FROM PANIC", "error is ", err.(string), "details", method)
		default:
			PoolLogger.Info("RECOVERED FROM PANIC")
		}

	}
)

func (gp *goPool) ShutDown() {
	gp.canclFunc()
	gp.wg.Wait()
}

func (gp *goPool) AddJob(method string, fn GoPoolFunc) {
	gp.wg.Add(1)

	go func() {
		defer func() {
			gp.wg.Done()
			handler := PanicHandler
			if handler != nil {
				if err := recover(); err != nil {
					handler(err, method)
				}
			}
		}()
		fn(gp.Context())
	}()
}
