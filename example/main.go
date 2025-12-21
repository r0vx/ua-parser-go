package main

import (
	"encoding/json"
	"fmt"
	"strings"

	uaparser "github.com/example/ua-parser-go"
)

func main() {
	// 示例 User-Agent 字符串
	userAgents := []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",
		"Mozilla/5.0 (iPhone; CPU iPhone OS 17_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 Mobile/15E148 Safari/604.1",
		"Mozilla/5.0 (Linux; Android 14; SM-S918B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.6099.43 Mobile Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 Safari/605.1.15",
	}

	fmt.Println("UA Parser Go - 示例")
	fmt.Println(strings.Repeat("=", 80))
	fmt.Println()

	for i, ua := range userAgents {
		fmt.Printf("示例 %d:\n", i+1)
		fmt.Printf("User-Agent: %s\n\n", ua)

		// 方法 1: 直接解析
		result := uaparser.Parse(ua)

		// 输出 JSON 格式
		jsonData, _ := json.MarshalIndent(result, "", "  ")
		fmt.Println("解析结果 (JSON):")
		fmt.Println(string(jsonData))

		// 方法 2: 使用 Parser 实例
		parser := uaparser.New(ua)
		
		fmt.Println("\n解析结果 (分类):")
		fmt.Printf("  浏览器: %s %s (major: %s)\n", 
			parser.GetBrowser().Name,
			parser.GetBrowser().Version,
			parser.GetBrowser().Major)
		
		fmt.Printf("  操作系统: %s %s\n", 
			parser.GetOS().Name,
			parser.GetOS().Version)
		
		fmt.Printf("  设备: %s %s (%s)\n", 
			parser.GetDevice().Vendor,
			parser.GetDevice().Model,
			parser.GetDevice().Type)
		
		fmt.Printf("  引擎: %s %s\n", 
			parser.GetEngine().Name,
			parser.GetEngine().Version)
		
		fmt.Printf("  CPU架构: %s\n", 
			parser.GetCPU().Architecture)

		fmt.Println()
		fmt.Println(strings.Repeat("-", 80))
		fmt.Println()
	}
}

