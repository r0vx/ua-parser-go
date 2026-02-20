package uaparser

import (
	"regexp"
	"strings"
)

var deviceRegexes = []RegexRule{
	// Samsung Tablets
	{
		Regexes: []*Regexp2{MustCompile(`\b(sch-i[89]0\d|shw-m380s|sm-[ptx]\w{2,4}|gt-[pn]\d{2,4}|sgh-t8[56]9|nexus 10)`)},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorSamsung},
			PropMapper{Prop: PropType, Value: DeviceTablet},
		},
	},
	// Samsung Mobiles (excluding wearables sm-[lr])
	{
		Regexes: []*Regexp2{
			MustCompile(`\b((?:s[cgp]h|gt|sm)-(?![lr])\w+|sc[g-]?[\d]+a?|galaxy nexus)`),
			MustCompile(`samsung[- ]((?!sm-[lr]|browser)[-\w]+)`),
			MustCompile(`sec-(sgh\w+)`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorSamsung},
			PropMapper{Prop: PropType, Value: DeviceMobile},
		},
	},
	// Apple iPhone/iPod
	{
		Regexes: []*Regexp2{MustCompile(`(?:\/|\()(ip(?:hone|od)[\w, ]*)[\/\);]`)},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorApple},
			PropMapper{Prop: PropType, Value: DeviceMobile},
		},
	},
	// Apple iPad
	{
		Regexes: []*Regexp2{
			MustCompile(`\b(?:ios|apple\w+)\/.+[\(\/](ipad)`),
			MustCompile(`\b(ipad)[\d,]*[;\] ].+(mac |i(pad)?)os`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorApple},
			PropMapper{Prop: PropType, Value: DeviceTablet},
		},
	},
	// Apple Macintosh
	{
		Regexes: []*Regexp2{MustCompile(`(macintosh);`)},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorApple},
		},
	},
	// Sharp
	{
		Regexes: []*Regexp2{MustCompile(`\b(sh-?[altvz]?\d\d[a-ekm]?)`)},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorSharp},
			PropMapper{Prop: PropType, Value: DeviceMobile},
		},
	},
	// Honor Tablets
	{
		Regexes: []*Regexp2{MustCompile(`\b((?:brt|eln|hey2?|gdi|jdn)-a?[lnw]09|(?:ag[rm]3?|jdn2|kob2)-a?[lw]0[09]hn)(?: bui|\)|;)`)},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorHonor},
			PropMapper{Prop: PropType, Value: DeviceTablet},
		},
	},
	// Honor Mobiles
	{
		Regexes: []*Regexp2{MustCompile(`honor([-\w ]+)[;\)]`)},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorHonor},
			PropMapper{Prop: PropType, Value: DeviceMobile},
		},
	},
	// Huawei Tablets
	{
		Regexes: []*Regexp2{MustCompile(`\b((?:ag[rs][2356]?k?|bah[234]?|bg[2o]|bt[kv]|cmr|cpn|db[ry]2?|jdn2|got|kob2?k?|mon|pce|scm|sht?|[tw]gr|vrd)-[ad]?[lw][0125][09]b?|605hw|bg2-u03|(?:gem|fdr|m2|ple|t1)-[7a]0[1-4][lu]|t1-a2[13][lw]|mediapad[\w\. ]*(?= bui|\)))\b(?!.+d\/s)`)},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorHuawei},
			PropMapper{Prop: PropType, Value: DeviceTablet},
		},
	},
	// Huawei Mobiles
	{
		Regexes: []*Regexp2{
			MustCompile(`(?:huawei) ?([-\w ]+)[;\)]`),
			MustCompile(`\b(nexus 6p|\w{2,4}e?-[atu]?[ln][\dx][\dc][adnt]?)\b(?!.+d\/s)`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorHuawei},
			PropMapper{Prop: PropType, Value: DeviceMobile},
		},
	},
	// Xiaomi Tablets
	{
		Regexes: []*Regexp2{
			MustCompile(`oid[^\)]+; (2[\dbc]{4}(182|283|rp\w{2})[cgl]|m2105k81a?c)(?: bui|\))`),
			MustCompile(`\b(?:xiao)?((?:red)?mi[-_ ]?pad[\w- ]*)(?: bui|\))`),
		},
		Props: []interface{}{
			PropMapper{Prop: PropModel, Pattern: regexp.MustCompile(`_`), Replace: " "},
			PropMapper{Prop: PropVendor, Value: VendorXiaomi},
			PropMapper{Prop: PropType, Value: DeviceTablet},
		},
	},
	// Xiaomi Mobiles
	{
		Regexes: []*Regexp2{
			MustCompile(`\b; (\w+) build\/hm\1`),
			MustCompile(`\b(hm[-_ ]?note?[_ ]?(?:\d\w)?) bui`),
			MustCompile(`oid[^\)]+; (redmi[\-_ ]?(?:note|k)?[\w_ ]+|m?[12]\d[01]\d\w{3,6}|poco[\w ]+|(shark )?\w{3}-[ah]0|qin ?[1-3](s\+|ultra| pro)?)( bui|; wv|\))`),
			MustCompile(`\b(mi[-_ ]?(?:a\d|one|one[_ ]plus|note|max|cc)?[_ ]?(?:\d{0,2}\w?)[_ ]?(?:plus|se|lite|pro)?( 5g|lte)?)(?: bui|\))`),
			MustCompile(` ([\w ]+) miui\/v?\d`),
		},
		Props: []interface{}{
			PropMapper{Prop: PropModel, Pattern: regexp.MustCompile(`_`), Replace: " "},
			PropMapper{Prop: PropVendor, Value: VendorXiaomi},
			PropMapper{Prop: PropType, Value: DeviceMobile},
		},
	},
	// Xiaomi Smart TVs
	{
		Regexes: []*Regexp2{MustCompile(`(mi(tv|box)-?\w+) bui`)},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorXiaomi},
			PropMapper{Prop: PropType, Value: DeviceSmartTV},
		},
	},
	// OnePlus
	{
		Regexes: []*Regexp2{
			MustCompile(`droid.+; (cph2[3-6]\d[13579]|((gm|hd)19|(ac|be|in|kb)20|(d[en]|eb|le|mt)21|ne22)[0-2]\d|p[g-k]\w[1m]10)\b`),
			MustCompile(`(?:one)?(?:plus)? (a\d0\d\d)(?: b|\))`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorOnePlus},
			PropMapper{Prop: PropType, Value: DeviceMobile},
		},
	},
	// OPPO
	{
		Regexes: []*Regexp2{
			MustCompile(`; (\w+) bui.+ oppo`),
			MustCompile(`\b(cph[12]\d{3}|p(?:af|c[al]|d\w|e[ar])[mt]\d0|x9007|a101op)\b`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorOppo},
			PropMapper{Prop: PropType, Value: DeviceMobile},
		},
	},
	// OPPO/OnePlus Tablets
	{
		Regexes: []*Regexp2{MustCompile(`\b(opd2(\d{3}a?))(?: bui|\))`)},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Fn: func(m string) string {
				if strings.Contains(m, "203") || strings.Contains(m, "304") || strings.Contains(m, "403") || strings.Contains(m, "404") || strings.Contains(m, "413") || strings.Contains(m, "415") {
					return VendorOnePlus
				}
				return VendorOppo
			}},
			PropMapper{Prop: PropType, Value: DeviceTablet},
		},
	},
	// BLU Vivo
	{
		Regexes: []*Regexp2{MustCompile(`(vivo (5r?|6|8l?|go|one|s|x[il]?[2-4]?)[\w\+ ]*)(?: bui|\))`)},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: "BLU"},
			PropMapper{Prop: PropType, Value: DeviceMobile},
		},
	},
	// Vivo
	{
		Regexes: []*Regexp2{
			MustCompile(`; vivo (\w+)(?: bui|\))`),
			MustCompile(`\b(v[12]\d{3}\w?[at])(?: bui|;)`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: "Vivo"},
			PropMapper{Prop: PropType, Value: DeviceMobile},
		},
	},
	// Realme
	{
		Regexes: []*Regexp2{MustCompile(`\b(rmx[1-3]\d{3})(?: bui|;|\))`)},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: "Realme"},
			PropMapper{Prop: PropType, Value: DeviceMobile},
		},
	},
	// Lenovo Tablets
	{
		Regexes: []*Regexp2{
			MustCompile(`(ideatab[-\w ]+|602lv|d-42a|a101lv|a2109a|a3500-hv|s[56]000|pb-6505[my]|tb-?x?\d{3,4}(?:f[cu]|xu|[av])|yt\d?-[jx]?\d+[lfmx])( bui|;|\)|\/)`),
			MustCompile(`lenovo ?(b[68]0[08]0-?[hf]?|tab(?:[\w- ]+?)|tb[\w-]{6,7})( bui|;|\)|\/)`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorLenovo},
			PropMapper{Prop: PropType, Value: DeviceTablet},
		},
	},
	// Lenovo Mobiles
	{
		Regexes: []*Regexp2{MustCompile(`lenovo[-_ ]?([-\w ]+?)(?: bui|\)|\/)`)},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorLenovo},
			PropMapper{Prop: PropType, Value: DeviceMobile},
		},
	},
	// Motorola
	{
		Regexes: []*Regexp2{
			MustCompile(`\b(milestone|droid(?:[2-4x]| (?:bionic|x2|pro|razr))?:?( 4g)?)\b[\w ]+build\/`),
			MustCompile(`\bmot(?:orola)?[- ]([\w\s]+)(\)| bui)`),
			MustCompile(`((?:moto(?! 360)[-\w\(\) ]+|xt\d{3,4}[cgkosw\+]?[-\d]*|nexus 6)(?= bui|\)))`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorMotorola},
			PropMapper{Prop: PropType, Value: DeviceMobile},
		},
	},
	// Motorola Tablets
	{
		Regexes: []*Regexp2{MustCompile(`\b(mz60\d|xoom[2 ]{0,2}) build\/`)},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorMotorola},
			PropMapper{Prop: PropType, Value: DeviceTablet},
		},
	},
	// LG Tablets
	{
		Regexes: []*Regexp2{MustCompile(`((?=lg)?[vl]k\-?\d{3}) bui| 3\.[-\w; ]{10}lg?-([06cv9]{3,4})`)},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorLG},
			PropMapper{Prop: PropType, Value: DeviceTablet},
		},
	},
	// LG Mobiles
	{
		Regexes: []*Regexp2{
			MustCompile(`(lm(?:-?f100[nv]?|-[\w\.]+)(?= bui|\))|nexus [45])`),
			MustCompile(`\blg[-e;\/ ]+(?!.*(?:browser|netcast|android tv|watch|webos))(\w+)`),
			MustCompile(`\blg-?([\d\w]+) bui`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorLG},
			PropMapper{Prop: PropType, Value: DeviceMobile},
		},
	},
	// Nokia Tablets
	{
		Regexes: []*Regexp2{MustCompile(`(nokia) (t[12][01])`)},
		Props: []interface{}{
			PropVendor,
			PropModel,
			PropMapper{Prop: PropType, Value: DeviceTablet},
		},
	},
	// Nokia Mobiles
	{
		Regexes: []*Regexp2{
			MustCompile(`(?:maemo|nokia).*(n900|lumia \d+|rm-\d+)`),
			MustCompile(`nokia[-_ ]?(([-\w\. ]*?))( bui|\)|;|\/)`),
		},
		Props: []interface{}{
			PropMapper{Prop: PropModel, Pattern: regexp.MustCompile(`_`), Replace: " "},
			PropMapper{Prop: PropType, Value: DeviceMobile},
			PropMapper{Prop: PropVendor, Value: "Nokia"},
		},
	},
	// Google Pixel Tablets
	{
		Regexes: []*Regexp2{MustCompile(`(pixel (c|tablet))\b`)},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorGoogle},
			PropMapper{Prop: PropType, Value: DeviceTablet},
		},
	},
	// Google Pixel Mobiles
	{
		Regexes: []*Regexp2{MustCompile(`droid.+;(?: google)? (g(01[13]a|020[aem]|025[jn]|1b60|1f8f|2ybb|4s1m|576d|5nz6|8hhn|8vou|a02099|c15s|d1yq|e2ae|ec77|gh2x|kv4x|p4bc|pj41|r83y|tt9q|ur25|wvk6)|pixel[\d ]*a?( pro)?( xl)?( fold)?( \(5g\))?)( bui|\))`)},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorGoogle},
			PropMapper{Prop: PropType, Value: DeviceMobile},
		},
	},
	// Google Pixelbook
	{
		Regexes: []*Regexp2{MustCompile(`(google) (pixelbook( go)?)`)},
		Props:   []interface{}{PropVendor, PropModel},
	},
	// Sony Mobiles
	{
		Regexes: []*Regexp2{MustCompile(`droid.+; (a?\d[0-2]{2}so|[c-g]\d{4}|so[-gl]\w+|xq-\w\w\d\d)(?= bui|\).+chrome\/(?![1-6]{0,1}\d\.))`)},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorSony},
			PropMapper{Prop: PropType, Value: DeviceMobile},
		},
	},
	// Sony Tablets
	{
		Regexes: []*Regexp2{
			MustCompile(`sony tablet [ps]`),
			MustCompile(`\b(?:sony)?sgp\w+(?: bui|\))`),
		},
		Props: []interface{}{
			PropMapper{Prop: PropModel, Value: "Xperia Tablet"},
			PropMapper{Prop: PropVendor, Value: VendorSony},
			PropMapper{Prop: PropType, Value: DeviceTablet},
		},
	},
	// Amazon Tablets
	{
		Regexes: []*Regexp2{
			MustCompile(`(alexa)webm`),
			MustCompile(`(kf[a-z]{2}wi|aeo(?!bc)\w\w)( bui|\))`),
			MustCompile(`(kf[a-z]+)( bui|\)).+silk\/`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorAmazon},
			PropMapper{Prop: PropType, Value: DeviceTablet},
		},
	},
	// Amazon Fire Phone
	{
		Regexes: []*Regexp2{MustCompile(`((?:sd|kf)[0349hijorstuw]+)( bui|\)).+silk\/`)},
		Props: []interface{}{
			PropMapper{Prop: PropModel, Fn: func(m string) string { return "Fire Phone " + m }},
			PropMapper{Prop: PropVendor, Value: VendorAmazon},
			PropMapper{Prop: PropType, Value: DeviceMobile},
		},
	},
	// BlackBerry PlayBook
	{
		Regexes: []*Regexp2{MustCompile(`(playbook);[-\w\),; ]+(rim)`)},
		Props: []interface{}{
			PropModel,
			PropVendor,
			PropMapper{Prop: PropType, Value: DeviceTablet},
		},
	},
	// BlackBerry Mobiles
	{
		Regexes: []*Regexp2{
			MustCompile(`\b((?:bb[a-f]|st[hv])100-\d)`),
			MustCompile(`(?:blackberry|\(bb10;) (\w+)`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorBlackBerry},
			PropMapper{Prop: PropType, Value: DeviceMobile},
		},
	},
	// ASUS Tablets
	{
		Regexes: []*Regexp2{MustCompile(`(?:\b|asus_)(transfo[prime ]{4,10} \w+|eeepc|slider \w+|nexus 7|padfone|p00[cj])`)},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorAsus},
			PropMapper{Prop: PropType, Value: DeviceTablet},
		},
	},
	// ASUS ZenFone
	{
		Regexes: []*Regexp2{MustCompile(` (z[bes]6[027][012][km][ls]|zenfone \d\w?)\b`)},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorAsus},
			PropMapper{Prop: PropType, Value: DeviceMobile},
		},
	},
	// HTC Nexus 9
	{
		Regexes: []*Regexp2{MustCompile(`(nexus 9)`)},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: "HTC"},
			PropMapper{Prop: PropType, Value: DeviceTablet},
		},
	},
	// HTC
	{
		Regexes: []*Regexp2{MustCompile(`(htc)[-;_ ]{1,2}([\w ]+(?=\)| bui)|\w+)`)},
		Props: []interface{}{
			PropVendor,
			PropMapper{Prop: PropModel, Pattern: regexp.MustCompile(`_`), Replace: " "},
			PropMapper{Prop: PropType, Value: DeviceMobile},
		},
	},
	// ZTE/Alcatel/GeeksPhone/Nexian/Panasonic/Sony
	{
		Regexes: []*Regexp2{
			MustCompile(`(zte)[- ]([\w ]+?)(?: bui|\/|\))`),
			MustCompile(`(alcatel|geeksphone|nexian|panasonic(?!(?:;|\.))|sony(?!-bra))[-_ ]?([-\w]*)`),
		},
		Props: []interface{}{
			PropVendor,
			PropMapper{Prop: PropModel, Pattern: regexp.MustCompile(`_`), Replace: " "},
			PropMapper{Prop: PropType, Value: DeviceMobile},
		},
	},
	// TCL Tablets
	{
		Regexes: []*Regexp2{
			MustCompile(`tcl (xess p17aa)`),
			MustCompile(`droid [\w\.]+; ((?:8[14]9[16]|9(?:0(?:48|60|8[01])|1(?:3[27]|66)|2(?:6[69]|9[56])|466))[gqswx])(_\w(\w|\w\w))?(\)| bui)`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: "TCL"},
			PropMapper{Prop: PropType, Value: DeviceTablet},
		},
	},
	// TCL Mobiles
	{
		Regexes: []*Regexp2{MustCompile(`droid [\w\.]+; (418(?:7d|8v)|5087z|5102l|61(?:02[dh]|25[adfh]|27[ai]|56[dh]|59k|65[ah])|a509dl|t(?:43(?:0w|1[adepqu])|50(?:6d|7[adju])|6(?:09dl|10k|12b|71[efho]|76[hjk])|7(?:66[ahju]|67[hw]|7[045][bh]|71[hk]|73o|76[ho]|79w|81[hks]?|82h|90[bhsy]|99b)|810[hs]))(_\w(\w|\w\w))?(\)| bui)`)},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: "TCL"},
			PropMapper{Prop: PropType, Value: DeviceMobile},
		},
	},
	// itel
	{
		Regexes: []*Regexp2{MustCompile(`(itel) ((\w+))`)},
		Props: []interface{}{
			PropMapper{Prop: PropVendor, Fn: strings.ToLower},
			PropModel,
			PropMapper{Prop: PropType, Fn: func(m string) string {
				ml := strings.ToLower(m)
				if ml == "p10001l" || ml == "w7001" {
					return DeviceTablet
				}
				return DeviceMobile
			}},
		},
	},
	// Acer Tablets
	{
		Regexes: []*Regexp2{MustCompile(`droid.+; ([ab][1-7]-?[0178a]\d\d?)`)},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: "Acer"},
			PropMapper{Prop: PropType, Value: DeviceTablet},
		},
	},
	// Meizu
	{
		Regexes: []*Regexp2{
			MustCompile(`droid.+; (m[1-5] note) bui`),
			MustCompile(`\bmz-([-\w]{2,})`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: "Meizu"},
			PropMapper{Prop: PropType, Value: DeviceMobile},
		},
	},
	// Ulefone
	{
		Regexes: []*Regexp2{MustCompile(`; ((?:power )?armor(?:[\w ]{0,8}))(?: bui|\))`)},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: "Ulefone"},
			PropMapper{Prop: PropType, Value: DeviceMobile},
		},
	},
	// Energizer
	{
		Regexes: []*Regexp2{
			MustCompile(`; (energy ?\w+)(?: bui|\))`),
			MustCompile(`; energizer ([\w ]+)(?: bui|\))`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: "Energizer"},
			PropMapper{Prop: PropType, Value: DeviceMobile},
		},
	},
	// Cat
	{
		Regexes: []*Regexp2{
			MustCompile(`; cat (b35);`),
			MustCompile(`; (b15q?|s22 flip|s48c|s62 pro)(?: bui|\))`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: "Cat"},
			PropMapper{Prop: PropType, Value: DeviceMobile},
		},
	},
	// Smartfren
	{
		Regexes: []*Regexp2{MustCompile(`((?:new )?andromax[\w- ]+)(?: bui|\))`)},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: "Smartfren"},
			PropMapper{Prop: PropType, Value: DeviceMobile},
		},
	},
	// Nothing
	{
		Regexes: []*Regexp2{MustCompile(`droid.+; (a(in)?(0(15|59|6[35])|142)p?)`)},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: "Nothing"},
			PropMapper{Prop: PropType, Value: DeviceMobile},
		},
	},
	// Archos Tablets
	{
		Regexes: []*Regexp2{
			MustCompile(`; (x67 5g|tikeasy \w+|ac[1789]\d\w+)( b|\))`),
			MustCompile(`archos ?(5|gamepad2?|([\w ]*[t1789]|hello) ?\d+[\w ]*)( b|\))`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: "Archos"},
			PropMapper{Prop: PropType, Value: DeviceTablet},
		},
	},
	// Archos Mobiles
	{
		Regexes: []*Regexp2{
			MustCompile(`archos ([\w ]+)( b|\))`),
			MustCompile(`; (ac[3-6]\d\w{2,8})( b|\))`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: "Archos"},
			PropMapper{Prop: PropType, Value: DeviceMobile},
		},
	},
	// HMD
	{
		Regexes: []*Regexp2{MustCompile(`; (n159v)`)},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: "HMD"},
			PropMapper{Prop: PropType, Value: DeviceMobile},
		},
	},
	// Mixed Tablets (IMO, Infinix, Tecno)
	{
		Regexes: []*Regexp2{
			MustCompile(`(imo) (tab \w+)`),
			MustCompile(`(infinix|tecno) (x1101b?|p904|dp(7c|8d|10a)( pro)?|p70[1-3]a?|p904|t1101)`),
		},
		Props: []interface{}{
			PropVendor,
			PropModel,
			PropMapper{Prop: PropType, Value: DeviceTablet},
		},
	},
	// Mixed Mobiles
	{
		Regexes: []*Regexp2{
			MustCompile(`(blackberry|benq|palm(?=\-)|sonyericsson|acer|asus(?! zenw)|dell|jolla|meizu|motorola|polytron|tecno|micromax|advan)[-_ ]?([-\w]*)`),
			MustCompile(`; (blu|hmd|imo|infinix|lava|oneplus|tcl|wiko)[_ ]([\w\+ ]+?)(?: bui|\)|; r)`),
			MustCompile(`(hp) ([\w ]+\w)`),
			MustCompile(`(microsoft); (lumia[\w ]+)`),
			MustCompile(`(oppo) ?([\w ]+) bui`),
			MustCompile(`(hisense) ([ehv][\w ]+)\)`),
			MustCompile(`droid[^;]+; (philips)[_ ]([sv-x][\d]{3,4}[xz]?)`),
		},
		Props: []interface{}{
			PropVendor,
			PropModel,
			PropMapper{Prop: PropType, Value: DeviceMobile},
		},
	},
	// Kobo, HP TouchPad, Kindle
	{
		Regexes: []*Regexp2{
			MustCompile(`(kobo)\s(ereader|touch)`),
			MustCompile(`(hp).+(touchpad(?!.+tablet)|tablet)`),
			MustCompile(`(kindle)\/([\w\.]+)`),
		},
		Props: []interface{}{
			PropVendor,
			PropModel,
			PropMapper{Prop: PropType, Value: DeviceTablet},
		},
	},
	// Surface Duo
	{
		Regexes: []*Regexp2{MustCompile(`(surface duo)`)},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorMicrosoft},
			PropMapper{Prop: PropType, Value: DeviceTablet},
		},
	},
	// Fairphone
	{
		Regexes: []*Regexp2{MustCompile(`droid [\d\.]+; (fp\du?)(?: b|\))`)},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: "Fairphone"},
			PropMapper{Prop: PropType, Value: DeviceMobile},
		},
	},
	// Nvidia Tablets
	{
		Regexes: []*Regexp2{MustCompile(`((?:tegranote|shield t(?!.+d tv))[\w- ]*?)(?: b|\))`)},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorNvidia},
			PropMapper{Prop: PropType, Value: DeviceTablet},
		},
	},
	// Sprint Phones
	{
		Regexes: []*Regexp2{MustCompile(`(sprint) (\w+)`)},
		Props: []interface{}{
			PropVendor,
			PropModel,
			PropMapper{Prop: PropType, Value: DeviceMobile},
		},
	},
	// Microsoft Kin
	{
		Regexes: []*Regexp2{MustCompile(`(kin\.[onetw]{3})`)},
		Props: []interface{}{
			PropMapper{Prop: PropModel, Pattern: regexp.MustCompile(`\.`), Replace: " "},
			PropMapper{Prop: PropVendor, Value: VendorMicrosoft},
			PropMapper{Prop: PropType, Value: DeviceMobile},
		},
	},
	// Zebra Tablets
	{
		Regexes: []*Regexp2{MustCompile(`droid.+; ([c6]+|et5[16]|mc[239][23]x?|vc8[03]x?)\)`)},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorZebra},
			PropMapper{Prop: PropType, Value: DeviceTablet},
		},
	},
	// Zebra Mobiles
	{
		Regexes: []*Regexp2{MustCompile(`droid.+; (ec30|ps20|tc[2-8]\d[kx])\)`)},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorZebra},
			PropMapper{Prop: PropType, Value: DeviceMobile},
		},
	},
	// Roku
	{
		Regexes: []*Regexp2{MustCompile(`\b(roku)[\dx]*[\)\/]((?:dvp-)?[\d\.]*)`)},
		Props: []interface{}{
			PropVendor,
			PropModel,
			PropMapper{Prop: PropType, Value: DeviceSmartTV},
		},
	},
	// TechniSat
	{
		Regexes: []*Regexp2{MustCompile(`[Hh]bbtv.*(technisat) (.*);`)},
		Props: []interface{}{
			PropVendor,
			PropModel,
			PropMapper{Prop: PropType, Value: DeviceSmartTV},
		},
	},
	// HbbTV devices (generic - Panasonic, Philips, Samsung, JVC, LOEWE, etc.)
	{
		Regexes: []*Regexp2{MustCompile(`hbbtv\/\d+\.\d+\.\d+ +\([\w\+ ]*; *([\w\d][^;]*);([^;]*)`)},
		Props: []interface{}{
			PropMapper{Prop: PropVendor, Fn: func(m string) string {
				// Extract last part after slash and map LGE -> LG
				v := m
				if idx := strings.LastIndex(m, "/"); idx >= 0 {
					v = m[idx+1:]
				}
				if strings.ToLower(v) == "lge" {
					return "LG"
				}
				return strings.TrimSpace(v)
			}},
			PropMapper{Prop: PropModel, Pattern: regexp.MustCompile(`^\s+|\s+$`), Replace: ""},
			PropMapper{Prop: PropType, Value: DeviceSmartTV},
		},
	},
	// Smart TVs
	{
		Regexes: []*Regexp2{
			MustCompile(`(philips)[\w ]+tv`),
			MustCompile(`smart-tv.+(samsung)`),
		},
		Props: []interface{}{
			PropVendor,
			PropMapper{Prop: PropType, Value: DeviceSmartTV},
		},
	},
	{
		Regexes: []*Regexp2{MustCompile(`hbbtv.+maple;(\d+)`)},
		Props: []interface{}{
			PropMapper{Prop: PropModel, Fn: func(m string) string { return "SmartTV" + m }},
			PropMapper{Prop: PropVendor, Value: VendorSamsung},
			PropMapper{Prop: PropType, Value: DeviceSmartTV},
		},
	},
	{
		Regexes: []*Regexp2{
			MustCompile(`(vizio)(?: |.+model\/)(\w+-\w+)`),
			MustCompile(`tcast.+(lg)e?. ([-\w]+)`),
		},
		Props: []interface{}{
			PropVendor,
			PropModel,
			PropMapper{Prop: PropType, Value: DeviceSmartTV},
		},
	},
	{
		Regexes: []*Regexp2{MustCompile(`(nux; netcast.+smarttv|lg (netcast\.tv-201\d|android tv))`)},
		Props: []interface{}{
			PropMapper{Prop: PropVendor, Value: VendorLG},
			PropMapper{Prop: PropType, Value: DeviceSmartTV},
		},
	},
	// LG webOS TV
	{
		Regexes: []*Regexp2{MustCompile(`web0s.+lg browser[^;]*;[^;]*;([^;]+)`)},
		Props: []interface{}{
			PropMapper{Prop: PropModel, Pattern: regexp.MustCompile(`^\s+|\s+$`), Replace: ""},
			PropMapper{Prop: PropVendor, Value: VendorLG},
			PropMapper{Prop: PropType, Value: DeviceSmartTV},
		},
	},
	// Apple TV
	{
		Regexes: []*Regexp2{MustCompile(`(apple) ?tv`)},
		Props: []interface{}{
			PropVendor,
			PropMapper{Prop: PropModel, Value: VendorApple + " TV"},
			PropMapper{Prop: PropType, Value: DeviceSmartTV},
		},
	},
	// Google Chromecast
	{
		Regexes: []*Regexp2{MustCompile(`crkey.*devicetype\/chromecast`)},
		Props: []interface{}{
			PropMapper{Prop: PropModel, Value: BrowserChromecast + " Third Generation"},
			PropMapper{Prop: PropVendor, Value: VendorGoogle},
			PropMapper{Prop: PropType, Value: DeviceSmartTV},
		},
	},
	{
		Regexes: []*Regexp2{MustCompile(`crkey.*devicetype\/([^/]*)`)},
		Props: []interface{}{
			PropMapper{Prop: PropModel, Fn: func(m string) string { return "Chromecast " + m }},
			PropMapper{Prop: PropVendor, Value: VendorGoogle},
			PropMapper{Prop: PropType, Value: DeviceSmartTV},
		},
	},
	{
		Regexes: []*Regexp2{MustCompile(`fuchsia.*crkey`)},
		Props: []interface{}{
			PropMapper{Prop: PropModel, Value: BrowserChromecast + " Nest Hub"},
			PropMapper{Prop: PropVendor, Value: VendorGoogle},
			PropMapper{Prop: PropType, Value: DeviceSmartTV},
		},
	},
	{
		Regexes: []*Regexp2{MustCompile(`crkey`)},
		Props: []interface{}{
			PropMapper{Prop: PropModel, Value: BrowserChromecast},
			PropMapper{Prop: PropVendor, Value: VendorGoogle},
			PropMapper{Prop: PropType, Value: DeviceSmartTV},
		},
	},
	// Facebook Portal TV
	{
		Regexes: []*Regexp2{MustCompile(`(portaltv)`)},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: BrowserFacebook},
			PropMapper{Prop: PropType, Value: DeviceSmartTV},
		},
	},
	// Amazon Fire TV
	{
		Regexes: []*Regexp2{MustCompile(`droid.+aft(\w+)( bui|\))`)},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorAmazon},
			PropMapper{Prop: PropType, Value: DeviceSmartTV},
		},
	},
	// Nvidia Shield TV
	{
		Regexes: []*Regexp2{MustCompile(`(shield \w+ tv)`)},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorNvidia},
			PropMapper{Prop: PropType, Value: DeviceSmartTV},
		},
	},
	// Sharp/Sony TV
	{
		Regexes: []*Regexp2{
			MustCompile(`\(dtv[\);].+(aquos)`),
			MustCompile(`(aquos-tv[\w ]+)\)`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorSharp},
			PropMapper{Prop: PropType, Value: DeviceSmartTV},
		},
	},
	{
		Regexes: []*Regexp2{MustCompile(`(bravia[\w ]+)( bui|\))`)},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorSony},
			PropMapper{Prop: PropType, Value: DeviceSmartTV},
		},
	},
	// Consoles
	{
		Regexes: []*Regexp2{MustCompile(`(playstation \w+)`)},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorSony},
			PropMapper{Prop: PropType, Value: DeviceConsole},
		},
	},
	{
		Regexes: []*Regexp2{MustCompile(`\b(xbox(?: one)?(?!; xbox))[\); ]`)},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorMicrosoft},
			PropMapper{Prop: PropType, Value: DeviceConsole},
		},
	},
	{
		Regexes: []*Regexp2{
			MustCompile(`(ouya)`),
			MustCompile(`(nintendo) (\w+)`),
			MustCompile(`(retroid) (pocket ([^\)]+))`),
			MustCompile(`(valve).+(steam deck)`),
		},
		Props: []interface{}{
			PropVendor,
			PropModel,
			PropMapper{Prop: PropType, Value: DeviceConsole},
		},
	},
	// Nvidia Shield/Anbernic/Logitech Consoles (with vendor mapping)
	{
		Regexes: []*Regexp2{MustCompile(`droid.+; ((shield|rgcube|gr0006))( bui|\))`)},
		Props: []interface{}{
			PropMapper{Prop: PropVendor, Fn: func(m string) string {
				ml := strings.ToLower(m)
				if ml == "shield" {
					return VendorNvidia
				}
				if ml == "rgcube" {
					return "Anbernic"
				}
				if ml == "gr0006" {
					return "Logitech"
				}
				return m
			}},
			PropModel,
			PropMapper{Prop: PropType, Value: DeviceConsole},
		},
	},
	// Wearables
	{
		Regexes: []*Regexp2{MustCompile(`\b(sm-[lr]\d\d[0156][fnuw]?s?|gear live)\b`)},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorSamsung},
			PropMapper{Prop: PropType, Value: DeviceWearable},
		},
	},
	{
		Regexes: []*Regexp2{MustCompile(`((pebble))app`)},
		Props: []interface{}{
			PropVendor,
			PropModel,
			PropMapper{Prop: PropType, Value: DeviceWearable},
		},
	},
	{
		Regexes: []*Regexp2{MustCompile(`(asus|google|lg|oppo|xiaomi) ((pixel |zen)?watch[\w ]*)( bui|\))`)},
		Props: []interface{}{
			PropVendor,
			PropModel,
			PropMapper{Prop: PropType, Value: DeviceWearable},
		},
	},
	{
		Regexes: []*Regexp2{MustCompile(`(ow(?:19|20)?we?[1-3]{1,3})`)},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorOppo},
			PropMapper{Prop: PropType, Value: DeviceWearable},
		},
	},
	// Apple Watch
	{
		Regexes: []*Regexp2{MustCompile(`(watch)(?: ?os[,\/]|\d,\d\/)[\d\.]+`)},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorApple},
			PropMapper{Prop: PropType, Value: DeviceWearable},
		},
	},
	// OnePlus Watch
	{
		Regexes: []*Regexp2{MustCompile(`(opwwe\d{3})`)},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorOnePlus},
			PropMapper{Prop: PropType, Value: DeviceWearable},
		},
	},
	// Motorola 360
	{
		Regexes: []*Regexp2{MustCompile(`(moto 360)`)},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorMotorola},
			PropMapper{Prop: PropType, Value: DeviceWearable},
		},
	},
	// Sony SmartWatch
	{
		Regexes: []*Regexp2{MustCompile(`(smartwatch 3)`)},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorSony},
			PropMapper{Prop: PropType, Value: DeviceWearable},
		},
	},
	// LG G Watch R
	{
		Regexes: []*Regexp2{MustCompile(`(g watch r)`)},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorLG},
			PropMapper{Prop: PropType, Value: DeviceWearable},
		},
	},
	// Zebra Wearables
	{
		Regexes: []*Regexp2{MustCompile(`droid.+; (wt63?0{2,3})\)`)},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorZebra},
			PropMapper{Prop: PropType, Value: DeviceWearable},
		},
	},
	// XR - Google Glass
	{
		Regexes: []*Regexp2{MustCompile(`droid.+; (glass) \d`)},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorGoogle},
			PropMapper{Prop: PropType, Value: DeviceXR},
		},
	},
	// XR - Pico
	{
		Regexes: []*Regexp2{MustCompile(`(pico) ([\w ]+) os\d`)},
		Props: []interface{}{
			PropVendor,
			PropModel,
			PropMapper{Prop: PropType, Value: DeviceXR},
		},
	},
	// XR - Meta Quest
	{
		Regexes: []*Regexp2{MustCompile(`(quest( \d| pro)?s?).+vr`)},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: BrowserFacebook},
			PropMapper{Prop: PropType, Value: DeviceXR},
		},
	},
	// XR - Firefox Reality/Wolvic
	{
		Regexes: []*Regexp2{MustCompile(`mobile vr; rv.+firefox`)},
		Props:   []interface{}{PropMapper{Prop: PropType, Value: DeviceXR}},
	},
	// Embedded - Tesla
	{
		Regexes: []*Regexp2{MustCompile(`(tesla)(?: qtcarbrowser|\/[-\w\.]+)`)},
		Props: []interface{}{
			PropVendor,
			PropMapper{Prop: PropType, Value: DeviceEmbedded},
		},
	},
	// Embedded - Echo Dot
	{
		Regexes: []*Regexp2{MustCompile(`(aeobc)\b`)},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorAmazon},
			PropMapper{Prop: PropType, Value: DeviceEmbedded},
		},
	},
	// Embedded - Apple HomePod
	{
		Regexes: []*Regexp2{MustCompile(`(homepod).+mac os`)},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorApple},
			PropMapper{Prop: PropType, Value: DeviceEmbedded},
		},
	},
	// Embedded - Windows IoT
	{
		Regexes: []*Regexp2{MustCompile(`windows iot`)},
		Props:   []interface{}{PropMapper{Prop: PropType, Value: DeviceEmbedded}},
	},
	// Generic SmartTV
	{
		Regexes: []*Regexp2{MustCompile(`droid.+; ([\w- ]+) (4k|android|smart|google)[- ]?tv`)},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropType, Value: DeviceSmartTV},
		},
	},
	{
		Regexes: []*Regexp2{MustCompile(`\b((4k|android|smart|opera)[- ]?tv|tv; rv:|large screen[\w ]+safari)\b`)},
		Props:   []interface{}{PropMapper{Prop: PropType, Value: DeviceSmartTV}},
	},
	// Generic Tablet/Mobile detection
	{
		Regexes: []*Regexp2{MustCompile(`droid .+?; ([^;]+?)(?: bui|; wv\)|\) applew|; hmsc).+?(mobile|vr|\d) safari`)},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropType, Fn: func(m string) string {
				ml := strings.ToLower(m)
				if ml == "mobile" {
					return DeviceMobile
				}
				if ml == "vr" {
					return DeviceXR
				}
				return DeviceTablet
			}},
		},
	},
	{
		Regexes: []*Regexp2{MustCompile(`\b((tablet|tab)[;\/]|focus\/\d(?!.+mobile))`)},
		Props:   []interface{}{PropMapper{Prop: PropType, Value: DeviceTablet}},
	},
	{
		Regexes: []*Regexp2{MustCompile(`(phone|mobile(?:[;\/]| [ \w\/\.]*safari)|pda(?=.+windows ce))`)},
		Props:   []interface{}{PropMapper{Prop: PropType, Value: DeviceMobile}},
	},
	// Generic Android Device
	{
		Regexes: []*Regexp2{MustCompile(`droid .+?; ([\w\. -]+)( bui|\))`)},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: "Generic"},
		},
	},
}
