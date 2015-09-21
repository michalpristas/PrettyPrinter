package prettyPrinter

import "fmt"

type PrettyString string

func (s PrettyString) colorOutput(color string, text string) PrettyString {
	return PrettyString(fmt.Sprintf("%s%s%s%s", s, color, text, none))
}

func (s PrettyString) changeBaground(color string) PrettyString {
	return PrettyString(fmt.Sprintf("%s%s", s, color))
}

func (s PrettyString) Format(args ...interface{}) PrettyString {
	return PrettyString(fmt.Sprintf(string(s), args...))
}

func (s PrettyString) String() string {
	return string(s)
}

func New() PrettyString {
	return PrettyString("")
}

func (s PrettyString) Black(text string) PrettyString     { return s.colorOutput(black, text) }
func (s PrettyString) Red(text string) PrettyString       { return s.colorOutput(red, text) }
func (s PrettyString) Green(text string) PrettyString     { return s.colorOutput(green, text) }
func (s PrettyString) Brown(text string) PrettyString     { return s.colorOutput(brown, text) }
func (s PrettyString) Blue(text string) PrettyString      { return s.colorOutput(blue, text) }
func (s PrettyString) Magenta(text string) PrettyString   { return s.colorOutput(magenta, text) }
func (s PrettyString) Cyan(text string) PrettyString      { return s.colorOutput(cyan, text) }
func (s PrettyString) Yellow(text string) PrettyString    { return s.colorOutput(yellow, text) }
func (s PrettyString) Purple(text string) PrettyString    { return s.colorOutput(purple, text) }
func (s PrettyString) White(text string) PrettyString     { return s.colorOutput(white, text) }
func (s PrettyString) LightGray(text string) PrettyString { return s.colorOutput(lightGray, text) }
func (s PrettyString) DarkGray(text string) PrettyString  { return s.colorOutput(darkGray, text) }
func (s PrettyString) BoldRed(text string) PrettyString   { return s.colorOutput(brightRed, text) }
func (s PrettyString) BoldGreen(text string) PrettyString { return s.colorOutput(brightGreen, text) }
func (s PrettyString) BoldBlue(text string) PrettyString  { return s.colorOutput(brightBlue, text) }
func (s PrettyString) BoldCyan(text string) PrettyString  { return s.colorOutput(brightCyan, text) }
func (s PrettyString) Text(text string) PrettyString      { return s.colorOutput(none, text) }

func (s PrettyString) BgRed() PrettyString     { return s.changeBaground(bgRed) }
func (s PrettyString) BgGreen() PrettyString   { return s.changeBaground(bgGreen) }
func (s PrettyString) BgYellow() PrettyString  { return s.changeBaground(bgYellow) }
func (s PrettyString) BgBlue() PrettyString    { return s.changeBaground(bgBlue) }
func (s PrettyString) BgMagenta() PrettyString { return s.changeBaground(bgMagenta) }
func (s PrettyString) BgCyan() PrettyString    { return s.changeBaground(bgCyan) }
func (s PrettyString) BgWhite() PrettyString   { return s.changeBaground(bgWhite) }

func (s PrettyString) Reset() PrettyString { return s.changeBaground(none) }
