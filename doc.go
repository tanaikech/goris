/*
Package main (doc.go) :
This is a CLI tool to search for images with Google Reverse Image Search (goris).

Images can be searched by image files and image URLs. Searched images display URLs and also can be downloaded as image files.

---------------------------------------------------------------

# Usage
Search images from an image file. You can select number of output URLs using ``-n``. The maximum number of output URLs is 100.

$ goris s -f [iamge file] -n 50

Search images from an image URL.

$ goris s -u [iamge URL]

Download searched images from an image file. Following sample downloads 10 searched images using an image file.

$ goris s -f [iamge file] -d -n 10

---------------------------------------------------------------
*/
package main
