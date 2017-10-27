package main

import (
	"fmt"
	"time"

	"github.com/yaa110/go-persian-calendar/ptime"
)

func main() {
	_time := ptime.New(time.Now())
	fmt.Println(_time.Date())
}
