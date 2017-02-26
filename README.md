# uryify (originally [chrisify](github.com/zikes/chrisify))

## Linux Install

```bash
$ sudo apt install libopencv-dev
$ go get github.com/UniversityRadioYork/uryify
$ go get github.com/lazywei/go-opencv
$ cd $GOPATH/src/github.com/UniversityRadioYork/uryify && go build
```

## Usage

Input/Output images have to be `.jpg` and the faces `.png` (with a transparent background obviously).

For best effect, use the [template](template.png) as it gives a nice margin around the face.

Simplest: `./uryify --faces /path/to/faces path/to/image.jpg > output.jpg`

If executed from any location besides the repository, you must tell it where to find the bundled Haar Cascade face recognition XML file.

`uryify --faces /path/to/faces --haar /path/to/haarcascade_frontalface_alt.xml /path/to/input.jpg > output.jpg`