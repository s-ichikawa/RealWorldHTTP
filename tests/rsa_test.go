package main

import (
    "crypto/rsa"
    "io"
    "crypto/rand"
    "testing"
    "crypto/md5"
)

func prepareRSA() (sourceData, label []byte, privateKey *rsa.PrivateKey) {
    sourceData = make([]byte, 128)
    label = []byte("")
    io.ReadFull(rand.Reader, sourceData)
    privateKey, _ = rsa.GenerateKey(rand.Reader, 2048)
    return
}

func BenchmarkRSAEncryption(b *testing.B) {
    sourceData, label, privateKey := prepareRSA()
    publicKey := &privateKey.PublicKey
    md5hash := md5.New()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        rsa.EncryptOAEP(md5hash, rand.Reader, publicKey, sourceData, label)
    }
}

func BenchmarkRSADecryption(b *testing.B) {
    sourceData, label, privateKey := prepareRSA()
    publicKey := &privateKey.PublicKey
    md5hash := md5.New()
    encrypted, _ := rsa.EncryptOAEP(md5hash, rand.Reader, publicKey, sourceData, label)
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        rsa.DecryptOAEP(md5hash, rand.Reader, privateKey, encrypted, label)
    }
}
