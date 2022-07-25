package cmd

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func init() {
	initTimeLocation()
}

func initTimeLocation() {
	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(errors.Wrap(err, "load time location"))
	}

	time.Local = loc
}

var rootCmd = &cobra.Command()

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(errors.Wrap(err, "cmd error"))
	}
}

func test() {
	fmt.Println("it's chaos")
}
