package pkg

import "strings"

func RemoveBOM(s string) string {
	// 定义 BOM 的 Unicode 表示
	bom := "\ufeff"

	// 检查字符串是否以 BOM 开头，如果是，则移除它
	if strings.HasPrefix(s, bom) {
		return s[len(bom):]
	}
	return s
}