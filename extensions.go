package uaparser

import "regexp"

// Extension 保存按检测项分组的扩展规则；扩展规则优先于核心规则执行。
type Extension struct {
	Browser []RegexRule
	CPU     []RegexRule
	Device  []RegexRule
	Engine  []RegexRule
	OS      []RegexRule
}

// CLIs 识别命令行工具（wget/curl/Lynx/ELinks/HTTPie/PowerShell）。
var CLIs = Extension{
	Browser: []RegexRule{
		{
			Regexes: []*Regexp2{MustCompile(`(wget|curl|lynx|elinks|httpie|powershell)[\/ ]\(?([\w\.-]+)`)},
			Props: []interface{}{
				PropName,
				PropVersion,
				PropMapper{Prop: PropType, Value: BrowserTypeCLI},
			},
		},
	},
}

// Crawlers 识别搜索引擎与 AI 爬虫。
var Crawlers = Extension{
	Browser: []RegexRule{
		{
			Regexes: []*Regexp2{
				MustCompile(`((?:adidx|ahrefs|amazon|(?:amzn|oai)-search|awario(?:smart|rss)?|bing|brave|cc|contx|coveo|criteo|dot|duckduck(?:go-favicons-)?|exa|facebook|gpt|iask|kagi|kangaroo |linkedin|mj12|mojeek|onespot-scraper|perplexity|sbintuitions|semrush|seznam|surdotly|swift|yep)bot)\/([\w\.-]+)`),
				MustCompile(`(algolia crawler(?: renderscript)?)\/?([\w\.]*)`),
				MustCompile(`(applebot(?:-extended)?)\/?([\w\.]*)`),
				MustCompile(`(baiduspider[-imagevdonwsfcpr]{0,7})\/?([\w\.]*)`),
				MustCompile(`(claude(?:bot|-searchbot|-web)|anthropic-ai)\/?([\w\.]*)`),
				MustCompile(`(coccocbot-(?:image|web))\/([\w\.]+)`),
				MustCompile(`(daum(?:oa)?(?:-image)?|hubspot crawler)[ \/]([\w\.]+)`),
				MustCompile(`(facebook(?:externalhit|catalog)|meta-(?:externalagent|externalads|webindexer))\/([\w\.]+)`),
				MustCompile(`(google(?:bot|other|-inspectiontool)(?:-image|-video|-news)?|storebot-google)\/?([\w\.]*)`),
				MustCompile(`(ia_archiver|archive\.org_bot)\/?([\w\.]*)`),
				MustCompile(`(oncrawl) mobile\/([\w\.]+)`),
				MustCompile(`(qwantbot(?:-news)?)[-\w]*\/?([\w\.]*)`),
				MustCompile(`((?:semrush|splitsignal)bot[-abcfimostw]*)\/?([\w\.-]*)`),
				MustCompile(`(sogou (?:pic|head|web|orion|news) spider)\/([\w\.]+)`),
				MustCompile(`(y!?j-(?:asr|br[uw]|dscv|mmp|vsidx|wsc))\/([\w\.]+)`),
				MustCompile(`(yandex(?:(?:mobile)?(?:accessibility|additional|com|renderresources|screenshot|sprav)?bot(?!.+mirror)|image(?:s|resizer)|adnet|blogs|favicons|market|media|metrika|news|ontodb(?:api)?|partner|rca|tracker|turbo|verti(?:cal)?s|webmaster|video(?:parser)?))\/([\w\.]+)`),
				MustCompile(`(yeti)\/([\w\.]+)`),
				MustCompile(`((?:aihit|blex|diff|huggingface-|msn|pangu|replicate-|runpod-|timpi|together-|xai-|you|zum)bot|(?:audisto |brightedge |magpie-|velenpublicweb)crawler|(?:chatglm-|line|screaming frog seo |yisou)spider|cloudflare-autorag|cotoyogi|(?:firecrawl|twin)agent|freespoke|omgili(?:bot)?|openai image downloader|startpageprivateimageproxy|webzio-extended)\/?([\w\.]*)`),
			},
			Props: []interface{}{
				PropName,
				PropVersion,
				PropMapper{Prop: PropType, Value: BrowserTypeCrawler},
			},
		},
		{
			Regexes: []*Regexp2{MustCompile(`(ev-crawler)\/([\w\.]+)`)},
			Props: []interface{}{
				PropMapper{Prop: PropName, Value: "Headline"},
				PropVersion,
				PropMapper{Prop: PropType, Value: BrowserTypeCrawler},
			},
		},
		{
			Regexes: []*Regexp2{MustCompile(`(yandexbot\/([\w\.]+); mirrordetector)`)},
			Props: []interface{}{
				PropMapper{Prop: PropName, Pattern: regexp.MustCompile(`(?i)/.+;`), Replace: ""},
				PropVersion,
				PropMapper{Prop: PropType, Value: BrowserTypeCrawler},
			},
		},
		{
			Regexes: []*Regexp2{
				MustCompile(`((?:adsbot|apis|mediapartners)-google(?:-mobile)?|google-?(?:other|cloudvertexbot|extended|notebooklm|safety))`),
				MustCompile(`\b((ai2|aspiegel|atlassian-|dataforseo|deepseek|imagesift|petal|seekport|turnitin|v0|yacy)bot|360spider-?(image|video)?|baidu-ads|botify|(byte|tiktok)spider|cohere-training-data-crawler|elastic(?=\/s)|marginalia|proximic|siteimprove(?=bot|\.com)|teoma|webzio|yahoo! slurp)`),
			},
			Props: []interface{}{
				PropName,
				PropMapper{Prop: PropType, Value: BrowserTypeCrawler},
			},
		},
	},
}

// ExtraDevices 补充核心规则未覆盖的平板/手机厂商。
var ExtraDevices = Extension{
	Device: []RegexRule{
		{
			Regexes: []*Regexp2{
				MustCompile(`(nook)[\w ]+build\/(\w+)`),
				MustCompile(`(dell) (strea[kpr\d ]*[\dko])`),
				MustCompile(`(le[- ]+pan)[- ]+(\w{1,9}) bui`),
				MustCompile(`(trinity)[- ]*(t\d{3}) bui`),
				MustCompile(`(gigaset)[- ]+(q\w{1,9}) bui`),
				MustCompile(`(vodafone) ([\w ]+)(?:\)| bui)`),
			},
			Props: []interface{}{
				PropVendor,
				PropModel,
				PropMapper{Prop: PropType, Value: DeviceTablet},
			},
		},
		{
			Regexes: []*Regexp2{MustCompile(`(u304aa)`)},
			Props: []interface{}{
				PropModel,
				PropMapper{Prop: PropVendor, Value: "AT&T"},
				PropMapper{Prop: PropType, Value: DeviceMobile},
			},
		},
		{
			Regexes: []*Regexp2{MustCompile(`\bsie-(\w*)`)},
			Props: []interface{}{
				PropModel,
				PropMapper{Prop: PropVendor, Value: "Siemens"},
				PropMapper{Prop: PropType, Value: DeviceMobile},
			},
		},
		{
			Regexes: []*Regexp2{MustCompile(`\b(rct\w+) b`)},
			Props: []interface{}{
				PropModel,
				PropMapper{Prop: PropVendor, Value: "RCA"},
				PropMapper{Prop: PropType, Value: DeviceTablet},
			},
		},
		{
			Regexes: []*Regexp2{MustCompile(`\b(venue[\d ]{2,7}) b`)},
			Props: []interface{}{
				PropModel,
				PropMapper{Prop: PropVendor, Value: "Dell"},
				PropMapper{Prop: PropType, Value: DeviceTablet},
			},
		},
		{
			Regexes: []*Regexp2{MustCompile(`\b(q(?:mv|ta)\w+) b`)},
			Props: []interface{}{
				PropModel,
				PropMapper{Prop: PropVendor, Value: "Verizon"},
				PropMapper{Prop: PropType, Value: DeviceTablet},
			},
		},
		{
			Regexes: []*Regexp2{MustCompile(`\b(?:barnes[& ]+noble |bn[rt])([\w\+ ]*) b`)},
			Props: []interface{}{
				PropModel,
				PropMapper{Prop: PropVendor, Value: "Barnes & Noble"},
				PropMapper{Prop: PropType, Value: DeviceTablet},
			},
		},
		{
			Regexes: []*Regexp2{MustCompile(`\b(tm\d{3}\w+) b`)},
			Props: []interface{}{
				PropModel,
				PropMapper{Prop: PropVendor, Value: "NuVision"},
				PropMapper{Prop: PropType, Value: DeviceTablet},
			},
		},
		{
			Regexes: []*Regexp2{MustCompile(`\b(k88) b`)},
			Props: []interface{}{
				PropModel,
				PropMapper{Prop: PropVendor, Value: "ZTE"},
				PropMapper{Prop: PropType, Value: DeviceTablet},
			},
		},
		{
			Regexes: []*Regexp2{MustCompile(`\b(nx\d{3}j) b`)},
			Props: []interface{}{
				PropModel,
				PropMapper{Prop: PropVendor, Value: "ZTE"},
				PropMapper{Prop: PropType, Value: DeviceMobile},
			},
		},
		{
			Regexes: []*Regexp2{MustCompile(`\b(gen\d{3}) b.+49h`)},
			Props: []interface{}{
				PropModel,
				PropMapper{Prop: PropVendor, Value: "Swiss"},
				PropMapper{Prop: PropType, Value: DeviceMobile},
			},
		},
		{
			Regexes: []*Regexp2{MustCompile(`\b(zur\d{3}) b`)},
			Props: []interface{}{
				PropModel,
				PropMapper{Prop: PropVendor, Value: "Swiss"},
				PropMapper{Prop: PropType, Value: DeviceTablet},
			},
		},
		{
			Regexes: []*Regexp2{MustCompile(`^((zeki)?tb.*\b) b`)},
			Props: []interface{}{
				PropModel,
				PropMapper{Prop: PropVendor, Value: "Zeki"},
				PropMapper{Prop: PropType, Value: DeviceTablet},
			},
		},
		{
			Regexes: []*Regexp2{
				MustCompile(`\b([yr]\d{2}) b`),
				MustCompile(`\b(?:dragon[- ]+touch |dt)(\w{5}) b`),
			},
			Props: []interface{}{
				PropModel,
				PropMapper{Prop: PropVendor, Value: "Dragon Touch"},
				PropMapper{Prop: PropType, Value: DeviceTablet},
			},
		},
		{
			Regexes: []*Regexp2{MustCompile(`\b(ns-?\w{0,9}) b`)},
			Props: []interface{}{
				PropModel,
				PropMapper{Prop: PropVendor, Value: "Insignia"},
				PropMapper{Prop: PropType, Value: DeviceTablet},
			},
		},
		{
			Regexes: []*Regexp2{MustCompile(`\b((nxa|next)-?\w{0,9}) b`)},
			Props: []interface{}{
				PropModel,
				PropMapper{Prop: PropVendor, Value: "NextBook"},
				PropMapper{Prop: PropType, Value: DeviceTablet},
			},
		},
		{
			Regexes: []*Regexp2{MustCompile(`\b(xtreme_)?(v(1[045]|2[015]|[3469]0|7[05])) b`)},
			Props: []interface{}{
				PropMapper{Prop: PropVendor, Value: "Voice"},
				PropModel,
				PropMapper{Prop: PropType, Value: DeviceMobile},
			},
		},
		{
			Regexes: []*Regexp2{MustCompile(`\b(lvtel\-)?(v1[12]) b`)},
			Props: []interface{}{
				PropMapper{Prop: PropVendor, Value: "LvTel"},
				PropModel,
				PropMapper{Prop: PropType, Value: DeviceMobile},
			},
		},
		{
			Regexes: []*Regexp2{MustCompile(`\b(ph-1) `)},
			Props: []interface{}{
				PropModel,
				PropMapper{Prop: PropVendor, Value: "Essential"},
				PropMapper{Prop: PropType, Value: DeviceMobile},
			},
		},
		{
			Regexes: []*Regexp2{MustCompile(`\b(v(100md|700na|7011|917g).*\b) b`)},
			Props: []interface{}{
				PropModel,
				PropMapper{Prop: PropVendor, Value: "Envizen"},
				PropMapper{Prop: PropType, Value: DeviceTablet},
			},
		},
		{
			Regexes: []*Regexp2{MustCompile(`\b(trio[-\w\. ]+) b`)},
			Props: []interface{}{
				PropModel,
				PropMapper{Prop: PropVendor, Value: "MachSpeed"},
				PropMapper{Prop: PropType, Value: DeviceTablet},
			},
		},
		{
			Regexes: []*Regexp2{MustCompile(`\btu_(1491) b`)},
			Props: []interface{}{
				PropModel,
				PropMapper{Prop: PropVendor, Value: "Rotor"},
				PropMapper{Prop: PropType, Value: DeviceTablet},
			},
		},
	},
}

// InApps 识别桌面与移动端的应用内浏览器。
var InApps = Extension{
	Browser: []RegexRule{
		{
			Regexes: []*Regexp2{
				MustCompile(`\b(discord|figma|mattermost|notion|postman|rambox|rocket.chat|slack|teams)\/([\w\.]+).+(electron\/|; ios)`),
				MustCompile(`(flipboard)\/([\w\.]+)`),
			},
			Props: []interface{}{
				PropName,
				PropVersion,
				PropMapper{Prop: PropType, Value: BrowserTypeInApp},
			},
		},
		{
			Regexes: []*Regexp2{
				MustCompile(`(evernote) win`),
				MustCompile(`(teams)mobile-(ios|and)`),
			},
			Props: []interface{}{
				PropName,
				PropMapper{Prop: PropType, Value: BrowserTypeInApp},
			},
		},
		{
			Regexes: []*Regexp2{MustCompile(`chatlyio\/([\d\.]+)`)},
			Props: []interface{}{
				PropVersion,
				PropMapper{Prop: PropName, Value: "Slack"},
				PropMapper{Prop: PropType, Value: BrowserTypeInApp},
			},
		},
		{
			Regexes: []*Regexp2{MustCompile(`ultralite app_version\/([\w\.]+)`)},
			Props: []interface{}{
				PropVersion,
				PropMapper{Prop: PropName, Value: "TikTok Lite"},
				PropMapper{Prop: PropType, Value: BrowserTypeInApp},
			},
		},
		{
			Regexes: []*Regexp2{MustCompile(`\) code\/([\d\.]+).+electron\/`)},
			Props: []interface{}{
				PropVersion,
				PropMapper{Prop: PropName, Value: "VS Code"},
				PropMapper{Prop: PropType, Value: BrowserTypeInApp},
			},
		},
		{
			Regexes: []*Regexp2{MustCompile(`jp\.co\.yahoo\.(?:android\.yjtop|ipn\.appli)\/([\d\.]+)`)},
			Props: []interface{}{
				PropVersion,
				PropMapper{Prop: PropName, Value: "Yahoo! Japan"},
				PropMapper{Prop: PropType, Value: BrowserTypeInApp},
			},
		},
	},
}

// Libraries 识别 HTTP 客户端与库（Scrapy、axios、node-fetch 等）。
var Libraries = Extension{
	Browser: []RegexRule{
		{
			Regexes: []*Regexp2{
				MustCompile(`^((?:apache|go|java)-http-?client|axios|bun|dart|deno|got|(?:guzzle|lua-resty-|ocaml-co|ok)http|hackney|http\.rb|java|jetty|libwww-perl|needle|node(?:\.js|-fetch|-superagent)|php-soap|postmanruntime|python-(?:httpx|urllib[23]?|requests)|rest-client|scrapy)\/([\w\.]+)`),
				MustCompile(`(adobeair|aiohttp|jsdom)\/([\w\.]+)`),
				MustCompile(`(nutch)-([\w\.-]+)(\(|$)`),
				MustCompile(`\((java)\/([\w\.]+)`),
			},
			Props: []interface{}{
				PropName,
				PropVersion,
				PropMapper{Prop: PropType, Value: BrowserTypeLibrary},
			},
		},
		{
			Regexes: []*Regexp2{MustCompile(`(node-fetch|phpcrawl|undici)`)},
			Props: []interface{}{
				PropName,
				PropMapper{Prop: PropType, Value: BrowserTypeLibrary},
			},
		},
	},
}

// Bots 合并所有机器人相关扩展集（CLIs + Fetchers + Crawlers + Libraries）。
var Bots = Extension{
	Browser: mergeRuleLists(CLIs.Browser, Fetchers.Browser, Crawlers.Browser, Libraries.Browser),
	OS:      Fetchers.OS,
}

// mergeRuleLists 依次拼接多个规则切片。
func mergeRuleLists(lists ...[]RegexRule) []RegexRule {
	total := 0
	for _, l := range lists {
		total += len(l)
	}
	merged := make([]RegexRule, 0, total)
	for _, l := range lists {
		merged = append(merged, l...)
	}
	return merged
}

// Fetchers 识别链接预览、抓取与截图类客户端。
var Fetchers = Extension{
	Browser: []RegexRule{
		{
			Regexes: []*Regexp2{
				MustCompile(`(asana|ahrefssiteaudit|(?:bing|microsoft)preview|blueno|(?:amzn|chatgpt|claude|mistralai|perplexity)-user|cohere-ai|flipboardproxy|hubspot page fetcher|mastodon|(?:bitly|bufferlinkpreview|discord|duckassist|linkedin|pinterest|reddit|roger|siteaudit|twitter|uptime(?:ro)?|zoom)bot|google-site-verification|iframely|kakaotalk-scrap|meta-externalfetcher|y!?j-dlc|yandex(?:calendar|direct(?:dyn)?|fordomain|pagechecker|searchshop)|yadirectfetcher|whatsapp)\/([\w\.]+)`),
				MustCompile(`(bluesky) cardyb\/([\w\.]+)`),
				MustCompile(`(feedly)(?:bot)?\/([\w\.]+)`),
				MustCompile(`agent-(novaact)\/([\w\.]+)`),
				MustCompile(`(skypeuripreview) preview\/([\w\.]+)`),
				MustCompile(`(slack(?:bot)?(?:-imgproxy|-linkexpanding)?) ([\w\.]+)`),
			},
			Props: []interface{}{
				PropName,
				PropVersion,
				PropMapper{Prop: PropType, Value: BrowserTypeFetcher},
			},
		},
		{
			Regexes: []*Regexp2{
				MustCompile(`((?:better uptime |keybase|telegram|vercel)bot|lighthouse$|feedfetcher-google|gemini-deep-research|google(?:docs|imageproxy|-read-aloud|-pagerenderer|producer)|snap url preview|vercel(flags|tracing|-(favicon|screenshot)-bot)|virustotal(?=cloud)|yandex(?:sitelinks|userproxy))`),
			},
			Props: []interface{}{
				PropName,
				PropMapper{Prop: PropType, Value: BrowserTypeFetcher},
			},
		},
	},
	OS: []RegexRule{
		{
			Regexes: []*Regexp2{MustCompile(`whatsapp\/[\d\.]+ (a|i)`)},
			Props: []interface{}{
				PropMapper{Prop: PropName, Fn: func(m string) string {
					if m == "A" {
						return "Android"
					}
					return "iOS"
				}},
			},
		},
	},
}
