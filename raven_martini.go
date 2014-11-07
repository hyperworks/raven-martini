package ravenrecover

import (
		"log"
		"fmt"
		"net/http"

		"github.com/go-martini/martini"
		"github.com/getsentry/raven-go"
)
var client *raven.Client

func trace() *raven.Stacktrace {
	return raven.NewStacktrace(0, 2, nil)
}

//Send in your raven dsn should look something like this
// "https://longnumber:lonnumber@app.getsentry.com/shortnumber"
func RecoverRaven(dsn string, logger *log.Logger) martini.Handler {
	var err error
	client, err = raven.NewClient(dsn, map[string]string{})

	if err != nil {
		logger.Printf("Error loading raven -%s \n", err.Error())
	}

	return func(c martini.Context, log *log.Logger, req *http.Request) {
		defer func() {
			if e := recover(); e != nil {
				if err, ok := e.(error); ok {
					log.Printf("Sending this error to get sentry ")
					packet := raven.NewPacket(err.Error(), raven.NewException(err, trace()), raven.NewHttp(req))
					client.Capture(packet, nil)
				} else if  strErr, ok := e.(string); ok {
					log.Printf("Sending this error to get sentry ")
					packet := raven.NewPacket(strErr, raven.NewException(fmt.Errorf(strErr), trace()), raven.NewHttp(req))
					client.Capture(packet, nil)
				}
				panic(e)
			}
		}()


		c.Next()
		return
	}
}