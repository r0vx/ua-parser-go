// Package uaparser provides User-Agent string parsing functionality.
// It can detect browser, engine, OS, CPU, and device information from User-Agent data.
//
// This is a Go port of UAParser.js v2.0.7 (https://github.com/faisalman/ua-parser-js)
package uaparser

import (
	"regexp"
	"strings"
)

const (
	Version      = "2.0.7"
	UAMaxLength  = 500
	Empty        = ""
	Unknown      = "?"
)

// Property names
const (
	PropName         = "name"
	PropType         = "type"
	PropVendor       = "vendor"
	PropVersion      = "version"
	PropArchitecture = "architecture"
	PropMajor        = "major"
	PropModel        = "model"
)

// Device types
const (
	DeviceConsole  = "console"
	DeviceMobile   = "mobile"
	DeviceTablet   = "tablet"
	DeviceSmartTV  = "smarttv"
	DeviceWearable = "wearable"
	DeviceXR       = "xr"
	DeviceEmbedded = "embedded"
)

// Browser types
const (
	BrowserTypeInApp = "inapp"
)

// Browser represents browser information
type Browser struct {
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
	Major   string `json:"major,omitempty"`
	Type    string `json:"type,omitempty"`
}

// CPU represents CPU architecture information
type CPU struct {
	Architecture string `json:"architecture,omitempty"`
}

// Device represents device information
type Device struct {
	Type   string `json:"type,omitempty"`
	Model  string `json:"model,omitempty"`
	Vendor string `json:"vendor,omitempty"`
}

// Engine represents browser engine information
type Engine struct {
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
}

// OS represents operating system information
type OS struct {
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
}

// Result represents the complete parsing result
type Result struct {
	UA      string  `json:"ua"`
	Browser Browser `json:"browser"`
	CPU     CPU     `json:"cpu"`
	Device  Device  `json:"device"`
	Engine  Engine  `json:"engine"`
	OS      OS      `json:"os"`
}

// UAParser is the main parser struct
type UAParser struct {
	ua string
}

// New creates a new UAParser with the given user agent string
func New(ua string) *UAParser {
	if len(ua) > UAMaxLength {
		ua = ua[:UAMaxLength]
	}
	return &UAParser{ua: strings.TrimSpace(ua)}
}

// Parse is a convenience function to parse a user agent string directly
func Parse(ua string) *Result {
	return New(ua).GetResult()
}

// SetUA sets the user agent string
func (p *UAParser) SetUA(ua string) *UAParser {
	if len(ua) > UAMaxLength {
		ua = ua[:UAMaxLength]
	}
	p.ua = strings.TrimSpace(ua)
	return p
}

// GetUA returns the current user agent string
func (p *UAParser) GetUA() string {
	return p.ua
}

// GetBrowser returns browser information
func (p *UAParser) GetBrowser() Browser {
	browser := Browser{}
	rgxMapper(&browser, p.ua, browserRegexes)
	if browser.Version != "" {
		browser.Major = majorize(browser.Version)
	}
	return browser
}

// GetCPU returns CPU architecture information
func (p *UAParser) GetCPU() CPU {
	cpu := CPU{}
	rgxMapper(&cpu, p.ua, cpuRegexes)
	return cpu
}

// GetDevice returns device information
func (p *UAParser) GetDevice() Device {
	device := Device{}
	rgxMapper(&device, p.ua, deviceRegexes)
	return device
}

// GetEngine returns browser engine information
func (p *UAParser) GetEngine() Engine {
	engine := Engine{}
	rgxMapper(&engine, p.ua, engineRegexes)
	return engine
}

// GetOS returns operating system information
func (p *UAParser) GetOS() OS {
	os := OS{}
	rgxMapper(&os, p.ua, osRegexes)
	return os
}

// GetResult returns the complete parsing result
func (p *UAParser) GetResult() *Result {
	return &Result{
		UA:      p.ua,
		Browser: p.GetBrowser(),
		CPU:     p.GetCPU(),
		Device:  p.GetDevice(),
		Engine:  p.GetEngine(),
		OS:      p.GetOS(),
	}
}

// Helper functions

func lowerize(s string) string {
	return strings.ToLower(s)
}

func majorize(version string) string {
	if version == "" {
		return ""
	}
	// Remove non-numeric prefix
	re := regexp.MustCompile(`[^\d\.]`)
	cleaned := re.ReplaceAllString(version, "")
	parts := strings.Split(cleaned, ".")
	if len(parts) > 0 && parts[0] != "" {
		return parts[0]
	}
	return ""
}

func trim(s string) string {
	return strings.TrimSpace(s)
}

func has(arr []string, str string) bool {
	str = lowerize(str)
	for _, a := range arr {
		if lowerize(a) == str {
			return true
		}
	}
	return false
}

