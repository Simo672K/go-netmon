package cmd

import (
	"fmt"

	"github.com/Simo672K/go-netmon/pkg/snmp"
)

func Run() {
	snmp := snmp.NewNetworkListner()

	err := snmp.EstablishConnection()
	if err != nil {
		fmt.Println(err)
		return
	}
}
