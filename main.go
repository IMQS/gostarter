package main

import (
	"github.com/IMQS/gostarter/starter"
	"github.com/IMQS/nf"
)

func main() {
	svc := starter.NewService()
	svc.LoadConfig()
	svc.Initialize()
	nf.RunService(svc.ListenAndServe)
}
