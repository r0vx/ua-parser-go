package uaparser

import (
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

// Sec-CH-UA* 请求头名称。
const (
	chBrands          = "sec-ch-ua"
	chFullVersionList = "sec-ch-ua-full-version-list"
	chArch            = "sec-ch-ua-arch"
	chBitness         = "sec-ch-ua-bitness"
	chFormFactors     = "sec-ch-ua-form-factors"
	chMobile          = "sec-ch-ua-mobile"
	chModel           = "sec-ch-ua-model"
	chPlatform        = "sec-ch-ua-platform"
	chPlatformVersion = "sec-ch-ua-platform-version"
	chUserAgent       = "user-agent"
)

// Brand 是 Sec-CH-UA 中的一个品牌项。
type Brand struct {
	Brand   string `json:"brand"`
	Version string `json:"version,omitempty"`
}

// ClientHints 保存从 Sec-CH-UA* 请求头解析出的客户端提示。
type ClientHints struct {
	Brands          []Brand  `json:"brands,omitempty"`
	FullVersionList []Brand  `json:"fullVersionList,omitempty"`
	Mobile          bool     `json:"mobile,omitempty"`
	Model           string   `json:"model,omitempty"`
	Platform        string   `json:"platform,omitempty"`
	PlatformVersion string   `json:"platformVersion,omitempty"`
	Architecture    string   `json:"architecture,omitempty"`
	FormFactors     []string `json:"formFactors,omitempty"`
	Bitness         string   `json:"bitness,omitempty"`
}

var (
	// headerQuoteRe 匹配可选反斜杠后的引号，用于去除头值中的引号。
	headerQuoteRe = regexp.MustCompile(`\\?"`)
	// notABrandRe 匹配 Chromium 的填充品牌（Not A Brand 的各种写法）。
	notABrandRe = regexp.MustCompile(`(?i)not.a.brand`)
	// chromiumRe 判断浏览器名是否属于 Chromium 系。
	chromiumRe = regexp.MustCompile(`Chrom`)
	// webView2Re 判断品牌是否为 Edge WebView2。
	webView2Re = regexp.MustCompile(`WebView2`)
)

// NewClientHints 从请求头集合解析客户端提示。头名称大小写不敏感。
func NewClientHints(headers map[string]string) ClientHints {
	return ClientHints{
		Brands:          parseBrandList(headerValue(headers, chBrands)),
		FullVersionList: parseBrandList(headerValue(headers, chFullVersionList)),
		Mobile:          strings.Contains(headerValue(headers, chMobile), "?1"),
		Model:           normalizeHeaderValue(headerValue(headers, chModel)),
		Platform:        normalizeHeaderValue(headerValue(headers, chPlatform)),
		PlatformVersion: normalizeHeaderValue(headerValue(headers, chPlatformVersion)),
		Architecture:    normalizeHeaderValue(headerValue(headers, chArch)),
		FormFactors:     parseTokenList(headerValue(headers, chFormFactors)),
		Bitness:         normalizeHeaderValue(headerValue(headers, chBitness)),
	}
}

// NewClientHintsFromHTTP 从 http.Header 解析客户端提示。
func NewClientHintsFromHTTP(headers http.Header) ClientHints {
	return NewClientHints(map[string]string{
		chBrands:          headers.Get(chBrands),
		chFullVersionList: headers.Get(chFullVersionList),
		chArch:            headers.Get(chArch),
		chBitness:         headers.Get(chBitness),
		chFormFactors:     headers.Get(chFormFactors),
		chMobile:          headers.Get(chMobile),
		chModel:           headers.Get(chModel),
		chPlatform:        headers.Get(chPlatform),
		chPlatformVersion: headers.Get(chPlatformVersion),
		chUserAgent:       headers.Get(chUserAgent),
	})
}

// headerValue 大小写不敏感地读取头值。
func headerValue(headers map[string]string, name string) string {
	if v, ok := headers[name]; ok {
		return v
	}
	for k, v := range headers {
		if strings.EqualFold(k, name) {
			return v
		}
	}
	return ""
}

// normalizeHeaderValue 去除引号、前导空白，并截断到 500 字符以防止超长头导致的正则开销。
func normalizeHeaderValue(s string) string {
	if s == "" {
		return ""
	}
	s = strings.TrimLeft(headerQuoteRe.ReplaceAllString(s, ""), " \t\r\n")
	if len(s) > UAMaxLength {
		s = s[:UAMaxLength]
	}
	return s
}

// parseBrandList 解析带版本的品牌列表，格式如 `"Chromium";v="130"`。
func parseBrandList(header string) []Brand {
	tokens := parseTokenList(header)
	if len(tokens) == 0 {
		return nil
	}
	brands := make([]Brand, 0, len(tokens))
	for _, tok := range tokens {
		brand, version, found := strings.Cut(tok, ";v=")
		if found {
			brands = append(brands, Brand{Brand: strings.TrimSpace(brand), Version: strings.TrimSpace(version)})
		} else {
			brands = append(brands, Brand{Brand: tok})
		}
	}
	return brands
}

// parseTokenList 按逗号拆分头值并去除每项两端的空白。
func parseTokenList(header string) []string {
	normalized := normalizeHeaderValue(header)
	if normalized == "" {
		return nil
	}
	tokens := make([]string, 0, strings.Count(normalized, ",")+1)
	for _, tok := range strings.Split(normalized, ",") {
		if tok = strings.TrimSpace(tok); tok != "" {
			tokens = append(tokens, tok)
		}
	}
	return tokens
}

// applyBrowserCH 用客户端提示覆盖浏览器名称、版本与主版本号。
func applyBrowserCH(b *Browser, ch *ClientHints) {
	brands := ch.FullVersionList
	if len(brands) == 0 {
		brands = ch.Brands
	}
	prevName := ""
	for _, brand := range brands {
		name := brand.Brand
		if notABrandRe.MatchString(name) {
			continue
		}
		// 只接受首个品牌，或 Chromium 系之后出现的更具体品牌，或 Edge 之后的 WebView2。
		if prevName != "" &&
			!(chromiumRe.MatchString(prevName) && name != "Chromium") &&
			!(prevName == BrowserEdge && webView2Re.MatchString(name)) {
			continue
		}
		mapped := strMapper(name, browserHintsMap)
		current := b.Name
		// 已有非 Chromium 系浏览器时，不让 Chromium 系品牌覆盖它。
		if !(current != "" && !chromiumRe.MatchString(current) && chromiumRe.MatchString(mapped)) {
			b.Name = mapped
			b.Version = brand.Version
			b.Major = majorize(brand.Version)
		}
		prevName = mapped
	}
}

// applyEngineCH 用客户端提示中的 Chromium 版本覆盖引擎版本。
func applyEngineCH(e *Engine, ch *ClientHints) {
	brands := ch.FullVersionList
	if len(brands) == 0 {
		brands = ch.Brands
	}
	for _, brand := range brands {
		if brand.Brand == BrowserChromium {
			e.Version = brand.Version
		}
	}
}

// applyCPUCH 用客户端提示中的架构与位数覆盖 CPU 架构。
func applyCPUCH(c *CPU, ch *ClientHints) {
	arch := ch.Architecture
	if arch != "" {
		if ch.Bitness == "64" {
			arch += "64"
		}
		rgxMapper(c, arch+";", cpuRegexes)
	}
}

// applyDeviceCH 用客户端提示覆盖设备类型、型号与厂商。
func applyDeviceCH(d *Device, ch *ClientHints) {
	if ch.Mobile {
		d.Type = DeviceMobile
	}
	if ch.Model != "" {
		d.Model = ch.Model
		if d.Type == "" || d.Vendor == "" {
			var fromModel Device
			rgxMapper(&fromModel, "droid 9; "+ch.Model+")", deviceRegexes)
			if d.Type == "" && fromModel.Type != "" {
				d.Type = fromModel.Type
			}
			if d.Vendor == "" && fromModel.Vendor != "" {
				d.Vendor = fromModel.Vendor
			}
		}
	}
	if len(ch.FormFactors) > 0 {
		ff := ""
		for _, v := range ch.FormFactors {
			if ff = strMapper(v, formFactorsMap); ff != "" {
				break
			}
		}
		d.Type = ff
	}
}

// applyOSCH 用客户端提示覆盖操作系统名称与版本。
func applyOSCH(o *OS, ch *ClientHints) {
	if ch.Platform != "" {
		version := ch.PlatformVersion
		if ch.Platform == BrowserWindows {
			// Windows 11 仍上报 NT 版本 10.0；平台版本 >= 13 视为 Windows 11。
			if major, err := strconv.Atoi(majorize(version)); err == nil && major >= 13 {
				version = "11"
			} else {
				version = "10"
			}
		}
		o.Name = ch.Platform
		o.Version = version
	}
	if o.Name == BrowserWindows && ch.Model == "Xbox" {
		o.Name = "Xbox"
		o.Version = ""
	}
}
