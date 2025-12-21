package uaparser

import (
	"regexp"
	"strings"
)

var cpuRegexes = []RegexRule{
	// AMD64 (x64)
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)\b((amd|x|x86[-_]?|wow|win)64)\b`),
		},
		Props: []interface{}{
			PropMapper{Prop: PropArchitecture, Value: "amd64"},
		},
	},
	// IA32 (x86)
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)\b((i[346]|x)86)\b`),
		},
		Props: []interface{}{
			PropMapper{Prop: PropArchitecture, Value: "ia32"},
		},
	},
	// ARM64
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)\b(aarch64|arm(v?[89]e?l?|_?64))\b`),
		},
		Props: []interface{}{
			PropMapper{Prop: PropArchitecture, Value: "arm64"},
		},
	},
	// ARMHF
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)\b(arm(v[67])?ht?n?[fl]p?)\b`),
		},
		Props: []interface{}{
			PropMapper{Prop: PropArchitecture, Value: "armhf"},
		},
	},
	// ARM (PocketPC)
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)( (ce|mobile); ppc;|\/[\w\.]+arm\b)`),
		},
		Props: []interface{}{
			PropMapper{Prop: PropArchitecture, Value: "arm"},
		},
	},
	// SPARC
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i) sun4\w[;\)]`),
		},
		Props: []interface{}{
			PropMapper{Prop: PropArchitecture, Value: "sparc"},
		},
	},
	// IA64, 68K, ARM, AVR, IRIX, MIPS, SPARC, PA-RISC
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)\b(avr32|ia64|68k|\barm|(irix|mips|sparc)(64)?\b|pa-risc)`),
			regexp.MustCompile(`(?i)((ppc|powerpc)(64)?)( mac|;|\))`),
			regexp.MustCompile(`(?i)(?:osf1|[freopnt]{3,4}bsd) (alpha)`),
		},
		Props: []interface{}{
			PropMapper{Prop: PropArchitecture, Fn: func(m string) string {
				return strings.ToLower(strings.ReplaceAll(m, "ower", ""))
			}},
		},
	},
	// 68K
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)mc680.0`),
		},
		Props: []interface{}{
			PropMapper{Prop: PropArchitecture, Value: "68k"},
		},
	},
	// Alpha
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)winnt.+\[axp`),
		},
		Props: []interface{}{
			PropMapper{Prop: PropArchitecture, Value: "alpha"},
		},
	},
}
