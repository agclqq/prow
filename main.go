package main

import (
	"os"

	"github.com/agclqq/prow-framework/prowjob/register"
	"github.com/agclqq/prowjob"
)

func main() {
	pj := prowjob.New()
	register.Register(pj)
	pj.Run("init:project")
	os.Remove("main.go")
}
