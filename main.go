package main

import (
	"flag"
	"fmt"
	"github.com/playwright-community/playwright-go"
	"log"
	"strings"
	"time"
)

func main() {
	url := flag.String("url", "", "The URL of the web page to be rendered")
	fileName := flag.String("file_name", "", "The name of the file to save screenshot")
	delay := flag.Float64("delay", 0.0, "Time in seconds to wait before rendering, default = 0")
	deviceType := flag.String("device", "desktop", "The type of device (supported values: iphone, android or desktop), default = desktop")
	browserType := flag.String("browser", "chromium", "The browser of choice (supported values: chromium, firefox, webkit), default = chromium")
	fullPage := flag.Bool("full_page", false, "Enable full page screenshot, disabled by default")

	flag.Parse()

	// required args
	if *url == "" {
		fmt.Println("Error: --url flag is required")
		flag.Usage()
		return
	}

	if *fileName == "" {
		fmt.Println("Error: --file_name flag is required")
		flag.Usage()
		return
	}

	path := *fileName

	fmt.Println("generating screenshot ...")

	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}

	// select device
	var browser playwright.Browser
	device := pw.Devices["Desktop Chrome"] // default
	if strings.ToLower(*browserType) == "firefox" {
		browser, err = pw.Firefox.Launch()
		device = pw.Devices["Desktop Firefox"]
	} else if strings.ToLower(*browserType) == "webkit" {
		browser, err = pw.WebKit.Launch()
		device = pw.Devices["Desktop Safari"]
	} else {
		browser, err = pw.Chromium.Launch()
	}

	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}

	// latest available versions
	if *deviceType != "" {
		if strings.ToLower(*deviceType) == "iphone" {
			device = pw.Devices["iPhone 14 Pro"]
		} else if strings.ToLower(*deviceType) == "android" {
			device = pw.Devices["Galaxy S9+"]
		}
	}

	context, err := browser.NewContext(playwright.BrowserNewContextOptions{
		Viewport:          device.Viewport,
		UserAgent:         playwright.String(device.UserAgent),
		DeviceScaleFactor: playwright.Float(device.DeviceScaleFactor),
		IsMobile:          playwright.Bool(device.IsMobile),
		HasTouch:          playwright.Bool(device.HasTouch),
	})
	if err != nil {
		log.Fatalf("could not create context: %v", err)
	}

	page, err := context.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}

	if _, err = page.Goto(*url, playwright.PageGotoOptions{
		WaitUntil: playwright.WaitUntilStateDomcontentloaded,
	}); err != nil {
		log.Fatalf("could visit %s: %v", url, err)
	}

	if *delay != 0.0 {
		time.Sleep(time.Duration(*delay*1000) * time.Millisecond)
	}

	// generate screenshot
	if _, err = page.Screenshot(playwright.PageScreenshotOptions{
		Path:     playwright.String(path),
		FullPage: fullPage,
	}); err != nil {
		log.Fatalf("could not create screenshot: %v", err)
	} else {
		fmt.Println("screenshot of " + *url + " saved to " + path)
	}
}
