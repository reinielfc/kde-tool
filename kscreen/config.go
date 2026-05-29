package kscreen

import "fmt"

// region Output

type Output struct {
	name             string
	enabled, primary bool
	mode             *Mode
	position         *Position
}

func NewOutput(name string) *Output {
	return &Output{name: name}
}

func (o *Output) Name() string        { return o.name }
func (o *Output) Enabled() bool       { return o.enabled }
func (o *Output) Primary() bool       { return o.primary }
func (o *Output) Mode() *Mode         { return o.mode }
func (o *Output) Position() *Position { return o.position }

func (o *Output) SetEnabled()             { o.enabled = true }
func (o *Output) SetDisabled()            { o.enabled = false }
func (o *Output) SetPrimary()             { o.primary = true }
func (o *Output) SetPosition(p *Position) { o.position = p }
func (o *Output) SetMode(m *Mode)         { o.mode = m }

func (o *Output) BuildArgs() []string {
	args := make([]string, 0, 4)
	prefix := fmt.Sprintf("output.%s", o.name)

	appendAsKey := func(key string) {
		args = append(args, fmt.Sprintf("%s.%s", prefix, key))
	}
	appendAsKeyValue := func(key, value string) {
		args = append(args, fmt.Sprintf("%s.%s.%s", prefix, key, value))
	}

	if !o.enabled {
		appendAsKey("disable")
	} else {
		appendAsKey("enable")

		if o.mode != nil {
			appendAsKeyValue("mode", o.mode.String())
		}
		if o.position != nil {
			appendAsKeyValue("position", o.position.String())
		}
		if o.primary {
			appendAsKey("primary")
		}
	}

	return args
}

// region Mode

type Mode struct {
	size        *Size
	refreshRate int
}

func NewMode(size *Size, refreshRate int) *Mode {
	return &Mode{size: size, refreshRate: refreshRate}
}

func (m *Mode) Size() *Size      { return m.size }
func (m *Mode) RefreshRate() int { return m.refreshRate }

func (m *Mode) String() string {
	return fmt.Sprintf("%v@%d", m.size, m.refreshRate)
}

// region Size

type Size struct {
	width, height int
}

func NewSize(width, height int) *Size {
	return &Size{width: width, height: height}
}

func (s *Size) Width() int  { return s.width }
func (s *Size) Height() int { return s.height }

func (s *Size) String() string {
	return fmt.Sprintf("%dx%d", s.width, s.height)
}

// region Position

type Position struct {
	x, y int
}

func NewPosition(x, y int) *Position {
	return &Position{x: x, y: y}
}

func NewTopLeftPosition() *Position {
	return &Position{x: 0, y: 0}
}

func NewRightOfOutputPosition(output *Output) *Position {
	if output.position == nil || output.mode == nil || output.mode.size == nil {
		return NewTopLeftPosition()
	}

	x := output.position.x + output.mode.size.width
	y := output.position.y

	return &Position{x: x, y: y}
}

func (p *Position) X() int { return p.x }
func (p *Position) Y() int { return p.y }

func (p *Position) String() string {
	return fmt.Sprintf("%d,%d", p.x, p.y)
}
