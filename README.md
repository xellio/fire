## fire - send http requests
Send HTTP requests

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
