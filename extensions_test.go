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
