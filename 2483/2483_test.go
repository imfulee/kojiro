package main

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func BestClosingTimeTest(t *testing.T, customers string, expect int) {
	closingHour := bestClosingTime(customers)
	if closingHour != expect {
		t.Errorf("Closing time for case %s should be %d but got %d", customers, expect, closingHour)
	}
}

func TestBestClosingTimeCase1(t *testing.T) {
	customers := "YYNY"
	expect := 2
	BestClosingTimeTest(t, customers, expect)
}

func TestBestClosingTimeCase2(t *testing.T) {
	customers := "NNNNN"
	expect := 0
	BestClosingTimeTest(t, customers, expect)
}

func TestBestClosingTimeCase3(t *testing.T) {
	customers := "YYYY"
	expect := 4
	BestClosingTimeTest(t, customers, expect)
}

func TestBestClosingTimeCase4(t *testing.T) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Errorf("Could not open file for large case")
	}
	dir := filepath.Dir(filename)
	dataFilepath := filepath.Join(dir, "./test/LargeCase.txt")
	file, err := os.Open(dataFilepath)
	if err != nil {
		t.Errorf("Could not open file for large case %s", dataFilepath)
	}

	r := bufio.NewReader(file)
	bytes, err := io.ReadAll(r)
	if err != nil {
		t.Errorf("Could not read file for large case %s: %v", dataFilepath, err)
	}
	customers := string(bytes)

	expect := 15074
	BestClosingTimeTest(t, customers, expect)
}
