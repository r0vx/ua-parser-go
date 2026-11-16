package uaparser

import (
	"encoding/json"
	"os"
	"strings"
	"testing"
)

// chCase 对应上游 test/data/ua-ch/headers.js 中的用例。
type chCase struct {
	Desc    string            `json:"desc"`
	Headers map[string]string `json:"headers"`
	Expect  struct {
		Browser struct {
			Name    string `json:"name"`
			Version string `json:"version"`
			Major   string `json:"major"`
			Type    string `json:"type"`
		} `json:"browser"`
	} `json:"expect"`
}

// TestClientHintsCases 逐条校验上游 UA-CH 测试数据。
func TestClientHintsCases(t *testing.T) {
	data, err := os.ReadFile("testdata/ua-ch.json")
	if err != nil {
		t.Fatalf("读取测试数据失败: %v", err)
	}
	var cases []chCase
	if err := json.Unmarshal(data, &cases); err != nil {
		t.Fatalf("解析测试数据失败: %v", err)
	}
	for _, tc := range cases {
		t.Run(tc.Desc, func(t *testing.T) {
			b := ParseWithClientHints("", tc.Headers).Browser
			if b.Name != tc.Expect.Browser.Name {
				t.Errorf("name = %q, want %q", b.Name, tc.Expect.Browser.Name)
			}
			if b.Version != tc.Expect.Browser.Version {
				t.Errorf("version = %q, want %q", b.Version, tc.Expect.Browser.Version)
			}
			if b.Major != tc.Expect.Browser.Major {
				t.Errorf("major = %q, want %q", b.Major, tc.Expect.Browser.Major)
			}
		})
	}
}

// TestClientHintsFull 覆盖浏览器、CPU、设备、引擎、系统的完整叠加。
func TestClientHintsFull(t *testing.T) {
	headers := map[string]string{
		"sec-ch-ua":                   `"Chromium";v="93", "Google Chrome";v="93", " Not;A Brand";v="99"`,
		"sec-ch-ua-full-version-list": `"Chromium";v="93.0.1.2", "Google Chrome";v="93.0.1.2", " Not;A Brand";v="99.0.1.2"`,
		"sec-ch-ua-arch":              `"arm"`,
		"sec-ch-ua-bitness":           `"64"`,
		"sec-ch-ua-mobile":            "?1",
		"sec-ch-ua-model":             `"Pixel 99"`,
		"sec-ch-ua-platform":          `"Windows"`,
		"sec-ch-ua-platform-version":  `"13"`,
		"user-agent":                  "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36",
	}
	r := ParseWithClientHints("", headers)

	if r.UA != headers["user-agent"] {
		t.Fatalf("UA = %q, want user-agent header", r.UA)
	}
	if r.Browser.Name != "Chrome" || r.Browser.Version != "93.0.1.2" || r.Browser.Major != "93" {
		t.Errorf("browser = %+v", r.Browser)
	}
	if r.CPU.Architecture != "arm64" {
		t.Errorf("cpu = %q, want arm64", r.CPU.Architecture)
	}
	if r.Device.Type != "mobile" || r.Device.Model != "Pixel 99" || r.Device.Vendor != "Google" {
		t.Errorf("device = %+v", r.Device)
	}
	if r.Engine.Name != "Blink" || r.Engine.Version != "93.0.1.2" {
		t.Errorf("engine = %+v", r.Engine)
	}
	if r.OS.Name != "Windows" || r.OS.Version != "11" {
		t.Errorf("os = %+v", r.OS)
	}
}

// TestClientHintsBrandOrder 校验不同品牌顺序下仍选中最具体的品牌。
func TestClientHintsBrandOrder(t *testing.T) {
	cases := []struct {
		list      string
		name, ver string
	}{
		{`"Not_A Brand;v=8, Chromium;v=120.0.6099.131, Google Chrome;v=120.0.6099.132"`, "Chrome", "120.0.6099.132"},
		{`"Chromium;v=120.0.6099.131, Not_A Brand;v=8, Google Chrome;v=120.0.6099.132"`, "Chrome", "120.0.6099.132"},
		{`"Google Chrome;v=120.0.6099.132, Chromium;v=120.0.6099.131, Not_A Brand;v=8"`, "Chrome", "120.0.6099.132"},
		{`"Microsoft Edge;v=120.0.6099.133, Google Chrome;v=120.0.6099.132, Chromium;v=120.0.6099.131, Not_A Brand;v=8"`, "Edge", "120.0.6099.133"},
		{`"Chromium;v=120.0.6099.131, Google Chrome;v=120.0.6099.132, Microsoft Edge;v=120.0.6099.133, Not_A Brand;v=8"`, "Edge", "120.0.6099.133"},
	}
	for _, tc := range cases {
		b := ParseWithClientHints("", map[string]string{"sec-ch-ua-full-version-list": tc.list}).Browser
		if b.Name != tc.name || b.Version != tc.ver {
			t.Errorf("%s: got %s %s, want %s %s", tc.list, b.Name, b.Version, tc.name, tc.ver)
		}
	}
}

// TestClientHintsFormFactors 校验设备类型来自 form-factors。
func TestClientHintsFormFactors(t *testing.T) {
	cases := map[string]string{
		`"VR"`:             "xr",
		`"Tablet", "EInk"`: "tablet",
		`"Unknown"`:        "",
	}
	for header, want := range cases {
		got := ParseWithClientHints("", map[string]string{"sec-ch-ua-form-factors": header}).Device.Type
		if got != want {
			t.Errorf("form-factors %s: type = %q, want %q", header, got, want)
		}
	}
}

// TestClientHintsFallback 校验无 CH 头时回退到 UA。
func TestClientHintsFallback(t *testing.T) {
	ua := "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36"
	r := ParseWithClientHints("", map[string]string{"user-agent": ua, "sec-ch-ua-mobile": "?1"})

	if r.Browser.Name != "Chrome" || r.Browser.Version != "110.0.0.0" {
		t.Errorf("browser = %+v", r.Browser)
	}
	if r.CPU.Architecture != "amd64" {
		t.Errorf("cpu = %q, want amd64", r.CPU.Architecture)
	}
	if r.OS.Name != "Linux" {
		t.Errorf("os = %+v", r.OS)
	}
	if r.Device.Type != "mobile" {
		t.Errorf("device type = %q, want mobile from sec-ch-ua-mobile", r.Device.Type)
	}
}

// TestClientHintsLongHeader 校验超长头被截断，避免正则灾难性回溯。
func TestClientHintsLongHeader(t *testing.T) {
	ch := NewClientHints(map[string]string{
		"sec-ch-ua-model": `"` + strings.Repeat("A ", 25000) + `"`,
	})
	if len(ch.Model) > UAMaxLength {
		t.Fatalf("model length = %d, want <= %d", len(ch.Model), UAMaxLength)
	}
}

// TestClientHintsAppleSilicon 校验 macOS + arm 客户端提示。
func TestClientHintsAppleSilicon(t *testing.T) {
	r := ParseWithClientHints("", map[string]string{
		"sec-ch-ua-arch":     "arm",
		"sec-ch-ua-platform": `"macOS"`,
		"user-agent":         "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:97.0) Gecko/20100101 Firefox/97.0",
	})
	if r.CPU.Architecture != "arm" {
		t.Errorf("cpu = %q, want arm", r.CPU.Architecture)
	}
	if r.OS.Name != "macOS" {
		t.Errorf("os = %+v", r.OS)
	}
}

// TestClientHintsDeviceFromModel 校验仅凭 model 推断厂商与类型。
func TestClientHintsDeviceFromModel(t *testing.T) {
	cases := []struct{ model, vendor, typ string }{
		{"Pixel 99", "Google", "mobile"},
		{"SM-S918B", "Samsung", "mobile"},
	}
	for _, tc := range cases {
		d := ParseWithClientHints("", map[string]string{"sec-ch-ua-model": `"` + tc.model + `"`}).Device
		if d.Vendor != tc.vendor || d.Type != tc.typ {
			t.Errorf("model %s: device = %+v, want vendor=%s type=%s", tc.model, d, tc.vendor, tc.typ)
		}
	}
}
