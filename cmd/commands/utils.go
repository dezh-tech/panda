package commands

import (
	"log"
	"os"
)

func ExitOnError(err error) {
	log.Printf("panda error: %s\n", err.Error()) //nolint
	os.Exit(1)
}
