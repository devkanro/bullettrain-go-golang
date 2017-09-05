package carGo

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"

	"github.com/bullettrain-sh/bullettrain-go-core/ansi"
)

const (
	carPaint      = "black:123"
	goSymbolPaint = "black:123"
	goSymbolIcon  = " "
)

// Car for Go
type Car struct {
	paint string
	// Current directory
	Pwd string
}

func paintedSymbol() string {
	var symbolIcon string
	if symbolIcon = os.Getenv("BULLETTRAIN_CAR_GO_ICON"); symbolIcon == "" {
		symbolIcon = goSymbolIcon
	}

	var symbolPaint string
	if symbolPaint = os.Getenv("BULLETTRAIN_CAR_GO_ICON_PAINT"); symbolPaint == "" {
		symbolPaint = goSymbolPaint
	}

	return ansi.Color(symbolIcon, symbolPaint)
}

// GetPaint returns the calculated end paint string for the car.
func (c *Car) GetPaint() string {
	if c.paint = os.Getenv("BULLETTRAIN_CAR_GO_PAINT"); c.paint == "" {
		c.paint = carPaint
	}

	return c.paint
}

// CanShow decides if this car needs to be displayed.
func (c *Car) CanShow() bool {
	if e := os.Getenv("BULLETTRAIN_CAR_GO_SHOW"); e == "true" {
		return true
	}

	var d string
	if d = c.Pwd; d == "" {
		return false
	}

	// Show when .go files exist in current directory
	p := fmt.Sprintf("%s%s*.go", d, string(os.PathSeparator))
	f, _ := filepath.Glob(p)
	if f != nil {
		return true
	}

	return false
}

// Render builds and passes the end product of a completely composed car onto
// the channel.
func (c *Car) Render(out chan<- string) {
	defer close(out) // Always close the channel!
	carPaint := ansi.ColorFunc(c.GetPaint())

	cmd := exec.Command("go", "version")
	cmdOut, err := cmd.Output()
	if err == nil {
		re := regexp.MustCompile(`([0-9.]+)`)
		versionArr := re.FindStringSubmatch(string(cmdOut))
		var version string
		if len(versionArr) > 0 {
			version = versionArr[1]
		}

		out <- fmt.Sprintf("%s%s", paintedSymbol(), carPaint(version))
	}
}

// GetSeparatorPaint overrides the Fg/Bg colours of the right hand side
// separator through ENV variables.
func (c *Car) GetSeparatorPaint() string {
	return os.Getenv("BULLETTRAIN_CAR_GO_SEPARATOR_PAINT")
}

// GetSeparatorSymbol overrides the symbol of the right hand side
// separator through ENV variables.
func (c *Car) GetSeparatorSymbol() string {
	return os.Getenv("BULLETTRAIN_CAR_GO_SEPARATOR_SYMBOL")
}
