package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
)

// 例如: "a6f5c93019d24c0c9d31b8b8e41234f2"
func GenerateAndroidID32() (string, error) {
	b := make([]byte, 16) // 16 bytes -> 32 hex chars
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

// 例如: "9774d56d682e549c"
func GenerateAndroidID16() (string, error) {
	b := make([]byte, 8) // 8 bytes -> 16 hex chars
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

//ANDROID_ID (32 hex): 46b0bc26c2c1f563fccba9a9d378b95c
//ANDROID_ID (16 hex): e0b752846758fb87
//

//5b0eb1a6882b963e432dc480458f7ed0
//e84e935091a6be0ba8d431445de4fa5a
//0e2f7484ccb1f5faec26773726acb56a
//d8ef58198e25e861a6ada91d4e626538
//2f6021fd5325f0eb215b599a56da51c5
//adf6e68dc58fa7367e44b4492e7b44c8
//f0048aed19a379ab6cdf4e3c1fa3e65c
//a372e675ca3b8cddf9222d22f5807354
//bc2da0b20e20f3469d55a45aabc2017a
//4c684e9d1206cdd5a596e679c8254827

func main() {
	// 生成并打印示例
	id32, err := GenerateAndroidID32()
	if err != nil {
		log.Fatalf("ANDROIDID: %v", err)
	}
	fmt.Println("ANDROIDID:", id32)

	id16, err := GenerateAndroidID16()
	if err != nil {
		log.Fatalf("ANDROIDID : %v", err)
	}
	fmt.Println("ANDROIDID:", id16)

}
