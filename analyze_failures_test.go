package uaparser

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// 分析失败用例 - 浏览器
func TestAnalyzeBrowserFailures(t *testing.T) {
	testFile := "/Users/rowei/Develop/js/ua-parser-js/test/data/ua/browser/browser-all.json"
	data, _ := os.ReadFile(testFile)
	var testCases []BrowserTestCase
	json.Unmarshal(data, &testCases)

	var failures []string
	for i, tc := range testCases {
		browser := New(tc.UA).GetBrowser()

		expectName := tc.Expect.Name
		expectVersion := tc.Expect.Version
		expectMajor := tc.Expect.Major
		if expectName == "undefined" {
			expectName = ""
		}
		if expectVersion == "undefined" {
			expectVersion = ""
		}
		if expectMajor == "undefined" {
			expectMajor = ""
		}

		nameMatch := browser.Name == expectName
		versionMatch := expectVersion == "" || browser.Version == expectVersion ||
			strings.HasPrefix(browser.Version, expectVersion)
		majorMatch := expectMajor == "" || browser.Major == expectMajor

		if !nameMatch || !versionMatch || !majorMatch {
			failures = append(failures, fmt.Sprintf("[%d] %s\n  UA: %s\n  期望: name=%q version=%q major=%q\n  实际: name=%q version=%q major=%q",
				i+1, tc.Desc, tc.UA,
				expectName, expectVersion, expectMajor,
				browser.Name, browser.Version, browser.Major))
		}
	}

	t.Logf("浏览器失败用例数: %d", len(failures))
	if len(failures) > 0 && len(failures) <= 30 {
		t.Logf("失败详情:\n%s", strings.Join(failures, "\n\n"))
	} else if len(failures) > 30 {
		t.Logf("失败详情（前30个）:\n%s", strings.Join(failures[:30], "\n\n"))
	}
}

// 分析失败用例 - CPU
func TestAnalyzeCPUFailures(t *testing.T) {
	testFile := "/Users/rowei/Develop/js/ua-parser-js/test/data/ua/cpu/cpu-all.json"
	data, _ := os.ReadFile(testFile)
	var testCases []CPUTestCase
	json.Unmarshal(data, &testCases)

	var failures []string
	for i, tc := range testCases {
		cpu := New(tc.UA).GetCPU()
		expectArch := tc.Expect.Architecture
		if expectArch == "undefined" {
			expectArch = ""
		}

		if cpu.Architecture != expectArch {
			failures = append(failures, fmt.Sprintf("[%d] %s\n  UA: %s\n  期望: arch=%q\n  实际: arch=%q",
				i+1, tc.Desc, tc.UA,
				expectArch,
				cpu.Architecture))
		}
	}

	t.Logf("CPU失败用例数: %d", len(failures))
	if len(failures) > 0 {
		t.Logf("失败详情:\n%s", strings.Join(failures, "\n\n"))
	}
}

// 分析失败用例 - 引擎
func TestAnalyzeEngineFailures(t *testing.T) {
	testFile := "/Users/rowei/Develop/js/ua-parser-js/test/data/ua/engine/engine-all.json"
	data, _ := os.ReadFile(testFile)
	var testCases []EngineTestCase
	json.Unmarshal(data, &testCases)

	var failures []string
	for i, tc := range testCases {
		engine := New(tc.UA).GetEngine()

		expectName := tc.Expect.Name
		expectVersion := tc.Expect.Version
		if expectName == "undefined" {
			expectName = ""
		}
		if expectVersion == "undefined" {
			expectVersion = ""
		}

		nameMatch := engine.Name == expectName
		versionMatch := expectVersion == "" || engine.Version == expectVersion ||
			strings.HasPrefix(engine.Version, expectVersion)

		if !nameMatch || !versionMatch {
			failures = append(failures, fmt.Sprintf("[%d] %s\n  UA: %s\n  期望: name=%q version=%q\n  实际: name=%q version=%q",
				i+1, tc.Desc, tc.UA,
				expectName, expectVersion,
				engine.Name, engine.Version))
		}
	}

	t.Logf("引擎失败用例数: %d", len(failures))
	if len(failures) > 0 {
		t.Logf("失败详情:\n%s", strings.Join(failures, "\n\n"))
	}
}

// 分析失败用例 - OS
func TestAnalyzeOSFailures(t *testing.T) {
	osDir := "/Users/rowei/Develop/js/ua-parser-js/test/data/ua/os"
	files, _ := os.ReadDir(osDir)

	var allFailures []string
	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".json") {
			continue
		}

		data, err := os.ReadFile(filepath.Join(osDir, file.Name()))
		if err != nil {
			continue
		}

		var testCases []OSTestCase
		if err := json.Unmarshal(data, &testCases); err != nil {
			continue
		}

		for i, tc := range testCases {
			os := New(tc.UA).GetOS()

			expectName := tc.Expect.Name
			expectVersion := tc.Expect.Version
			if expectName == "undefined" {
				expectName = ""
			}
			if expectVersion == "undefined" {
				expectVersion = ""
			}

			nameMatch := os.Name == expectName
			versionMatch := expectVersion == "" || os.Version == expectVersion ||
				strings.HasPrefix(os.Version, strings.Split(expectVersion, ".")[0])

			if !nameMatch || !versionMatch {
				allFailures = append(allFailures, fmt.Sprintf("[%s:%d] %s\n  UA: %s\n  期望: name=%q version=%q\n  实际: name=%q version=%q",
					file.Name(), i+1, tc.Desc, tc.UA,
					expectName, expectVersion,
					os.Name, os.Version))
			}
		}
	}

	t.Logf("OS失败用例数: %d", len(allFailures))
	if len(allFailures) > 0 && len(allFailures) <= 30 {
		t.Logf("失败详情:\n%s", strings.Join(allFailures, "\n\n"))
	} else if len(allFailures) > 30 {
		t.Logf("失败详情（前30个）:\n%s", strings.Join(allFailures[:30], "\n\n"))
	}
}

// 分析失败用例 - 设备
func TestAnalyzeDeviceFailures(t *testing.T) {
	deviceDir := "/Users/rowei/Develop/js/ua-parser-js/test/data/ua/device"
	files, _ := os.ReadDir(deviceDir)

	failuresByFile := make(map[string][]string)
	totalFailures := 0

	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".json") {
			continue
		}

		data, err := os.ReadFile(filepath.Join(deviceDir, file.Name()))
		if err != nil {
			continue
		}

		var testCases []DeviceTestCase
		if err := json.Unmarshal(data, &testCases); err != nil {
			continue
		}

		fileFailures := 0
		for i, tc := range testCases {
			device := New(tc.UA).GetDevice()

			expectVendor := tc.Expect.Vendor
			expectModel := tc.Expect.Model
			expectType := tc.Expect.Type
			if expectVendor == "undefined" {
				expectVendor = ""
			}
			if expectModel == "undefined" {
				expectModel = ""
			}
			if expectType == "undefined" {
				expectType = ""
			}

			vendorMatch := expectVendor == "" || device.Vendor == expectVendor
			modelMatch := expectModel == "" || device.Model == expectModel
			typeMatch := expectType == "" || device.Type == expectType

			if !vendorMatch || !modelMatch || !typeMatch {
				fileFailures++
				totalFailures++
				if len(failuresByFile[file.Name()]) < 3 {
					failuresByFile[file.Name()] = append(failuresByFile[file.Name()],
						fmt.Sprintf("  [%d] %s\n    UA: %s\n    期望: vendor=%q model=%q type=%q\n    实际: vendor=%q model=%q type=%q",
							i+1, tc.Desc, tc.UA[:min(80, len(tc.UA))],
							expectVendor, expectModel, expectType,
							device.Vendor, device.Model, device.Type))
				}
			}
		}

		if fileFailures > 0 {
			summary := fmt.Sprintf("%s: %d个失败", file.Name(), fileFailures)
			failuresByFile[file.Name()] = append([]string{summary}, failuresByFile[file.Name()]...)
		}
	}

	t.Logf("设备失败用例总数: %d", totalFailures)
	t.Logf("按文件统计失败情况（前3个示例）:")
	for _, failures := range failuresByFile {
		if len(failures) > 0 {
			t.Logf("\n%s", strings.Join(failures, "\n"))
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
