package main

import (
	"github.com/godbus/dbus/v5"
	"github.com/reinielfc/kde-tool/qdbus"
	"github.com/spf13/cobra"
)

func main() {
	var nextActivity, previousActivity bool
	var toggleNightLight bool

	cmd := &cobra.Command{
		Use:   "kde",
		Short: "KDE utilities",
		Run: func(cmd *cobra.Command, args []string) {
			var busAction func(*dbus.Conn) error

			switch {
			case nextActivity:
				busAction = func(conn *dbus.Conn) error {
					return qdbus.NewActivityClient(conn).NextActivity()
				}

			case previousActivity:
				busAction = func(conn *dbus.Conn) error {
					return qdbus.NewActivityClient(conn).PreviousActivity()
				}

			case toggleNightLight:
				busAction = func(conn *dbus.Conn) error {
					return qdbus.NewNightLightClient(conn).Toggle()
				}

			default:
				panic("no action specified")
			}

			var err error

			if busAction != nil {
				conn, connErr := dbus.SessionBus()
				if connErr != nil {
					panic(connErr)
				}

				defer conn.Close()

				err = busAction(conn)
			}

			if err != nil {
				panic(err)
			}

		},
	}

	cmd.Flags().BoolVar(&nextActivity, "next-activity", false, "switch to the next activity")
	cmd.Flags().BoolVar(&previousActivity, "prev-activity", false, "switch to the previous activity")
	cmd.Flags().BoolVar(&toggleNightLight, "toggle-nightlight", false, "toggle night light")

	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
