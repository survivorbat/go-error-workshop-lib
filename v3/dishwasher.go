package dishwasksdk

import (
	"fmt"
	"strings"
	"time"
)

const maxIntensity = 4

var programs = []string{"fast", "basic"}

// RunProgram will run a dishwasher program. The program must exist, and the intensity must be between
func RunProgram(program string, intensity int) error {
	if intensity < 1 || intensity > maxIntensity {
		return &IntensityError{Minimum: 1, Maximum: maxIntensity, Value: intensity}
	}

	if program != "fast" && program != "basic" {
		return &ProgramError{Allowed: programs, Value: program}
	}

	time.Sleep(20 * time.Millisecond)

	return nil
}

type IntensityError struct {
	Minimum int
	Maximum int
	Value   int
}

func (i *IntensityError) Error() string {
	return fmt.Sprintf("intensity must be between %d and %d, provided value: %d", i.Minimum, i.Maximum, i.Value)
}

type ProgramError struct {
	Allowed []string
	Value   string
}

func (p *ProgramError) Error() string {
	return fmt.Sprintf("program %s does not exist, allowed: %s", p.Value, strings.Join(p.Allowed, ", "))
}
