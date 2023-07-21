package helper

import (
	"fmt"
	"os"
)

func PrintHostname() {
	fmt.Println(os.Hostname())
}
