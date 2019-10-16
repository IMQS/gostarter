package starter

import "github.com/IMQS/nf/nfdb"

type frogType struct {
	nfdb.Model
	Description string
}

type frog struct {
	nfdb.Model
	Description string
	FrogTypeID  int64
}
