package qdbus

type ShortcutClient struct {
	*QDBusClient
}

func NewShortcutClient() *ShortcutClient {
	return &ShortcutClient{
		QDBusClient: NewQDBusClient(
			"org.kde.kglobalaccel",
			"/kglobalaccel",
			"org.kde.kglobalaccel.Component",
		),
	}
}

// Methods

func (c *ShortcutClient) InvokeShortcut(name string) error {
	return c.run("invokeShortcut", name)
}

// Shortcuts

func (c *ShortcutClient) ToggleNightColor() error {
	return c.InvokeShortcut("Toggle Night Color")
}
