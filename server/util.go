package server

import (
	"fmt"
)

const (
	Port = 10000
)

var (
	DemoAddr string = fmt.Sprintf("localhost:%d", Port)
)
