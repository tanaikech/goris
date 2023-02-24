# goris

[![Build Status](https://travis-ci.com/tanaikech/goris.svg?branch=master)](https://travis-ci.com/tanaikech/goris)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENCE)

<a name="top"></a>

# Overview

This is a CLI tool to search for images with **Go**ogle **R**everse **I**mage **S**earch.

# Description

Images can be searched by image files and image URLs. Searched images display URLs and also can be downloaded as image files.

# How to Install

Download an executable file from [the release page](https://github.com/tanaikech/goris/releases) and put to a directory with path.

or

Use go get.

```bash
$ go install github.com/tanaikech/goris@latest
```

# Usage

Search images from an image file. You can select number of output URLs using `-n`. The maximum number of output URLs is 100. (The current default value is 10. So when `-n` is not used, 10 URLs are returned.)

```bash
$ goris s -f [image file] -n 50
```

Search images from an image URL.

```bash
$ goris s -u [image URL]
```

Download searched images from an image file. Following sample downloads 10 searched images using an image file.

```bash
$ goris s -f [image file] -d -n 10
```

Retrieve web pages with matching images on Google top page. When this is not used, images are retrieved.

```bash
$ goris s -u [image URL] -w
```

```bash
$ goris s -f [image file] -w
```

<a name="Licence"></a>

# Licence

[MIT](LICENCE)

<a name="Author"></a>

# Author

[TANAIKE](https://github.com/tanaikech)

If you have any questions and commissions for me, feel free to tell me using e-mail of tanaike@hotmail.com

<a name="Update_History"></a>

# Update History

- v1.0.0 (April 26, 2017)

  Initial release.

- v1.0.1 (May 16, 2017)

  1. A bugfix
     - When number of retrieved URLs is smaller than number of default output, an error had occurred. This was fixed. (Thank you! [Steve Davis](https://github.com/OptumCS))

- v1.1.0 (June 13, 2017)
  1. Add option
     - When images are matched to a searched image, web pages with matching images are retrieved. These are web pages displayed on Google top page. When this is not used, images are retrieved. This was added as a boolean option. (This was added by a request.)

<a name="v111"></a>

- v1.1.1 (January 3, 2020)

  1. It seems that the specification of `github.com/urfave/cli` was changed by the update of [https://github.com/urfave/cli](https://github.com/urfave/cli). By this, when `go get -u github.com/tanaikech/goris` is run, an error occurred. So this error was removed.

<a name="v200"></a>

- v2.0.0 (April 23, 2020)

  1. The specification for running the reverse image search was changed at Google side. By this, this application was also changed.

<a name="v300"></a>

- v3.0.0 (February 23, 2022)

  1. The specification for running the reverse image search was changed at Google side. By this, this application was also changed.

<a name="v301"></a>

- v3.0.1 (May 2, 2022)

  1. About the option `-w`, a bug was resolved. By this, the URLs of the related Web site are returned. As the default, 10 URLs are returned. If you want to retrieve more, please use the option `-n` like `-n 20`.

<a name="v302"></a>

- v3.0.2 (February 2, 2023)

  1. Modified go.mod and go.sum.

[TOP](#top)
