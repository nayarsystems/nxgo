// +build js

package nxgo

import (
	"fmt"
	"net/url"

	"github.com/jaracil/wsck"
	"github.com/nayarsystems/nxgo/nxcore"
)

var ErrVersionIncompatible = fmt.Errorf("incompatible version")

func Dial(s string, _ interface{}) (*nxcore.NexusConn, error) {

	u, err := url.Parse(s)
	if err != nil {
		return nil, err
	}

	conn, err := websocket.Dial(u.String(), "http://gopherjs.nexus")

	if err != nil {
		return nil, err
	}

	nxconn := nxcore.NewNexusConn(conn)

	nxconn.NexusVersion = getNexusVersion(nxconn)
	if !isVersionCompatible(nxconn.NexusVersion) {
		return nxconn, ErrVersionIncompatible
	}

	return nxconn, nil
}
