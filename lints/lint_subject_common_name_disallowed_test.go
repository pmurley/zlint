// lint_subject_common_name_disallowed_test.go
package lints

import (
	"testing"
)

func TestCommonNameInSAN(t *testing.T) {
	inputPath := "../testlint/testCerts/commonNameInSAN.pem"
	desEnum := Pass
	out, _ := Lints["e_subject_common_name_disallowed"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}

func TestCommonNameNotInSAN(t *testing.T) {
	inputPath := "../testlint/testCerts/utcHasSeconds.pem"
	desEnum := Error
	out, _ := Lints["e_subject_common_name_disallowed"].ExecuteTest(ReadCertificate(inputPath))
	if out.Result != desEnum {
		t.Error(
			"For", inputPath, /* input path*/
			"expected", desEnum, /* The enum you expected */
			"got", out.Result, /* Actual Result */
		)
	}
}
