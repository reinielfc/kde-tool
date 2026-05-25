package main

import (
	"github.com/reinielfc/kde-tools/qdbus"
	"github.com/spf13/cobra"
)

func main() {
	var nextActivity, previousActivity bool
	var toggleNightLight bool

	cmd := &cobra.Command{
		Use:   "kde",
		Short: "KDE utilities",
		Run: func(cmd *cobra.Command, args []string) {
			var err error
			switch {
			case nextActivity:
				err = qdbus.NewActivityClient().NextActivity()

			case previousActivity:
				err = qdbus.NewActivityClient().PreviousActivity()

			case toggleNightLight:
				err = qdbus.NewShortcutClient().ToggleNightColor()

			default:
				panic("no action specified")
			}

			if err != nil {
				panic(err)
			}
		},
	}

	cmd.Flags().BoolVarP(&nextActivity, "next-activity", "a", false, "switch to the next activity")
	cmd.Flags().BoolVarP(&previousActivity, "prev-activity", "A", false, "switch to the previous activity")
	cmd.Flags().BoolVarP(&toggleNightLight, "nightlight", "l", false, "toggle night light")

	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
