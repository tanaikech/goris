// Package ris (ris.go) :
// These methods are for retrieving images from url and file.
package ris

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/urfave/cli"
)

const (
	baseurl = "https://www.google.com"
)

// requestParams : Parameters for fetchURL
type requestParams struct {
	Method      string
	URL         string
	Contenttype string
	Data        io.Reader
	Client      *http.Client
}

// Imgdata : Image URL
type Imgdata struct {
	OU      string `json:"ou"`
	WebPage bool
}

// DefImg : Initialize imagdata.
func DefImg(c *cli.Context) *Imgdata {
	return &Imgdata{
		WebPage: c.Bool("webpages"),
	}
}

// fetchURL : Fetch method
func (r *requestParams) fetchURL() *http.Response {
	req, err := http.NewRequest(
		r.Method,
		r.URL,
		r.Data,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v. ", err)
		os.Exit(1)
	}
	if len(r.Contenttype) > 0 {
		req.Header.Set("Content-Type", r.Contenttype)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 Firefox/26.0")
	res, _ := r.Client.Do(req)
	return res
}

// ImgFromURL : Search images from an image URL
func (im *Imgdata) ImgFromURL(searchimage string) []string {
	var url string
	r := &requestParams{
		Method: "GET",
		URL:    baseurl + "/searchbyimage?&image_url=" + searchimage,
		Data:   nil,
		Client: &http.Client{
			Timeout:       time.Duration(10) * time.Second,
			CheckRedirect: func(req *http.Request, via []*http.Request) error { return errors.New("Redirect") },
		},
	}
	var res *http.Response
	for {
		res = r.fetchURL()
		if res.StatusCode == 200 {
			break
		}
		reurl, _ := res.Location()
		r.URL = reurl.String()
	}
	defer res.Body.Close()
	doc, _ := goquery.NewDocumentFromResponse(res)
	var ar []string
	if im.WebPage {
		ar = getWebPages(doc)
	} else {
		doc.Find(".iu-card-header").Each(func(_ int, s *goquery.Selection) {
			url, _ = s.Attr("href")
		})
		r.URL = baseurl + url
		r.Client = &http.Client{Timeout: time.Duration(10) * time.Second}
		res = r.fetchURL()
		doc, _ = goquery.NewDocumentFromResponse(res)
		doc.Find(".rg_meta").Each(func(_ int, s *goquery.Selection) {
			json.Unmarshal([]byte(s.Text()), &im)
			ar = append(ar, im.OU)
		})
	}
	return ar
}

// ImgFromFile : Search images from an image file
func (im *Imgdata) ImgFromFile(file string) []string {
	var url string
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	fs, err := os.Open(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v. ", err)
		os.Exit(1)
	}
	defer fs.Close()
	data, err := w.CreateFormFile("encoded_image", file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v. ", err)
		os.Exit(1)
	}
	if _, err = io.Copy(data, fs); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v. ", err)
		os.Exit(1)
	}
	w.Close()
	r := &requestParams{
		Method: "POST",
		URL:    baseurl + "/searchbyimage/upload",
		Data:   &b,
		Client: &http.Client{
			Timeout:       time.Duration(10) * time.Second,
			CheckRedirect: func(req *http.Request, via []*http.Request) error { return errors.New("Redirect") },
		},
		Contenttype: w.FormDataContentType(),
	}
	var res *http.Response
	for {
		res = r.fetchURL()
		if res.StatusCode == 200 {
			break
		}
		reurl, _ := res.Location()
		r.URL = reurl.String()
		r.Method = "GET"
		r.Data = nil
		r.Contenttype = ""
	}
	defer res.Body.Close()
	doc, _ := goquery.NewDocumentFromResponse(res)
	var ar []string
	if im.WebPage {
		ar = getWebPages(doc)
	} else {
		doc.Find(".iu-card-header").Each(func(_ int, s *goquery.Selection) {
			url, _ = s.Attr("href")
		})
		r.URL = baseurl + url
		r.Client = &http.Client{Timeout: time.Duration(10) * time.Second}
		res = r.fetchURL()
		doc, _ = goquery.NewDocumentFromResponse(res)
		doc.Find(".rg_meta").Each(func(_ int, s *goquery.Selection) {
			json.Unmarshal([]byte(s.Text()), &im)
			ar = append(ar, im.OU)
		})
	}
	return ar
}

// Download : Download image files from searched image URLs
func Download(r []string, c int) {
	var wg sync.WaitGroup
	dlch := make(chan string, len(r))
	workers := 2
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, dlch chan string) {
			defer wg.Done()
			var res *http.Response
			for {
				dlurl, fin := <-dlch
				if !fin {
					return
				}
				filename := filepath.Base(dlurl)
				conv := strings.Replace(strings.TrimSpace(dlurl), filename, "", -1)
				conv = strings.Replace(strings.TrimSpace(conv), "http://", "", -1)
				conv = strings.Replace(strings.TrimSpace(conv), "https://", "", -1)
				conv = strings.Replace(strings.TrimSpace(conv), "/", "_", -1)
				conv = strings.Replace(strings.TrimSpace(conv), ".", "-", -1)
				conv += filename
				r := &requestParams{
					Method: "GET",
					URL:    dlurl,
					Data:   nil,
					Client: &http.Client{Timeout: time.Duration(10) * time.Second},
				}
				res = r.fetchURL()
				body, err := ioutil.ReadAll(res.Body)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error: %v. ", err)
					os.Exit(1)
				}
				ioutil.WriteFile(conv, body, 0777)
			}
			defer res.Body.Close()
		}(&wg, dlch)
	}
	for i := 0; i < c; i++ {
		dlch <- r[i]
	}
	close(dlch)
	wg.Wait()
}

// getWebPages : Retrieve web pages with matching images on Google top page. When this is not used, images are retrieved.
func getWebPages(doc *goquery.Document) []string {
	var ar []string
	doc.Find("h3.r").Each(func(i int, s *goquery.Selection) {
		s.Find("a").Each(func(_ int, s *goquery.Selection) {
			url, _ := s.Attr("href")
			ar = append(ar, url)
		})
	})
	return ar
}
