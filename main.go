package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/golang/protobuf/ptypes"
)

func main() {
	var timeformat = "2006/01/02"
	stdin := bufio.NewScanner(os.Stdin)
	fmt.Print("(%YY/%mm/%dd)> ")
	for stdin.Scan() {
		text := stdin.Text()
		t, _ := time.Parse(timeformat, text)
		fmt.Println(ptypes.TimestampProto(t))
		fmt.Print("(%YY/%mm/%dd)> ")
	}
}
