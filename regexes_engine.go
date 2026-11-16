package uaparser

var engineRegexes = []RegexRule{
	// EdgeHTML
	{
		Regexes: []*Regexp2{MustCompile(`windows.+ edge\/([\w\.]+)`)},
		Props:   []interface{}{PropVersion, PropMapper{Prop: PropName, Value: BrowserEdge + "HTML"}},
	},
	// ArkWeb
	{
		Regexes: []*Regexp2{MustCompile(`(arkweb)\/([\w\.]+)`)},
		Props:   []interface{}{PropName, PropVersion},
	},
	// Blink (Chrome 28+, excluding Chrome 27)
	{
		Regexes: []*Regexp2{MustCompile(`webkit\/537\.36.+chrome\/(?!27)([\w\.]+)`)},
		Props:   []interface{}{PropVersion, PropMapper{Prop: PropName, Value: "Blink"}},
	},
	// Presto, WebKit, Trident, NetFront, NetSurf, Amaya, Lynx, w3m, Goanna, Servo
	{
		Regexes: []*Regexp2{
			MustCompile(`(presto)\/([\w\.]+)`),
			MustCompile(`(webkit|trident|netfront|netsurf|amaya|lynx|w3m|goanna|servo)\/([\w\.]+)`),
			MustCompile(`ekioh(flow)\/([\w\.]+)`),
			MustCompile(`(khtml|tasman|links|dillo)[\/ ]\(?([\w\.]+)`),
			MustCompile(`(icab)[\/ ]([23]\.[\d\.]+)`),
		},
		Props: []interface{}{PropName, PropVersion},
	},
	// LibWeb
	{
		Regexes: []*Regexp2{MustCompile(`\b(libweb)`)},
		Props:   []interface{}{PropName, PropVersion},
	},
	// Ladybird (LibWeb)
	{
		Regexes: []*Regexp2{MustCompile(`ladybird\/`)},
		Props:   []interface{}{PropMapper{Prop: PropName, Value: "LibWeb"}},
	},
	// Gecko
	{
		Regexes: []*Regexp2{MustCompile(`rv\:([\w\.]{1,9})\b.+(gecko)`)},
		Props:   []interface{}{PropVersion, PropName},
	},
}
