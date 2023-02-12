package core

import "time"

type Options struct {
	APIURL            string
	AppID             string
	AppKey            string
	EventLoopDuration time.Duration
	Test              bool
	API               API
	RobotID           string
}

func NewOptions(opts ...Option) *Options {
	optList := &Options{
		APIURL:            "http://175.27.245.108:8081/v1",
		EventLoopDuration: time.Millisecond * 300,
	}

	if len(opts) > 0 {
		for _, opt := range opts {
			opt(optList)
		}
	}

	return optList
}

type Option func(opts *Options)

func WithAppID(appID string) Option {

	return func(opts *Options) {
		opts.AppID = appID
	}
}
func WithRobotID(robotID string) Option {

	return func(opts *Options) {
		opts.RobotID = robotID
	}
}

func WithAppKey(appKey string) Option {

	return func(opts *Options) {
		opts.AppKey = appKey
	}
}

func WithTest(test bool) Option {
	return func(opts *Options) {
		opts.Test = test
	}
}
