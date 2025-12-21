package uaparser

import "regexp"

// Common browser constants
const (
	PrefixMobile  = "Mobile "
	SuffixBrowser = " Browser"

	// Vendor names
	VendorAmazon    = "Amazon"
	VendorApple     = "Apple"
	VendorAsus      = "ASUS"
	VendorBlackBerry = "BlackBerry"
	VendorGoogle    = "Google"
	VendorHuawei    = "Huawei"
	VendorLenovo    = "Lenovo"
	VendorHonor     = "Honor"
	VendorLG        = "LG"
	VendorMicrosoft = "Microsoft"
	VendorMotorola  = "Motorola"
	VendorNvidia    = "Nvidia"
	VendorOnePlus   = "OnePlus"
	VendorOppo      = "OPPO"
	VendorSamsung   = "Samsung"
	VendorSharp     = "Sharp"
	VendorSony      = "Sony"
	VendorXiaomi    = "Xiaomi"
	VendorZebra     = "Zebra"

	// Browser names
	BrowserChrome     = "Chrome"
	BrowserChromium   = "Chromium"
	BrowserChromecast = "Chromecast"
	BrowserEdge       = "Edge"
	BrowserFirefox    = "Firefox"
	BrowserOpera      = "Opera"
	BrowserFacebook   = "Facebook"
	BrowserSogou      = "Sogou"
	BrowserWindows    = "Windows"
)

var browserRegexes = []RegexRule{
	// Chrome for Android/iOS
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)\b(?:crmo|crios)\/([\w\.]+)`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: PrefixMobile + BrowserChrome},
		},
	},
	// Edge WebView
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)webview.+edge\/([\w\.]+)`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: BrowserEdge + " WebView"},
		},
	},
	// Edge
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)edg(?:e|ios|a)?\/([\w\.]+)`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: BrowserEdge},
		},
	},
	// Opera Mini
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)(opera mini)\/([-\w\.]+)`),
			regexp.MustCompile(`(?i)(opera [mobiletab]{3,6})\b.+version\/([-\w\.]+)`),
			regexp.MustCompile(`(?i)(opera)(?:.+version\/|[\/ ]+)([\w\.]+)`),
		},
		Props: []interface{}{PropName, PropVersion},
	},
	// Opera mini on iPhone
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)opios[\/ ]+([\w\.]+)`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: BrowserOpera + " Mini"},
		},
	},
	// Opera GX
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)\bop(?:rg)?x\/([\w\.]+)`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: BrowserOpera + " GX"},
		},
	},
	// Opera Webkit
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)\bopr\/([\w\.]+)`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: BrowserOpera},
		},
	},
	// Baidu
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)\bb[ai]*d(?:uhd|[ub]*[aekoprswx]{5,6})[\/ ]?([\w\.]+)`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: "Baidu"},
		},
	},
	// Maxthon
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)\b(?:mxbrowser|mxios|myie2)\/?([-\w\.]*)\b`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: "Maxthon"},
		},
	},
	// Kindle, Lunascape, Maxthon, Netfront, Jasmine, Blazer, Sleipnir
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)(kindle)\/([\w\.]+)`),
			regexp.MustCompile(`(?i)(lunascape|maxthon|netfront|jasmine|blazer|sleipnir)[\/ ]?([\w\.]*)`),
			regexp.MustCompile(`(?i)(avant|iemobile|slim(?:browser|boat|jet))[\/ ]?([\d\.]*)`),
			regexp.MustCompile(`(?i)(?:ms|\()(ie) ([\w\.]+)`),
		},
		Props: []interface{}{PropName, PropVersion},
	},
	// Various browsers
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)(atlas|flock|rockmelt|midori|epiphany|silk|skyfire|ovibrowser|bolt|iron|vivaldi|iridium|phantomjs|bowser|qupzilla|falkon|rekonq|puffin|brave|whale|qqbrowserlite|duckduckgo|klar|helio|dragon|otter|dooble|(?:lg |qute)browser|palemoon)\/([-\w\.]+)`),
			regexp.MustCompile(`(?i)(heytap|ovi|115|surf|qwant)browser\/([\d\.]+)`),
			regexp.MustCompile(`(?i)(qwant)(?:ios|mobile)\/([\d\.]+)`),
			regexp.MustCompile(`(?i)(ecosia|weibo)(?:__| \w+@)([\d\.]+)`),
		},
		Props: []interface{}{PropName, PropVersion},
	},
	// Quark
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)quark(?:pc)?\/([-\w\.]+)`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: "Quark"},
		},
	},
	// DuckDuckGo
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)\bddg\/([\w\.]+)`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: "DuckDuckGo"},
		},
	},
	// UCBrowser
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)(?:\buc? ?browser|(?:juc.+)ucweb)[\/ ]?([\w\.]+)`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: "UCBrowser"},
		},
	},
	// WeChat
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)microm.+\bqbcore\/([\w\.]+)`),
			regexp.MustCompile(`(?i)\bqbcore\/([\w\.]+).+microm`),
			regexp.MustCompile(`(?i)micromessenger\/([\w\.]+)`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: "WeChat"},
		},
	},
	// Konqueror
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)konqueror\/([\w\.]+)`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: "Konqueror"},
		},
	},
	// IE11
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)trident.+rv[: ]([\w\.]{1,9})\b.+like gecko`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: "IE"},
		},
	},
	// Yandex
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)ya(?:search)?browser\/([\w\.]+)`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: "Yandex"},
		},
	},
	// Smart Lenovo Browser
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)slbrowser\/([\w\.]+)`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: "Smart " + VendorLenovo + SuffixBrowser},
		},
	},
	// Avast/AVG Secure Browser
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)(avast|avg)\/([\w\.]+)`),
		},
		Props: []interface{}{
			PropMapper{Prop: PropName, Fn: func(m string) string { return m + " Secure" + SuffixBrowser }},
			PropVersion,
		},
	},
	// Firefox Focus
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)\bfocus\/([\w\.]+)`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: BrowserFirefox + " Focus"},
		},
	},
	// Opera Touch
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)\bopt\/([\w\.]+)`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: BrowserOpera + " Touch"},
		},
	},
	// Coc Coc Browser
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)coc_coc\w+\/([\w\.]+)`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: "Coc Coc"},
		},
	},
	// Dolphin
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)dolfin\/([\w\.]+)`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: "Dolphin"},
		},
	},
	// Opera Coast
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)coast\/([\w\.]+)`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: BrowserOpera + " Coast"},
		},
	},
	// MIUI Browser
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)miuibrowser\/([\w\.]+)`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: "MIUI" + SuffixBrowser},
		},
	},
	// Firefox for iOS
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)fxios\/([\w\.-]+)`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: PrefixMobile + BrowserFirefox},
		},
	},
	// 360
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)\bqihoobrowser\/?([\w\.]*)`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: "360"},
		},
	},
	// QQ
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)\b(qq)\/([\w\.]+)`),
		},
		Props: []interface{}{
			PropMapper{Prop: PropName, Fn: func(m string) string { return m + "Browser" }},
			PropVersion,
		},
	},
	// Oculus/Sailfish/HuaweiBrowser/VivoBrowser/PicoBrowser
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)(oculus|sailfish|huawei|vivo|pico)browser\/([\w\.]+)`),
		},
		Props: []interface{}{
			PropMapper{Prop: PropName, Fn: func(m string) string { return m + SuffixBrowser }},
			PropVersion,
		},
	},
	// Samsung Internet
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)samsungbrowser\/([\w\.]+)`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: VendorSamsung + " Internet"},
		},
	},
	// Sogou Explorer
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)metasr[\/ ]?([\d\.]+)`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: BrowserSogou + " Explorer"},
		},
	},
	// Sogou Mobile
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)(sogou)mo\w+\/([\d\.]+)`),
		},
		Props: []interface{}{
			PropMapper{Prop: PropName, Value: BrowserSogou + " Mobile"},
			PropVersion,
		},
	},
	// Electron, Tesla
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)(electron)\/([\w\.]+) safari`),
			regexp.MustCompile(`(?i)(tesla)(?: qtcarbrowser|\/(20\d\d\.[-\w\.]+))`),
		},
		Props: []interface{}{PropName, PropVersion},
	},
	// QQBrowser, 2345Browser
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)m?(qqbrowser|2345browser|2345chrome|2345explorer)\w*[\/ ]?v?([\w\.]+)`),
		},
		Props: []interface{}{PropName, PropVersion},
	},
	// LieBao Browser/Rekonq/Steam
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)(lbbrowser|rekonq|steam)`),
		},
		Props: []interface{}{PropName},
	},
	// Iron, 360
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)ome\/([\w\.]+) \w* ?(iron) saf`),
			regexp.MustCompile(`(?i)ome\/([\w\.]+).+qihu (360)[es]e`),
		},
		Props: []interface{}{PropVersion, PropName},
	},
	// Facebook App for iOS & Android (simplified)
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i);fbav\/([\w\.]+);`),
			regexp.MustCompile(`(?i)fban\/fbios`),
			regexp.MustCompile(`(?i)fb_iab\/fb4a`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: BrowserFacebook},
			PropMapper{Prop: PropType, Value: BrowserTypeInApp},
		},
	},
	// Various In-App browsers
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)(kakao(?:talk|story))[\/ ]([\w\.]+)`),
			regexp.MustCompile(`(?i)(naver)\(.*?(\d+\.[\w\.]+).*\)`),
			regexp.MustCompile(`(?i)(daum)apps[\/ ]([\w\.]+)`),
			regexp.MustCompile(`(?i)safari (line)\/([\w\.]+)`),
			regexp.MustCompile(`(?i)\b(line)\/([\w\.]+)\/iab`),
			regexp.MustCompile(`(?i)(alipay)client\/([\w\.]+)`),
			regexp.MustCompile(`(?i)(twitter)(?:and| f.+e\/([\w\.]+))`),
			regexp.MustCompile(`(?i)(bing)(?:web|sapphire)\/([\w\.]+)`),
			regexp.MustCompile(`(?i)(instagram|snapchat|klarna)[\/ ]([-\w\.]+)`),
		},
		Props: []interface{}{
			PropName,
			PropVersion,
			PropMapper{Prop: PropType, Value: BrowserTypeInApp},
		},
	},
	// GSA
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)\bgsa\/([\w\.]+) .*safari\/`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: "GSA"},
			PropMapper{Prop: PropType, Value: BrowserTypeInApp},
		},
	},
	// TikTok
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)musical_ly(?:.+app_?version\/|_)([\w\.]+)`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: "TikTok"},
			PropMapper{Prop: PropType, Value: BrowserTypeInApp},
		},
	},
	// LinkedIn
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)\[(linkedin)app\]`),
		},
		Props: []interface{}{
			PropName,
			PropMapper{Prop: PropType, Value: BrowserTypeInApp},
		},
	},
	// Zalo
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)(zalo(?:app)?)[\/\sa-z]*([\w\.-]+)`),
		},
		Props: []interface{}{
			PropMapper{Prop: PropName, Value: "Zalo"},
			PropVersion,
			PropMapper{Prop: PropType, Value: BrowserTypeInApp},
		},
	},
	// Chromium
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)(chromium)[\/ ]([-\w\.]+)`),
		},
		Props: []interface{}{PropName, PropVersion},
	},
	// Chrome Headless
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)headlesschrome(?:\/([\w\.]+)| )`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: BrowserChrome + " Headless"},
		},
	},
	// Edge WebView2
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)wv\).+chrome\/([\w\.]+).+edgw\/`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: BrowserEdge + " WebView2"},
		},
	},
	// Chrome WebView
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i) wv\).+(chrome)\/([\w\.]+)`),
		},
		Props: []interface{}{
			PropMapper{Prop: PropName, Value: BrowserChrome + " WebView"},
			PropVersion,
		},
	},
	// Android Browser
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)droid.+ version\/([\w\.]+)\b.+(?:mobile safari|safari)`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: "Android" + SuffixBrowser},
		},
	},
	// Chrome Mobile
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)chrome\/([\w\.]+) mobile`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: PrefixMobile + BrowserChrome},
		},
	},
	// Chrome/OmniWeb/Arora/Tizen/Nokia
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)(chrome|omniweb|arora|[tizenoka]{5} ?browser)\/v?([\w\.]+)`),
		},
		Props: []interface{}{PropName, PropVersion},
	},
	// Safari Mobile
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)version\/([\w\.\,]+) .*mobile(?:\/\w+ | ?)safari`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: PrefixMobile + "Safari"},
		},
	},
	// Safari Mobile (no version)
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)iphone .*mobile(?:\/\w+ | ?)safari`),
		},
		Props: []interface{}{
			PropMapper{Prop: PropName, Value: PrefixMobile + "Safari"},
		},
	},
	// Safari
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)version\/([\w\.\,]+) .*(safari)`),
		},
		Props: []interface{}{PropVersion, PropName},
	},
	// Safari < 3.0
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)webkit.+?(mobile ?safari|safari)(\/[\w\.]+)`),
		},
		Props: []interface{}{
			PropName,
			PropMapper{Prop: PropVersion, Value: "1"},
		},
	},
	// WebKit/KHTML
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)(webkit|khtml)\/([\w\.]+)`),
		},
		Props: []interface{}{PropName, PropVersion},
	},
	// Firefox Mobile
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)(?:mobile|tablet);.*(firefox)\/([\w\.-]+)`),
		},
		Props: []interface{}{
			PropMapper{Prop: PropName, Value: PrefixMobile + BrowserFirefox},
			PropVersion,
		},
	},
	// Netscape
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)(navigator|netscape\d?)\/([-\w\.]+)`),
		},
		Props: []interface{}{
			PropMapper{Prop: PropName, Value: "Netscape"},
			PropVersion,
		},
	},
	// Wolvic/LibreWolf
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)(wolvic|librewolf)\/([\w\.]+)`),
		},
		Props: []interface{}{PropName, PropVersion},
	},
	// Firefox Reality
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)mobile vr; rv:([\w\.]+)\).+firefox`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: BrowserFirefox + " Reality"},
		},
	},
	// Other Firefox-based
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)ekiohf.+(flow)\/([\w\.]+)`),
			regexp.MustCompile(`(?i)(swiftfox)`),
			regexp.MustCompile(`(?i)(icedragon|iceweasel|camino|chimera|fennec|maemo browser|minimo|conkeror)[\/ ]?([\w\.\+]+)`),
			regexp.MustCompile(`(?i)(seamonkey|k-meleon|icecat|iceape|firebird|phoenix|basilisk|waterfox)\/([-\w\.]+)$`),
			regexp.MustCompile(`(?i)(firefox)\/([\w\.]+)`),
			regexp.MustCompile(`(?i)(mozilla)\/([\w\.]+)`),
		},
		Props: []interface{}{PropName, PropVersion},
	},
	// Other browsers
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)(amaya|dillo|doris|icab|ladybird|lynx|mosaic|netsurf|obigo|polaris|w3m|(?:go|ice|up)[\. ]?browser)[-\/ ]?v?([\w\.]+)`),
			regexp.MustCompile(`(?i)\b(links) \(([\w\.]+)`),
		},
		Props: []interface{}{
			PropName,
			PropMapper{Prop: PropVersion, Pattern: regexp.MustCompile(`_`), Replace: "."},
		},
	},
	// Cobalt
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)(cobalt)\/([\w\.]+)`),
		},
		Props: []interface{}{
			PropName,
			PropMapper{Prop: PropVersion, Pattern: regexp.MustCompile(`[^\d\.]+.`), Replace: ""},
		},
	},
}
