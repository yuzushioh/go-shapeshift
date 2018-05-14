package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/pkg/errors"
	shapeshift "github.com/yuzushioh/go-shapeshift"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}
}

func run() error {
	if len(os.Args) == 1 {
		return errors.New("Specify a subcommand")
	}

	switch os.Args[1] {
	case "getrate":
		cli := shapeshift.New()
		if os.Args[2] == "" {
			return errors.New("Specify a currency pair sush as btc_ltc")
		}
		res, err := cli.GetRate(context.Background(), os.Args[2])
		if err != nil {
			return err
		}
		if err := json.NewEncoder(os.Stdout).Encode(res); err != nil {
			return err
		}

	case "getdepositlimit":
		cli := shapeshift.New()
		if os.Args[2] == "" {
			return errors.New("Specify a currency pair sush as btc_ltc")
		}
		res, err := cli.GetDepositLimit(context.Background(), os.Args[2])
		if err != nil {
			return err
		}
		if err := json.NewEncoder(os.Stdout).Encode(res); err != nil {
			return err
		}

	default:
		return errors.Errorf("does not support %q", os.Args[1])
	}

	return nil
}
