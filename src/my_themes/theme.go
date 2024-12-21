// Copyright (c) 2024 Rui Barroso
// This code is licensed under the MIT License.

package mythemes

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type AppTheme struct {
}

func (*AppTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {

	switch n {
	case theme.ColorNameBackground:
		if v == theme.VariantDark {
			return color.RGBA{R: 30, G: 30, B: 30, A: 255} // Dark background color
		}
		return color.RGBA{R: 255, G: 255, B: 255, A: 255} // Light background color
	default:
		// Return default colors for unspecified items.
		return theme.DefaultTheme().Color(n, v)
	}
}
func (*AppTheme) Font(s fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(s)
}
func (*AppTheme) Icon(i fyne.ThemeIconName) fyne.Resource {
	switch i {
	case theme.IconNameInfo:
		return theme.DefaultTheme().Icon(i) // Fallback to default if there's an error
	default:
		return theme.DefaultTheme().Icon(i) // Default icons for other cases
	}
}
func (*AppTheme) Size(n fyne.ThemeSizeName) float32 {
	switch n {
	case theme.SizeNamePadding:
		return 15 // Custom padding size
	default:
		return theme.DefaultTheme().Size(n) // Default sizes
	}
}
