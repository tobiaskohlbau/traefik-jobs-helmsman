package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// extracted from https://github.com/golang/go/blob/139cd0e12ff9d7628c321abbfb8d2f4ada461543/src/cmd/go/internal/modload/build.go#L29-L30
var (
	infoStart, _ = hex.DecodeString("3077af0c9274080241e1c107e6d618e6")
	infoEnd, _   = hex.DecodeString("f932433186182072008242104116d8f2")
)

func execute() error {
	file, err := os.Open(os.Args[1])
	if err != nil {
		return fmt.Errorf("failed to open input: %w", err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("failed to read input: %w", err)
	}

	info := string(bytes.Split(bytes.SplitAfter(data, infoStart)[1], infoEnd)[0])

	fmt.Printf("%s\n", info)

	return nil
}

func main() {
	if err := execute(); err != nil {
		log.Fatal(err)
	}
}
