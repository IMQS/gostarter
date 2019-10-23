package starter

import "github.com/IMQS/nf/nfdb"

type frogType struct {
	nfdb.Model
	Description *string `json:"description"`
}

type frog struct {
	nfdb.Model
	Description *string `json:"description"`
	FrogTypeID  *int64  `json:"frogTypeID"`
}
