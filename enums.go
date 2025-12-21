package uaparser

// BrowserName constants for common browser names
var BrowserName = struct {
	Chrome         string
	ChromeHeadless string
	ChromeMobile   string
	ChromeWebView  string
	Chromium       string
	Edge           string
	EdgeWebView    string
	Firefox        string
	FirefoxFocus   string
	FirefoxMobile  string
	Opera          string
	OperaMini      string
	Safari         string
	SafariMobile   string
	Samsung        string
	UCBrowser      string
	Yandex         string
	IE             string
	Brave          string
	Vivaldi        string
	DuckDuckGo     string
	WeChat         string
	QQBrowser      string
	Baidu          string
}{
	Chrome:         "Chrome",
	ChromeHeadless: "Chrome Headless",
	ChromeMobile:   "Mobile Chrome",
	ChromeWebView:  "Chrome WebView",
	Chromium:       "Chromium",
	Edge:           "Edge",
	EdgeWebView:    "Edge WebView",
	Firefox:        "Firefox",
	FirefoxFocus:   "Firefox Focus",
	FirefoxMobile:  "Mobile Firefox",
	Opera:          "Opera",
	OperaMini:      "Opera Mini",
	Safari:         "Safari",
	SafariMobile:   "Mobile Safari",
	Samsung:        "Samsung Internet",
	UCBrowser:      "UCBrowser",
	Yandex:         "Yandex",
	IE:             "IE",
	Brave:          "Brave",
	Vivaldi:        "Vivaldi",
	DuckDuckGo:     "DuckDuckGo",
	WeChat:         "WeChat",
	QQBrowser:      "QQBrowser",
	Baidu:          "Baidu",
}

// BrowserType constants for browser types
var BrowserType = struct {
	Crawler     string
	CLI         string
	Email       string
	Fetcher     string
	InApp       string
	MediaPlayer string
	Library     string
}{
	Crawler:     "crawler",
	CLI:         "cli",
	Email:       "email",
	Fetcher:     "fetcher",
	InApp:       "inapp",
	MediaPlayer: "mediaplayer",
	Library:     "library",
}

// CPUArch constants for CPU architectures
var CPUArch = struct {
	AMD64  string
	IA32   string
	ARM64  string
	ARMHF  string
	ARM    string
	SPARC  string
	IA64   string
	MIPS   string
	PPC    string
	Alpha  string
}{
	AMD64:  "amd64",
	IA32:   "ia32",
	ARM64:  "arm64",
	ARMHF:  "armhf",
	ARM:    "arm",
	SPARC:  "sparc",
	IA64:   "ia64",
	MIPS:   "mips",
	PPC:    "ppc",
	Alpha:  "alpha",
}

// DeviceType constants for device types
var DeviceType = struct {
	Console  string
	Desktop  string
	Embedded string
	Mobile   string
	SmartTV  string
	Tablet   string
	Wearable string
	XR       string
}{
	Console:  "console",
	Desktop:  "desktop",
	Embedded: "embedded",
	Mobile:   "mobile",
	SmartTV:  "smarttv",
	Tablet:   "tablet",
	Wearable: "wearable",
	XR:       "xr",
}

// DeviceVendor constants for device vendors
var DeviceVendor = struct {
	Amazon    string
	Apple     string
	Asus      string
	BlackBerry string
	Google    string
	Huawei    string
	Lenovo    string
	LG        string
	Microsoft string
	Motorola  string
	Nokia     string
	OnePlus   string
	Oppo      string
	Samsung   string
	Sony      string
	Xiaomi    string
}{
	Amazon:    "Amazon",
	Apple:     "Apple",
	Asus:      "ASUS",
	BlackBerry: "BlackBerry",
	Google:    "Google",
	Huawei:    "Huawei",
	Lenovo:    "Lenovo",
	LG:        "LG",
	Microsoft: "Microsoft",
	Motorola:  "Motorola",
	Nokia:     "Nokia",
	OnePlus:   "OnePlus",
	Oppo:      "OPPO",
	Samsung:   "Samsung",
	Sony:      "Sony",
	Xiaomi:    "Xiaomi",
}

// EngineName constants for browser engine names
var EngineName = struct {
	Blink    string
	EdgeHTML string
	Gecko    string
	Presto   string
	Trident  string
	WebKit   string
}{
	Blink:    "Blink",
	EdgeHTML: "EdgeHTML",
	Gecko:    "Gecko",
	Presto:   "Presto",
	Trident:  "Trident",
	WebKit:   "WebKit",
}

// OSName constants for operating system names
var OSName = struct {
	Android   string
	ChromeOS  string
	iOS       string
	Linux     string
	MacOS     string
	Windows   string
	WindowsPhone string
}{
	Android:   "Android",
	ChromeOS:  "Chrome OS",
	iOS:       "iOS",
	Linux:     "Linux",
	MacOS:     "macOS",
	Windows:   "Windows",
	WindowsPhone: "Windows Phone",
}

