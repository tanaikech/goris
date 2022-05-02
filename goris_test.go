package main

import (
	"testing"

	"goris/ris"
)

func TestImgFromURL(t *testing.T) {
	// This sample image is a logo of Stackoverflow. https://stackoverflow.design/brand/logo/
	imageurl := "https://stackoverflow.design/assets/img/logos/so/logo-stackoverflow.png"
	webpages := false
	results, _ := ris.DefImg(webpages).ImgFromURL(imageurl)
	t.Log(results)
}
