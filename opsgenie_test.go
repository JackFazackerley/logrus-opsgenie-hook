package opsgenie

import (
	"github.com/go-test/deep"
	"github.com/sirupsen/logrus"
	"testing"

	"github.com/opsgenie/opsgenie-go-sdk/alertsv2"
)

func TestCast(t *testing.T) {
	details := make(map[string]string, 2)
	details["test1"] = "val"
	details["test2"] = "val"

	alert := alertsv2.CreateAlertRequest{
		Message:     "some error here",
		Alias:       "some alias here",
		Description: "some description here",
		Teams: []alertsv2.TeamRecipient{
			&alertsv2.Team{
				Name: "first",
			},
			&alertsv2.Team{
				Name: "second",
			},
		},
		VisibleTo: []alertsv2.Recipient{
			&alertsv2.Team{
				Name: "first",
			},
			&alertsv2.Team{
				Name: "second",
			},
			&alertsv2.User{
				Username: "user1",
			},
			&alertsv2.User{
				Username: "user2",
			},
		},
		Actions:  []string{"action1", "action2"},
		Tags:     []string{"tag1", "tag2"},
		Details:  details,
		Entity:   "some entity here",
		Source:   "some source here",
		Priority: alertsv2.P5,
		User:     "some user here",
		Note:     "some note here",
	}

	entry := logrus.WithField("alert", alert)

	castedAlert, ok := entry.Data["alert"].(alertsv2.CreateAlertRequest)
	if !ok {
		t.Error("doesn't exist")
	}

	if diff := deep.Equal(castedAlert, alert); diff != nil {
		t.Error(diff)
	}
}
