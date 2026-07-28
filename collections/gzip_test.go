package collections

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"testing"
	"time"
)

// 用于演示和验证Go HTTP客户端自动处理gzip压缩内容的功能
func TestGzipAutoDecompress(t *testing.T) {
	// 创建服务器
	go func() {
		http.HandleFunc("/gzipped-content", handleGzippedContent)
		log.Println("Server starting on :8080")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal("Server failed to start:", err)
		}
	}()

	// 等待服务器启动
	time.Sleep(100 * time.Millisecond)

	// 创建客户端并发送请求
	client := &http.Client{}

	fmt.Println("=== 开始测试Go客户端自动解压缩gzip功能 ===")

	// 发送请求到服务器
	resp, err := client.Get("http://localhost:8080/gzipped-content")
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}
	defer resp.Body.Close()

	// 检查响应头
	fmt.Printf("Response Status: %s\n", resp.Status)
	fmt.Println("Response Headers:")
	for key, values := range resp.Header {
		for _, value := range values {
			fmt.Printf("  %s: %s\n", key, value)
		}
	}

	// 读取响应体（这个过程会自动解压gzip）
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	fmt.Printf("Response Body (自动解压后): %q\n", string(body))
	fmt.Printf("Response Body Length: %d bytes\n", len(body))

	// 验证内容是否正确解压
	expected := "This is a test string that will be compressed using gzip. "
	expected += strings.Repeat("This is repeated content to make the compression more effective. ", 10)

	if string(body) != expected {
		t.Errorf("Response body does not match expected. Got: %q, Expected: %q", string(body), expected)
	} else {
		fmt.Println("✅ 测试通过：客户端自动解压缩了gzip内容")
	}
}

// 处理函数，返回gzip压缩的内容
func handleGzippedContent(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request from %s for /gzipped-content", r.RemoteAddr)

	// 打印请求头信息
	log.Println("Request Headers:")
	for key, values := range r.Header {
		for _, value := range values {
			log.Printf("  %s: %s", key, value)
		}
	}

	// 创建一个字符串作为原始内容
	originalContent := "This is a test string that will be compressed using gzip. "
	originalContent += strings.Repeat("This is repeated content to make the compression more effective. ", 10)

	// 使用gzip压缩内容
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)

	_, err := gz.Write([]byte(originalContent))
	if err != nil {
		http.Error(w, "Failed to compress content", http.StatusInternalServerError)
		return
	}

	err = gz.Close()
	if err != nil {
		http.Error(w, "Failed to close gzip writer", http.StatusInternalServerError)
		return
	}

	// 设置响应头表示内容是gzip压缩的
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", buf.Len()))

	// 记录压缩后数据大小
	log.Printf("Sending gzipped response, compressed size: %d bytes", buf.Len())
	log.Printf("Original uncompressed content: %q", originalContent)
	log.Printf("Original uncompressed length: %d bytes", len(originalContent))

	// 写入压缩后的内容
	_, err = w.Write(buf.Bytes())
	if err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}