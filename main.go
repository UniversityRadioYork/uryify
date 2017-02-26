package main

import (
	"flag"
	"image"
	"image/draw"
	"image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"

	"github.com/UniversityRadioYork/uryify/facefinder"

	"github.com/disintegration/imaging"
)

var haarCascade = flag.String("haar", "haarcascade_frontalface_alt.xml", "The location of the Haar Cascade XML configuration to be provided to OpenCV.")
var facesDir = flag.String("faces", "", "The directory to search for faces.")

func main() {
	flag.Parse()

	var faces FaceList

	var facesPath string
	var err error

	if *facesDir != "" {
		facesPath, err = filepath.Abs(*facesDir)
		if err != nil {
			panic(err)
		}
	}

	err = faces.Load(facesPath)
	if err != nil {
		panic(err)
	}
	if len(faces) == 0 {
		panic("no faces found")
	}

	file := flag.Arg(0)

	finder := facefinder.NewFinder(*haarCascade)

	baseImage := loadImage(file)

	facesFound := finder.Detect(baseImage)

	bounds := baseImage.Bounds()

	canvas := canvasFromImage(baseImage)

	for _, faceFound := range facesFound {

		h := int(0.8 * float64(faceFound.Dy()))

		faceFound.Min.Y += h
		faceFound.Max.Y += h

		rect := rectMargin(100.0, faceFound)

		newFace := faces.Random()
		if newFace == nil {
			panic("nil face")
		}
		chrisFace := imaging.Fit(newFace, rect.Dx(), rect.Dy(), imaging.Lanczos)

		draw.Draw(
			canvas,
			rect,
			chrisFace,
			bounds.Min,
			draw.Over,
		)
	}

	if len(facesFound) == 0 {
		face := imaging.Resize(
			faces.Random(),
			bounds.Dx()/3,
			0,
			imaging.Lanczos,
		)
		face_bounds := face.Bounds()
		draw.Draw(
			canvas,
			bounds,
			face,
			bounds.Min.Add(image.Pt(-bounds.Max.X/2+face_bounds.Max.X/2, -bounds.Max.Y+int(float64(face_bounds.Max.Y)/1.9))),
			draw.Over,
		)
	}

	jpeg.Encode(os.Stdout, canvas, &jpeg.Options{jpeg.DefaultQuality})
}
