package qdbus

type NightLightClient struct {
	*QDBusClient
}

func NewNightLightClient() *NightLightClient {
	return &NightLightClient{
		QDBusClient: NewQDBusClient(
			"org.kde.kded",
			"/modules/nightlight",
			"org.kde.kded.nightlight",
		),
	}
}

// Methods

func (c *NightLightClient) Running() (bool, error) {
	output, err := c.output("running")
	if err != nil {
		return false, err
	}
	return output == "true", nil
}
