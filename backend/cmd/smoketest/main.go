package main

import (
	"flag"
	"fmt"
	"os"
)

type SmokeTest interface {
	RunTest(baseUrl string) bool
	GetTestName() string
}

func main() {
	baseUrl := flag.String("base-url", "http://localhost:8080", "The url to access the instance which should be smoke tested.")
	flag.Parse()

	fmt.Println("Starting smoke tests...")
	fmt.Printf("Base url: %s\n", *baseUrl)

	smokeTestSteps := []SmokeTest{
		&IndexHtmlTest{},
		&UsersTest{},
		&ArticlesTest{},
	}

	for idx, step := range smokeTestSteps {
		fmt.Printf("Running smoke test (%d/%d): %s\n", idx+1, len(smokeTestSteps), step.GetTestName())
		if step.RunTest(*baseUrl) {
			fmt.Println("Success.")
			continue
		}
		fmt.Println("Smoke tests failed.")
		os.Exit(1)
	}

	fmt.Println("All smoke tests passed.")
	os.Exit(0)
}
