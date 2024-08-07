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

var s256 = flag.Bool("sha256", false, "Флаг активирует вывод дайджеста в формате sha256")
var s384 = flag.Bool("sha384", false, "Флаг активирует вывод дайджеста в формате sha384")
var s512 = flag.Bool("sha512", false, "Флаг активирует вывод дайджеста в формате sha512")
var verb = flag.Bool("v", false, "При активации флага активируется подробный отладочный вывод")

func Checksum(s string) string {
	flag.Parse()

	if !*s256 && !*s384 && !*s512 { // по умолчанию конвертируем в контрольную сумму sha256
		if *verb {
			fmt.Printf("Входящие данные: %s конвертируем в контрольную сумму(sha256 checksum):\n%x\n\n", s, sha256.Sum256([]byte(s)))
		}
		b := sha256.Sum256([]byte(s))
		return fmt.Sprintf("%x", b)
	} else if *s256 {
		if *verb {
			fmt.Printf("Входящие данные: %s, конвертируем в контрольную сумму(sha256 checksum):\n%x\n\n", s, sha256.Sum256([]byte(s)))
		}
		b := sha256.Sum256([]byte(s))
		return fmt.Sprintf("%x", b)
	} else if *s384 {
		if *verb {
			fmt.Printf("Входящие данные: %s, конвертируем в контрольную сумму(sha384 checksum):\n%x\n\n", s, sha512.Sum384([]byte(s)))
		}
		b := sha512.Sum384([]byte(s))
		return fmt.Sprintf("%x", b)
	} else if *s512 {
		if *verb {
			fmt.Printf("Входящие данные: %s конвертируем в контрольную сумму(sha512 checksum):\n%x\n\n", s, sha512.Sum512([]byte(s)))
		}
		b := sha512.Sum512([]byte(s))
		return fmt.Sprintf("%x", b)
	}
	return ""
}

