package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"hash"
	"io"
	"os"

	"github.com/c0mm4nd/go-ripemd"
	"github.com/martinlindhe/gogost/gost34112012256"
	"github.com/martinlindhe/gogost/gost34112012512"
	"github.com/tjfoc/gmsm/sm3"
)

func main() {
	var use512 bool
	var useSM3 bool
	var useRIPEMD160 bool
	var useRIPEMD256 bool
	var useRIPEMD320 bool
	flag.BoolVar(&use512, "512", false, "Use GOST R 34.11-2012 512-bit")
	flag.BoolVar(&useSM3, "sm3", false, "Use Chinese SM3 hash algorithm")
	flag.BoolVar(&useRIPEMD160, "ripemd160", false, "Use RIPEMD-160 hash algorithm")
	flag.BoolVar(&useRIPEMD256, "ripemd256", false, "Use RIPEMD-256 hash algorithm")
	flag.BoolVar(&useRIPEMD320, "ripemd320", false, "Use RIPEMD-320 hash algorithm")
	flag.Parse()

	var hasher hash.Hash

	switch {
	case useRIPEMD160:
		hasher = ripemd.New160()
	case useRIPEMD256:
		hasher = ripemd.New256()
	case useRIPEMD320:
		hasher = ripemd.New320()
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
