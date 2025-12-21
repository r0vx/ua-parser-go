package uaparser

import "regexp"

var engineRegexes = []RegexRule{
	// EdgeHTML
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)windows.+ edge\/([\w\.]+)`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: BrowserEdge + "HTML"},
		},
	},
	// ArkWeb
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)(arkweb)\/([\w\.]+)`),
		},
		Props: []interface{}{PropName, PropVersion},
	},
	// Blink
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)webkit\/537\.36.+chrome\/([\w\.]+)`),
		},
		Props: []interface{}{
			PropVersion,
			PropMapper{Prop: PropName, Value: "Blink"},
		},
	},
	// Presto, WebKit, Trident, NetFront, NetSurf, Amaya, Lynx, w3m, Goanna, Servo
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)(presto)\/([\w\.]+)`),
			regexp.MustCompile(`(?i)(webkit|trident|netfront|netsurf|amaya|lynx|w3m|goanna|servo)\/([\w\.]+)`),
			regexp.MustCompile(`(?i)ekioh(flow)\/([\w\.]+)`),
			regexp.MustCompile(`(?i)(khtml|tasman|links)[\/ ]\(?([\w\.]+)`),
			regexp.MustCompile(`(?i)(icab)[\/ ]([23]\.[\d\.]+)`),
		},
		Props: []interface{}{PropName, PropVersion},
	},
	// LibWeb
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)\b(libweb)`),
		},
		Props: []interface{}{PropName, PropVersion},
	},
	// Ladybird (LibWeb)
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)ladybird\/`),
		},
		Props: []interface{}{
			PropMapper{Prop: PropName, Value: "LibWeb"},
		},
	},
	// Gecko
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)rv\:([\w\.]{1,9})\b.+(gecko)`),
		},
		Props: []interface{}{PropVersion, PropName},
	},
}
