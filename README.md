
# Simple url shortener bindings for go - [![Go](https://github.com/illiafox/shorturl/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/illiafox/gostrings/actions/workflows/go.yml)
##### *Services:* 3 | *Last version:* v0.0.3 (26.09.2021)
## QuickList
* [Install](#install)
* [Usage](#usage)
* [Errors](#errors)
#### Services:
* [Cleanuri](#cleanuri)
* [Gotiny](#gotiny)
* [OnePt](#onept)
## Install:
```
go get github.com/illiafox/shorturl
```

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
returns short url (cleanuri.com/) , error (nil - no error)
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
   Others: **[⬇ Errors](#errors)**

**[⬆ Back to QuickList](#quicklist)**

## GoTiny:
#### Limits: Unknown
 ```
GoTiny(url string) (string,err)
url = Url | Text to shorten

if url: https://www.example.me
else: "text"

returns 6-digit code link (gotiny.cc/abc123) , error (nil - no error)
```
You can short you **text** instead of link

  * **Usage:**
```go
link, err := shorturl.GoTiny("https://github.com/illiafox")
if err != nil {
// panic
}

```

```
OUTPUT: gotiny.cc/a63c2k
```

```go
link, err := shorturl.GoTiny("Hello World!")
if err != nil {
// panic
}

```


 **Errors:** only from **[⬇ Errors](#errors)**
   
**[⬆ Back to QuickList](#quicklist)**




## OnePt (currently not working)
#### Warning: couldn't check this function due to server 500 error (if works please tell about this in [issues](https://github.com/illiafox/shorturl/issues) )
#### Limits: unknown
 ```
Pt1(url string,short string) (string,err)

url = URL to shorten https://www.example.me

short(optional) - The part after 1pt.co/ that will redirect to your long URL

returns short url (cleanuri.com/) , error (nil - no error)
```
If paramter **short** is not provided or the requested short URL is already taken, it will return a **random 5-letter**

  * **Usage (NOT TESTED):**
  ```go
link, err := shorturl.Pt1("https://github.com/illiafox","abc12") // or ("link","")
if err != nil {
// panic
}
```

```
OUTPUT (not checked): 1pt.co/XgpYj6
```

 * **Errors:**
 There are only one error that we know: **500 code**

Nowadays it always returns this code

Others: **[⬇ Errors](#errors)**
   
**[⬆ Back to QuickList](#quicklist)**


## Errors
These bindings can return errors from:
* [net/http](https://pkg.go.dev/net/http)
* [encoding/json](https://pkg.go.dev/encoding/json)
* [io/ioutill](https://pkg.go.dev/io/ioutil)
