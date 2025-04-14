package theme

import (
	"akai-ito/config"
	"encoding/hex"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

type AkaiItoTheme struct{}

var _ fyne.Theme = (*AkaiItoTheme)(nil)

func hexToColor(hexStr string) color.Color {
	hexStr = hexStr[1:] // Remove # prefix
	b, _ := hex.DecodeString(hexStr)
	return color.RGBA{R: b[0], G: b[1], B: b[2], A: 0xff}
}

func (m AkaiItoTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return hexToColor(config.Config.Colors.Base00)
	case theme.ColorNameButton:
		return hexToColor(config.Config.Colors.Base0D)
	case theme.ColorNameDisabled:
		return hexToColor(config.Config.Colors.Base03)
	case theme.ColorNamePrimary:
		return hexToColor(config.Config.Colors.Base05)
	}
	return theme.DefaultTheme().Color(name, variant)
}

func (m AkaiItoTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m AkaiItoTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}

func (m AkaiItoTheme) Icon(name fyne.ThemeIconName) fyne.Resource {

	return theme.DefaultTheme().Icon(name)
}
