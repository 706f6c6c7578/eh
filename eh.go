package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"hash"
	"io"
	"os"

	"github.com/martinlindhe/gogost/gost34112012256"
	"github.com/martinlindhe/gogost/gost34112012512"
	"github.com/tjfoc/gmsm/sm3"
	"golang.org/x/crypto/ripemd160"
)

func main() {
	var use512 bool
	var useSM3 bool
	var useRIPEMD160 bool
	flag.BoolVar(&use512, "512", false, "Use GOST R 34.11-2012 512-bit")
	flag.BoolVar(&useSM3, "sm3", false, "Use Chinese SM3 hash algorithm")
	flag.BoolVar(&useRIPEMD160, "ripemd160", false, "Use RIPEMD-160 hash algorithm")
	flag.Parse()

	var hasher hash.Hash

	switch {
	case useRIPEMD160:
		hasher = ripemd160.New()
	case useSM3:
		hasher = sm3.New()
	case use512:
		hasher = gost34112012512.New()
	default:
		hasher = gost34112012256.New()
	}

	processInput(hasher)
	hash := hasher.Sum(nil)

	fmt.Println(hex.EncodeToString(hash))
}

func processInput(hasher io.Writer) {
	if len(flag.Args()) > 0 {
		// Hash string from command line argument
		io.WriteString(hasher, flag.Args()[0])
	} else {
		// Hash data from stdin
		if _, err := io.Copy(hasher, os.Stdin); err != nil {
			fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
			os.Exit(1)
		}
	}
}
