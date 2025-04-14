package main

import (
	"fmt"
	"log"
	"slices"

	"akai-ito/config"
	"akai-ito/theme"
	"akai-ito/utils"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Wifx/gonetworkmanager"
)

func main() {
	config.Init()
	// Create a new Fyne app
	a := app.NewWithID("akai-ito")
	a.Settings().SetTheme(theme.AkaiItoTheme{})
	w := a.NewWindow("akai-ito")
	w.Resize(fyne.NewSize(float32(config.Config.Size.Width), float32(config.Config.Size.Height)))

	// Create new instance of gonetworkmanager
	nm, err := gonetworkmanager.NewNetworkManager()
	utils.CheckError(err)

	// Get active connections
	activeConnections, err := nm.GetPropertyActiveConnections()
	utils.CheckError(err)

	// List of active connection IDs
	activeConnectionIDs := []string{}
	activeConnectionMap := make(map[string]gonetworkmanager.ActiveConnection)
	for _, connection := range activeConnections {
		id, err := connection.GetPropertyID()
		if err != nil {
			log.Fatal(err)
		}
		activeConnectionIDs = append(activeConnectionIDs, id)
		activeConnectionMap[id] = connection
	}

	// Create a list to show active connections first
	var activeItems []string
	var inactiveItems []string

	// Fetch all connections
	settings, err := gonetworkmanager.NewSettings()
	utils.CheckError(err)

	connections, err := settings.ListConnections()
	utils.CheckError(err)

	connectionMap := make(map[string]gonetworkmanager.Connection)

	// Sort connections into active and inactive
	for _, connection := range connections {
		settings, err := connection.GetSettings()
		utils.CheckError(err)
		connID := settings["connection"]["id"].(string)

		connectionMap[connID] = connection
		if slices.Contains(activeConnectionIDs, connID) {
			activeItems = append(activeItems, fmt.Sprintf("%s", connID))
		} else {
			inactiveItems = append(inactiveItems, fmt.Sprintf("%s", connID))
		}
	}

	activeList := widget.NewList(
		func() int { return len(activeItems) },
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(i int, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(activeItems[i] + " (active)")
		},
	)

	// Set a click handler for active connections
	activeList.OnSelected = func(id widget.ListItemID) {
		selectedConnection := activeItems[id]
		activeList.Unselect(id)
		fmt.Println("Selected active connection:", selectedConnection)
		nm.DeactivateConnection(activeConnectionMap[selectedConnection])
		w.Close()
	}
	inactiveList := widget.NewList(
		func() int { return len(inactiveItems) },
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(i int, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(inactiveItems[i] + " (inactive)")
		},
	)

	// Set a click handler for inactive connections
	inactiveList.OnSelected = func(id widget.ListItemID) {
		selectedConnection := inactiveItems[id]
		inactiveList.Unselect(id)
		fmt.Println("Selected inactive connection:", selectedConnection)
		devices, err := nm.GetDevices()
		utils.CheckError(err)
		nm.ActivateConnection(connectionMap[selectedConnection], devices[config.Config.DeviceIndex], nil)
		w.Close()
	}

	// Create the separator
	// separator := widget.NewSeparator()

	// Combine all elements: Active, Separator, and Inactive lists
	content := container.NewVSplit(
		activeList,
		inactiveList,
	)

	// Set the content of the window to the scrollable container
	w.SetContent(content)

	// Show the window
	w.ShowAndRun()
}
