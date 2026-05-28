package qdbus

import (
	"github.com/godbus/dbus/v5"
)

const (
	propertiesInterface = "org.freedesktop.DBus.Properties"
)

type QDBusInterfaceClient struct {
	conn      *dbus.Conn
	svcName   string
	objPath   string
	ifaceName string
}

func NewQDBusInterfaceClient(conn *dbus.Conn, svcName, objPath, ifaceName string) *QDBusInterfaceClient {
	return &QDBusInterfaceClient{
		conn:      conn,
		svcName:   svcName,
		objPath:   objPath,
		ifaceName: ifaceName,
	}
}

func (c *QDBusInterfaceClient) callInterface(iface, method string, flags dbus.Flags, args ...any) (*dbus.Call, error) {
	call := c.object().Call(iface+"."+method, flags, args...)
	if call.Err != nil {
		return nil, call.Err
	}
	return call, nil
}

func (c *QDBusInterfaceClient) callGetter(flags dbus.Flags, property string) (*dbus.Call, error) {
	return c.callInterface(propertiesInterface, "Get", flags, c.ifaceName, property)
}

func (c *QDBusInterfaceClient) callMethod(method string, flags dbus.Flags, args ...any) (*dbus.Call, error) {
	return c.callInterface(c.ifaceName, method, flags, args...)
}

func (c *QDBusInterfaceClient) object() dbus.BusObject {
	return c.conn.Object(c.svcName, dbus.ObjectPath(c.objPath))
}

// typed calls and properties

func (c *QDBusInterfaceClient) boolProperty(property string, flags dbus.Flags) (bool, error) {
	return typedProperty[bool](c, property, flags)
}

func (c *QDBusInterfaceClient) voidCall(method string, flags dbus.Flags, args ...any) error {
	_, err := c.callMethod(method, flags, args...)
	return err
}

func (c *QDBusInterfaceClient) uintCall(method string, flags dbus.Flags, args ...any) (uint32, error) {
	return typedCall[uint32](c, method, flags, args...)
}

func (c *QDBusInterfaceClient) stringCall(method string, flags dbus.Flags, args ...any) (string, error) {
	return typedCall[string](c, method, flags, args...)
}

func (c *QDBusInterfaceClient) stringListCall(method string, flags dbus.Flags, args ...any) ([]string, error) {
	return typedCall[[]string](c, method, flags, args...)
}

func typedProperty[T any](c *QDBusInterfaceClient, property string, flags dbus.Flags) (T, error) {
	call, err := c.callGetter(flags, property)
	if err != nil {
		var zero T
		return zero, err
	}
	var value T
	err = call.Store(&value)
	if err != nil {
		var zero T
		return zero, err
	}
	return value, nil
}

func typedCall[T any](c *QDBusInterfaceClient, method string, flags dbus.Flags, args ...any) (T, error) {
	call, err := c.callMethod(method, flags, args...)
	if err != nil {
		var zero T
		return zero, err
	}
	var result T
	err = call.Store(&result)
	if err != nil {
		var zero T
		return zero, err
	}
	return result, nil
}
