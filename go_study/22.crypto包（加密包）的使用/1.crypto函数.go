package main

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
)

func main() {

	fmt.Println("=== 哈希计算 ===")

	data := "Hello World 中文"

	// 1. MD5 哈希（128位，不提交用于安全场景）
	fmt.Println("1. MD5 哈希:")
	md5Hash := md5.New()
	md5Hash.Write([]byte(data))
	fmt.Printf("MD5(\"%s\") = %x\n ", data, md5Hash.Sum(nil))

	// 2. SHA1 哈希（160位，已经被认为不够安全）
	fmt.Println("\n2. SHA1 哈希:")
	sha1Hash := sha1.New()
	sha1Hash.Write([]byte(data))
	fmt.Printf("SHA1(\"%s\") = %x\n ", data, sha1Hash.Sum(nil))

	// 3. SHA256 哈希（256位，常用且安全）
	fmt.Println("\n3. SHA256 哈希:")
	sha256Hash := sha256.New()
	sha256Hash.Write([]byte(data))
	fmt.Printf("SHA256(\"%s\") = %x\n ", data, sha256Hash.Sum(nil))

	// 4. SHA512 哈希（512位，更安全但更长）
	fmt.Println("\n4. SHA512 哈希:")
	sha512Hash := sha512.New()
	sha512Hash.Write([]byte(data))
	fmt.Printf("SHA512(\"%s\") = %x\n ", data, sha512Hash.Sum(nil))

	// 5. 使用 hex 编码展示不同的输出格式
	fmt.Println("\n5. 不同编码格式:")
	sha256Hash2 := sha256.Sum256([]byte(data))
	fmt.Printf("十六进制: %x\n", sha256Hash2)
	fmt.Printf("十六进制(大写): %X\n", sha256Hash2)
	fmt.Printf("hex.EncodeToString: %s\n", hex.EncodeToString(sha256Hash2[:]))

	// 6. 演示相同数据产生相同哈希
	fmt.Println("\n6. 哈希一致性验证:")
	data1 := "password123"
	data2 := "password123"
	data3 := "password124" // 只有一个字符不同

	hash1 := sha256.Sum256([]byte(data1))
	hash2 := sha256.Sum256([]byte(data2))
	hash3 := sha256.Sum256([]byte(data3))

	fmt.Printf("   \"%s\" -> %x\n", data1, hash1)
	fmt.Printf("   \"%s\" -> %x\n", data2, hash2)
	fmt.Printf("   \"%s\" -> %x\n", data3, hash3)
	fmt.Printf("   data1 == data2: %v (哈希相同)\n", hash1 == hash2)
	fmt.Printf("   data1 == data3: %v (哈希不同，即使只差一个字符)\n", hash1 == hash3)

	// 7. HMAC - 带密钥的哈希消息认证码
	fmt.Println("\n7. HMAC 消息认证:")
	secret := []byte("my-secret-key")
	message := "Important message"

	// 使用 HMAC-SHA256
	h := hmac.New(sha256.New, secret)
	h.Write([]byte(message))
	signature := h.Sum(nil)
	fmt.Printf("   消息: \"%s\"\n", message)
	fmt.Printf("   密钥: \"%s\"\n", string(secret))
	fmt.Printf("   HMAC-SHA256: %x\n", signature)

	// 验证 HMAC
	h2 := hmac.New(sha256.New, secret)
	h2.Write([]byte(message))
	expectedSignature := h2.Sum(nil)
	isValid := hmac.Equal(signature, expectedSignature)
	fmt.Printf("   签名验证: %v\n", isValid)

	// 使用错误的密钥验证
	wrongSecret := []byte("wrong-key")
	h3 := hmac.New(sha256.New, wrongSecret)
	h3.Write([]byte(message))
	wrongSignature := h3.Sum(nil)
	isValidWrong := hmac.Equal(signature, wrongSignature)
	fmt.Printf("   错误密钥验证: %v\n", isValidWrong)

	// 8. 文件内容哈希计算
	fmt.Println("\n8. 文件内容哈希:")

	// 9. 多次写入累积哈希
	fmt.Println("\n9. 流式哈希计算（多次写入）:")
	streamHash := sha256.New()
	chunks := []string{"Hello", " ", "World", "!"}
	for i, chunk := range chunks {
		streamHash.Write([]byte(chunk))
		fmt.Printf("   写入第 %d 块: \"%s\"\n", i+1, chunk)
	}
	finalHash := streamHash.Sum(nil)
	fmt.Printf("   最终哈希: %x\n", finalHash)

	// 对比一次性哈希
	oneTimeHash := sha256.Sum256([]byte("Hello World!"))
	fmt.Printf("   一次性哈希: %x\n", oneTimeHash)
	fmt.Printf("   两种方式结果相同: %v\n", hex.EncodeToString(finalHash) == hex.EncodeToString(oneTimeHash[:]))

	fmt.Println()
}
