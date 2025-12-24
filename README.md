# UA Parser Go

🚀 **ua-parser-js 的 Go 语言完整移植版本**

一个高性能、类型安全的 User-Agent 解析库，用于检测浏览器、引擎、操作系统、CPU 和设备类型。

## ✨ 特性

- 🎯 **100% 准确率**：与原版 ua-parser-js 完全一致
- ⚡ **高性能**：完整解析仅需 ~0.5ms
- 🔒 **类型安全**：利用 Go 的强类型系统
- 🌐 **全面支持**：识别 100+ 种浏览器、设备和操作系统
- 🧪 **测试完善**：通过全部 1563 个原版测试用例

## 📊 准确率

基于原始 ua-parser-js 的测试数据集验证：

| 模块 | 通过 | 总数 | 准确率 |
|------|------|------|--------|
| 浏览器检测 | 282 | 282 | **100%** ✅ |
| 操作系统检测 | 258 | 258 | **100%** ✅ |
| 设备检测 | 963 | 963 | **100%** ✅ |
| 引擎检测 | 21 | 21 | **100%** ✅ |
| CPU架构检测 | 39 | 39 | **100%** ✅ |
| **总计** | **1563** | **1563** | **100%** ✅ |

## 🚀 快速开始

### 安装

```bash
go get github.com/r0vx/ua-parser-go
```

### 基本使用

```go
package main

import (
    "fmt"
    uaparser "github.com/r0vx/ua-parser-go"
)

func main() {
    ua := "Mozilla/5.0 (iPhone; CPU iPhone OS 17_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 Mobile/15E148 Safari/604.1"
    
    // 方式 1: 直接解析所有信息
    result := uaparser.Parse(ua)
    fmt.Printf("浏览器: %s %s\n", result.Browser.Name, result.Browser.Version)
    fmt.Printf("操作系统: %s %s\n", result.OS.Name, result.OS.Version)
    fmt.Printf("设备: %s %s (%s)\n", result.Device.Vendor, result.Device.Model, result.Device.Type)
    
    // 方式 2: 使用 Parser 实例（按需解析）
    parser := uaparser.New(ua)
    browser := parser.GetBrowser()
    os := parser.GetOS()
    device := parser.GetDevice()
    engine := parser.GetEngine()
    cpu := parser.GetCPU()
}
```

### 输出示例

```go
result := uaparser.Parse("Mozilla/5.0 (Linux; Android 14; SM-S918B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.6099.43 Mobile Safari/537.36")

// result.Browser
// {
//   "name": "Mobile Chrome",
//   "version": "120.0.6099.43",
//   "major": "120",
//   "type": ""
// }

// result.OS
// {
//   "name": "Android",
//   "version": "14"
// }

// result.Device
// {
//   "vendor": "Samsung",
//   "model": "SM-S918B",
//   "type": "mobile"
// }

// result.Engine
// {
//   "name": "Blink",
//   "version": "120.0.6099.43"
// }

// result.CPU
// {
//   "architecture": "arm64"
// }
```

## 📖 API 文档

### 主要函数

#### `Parse(ua string) *Result`
一次性解析所有信息，返回完整结果。

#### `New(ua string) *UAParser`
创建解析器实例，支持按需解析。

### Parser 方法

- `GetBrowser() Browser` - 获取浏览器信息
- `GetOS() OS` - 获取操作系统信息
- `GetDevice() Device` - 获取设备信息
- `GetEngine() Engine` - 获取浏览器引擎信息
- `GetCPU() CPU` - 获取 CPU 架构信息
- `GetResult() *Result` - 获取完整解析结果
- `SetUA(ua string) *UAParser` - 设置新的 UA 字符串
- `GetUA() string` - 获取当前 UA 字符串

### 数据结构

```go
type Result struct {
    UA      string  `json:"ua"`      // 原始 User-Agent 字符串
    Browser Browser `json:"browser"`
    CPU     CPU     `json:"cpu"`
    Device  Device  `json:"device"`
    Engine  Engine  `json:"engine"`
    OS      OS      `json:"os"`
}

type Browser struct {
    Name    string `json:"name"`    // 浏览器名称: "Chrome", "Safari", "WeChat" 等
    Version string `json:"version"` // 完整版本: "120.0.6099.43"
    Major   string `json:"major"`   // 主版本号: "120"
    Type    string `json:"type"`    // 类型: "inapp" (应用内浏览器) 或 ""
}

type OS struct {
    Name    string `json:"name"`    // 系统名称: "iOS", "Android", "Windows" 等
    Version string `json:"version"` // 系统版本: "17.0", "14", "10" 等
}

type Device struct {
    Type   string `json:"type"`   // 设备类型 (见下表)
    Model  string `json:"model"`  // 型号: "iPhone", "Galaxy S21" 等
    Vendor string `json:"vendor"` // 厂商: "Apple", "Samsung" 等
}

type Engine struct {
    Name    string `json:"name"`    // 引擎名称: "Blink", "WebKit", "Gecko" 等
    Version string `json:"version"` // 引擎版本
}

type CPU struct {
    Architecture string `json:"architecture"` // CPU 架构: "amd64", "arm64" 等
}
```

### Device.Type 可能值

| 值 | 说明 |
|----|------|
| `"mobile"` | 手机 |
| `"tablet"` | 平板 |
| `"smarttv"` | 智能电视 |
| `"console"` | 游戏机 |
| `"wearable"` | 可穿戴设备 |
| `"xr"` | VR/AR 设备 |
| `"embedded"` | 嵌入式设备 |
| `""` | 桌面/未知 |

### Browser.Type 可能值

| 值 | 说明 |
|----|------|
| `"inapp"` | 应用内浏览器 (微信、支付宝、Facebook 等) |
| `""` | 普通浏览器 |

## 🎯 支持的浏览器

**桌面浏览器**
- Chrome / Chromium / Chrome Headless
- Firefox / LibreWolf
- Safari
- Edge / Edge WebView
- Opera / Opera GX / Opera Touch
- Brave
- Vivaldi
- Yandex
- 360浏览器
- QQ浏览器
- 搜狗浏览器
- IE
- 更多...

**移动浏览器**
- Mobile Chrome / Chrome WebView
- Mobile Safari
- Mobile Firefox
- Samsung Internet
- UCBrowser
- MIUI浏览器
- Huawei浏览器
- Vivo浏览器
- Pico浏览器
- 更多...

**应用内浏览器** (`Browser.Type = "inapp"`)
- WeChat (微信)
- Alipay (支付宝)
- Facebook / Instagram
- TikTok (抖音)
- Twitter
- LINE
- Snapchat
- LinkedIn
- KakaoTalk
- 更多...

## 🔧 技术细节

### 依赖

本项目使用 [dlclark/regexp2](https://github.com/dlclark/regexp2) 库来支持完整的 PCRE 正则语法，确保与原版 JavaScript 正则 100% 兼容：

- ✅ 负向前瞻 `(?!...)`
- ✅ 正向前瞻 `(?=...)`
- ✅ 后向引用 `\1`, `\2`

### 性能基准

```
BenchmarkParse-12         2379    502586 ns/op    1465 B/op    25 allocs/op
BenchmarkGetBrowser-12    5362    223390 ns/op    1040 B/op    17 allocs/op
```

在 Apple M2 Max 上：
- 完整解析：~0.5ms
- 单项解析（浏览器）：~0.22ms

## 🧪 测试

```bash
# 运行所有测试
go test -v ./...

# 运行完整验证测试（1563 个用例）
go test -v -run='TestFull'

# 运行基准测试
go test -bench=. -benchmem

# 查看测试覆盖率
go test -cover

# 运行示例程序
cd example && go run main.go
```

## 📁 项目结构

```
ua-parser-go/
├── parser.go              # 主解析器
├── mapper.go              # 正则映射器 (使用 regexp2)
├── enums.go               # 常量定义
├── regexes_browser.go     # 浏览器规则
├── regexes_os.go          # 操作系统规则
├── regexes_device.go      # 设备规则
├── regexes_cpu.go         # CPU 规则
├── regexes_engine.go      # 引擎规则
├── parser_test.go         # 单元测试
├── full_validate_test.go  # 完整验证测试
├── example/               # 示例代码
├── go.mod
└── README.md
```

## 🤝 对比原始项目

本项目是 [ua-parser-js](https://github.com/faisalman/ua-parser-js) v2.0.7 的 Go 完整移植版本。

### 优势
- ✅ 编译型语言，性能更高
- ✅ 类型安全，减少运行时错误
- ✅ 内置并发支持
- ✅ 单一二进制文件，部署简单
- ✅ **100% 测试通过率**，与原版完全一致

### 差异
- 空值使用空字符串 `""` 而非 JavaScript 的 `undefined`
- 目前不支持 Client Hints (UACH)

## 📝 许可证

本项目采用与原始 ua-parser-js 项目相同的许可协议 (AGPLv3)。

## 🙏 致谢

本项目基于 [faisalman/ua-parser-js](https://github.com/faisalman/ua-parser-js) 移植，感谢原作者的出色工作！

---

**Made with ❤️ for the Go community**
