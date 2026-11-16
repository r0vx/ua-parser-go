package uaparser

import (
	"reflect"
	"regexp"
	"strings"

	"github.com/dlclark/regexp2"
)

// Regexp2 is a type alias for regexp2.Regexp
type Regexp2 = regexp2.Regexp

// PropMapper defines how to map a regex match to a property
type PropMapper struct {
	// Property name to set
	Prop string
	// Static value to use (if set, ignores match)
	Value string
	// Function to transform the match
	Fn func(match string) string
	// Regex pattern for replacement (standard regexp for simple replacements)
	Pattern *regexp.Regexp
	// Replacement string
	Replace string
	// String mapper for lookup
	StrMap map[string]interface{}
}

// RegexRule defines a regex rule with property mappings
type RegexRule struct {
	Regexes []*regexp2.Regexp
	Props   []interface{} // can be string or PropMapper
}

// MustCompile compiles a regexp2 pattern with case-insensitive flag
func MustCompile(pattern string) *regexp2.Regexp {
	re, err := regexp2.Compile(pattern, regexp2.IgnoreCase)
	if err != nil {
		panic("regexp2: Compile(" + pattern + "): " + err.Error())
	}
	return re
}

// rgxMapper applies regex rules to extract values
func rgxMapper(result interface{}, ua string, rules []RegexRule) {
	if ua == "" || len(rules) == 0 {
		return
	}

	rv := reflect.ValueOf(result)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return
	}
	rv = rv.Elem()

	for _, rule := range rules {
		for _, regex := range rule.Regexes {
			if regex == nil {
				continue
			}
			match, err := regex.FindStringMatch(ua)
			if err != nil || match == nil {
				continue
			}

			// Get all groups
			groups := match.Groups()

			// Apply property mappings
			matchIdx := 0
			for _, prop := range rule.Props {
				matchIdx++
				var matchStr string
				if matchIdx < len(groups) {
					matchStr = groups[matchIdx].String()
				}

				switch p := prop.(type) {
				case string:
					// Simple property name
					setField(rv, p, matchStr)
				case PropMapper:
					// Complex property mapping
					applyPropMapper(rv, p, matchStr)
				case []interface{}:
					// Array format: [propName, value] or [propName, pattern, replace] etc.
					applyArrayMapper(rv, p, matchStr)
				}
			}
			return // Found a match, stop processing
		}
	}
}

func setField(rv reflect.Value, name string, value string) {
	// Map property names to struct field names
	fieldName := propToFieldName(name)
	field := rv.FieldByName(fieldName)
	if field.IsValid() && field.CanSet() && field.Kind() == reflect.String {
		field.SetString(value)
	}
}

func propToFieldName(prop string) string {
	switch prop {
	case PropName:
		return "Name"
	case PropVersion:
		return "Version"
	case PropMajor:
		return "Major"
	case PropType:
		return "Type"
	case PropVendor:
		return "Vendor"
	case PropModel:
		return "Model"
	case PropArchitecture:
		return "Architecture"
	default:
		// Capitalize first letter
		if len(prop) > 0 {
			return strings.ToUpper(prop[:1]) + prop[1:]
		}
		return prop
	}
}

func applyPropMapper(rv reflect.Value, pm PropMapper, match string) {
	value := match

	if pm.Value != "" {
		value = pm.Value
	} else {
		if pm.Fn != nil {
			value = pm.Fn(value)
		}
		if pm.Pattern != nil && value != "" {
			value = pm.Pattern.ReplaceAllString(value, pm.Replace)
		}
		if pm.StrMap != nil && value != "" {
			value = strMapper(value, pm.StrMap)
		}
	}

	setField(rv, pm.Prop, value)
}

func applyArrayMapper(rv reflect.Value, arr []interface{}, match string) {
	if len(arr) < 2 {
		return
	}

	propName, ok := arr[0].(string)
	if !ok {
		return
	}

	var value string

	switch v := arr[1].(type) {
	case string:
		// Static value
		value = v
	case func(string) string:
		// Transform function
		if match != "" {
			value = v(match)
		}
	case *regexp.Regexp:
		// Pattern replacement: [propName, pattern, replace]
		if len(arr) >= 3 && match != "" {
			replace, _ := arr[2].(string)
			value = v.ReplaceAllString(match, replace)
		}
	case map[string]interface{}:
		// String map lookup
		if match != "" {
			value = strMapper(match, v)
		}
	default:
		value = match
	}

	setField(rv, propName, value)
}

// strMapper maps a string value using a lookup map
func strMapper(str string, m map[string]interface{}) string {
	if str == "" {
		return ""
	}

	strLower := lowerize(str)

	for key, val := range m {
		switch v := val.(type) {
		case string:
			if lowerize(v) == strLower {
				if key == Unknown {
					return ""
				}
				return key
			}
		case []string:
			for _, s := range v {
				if lowerize(s) == strLower {
					if key == Unknown {
						return ""
					}
					return key
				}
			}
		}
	}

	// Check for wildcard
	if wildcard, ok := m["*"]; ok {
		if wc, ok := wildcard.(string); ok {
			return wc
		}
	}

	return str
}

// Windows version map
var windowsVersionMap = map[string]interface{}{
	"ME":      "4.90",
	"NT 3.51": "3.51",
	"NT 4.0":  "4.0",
	"2000":    []string{"5.0", "5.01"},
	"XP":      []string{"5.1", "5.2"},
	"Vista":   "6.0",
	"7":       "6.1",
	"8":       "6.2",
	"8.1":     "6.3",
	"10":      []string{"6.4", "10.0"},
	"NT":      "",
}

// Browser hints map：把 Sec-CH-UA 中的品牌名映射回标准浏览器名。
var browserHintsMap = map[string]interface{}{
	"Chrome":          "Google Chrome",
	"Edge":            "Microsoft Edge",
	"Edge WebView2":   "Microsoft Edge WebView2",
	"Chrome WebView":  "Android WebView",
	"Chrome Headless": "HeadlessChrome",
	"Huawei Browser":  "HuaweiBrowser",
	"MIUI Browser":    "Miui Browser",
	"Opera Mobi":      "OperaMobile",
	"Yandex":          "YaBrowser",
}

// Form factors map
var formFactorsMap = map[string]interface{}{
	DeviceEmbedded: "Automotive",
	DeviceMobile:   "Mobile",
	DeviceTablet:   []string{"Tablet", "EInk"},
	DeviceSmartTV:  "TV",
	DeviceWearable: "Watch",
	DeviceXR:       []string{"VR", "XR"},
	Unknown:        []string{"Desktop", "Unknown"},
	"*":            "",
}
