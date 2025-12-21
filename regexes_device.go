package uaparser

import "regexp"

var deviceRegexes = []RegexRule{
	// Samsung Tablets
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)\b(sch-i[89]0\d|shw-m380s|sm-[ptx]\w{2,4}|gt-[pn]\d{2,4}|sgh-t8[56]9|nexus 10)`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorSamsung},
			PropMapper{Prop: PropType, Value: DeviceTablet},
		},
	},
	// Samsung Mobiles
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)\b((?:s[cgp]h|gt|sm)-\w+|sc[g-]?[\d]+a?|galaxy nexus)`),
			regexp.MustCompile(`(?i)samsung[- ]([-\w]+)`),
			regexp.MustCompile(`(?i)sec-(sgh\w+)`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorSamsung},
			PropMapper{Prop: PropType, Value: DeviceMobile},
		},
	},
	// Apple iPod/iPhone
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)(?:\/|\()(ip(?:hone|od)[\w, ]*)[\/\);]`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorApple},
			PropMapper{Prop: PropType, Value: DeviceMobile},
		},
	},
	// Apple iPad
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)\b(?:ios|apple\w+)\/.+[\(\/](ipad)`),
			regexp.MustCompile(`(?i)\b(ipad)[\d,]*[;\] ].+(mac |i(pad)?)os`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorApple},
			PropMapper{Prop: PropType, Value: DeviceTablet},
		},
	},
	// Apple Macintosh
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)(macintosh);`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorApple},
		},
	},
	// Sharp
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)\b(sh-?[altvz]?\d\d[a-ekm]?)`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorSharp},
			PropMapper{Prop: PropType, Value: DeviceMobile},
		},
	},
	// Honor Tablets
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)\b((?:brt|eln|hey2?|gdi|jdn)-a?[lnw]09|(?:ag[rm]3?|jdn2|kob2)-a?[lw]0[09]hn)(?: bui|\)|;)`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorHonor},
			PropMapper{Prop: PropType, Value: DeviceTablet},
		},
	},
	// Honor Mobiles
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)honor([-\w ]+)[;\)]`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorHonor},
			PropMapper{Prop: PropType, Value: DeviceMobile},
		},
	},
	// Huawei Tablets
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)\b((?:ag[rs][2356]?k?|bah[234]?|bg[2o]|bt[kv]|cmr|cpn|db[ry]2?|jdn2|got|kob2?k?|mon|pce|scm|sht?|[tw]gr|vrd)-[ad]?[lw][0125][09]b?|605hw|bg2-u03|(?:gem|fdr|m2|ple|t1)-[7a]0[1-4][lu]|t1-a2[13][lw]|mediapad[\w\. ]*)\b`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorHuawei},
			PropMapper{Prop: PropType, Value: DeviceTablet},
		},
	},
	// Huawei Mobiles
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)(?:huawei) ?([-\w ]+)[;\)]`),
			regexp.MustCompile(`(?i)\b(nexus 6p|\w{2,4}e?-[atu]?[ln][\dx][\dc][adnt]?)\b`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorHuawei},
			PropMapper{Prop: PropType, Value: DeviceMobile},
		},
	},
	// Xiaomi Tablets
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)oid[^\)]+; (2[\dbc]{4}(182|283|rp\w{2})[cgl]|m2105k81a?c)(?: bui|\))`),
			regexp.MustCompile(`(?i)\b(?:xiao)?((?:red)?mi[-_ ]?pad[\w- ]*)(?: bui|\))`),
		},
		Props: []interface{}{
			PropMapper{Prop: PropModel, Pattern: regexp.MustCompile(`_`), Replace: " "},
			PropMapper{Prop: PropVendor, Value: VendorXiaomi},
			PropMapper{Prop: PropType, Value: DeviceTablet},
		},
	},
	// Xiaomi Mobiles
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)\b; (hm[\-_ ]?\w+) build/`),
			regexp.MustCompile(`(?i)\b(hm[-_ ]?note?[_ ]?(?:\d\w)?) bui`),
			regexp.MustCompile(`(?i)oid[^\)]+; (redmi[\-_ ]?(?:note|k)?[\w_ ]+|m?[12]\d[01]\d\w{3,6}|poco[\w ]+|(shark )?\w{3}-[ah]0|qin ?[1-3](s\+|ultra| pro)?)( bui|; wv|\))`),
			regexp.MustCompile(`(?i)\b(mi[-_ ]?(?:a\d|one|one[_ ]plus|note|max|cc)?[_ ]?(?:\d{0,2}\w?)[_ ]?(?:plus|se|lite|pro)?( 5g|lte)?)(?: bui|\))`),
			regexp.MustCompile(`(?i) ([\w ]+) miui\/v?\d`),
		},
		Props: []interface{}{
			PropMapper{Prop: PropModel, Pattern: regexp.MustCompile(`_`), Replace: " "},
			PropMapper{Prop: PropVendor, Value: VendorXiaomi},
			PropMapper{Prop: PropType, Value: DeviceMobile},
		},
	},
	// OnePlus
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)droid.+; (cph2[3-6]\d[13579]|((gm|hd)19|(ac|be|in|kb)20|(d[en]|eb|le|mt)21|ne22)[0-2]\d|p[g-k]\w[1m]10)\b`),
			regexp.MustCompile(`(?i)(?:one)?(?:plus)? (a\d0\d\d)(?: b|\))`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorOnePlus},
			PropMapper{Prop: PropType, Value: DeviceMobile},
		},
	},
	// OPPO
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i); (\w+) bui.+ oppo`),
			regexp.MustCompile(`(?i)\b(cph[12]\d{3}|p(?:af|c[al]|d\w|e[ar])[mt]\d0|x9007|a101op)\b`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorOppo},
			PropMapper{Prop: PropType, Value: DeviceMobile},
		},
	},
	// OPPO Tablets
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)\b(opd2(\d{3}a?))(?: bui|\))`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorOppo},
			PropMapper{Prop: PropType, Value: DeviceTablet},
		},
	},
	// Vivo
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i); vivo (\w+)(?: bui|\))`),
			regexp.MustCompile(`(?i)\b(v[12]\d{3}\w?[at])(?: bui|;)`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: "Vivo"},
			PropMapper{Prop: PropType, Value: DeviceMobile},
		},
	},
	// Realme
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)\b(rmx[1-3]\d{3})(?: bui|;|\))`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: "Realme"},
			PropMapper{Prop: PropType, Value: DeviceMobile},
		},
	},
	// Lenovo Tablets
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)(ideatab[-\w ]+|602lv|d-42a|a101lv|a2109a|a3500-hv|s[56]000|pb-6505[my]|tb-?x?\d{3,4}(?:f[cu]|xu|[av])|yt\d?-[jx]?\d+[lfmx])( bui|;|\)|\/)`),
			regexp.MustCompile(`(?i)lenovo ?(b[68]0[08]0-?[hf]?|tab(?:[\w- ]+?)|tb[\w-]{6,7})( bui|;|\)|\/)`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorLenovo},
			PropMapper{Prop: PropType, Value: DeviceTablet},
		},
	},
	// Lenovo Mobiles
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)lenovo[-_ ]?([-\w ]+?)(?: bui|\)|\/)`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorLenovo},
			PropMapper{Prop: PropType, Value: DeviceMobile},
		},
	},
	// Motorola
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)\b(milestone|droid(?:[2-4x]| (?:bionic|x2|pro|razr))?:?( 4g)?)\b[\w ]+build\/`),
			regexp.MustCompile(`(?i)\bmot(?:orola)?[- ]([\w\s]+)(\)| bui)`),
			regexp.MustCompile(`(?i)(moto[-\w\(\) ]+|xt\d{3,4}[cgkosw\+]?[-\d]*)`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorMotorola},
			PropMapper{Prop: PropType, Value: DeviceMobile},
		},
	},
	// Motorola Tablets
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)\b(mz60\d|xoom[2 ]{0,2}) build\/`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorMotorola},
			PropMapper{Prop: PropType, Value: DeviceTablet},
		},
	},
	// LG Tablets
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)([vl]k\-?\d{3}) bui| 3\.[-\w; ]{10}lg?-([06cv9]{3,4})`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorLG},
			PropMapper{Prop: PropType, Value: DeviceTablet},
		},
	},
	// LG Mobiles
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)(lm(?:-?f100[nv]?|-[\w\.]+)|nexus [45])`),
			regexp.MustCompile(`(?i)\blg[-e;\/ ]+(\w+)`),
			regexp.MustCompile(`(?i)\blg-?([\d\w]+) bui`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorLG},
			PropMapper{Prop: PropType, Value: DeviceMobile},
		},
	},
	// Nokia Tablets
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)(nokia) (t[12][01])`),
		},
		Props: []interface{}{
			PropVendor,
			PropModel,
			PropMapper{Prop: PropType, Value: DeviceTablet},
		},
	},
	// Nokia Mobiles
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)(?:maemo|nokia).*(n900|lumia \d+|rm-\d+)`),
			regexp.MustCompile(`(?i)nokia[-_ ]?(([-\w\. ]*?))( bui|\)|;|\/)`),
		},
		Props: []interface{}{
			PropMapper{Prop: PropModel, Pattern: regexp.MustCompile(`_`), Replace: " "},
			PropMapper{Prop: PropType, Value: DeviceMobile},
			PropMapper{Prop: PropVendor, Value: "Nokia"},
		},
	},
	// Google Pixel Tablet
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)(pixel (c|tablet))\b`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorGoogle},
			PropMapper{Prop: PropType, Value: DeviceTablet},
		},
	},
	// Google Pixel
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)droid.+;(?: google)? (g(01[13]a|020[aem]|025[jn]|1b60|1f8f|2ybb|4s1m|576d|5nz6|8hhn|8vou|a02099|c15s|d1yq|e2ae|ec77|gh2x|kv4x|p4bc|pj41|r83y|tt9q|ur25|wvk6)|pixel[\d ]*a?( pro)?( xl)?( fold)?( \(5g\))?)( bui|\))`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorGoogle},
			PropMapper{Prop: PropType, Value: DeviceMobile},
		},
	},
	// Google Pixelbook
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)(google) (pixelbook( go)?)`),
		},
		Props: []interface{}{PropVendor, PropModel},
	},
	// Sony Mobile
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)droid.+; (a?\d[0-2]{2}so|[c-g]\d{4}|so[-gl]\w+|xq-\w\w\d\d)`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorSony},
			PropMapper{Prop: PropType, Value: DeviceMobile},
		},
	},
	// Sony Tablets
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)sony tablet [ps]`),
			regexp.MustCompile(`(?i)\b(?:sony)?sgp\w+(?: bui|\))`),
		},
		Props: []interface{}{
			PropMapper{Prop: PropModel, Value: "Xperia Tablet"},
			PropMapper{Prop: PropVendor, Value: VendorSony},
			PropMapper{Prop: PropType, Value: DeviceTablet},
		},
	},
	// Amazon Tablets
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)(alexa)webm`),
			regexp.MustCompile(`(?i)(kf[a-z]{2}wi|aeo\w\w)( bui|\))`),
			regexp.MustCompile(`(?i)(kf[a-z]+)( bui|\)).+silk\/`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorAmazon},
			PropMapper{Prop: PropType, Value: DeviceTablet},
		},
	},
	// Amazon Fire Phone
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)((?:sd|kf)[0349hijorstuw]+)( bui|\)).+silk\/`),
		},
		Props: []interface{}{
			PropMapper{Prop: PropModel, Fn: func(m string) string { return "Fire Phone " + m }},
			PropMapper{Prop: PropVendor, Value: VendorAmazon},
			PropMapper{Prop: PropType, Value: DeviceMobile},
		},
	},
	// BlackBerry PlayBook
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)(playbook);[-\w\),; ]+(rim)`),
		},
		Props: []interface{}{
			PropModel,
			PropVendor,
			PropMapper{Prop: PropType, Value: DeviceTablet},
		},
	},
	// BlackBerry
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)\b((?:bb[a-f]|st[hv])100-\d)`),
			regexp.MustCompile(`(?i)(?:blackberry|\(bb10;) (\w+)`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorBlackBerry},
			PropMapper{Prop: PropType, Value: DeviceMobile},
		},
	},
	// Asus Tablets
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)(?:\b|asus_)(transfo[prime ]{4,10} \w+|eeepc|slider \w+|nexus 7|padfone|p00[cj])`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorAsus},
			PropMapper{Prop: PropType, Value: DeviceTablet},
		},
	},
	// Asus Mobiles
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i) (z[bes]6[027][012][km][ls]|zenfone \d\w?)\b`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorAsus},
			PropMapper{Prop: PropType, Value: DeviceMobile},
		},
	},
	// HTC Nexus 9
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)(nexus 9)`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: "HTC"},
			PropMapper{Prop: PropType, Value: DeviceTablet},
		},
	},
	// HTC, ZTE, Alcatel, etc.
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)(htc)[-;_ ]{1,2}([\w ]+)`),
			regexp.MustCompile(`(?i)(zte)[- ]([\w ]+?)(?: bui|\/|\))`),
			regexp.MustCompile(`(?i)(alcatel|geeksphone|nexian|panasonic|sony)[-_ ]?([-\w]*)`),
		},
		Props: []interface{}{
			PropVendor,
			PropMapper{Prop: PropModel, Pattern: regexp.MustCompile(`_`), Replace: " "},
			PropMapper{Prop: PropType, Value: DeviceMobile},
		},
	},
	// Smart TVs
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)(philips)[\w ]+tv`),
			regexp.MustCompile(`(?i)smart-tv.+(samsung)`),
		},
		Props: []interface{}{
			PropVendor,
			PropMapper{Prop: PropType, Value: DeviceSmartTV},
		},
	},
	// Samsung SmartTV
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)hbbtv.+maple;(\d+)`),
		},
		Props: []interface{}{
			PropMapper{Prop: PropModel, Fn: func(m string) string { return "SmartTV" + m }},
			PropMapper{Prop: PropVendor, Value: VendorSamsung},
			PropMapper{Prop: PropType, Value: DeviceSmartTV},
		},
	},
	// Apple TV
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)(apple) ?tv`),
		},
		Props: []interface{}{
			PropVendor,
			PropMapper{Prop: PropModel, Value: VendorApple + " TV"},
			PropMapper{Prop: PropType, Value: DeviceSmartTV},
		},
	},
	// Google Chromecast
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)crkey.*devicetype\/chromecast`),
		},
		Props: []interface{}{
			PropMapper{Prop: PropModel, Value: BrowserChromecast + " Third Generation"},
			PropMapper{Prop: PropVendor, Value: VendorGoogle},
			PropMapper{Prop: PropType, Value: DeviceSmartTV},
		},
	},
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)crkey.*devicetype\/([^/]*)`),
		},
		Props: []interface{}{
			PropMapper{Prop: PropModel, Fn: func(m string) string { return "Chromecast " + m }},
			PropMapper{Prop: PropVendor, Value: VendorGoogle},
			PropMapper{Prop: PropType, Value: DeviceSmartTV},
		},
	},
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)fuchsia.*crkey`),
		},
		Props: []interface{}{
			PropMapper{Prop: PropModel, Value: BrowserChromecast + " Nest Hub"},
			PropMapper{Prop: PropVendor, Value: VendorGoogle},
			PropMapper{Prop: PropType, Value: DeviceSmartTV},
		},
	},
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)crkey`),
		},
		Props: []interface{}{
			PropMapper{Prop: PropModel, Value: BrowserChromecast},
			PropMapper{Prop: PropVendor, Value: VendorGoogle},
			PropMapper{Prop: PropType, Value: DeviceSmartTV},
		},
	},
	// Amazon Fire TV
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)droid.+aft(\w+)( bui|\))`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorAmazon},
			PropMapper{Prop: PropType, Value: DeviceSmartTV},
		},
	},
	// Consoles
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)(playstation \w+)`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorSony},
			PropMapper{Prop: PropType, Value: DeviceConsole},
		},
	},
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)\b(xbox one)\b`),
			regexp.MustCompile(`(?i)\b(xbox)\b`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorMicrosoft},
			PropMapper{Prop: PropType, Value: DeviceConsole},
		},
	},
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)(ouya)`),
			regexp.MustCompile(`(?i)(nintendo) (\w+)`),
		},
		Props: []interface{}{PropVendor, PropModel, PropMapper{Prop: PropType, Value: DeviceConsole}},
	},
	// Wearables
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)\b(sm-[lr]\d\d[0156][fnuw]?s?|gear live)\b`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorSamsung},
			PropMapper{Prop: PropType, Value: DeviceWearable},
		},
	},
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)((pebble))app`),
			regexp.MustCompile(`(?i)(asus|google|lg|oppo|xiaomi) ((pixel |zen)?watch[\w ]*)( bui|\))`),
		},
		Props: []interface{}{
			PropVendor,
			PropModel,
			PropMapper{Prop: PropType, Value: DeviceWearable},
		},
	},
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)(watch)(?: ?os[,\/]|\d,\d\/)[\d\.]+`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorApple},
			PropMapper{Prop: PropType, Value: DeviceWearable},
		},
	},
	// XR
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)droid.+; (glass) \d`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorGoogle},
			PropMapper{Prop: PropType, Value: DeviceXR},
		},
	},
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)(pico) ([\w ]+) os\d`),
		},
		Props: []interface{}{
			PropVendor,
			PropModel,
			PropMapper{Prop: PropType, Value: DeviceXR},
		},
	},
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)(quest( \d| pro)?s?).+vr`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: BrowserFacebook},
			PropMapper{Prop: PropType, Value: DeviceXR},
		},
	},
	// Embedded
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)(tesla)(?: qtcarbrowser|\/[-\w\.]+)`),
		},
		Props: []interface{}{
			PropVendor,
			PropMapper{Prop: PropType, Value: DeviceEmbedded},
		},
	},
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)(aeobc)\b`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorAmazon},
			PropMapper{Prop: PropType, Value: DeviceEmbedded},
		},
	},
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)(homepod).+mac os`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: VendorApple},
			PropMapper{Prop: PropType, Value: DeviceEmbedded},
		},
	},
	// Generic Android device
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)droid .+?; ([\w\. -]+)( bui|\))`),
		},
		Props: []interface{}{
			PropModel,
			PropMapper{Prop: PropVendor, Value: "Generic"},
		},
	},
	// Unidentifiable Tablet
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)\b(tablet|tab)[;\/]`),
		},
		Props: []interface{}{
			PropMapper{Prop: PropType, Value: DeviceTablet},
		},
	},
	// Unidentifiable Mobile
	{
		Regexes: []*regexp.Regexp{
			regexp.MustCompile(`(?i)(phone|mobile(?:[;\/]| [ \w\/\.]*safari))`),
		},
		Props: []interface{}{
			PropMapper{Prop: PropType, Value: DeviceMobile},
		},
	},
}
