package dishwasksdk

import (
	"errors"
	"fmt"
	"strings"
	"time"
)

const maxIntensity = 4

var programs = []string{"fast", "basic"}

// RunProgram will run a dishwasher program. The program must exist, and the intensity must be between
func RunProgram(program string, intensity int) error {
	if intensity < 1 {
		return errors.New("intensity must be a positive number")
	}

	if intensity > maxIntensity {
		return fmt.Errorf("intensity %d is too great, the maximum is %d", intensity, maxIntensity)
	}

	if program != "fast" && program != "basic" {
		return fmt.Errorf("program %q does not exist, allowed programs are: %s", program, strings.Join(programs, ", "))
	}

	time.Sleep(20 * time.Millisecond)

	return nil
}
