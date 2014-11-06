package raven-martini

import (
		"log"

		"github.com/go-martini/martini"
		"github.com/getsentry/raven-go"
)
var client *raven.Client

//Send in your raven dsn should look something like this
// "https://longnumber:lonnumber@app.getsentry.com/shortnumber"
func RecoverRaven(dsn string, logger *log.Logger) martini.Handler {
	var err error
	client, err = raven.NewClient(dsn, map[string]string{})

	if err != nil {
		logger.Printf("Error loading raven -%s \n", err.Error())
	}

	return func(c martini.Context, log *log.Logger) {
		defer func() {
			if e := recover(); e != nil {
				if err, ok := e.(error); ok {
					log.Printf("Sending this error to get sentry ")
					client.CaptureError(err, map[string]string{})
				} else if  strErr, ok := e.(string); ok {
					log.Printf("Sending this error to get sentry ")
					client.CaptureMessage(strErr, map[string]string{})
				}
				panic(e)
			}
		}()


		c.Next()
		return
	}
}