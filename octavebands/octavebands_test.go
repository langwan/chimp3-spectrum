package octavebands

import (
	"testing"
)

func TestGenBanks(t *testing.T) {
	t.Log(GenBanks(1))
	t.Log(GenBanks(1.0 / 3.0))
	t.Log(GenBanks(1.0 / 24.0))
}
