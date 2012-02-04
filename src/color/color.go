/*
  GoMUD: color.go
         Color mappings and functions.
*/
package color

const (
	Escape     = "\x1b"
	Reset      = Escape + "[0m"
	Bright     = Escape + "[1m"
	Dim        = Escape + "[2m"
	Underscore = Escape + "[4m"
	Blink      = Escape + "[5m"
	Reverse    = Escape + "[7m"
	Hidden     = Escape + "[8m"

	FgBlack   = Escape + "[30m"
	FgRed     = Escape + "[31m"
	FgGreen   = Escape + "[32m"
	FgYellow  = Escape + "[33m"
	FgBlue    = Escape + "[34m"
	FgMagenta = Escape + "[35m"
	FgCyan    = Escape + "[36m"
	FgWhite   = Escape + "[37m"

	BgBlack   = Escape + "[40m"
	BgRed     = Escape + "[41m"
	BgGreen   = Escape + "[42m"
	BgYellow  = Escape + "[43m"
	BgBlue    = Escape + "[44m"
	BgMagenta = Escape + "[45m"
	BgCyan    = Escape + "[46m"
	BgWhite   = Escape + "[47m"
)
