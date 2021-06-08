# Introduction

This package contains a ready to use hook for the[logrus](https://github.com/sirupsen/logrus) logging package that will collect and send logs to datadog via the http intake. It will batch logs up on a timer or to the maximum amount per request before sending to datadog, Which ever is sooner.

It also modifies the log format to ensure that datadog can properly read the message attribute.

It also allows for basic global tags and attributes to be added to all logs. Some typical ones can be added via environment variables.

Some other behavior can also be controlled via environment variables. All of the available variables are listed below.

| Environment Variable | Description | Example | Required | Default |
|---|---|---|---|---|
| `SERVICE` | Will be used as the `service` tag, and attribute on all logs | `my-service` | false | `unknown` |
| `ENVIRONMENT` | Will be used as the `environment` tag and attribute on all logs | `dev` | false | `unknown` |
| `MAINTAINER` | Will be used as the `maintainer` tag and attribute on all logs | `gfs` | false | `unknown` |
| `APPLICATION` | Will be used as the `application` tag and attribute on all logs | `my-application` | false | `unknown` |
| `HOST` | Will be used as the `hostname` tag and attribute on all logs | `my-host` | false | `0.0.0.0` |
| `DATADOG_REGION` | Is used to determine the region to send the logs to, accepts `eu` or `us` | `eu` | false | `us` |
| `DATADOG_API_KEY` | Is used to set the api key to use to authenticate the log post request, this is only used if the apiKey provided to the `New` method of the `datadog` module is `nil`. Without an api key this integration will not work. | `your-api-key` | true if apiKey is `nil`, otherwise false | No default |
| `DATADOG_MAX_RETRIES` | This is used to determine how many times a post request should be retried after failure | `3` | false | `5` |

# Installing the module

```
> go get github.com/GlobalFreightSolutions/logrus-datadog-hook
```

# Using the Module

```go
package main

import (
  "time"

  "github.com/GlobalFreightSolutions/logrus-datadog-hook/datadog"
  "github.com/sirupsen/logrus"
)

func main() {
  apiKey := "YOUR_API_KEY_HERE"
  hook, err := datadog.New(&apiKey, logrus.InfoLevel)
  if err !- nil {
    panic(err.Error())
  }

  logger := logrus.New()
  logger.AddHook(hook)
  
  // This ensures that the logger exits gracefully and all buffered logs are sent before closing down
	logrus.DeferExitHandler(hook.Close)

  for {
    logger.Info("This is a log sent to datadog")
    time.Sleep(30 * time.Second)
  }
}
```
