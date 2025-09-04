package theme

type ThemeManagerDeps struct {
	SC interface{ GetScale() float64 }
}

type ThemeManager struct {
	d ThemeManagerDeps

	fonts   *FontSet
	picker  *ThemePicker
	current Theme
}

func NewThemeService(d ThemeManagerDeps) *ThemeManager {
	tm := &ThemeManager{
		picker: NewThemePicker(),
		fonts:  InitFonts(d.SC.GetScale()),
	}

	tm.current = tm.picker.GetCurrentTheme()

	return tm
}

func (tm *ThemeManager) UpdateFonts() {
	tm.fonts = InitFonts(tm.d.SC.GetScale())
}

func (tm *ThemeManager) Current() Theme {
	return tm.current
}

func (tm *ThemeManager) Fonts() *FontSet {
	return tm.fonts
}

func (tm *ThemeManager) NextFont() {
	tm.current = tm.picker.IncrementCurrentTheme()
}
