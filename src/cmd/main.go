// Copyright (c) 2024 Rui Barroso
// This code is licensed under the MIT License.

package main

import (
	mythemes "calculator/src/my_themes"
	"calculator/src/views"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	// "fyne.io/fyne/v2/canvas"
	// "fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/layout"
)

func main() {
	log.Println("Init App...")
	RunApp()
	log.Println("Close App...")
}

func RunApp() {
	myApp := app.New()
	myApp.Settings().SetTheme(&mythemes.AppTheme{})
	w := myApp.NewWindow("Calculator")

	w.SetContent(views.CreateApp())

	iconResource, err := LoadIconAsset()
	if err != nil {
		fyne.LogError("Failed to load icon", err)
	} else {

		w.SetIcon(iconResource)
	}

	w.Resize(fyne.NewSize(350, 550))
	w.SetFixedSize(true)

	w.ShowAndRun()
}

func LoadIconAsset() (*fyne.StaticResource, error) {
	// VS Code Launcher
	iconData, err := os.ReadFile("../../assets/icon.jpg")

	//Terminal
	if err != nil {
		iconData, err = os.ReadFile("./assets/icon.jpg")
	}

	if err != nil {
		return nil, err
	} else {
		iconResource := fyne.NewStaticResource("window_icon", iconData)
		return iconResource, nil
	}
}
