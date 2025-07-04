package create

import (
	"fmt"
	"net"
	"strconv"

	"github.com/moby/moby/pkg/namesgenerator"
)

// getFreePort finds a free TCP port and checks if port+100 is also free.
// Returns the base port as a string, or an error if not found.
// Retries 100 times
func getFreePort() (int, error) {
	for i := 0; i < 100; i++ {
		l, err := net.Listen("tcp", ":0")
		if err != nil {
			return 0, fmt.Errorf("could not get free port: %v", err)
		}
		port := l.Addr().(*net.TCPAddr).Port
		l.Close()

		// Check if port+100 is also free
		// This is a requirement from GraphDB itself.
		// The additional port is used for clustering.
		plus100 := port + 100
		l2, err := net.Listen("tcp", ":"+strconv.Itoa(plus100))
		if err == nil {
			l2.Close()
			return port, nil
		}

		// If the port+100 was not free, it will retry again
	}
	return 0, fmt.Errorf("could not find suitable port pair")
}

// generateRandomName returns a string that consist of
// two words, concatenated with underscore.
func generateRandomName() string {
	return namesgenerator.GetRandomName(0)
}
