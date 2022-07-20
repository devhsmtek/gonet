# gonet/port

Port Scanner Module

Usage :

```go
package main

import (
	"fmt"

	hsmport "github.com/husamettinarabaci/gonet/port"
)

func init() {
	hsmport.SetCb(portScanCb)
}

func main() {
	err := hsmport.Scan()
	if err != nil {
		fmt.Println(err)
	}
}

func portScanCb(portStatuses hsmport.SStatuses, err error) {

	if err != nil {
		//
	} else {
		var resStr string = ""
		for _, portStatus := range portStatuses.Statuses {
			resStr += fmt.Sprintf("Statu: %s - Local Address: %s:%s - Remote Address: %s:%s\n", portStatus.Statu, portStatus.LocalAddress, portStatus.LocalPort, portStatus.RemoteAddress, portStatus.RemotePort)
		}
		fmt.Println(resStr)
	}
}
```

## Results:

```shell
Statu: LISTENING - Local Address: 127.0.0.53%lo:53 - Remote Address: 0.0.0.0:*
Statu: LISTENING - Local Address: 127.0.0.1:631 - Remote Address: 0.0.0.0:*
Statu: TIME-WAIT - Local Address: 192.168.1.54:46576 - Remote Address: 172.217.17.97:443
Statu: ESTABLISHED - Local Address: 192.168.1.54:49198 - Remote Address: 140.82.112.25:443
```

## Port Statu Types:

```go
type EStatuType string

const (
	ESTABLISHED EStatuType = "ESTABLISHED"
	LISTENING   EStatuType = "LISTENING"
	SYN_SENT               = "SYN-SENT"
	SYN_RECV               = "SYN-RECV"
	LAST_ACK               = "LAST-ACK"
	FIN_WAIT_1             = "FIN-WAIT-1"
	FIN_WAIT_2             = "FIN-WAIT-2"
	TIME_WAIT              = "TIME-WAIT"
	CLOSE_WAIT             = "CLOSE_WAIT"
	CLOSED                 = "CLOSED"
	CLOSING                = "CLOSING"
	UNKNOWN                = "UNKNOWN"
)
```

## Result Type:

```go
type SStatus struct {
	Statu         EStatuType `protobuf:"bytes,1,opt,name=statu,proto3" json:"statu"`
	LocalAddress  string     `protobuf:"bytes,2,opt,name=local_address,proto3" json:"local_address"`
	LocalPort     string     `protobuf:"bytes,3,opt,name=local_port,proto3" json:"local_port"`
	RemoteAddress string     `protobuf:"bytes,4,opt,name=remote_address,proto3" json:"remote_address"`
	RemotePort    string     `protobuf:"bytes,5,opt,name=remote_port,proto3" json:"remote_port"`
}
```

## Installation:

```shell
go get "github.com/husamettinarabaci/gonet/port@latest"
```

## Developers:
<img src="https://github.com/husamettinarabaci/husamettinarabaci/blob/main/hsmtek-logo.png?raw=true" width="200"/>
Web Site        : www.hsmteknoloji.com <br />