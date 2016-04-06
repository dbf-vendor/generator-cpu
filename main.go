package main

import (
	"github.com/dbf-vendor/generator-cpu/general"
	"github.com/dbf-vendor/generator-cpu/global"
	"github.com/dbf-vendor/generator-cpu/policy"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "--generator-policy":
			global.GENERATOR = global.GENERATOR_POLICY
			os.Args = append(os.Args[:i], os.Args[i+1:]...) // Remove from args
		}
	}

	global.InitGlobal()

	switch global.GENERATOR {
	case global.GENERATOR_POLICY:
		policy.Main()
	default:
		general.Main()
	}
}
