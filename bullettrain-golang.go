package bullettrain_go_golang

import (
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

type Segment struct {
	Fg, Bg color.Attribute
}

func (p *Segment) SetFg(fg color.Attribute) {
	p.Fg = fg
}

func (p *Segment) SetBg(bg color.Attribute) {
	p.Bg = bg
}

func (p *Segment) Render(ch chan<- string) {
	const golang_symbol string = "🐹"
	defer close(ch) // Always close the channel!

	col := color.New(p.Fg, p.Bg)

	golangCmd := exec.Command("go", "version")
	versions_info, err := golangCmd.Output()
	if err == nil {
		ch <- col.Sprintf(" %s %s ", golang_symbol,
			strings.Trim(string(versions_info), "\n"))
		return
	}
}
