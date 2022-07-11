package port

import (
	"encoding/json"

	portpb "github.com/husamettinarabaci/gonet/port/proto"
)

//
// TODO
//
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

//
// TODO
//
func StatuTypeConvert(statu string) EStatuType {
	var pS EStatuType
	switch statu {
	case "LISTENING":
	case "LISTEN":
		pS = LISTENING
	case "CLOSING":
	case "CLOSE":
		pS = CLOSING
	case "FIN-WAIT-1":
	case "FIN_WAIT_1":
		pS = FIN_WAIT_1
	case "FIN-WAIT-2":
	case "FIN_WAIT_2":
		pS = FIN_WAIT_2
	case "TIME_WAIT":
	case "TIME-WAIT":
		pS = TIME_WAIT
	case "CLOSE-WAIT":
	case "CLOSE_WAIT":
		pS = CLOSE_WAIT
	case "LAST-ACK":
	case "LAST_ACK":
		pS = LAST_ACK
	case "SYN_RECEIVED":
	case "SYN-RECEIVED":
	case "SYN_RECV":
	case "SYN-RECV":
		pS = SYN_RECV
	case "SYN_SEND":
	case "SYN-SEND":
	case "SYN_SENT":
	case "SYN-SENT":
		pS = SYN_SENT
	case "ESTABLISHED":
	case "ESTAB":
		pS = ESTABLISHED
	default:
		pS = UNKNOWN
	}
	return pS
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
