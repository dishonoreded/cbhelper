package lib

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// BytesArrayDecode 将形如 "[12 123 23]" 的字符串解码为对应的字符串
func BytesArrayDecode(src []byte) ([]byte, error) {
	input := string(src)
	
	// 移除首尾的方括号和空白字符
	input = strings.TrimSpace(input)
	if !strings.HasPrefix(input, "[") || !strings.HasSuffix(input, "]") {
		return nil, fmt.Errorf("invalid bytes array format, expected format: [12 123 23]")
	}
	
	// 提取方括号内的内容
	content := strings.TrimSpace(input[1 : len(input)-1])
	if content == "" {
		return []byte{}, nil
	}
	
	// 使用正则表达式分割数字，支持空格、逗号等分隔符
	re := regexp.MustCompile(`\s+|,\s*`)
	parts := re.Split(content, -1)
	
	var result []byte
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		
		// 将字符串转换为数字
		num, err := strconv.Atoi(part)
		if err != nil {
			return nil, fmt.Errorf("invalid number in bytes array: %s", part)
		}
		
		// 检查数字范围 (0-255)
		if num < 0 || num > 255 {
			return nil, fmt.Errorf("byte value out of range (0-255): %d", num)
		}
		
		result = append(result, byte(num))
	}
	
	return result, nil
}