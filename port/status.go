package port

import (
	"encoding/json"
	"fmt"
	"time"

	portpb "github.com/husamettinarabaci/gonet/port/proto"
)

//
// TODO
//
type EStatuType string

const (
	ESTAB     EStatuType = "ESTAB"
	LISTEN    EStatuType = "LISTEN"
	SYN_SENT             = "SYN-SENT"
	SYN_RECV             = "SYN-RECV"
	LAST_ACK             = "LAST-ACK"
	FIN_WAIT             = "FIN-WAIT"
	TIME_WAIT            = "TIME-WAIT"
	CLOSE                = "CLOSE"
	UNKNOWN              = "UNKNOWN"
)

//
// TODO
//
func (pS *EStatuType) StatuTypeConvert(statu string) {
	switch statu {
	case "LISTENING", "LISTEN":
		*pS = LISTEN
	case "CLOSING", "CLOSE", "CLOSED", "CLOSE-WAIT", "CLOSE_WAIT":
		*pS = CLOSE
	case "FIN-WAIT-1", "FIN_WAIT_1", "FIN-WAIT-2", "FIN_WAIT_2":
		*pS = FIN_WAIT
	case "TIME_WAIT", "TIME-WAIT":
		*pS = TIME_WAIT
	case "LAST-ACK", "LAST_ACK":
		*pS = LAST_ACK
	case "SYN_RECEIVED", "SYN-RECEIVED", "SYN_RECV", "SYN-RECV":
		*pS = SYN_RECV
	case "SYN_SEND", "SYN-SEND", "SYN_SENT", "SYN-SENT":
		*pS = SYN_SENT
	case "ESTABLISHED", "ESTAB":
		*pS = ESTAB
	default:
		*pS = UNKNOWN
	}
}

//
// TODO
//
type SStatus struct {
	Statu         EStatuType `protobuf:"bytes,1,opt,name=statu,proto3" json:"statu"`
	LocalAddress  string     `protobuf:"bytes,2,opt,name=local_address,proto3" json:"local_address"`
	LocalPort     string     `protobuf:"bytes,3,opt,name=local_port,proto3" json:"local_port"`
	RemoteAddress string     `protobuf:"bytes,4,opt,name=remote_address,proto3" json:"remote_address"`
	RemotePort    string     `protobuf:"bytes,5,opt,name=remote_port,proto3" json:"remote_port"`
	OpeningTime   time.Time  `protobuf:"bytes,6,opt,name=opening_time,proto3" json:"opening_time"`
	ClosingTime   time.Time  `protobuf:"bytes,7,opt,name=closing_time,proto3" json:"closing_time"`
	LastScanTime  time.Time  `protobuf:"bytes,8,opt,name=last_scan_time,proto3" json:"last_scan_time"`
	IsOpen        bool       `protobuf:"bytes,9,opt,name=isopen,proto3" json:"isopen"`
}

//
//TODO
//
func (ps SStatus) String() string {
	return fmt.Sprintf("Statu: %s - Local Address: %s:%s - Remote Address: %s:%s - IsOpen: %v - OpeningTime: %s - ClosingTime: %s - LastScanTime: %s\n", ps.Statu, ps.LocalAddress, ps.LocalPort, ps.RemoteAddress, ps.RemotePort, ps.IsOpen, ps.OpeningTime, ps.ClosingTime, ps.LastScanTime)
}

func (ps SStatus) EqualTo(in SStatus) bool {
	if ps.LocalAddress == in.LocalAddress &&
		ps.LocalPort == in.LocalPort &&
		ps.RemoteAddress == in.RemoteAddress &&
		ps.RemotePort == in.RemotePort {
		return true
	} else {
		return false
	}
}

//
//TODO
//
func (ps *SStatus) SetStatu(statu string) {
	ps.Statu.StatuTypeConvert(statu)
	if ps.Statu == CLOSE {
		ps.IsOpen = false
	} else {
		ps.IsOpen = true
	}

}

//
// TODO
//
func (ps SStatus) ToProto() *portpb.PbStatus {
	b, _ := json.Marshal(ps)
	myProto := &portpb.PbStatus{}
	json.Unmarshal(b, myProto)
	return myProto
}

//
// TODO
//
func (ps *SStatus) ToStruct(myProto *portpb.PbStatus) {
	b, _ := json.Marshal(myProto)
	json.Unmarshal(b, ps)
}

//
// TODO
//
type SStatuses struct {
	Statuses []SStatus `protobuf:"bytes,1,opt,name=statuses,proto3" json:"statuses"`
}

//
//TODO
//
func (ps SStatuses) String() string {
	retVal := ""
	for _, statu := range ps.Statuses {
		retVal += statu.String()
	}
	return retVal
}

//
// TODO
//
func (ps SStatuses) ToProto() *portpb.PbStatuses {
	b, _ := json.Marshal(ps)
	myProto := &portpb.PbStatuses{}
	json.Unmarshal(b, myProto)
	return myProto
}

//
// TODO
//
func (ps *SStatuses) ToStruct(myProto *portpb.PbStatuses) {
	b, _ := json.Marshal(myProto)
	json.Unmarshal(b, ps)
}
