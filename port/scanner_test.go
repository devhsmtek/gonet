package port

import (
	"fmt"
	"testing"
)

//
// TODO
//
func TestScanner(t *testing.T) {

	t.Run("Scanner Test", func(t *testing.T) {
		var cb func(SStatuses, error) = func(portStatuses SStatuses, err error) {
			if err != nil {
				t.Error(err)
			} else {
				var resStr string = ""
				for _, portStatu := range portStatuses.Statuses {
					resStr += fmt.Sprintf("Statu: %s - Local Address: %s:%s - Remote Address: %s:%s\n", portStatu.Statu, portStatu.LocalAddress, portStatu.LocalPort, portStatu.RemoteAddress, portStatu.RemotePort)
				}
				fmt.Println(resStr)
			}
		}
		var sc Scanner = *NewScanner(cb)
		sc.Scan()
		if sc.err != nil {
			t.Error(sc.err)
		}
	})
}
