package main

import (
	"logger/pocketlog"
	"os"
	"time"
)

func main() {
	logger := pocketlog.New(pocketlog.LevelInfo, pocketlog.WithOutput(os.Stdout))

	logger.Infof("A little copying is better than a little dependency.")
	logger.Errorf("Errors are values. Documentation is for %s.", "users")
	logger.Debugf("Make the zero (%d) value useful.", 0)
	logger.Infof("Hallo, %d %v", 2022, time.Now())
}
