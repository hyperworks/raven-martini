raven-go-middleware
===================

[![GoDoc](https://godoc.org/github.com/hyperworks/raven-martini?status.svg)](https://godoc.org/github.com/hyperworks/raven-martini) [![Build Status](https://travis-ci.org/hyperworks/raven-martini.svg)](https://travis-ci.org/hyperworks/raven-martini)

Easy Middleware for GetSentry (Raven) for GO Web framework Martini

```
import ravenrecover "github.com/hyperworks/raven-martini"

ravenDsn := "https://longnumber:lonnumber@app.getsentry.com/shortnumber"
M := martini.Classic()
logger := log.New(os.Stdout, "[martini] ", 0)

M.Use(ravenrecover.RecoverRaven(ravenDsn, logger))
```
