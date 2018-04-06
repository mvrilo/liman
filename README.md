# docps
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com) [![Build Status](https://travis-ci.org/salihciftci/docps.svg?branch=master)](https://travis-ci.org/salihciftci/docps) [![Go Report Card](https://goreportcard.com/badge/github.com/salihciftci/docps)](https://goreportcard.com/report/github.com/salihciftci/docps)

docps is basic docker monitoring web application. Written with Go.

### Features
- Process State
- Images
- Volumes

## Installation

``` bash
$ go get github.com/salihciftci/docps
$ go build
$ ./docps
```

### Deploy

``` bash
$ docker build -t docps .
$ docker run -dit --name docps --restart always -p 8080:8080 -v /var/run/docker.sock:/var/run/docker.sock docps
```
Note: the `-v /var/run/docker.sock:/var/run/docker.sock` option can be used in Linux environments only. 

## Screenshots

![](https://raw.githubusercontent.com/salihciftci/docps/master/screenshots/dashboard.png)
![](https://raw.githubusercontent.com/salihciftci/docps/master/screenshots/containers.png)
![](https://raw.githubusercontent.com/salihciftci/docps/master/screenshots/images.png)