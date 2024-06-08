# Introduction

A simple script written in Go for generating screenshot from URL using [Playwright](https://github.com/playwright-community/playwright-go) library.

## Arguments

* url: The URL to generate screenshot
* browser: The browser of choice (chromium, firefox, webkit). Default: chromium
* device: Which device to render the webpage (iphone, android, desktop), Default: desktop
* full_page: For full page screenshot. Disabled by default
* file_name: Name of the file to save the screenshot to
* delay: Number of seconds to wait before taking screenshot. Default: 0

## Examples:

* Default

      go run main.go --url="https://go.dev/" --file_name="go_dev.png"

* Full Page on Firefox

      go run main.go --url="https://go.dev/" --full_page --file_name="go_full_page.png" --browser="firefox"


* Iphone, Wait for 5 seconds

      go run main.go --url="https://go.dev/" --device="iphone" --file_name="go_iphone.png" --delay=5


* Full page on android

      go run main.go --url="https://go.dev/" --device="android" --full_page --file_name="go_mobile_full_page.png"

## References
* [Playwright Documentation](https://pkg.go.dev/github.com/playwright-community/playwright-go)
* [Example screenshot code](https://github.com/playwright-community/playwright-go/blob/main/examples/screenshot/main.go)  
