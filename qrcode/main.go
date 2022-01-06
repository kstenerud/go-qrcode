// go-qrcode
// Copyright 2014 Tom Harwood

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	qrcode "github.com/skip2/go-qrcode"
)

func main() {
	outFile := flag.String("o", "", "out PNG file prefix, empty for stdout")
	size := flag.Int("s", 256, "image size (pixel)")
	textArt := flag.Bool("t", false, "print as text-art on stdout")
	negative := flag.Bool("i", false, "invert black and white")
	borderSize := flag.Int("b", qrcode.DefaultQuietZoneSize, "QR Code border size")
	erLevel := flag.Int("e", int(qrcode.Highest), "error recovery level (1=lowest, 4=highest")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `qrcode -- QR Code encoder in Go

Flags:
`)
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, `
Usage:
  1. Data is read from STDIN and used to generate the QR code.
     Default output is STDOUT.

	 echo "hello world" | qrcode > out.png

  2. Pipe to imagemagick command "display" to display on any X server.

     echo "hello world" | qrcode | display

`)
	}
	flag.Parse()

	recoveryLevel := qrcode.RecoveryLevel(*erLevel)
	if recoveryLevel < qrcode.Low || recoveryLevel > qrcode.Highest {
		panic(fmt.Errorf("error recovery level %v is out of range (%v to %v)", recoveryLevel, qrcode.Low, qrcode.Highest))
	}

	content, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(fmt.Errorf("could not read from stdin: %v", err))
	}

	var q *qrcode.QRCode
	q, err = qrcode.New(content, recoveryLevel)
	checkError(err)

	q.BorderSize = *borderSize

	if *textArt {
		art := q.ToString(*negative)
		fmt.Println(art)
		return
	}

	if *negative {
		q.ForegroundColor, q.BackgroundColor = q.BackgroundColor, q.ForegroundColor
	}

	var png []byte
	png, err = q.PNG(*size)
	checkError(err)

	if *outFile == "" {
		os.Stdout.Write(png)
	} else {
		var fh *os.File
		fh, err = os.Create(*outFile + ".png")
		checkError(err)
		defer fh.Close()
		fh.Write(png)
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
