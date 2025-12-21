package uaparser

import (
	"regexp"
)

// Common browser constants
const (
	PrefixMobile  = "Mobile "
	SuffixBrowser = " Browser"

	// Vendor names
	VendorAmazon     = "Amazon"
	VendorApple      = "Apple"
	VendorAsus       = "ASUS"
	VendorBlackBerry = "BlackBerry"
	VendorGoogle     = "Google"
	VendorHuawei     = "Huawei"
	VendorLenovo     = "Lenovo"
	VendorHonor      = "Honor"
	VendorLG         = "LG"
	VendorMicrosoft  = "Microsoft"
	VendorMotorola   = "Motorola"
	VendorNvidia     = "Nvidia"
	VendorOnePlus    = "OnePlus"
	VendorOppo       = "OPPO"
	VendorSamsung    = "Samsung"
	VendorSharp      = "Sharp"
	VendorSony       = "Sony"
	VendorXiaomi     = "Xiaomi"
	VendorZebra      = "Zebra"

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
		Regexes: []*Regexp2{MustCompile(`\b(?:crmo|crios)\/([\w\.]+)`)},
		Props:   []interface{}{PropVersion, PropMapper{Prop: PropName, Value: PrefixMobile + BrowserChrome}},
	},
	// Edge WebView
	{
		Regexes: []*Regexp2{MustCompile(`webview.+edge\/([\w\.]+)`)},
		Props:   []interface{}{PropVersion, PropMapper{Prop: PropName, Value: BrowserEdge + " WebView"}},
	},
	// Edge
	{
		Regexes: []*Regexp2{MustCompile(`edg(?:e|ios|a)?\/([\w\.]+)`)},
		Props:   []interface{}{PropVersion, PropMapper{Prop: PropName, Value: BrowserEdge}},
	},
	// Opera
	{
		Regexes: []*Regexp2{
			MustCompile(`(opera mini)\/([-\w\.]+)`),
			MustCompile(`(opera [mobiletab]{3,6})\b.+version\/([-\w\.]+)`),
			MustCompile(`(opera)(?:.+version\/|[\/ ]+)([\w\.]+)`),
		},
		Props: []interface{}{PropName, PropVersion},
	},
	// Opera mini on iPhone
	{
		Regexes: []*Regexp2{MustCompile(`opios[\/ ]+([\w\.]+)`)},
		Props:   []interface{}{PropVersion, PropMapper{Prop: PropName, Value: BrowserOpera + " Mini"}},
	},
	// Opera GX
	{
		Regexes: []*Regexp2{MustCompile(`\bop(?:rg)?x\/([\w\.]+)`)},
		Props:   []interface{}{PropVersion, PropMapper{Prop: PropName, Value: BrowserOpera + " GX"}},
	},
	// Opera Webkit
	{
		Regexes: []*Regexp2{MustCompile(`\bopr\/([\w\.]+)`)},
		Props:   []interface{}{PropVersion, PropMapper{Prop: PropName, Value: BrowserOpera}},
	},
	// Baidu
	{
		Regexes: []*Regexp2{MustCompile(`\bb[ai]*d(?:uhd|[ub]*[aekoprswx]{5,6})[\/ ]?([\w\.]+)`)},
		Props:   []interface{}{PropVersion, PropMapper{Prop: PropName, Value: "Baidu"}},
	},
	// Maxthon
	{
		Regexes: []*Regexp2{MustCompile(`\b(?:mxbrowser|mxios|myie2)\/?([-\w\.]*)\b`)},
		Props:   []interface{}{PropVersion, PropMapper{Prop: PropName, Value: "Maxthon"}},
	},
	// Kindle, Lunascape, Maxthon, Netfront, etc.
	{
		Regexes: []*Regexp2{
			MustCompile(`(kindle)\/([\w\.]+)`),
			MustCompile(`(lunascape|maxthon|netfront|jasmine|blazer|sleipnir)[\/ ]?([\w\.]*)`),
			MustCompile(`(avant|iemobile|slim(?:browser|boat|jet))[\/ ]?([\d\.]*)`),
			MustCompile(`(?:ms|\()(ie) ([\w\.]+)`),
		},
		Props: []interface{}{PropName, PropVersion},
	},
	// Various browsers
	{
		Regexes: []*Regexp2{
			MustCompile(`(atlas|flock|rockmelt|midori|epiphany|silk|skyfire|ovibrowser|bolt|iron|vivaldi|iridium|phantomjs|bowser|qupzilla|falkon|rekonq|puffin|brave|whale(?!.+naver)|qqbrowserlite|duckduckgo|klar|helio|(?=comodo_)?dragon|otter|dooble|(?:lg |qute)browser|palemoon)\/([-\w\.]+)`),
			MustCompile(`(heytap|ovi|115|surf|qwant)browser\/([\d\.]+)`),
			MustCompile(`(qwant)(?:ios|mobile)\/([\d\.]+)`),
			MustCompile(`(ecosia|weibo)(?:__| \w+@)([\d\.]+)`),
		},
		Props: []interface{}{PropName, PropVersion},
	},
	// Quark
	{
		Regexes: []*Regexp2{MustCompile(`quark(?:pc)?\/([-\w\.]+)`)},
		Props:   []interface{}{PropVersion, PropMapper{Prop: PropName, Value: "Quark"}},
	},
	// DuckDuckGo
	{
		Regexes: []*Regexp2{MustCompile(`\bddg\/([\w\.]+)`)},
		Props:   []interface{}{PropVersion, PropMapper{Prop: PropName, Value: "DuckDuckGo"}},
	},
	// UCBrowser
	{
		Regexes: []*Regexp2{MustCompile(`(?:\buc? ?browser|(?:juc.+)ucweb)[\/ ]?([\w\.]+)`)},
		Props:   []interface{}{PropVersion, PropMapper{Prop: PropName, Value: "UCBrowser"}},
	},
	// WeChat
	{
		Regexes: []*Regexp2{
			MustCompile(`microm.+\bqbcore\/([\w\.]+)`),
			MustCompile(`\bqbcore\/([\w\.]+).+microm`),
			MustCompile(`micromessenger\/([\w\.]+)`),
		},
		Props: []interface{}{PropVersion, PropMapper{Prop: PropName, Value: "WeChat"}},
	},
	// Konqueror
	{
		Regexes: []*Regexp2{MustCompile(`konqueror\/([\w\.]+)`)},
		Props:   []interface{}{PropVersion, PropMapper{Prop: PropName, Value: "Konqueror"}},
	},
	// IE11
	{
		Regexes: []*Regexp2{MustCompile(`trident.+rv[: ]([\w\.]{1,9})\b.+like gecko`)},
		Props:   []interface{}{PropVersion, PropMapper{Prop: PropName, Value: "IE"}},
	},
	// Yandex
	{
		Regexes: []*Regexp2{MustCompile(`ya(?:search)?browser\/([\w\.]+)`)},
		Props:   []interface{}{PropVersion, PropMapper{Prop: PropName, Value: "Yandex"}},
	},
	// Smart Lenovo Browser
	{
		Regexes: []*Regexp2{MustCompile(`slbrowser\/([\w\.]+)`)},
		Props:   []interface{}{PropVersion, PropMapper{Prop: PropName, Value: "Smart " + VendorLenovo + SuffixBrowser}},
	},
	// Avast/AVG Secure Browser
	{
		Regexes: []*Regexp2{MustCompile(`(avast|avg)\/([\w\.]+)`)},
		Props: []interface{}{
			PropMapper{Prop: PropName, Fn: func(m string) string { return m + " Secure" + SuffixBrowser }},
			PropVersion,
		},
	},
	// Firefox Focus
	{
		Regexes: []*Regexp2{MustCompile(`\bfocus\/([\w\.]+)`)},
		Props:   []interface{}{PropVersion, PropMapper{Prop: PropName, Value: BrowserFirefox + " Focus"}},
	},
	// Opera Touch
	{
		Regexes: []*Regexp2{MustCompile(`\bopt\/([\w\.]+)`)},
		Props:   []interface{}{PropVersion, PropMapper{Prop: PropName, Value: BrowserOpera + " Touch"}},
	},
	// Coc Coc Browser
	{
		Regexes: []*Regexp2{MustCompile(`coc_coc\w+\/([\w\.]+)`)},
		Props:   []interface{}{PropVersion, PropMapper{Prop: PropName, Value: "Coc Coc"}},
	},
	// Dolphin
	{
		Regexes: []*Regexp2{MustCompile(`dolfin\/([\w\.]+)`)},
		Props:   []interface{}{PropVersion, PropMapper{Prop: PropName, Value: "Dolphin"}},
	},
	// Opera Coast
	{
		Regexes: []*Regexp2{MustCompile(`coast\/([\w\.]+)`)},
		Props:   []interface{}{PropVersion, PropMapper{Prop: PropName, Value: BrowserOpera + " Coast"}},
	},
	// MIUI Browser
	{
		Regexes: []*Regexp2{MustCompile(`miuibrowser\/([\w\.]+)`)},
		Props:   []interface{}{PropVersion, PropMapper{Prop: PropName, Value: "MIUI" + SuffixBrowser}},
	},
	// Firefox for iOS
	{
		Regexes: []*Regexp2{MustCompile(`fxios\/([\w\.-]+)`)},
		Props:   []interface{}{PropVersion, PropMapper{Prop: PropName, Value: PrefixMobile + BrowserFirefox}},
	},
	// 360
	{
		Regexes: []*Regexp2{MustCompile(`\bqihoobrowser\/?([\w\.]*)`)},
		Props:   []interface{}{PropVersion, PropMapper{Prop: PropName, Value: "360"}},
	},
	// QQ
	{
		Regexes: []*Regexp2{MustCompile(`\b(qq)\/([\w\.]+)`)},
		Props: []interface{}{
			PropMapper{Prop: PropName, Fn: func(m string) string { return m + "Browser" }},
			PropVersion,
		},
	},
	// Oculus/Sailfish/HuaweiBrowser/VivoBrowser/PicoBrowser
	{
		Regexes: []*Regexp2{MustCompile(`(oculus|sailfish|huawei|vivo|pico)browser\/([\w\.]+)`)},
		Props: []interface{}{
			PropMapper{Prop: PropName, Fn: func(m string) string { return m + SuffixBrowser }},
			PropVersion,
		},
	},
	// Samsung Internet
	{
		Regexes: []*Regexp2{MustCompile(`samsungbrowser\/([\w\.]+)`)},
		Props:   []interface{}{PropVersion, PropMapper{Prop: PropName, Value: VendorSamsung + " Internet"}},
	},
	// Sogou Explorer
	{
		Regexes: []*Regexp2{MustCompile(`metasr[\/ ]?([\d\.]+)`)},
		Props:   []interface{}{PropVersion, PropMapper{Prop: PropName, Value: BrowserSogou + " Explorer"}},
	},
	// Sogou Mobile
	{
		Regexes: []*Regexp2{MustCompile(`(sogou)mo\w+\/([\d\.]+)`)},
		Props:   []interface{}{PropMapper{Prop: PropName, Value: BrowserSogou + " Mobile"}, PropVersion},
	},
	// Electron, Tesla, QQBrowser, 2345
	{
		Regexes: []*Regexp2{
			MustCompile(`(electron)\/([\w\.]+) safari`),
			MustCompile(`(tesla)(?: qtcarbrowser|\/(20\d\d\.[-\w\.]+))`),
			MustCompile(`m?(qqbrowser|2345(?=browser|chrome|explorer))\w*[\/ ]?v?([\w\.]+)`),
		},
		Props: []interface{}{PropName, PropVersion},
	},
	// LieBao Browser/Rekonq/Steam
	{
		Regexes: []*Regexp2{MustCompile(`(lbbrowser|rekonq|steam(?= (clie|tenf|gameo)))`)},
		Props:   []interface{}{PropName},
	},
	// Iron, 360
	{
		Regexes: []*Regexp2{
			MustCompile(`ome\/([\w\.]+) \w* ?(iron) saf`),
			MustCompile(`ome\/([\w\.]+).+qihu (360)[es]e`),
		},
		Props: []interface{}{PropVersion, PropName},
	},
	// Facebook App
	{
		Regexes: []*Regexp2{MustCompile(`((?:fban\/fbios|fb_iab\/fb4a)(?!.+fbav)|;fbav\/([\w\.]+);)`)},
		Props: []interface{}{
			PropMapper{Prop: PropName, Value: BrowserFacebook},
			PropVersion,
			PropMapper{Prop: PropType, Value: BrowserTypeInApp},
		},
	},
	// Various In-App browsers
	{
		Regexes: []*Regexp2{
			MustCompile(`(kakao(?:talk|story))[\/ ]([\w\.]+)`),
			MustCompile(`(naver)\(.*?(\d+\.[\w\.]+).*\)`),
			MustCompile(`(daum)apps[\/ ]([\w\.]+)`),
			MustCompile(`safari (line)\/([\w\.]+)`),
			MustCompile(`\b(line)\/([\w\.]+)\/iab`),
			MustCompile(`(alipay)client\/([\w\.]+)`),
			MustCompile(`(twitter)(?:and| f.+e\/([\w\.]+))`),
			MustCompile(`(bing)(?:web|sapphire)\/([\w\.]+)`),
			MustCompile(`(instagram|snapchat|klarna)[\/ ]([-\w\.]+)`),
		},
		Props: []interface{}{PropName, PropVersion, PropMapper{Prop: PropType, Value: BrowserTypeInApp}},
	},
	// GSA
	{
		Regexes: []*Regexp2{MustCompile(`\bgsa\/([\w\.]+) .*safari\/`)},
		Props:   []interface{}{PropVersion, PropMapper{Prop: PropName, Value: "GSA"}, PropMapper{Prop: PropType, Value: BrowserTypeInApp}},
	},
	// TikTok
	{
		Regexes: []*Regexp2{MustCompile(`musical_ly(?:.+app_?version\/|_)([\w\.]+)`)},
		Props:   []interface{}{PropVersion, PropMapper{Prop: PropName, Value: "TikTok"}, PropMapper{Prop: PropType, Value: BrowserTypeInApp}},
	},
	// LinkedIn
	{
		Regexes: []*Regexp2{MustCompile(`\[(linkedin)app\]`)},
		Props:   []interface{}{PropName, PropMapper{Prop: PropType, Value: BrowserTypeInApp}},
	},
	// Zalo
	{
		Regexes: []*Regexp2{MustCompile(`(zalo(?:app)?)[\/\sa-z]*([\w\.-]+)`)},
		Props:   []interface{}{PropMapper{Prop: PropName, Value: "Zalo"}, PropVersion, PropMapper{Prop: PropType, Value: BrowserTypeInApp}},
	},
	// Chromium
	{
		Regexes: []*Regexp2{MustCompile(`(chromium)[\/ ]([-\w\.]+)`)},
		Props:   []interface{}{PropName, PropVersion},
	},
	// Chrome Headless
	{
		Regexes: []*Regexp2{MustCompile(`headlesschrome(?:\/([\w\.]+)| )`)},
		Props:   []interface{}{PropVersion, PropMapper{Prop: PropName, Value: BrowserChrome + " Headless"}},
	},
	// Edge WebView2
	{
		Regexes: []*Regexp2{MustCompile(`wv\).+chrome\/([\w\.]+).+edgw\/`)},
		Props:   []interface{}{PropVersion, PropMapper{Prop: PropName, Value: BrowserEdge + " WebView2"}},
	},
	// Chrome WebView
	{
		Regexes: []*Regexp2{MustCompile(` wv\).+(chrome)\/([\w\.]+)`)},
		Props:   []interface{}{PropMapper{Prop: PropName, Value: BrowserChrome + " WebView"}, PropVersion},
	},
	// Android Browser
	{
		Regexes: []*Regexp2{MustCompile(`droid.+ version\/([\w\.]+)\b.+(?:mobile safari|safari)`)},
		Props:   []interface{}{PropVersion, PropMapper{Prop: PropName, Value: "Android" + SuffixBrowser}},
	},
	// Chrome Mobile
	{
		Regexes: []*Regexp2{MustCompile(`chrome\/([\w\.]+) mobile`)},
		Props:   []interface{}{PropVersion, PropMapper{Prop: PropName, Value: PrefixMobile + BrowserChrome}},
	},
	// Chrome/OmniWeb/Arora/Tizen/Nokia
	{
		Regexes: []*Regexp2{MustCompile(`(chrome|omniweb|arora|[tizenoka]{5} ?browser)\/v?([\w\.]+)`)},
		Props:   []interface{}{PropName, PropVersion},
	},
	// Safari Mobile
	{
		Regexes: []*Regexp2{MustCompile(`version\/([\w\.\,]+) .*mobile(?:\/\w+ | ?)safari`)},
		Props:   []interface{}{PropVersion, PropMapper{Prop: PropName, Value: PrefixMobile + "Safari"}},
	},
	// Safari Mobile (no version)
	{
		Regexes: []*Regexp2{MustCompile(`iphone .*mobile(?:\/\w+ | ?)safari`)},
		Props:   []interface{}{PropMapper{Prop: PropName, Value: PrefixMobile + "Safari"}},
	},
	// Safari
	{
		Regexes: []*Regexp2{MustCompile(`version\/([\w\.\,]+) .*(safari)`)},
		Props:   []interface{}{PropVersion, PropName},
	},
	// Safari < 3.0
	{
		Regexes: []*Regexp2{MustCompile(`webkit.+?(mobile ?safari|safari)(\/[\w\.]+)`)},
		Props:   []interface{}{PropName, PropMapper{Prop: PropVersion, Value: "1"}},
	},
	// WebKit/KHTML
	{
		Regexes: []*Regexp2{MustCompile(`(webkit|khtml)\/([\w\.]+)`)},
		Props:   []interface{}{PropName, PropVersion},
	},
	// Firefox Mobile
	{
		Regexes: []*Regexp2{MustCompile(`(?:mobile|tablet);.*(firefox)\/([\w\.-]+)`)},
		Props:   []interface{}{PropMapper{Prop: PropName, Value: PrefixMobile + BrowserFirefox}, PropVersion},
	},
	// Netscape
	{
		Regexes: []*Regexp2{MustCompile(`(navigator|netscape\d?)\/([-\w\.]+)`)},
		Props:   []interface{}{PropMapper{Prop: PropName, Value: "Netscape"}, PropVersion},
	},
	// Wolvic/LibreWolf
	{
		Regexes: []*Regexp2{MustCompile(`(wolvic|librewolf)\/([\w\.]+)`)},
		Props:   []interface{}{PropName, PropVersion},
	},
	// Firefox Reality
	{
		Regexes: []*Regexp2{MustCompile(`mobile vr; rv:([\w\.]+)\).+firefox`)},
		Props:   []interface{}{PropVersion, PropMapper{Prop: PropName, Value: BrowserFirefox + " Reality"}},
	},
	// Other Firefox-based
	{
		Regexes: []*Regexp2{
			MustCompile(`ekiohf.+(flow)\/([\w\.]+)`),
			MustCompile(`(swiftfox)`),
			MustCompile(`(icedragon|iceweasel|camino|chimera|fennec|maemo browser|minimo|conkeror)[\/ ]?([\w\.\+]+)`),
			MustCompile(`(seamonkey|k-meleon|icecat|iceape|firebird|phoenix|basilisk|waterfox)\/([-\w\.]+)$`),
			MustCompile(`(firefox)\/([\w\.]+)`),
			MustCompile(`(mozilla)\/([\w\.]+(?= .+rv\:.+gecko\/\d+)|[0-4][\w\.]+(?!.+compatible))`),
		},
		Props: []interface{}{PropName, PropVersion},
	},
	// Other browsers
	{
		Regexes: []*Regexp2{
			MustCompile(`(amaya|dillo|doris|icab|ladybird|lynx|mosaic|netsurf|obigo|polaris|w3m|(?:go|ice|up)[\. ]?browser)[-\/ ]?v?([\w\.]+)`),
			MustCompile(`\b(links) \(([\w\.]+)`),
		},
		Props: []interface{}{PropName, PropMapper{Prop: PropVersion, Pattern: regexp.MustCompile(`_`), Replace: "."}},
	},
	// Cobalt
	{
		Regexes: []*Regexp2{MustCompile(`(cobalt)\/([\w\.]+)`)},
		Props:   []interface{}{PropName, PropMapper{Prop: PropVersion, Pattern: regexp.MustCompile(`[^\d\.]+.`), Replace: ""}},
	},
}
