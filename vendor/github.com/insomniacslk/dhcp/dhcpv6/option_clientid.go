package dhcpv6

import (
	"fmt"
)

// OptClientID represents a Client Identifier option as defined by RFC 3315
// Section 22.2.
func OptClientID(d DUID) Option {
	return &optClientID{d}
}

type optClientID struct {
	DUID
}

func (*optClientID) Code() OptionCode {
	return OptionClientID
}

func (op *optClientID) String() string {
	return fmt.Sprintf("%s: %s", op.Code(), op.DUID)
}

// parseOptClientID builds an OptClientId structure from a sequence
// of bytes. The input data does not include option code and length
// bytes.
func parseOptClientID(data []byte) (*optClientID, error) {
	cid, err := DUIDFromBytes(data)
	if err != nil {
		return nil, err
	}
	return &optClientID{cid}, nil
}
