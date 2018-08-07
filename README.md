## fire - send http requests
Send HTTP requests

[![go report card](https://goreportcard.com/badge/github.com/xellio/fire "go report card")](https://goreportcard.com/report/github.com/xellio/fire)
[![MIT license](http://img.shields.io/badge/license-MIT-brightgreen.svg)](http://opensource.org/licenses/MIT)
[![GoDoc](https://godoc.org/github.com/xellio/fire?status.svg)](https://godoc.org/github.com/xellio/fire)


### build
```
make
```
if not using upx run
```
make build
```
the binary is generated in ./bin/

### usage
define requests in a JSON file
```
[
    {
        "METHOD": "GET",
        "URL": "http://www.prillwitz.io",
        "HEADERS": {
            "User-Agent": "fire"
        },        
        "PAYLOAD": {
            "block": "true"
        }
    }
]
```
Run
```
fire ./path/to/requests.json
```
Output
```
REQUEST:        [GET]   http://www.prillwitz.io?block=true
HEADERS:        {"User-Agent":["fire"]}
PAYLOAD:        null
RESPONSE:       503
```
