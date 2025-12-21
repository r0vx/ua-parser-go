package uaparser

import "regexp"

var osRegexes = []RegexRule{
	// Windows RT
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)(windows nt) (6\.[23]); arm`),
		},
		Props: []interface{}{
			PropMapper{Prop: PropName, Pattern: regexp.MustCompile(`N`), Replace: "R"},
			PropMapper{Prop: PropVersion, StrMap: windowsVersionMap},
		},
	},
	// Windows IoT/Mobile/Phone
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)(windows (?:phone|mobile|iot))(?: os)?[\/ ]?([\d\.]*( se)?)`),
			regexp.MustCompile(`(?i)(windows)[\/ ](1[01]|2000|3\.1|7|8(\.1)?|9[58]|me|server 20\d\d( r2)?|vista|xp)`),
		},
		Props: []interface{}{PropName, PropVersion},
	},
	// Windows NT
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)windows nt ?([\d\.]*)`),
			regexp.MustCompile(`(?i)\bwin(?:nt| 9x )?([\d\.;]*)`),
		},
		Props: []interface{}{
			PropMapper{Prop: PropVersion, Pattern: regexp.MustCompile(`[;)]`), Replace: "", StrMap: windowsVersionMap},
			PropMapper{Prop: PropName, Value: BrowserWindows},
		},
	},
	// Windows CE
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)(windows ce)\/?([\d\.]*)`),
		},
		Props: []interface{}{PropName, PropVersion},
	},
	// iOS
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)[adehimnop]{4,7}\b(?:.*os ([\w]+) like mac|; opera)`),
			regexp.MustCompile(`(?i)(?:ios;fbsv|ip(?:ad|hone)(?: |.+i(?:pad)?)os|apple ?tv.+ios)[\/ ]([\w\.]+)`),
			regexp.MustCompile(`(?i)\btvos ?([\w\.]+)`),
			regexp.MustCompile(`(?i)cfnetwork\/.+darwin`),
		},
		Props: []interface{}{
			PropMapper{Prop: PropVersion, Pattern: regexp.MustCompile(`_`), Replace: "."},
			PropMapper{Prop: PropName, Value: "iOS"},
		},
	},
	// macOS
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)(mac os x) ?([\w\. ]*)`),
			regexp.MustCompile(`(?i)(macintosh|mac_powerpc\b)`),
		},
		Props: []interface{}{
			PropMapper{Prop: PropName, Value: "macOS"},
			PropMapper{Prop: PropVersion, Pattern: regexp.MustCompile(`_`), Replace: "."},
		},
	},
	// Chromecast Android
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)android ([\d\.]+).*crkey`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: BrowserChromecast + " Android"},
		},
	},
	// Chromecast Fuchsia
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)fuchsia.*crkey\/([\d\.]+)`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: BrowserChromecast + " Fuchsia"},
		},
	},
	// Chromecast SmartSpeaker
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)crkey\/([\d\.]+).*devicetype\/smartspeaker`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: BrowserChromecast + " SmartSpeaker"},
		},
	},
	// Chromecast Linux
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)linux.*crkey\/([\d\.]+)`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: BrowserChromecast + " Linux"},
		},
	},
	// Chromecast unknown
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)crkey\/([\d\.]+)`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: BrowserChromecast},
		},
	},
	// Android-x86
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)droid ([\w\.]+)\b.+(android[- ]x86)`),
		},
		Props: []interface{}{PropVersion, PropName},
	},
	// Ubuntu Touch
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)(ubuntu) ([\w\.]+) like android`),
		},
		Props: []interface{}{
			PropMapper{Prop: PropName, Fn: func(m string) string { return m + " Touch" }},
			PropVersion,
		},
	},
	// HarmonyOS, Android, BlackBerry, etc.
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)(harmonyos)[\/ ]?([\d\.]*)`),
			regexp.MustCompile(`(?i)(android|bada|blackberry|kaios|maemo|meego|openharmony|qnx|rim tablet os|sailfish|series40|symbian|tizen)\w*[-\/\.; ]?([\d\.]*)`),
		},
		Props: []interface{}{PropName, PropVersion},
	},
	// BlackBerry 10
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)\(bb(10);`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: VendorBlackBerry},
		},
	},
	// Symbian
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)(?:symbian ?os|symbos|s60;|series ?60)[-\/ ]?([\w\.]*)`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: "Symbian"},
		},
	},
	// Firefox OS
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)mozilla\/[\d\.]+ \((?:mobile|tablet|tv|mobile; [\w ]+); rv:.+ gecko\/([\w\.]+)`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: BrowserFirefox + " OS"},
		},
	},
	// webOS
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)\b(?:hp)?wos(?:browser)?\/([\w\.]+)`),
			regexp.MustCompile(`(?i)webos(?:[ \/]?|\.tv-20)(\d[\d\.]*)`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: "webOS"},
		},
	},
	// webOS TV
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)web0s;.+?(?:chr[o0]me|safari)\/(\d+)`),
		},
		Props: []interface{}{
			PropMapper{
				Prop: PropVersion,
				StrMap: map[string]interface{}{
					"120": "25",
					"108": "24",
					"94":  "23",
					"87":  "22",
					"79":  "6",
					"68":  "5",
					"53":  "4",
					"38":  "3",
					"538": "2",
					"537": "1",
					"*":   "TV",
				},
			},
			PropMapper{Prop: PropName, Value: "webOS"},
		},
	},
	// watchOS
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)watch(?: ?os[,\/]|\d,\d\/)([\d\.]+)`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: "watchOS"},
		},
	},
	// Chrome OS
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)(cros) [\w]+(?:\)| ([\w\.]+)\b)`),
		},
		Props: []interface{}{
			PropMapper{Prop: PropName, Value: "Chrome OS"},
			PropVersion,
		},
	},
	// Smart TVs, Consoles
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)panasonic;(viera)`),
			regexp.MustCompile(`(?i)(netrange)mmh`),
			regexp.MustCompile(`(?i)(nettv)\/(\d+\.[\w\.]+)`),
			regexp.MustCompile(`(?i)(nintendo|playstation) (\w+)`),
			regexp.MustCompile(`(?i)(xbox); +xbox ([^\);]+)`),
			regexp.MustCompile(`(?i)(pico) .+os([\w\.]+)`),
		},
		Props: []interface{}{PropName, PropVersion},
	},
	// Various Linux distros
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)\b(joli|palm)\b ?(?:os)?\/?([\w\.]*)`),
			regexp.MustCompile(`(?i)linux.+(mint)[\/\(\) ]?([\w\.]*)`),
			regexp.MustCompile(`(?i)(mageia|vectorlinux|fuchsia|arcaos|arch)[;l ]([\d\.]*)`),
			regexp.MustCompile(`(?i)([kxln]?ubuntu|debian|suse|opensuse|gentoo|slackware|fedora|mandriva|centos|pclinuxos|red ?hat|zenwalk|linpus|raspbian|plan 9|minix|risc os|contiki|deepin|manjaro|elementary os|sabayon|linspire|knoppix)(?: gnu[\/ ]linux)?(?: enterprise)?(?:[- ]linux)?(?:-gnu)?[-\/ ]?([-\w\.]*)`),
			regexp.MustCompile(`(?i)((?:open)?solaris)[-\/ ]?([\w\.]*)`),
			regexp.MustCompile(`(?i)\b(aix)[; ]([1-9\.]{0,4})`),
			regexp.MustCompile(`(?i)(hurd|linux|morphos)(?: (?:arm|x86|ppc)\w*| ?)([\w\.]*)`),
			regexp.MustCompile(`(?i)(gnu) ?([\w\.]*)`),
			regexp.MustCompile(`(?i)\b([-frentopcghs]{0,5}bsd|dragonfly)[\/ ]?([\w\.]*)`),
			regexp.MustCompile(`(?i)(haiku) ?(r\d)?`),
		},
		Props: []interface{}{PropName, PropVersion},
	},
	// Solaris
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)(sunos) ?([\d\.]*)`),
		},
		Props: []interface{}{
			PropMapper{Prop: PropName, Value: "Solaris"},
			PropVersion,
		},
	},
	// BeOS/OS2/AmigaOS/OpenVMS/HP-UX/SerenityOS, Unix
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)\b(beos|os\/2|amigaos|openvms|hp-ux|serenityos)`),
			regexp.MustCompile(`(?i)(unix) ?([\w\.]*)`),
		},
		Props: []interface{}{PropName, PropVersion},
	},
}
