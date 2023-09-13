package cmd

import (
	"fmt"
	"os"

	"github.com/gotohr/fritzctl/cmd/printer"
	"github.com/gotohr/fritzctl/fritz"
	"github.com/gotohr/fritzctl/internal/console"
	"github.com/gotohr/fritzctl/logger"
	"github.com/spf13/cobra"
)

var listSwitchesCmd = &cobra.Command{
	Use:   "switches",
	Short: "List the available smart home switches",
	Long:  "List the available smart home devices [switches] and associated data.",
	Example: `fritzctl list switches
fritzctl list switches --output=json`,
	RunE: listSwitches,
}

func init() {
	listSwitchesCmd.Flags().StringP("output", "o", "", "specify output format")
	listCmd.AddCommand(listSwitchesCmd)
}

func listSwitches(cmd *cobra.Command, _ []string) error {
	devs := mustList()
	logger.Success("Device data:")
	data := selectFmt(cmd, devs.Switches(), switchTable)
	printer.Print(data, os.Stdout)
	return nil
}

func switchTable(devs []fritz.Device) interface{} {
	table := console.NewTable(console.Headers(
		"NAME",
		"PRODUCT",
		"PRESENT",
		"STATE",
		"LOCK (BOX/DEV)",
		"MODE",
		"POWER",
		"ENERGY",
		"TEMP",
		"OFFSET",
	))
	appendSwitches(devs, table)
	return table
}

func appendSwitches(devs []fritz.Device, table *console.Table) {
	for _, dev := range devs {
		table.Append(switchColumns(dev))
	}
}

func switchColumns(dev fritz.Device) []string {
	return []string{
		dev.Name,
		fmt.Sprintf("%s %s", dev.Manufacturer, dev.Productname),
		console.IntToCheckmark(dev.Present),
		console.StringToCheckmark(dev.Switch.State),
		console.StringToCheckmark(dev.Switch.Lock) + "/" + console.StringToCheckmark(dev.Switch.DeviceLock),
		dev.Switch.Mode,
		fmtUnit(dev.Powermeter.FmtPowerW, "W"),
		fmtUnit(dev.Powermeter.FmtEnergyWh, "Wh"),
		fmtUnit(dev.Temperature.FmtCelsius, "°C"),
		fmtUnit(dev.Temperature.FmtOffset, "°C"),
	}
}
