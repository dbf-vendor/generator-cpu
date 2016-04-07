package general

import (
	"github.com/dbf-vendor/generator-cpu/global"
	"log"
	"os"
	"os/exec"
)

var (
	cracker string = "hashcat"
)

func Main() {
	cracker = global.CURRENT_PATH + cracker + global.EXT

	cmd := exec.Command(cracker, os.Args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Printf("%s\n", err)
	}
}
