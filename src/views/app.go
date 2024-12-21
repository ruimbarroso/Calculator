// Copyright (c) 2024 Rui Barroso
// This code is licensed under the MIT License.

package views

import (
	"calculator/src/controller"

	"fyne.io/fyne/v2"
	// "fyne.io/fyne/v2/app"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func CreateApp() fyne.CanvasObject {

	display := widget.NewMultiLineEntry()
	display.SetText("")
	display.Disable()

	ctr := controller.New(display)
	app := container.New(
		layout.NewVBoxLayout(),
		container.NewStack(display),
		container.NewHBox(
			widget.NewButtonWithIcon("", theme.ContentUndoIcon(), func() { ctr.GoBack() }),
			widget.NewButtonWithIcon("", theme.ContentRedoIcon(), func() { ctr.GoFront() }),
		),
		container.NewGridWithRows(5,
			container.NewHBox(
				CreateDefaultBtn("<-", func() { ctr.MoveCursorLeft() }),
				CreateDefaultBtn("->", func() { ctr.MoveCursorRigth() }),
				CreateDefaultBtn("D", func() { ctr.Delete() }),
				CreateDefaultBtn("C", func() { ctr.Clear() }),
				CreateDefaultBtn("=", func() { ctr.Calculate() }),
			),
			container.NewHBox(
				CreateDefaultBtn("1", func() { ctr.Insert("1") }),
				CreateDefaultBtn("2", func() { ctr.Insert("2") }),
				CreateDefaultBtn("3", func() { ctr.Insert("3") }),
				CreateDefaultBtn("+", func() { ctr.Insert("+") }),
				CreateDefaultBtn("(", func() { ctr.Insert("(") }),
			),
			container.NewHBox(
				CreateDefaultBtn("4", func() { ctr.Insert("4") }),
				CreateDefaultBtn("5", func() { ctr.Insert("5") }),
				CreateDefaultBtn("6", func() { ctr.Insert("6") }),
				CreateDefaultBtn("-", func() { ctr.Insert("-") }),
				CreateDefaultBtn(")", func() { ctr.Insert(")") }),
			),
			container.NewHBox(
				CreateDefaultBtn("7", func() { ctr.Insert("7") }),
				CreateDefaultBtn("8", func() { ctr.Insert("8") }),
				CreateDefaultBtn("9", func() { ctr.Insert("9") }),
				CreateDefaultBtn("*", func() { ctr.Insert("*") }),
				CreateDefaultBtn("^", func() { ctr.Insert("^") }),
			),
			container.NewHBox(
				CreateDefaultBtn("log", func() { ctr.Insert("l") }),
				CreateDefaultBtn("0", func() { ctr.Insert("0") }),
				CreateDefaultBtn(".", func() { ctr.Insert(".") }),
				CreateDefaultBtn("/", func() { ctr.Insert("/") }),
				CreateDefaultBtn("rt", func() { ctr.Insert("r") }),
			),
		),
	)
	ctr.WriteInDisplay()
	return app
}
func CreateDefaultBtn(label string, onClick func()) fyne.CanvasObject {
	return CreateBtn(label, onClick, 50, 50)
}

func CreateBtn(label string, onClick func(), width float32, height float32) fyne.CanvasObject {
	btn := widget.NewButton(label, onClick)
	return container.New(layout.NewGridWrapLayout(fyne.NewSize(width, height)), btn)
}
