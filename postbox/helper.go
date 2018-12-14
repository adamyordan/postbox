package postbox

import (
	"encoding/binary"
	"os"
	"os/user"
	"path"
)

func itob(v uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) uint64 {
	return binary.BigEndian.Uint64(b)
}

func getPostboxDir() string {
	var dir string
	usr, err := user.Current()
	if err != nil {
		dir = "postbox"
	} else {
		dir = path.Join(usr.HomeDir, ".postbox")
	}
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0755)
	}
	return dir
}
