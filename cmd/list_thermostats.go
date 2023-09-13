package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/gotohr/fritzctl/cmd/printer"
	"github.com/gotohr/fritzctl/fritz"
	"github.com/gotohr/fritzctl/internal/console"
	"github.com/gotohr/fritzctl/logger"
	"github.com/spf13/cobra"
)

var listThermostatsCmd = &cobra.Command{
	Use:   "thermostats",
	Short: "List the available smart home thermostats",
	Long:  "List the available smart home devices [thermostats] and associated data.",
	Example: `fritzctl list thermostats
fritzctl list thermostats --output=json`,
	RunE: listThermostats,
}

func init() {
	listThermostatsCmd.Flags().StringP("output", "o", "", "specify output format")
	listCmd.AddCommand(listThermostatsCmd)
}

func listThermostats(cmd *cobra.Command, _ []string) error {
	devs := mustList()
	data := selectFmt(cmd, devs.Thermostats(), thermostatsTable)
	logger.Success("Device data:")
	printer.Print(data, os.Stdout)
	return nil
}

func thermostatsTable(devs []fritz.Device) interface{} {
	table := console.NewTable(console.Headers(
		"NAME",
		"PRODUCT",
		"PRESENT",
		"LOCK (BOX/DEV)",
		"MEASURED",
		"OFFSET",
		"WANT",
		"SAVING",
		"COMFORT",
		"NEXT",
		"STATE",
		"BATTERY",
	))
	appendThermostats(devs, table)
	return table
}

func appendThermostats(devs []fritz.Device, table *console.Table) {
	for _, dev := range devs {
		columns := thermostatColumns(dev)
		table.Append(columns)
	}
}

func thermostatColumns(dev fritz.Device) []string {
	var columnValues []string
	columnValues = appendMetadata(columnValues, dev)
	columnValues = appendRuntimeFlags(columnValues, dev)
	columnValues = appendTemperatureValues(columnValues, dev)
	columnValues = appendRuntimeWarnings(columnValues, dev)
	return columnValues
}

func appendMetadata(cols []string, dev fritz.Device) []string {
	return append(cols, dev.Name, fmt.Sprintf("%s %s", dev.Manufacturer, dev.Productname))
}

func appendRuntimeFlags(cols []string, dev fritz.Device) []string {
	return append(cols,
		console.IntToCheckmark(dev.Present),
		console.StringToCheckmark(dev.Thermostat.Lock)+"/"+console.StringToCheckmark(dev.Thermostat.DeviceLock))
}

func appendRuntimeWarnings(cols []string, dev fritz.Device) []string {
	return append(cols, errorCode(dev.Thermostat.ErrorCode), batteryState(dev.Thermostat))
}

func appendTemperatureValues(cols []string, dev fritz.Device) []string {
	return append(cols,
		fmtUnit(dev.Thermostat.FmtMeasuredTemperature, "°C"),
		fmtUnit(dev.Temperature.FmtOffset, "°C"),
		fmtUnit(dev.Thermostat.FmtGoalTemperature, "°C"),
		fmtUnit(dev.Thermostat.FmtSavingTemperature, "°C"),
		fmtUnit(dev.Thermostat.FmtComfortTemperature, "°C"),
		fmtNextChange(dev.Thermostat.NextChange))
}

func fmtNextChange(n fritz.NextChange) string {
	ts := n.FmtTimestamp(time.Now())
	if ts == "" {
		return "?"
	}
	return ts + " -> " + fmtUnit(n.FmtGoalTemperature, "°C")
}

func errorCode(ec string) string {
	checkMark := console.Stoc(ec).Inverse()
	return checkMark.String() + fritz.HkrErrorDescriptions[ec]
}

func batteryState(thermostat fritz.Thermostat) string {
	return fmt.Sprintf("%s%% %s", thermostat.BatteryChargeLevel, console.Stoc(thermostat.BatteryLow).Inverse().String())
}
