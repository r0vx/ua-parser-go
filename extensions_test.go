package uaparser

import (
	"encoding/json"
	"os"
	"testing"
)

// extensionCase 对应上游 test/data/ua/extension/*.json 中的用例。
type extensionCase struct {
	Desc   string `json:"desc"`
	UA     string `json:"ua"`
	Expect struct {
		Name    string `json:"name"`
		Version string `json:"version"`
		Major   string `json:"major"`
		Type    string `json:"type"`
	} `json:"expect"`
}

// normalizeUndefined 把上游测试数据中的 "undefined" 转为 Go 的空字符串。
func normalizeUndefined(v string) string {
	if v == "undefined" {
		return ""
	}
	return v
}

// TestExtensionSets 用上游测试数据校验 CLIs / Crawlers / Fetchers 规则集。
func TestExtensionSets(t *testing.T) {
	sets := []struct {
		name string
		file string
		ext  Extension
	}{
		{"CLIs", "testdata/extensions/cli.json", CLIs},
		{"Crawlers", "testdata/extensions/crawler.json", Crawlers},
		{"Fetchers", "testdata/extensions/fetcher.json", Fetchers},
	}
	for _, set := range sets {
		t.Run(set.name, func(t *testing.T) {
			data, err := os.ReadFile(set.file)
			if err != nil {
				t.Fatalf("读取测试数据失败: %v", err)
			}
			var cases []extensionCase
			if err := json.Unmarshal(data, &cases); err != nil {
				t.Fatalf("解析测试数据失败: %v", err)
			}
			for _, tc := range cases {
				expectName := normalizeUndefined(tc.Expect.Name)
				expectVersion := normalizeUndefined(tc.Expect.Version)
				expectType := normalizeUndefined(tc.Expect.Type)
				b := New(tc.UA).UseExtension(set.ext).GetBrowser()
				if b.Name != expectName {
					t.Errorf("%s: name = %q, want %q", tc.Desc, b.Name, expectName)
				}
				if b.Version != expectVersion {
					t.Errorf("%s: version = %q, want %q", tc.Desc, b.Version, expectVersion)
				}
				if b.Type != expectType {
					t.Errorf("%s: type = %q, want %q", tc.Desc, b.Type, expectType)
				}
				if tc.Expect.Major != "" && tc.Expect.Major != "undefined" && b.Major != tc.Expect.Major {
					t.Errorf("%s: major = %q, want %q", tc.Desc, b.Major, tc.Expect.Major)
				}
			}
		})
	}
}

// TestExtraDevices 用上游测试数据校验 ExtraDevices 设备规则。
func TestExtraDevices(t *testing.T) {
	data, err := os.ReadFile("testdata/extensions/extra-devices.json")
	if err != nil {
		t.Fatalf("读取测试数据失败: %v", err)
	}
	var cases []DeviceTestCase
	if err := json.Unmarshal(data, &cases); err != nil {
		t.Fatalf("解析测试数据失败: %v", err)
	}
	for _, tc := range cases {
		expectVendor := normalizeUndefined(tc.Expect.Vendor)
		expectModel := normalizeUndefined(tc.Expect.Model)
		expectType := normalizeUndefined(tc.Expect.Type)
		d := New(tc.UA).UseExtension(ExtraDevices).GetDevice()
		if d.Vendor != expectVendor {
			t.Errorf("%s: vendor = %q, want %q", tc.Desc, d.Vendor, expectVendor)
		}
		if d.Model != expectModel {
			t.Errorf("%s: model = %q, want %q", tc.Desc, d.Model, expectModel)
		}
		if d.Type != expectType {
			t.Errorf("%s: type = %q, want %q", tc.Desc, d.Type, expectType)
		}
	}
}

// TestInApps 用上游测试数据校验 InApps 浏览器规则。
func TestInApps(t *testing.T) {
	data, err := os.ReadFile("testdata/extensions/inapp.json")
	if err != nil {
		t.Fatalf("读取测试数据失败: %v", err)
	}
	var cases []extensionCase
	if err := json.Unmarshal(data, &cases); err != nil {
		t.Fatalf("解析测试数据失败: %v", err)
	}
	for _, tc := range cases {
		expectName := normalizeUndefined(tc.Expect.Name)
		expectVersion := normalizeUndefined(tc.Expect.Version)
		expectType := normalizeUndefined(tc.Expect.Type)
		b := New(tc.UA).UseExtension(InApps).GetBrowser()
		if b.Name != expectName {
			t.Errorf("%s: name = %q, want %q", tc.Desc, b.Name, expectName)
		}
		if b.Version != expectVersion {
			t.Errorf("%s: version = %q, want %q", tc.Desc, b.Version, expectVersion)
		}
		if b.Type != expectType {
			t.Errorf("%s: type = %q, want %q", tc.Desc, b.Type, expectType)
		}
	}
}

// readExtensionCases 读取浏览器类扩展测试数据。
func readExtensionCases(t *testing.T, file string) []extensionCase {
	t.Helper()
	data, err := os.ReadFile(file)
	if err != nil {
		t.Fatalf("读取测试数据失败: %v", err)
	}
	var cases []extensionCase
	if err := json.Unmarshal(data, &cases); err != nil {
		t.Fatalf("解析测试数据失败: %v", err)
	}
	return cases
}

// assertBrowserCases 逐条校验浏览器名称、版本与类型。
func assertBrowserCases(t *testing.T, cases []extensionCase, ext Extension, checkType bool) {
	t.Helper()
	for _, tc := range cases {
		expectName := normalizeUndefined(tc.Expect.Name)
		expectVersion := normalizeUndefined(tc.Expect.Version)
		expectMajor := normalizeUndefined(tc.Expect.Major)
		expectType := normalizeUndefined(tc.Expect.Type)
		b := New(tc.UA).UseExtension(ext).GetBrowser()
		if b.Name != expectName {
			t.Errorf("%s: name = %q, want %q", tc.Desc, b.Name, expectName)
		}
		if b.Version != expectVersion {
			t.Errorf("%s: version = %q, want %q", tc.Desc, b.Version, expectVersion)
		}
		if expectMajor != "" && b.Major != expectMajor {
			t.Errorf("%s: major = %q, want %q", tc.Desc, b.Major, expectMajor)
		}
		if checkType && b.Type != expectType {
			t.Errorf("%s: type = %q, want %q", tc.Desc, b.Type, expectType)
		}
	}
}

// TestEmails 用上游测试数据校验 Emails 规则。
func TestEmails(t *testing.T) {
	assertBrowserCases(t, readExtensionCases(t, "testdata/extensions/email.json"), Emails, true)
}

// TestMediaPlayers 校验 MediaPlayers 规则。
// 上游未把该数据接入测试，期望值取自 2.0.10 实际输出快照。
func TestMediaPlayers(t *testing.T) {
	assertBrowserCases(t, readExtensionCases(t, "testdata/extensions/mediaplayer.json"), MediaPlayers, true)
}

// TestVehicles 用上游测试数据校验 Vehicles 设备规则。
func TestVehicles(t *testing.T) {
	data, err := os.ReadFile("testdata/extensions/vehicle.json")
	if err != nil {
		t.Fatalf("读取测试数据失败: %v", err)
	}
	var cases []DeviceTestCase
	if err := json.Unmarshal(data, &cases); err != nil {
		t.Fatalf("解析测试数据失败: %v", err)
	}
	for _, tc := range cases {
		expectVendor := normalizeUndefined(tc.Expect.Vendor)
		expectModel := normalizeUndefined(tc.Expect.Model)
		expectType := normalizeUndefined(tc.Expect.Type)
		d := New(tc.UA).UseExtension(Vehicles).GetDevice()
		if d.Vendor != expectVendor {
			t.Errorf("%s: vendor = %q, want %q", tc.Desc, d.Vendor, expectVendor)
		}
		if d.Model != expectModel {
			t.Errorf("%s: model = %q, want %q", tc.Desc, d.Model, expectModel)
		}
		if d.Type != expectType {
			t.Errorf("%s: type = %q, want %q", tc.Desc, d.Type, expectType)
		}
	}
}

// TestExtensionPriority 校验扩展规则优先于核心规则，后加入的扩展优先。
func TestExtensionPriority(t *testing.T) {
	ua := "MyApp/1.2.3"
	if b := New(ua).GetBrowser(); b.Name != "" {
		t.Fatalf("核心规则意外匹配: %+v", b)
	}

	rule := func(name string) Extension {
		return Extension{Browser: []RegexRule{{
			Regexes: []*Regexp2{MustCompile(`(myapp)\/([\w\.]+)`)},
			Props:   []interface{}{PropName, PropVersion},
		}}}
	}
	first := Extension{Browser: []RegexRule{{
		Regexes: []*Regexp2{MustCompile(`(myapp)\/([\w\.]+)`)},
		Props: []interface{}{
			PropMapper{Prop: PropName, Value: "First"},
			PropVersion,
		},
	}}}
	b := New(ua).UseExtension(first, rule("Second")).GetBrowser()
	if b.Name != "MyApp" {
		t.Errorf("name = %q, want last extension to win (MyApp)", b.Name)
	}

	b = New(ua).UseExtension(first).GetBrowser()
	if b.Name != "First" || b.Version != "1.2.3" {
		t.Errorf("browser = %+v, want name=First version=1.2.3", b)
	}
}

// TestExtensionFetcherOS 校验 Fetchers 的操作系统规则。
func TestExtensionFetcherOS(t *testing.T) {
	cases := map[string]string{
		"WhatsApp/2.23.20.0 A": "Android",
		"WhatsApp/2.23.20.0 I": "iOS",
	}
	for ua, want := range cases {
		if got := New(ua).UseExtension(Fetchers).GetOS().Name; got != want {
			t.Errorf("%s: os = %q, want %q", ua, got, want)
		}
	}
}
