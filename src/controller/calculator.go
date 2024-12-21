// Copyright (c) 2024 Rui Barroso
// This code is licensed under the MIT License.

package controller

import (
	"calculator/src/model"
	"fmt"
	"strings"

	"fyne.io/fyne/v2/widget"
)

const (
	Cursor   = "|"
	ErrorMSG = "F!!!"
)

type CalculatorController struct {
	equation     model.Equation
	Display      *widget.Entry
	History      []model.Equation
	historyIndex int
	cursorIndex  int
}

func (t *CalculatorController) Calculate() {

	res, err := model.Evaluate(t.equation.ParseEquation(Cursor))

	t.InsertInHistory()
	t.Clear()

	if err != nil {
		t.Display.SetText(ErrorMSG)
	} else {
		t.Display.SetText(fmt.Sprintf("%g", res))
	}

}
func New(display *widget.Entry) *CalculatorController {
	return &CalculatorController{
		Display:      display,
		equation:     model.Equation{Equation: Cursor},
		cursorIndex:  0,
		History:      make([]model.Equation, 0),
		historyIndex: -1,
	}
}
func (t *CalculatorController) WriteInDisplay() {
	t.Display.SetText(t.equation.Equation)
}
func (t *CalculatorController) Insert(c string) {
	res := strings.Split(t.equation.Equation, Cursor)

	t.equation.Equation = res[0] + c + Cursor + res[1]
	t.WriteInDisplay()
}
func (t *CalculatorController) Delete() {
	res := strings.Split(t.equation.Equation, Cursor)

	if res[0] == "" {
		t.equation.Equation = res[0] + Cursor + res[1]
	} else {
		t.equation.Equation = res[0][:len(res[0])-1] + Cursor + res[1]
	}

	t.WriteInDisplay()
}
func (t *CalculatorController) Clear() {
	t.equation.Equation = Cursor
	t.WriteInDisplay()
}

func (t *CalculatorController) MoveCursorLeft() {
	res := strings.Split(t.equation.Equation, Cursor)

	if res[0] != "" {
		t.equation.Equation = res[0][:len(res[0])-1] + Cursor + res[0][len(res[0])-1:] + res[1]
	}

	t.WriteInDisplay()
}
func (t *CalculatorController) MoveCursorRigth() {
	res := strings.Split(t.equation.Equation, Cursor)

	if res[1] != "" {
		t.equation.Equation = res[0] + res[1][:1] + Cursor + res[1][1:]
	}
	t.WriteInDisplay()
}

func (t *CalculatorController) InsertInHistory() {
	t.historyIndex++
	if t.historyIndex < len(t.History) {
		t.History[t.historyIndex] = t.equation
		t.History = t.History[:t.historyIndex+1]
	} else {
		t.History = append(t.History, t.equation)
	}

}
func (t *CalculatorController) GoBack() {

	if t.historyIndex == -1 {
		t.Clear()
		return
	} else if t.historyIndex == 0 {
		t.Clear()
		t.historyIndex--
		return
	}

	t.historyIndex--
	t.equation = t.History[t.historyIndex]
	t.WriteInDisplay()

}
func (t *CalculatorController) GoFront() {
	if t.historyIndex == len(t.History)-1 {
		return
	}
	t.historyIndex++

	t.equation = t.History[t.historyIndex]
	t.WriteInDisplay()
}
