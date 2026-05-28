package qdbus

import (
	"fmt"

	"github.com/godbus/dbus/v5"
)

// region Client

const (
	activityMgrSvcName  = "org.kde.ActivityManager"
	activitiesObjPath   = "/ActivityManager/Activities"
	activitiesIfaceName = "org.kde.ActivityManager.Activities"
)

type ActivityClient struct {
	activities *QDBusInterfaceClient
}

func NewActivityClient(conn *dbus.Conn) *ActivityClient {
	return &ActivityClient{
		activities: &QDBusInterfaceClient{
			conn:      conn,
			svcName:   activityMgrSvcName,
			objPath:   activitiesObjPath,
			ifaceName: activitiesIfaceName,
		},
	}
}

func (c *ActivityClient) SetCurrentActivity(activityID string) error {
	return c.activities.voidCall("SetCurrentActivity", 0, activityID)
}

func (c *ActivityClient) CurrentActivity() (string, error) {
	return c.activities.stringCall("CurrentActivity", 0)
}

func (c *ActivityClient) ListActivities() ([]string, error) {
	return c.activities.stringListCall("ListActivities", 0)
}

func (c *ActivityClient) ActivityName(activityID string) (string, error) {
	return c.activities.stringCall("ActivityName", 0, activityID)
}

func (c *ActivityClient) NextActivity() error {
	return c.activities.voidCall("NextActivity", 0)
}

func (c *ActivityClient) PreviousActivity() error {
	return c.activities.voidCall("PreviousActivity", 0)
}

func (c *ActivityClient) SetCurrentActivityByName(activityName string) error {
	activityID, err := c.activityID(activityName)
	if err != nil {
		return err
	}
	return c.SetCurrentActivity(activityID)
}

func (c *ActivityClient) activityID(activityName string) (string, error) {
	activities, err := c.ListActivities()
	if err != nil {
		return "", err
	}

	for _, id := range activities {
		name, err := c.ActivityName(id)
		if err != nil {
			return "", err
		}
		if name == activityName {
			return id, nil
		}
	}
	return "", fmt.Errorf("activity not found: %s", activityName)
}

func (c *ActivityClient) CurrentActivityName() (string, error) {
	id, err := c.CurrentActivity()
	if err != nil {
		return "", err
	}
	return c.ActivityName(id)
}
