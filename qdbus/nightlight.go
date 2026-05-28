package qdbus

import "github.com/godbus/dbus/v5"

// region Client

const (
	nightLightSvcName   = "org.kde.KWin.NightLight"
	nightLightObjPath   = "/org/kde/KWin/NightLight"
	nightLightIfaceName = "org.kde.KWin.NightLight"
)

const (
	shortcutSvcName   = "org.kde.kglobalaccel"
	shortcutObjPath   = "/component/kwin"
	shortcutIfaceName = "org.kde.kglobalaccel.Component"
)

type NightLightClient struct {
	nightLight *QDBusInterfaceClient
	shortcut   *QDBusInterfaceClient
}

func NewNightLightClient(conn *dbus.Conn) *NightLightClient {
	return &NightLightClient{
		nightLight: &QDBusInterfaceClient{
			conn:      conn,
			svcName:   nightLightSvcName,
			objPath:   nightLightObjPath,
			ifaceName: nightLightIfaceName,
		},
		shortcut: &QDBusInterfaceClient{
			conn:      conn,
			svcName:   shortcutSvcName,
			objPath:   shortcutObjPath,
			ifaceName: shortcutIfaceName,
		},
	}
}

func (c *NightLightClient) Toggle() error {
	return c.shortcut.voidCall("invokeShortcut", 0, "Toggle Night Color")
}

func (c *NightLightClient) Running() (bool, error) {
	return c.nightLight.boolProperty("running", 0)
}

func (c *NightLightClient) TurnOn() error {
	return c.setState(true)
}

func (c *NightLightClient) TurnOff() error {
	return c.setState(false)
}

func (c *NightLightClient) setState(state bool) error {
	running, err := c.Running()
	if err != nil {
		return err
	}

	if running == state {
		return nil
	}

	return c.Toggle()
}
