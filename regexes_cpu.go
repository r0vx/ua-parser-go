package uaparser

import (
	"strings"
)

var cpuRegexes = []RegexRule{
	// AMD64 (x64)
	{
		Regexes: []*Regexp2{MustCompile(`\b((amd|x|x86[-_]?|wow|win)64)\b`)},
		Props:   []interface{}{PropMapper{Prop: PropArchitecture, Value: "amd64"}},
	},
	// IA32 (x86)
	{
		Regexes: []*Regexp2{
			MustCompile(`(ia32(?=;))`),
			MustCompile(`\b((i[346]|x)86)(pc)?\b`),
		},
		Props: []interface{}{PropMapper{Prop: PropArchitecture, Value: "ia32"}},
	},
	// ARM64
	{
		Regexes: []*Regexp2{MustCompile(`\b(aarch64|arm(v?[89]e?l?|_?64))\b`)},
		Props:   []interface{}{PropMapper{Prop: PropArchitecture, Value: "arm64"}},
	},
	// ARMHF
	{
		Regexes: []*Regexp2{MustCompile(`\b(arm(v[67])?ht?n?[fl]p?)\b`)},
		Props:   []interface{}{PropMapper{Prop: PropArchitecture, Value: "armhf"}},
	},
	// ARM (PocketPC)
	{
		Regexes: []*Regexp2{MustCompile(`( (ce|mobile); ppc;|\/[\w\.]+arm\b)`)},
		Props:   []interface{}{PropMapper{Prop: PropArchitecture, Value: "arm"}},
	},
	// SPARC
	{
		Regexes: []*Regexp2{MustCompile(` sun4\w[;\)]`)},
		Props:   []interface{}{PropMapper{Prop: PropArchitecture, Value: "sparc"}},
	},
	// IA64, 68K, ARM, AVR, IRIX, MIPS, SPARC, PA-RISC
	{
		Regexes: []*Regexp2{
			MustCompile(`\b(avr32|ia64(?=;)|68k(?=\))|\barm(?=v([1-7]|[5-7]1)l?|;|eabi)|(irix|mips|sparc)(64)?\b|pa-risc)`),
			MustCompile(`((ppc|powerpc)(64)?)( mac|;|\))`),
			MustCompile(`(?:osf1|[freopnt]{3,4}bsd) (alpha)`),
		},
		Props: []interface{}{
			PropMapper{Prop: PropArchitecture, Fn: func(m string) string {
				return strings.ToLower(strings.ReplaceAll(m, "ower", ""))
			}},
		},
	},
	// 68K
	{
		Regexes: []*Regexp2{MustCompile(`mc680.0`)},
		Props:   []interface{}{PropMapper{Prop: PropArchitecture, Value: "68k"}},
	},
	// Alpha
	{
		Regexes: []*Regexp2{MustCompile(`winnt.+\[axp`)},
		Props:   []interface{}{PropMapper{Prop: PropArchitecture, Value: "alpha"}},
	},
}
