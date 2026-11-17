package uaparser

import (
	"regexp"
	"strconv"
	"strings"
)

var (
	frozenUARe      = regexp.MustCompile(`^Mozilla\/5\.0 \((Windows NT 10\.0; Win64; x64|Macintosh; Intel Mac OS X 10_15_7|X11; Linux x86_64|X11; CrOS x86_64 14541\.0\.0|Fuchsia|Linux; Android 10; K)\) AppleWebKit\/537\.36 \(KHTML, like Gecko\) Chrome\/\d+\.0\.0\.0 (Mobile )?Safari\/537\.36`)
	electronRe      = regexp.MustCompile(`(?i) electron\/`)
	microsoftNameRe = regexp.MustCompile(`(?i)microsoft\s+`)
)

// botTypes 是 isBot 认可的浏览器类型。
var botTypes = map[string]struct{}{
	BrowserTypeCLI:     {},
	BrowserTypeCrawler: {},
	BrowserTypeFetcher: {},
	BrowserTypeLibrary: {},
}

// aiCrawlers 是上游 bot-detection 认定的 AI 爬虫名称（小写）。
var aiCrawlers = toLowerSet(
	"AI2Bot", "Amazonbot", "anthropic-ai", "ClaudeBot", "Claude-SearchBot", "Claude-Web",
	"Applebot", "Applebot-Extended", "atlassian-bot", "Bravebot", "Bytespider", "TikTokSpider",
	"Cloudflare-AutoRAG", "cohere-training-data-crawler", "CCBot", "Coveobot", "DataForSeoBot",
	"DeepSeekBot", "Diffbot", "Google-Extended", "Google-NotebookLM", "GoogleOther",
	"GoogleOther-Image", "GoogleOther-Video", "Google-CloudVertexBot", "ImagesiftBot", "PetalBot",
	"PanguBot", "HuggingFace-Bot", "Kangaroo Bot", "FirecrawlAgent", "FacebookBot",
	"meta-externalagent", "meta-webindexer", "GPTBot", "OAI-SearchBot", "PerplexityBot",
	"Replicate-Bot", "RunPod-Bot", "SBIntuitionsBot", "SemrushBot-OCOB", "SemrushBot-SWA",
	"Timpibot", "Together-Bot", "VelenPublicWebCrawler", "v0bot", "omgili", "omgilibot",
	"Webzio-Extended", "xAI-Bot", "YouBot", "ChatGLM-Spider",
)

// aiAssistants 是上游 bot-detection 认定的 AI 助手名称（小写）。
var aiAssistants = toLowerSet(
	"NovaAct", "Claude-User", "Cohere-AI", "DuckAssistBot", "Gemini-Deep-Research",
	"MistralAI-User", "ChatGPT-User", "Perplexity-User",
)

// toLowerSet 把名称列表转成小写集合，便于大小写不敏感查找。
func toLowerSet(names ...string) map[string]struct{} {
	set := make(map[string]struct{}, len(names))
	for _, n := range names {
		set[strings.ToLower(n)] = struct{}{}
	}
	return set
}

// IsBot 判断 UA 是否属于机器人（CLI / 爬虫 / 抓取器 / 库）。
func IsBot(ua string) bool {
	b := New(ua).UseExtension(Bots).GetBrowser()
	_, ok := botTypes[strings.ToLower(b.Type)]
	return ok
}

// IsAICrawler 判断 UA 是否为 AI 训练/检索爬虫。
func IsAICrawler(ua string) bool {
	b := New(ua).UseExtension(Crawlers).GetBrowser()
	_, ok := aiCrawlers[strings.ToLower(b.Name)]
	return ok
}

// IsAIBot 是 IsAICrawler 的兼容别名。
func IsAIBot(ua string) bool {
	return IsAICrawler(ua)
}

// IsAIAssistant 判断 UA 是否为 AI 助手类抓取器。
func IsAIAssistant(ua string) bool {
	b := New(ua).UseExtension(Fetchers).GetBrowser()
	_, ok := aiAssistants[strings.ToLower(b.Name)]
	return ok
}

// IsChromeFamily 判断 UA 是否使用 Blink 引擎（Chrome 系）。
func IsChromeFamily(ua string) bool {
	return New(ua).GetEngine().Name == EngineName.Blink
}

// IsElectron 判断 UA 是否包含 Electron 标识。
func IsElectron(ua string) bool {
	return electronRe.MatchString(ua)
}

// IsFrozenUA 判断 UA 是否为浏览器冻结/简化后的 UA。
func IsFrozenUA(ua string) bool {
	return frozenUARe.MatchString(ua)
}

// GetDeviceVendor 根据设备型号猜测厂商；无法识别时返回空字符串。
func GetDeviceVendor(model string) string {
	const uaPrefix = "Mozilla/5.0 (Linux; Android 10; "
	const uaSuffix = ") AppleWebKit/537.36 (KHTML, like Gecko) Chrome/100.0.0.0 Safari/537.36"
	return New(uaPrefix + model + uaSuffix).GetDevice().Vendor
}

// IsAppleSilicon 根据解析结果判断 Apple Silicon。
// 建议用 ParseWithHTTPHeaders 取得带 Client Hints 的结果，UA 本身无法区分。
func IsAppleSilicon(r *Result) bool {
	if r == nil {
		return false
	}
	return r.OS.Name == OSName.MacOS && r.CPU.Architecture == CPUArch.ARM
}

// IsAppleSiliconUA 从 UA 字符串判断 Apple Silicon；无 Client Hints 时准确率有限。
func IsAppleSiliconUA(ua string) bool {
	return IsAppleSilicon(Parse(ua))
}

// GetOutlookEdition 把 Outlook 名称/版本映射为可读版本（如 "Outlook 2019 (Modern Word)"）。
func GetOutlookEdition(name, version string) string {
	if name == "" || version == "" {
		return name
	}
	cleanName := microsoftNameRe.ReplaceAllString(strings.ToLower(name), "")

	if cleanName == "macoutlook" {
		major := leadingInt(version)
		if major >= 16 {
			return "Outlook for Mac (Modern)"
		}
		return "Outlook for Mac (Legacy)"
	}

	if cleanName == "outlook" {
		parts := strings.Split(version, ".")
		major, err := strconv.Atoi(parts[0])
		if err != nil {
			return name
		}
		build := 0
		if len(parts) > 2 {
			build, _ = strconv.Atoi(parts[2])
		}

		switch {
		case major == 15:
			return "Outlook 2013"
		case major == 14:
			return "Outlook 2010"
		case major == 12:
			return "Outlook 2007"
		case major < 12:
			return "Outlook (Legacy)"
		case major == 16:
			if build < 10000 {
				return "Outlook 2016 (MSI / Volume License)"
			}
			return "Outlook 365 / 2019+ (Modern)"
		}
	}

	return name
}

// leadingInt 解析版本号首段的整数，失败时返回 0。
func leadingInt(version string) int {
	major, err := strconv.Atoi(strings.Split(version, ".")[0])
	if err != nil {
		return 0
	}
	return major
}
