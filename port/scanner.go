package port

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"sync"
	"time"
)

//
// TODO
//
type Scanner struct {
	os            string
	isLinux       bool
	isWindows     bool
	isMac         bool
	isSupportedOs bool
	cb            func(SStatuses, error)
	err           error
	command       string
	commandArgs   []string
}

//Create new instance of Scanner
//Params:
// - cb func(PortStatuses, hsm.H_Error):
//   Scanner struct's cb field
//Returns:
// - *Scanner:
//   New *Scanner instance
func NewScanner(cb func(SStatuses, error)) *Scanner {
	var s Scanner
	s.os = runtime.GOOS
	s.isLinux = s.os == "linux"
	s.isWindows = s.os == "windows"
	s.isMac = s.os == "darwin"
	s.isSupportedOs = s.isLinux || s.isWindows || s.isMac
	s.cb = cb
	s.generateCommand()
	return &s
}

//Generate Scanner's Command for Os Type
func (s *Scanner) generateCommand() {
	if s.isLinux || s.isWindows || s.isMac {
		if s.isLinux {
			s.command = "ss"
			s.commandArgs = []string{"-atn"}
		} else if s.isWindows {
			s.command = "netstat"
			s.commandArgs = []string{"-ano", "-p", "tcp"}
		} else {
			s.command = "netstat"
			s.commandArgs = []string{"-an", "-p", "tcp", "-f", "inet"}
		}
	} else {
		s.err = fmt.Errorf("%s is not supported", s.os)
	}

}

//Scan ports
//Call to Scanner's cb function
func (s *Scanner) Scan() {
	var portStatuses SStatuses
	if s.isSupportedOs {
		cmd := exec.Command(s.command, s.commandArgs...)
		var outb, errb bytes.Buffer
		cmd.Stdout = &outb
		cmd.Stderr = &errb
		err := cmd.Run()
		if err != nil {
			s.cb(portStatuses, err)
		} else {
			for rI, row := range strings.Split(outb.String(), "\n") {
				if ((s.isLinux && rI > 0) || (s.isWindows && rI > 3) || (s.isMac && rI > 1)) && row != "" && len(row) > 0 {
					var pRes SStatus
					i := 0
					for _, col := range strings.Split(row, " ") {
						if col != "" {
							if s.isLinux {
								switch i {
								case 0:
									pRes.SetStatu(col)
								case 3:
									sepInd := strings.LastIndex(col, ":")
									pRes.LocalAddress = col[:sepInd]
									pRes.LocalPort = col[sepInd+1:]
								case 4:
									sepInd := strings.LastIndex(col, ":")
									pRes.RemoteAddress = col[:sepInd]
									pRes.RemotePort = col[sepInd+1:]
								}
							} else if s.isWindows {
								switch i {
								case 3:
									pRes.SetStatu(col)
								case 1:
									sepInd := strings.LastIndex(col, ":")
									pRes.LocalAddress = col[:sepInd]
									pRes.LocalPort = col[sepInd+1:]
								case 2:
									sepInd := strings.LastIndex(col, ":")
									pRes.RemoteAddress = col[:sepInd]
									pRes.RemotePort = col[sepInd+1:]
								}
							} else {
								switch i {
								case 5:
									pRes.SetStatu(col)
								case 3:
									sepInd := strings.LastIndex(col, ".")
									pRes.LocalAddress = col[:sepInd]
									pRes.LocalPort = col[sepInd+1:]
								case 4:
									sepInd := strings.LastIndex(col, ".")
									pRes.RemoteAddress = col[:sepInd]
									pRes.RemotePort = col[sepInd+1:]
								}
							}
							i++
						}

					}
					pRes.LastScanTime = time.Now()
					portStatuses.Statuses = append(portStatuses.Statuses, pRes)
				}
			}
			s.cb(portStatuses, nil)
		}

	} else {
		s.cb(portStatuses, s.err)
	}

}

var instance *Scanner
var lc sync.RWMutex
var cb func(SStatuses, error)

//
// TODO
//
func SetCb(_cb func(SStatuses, error)) {
	cb = _cb
}

//
// TODO
//
func GetInstace() *Scanner {
	lc.Lock()
	defer lc.Unlock()
	if instance == nil {
		instance = NewScanner(cb)
	}
	return instance
}

//
// TODO
//
func Scan() error {
	if cb != nil {
		GetInstace().Scan()
		return nil
	} else {
		return fmt.Errorf("%s", "Call back func is nil")
	}
}
