package binchecker

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestBinCheckerWithoutEnvVar(t *testing.T) {
	backupEnv := ""
	currentEnv, ok := os.LookupEnv("PROMPTAPI_TOKEN")
	if ok {
		backupEnv = currentEnv
		os.Unsetenv("PROMPTAPI_TOKEN")
	}

	r := new(Result)
	got := BinChecker("1234", r)
	if got.Error() != "You need to set PROMPTAPI_TOKEN environment variable" {
		t.Errorf("got: %v", got)
	}

	if ok {
		os.Setenv("PROMPTAPI_TOKEN", backupEnv)
	}
}

func ExampleBinChecker() {
	result := new(Result)
	if err := BinChecker("302596", result); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", result)

	fmt.Printf("BankName: %s\n", result.BankName)
	fmt.Printf("Country: %s\n", result.Country)
	fmt.Printf("URL: %s\n", result.URL)
	fmt.Printf("Type: %s\n", result.Type)
	fmt.Printf("Scheme: %s\n", result.Scheme)
	fmt.Printf("Bin: %s\n", result.Bin)
}
