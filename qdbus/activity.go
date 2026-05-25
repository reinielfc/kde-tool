package qdbus

import (
	"fmt"
)

type ActivityClient struct {
	*QDBusClient
}

func NewActivityClient() *ActivityClient {
	return &ActivityClient{
		QDBusClient: NewQDBusClient(
			"org.kde.ActivityManager",
			"/ActivityManager/Activities",
			"org.kde.ActivityManager.Activities",
		),
	}
}

// Methods

func (c *ActivityClient) SetCurrentActivity(activityID string) error {
	return c.run("SetCurrentActivity", activityID)
}

func (c *ActivityClient) CurrentActivity() (string, error) {
	return c.output("CurrentActivity")
}

func (c *ActivityClient) ListActivities() ([]string, error) {
	return c.outputLines("ListActivities")
}

func (c *ActivityClient) ActivityName(activityID string) (string, error) {
	return c.output("ActivityName", activityID)
}

// Additional Helpers

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

func (c *ActivityClient) NextActivity() error {
	return c.cycleActivities(1)
}

func (c *ActivityClient) PreviousActivity() error {
	return c.cycleActivities(-1)
}

func (c *ActivityClient) cycleActivities(offset int) error {
	activities, err := c.ListActivities()
	if err != nil {
		return err
	}

	current, err := c.CurrentActivity()
	if err != nil {
		return err
	}

	index := -1
	for i, id := range activities {
		if id == current {
			index = i
			break
		}
	}
	if index == -1 {
		return fmt.Errorf("current activity not found in list of activities")
	}

	n := len(activities)
	nextIndex := ((index+offset)%n + n) % n // handles any offset magnitude
	return c.SetCurrentActivity(activities[nextIndex])
}
