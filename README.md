# OpsGenie Hook for Logrus <img src="http://i.imgur.com/hTeVwmJ.png" width="40" height="40" alt=":walrus:" class="emoji" title=":walrus:" />&nbsp;[![Build Status](https://travis-ci.org/JackFazackerley/logrus-opsgenie-hook.svg?branch=master)](https://travis-ci.org/JackFazackerley/logrus-opsgenie-hook)

This hook is used to send your errors to [OpsGenie](https://www.opsgenie.com/) as an alert. It uses the [opsgenie-go-sdk](https://github.com/opsgenie/opsgenie-go-sdk) to handle the requests. The levels that are blocked by this hook are `log.Error`, `log.Fatal`, and `log.Panic`.

## Usage

The only configuration needed for this hook is the OpsGenie API key you wish to use. However the alert struct must be created yourself with the fields you wish to use and added as a `logrus.Entry` using `WithField("request", alert)`.

The message of the alert will default to the log level message: `Error("some error")`, although if `WithError(err)` is used that will become a priority.

```go
import (
    "fmt"

    "github.com/jackfazackerley/logrus-opsgenie-hook"
    "github.com/opsgenie/opsgenie-go-sdk/alertsv2"
    "github.com/sirupsen/logrus"
)

func main() {
    log := logrus.New()

    hook, err := opsgenie.NewHook("some API key")
    if err != nil {
        panic(err)
    }

    log.AddHook(hook)

    alert := alertsv2.CreateAlertRequest{
        Alias: "some alias here",
        Description: "some description here",
        Teams: []alertsv2.TeamRecipient{
            &alertsv2.Team{
                Name: "Dev",
            },
        },
        Source:   "some source here",
        Priority: alertsv2.P5,
    }

    log.WithField("alert", alert).Error("default message value")

    log.WithField("alert", alert).WithError(fmt.Errorf("made priority")).Error("default message value")
}
```

The structure for the OpsGenie alert is documented [Here](https://godoc.org/github.com/opsgenie/opsgenie-go-sdk/alertsv2#CreateAlertRequest).