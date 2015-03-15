package main

import (
	"fmt"
	"golang.org/x/crypto/twofish"
	"crypto/cipher"
)

func main() {

/* using 32 byte (or 256 bits) crypto key and initialization vector*/
crypto_cipher_key := []byte("12345678901234567890123456789012")
crypto_cipher_text := []byte("12345678901234567890123456789012")

/* the top secret data to transmit */
plain_msg := "The original plaintext message is here, although the compression doesn't seem to make much difference to output size - in fact somewhat expands the required space!"

crypto_cipher_block, err := twofish.NewCipher(crypto_cipher_key)
if err != nil {
	panic(err)
}

crypto_cipher_iv := crypto_cipher_text[:twofish.BlockSize]
crypto_cipher_encrypter := cipher.NewCFBEncrypter(crypto_cipher_block, crypto_cipher_iv)
crypto_cipher_output_buffer := make([]byte, len(plain_msg))
crypto_cipher_encrypter.XORKeyStream(crypto_cipher_output_buffer, []byte(plain_msg))

/* show the encrypted output of our zlib compressed plaintext secret message */
fmt.Printf("\n'%s'\n is encrypted to \n'%v'\n\n", plain_msg, crypto_cipher_output_buffer)


}
