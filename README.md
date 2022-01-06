# go-qrcode #

<img src='https://skip.org/img/nyancat-youtube-qr.png' align='right'>

Package qrcode implements a QR Code encoder. [![Build Status](https://travis-ci.org/kstenerud/go-qrcode.svg?branch=master)](https://travis-ci.org/kstenerud/go-qrcode)

A QR Code is a matrix (two-dimensional) barcode. Arbitrary content may be encoded, with URLs being a popular choice :)

Each QR Code contains error recovery information to aid reading damaged or obscured codes. There are four levels of error recovery: Low, medium, high and highest. QR Codes with a higher recovery level are more robust to damage, at the cost of being physically larger.

## Install

    go get -u github.com/kstenerud/go-qrcode/...

A command-line tool `qrcode` will be built into `$GOPATH/bin/`.

## Usage

    import qrcode "github.com/kstenerud/go-qrcode"

- **Create a 256x256 PNG image:**

        var png []byte
        png, err := qrcode.Encode([]byte("https://example.org"), qrcode.Medium, 256)

- **Create a 256x256 PNG image and write to a file:**

        err := qrcode.WriteFile([]byte("https://example.org"), qrcode.Medium, 256, "qr.png")

- **Create a 256x256 PNG image with custom colors and write to file:**

        err := qrcode.WriteColorFile([]byte("https://example.org"), qrcode.Medium, 256, color.Black, color.White, "qr.png")

All examples use the qrcode.Medium error Recovery Level and create a fixed 256x256px size QR Code. The last function creates a white on black instead of black on white QR Code.

## Important!

QR codes are only meant to support ISO 8859-1 text and a special kanji encoding! If you pass in UTF-8 text, it may not be decoded properly by a QR code scanner. Characters from the ASCII range (0x20-0x7e) will always work.

## Documentation

[![godoc](https://godoc.org/github.com/kstenerud/go-qrcode?status.png)](https://godoc.org/github.com/kstenerud/go-qrcode)

## Demoapp

[http://go-qrcode.appspot.com](http://go-qrcode.appspot.com)

## CLI

A command-line tool `qrcode` will be built into `$GOPATH/bin/`.

```
qrcode -- QR Code encoder in Go

Flags:
  -b int
        QR Code border size (default 4)
  -e int
        error recovery level: 0=lowest, 3=highest (default 3)
  -i    invert black and white
  -o string
        out PNG file prefix, empty for stdout
  -s int
        image size (pixel) (default 256)
  -t    print as text-art on stdout

Usage:
  1. Data is read from STDIN and used to generate the QR code.
     Default output is STDOUT.

     echo "hello world" | qrcode > out.png

  2. Pipe to imagemagick command "display" to display on any X server.

     echo "hello world" | qrcode | display
```
## Maximum capacity
The maximum capacity of a QR Code varies according to the content encoded and the error recovery level. The maximum capacity is 2,953 bytes, 4,296 alphanumeric characters, 7,089 numeric digits, or a combination of these.

## Borderless QR Codes

To aid QR Code reading software, QR codes have a built in whitespace border. To disable it completely, set the border size to 0.

## Links

- [http://en.wikipedia.org/wiki/QR_code](http://en.wikipedia.org/wiki/QR_code)
- [ISO/IEC 18004:2006](http://www.iso.org/iso/catalogue_detail.htm?csnumber=43655) - Main QR Code specification (approx CHF 198,00)<br>
- [https://github.com/qpliu/qrencode-go/](https://github.com/qpliu/qrencode-go/) - alternative Go QR encoding library based on [ZXing](https://github.com/zxing/zxing)
