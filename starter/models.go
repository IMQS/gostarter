package starter

import "github.com/IMQS/nf"

type frogType struct {
	nf.Model
	Description string
}

type frog struct {
	nf.Model
	Description string
	FrogTypeID  int64
}
