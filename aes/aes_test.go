package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	. "testing"

	"github.com/davecgh/go-spew/spew"
)

func Test_EncryptCorrectness(t *T) {
	key := make([]byte, 32)
	original, err := aes.NewCipher(key)
	if err != nil {
		t.Fatal("Could not create cipher:", err)
	}
	openssl, err := NewCipher(key)
	if err != nil {
		t.Fatal("Could not create cipher:", err)
	}

	plaintext := make([]byte, 32)
	out1 := make([]byte, len(plaintext))
	out2 := make([]byte, len(plaintext))
	original.Encrypt(out1, plaintext)
	openssl.Encrypt(out2, plaintext)

	if !bytes.Equal(out1, out2) {
		t.Error("Encoded output fails.")
		spew.Dump(out1)
		spew.Dump(out2)
	}
}

func Test_DecryptCorrectness(t *T) {
	key := make([]byte, 32)
	original, err := aes.NewCipher(key)
	if err != nil {
		t.Fatal("Could not create cipher:", err)
	}
	openssl, err := NewCipher(key)
	if err != nil {
		t.Fatal("Could not create cipher:", err)
	}

	ciphertext := []byte{
		0xdc, 0x95, 0xc0, 0x78, 0xa2, 0x40, 0x89, 0x89,
		0xad, 0x48, 0xa2, 0x14, 0x92, 0x84, 0x20, 0x87,
	}
	out1 := make([]byte, len(ciphertext))
	out2 := make([]byte, len(ciphertext))
	original.Decrypt(out1, ciphertext)
	openssl.Decrypt(out2, ciphertext)

	for i := 0; i < len(ciphertext); i++ {
		if out1[i] != 0 {
			t.Error("[out1] Expected all bytes to be zero, got:", out1[i])
		}
		if out2[i] != 0 {
			t.Error("[out2] Expected all bytes to be zero, got:", out2[i])
		}
	}
	if !bytes.Equal(out1, out2) {
		t.Error("Encoded output fails.")
	}
}

func Test_EncryptCorrectnessCTR(t *T) {
	key := make([]byte, 32)
	nonce := make([]byte, 16)

	original, err := aes.NewCipher(key)
	if err != nil {
		t.Fatal("Could not create cipher:", err)
	}
	ctroriginal := cipher.NewCTR(original, nonce)
	ctropenssl, err := NewCipherCTR(key, nonce)
	if err != nil {
		t.Fatal("Could not create cipher:", err)
	}

	plaintext := make([]byte, 64)
	out1 := make([]byte, len(plaintext))
	out2 := make([]byte, len(plaintext))
	ctroriginal.XORKeyStream(out1, plaintext)
	ctropenssl.XORKeyStream(out2, plaintext)

	for i := 0; i < len(plaintext); i++ {
		if out1[i] == 0 {
			t.Error("[out1] Expected all bytes to be non-zero, got:", out1[i])
		}
		if out2[i] == 0 {
			t.Error("[out2] Expected all bytes to be non-zero, got:", out2[i])
		}
	}

	if !bytes.Equal(out1, out2) {
		t.Error("Encoded output fails.")
	}
}

func Test_DecryptCorrectnessCTR(t *T) {
	key := make([]byte, 32)
	nonce := make([]byte, 16)

	original, err := aes.NewCipher(key)
	if err != nil {
		t.Fatal("Could not create cipher:", err)
	}
	ctroriginal := cipher.NewCTR(original, nonce)
	ctropenssl, err := NewCipherCTR(key, nonce)
	if err != nil {
		t.Fatal("Could not create cipher:", err)
	}

	plaintext := make([]byte, 64)
	ciphertext := make([]byte, 64)
	ctroriginal.XORKeyStream(ciphertext, plaintext)
	out1 := make([]byte, len(ciphertext))
	out2 := make([]byte, len(ciphertext))

	ctroriginal = cipher.NewCTR(original, nonce)
	ctroriginal.XORKeyStream(out1, ciphertext)
	ctropenssl.XORKeyStream(out2, ciphertext)

	for i := 0; i < len(plaintext); i++ {
		if out1[i] != 0 {
			t.Error("[out1] Expected all bytes to be zero, got:", out1[i])
		}
		if out2[i] != 0 {
			t.Error("[out2] Expected all bytes to be zero, got:", out2[i])
		}
	}

	if !bytes.Equal(out1, out2) {
		t.Error("Encoded output fails.")
	}
}
