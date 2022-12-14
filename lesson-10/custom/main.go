package main

import (
	"fmt"
)

type badError struct {
	message           string
	additionalContext string
}

func (e *badError) Error() string {
	return e.message
}

type worseError struct {
	message           string
	warpDriveUnstable bool
}

func (e *worseError) Error() string {
	return e.message
}

func main() {
	success, err := doRiskyManeuver(riskLevelMedium)

	if err != nil {
		switch err {
		case err.(*badError):
			fmt.Printf("It's bad Captain: %s %s\n", err, err.(*badError).additionalContext)
		case err.(*worseError):
			fmt.Printf("It's realy bad Captain: %s\n", err)
			if err.(*worseError).warpDriveUnstable {
				fmt.Println("The warp drive is unstable Captain! What do we do?")
			}
		default:
			fmt.Println("I don't know this error.")
		}
	}

	if success {
		fmt.Println("We made it to port, Captain!")
	}
}

type riskLevel int

const (
	riskLevelLow riskLevel = iota
	riskLevelMedium
	riskLevelHigh
)

func doRiskyManeuver(rl riskLevel) (bool, error) {
	switch rl {
	case riskLevelMedium:
		return false, &badError{
			message:           "We're running on empty.",
			additionalContext: "Not enough dillitium crystals",
		}
	case riskLevelHigh:
		return false, &worseError{
			message:           "Plasma chamber's melting",
			warpDriveUnstable: true,
		}
	default:
		return true, nil
	}

}
