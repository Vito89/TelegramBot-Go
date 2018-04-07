package wiki

import (
	"testing"
)

func TestUrl_encoder(t *testing.T) {

	if _, err := UrlEncoded("test"); err != nil {
		t.Fail()
	}

	if _, err := UrlEncoded("$#%B^&M*(>(*)>(*?(?```````HGKKJ:"); err == nil {
		t.Fail()
	}
}
