package uaparser

import "testing"

// TestIsFrozenUA 对应上游 helpers 测试。
func TestIsFrozenUA(t *testing.T) {
	regular := "Mozilla/5.0 (Linux; Android 9; SM-A205U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.1234.56 Mobile Safari/537.36"
	frozen := "Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.0.0 Mobile Safari/537.36"
	if IsFrozenUA(regular) {
		t.Errorf("regular UA misdetected as frozen")
	}
	if !IsFrozenUA(frozen) {
		t.Errorf("frozen UA not detected")
	}
}

// TestGetOutlookEdition 对应上游版本映射测试。
func TestGetOutlookEdition(t *testing.T) {
	cases := []struct{ name, version, want string }{
		{"Microsoft Outlook", "16.0.4266.1001", "Outlook 2016 (MSI / Volume License)"},
		{"Microsoft Outlook", "16.0.14326.20000", "Outlook 365 / 2019+ (Modern)"},
		{"Microsoft Outlook", "15.0.4569.1506", "Outlook 2013"},
		{"MacOutlook", "16.61", "Outlook for Mac (Modern)"},
		{"MacOutlook", "15.4", "Outlook for Mac (Legacy)"},
		{"Thunderbird", "91.0", "Thunderbird"},
		{"Microsoft Outlook", "SomeRandomString", "Microsoft Outlook"},
	}
	for _, tc := range cases {
		if got := GetOutlookEdition(tc.name, tc.version); got != tc.want {
			t.Errorf("GetOutlookEdition(%q, %q) = %q, want %q", tc.name, tc.version, got, tc.want)
		}
	}
}

// TestGetDeviceVendor 对应上游型号推断测试。
func TestGetDeviceVendor(t *testing.T) {
	cases := map[string]string{
		"SM-A605G":     "Samsung",
		"Redmi Note 8": "Xiaomi",
		"Nexus 6P":     "Huawei",
		"AQUOS-TVX19B": "Sharp",
	}
	for model, want := range cases {
		if got := GetDeviceVendor(model); got != want {
			t.Errorf("GetDeviceVendor(%q) = %q, want %q", model, got, want)
		}
	}
}

// TestIsAppleSilicon 校验 UA 与 Client Hints 两种判断路径。
func TestIsAppleSilicon(t *testing.T) {
	macIntel := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:97.0) Gecko/20100101 Firefox/97.0"
	macARM := "Mozilla/5.0 (Macintosh; ARM; Mac OS X 10.15; rv:97.0) Gecko/20100101 Firefox/97.0"
	if IsAppleSiliconUA(macIntel) {
		t.Errorf("Intel Mac misdetected as Apple Silicon")
	}
	if !IsAppleSiliconUA(macARM) {
		t.Errorf("ARM Mac not detected")
	}
	withCH := ParseWithClientHints("", map[string]string{
		"sec-ch-ua-arch":     "arm",
		"sec-ch-ua-platform": `"macOS"`,
	})
	if !IsAppleSilicon(withCH) {
		t.Errorf("Client Hints arm + macOS not detected")
	}
	if IsAppleSilicon(nil) {
		t.Errorf("nil result should be false")
	}
}

// TestIsChromeFamily 校验 Blink 引擎判断。
func TestIsChromeFamily(t *testing.T) {
	edge := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36 Edg/119.0.2151.58"
	firefox := "Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/111.0"
	if !IsChromeFamily(edge) {
		t.Errorf("Edge not detected as Chrome family")
	}
	if IsChromeFamily(firefox) {
		t.Errorf("Firefox misdetected as Chrome family")
	}
}

// TestIsElectron 校验 Electron 标识。
func TestIsElectron(t *testing.T) {
	if !IsElectron("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) app/1.0.0 Chrome/114.0.0.0 Safari/537.36 Electron/25.0.0") {
		t.Errorf("Electron UA not detected")
	}
	if IsElectron("Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/111.0") {
		t.Errorf("Firefox UA misdetected as Electron")
	}
}

// TestIsBot 对应上游 bot 判断测试。
func TestIsBot(t *testing.T) {
	bots := []string{
		"Mozilla/5.0 (compatible; AhrefsBot/7.0; +http://ahrefs.com/robot/)",
		"Scrapy/1.5.0 (+https://scrapy.org)",
	}
	for _, ua := range bots {
		if !IsBot(ua) {
			t.Errorf("bot not detected: %s", ua)
		}
	}
	nonBots := []string{
		"Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/111.0",
		"Mozilla/5.0 (X11; Linux x86_64; rv:78.0) Gecko/20100101 Thunderbird/78.13.0",
	}
	for _, ua := range nonBots {
		if IsBot(ua) {
			t.Errorf("non-bot misdetected: %s", ua)
		}
	}
}

// TestIsAICrawler 对应上游 AI 爬虫测试。
func TestIsAICrawler(t *testing.T) {
	crawlers := []string{
		"Mozilla/5.0 AppleWebKit/537.36 (KHTML, like Gecko; compatible; ClaudeBot/1.0; +claudebot@anthropic.com)",
		"Mozilla/5.0 AppleWebKit/537.36 (KHTML, like Gecko); compatible; OAI-SearchBot/1.0; +https://openai.com/searchbot",
		"Mozilla/5.0 (compatible; SemrushBot-OCOB/1; +https://www.semrush.com/bot/)",
	}
	for _, ua := range crawlers {
		if !IsAICrawler(ua) {
			t.Errorf("AI crawler not detected: %s", ua)
		}
	}
	if IsAICrawler("Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/111.0") {
		t.Errorf("Firefox misdetected as AI crawler")
	}
	if !IsAIBot(crawlers[0]) {
		t.Errorf("IsAIBot alias does not match IsAICrawler")
	}
}

// TestIsAIAssistant 校验 AI 助手类抓取器。
func TestIsAIAssistant(t *testing.T) {
	ua := "Mozilla/5.0 AppleWebKit/537.36 (KHTML, like Gecko; compatible; ChatGPT-User/1.0; +https://openai.com/bot) Chrome/119.0.6045.214 Safari/537.36"
	if !IsAIAssistant(ua) {
		t.Errorf("ChatGPT-User not detected as AI assistant")
	}
	if IsAIAssistant("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)") {
		t.Errorf("Googlebot misdetected as AI assistant")
	}
}
