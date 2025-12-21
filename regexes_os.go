package uaparser

import (
	"regexp"
	"strings"
)

var osRegexes = []RegexRule{
	// Windows RT
	{
		Regexes: []*Regexp2{
			MustCompile(`(windows nt) (6\.[23]); arm`),
		},
		Props: []interface{}{
			PropMapper{Prop: PropName, Pattern: regexp.MustCompile(`(?i)NT`), Replace: "RT"},
			PropMapper{Prop: PropVersion, StrMap: windowsVersionMap},
		},
	},
	// Windows IoT/Mobile/Phone and Windows versions
	{
		Regexes: []*Regexp2{
			MustCompile(`(windows (?:phone|mobile|iot))(?: os)?[\/ ]?([\d\.]*( se)?)`),
			MustCompile(`(windows)[\/ ](1[01]|2000|3\.1|7|8(\.1)?|9[58]|me|server 20\d\d( r2)?|vista|xp)`),
		},
		Props: []interface{}{PropName, PropVersion},
	},
	// Windows NT (plain, no version number - should return "NT")
	{
		Regexes: []*Regexp2{
			MustCompile(`(windows nt)\)(?!.+xbox)`),
			MustCompile(`\b(winnt)[;\) ]`),
		},
		Props: []interface{}{
			PropMapper{Prop: PropName, Value: BrowserWindows},
			PropMapper{Prop: PropVersion, Value: "NT"},
		},
	},
	// Windows NT with version (excluding Xbox)
	{
		Regexes: []*Regexp2{
			MustCompile(`windows nt ?([\d\.]+)(?!.+xbox)`),
			MustCompile(`\bwinnt([\d\.]+)`),
			MustCompile(`\bwin(?=3| ?9)(?:nt| 9x )?([\d\.;]*)`),
		},
		Props: []interface{}{
			PropMapper{Prop: PropVersion, Pattern: regexp.MustCompile(`[;)]`), Replace: "", StrMap: windowsVersionMap},
			PropMapper{Prop: PropName, Value: BrowserWindows},
		},
	},
	// Windows CE
	{
		Regexes: []*Regexp2{
			MustCompile(`(windows ce)\/?([\d\.]*)`),
		},
		Props: []interface{}{PropName, PropVersion},
	},
	// iOS/macOS
	{
		Regexes: []*Regexp2{
			MustCompile(`[adehimnop]{4,7}\b(?:.*os ([\w]+) like mac|; opera)`),
			MustCompile(`(?:ios;fbsv|ios(?=.+ip(?:ad|hone)|.+apple ?tv)|ip(?:ad|hone)(?: |.+i(?:pad)?)os|apple ?tv.+ios)[\/ ]([\w\.]+)`),
			MustCompile(`\btvos ?([\w\.]+)`),
			MustCompile(`cfnetwork\/.+darwin`),
		},
		Props: []interface{}{
			PropMapper{Prop: PropVersion, Pattern: regexp.MustCompile(`_`), Replace: "."},
			PropMapper{Prop: PropName, Value: "iOS"},
		},
	},
	// macOS (excluding haiku and morphos)
	{
		Regexes: []*Regexp2{
			MustCompile(`(mac os x) ?([\w\. ]*)`),
			MustCompile(`(macintosh|mac_powerpc\b)(?!.+(haiku|morphos))`),
		},
		Props: []interface{}{
			PropMapper{Prop: PropName, Value: "macOS"},
			PropMapper{Prop: PropVersion, Pattern: regexp.MustCompile(`_`), Replace: "."},
		},
	},
	// Google Chromecast
	{
		Regexes: []*Regexp2{
			MustCompile(`android ([\d\.]+).*crkey`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: BrowserChromecast + " Android"},
		},
	},
	{
		Regexes: []*Regexp2{
			MustCompile(`fuchsia.*crkey\/([\d\.]+)`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: BrowserChromecast + " Fuchsia"},
		},
	},
	{
		Regexes: []*Regexp2{
			MustCompile(`crkey\/([\d\.]+).*devicetype\/smartspeaker`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: BrowserChromecast + " SmartSpeaker"},
		},
	},
	{
		Regexes: []*Regexp2{
			MustCompile(`linux.*crkey\/([\d\.]+)`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: BrowserChromecast + " Linux"},
		},
	},
	{
		Regexes: []*Regexp2{
			MustCompile(`crkey\/([\d\.]+)`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: BrowserChromecast},
		},
	},
	// Android-x86
	{
		Regexes: []*Regexp2{
			MustCompile(`droid ([\w\.]+)\b.+(android[- ]x86)`),
		},
		Props: []interface{}{PropVersion, PropName},
	},
	// Ubuntu Touch
	{
		Regexes: []*Regexp2{
			MustCompile(`(ubuntu) ([\w\.]+) like android`),
		},
		Props: []interface{}{
			PropMapper{Prop: PropName, Fn: func(m string) string { return m + " Touch" }},
			PropVersion,
		},
	},
	// HarmonyOS and other mobile OSes
	{
		Regexes: []*Regexp2{
			MustCompile(`(harmonyos)[\/ ]?([\d\.]*)`),
			MustCompile(`(android|bada|blackberry|kaios|maemo|meego|openharmony|qnx|rim tablet os|sailfish|series40|symbian|tizen)\w*[-\/\.; ]?([\d\.]*)`),
		},
		Props: []interface{}{PropName, PropVersion},
	},
	// BlackBerry 10
	{
		Regexes: []*Regexp2{
			MustCompile(`\(bb(10);`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: VendorBlackBerry},
		},
	},
	// Symbian
	{
		Regexes: []*Regexp2{
			MustCompile(`(?:symbian ?os|symbos|s60(?=;)|series ?60)[-\/ ]?([\w\.]*)`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: "Symbian"},
		},
	},
	// Firefox OS
	{
		Regexes: []*Regexp2{
			MustCompile(`mozilla\/[\d\.]+ \((?:mobile|tablet|tv|mobile; [\w ]+); rv:.+ gecko\/([\w\.]+)`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: BrowserFirefox + " OS"},
		},
	},
	// WebOS
	{
		Regexes: []*Regexp2{
			MustCompile(`\b(?:hp)?wos(?:browser)?\/([\w\.]+)`),
			MustCompile(`webos(?:[ \/]?|\.tv-20(?=2[2-9]))(\d[\d\.]*)`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: "webOS"},
		},
	},
	// WebOS TV (Chrome version -> webOS version mapping)
	{
		Regexes: []*Regexp2{
			MustCompile(`web0s;.+?(?:chr[o0]me|safari)\/(\d+)`),
		},
		Props: []interface{}{
			PropMapper{
				Prop: PropVersion,
				StrMap: map[string]interface{}{
					"25": "120",
					"24": "108",
					"23": "94",
					"22": "87",
					"6":  "79",
					"5":  "68",
					"4":  "53",
					"3":  "38",
					"2":  "538",
					"1":  "537",
					"*":  "TV",
				},
			},
			PropMapper{Prop: PropName, Value: "webOS"},
		},
	},
	// watchOS
	{
		Regexes: []*Regexp2{
			MustCompile(`watch(?: ?os[,\/]|\d,\d\/)([\d\.]+)`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: "watchOS"},
		},
	},
	// Chrome OS
	{
		Regexes: []*Regexp2{
			MustCompile(`(cros) [\w]+(?:\)| ([\w\.]+)\b)`),
		},
		Props: []interface{}{
			PropMapper{Prop: PropName, Value: "Chrome OS"},
			PropVersion,
		},
	},
	// Smart TVs and Consoles
	{
		Regexes: []*Regexp2{
			MustCompile(`panasonic;(viera)`),
			MustCompile(`(netrange)mmh`),
			MustCompile(`(nettv)\/(\d+\.[\w\.]+)`),
			MustCompile(`(nintendo|playstation) (\w+)`),
			MustCompile(`(xbox); +xbox ([^\);]+)`),
			MustCompile(`(pico) .+os([\w\.]+)`),
		},
		Props: []interface{}{PropName, PropVersion},
	},
	// Linux distros and other OSes
	{
		Regexes: []*Regexp2{
			MustCompile(`\b(joli|palm)\b ?(?:os)?\/?([\w\.]*)`),
			MustCompile(`linux.+(mint)[\/\(\) ]?([\w\.]*)`),
			MustCompile(`(mageia|vectorlinux|fuchsia|arcaos|arch(?= ?linux))[;l ]([\d\.]*)`),
			MustCompile(`([kxln]?ubuntu|debian|suse|opensuse|gentoo|slackware|fedora|mandriva|centos|pclinuxos|red ?hat|zenwalk|linpus|raspbian|plan 9|minix|risc os|contiki|deepin|manjaro|elementary os|sabayon|linspire|knoppix)(?: gnu[\/ ]linux)?(?: enterprise)?(?:[- ]linux)?(?:-gnu)?[-\/ ]?(?!chrom|package)([-\w\.]*)`),
			MustCompile(`((?:open)?solaris)[-\/ ]?([\w\.]*)`),
			MustCompile(`\b(aix)[; ]([1-9\.]{0,4})`),
			MustCompile(`(hurd|linux|morphos)(?: (?:arm|x86|ppc)\w*| ?)([\w\.]*)`),
			MustCompile(`(gnu) ?([\w\.]*)`),
			MustCompile(`\b([-frentopcghs]{0,5}bsd|dragonfly)[\/ ]?(?!amd|[ix346]{1,2}86)([\w\.]*)`),
			MustCompile(`(haiku) ?(r\d)?`),
		},
		Props: []interface{}{
			PropName,
			PropMapper{Prop: PropVersion, Fn: func(v string) string {
				vLower := strings.ToLower(v)
				// Exclude architecture names
				if vLower == "amd" || vLower == "i386" || vLower == "i486" || vLower == "i586" || vLower == "i686" ||
					vLower == "x86_64" || vLower == "chrom" || vLower == "package" ||
					strings.HasPrefix(vLower, "chrom") {
					return ""
				}
				return v
			}},
		},
	},
	// Solaris
	{
		Regexes: []*Regexp2{
			MustCompile(`(sunos) ?([\d\.]*)`),
		},
		Props: []interface{}{
			PropMapper{Prop: PropName, Value: "Solaris"},
			PropVersion,
		},
	},
	// BeOS/OS2/AmigaOS/OpenVMS/HP-UX/SerenityOS, Unix
	{
		Regexes: []*Regexp2{
			MustCompile(`\b(beos|os\/2|amigaos|openvms|hp-ux|serenityos)`),
			MustCompile(`(unix) ?([\w\.]*)`),
		},
		Props: []interface{}{PropName, PropVersion},
	},
}
