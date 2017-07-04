package main

import (
    "io"
    "crypto/rand"
    "testing"
    "crypto/aes"
    "crypto/cipher"
)

func prepareAES() (sourceData, nonce []byte, gcm cipher.AEAD) {
    sourceData = make([]byte, 128)
    io.ReadFull(rand.Reader, sourceData)
    key := make([]byte, 32)
    io.ReadFull(rand.Reader, sourceData)
    nonce = make([]byte, 12)
    io.ReadFull(rand.Reader, nonce)
    block, _ := aes.NewCipher(key)
    gcm, _ = cipher.NewGCM(block)
    return
}

func BenchmarkAESEncryption(b *testing.B) {
    sourceData, nonce, gcm := prepareAES()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        gcm.Seal(nil, nonce, sourceData, nil)
    }
}

func BenchmarkAESDecryption(b *testing.B) {
    sourceData, nonce, gcm := prepareAES()
    encrypted := gcm.Seal(nil, nonce, sourceData, nil)
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        gcm.Open(nil, nonce, encrypted, nil)
    }
}
