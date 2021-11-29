package id

import gonanoid "github.com/matoous/go-nanoid/v2"

func ID() string {
	return gonanoid.Must()
}
