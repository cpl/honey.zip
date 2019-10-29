# honey.zip

A minimal Go web-server that will respond to any and all requests with a simple `200` and a file (which I recommend to be a [zip bomb](https://en.wikipedia.org/wiki/Zip_bomb) with a masked name).

**Disclaimer**, do not use for malicious purposes. Only for educational, pen-testing and proof-of-concept purposes.

## Install

* You can download a pre-compiled binary from the [GitHub releases]()
* If you have Go installed, simply run `go get cpl.li/go/honey.zip`
* You can build from source by cloning the repo and running `go build .`

## Use

```bash
# this is the most basic way you can call honey.zip
$ honey.zip -fpath=/tmp/zipbomb.zip
```

```bash
$ honey.zip --help

Usage of honey.zip:
  -addr string
    	web server listening address (default ":8091")
  -delay duration
    	delay for the response
  -fname string
    	filename used to serve zip bomb (default "sys-memory-dump.zip")
  -fpath string
    	filepath to zip bomb
  -websrv string
    	web server header value (default "Apache/1.8")
```