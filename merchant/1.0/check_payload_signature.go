package revolut_merchant

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"regexp"
	"strings"
)

func (c *Merchant) CheckPayloadSignature(
	revolutSignatureHeader string,
	revolutTimestampHeader string,
	payload []byte,
) bool {
	revolutSignatures := strings.Split(revolutSignatureHeader, ",")
	for _, signature := range revolutSignatures {
		hash := hmac.New(sha256.New, []byte(c.clientRequest.GetSecret()))
		payloadToHash := fmt.Sprintf(
			"%s.%s.%s",
			"v1",
			revolutTimestampHeader,
			string(payload),
		)
		hash.Write([]byte(payloadToHash))
		sign := hex.EncodeToString(hash.Sum(nil))
		cleanSignReg := regexp.MustCompile(`v\d+=`)

		if hmac.Equal(
			[]byte(sign),
			[]byte(cleanSignReg.ReplaceAllString(signature, "")),
		) {
			return true
		}
	}
	return false
}
