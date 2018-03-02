package main

import (
	"fmt"
	"os"
	"log"
)

func main() {

	fmt.Println("Programmet skal returnere informasjon om en fil")

	Info()
}

func Info() {
	fi, err := os.Lstat("Oblig2/src/oppgave1/randomstuff.txt")
	if err != nil {
		log.Fatal(err)
	}

	mode := fi.Mode()

	if mode.IsRegular() {
		fmt.Println("is a regular file")
	} else {
		fmt.Println("is not a regular file")
	}

	if mode.IsDir() {
		fmt.Println("is a directory")
	} else {
		fmt.Println("is not a directory")
	}

	if mode&os.ModeSymlink != 0 {
		fmt.Println("is a symbolic link")
	} else {
		fmt.Println("is not a symbolic link")
	}

	if mode&os.ModeNamedPipe != 0 {
		fmt.Println("is a named pipe")
	} else {
		fmt.Println("is not a named pipe")
	}

	if mode&os.ModeAppend != 0 {
		fmt.Println("is append-only")
	} else {
		fmt.Println("is not append-only")
	}

	if mode&os.ModeDevice != 0 {
		fmt.Println("is a device file")
	} else {
		fmt.Println("is not a device file")
	}

	/*
	case mode.IsDir():
		fmt.Println("directory")
	case mode&os.ModeSymlink != 0:
		fmt.Println("symbolic link")
	case mode&os.ModeNamedPipe != 0:
		fmt.Println("named pipe")*/

}

func mainn() {
	fi, err := os.Lstat("Oblig2/src/oppgave1/randomstuff.txt")
	if err != nil {
		log.Fatal(err)
	}

	switch mode := fi.Mode(); {
	case mode.IsRegular():
		fmt.Println("regular file")
	case mode.IsDir():
		fmt.Println("directory")
	case mode&os.ModeSymlink != 0:
		fmt.Println("symbolic link")
	case mode&os.ModeNamedPipe != 0:
		fmt.Println("named pipe")
	}
}
