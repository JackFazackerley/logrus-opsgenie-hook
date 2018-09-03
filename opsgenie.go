package opsgenie

import (
	"errors"
	"fmt"
	"github.com/opsgenie/opsgenie-go-sdk/alertsv2"
	ogcli "github.com/opsgenie/opsgenie-go-sdk/client"
	"github.com/sirupsen/logrus"
)

// OpsGenieHook used to hold the OpsGenie client, errors will be sent using the client
type opsgenieHook struct {
	client *ogcli.OpsGenieAlertV2Client
}

// NewHook creates and returns a new OpsGenie alert client with the supplied apiKey
func NewHook(apiKey string) (*opsgenieHook, error) {
	cli := new(ogcli.OpsGenieClient)
	cli.SetAPIKey(apiKey)

	client, err := cli.AlertV2()
	if err != nil {
		return nil, err
	}

	return &opsgenieHook{
		client: client,
	}, nil
}

func (hook *opsgenieHook) Fire(entry *logrus.Entry) error {
	alert, ok := entry.Data["alert"].(alertsv2.CreateAlertRequest)
	if !ok {
		return fmt.Errorf("alert must be an instace of CreateAlertRequest")
	}

	var message error
	// Check to see if WithError has been used, if so then prioritise that
	val, ok := entry.Data["error"].(error)
	if ok {
		message = val
	} else {
		message = errors.New(entry.Message)
	}

	alert.Message = message.Error()

	_, err := hook.client.Create(alert)
	if err != nil {
		return err
	}

	return nil
}

func (hook *opsgenieHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.ErrorLevel,
		logrus.FatalLevel,
		logrus.PanicLevel,
	}
}
