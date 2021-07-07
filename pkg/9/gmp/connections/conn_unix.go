package connections

import (
	"net"

	"gogmplib/pkg/9/gmp"
	"gogmplib/pkg/9/gmp/connections/internal/implementation"
)

// NewUnixConnection returns an instance of `gmp.Connection`. The `socket` parameter refers to the file path of the Unix Socket
// where Openvas GVMD is listening. Ex: "/var/run/gvmd.sock"
func NewUnixConnection(socket string) (gmp.Connection, error) {
	conn, err := net.Dial("unix", socket)
	if err != nil {
		return nil, err
	}

	c := &implementation.Connection{}
	c.SetRawConn(conn)

	return c, nil
}
