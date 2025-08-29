# cbhelper

一个用于处理剪贴板数据的命令行工具，支持多种编码和解码操作。

## 功能特性

### 解码选项 (-d)
- `b`: Base64 解码
- `z`: Gzip 解压缩
- `u`: URL 解码
- `y`: Bytes 数组解码 - 将形如 `[12 123 23]` 的字节数组转换为字符串

### 编码选项 (-e)
- `b`: Base64 编码
- `z`: Gzip 压缩
- `u`: URL 编码

### 输出选项 (-o)
- `clipboard`: 输出到剪贴板 (默认)
- `stdout`: 输出到标准输出

### 格式选项 (-f)
- `raw`: 原始格式 (默认)
- `json`: JSON 格式
- `pretty-json`: 美化的 JSON 格式

## 使用示例

### Bytes 数组解码
```bash
# 将 [72 101 108 108 111] 解码为 "Hello"
echo "[72 101 108 108 111]" | pbcopy
cbhelper -d y -o stdout

# 支持逗号分隔的格式
echo "[72, 101, 108, 108, 111]" | pbcopy  
cbhelper -d y -o stdout

# 将结果输出到剪贴板
echo "[87 111 114 108 100]" | pbcopy
cbhelper -d y  # 结果会保存到剪贴板
```

### 其他编解码操作
```bash
# Base64 解码
cbhelper -d b -o stdout

# URL 编码
cbhelper -e u -o stdout

# 组合操作：先 Base64 解码，再 Gzip 解压
cbhelper -d bz -o stdout
```