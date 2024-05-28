package sub

import (
	"log"

	"github.com/containerd/containerd"
)

// Const is a constant
const Const = "Hello, Sub World!"

func Func() {
	_, err := containerd.New("")
	if err != nil {
		log.Fatal(err)
	}

}
