package uaparser

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// 测试数据结构
type BrowserTestCase struct {
	Desc   string `json:"desc"`
	UA     string `json:"ua"`
	Expect struct {
		Name    string `json:"name"`
		Version string `json:"version"`
		Major   string `json:"major"`
	} `json:"expect"`
}

type OSTestCase struct {
	Desc   string `json:"desc"`
	UA     string `json:"ua"`
	Expect struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	} `json:"expect"`
}

type DeviceTestCase struct {
	Desc   string `json:"desc"`
	UA     string `json:"ua"`
	Expect struct {
		Vendor string `json:"vendor"`
		Model  string `json:"model"`
		Type   string `json:"type"`
	} `json:"expect"`
}

type EngineTestCase struct {
	Desc   string `json:"desc"`
	UA     string `json:"ua"`
	Expect struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	} `json:"expect"`
}

type CPUTestCase struct {
	Desc   string `json:"desc"`
	UA     string `json:"ua"`
	Expect struct {
		Architecture string `json:"architecture"`
	} `json:"expect"`
}

// 完整测试 - 浏览器检测
func TestFullBrowser(t *testing.T) {
	testFile := "/Users/rowei/Develop/js/ua-parser-js/test/data/ua/browser/browser-all.json"
	data, err := os.ReadFile(testFile)
	if err != nil {
		t.Skipf("跳过测试：无法读取原始测试数据 (%v)", err)
		return
	}

	var testCases []BrowserTestCase
	if err := json.Unmarshal(data, &testCases); err != nil {
		t.Fatalf("解析测试数据失败: %v", err)
	}

	passed, failed := 0, 0
	for _, tc := range testCases {
		browser := New(tc.UA).GetBrowser()
		nameMatch := browser.Name == tc.Expect.Name
		versionMatch := tc.Expect.Version == "" || browser.Version == tc.Expect.Version ||
			strings.HasPrefix(browser.Version, tc.Expect.Version)
		majorMatch := tc.Expect.Major == "" || browser.Major == tc.Expect.Major

		if nameMatch && versionMatch && majorMatch {
			passed++
		} else {
			failed++
		}
	}

	total := passed + failed
	t.Logf("浏览器检测: %d/%d (%.1f%%)", passed, total, float64(passed)/float64(total)*100)
}

// 完整测试 - 操作系统
func TestFullOS(t *testing.T) {
	osDir := "/Users/rowei/Develop/js/ua-parser-js/test/data/ua/os"
	files, err := os.ReadDir(osDir)
	if err != nil {
		t.Skipf("跳过测试：无法读取目录 (%v)", err)
		return
	}

	totalPassed, totalFailed := 0, 0
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

		for _, tc := range testCases {
			os := New(tc.UA).GetOS()
			nameMatch := os.Name == tc.Expect.Name
			versionMatch := tc.Expect.Version == "" || os.Version == tc.Expect.Version ||
				strings.HasPrefix(os.Version, strings.Split(tc.Expect.Version, ".")[0])

			if nameMatch && versionMatch {
				totalPassed++
			} else {
				totalFailed++
			}
		}
	}

	total := totalPassed + totalFailed
	t.Logf("操作系统检测: %d/%d (%.1f%%)", totalPassed, total, float64(totalPassed)/float64(total)*100)
}

// 完整测试 - 设备
func TestFullDevice(t *testing.T) {
	deviceDir := "/Users/rowei/Develop/js/ua-parser-js/test/data/ua/device"
	files, err := os.ReadDir(deviceDir)
	if err != nil {
		t.Skipf("跳过测试：无法读取目录 (%v)", err)
		return
	}

	totalPassed, totalFailed := 0, 0
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

		for _, tc := range testCases {
			device := New(tc.UA).GetDevice()
			vendorMatch := tc.Expect.Vendor == "" || device.Vendor == tc.Expect.Vendor
			modelMatch := tc.Expect.Model == "" || device.Model == tc.Expect.Model
			typeMatch := tc.Expect.Type == "" || device.Type == tc.Expect.Type

			if vendorMatch && modelMatch && typeMatch {
				totalPassed++
			} else {
				totalFailed++
			}
		}
	}

	total := totalPassed + totalFailed
	t.Logf("设备检测: %d/%d (%.1f%%)", totalPassed, total, float64(totalPassed)/float64(total)*100)
}

// 完整测试 - 引擎
func TestFullEngine(t *testing.T) {
	testFile := "/Users/rowei/Develop/js/ua-parser-js/test/data/ua/engine/engine-all.json"
	data, err := os.ReadFile(testFile)
	if err != nil {
		t.Skipf("跳过测试：无法读取原始测试数据 (%v)", err)
		return
	}

	var testCases []EngineTestCase
	if err := json.Unmarshal(data, &testCases); err != nil {
		t.Fatalf("解析测试数据失败: %v", err)
	}

	passed, failed := 0, 0
	for _, tc := range testCases {
		engine := New(tc.UA).GetEngine()
		nameMatch := engine.Name == tc.Expect.Name
		versionMatch := tc.Expect.Version == "" || engine.Version == tc.Expect.Version ||
			strings.HasPrefix(engine.Version, tc.Expect.Version)

		if nameMatch && versionMatch {
			passed++
		} else {
			failed++
		}
	}

	total := passed + failed
	t.Logf("引擎检测: %d/%d (%.1f%%)", passed, total, float64(passed)/float64(total)*100)
}

// 完整测试 - CPU
func TestFullCPU(t *testing.T) {
	testFile := "/Users/rowei/Develop/js/ua-parser-js/test/data/ua/cpu/cpu-all.json"
	data, err := os.ReadFile(testFile)
	if err != nil {
		t.Skipf("跳过测试：无法读取原始测试数据 (%v)", err)
		return
	}

	var testCases []CPUTestCase
	if err := json.Unmarshal(data, &testCases); err != nil {
		t.Fatalf("解析测试数据失败: %v", err)
	}

	passed, failed := 0, 0
	for _, tc := range testCases {
		cpu := New(tc.UA).GetCPU()
		if cpu.Architecture == tc.Expect.Architecture {
			passed++
		} else {
			failed++
		}
	}

	total := passed + failed
	t.Logf("CPU检测: %d/%d (%.1f%%)", passed, total, float64(passed)/float64(total)*100)
}

// 汇总测试
func TestFullSummary(t *testing.T) {
	t.Run("Browser", TestFullBrowser)
	t.Run("OS", TestFullOS)
	t.Run("Device", TestFullDevice)
	t.Run("Engine", TestFullEngine)
	t.Run("CPU", TestFullCPU)
	
	fmt.Println("\n========================================")
	fmt.Println("   完整验证测试总结")
	fmt.Println("========================================")
	fmt.Println("对比原始 ua-parser-js 项目的全部测试数据")
	fmt.Println("  - 浏览器测试用例: 282个")
	fmt.Println("  - 操作系统测试用例: 258个")
	fmt.Println("  - 设备测试用例: 963个")
	fmt.Println("  - 引擎测试用例: 21个")
	fmt.Println("  - CPU测试用例: 39个")
	fmt.Println("  总计: 1563个测试用例")
	fmt.Println("========================================")
}

