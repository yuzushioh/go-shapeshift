package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

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

	case "getmarketinfo":
		cli := shapeshift.New()
		if os.Args[2] == "" {
			return errors.New("Specify a currency pair sush as btc_ltc")
		}
		res, err := cli.GetMarketInfo(context.Background(), os.Args[2])
		if err != nil {
			return err
		}
		if err := json.NewEncoder(os.Stdout).Encode(res); err != nil {
			return nil
		}

	case "getrecenttransactions":
		cli := shapeshift.New()
		if os.Args[2] == "" {
			return errors.New("Specify the transaction limit number")
		}
		limit, err := strconv.Atoi(os.Args[2])
		if err != nil {
			return err
		}
		res, err := cli.GetTransactionList(context.Background(), limit)
		if err != nil {
			return err
		}
		if err := json.NewEncoder(os.Stdout).Encode(res); err != nil {
			return nil
		}

	default:
		return errors.Errorf("does not support %q", os.Args[1])
	}

	return nil
}
