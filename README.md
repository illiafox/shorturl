
# Simple url shortener for go - [![Go](https://github.com/illiafox/shorturl/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/illiafox/gostrings/actions/workflows/go.yml)
##### *Services:* 3 | *Last version:* v0.0.3 (26.09.2021)
## QuickList
* [Install](#install)
* [Usage](#usage)
#### Services:
* [Cleanuri](#cleanuri)

## Install:
```
go get github.com/illiafox/shorturl
```
**[⬆ Back to QuickList](#quicklist)**

## Usage:
```go
import "github.com/illiafox/shorturl"
shorturl.func()
```


## Cleanuri:
#### Limits: 2 requests per second (per IP)
 ```
Cleanuri(url string) (string,err)

url = URL to shorten https://www.example.me
returns short url (cleanuri.com/) and error (nil - no error)
```

  * **Usage:**
  ```go
link, err := shorturl.Cleanuri("https://github.com/illiafox")
if err != nil {
// panic
}
```

```
OUTPUT: https://cleanuri.com/XgpYj6
```

 * **Errors:**
 
   All api errors is **json**:
   ```json
   {"error":"API Error: URL is invalid. (check #2)"}
   ```
   Other errors are from [net/http](https://pkg.go.dev/net/http) package
   
**[⬆ Back to QuickList](#quicklist)**
