package snmp

import (
	"fmt"
	"log"

	"github.com/gosnmp/gosnmp"
)

type SNMPApi struct {
	IpAddr string
}

func NewNetworkListner() *SNMPApi {
	lIpAdrr := "127.0.0.1"
	// if err != nil {
	// 	fmt.Println("An error accured while trying to get ip addr.")
	// 	return nil
	// }
	return &SNMPApi{
		IpAddr: lIpAdrr,
	}
}

func (s *SNMPApi) EstablishConnection() error {
	gosnmp.Default.Target = s.IpAddr
	fmt.Println(s.IpAddr)

	// connecting to snmp
	if err := gosnmp.Default.Connect(); err != nil {
		return fmt.Errorf("An error accured while establshing connection to snmp: %s", err)
	}

	defer gosnmp.Default.Conn.Close()
	/*
	   1.3.6.1.2.1.25
	   1.3.6.1.4.1.2021
	   1.3.6.1.2.1.2
	   1.3.6.1.2.1.1
	*/
	oids := []string{"1.3.6.1.2.1.2.2.1"}
	result, err2 := gosnmp.Default.Get(oids) // Get() accepts up to g.MAX_OIDS
	if err2 != nil {
		log.Fatalf("Get() err: %v", err2)
	}

	for i, variable := range result.Variables {
		fmt.Printf("%d: oid: %s ", i, variable.Name)

		// the Value of each variable returned by Get() implements
		// interface{}. You could do a type switch...
		switch variable.Type {
		case gosnmp.OctetString:
			bytes := variable.Value.([]byte)
			fmt.Printf("string: %s\n", string(bytes))
		default:
			// ... or often you're just interested in numeric values.
			// ToBigInt() will return the Value as a BigInt, for plugging
			// into your calculations.
			fmt.Println(variable.Type.String())
			fmt.Printf("number: %d\n", gosnmp.ToBigInt(variable.Value))
		}
	}

	return nil
}
