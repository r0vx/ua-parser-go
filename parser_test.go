package uaparser

import (
	"testing"
)

func TestParseBrowser(t *testing.T) {
	tests := []struct {
		name    string
		ua      string
		browser Browser
	}{
		{
			name: "Chrome on Windows",
			ua:   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
			browser: Browser{
				Name:    "Chrome",
				Version: "120.0.0.0",
				Major:   "120",
			},
		},
		{
			name: "Chrome Mobile",
			ua:   "Mozilla/5.0 (Linux; Android 10; SM-G960U) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.181 Mobile Safari/537.36",
			browser: Browser{
				Name:    "Mobile Chrome",
				Version: "88.0.4324.181",
				Major:   "88",
			},
		},
		{
			name: "Firefox",
			ua:   "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/119.0",
			browser: Browser{
				Name:    "Firefox",
				Version: "119.0",
				Major:   "119",
			},
		},
		{
			name: "Safari",
			ua:   "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 Safari/605.1.15",
			browser: Browser{
				Name:    "Safari",
				Version: "17.0",
				Major:   "17",
			},
		},
		{
			name: "Safari Mobile",
			ua:   "Mozilla/5.0 (iPhone; CPU iPhone OS 17_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 Mobile/15E148 Safari/604.1",
			browser: Browser{
				Name:    "Mobile Safari",
				Version: "17.0",
				Major:   "17",
			},
		},
		{
			name: "Edge",
			ua:   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36 Edg/120.0.0.0",
			browser: Browser{
				Name:    "Edge",
				Version: "120.0.0.0",
				Major:   "120",
			},
		},
		{
			name: "Opera",
			ua:   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36 OPR/106.0.0.0",
			browser: Browser{
				Name:    "Opera",
				Version: "106.0.0.0",
				Major:   "106",
			},
		},
		{
			name: "WeChat",
			ua:   "Mozilla/5.0 (Linux; Android 10; SM-G960U) AppleWebKit/537.36 (KHTML, like Gecko) MicroMessenger/8.0.0 Safari/537.36",
			browser: Browser{
				Name:    "WeChat",
				Version: "8.0.0",
				Major:   "8",
			},
		},
		{
			name: "UCBrowser",
			ua:   "Mozilla/5.0 (Linux; U; Android 10; en-US) AppleWebKit/537.36 (KHTML, like Gecko) UCBrowser/12.10.0.1163 Mobile Safari/537.36",
			browser: Browser{
				Name:    "UCBrowser",
				Version: "12.10.0.1163",
				Major:   "12",
			},
		},
		{
			name: "Samsung Internet",
			ua:   "Mozilla/5.0 (Linux; Android 11; SAMSUNG SM-G991B) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/16.0 Chrome/92.0.4515.166 Mobile Safari/537.36",
			browser: Browser{
				Name:    "Samsung Internet",
				Version: "16.0",
				Major:   "16",
			},
		},
		{
			name: "Chrome iOS",
			ua:   "Mozilla/5.0 (iPhone; CPU iPhone OS 17_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/119.0.6045.109 Mobile/15E148 Safari/604.1",
			browser: Browser{
				Name:    "Mobile Chrome",
				Version: "119.0.6045.109",
				Major:   "119",
			},
		},
		{
			name: "Firefox iOS",
			ua:   "Mozilla/5.0 (iPhone; CPU iPhone OS 17_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/119.0 Mobile/15E148 Safari/605.1.15",
			browser: Browser{
				Name:    "Mobile Firefox",
				Version: "119.0",
				Major:   "119",
			},
		},
		{
			name: "Edge iOS",
			ua:   "Mozilla/5.0 (iPhone; CPU iPhone OS 16_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) EdgiOS/46.0.1 Mobile/15E148 Safari/605.1.15",
			browser: Browser{
				Name:    "Edge",
				Version: "46.0.1",
				Major:   "46",
			},
		},
		{
			name: "Brave",
			ua:   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36 Brave/120.0.0.0",
			browser: Browser{
				Name:    "Brave",
				Version: "120.0.0.0",
				Major:   "120",
			},
		},
		{
			name: "Vivaldi",
			ua:   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36 Vivaldi/6.5.3206.42",
			browser: Browser{
				Name:    "Vivaldi",
				Version: "6.5.3206.42",
				Major:   "6",
			},
		},
		{
			name: "Yandex",
			ua:   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 YaBrowser/24.1.0.0 Safari/537.36",
			browser: Browser{
				Name:    "Yandex",
				Version: "24.1.0.0",
				Major:   "24",
			},
		},
		{
			name: "DuckDuckGo",
			ua:   "Mozilla/5.0 (iPhone; CPU iPhone OS 17_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 DuckDuckGo/7 Mobile/15E148 Safari/605.1.15",
			browser: Browser{
				Name:    "DuckDuckGo",
				Version: "7",
				Major:   "7",
			},
		},
		{
			name: "Chrome Headless",
			ua:   "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) HeadlessChrome/120.0.0.0 Safari/537.36",
			browser: Browser{
				Name:    "Chrome Headless",
				Version: "120.0.0.0",
				Major:   "120",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := New(tt.ua)
			browser := parser.GetBrowser()

			if browser.Name != tt.browser.Name {
				t.Errorf("Browser.Name = %q, want %q", browser.Name, tt.browser.Name)
			}
			if browser.Version != tt.browser.Version {
				t.Errorf("Browser.Version = %q, want %q", browser.Version, tt.browser.Version)
			}
			if browser.Major != tt.browser.Major {
				t.Errorf("Browser.Major = %q, want %q", browser.Major, tt.browser.Major)
			}
		})
	}
}

func TestParseOS(t *testing.T) {
	tests := []struct {
		name string
		ua   string
		os   OS
	}{
		{
			name: "Windows 10",
			ua:   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
			os: OS{
				Name:    "Windows",
				Version: "10",
			},
		},
		{
			name: "Windows 11",
			ua:   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
			os: OS{
				Name:    "Windows",
				Version: "10",
			},
		},
		{
			name: "macOS",
			ua:   "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
			os: OS{
				Name:    "macOS",
				Version: "10.15.7",
			},
		},
		{
			name: "iOS",
			ua:   "Mozilla/5.0 (iPhone; CPU iPhone OS 17_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 Mobile/15E148 Safari/604.1",
			os: OS{
				Name:    "iOS",
				Version: "17.0",
			},
		},
		{
			name: "Android",
			ua:   "Mozilla/5.0 (Linux; Android 14; SM-S918B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.6099.43 Mobile Safari/537.36",
			os: OS{
				Name:    "Android",
				Version: "14",
			},
		},
		{
			name: "Linux",
			ua:   "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
			os: OS{
				Name:    "Linux",
			},
		},
		{
			name: "Chrome OS",
			ua:   "Mozilla/5.0 (X11; CrOS x86_64 15633.0.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
			os: OS{
				Name:    "Chrome OS",
				Version: "15633.0.0",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := New(tt.ua)
			os := parser.GetOS()

			if os.Name != tt.os.Name {
				t.Errorf("OS.Name = %q, want %q", os.Name, tt.os.Name)
			}
			if tt.os.Version != "" && os.Version != tt.os.Version {
				t.Errorf("OS.Version = %q, want %q", os.Version, tt.os.Version)
			}
		})
	}
}

func TestParseDevice(t *testing.T) {
	tests := []struct {
		name   string
		ua     string
		device Device
	}{
		{
			name: "Samsung Galaxy",
			ua:   "Mozilla/5.0 (Linux; Android 14; SM-S918B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.6099.43 Mobile Safari/537.36",
			device: Device{
				Vendor: "Samsung",
				Model:  "SM-S918B",
				Type:   "mobile",
			},
		},
		{
			name: "iPhone",
			ua:   "Mozilla/5.0 (iPhone; CPU iPhone OS 17_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 Mobile/15E148 Safari/604.1",
			device: Device{
				Vendor: "Apple",
				Model:  "iPhone",
				Type:   "mobile",
			},
		},
		{
			name: "iPad",
			ua:   "Mozilla/5.0 (iPad; CPU OS 17_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 Mobile/15E148 Safari/604.1",
			device: Device{
				Vendor: "Apple",
				Model:  "iPad",
				Type:   "tablet",
			},
		},
		{
			name: "Google Pixel",
			ua:   "Mozilla/5.0 (Linux; Android 14; Pixel 8) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.6099.43 Mobile Safari/537.36",
			device: Device{
				Vendor: "Google",
				Model:  "Pixel 8",
				Type:   "mobile",
			},
		},
		{
			name: "PlayStation",
			ua:   "Mozilla/5.0 (PlayStation 5 6.10) AppleWebKit/605.1.15 (KHTML, like Gecko)",
			device: Device{
				Vendor: "Sony",
				Model:  "PlayStation 5",
				Type:   "console",
			},
		},
		{
			name: "Xbox",
			ua:   "Mozilla/5.0 (Windows NT 10.0; Win64; x64; Xbox; Xbox One) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36 Edge/44.18363.8131",
			device: Device{
				Vendor: "Microsoft",
				Model:  "Xbox One",
				Type:   "console",
			},
		},
		{
			name: "Samsung Galaxy Tab",
			ua:   "Mozilla/5.0 (Linux; Android 13; SM-T970) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
			device: Device{
				Vendor: "Samsung",
				Model:  "SM-T970",
				Type:   "tablet",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := New(tt.ua)
			device := parser.GetDevice()

			if tt.device.Vendor != "" && device.Vendor != tt.device.Vendor {
				t.Errorf("Device.Vendor = %q, want %q", device.Vendor, tt.device.Vendor)
			}
			if tt.device.Model != "" && device.Model != tt.device.Model {
				t.Errorf("Device.Model = %q, want %q", device.Model, tt.device.Model)
			}
			if tt.device.Type != "" && device.Type != tt.device.Type {
				t.Errorf("Device.Type = %q, want %q", device.Type, tt.device.Type)
			}
		})
	}
}

func TestParseEngine(t *testing.T) {
	tests := []struct {
		name   string
		ua     string
		engine Engine
	}{
		{
			name: "Blink",
			ua:   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
			engine: Engine{
				Name:    "Blink",
				Version: "120.0.0.0",
			},
		},
		{
			name: "Gecko",
			ua:   "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:120.0) Gecko/20100101 Firefox/120.0",
			engine: Engine{
				Name:    "Gecko",
				Version: "120.0",
			},
		},
		{
			name: "WebKit",
			ua:   "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 Safari/605.1.15",
			engine: Engine{
				Name:    "WebKit",
				Version: "605.1.15",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := New(tt.ua)
			engine := parser.GetEngine()

			if engine.Name != tt.engine.Name {
				t.Errorf("Engine.Name = %q, want %q", engine.Name, tt.engine.Name)
			}
			if tt.engine.Version != "" && engine.Version != tt.engine.Version {
				t.Errorf("Engine.Version = %q, want %q", engine.Version, tt.engine.Version)
			}
		})
	}
}

func TestParseCPU(t *testing.T) {
	tests := []struct {
		name string
		ua   string
		cpu  CPU
	}{
		{
			name: "AMD64",
			ua:   "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
			cpu: CPU{
				Architecture: "amd64",
			},
		},
		{
			name: "x86_64",
			ua:   "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
			cpu: CPU{
				Architecture: "amd64",
			},
		},
		{
			name: "ARM64",
			ua:   "Mozilla/5.0 (Linux; Android 14; SM-S918B; arm64-v8a) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.6099.43 Mobile Safari/537.36",
			cpu: CPU{
				Architecture: "arm64",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := New(tt.ua)
			cpu := parser.GetCPU()

			if cpu.Architecture != tt.cpu.Architecture {
				t.Errorf("CPU.Architecture = %q, want %q", cpu.Architecture, tt.cpu.Architecture)
			}
		})
	}
}

func TestParseResult(t *testing.T) {
	ua := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"
	result := Parse(ua)

	if result.UA != ua {
		t.Errorf("Result.UA = %q, want %q", result.UA, ua)
	}
	if result.Browser.Name != "Chrome" {
		t.Errorf("Result.Browser.Name = %q, want %q", result.Browser.Name, "Chrome")
	}
	if result.OS.Name != "Windows" {
		t.Errorf("Result.OS.Name = %q, want %q", result.OS.Name, "Windows")
	}
	if result.Engine.Name != "Blink" {
		t.Errorf("Result.Engine.Name = %q, want %q", result.Engine.Name, "Blink")
	}
	if result.CPU.Architecture != "amd64" {
		t.Errorf("Result.CPU.Architecture = %q, want %q", result.CPU.Architecture, "amd64")
	}
}

func TestUAMaxLength(t *testing.T) {
	// Create a very long UA string
	longUA := ""
	for i := 0; i < 1000; i++ {
		longUA += "a"
	}
	
	parser := New(longUA)
	ua := parser.GetUA()
	
	if len(ua) != UAMaxLength {
		t.Errorf("UA length = %d, want %d", len(ua), UAMaxLength)
	}
}

func BenchmarkParse(b *testing.B) {
	ua := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"
	
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Parse(ua)
	}
}

func BenchmarkGetBrowser(b *testing.B) {
	ua := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"
	parser := New(ua)
	
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		parser.GetBrowser()
	}
}
