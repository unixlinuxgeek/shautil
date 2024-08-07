// Упражнение 4.2.
//
// Напишите функцию, которая по умолчанию выводит дайджест SHA256 для входных данных,
// но при использовании соответствующих флагов командной строки выводит SHA384 или SHA512.

package shautil

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

var s256 = flag.Bool("sha256", false, "usage: -sha256 <symbol>, example: -sha256 1")
var s384 = flag.Bool("sha384", false, "usage: -sha384 <symbol>, example: -sha384 2")
var s512 = flag.Bool("sha512", false, "usage: -sha512 <symbol>, example: -sha512 3")

func Checksum(s string) string {
	flag.Parse()

	if !*s256 && !*s384 && !*s512 { // по умолчанию конвертируем в контрольную сумму sha256
		fmt.Printf("Входящие данные: %s конвертируем в контрольную сумму(sha256 checksum):\n%x\n\n", s, sha256.Sum256([]byte(s)))
		b := sha256.Sum256([]byte(s))
		return fmt.Sprintf("%x", b)
	} else if *s256 {
		fmt.Printf("Входящие данные: %s, конвертируем в контрольную сумму(sha256 checksum):\n%x\n\n", s, sha256.Sum256([]byte(s)))
		b := sha256.Sum256([]byte(s))
		return fmt.Sprintf("%x", b)
	} else if *s384 {
		fmt.Printf("Входящие данные: %s, конвертируем в контрольную сумму(sha384 checksum):\n%x\n\n", s, sha512.Sum384([]byte(s)))
		b := sha512.Sum384([]byte(s))
		return fmt.Sprintf("%x", b)
	} else if *s512 {
		fmt.Printf("Входящие данные: %s конвертируем в контрольную сумму(sha512 checksum):\n%x\n\n", s, sha512.Sum512([]byte(s)))
		b := sha512.Sum512([]byte(s))
		return fmt.Sprintf("%x", b)
	}
	return ""
}

