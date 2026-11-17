// Package uaparser provides User-Agent string parsing functionality.
// It can detect browser, engine, OS, CPU, and device information from User-Agent data.
//
// This is a Go port of UAParser.js v2.0.10 (https://github.com/faisalman/ua-parser-js)
package uaparser

import (
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

const (
	Version     = "2.0.10"
	UAMaxLength = 500
	Empty       = ""
	Unknown     = "?"
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
	BrowserTypeCLI         = "cli"
	BrowserTypeCrawler     = "crawler"
	BrowserTypeEmail       = "email"
	BrowserTypeInApp       = "inapp"
	BrowserTypeLibrary     = "library"
	BrowserTypeMediaPlayer = "mediaplayer"
	// BrowserTypeFetcher marks fetching tools such as Chrome Lighthouse.
	BrowserTypeFetcher = "fetcher"
)

// iOS 26+ freezes the OS version reported in Safari UAs at 18/19;
// the real version is taken from the Safari version token instead.
var (
	iOSFrozenVersionRe = regexp.MustCompile(`^1[89][^\d]`)
	safariVersionRe    = regexp.MustCompile(`\) Version\/((\d+)[\d\.]*)`)
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
	ua    string
	ch    *ClientHints
	rules *ruleSet
}

// ruleSet 保存叠加扩展后的规则集；nil 时使用核心规则。
type ruleSet struct {
	browser, cpu, device, engine, os []RegexRule
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

// ParseWithClientHints 解析 UA 并叠加 Sec-CH-UA* 客户端提示。
// ua 为空时使用 headers 中的 user-agent。
func ParseWithClientHints(ua string, headers map[string]string) *Result {
	if ua == "" {
		ua = headerValue(headers, chUserAgent)
	}
	return New(ua).SetClientHints(NewClientHints(headers)).GetResult()
}

// ParseWithHTTPHeaders 从 http.Header 解析 UA 与客户端提示。
func ParseWithHTTPHeaders(ua string, headers http.Header) *Result {
	if ua == "" {
		ua = headers.Get(chUserAgent)
	}
	return New(ua).SetClientHints(NewClientHintsFromHTTP(headers)).GetResult()
}

// UseExtension 在核心规则之前追加扩展规则，可多次调用；后加入的扩展优先。
func (p *UAParser) UseExtension(exts ...Extension) *UAParser {
	if p.rules == nil {
		p.rules = &ruleSet{
			browser: browserRegexes,
			cpu:     cpuRegexes,
			device:  deviceRegexes,
			engine:  engineRegexes,
			os:      osRegexes,
		}
	}
	for _, ext := range exts {
		p.rules.browser = prependRules(ext.Browser, p.rules.browser)
		p.rules.cpu = prependRules(ext.CPU, p.rules.cpu)
		p.rules.device = prependRules(ext.Device, p.rules.device)
		p.rules.engine = prependRules(ext.Engine, p.rules.engine)
		p.rules.os = prependRules(ext.OS, p.rules.os)
	}
	return p
}

// prependRules 返回 extra 在前、base 在后的新切片，避免修改调用方切片。
func prependRules(extra, base []RegexRule) []RegexRule {
	if len(extra) == 0 {
		return base
	}
	merged := make([]RegexRule, 0, len(extra)+len(base))
	merged = append(merged, extra...)
	return append(merged, base...)
}

// SetClientHints 设置客户端提示，后续 Get* 会在 UA 解析结果上叠加。
func (p *UAParser) SetClientHints(ch ClientHints) *UAParser {
	p.ch = &ch
	return p
}

// rulesFor 返回叠加扩展后的规则；未使用扩展时回退到核心规则。
func (p *UAParser) rulesFor(pick func(*ruleSet) []RegexRule, fallback []RegexRule) []RegexRule {
	if p.rules == nil {
		return fallback
	}
	return pick(p.rules)
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
	rgxMapper(&browser, p.ua, p.rulesFor(func(r *ruleSet) []RegexRule { return r.browser }, browserRegexes))
	if browser.Version != "" {
		browser.Major = majorize(browser.Version)
	}
	if p.ch != nil {
		applyBrowserCH(&browser, p.ch)
	}
	return browser
}

// GetCPU returns CPU architecture information
func (p *UAParser) GetCPU() CPU {
	cpu := CPU{}
	rgxMapper(&cpu, p.ua, p.rulesFor(func(r *ruleSet) []RegexRule { return r.cpu }, cpuRegexes))
	if p.ch != nil {
		applyCPUCH(&cpu, p.ch)
	}
	return cpu
}

// GetDevice returns device information
func (p *UAParser) GetDevice() Device {
	device := Device{}
	rgxMapper(&device, p.ua, p.rulesFor(func(r *ruleSet) []RegexRule { return r.device }, deviceRegexes))
	if p.ch != nil {
		applyDeviceCH(&device, p.ch)
	}
	return device
}

// GetEngine returns browser engine information
func (p *UAParser) GetEngine() Engine {
	engine := Engine{}
	rgxMapper(&engine, p.ua, p.rulesFor(func(r *ruleSet) []RegexRule { return r.engine }, engineRegexes))
	if p.ch != nil {
		applyEngineCH(&engine, p.ch)
	}
	return engine
}

// GetOS returns operating system information
func (p *UAParser) GetOS() OS {
	os := OS{}
	rgxMapper(&os, p.ua, p.rulesFor(func(r *ruleSet) []RegexRule { return r.os }, osRegexes))

	// Since iOS 26, Safari's UA reports the OS version as frozen at 18:
	// https://webkit.org/blog/17333/webkit-features-in-safari-26-0/#update-to-ua-string
	if os.Name == "iOS" && os.Version != "" && iOSFrozenVersionRe.MatchString(os.Version) {
		match := safariVersionRe.FindStringSubmatch(p.ua)
		if len(match) > 2 {
			if major, err := strconv.Atoi(match[2]); err == nil && major >= 26 {
				os.Version = match[1]
			}
		}
	}
	if p.ch != nil {
		applyOSCH(&os, p.ch)
	}

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
