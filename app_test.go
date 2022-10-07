package main

import (
	"changeme/utils"
	"fmt"
	"testing"
)

func TestGetDirList(t *testing.T) {
	s := utils.GetPathData()
	fmt.Printf("s: %v\n", s)
}
