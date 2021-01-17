package main

import (
	"github.com/go-kit/kit/log"
	"time"
)

type loggingMiddleware struct {
	logger log.Logger
	next   StringService
}

func loggingMiddlewareCre(logger log.Logger) ServiceMiddleware {
	return func(next StringService) StringService {
		return loggingMiddleware{
			logger: logger,
			next:   next,
		}
	}
}
func (mw loggingMiddleware) Uppercase(s string) (output string, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "uppercase",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.next.Uppercase(s)
	return
}

func (mw loggingMiddleware) Count(s string) (n int) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "count",
			"input", s,
			"n", n,
			"took", time.Since(begin),
		)
	}(time.Now())
	n = mw.next.Count(s)
	return
}

func (mw loggingMiddleware) Status() (s string) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "Status",
			"took", time.Since(begin),
		)
	}(time.Now())
	s = mw.next.Status()
	return
}
