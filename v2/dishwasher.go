package dishwasksdk

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

const maxIntensity = 4

var programs = []string{"fast", "basic"}

var (
	ErrNegativeIntensity = errors.New("intensity must be a positive number")
	ErrIntensityTooGreat = errors.New("intensity is too great")
	ErrInvalidProgram    = errors.New("invalid program")
)

// RunProgram will run a dishwasher program. The program must exist, and the intensity must be between
func RunProgram(program string, intensity int) error {
	if intensity < 1 {
		return ErrNegativeIntensity
	}

	if intensity > maxIntensity {
		return fmt.Errorf("intensity %d is too great, the maximum is %d: %w", intensity, maxIntensity, ErrIntensityTooGreat)
	}

	if program != "fast" && program != "basic" {
		return fmt.Errorf("program %q does not exist, allowed programs are %s: %w", program, strings.Join(programs, ", "), ErrInvalidProgram)
	}

	time.Sleep(20 * time.Millisecond)

	return nil
}
