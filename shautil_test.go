// Модульные тесты

package shautil

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"testing"
)

// Flags Not Set
func TestChecksum(t *testing.T) {
	*s256 = false
	*s384 = false
	*s512 = false

	if !*s256 && !*s384 && !*s512 {
		chSum := Checksum("1")
		ok := fmt.Sprintf("%x", sha256.Sum256([]byte("1")))

		t.Logf("Входящие данные: %s\n", chSum)
		t.Logf("Проверочные данные: %s\n", ok)

		if chSum != ok {
			t.Errorf("%s: %s not equal %s", t.Name(), chSum, ok)
		} else {
			t.Logf("Тест %s успешно пройден!!!", t.Name())
		}
	}
}

// Flag sha256 is set
func TestChecksum256(t *testing.T) {
	*s256 = true
	*s384 = false
	*s512 = false

	if *s256 && !*s384 && !*s512 {
		chSum := Checksum("1")
		ok := fmt.Sprintf("%x", sha256.Sum256([]byte("1")))

		t.Logf("Входящие данные: %s\n", chSum)
		t.Logf("Проверочные данные: %s\n", ok)

		if chSum != ok {
			t.Errorf("%s: %s not equal %s", t.Name(), chSum, ok)
		} else {
			t.Logf("Тест %s успешно пройден!!!", t.Name())
		}
	}
}

// Flag sha384 is set
func TestChecksum384(t *testing.T) {
	*s256 = false
	*s384 = true
	*s512 = false

	if *s256 && *s384 && !*s512 {
		chSum := Checksum("1")
		ok := fmt.Sprintf("%x", sha512.Sum384([]byte("1"))) // correct

		t.Logf("Входящие данные: %s\n", chSum)
		t.Logf("Проверочные данные: %s\n", ok)

		if chSum != ok {
			t.Errorf("%s: %s not equal %s", t.Name(), chSum, ok)
		} else {
			t.Logf("Тест %s успешно пройден!!!", t.Name())
		}
	}
}

// Flag sha512 is set
func TestChecksum512(t *testing.T) {
	*s256 = false
	*s384 = false
	*s512 = true

	if !*s256 && !*s384 && *s512 {
		chSum := Checksum("1")
		ok := fmt.Sprintf("%x", sha512.Sum512([]byte("1"))) // correct

		t.Logf("Входящие данные: %s\n", chSum)
		t.Logf("Проверочные данные: %s\n", ok)

		if chSum != ok {
			t.Errorf("%s: %s not equal %s", t.Name(), chSum, ok)
		} else {
			t.Logf("Тест %s успешно пройден!!!", t.Name())
		}
	}
}
