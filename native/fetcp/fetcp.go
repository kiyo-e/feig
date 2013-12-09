// mksyscall_windows.pl -l32 fetcp_api.go
// MACHINE GENERATED BY THE COMMAND ABOVE; DO NOT EDIT

package fetcp

import "unsafe"
import "syscall"

var (
	modfetcp = syscall.NewLazyDLL("fetcp.dll")

	procFETCP_Connect        = modfetcp.NewProc("FETCP_Connect")
	procFETCP_DisConnect     = modfetcp.NewProc("FETCP_DisConnect")
	procFETCP_Detect         = modfetcp.NewProc("FETCP_Detect")
	procFETCP_GetSocketState = modfetcp.NewProc("FETCP_GetSocketState")
	procFETCP_GetSocketList  = modfetcp.NewProc("FETCP_GetSocketList")
	procFETCP_GetDLLVersion  = modfetcp.NewProc("FETCP_GetDLLVersion")
	procFETCP_GetErrorText   = modfetcp.NewProc("FETCP_GetErrorText")
	procFETCP_GetLastError   = modfetcp.NewProc("FETCP_GetLastError")
	procFETCP_GetSocketHnd   = modfetcp.NewProc("FETCP_GetSocketHnd")
	procFETCP_GetSocketPara  = modfetcp.NewProc("FETCP_GetSocketPara")
	procFETCP_SetSocketPara  = modfetcp.NewProc("FETCP_SetSocketPara")
	procFETCP_Transceive     = modfetcp.NewProc("FETCP_Transceive")
	procFETCP_Transmit       = modfetcp.NewProc("FETCP_Transmit")
	procFETCP_Receive        = modfetcp.NewProc("FETCP_Receive")
)

func FETCP_Connect(cHostAdr *byte, iPortNr int) (result int) {
	r0, _, _ := syscall.Syscall(procFETCP_Connect.Addr(), 2, uintptr(unsafe.Pointer(cHostAdr)), uintptr(iPortNr), 0)
	result = int(r0)
	return
}

func FETCP_Disconnect(iSocketHnd int) (result int) {
	r0, _, _ := syscall.Syscall(procFETCP_DisConnect.Addr(), 1, uintptr(iSocketHnd), 0, 0)
	result = int(r0)
	return
}

func FETCP_Detect(cHostAdr *byte, iPortNr int) (result int) {
	r0, _, _ := syscall.Syscall(procFETCP_Detect.Addr(), 2, uintptr(unsafe.Pointer(cHostAdr)), uintptr(iPortNr), 0)
	result = int(r0)
	return
}

func FETCP_GetSocketState(cHostAdr *byte, iPortNr int) (result int) {
	r0, _, _ := syscall.Syscall(procFETCP_GetSocketState.Addr(), 2, uintptr(unsafe.Pointer(cHostAdr)), uintptr(iPortNr), 0)
	result = int(r0)
	return
}

func FETCP_GetSocketList(iNext int) (result int) {
	r0, _, _ := syscall.Syscall(procFETCP_GetSocketList.Addr(), 1, uintptr(iNext), 0, 0)
	result = int(r0)
	return
}

func FETCP_GetDLLVersion(cVersion *byte) {
	syscall.Syscall(procFETCP_GetDLLVersion.Addr(), 1, uintptr(unsafe.Pointer(cVersion)), 0, 0)
	return
}

func FETCP_GetErrorText(iErrorCode int, cErrorText *byte) (result int) {
	r0, _, _ := syscall.Syscall(procFETCP_GetErrorText.Addr(), 2, uintptr(iErrorCode), uintptr(unsafe.Pointer(cErrorText)), 0)
	result = int(r0)
	return
}

func FETCP_GetLastError(iSocketHnd int, iErrorCode *int, cErrorText *byte) (result int) {
	r0, _, _ := syscall.Syscall(procFETCP_GetLastError.Addr(), 3, uintptr(iSocketHnd), uintptr(unsafe.Pointer(iErrorCode)), uintptr(unsafe.Pointer(cErrorText)))
	result = int(r0)
	return
}

func FETCP_GetSocketHnd(cHostAdr *byte, iPortNr int) (result int) {
	r0, _, _ := syscall.Syscall(procFETCP_GetSocketHnd.Addr(), 2, uintptr(unsafe.Pointer(cHostAdr)), uintptr(iPortNr), 0)
	result = int(r0)
	return
}

func FETCP_GetSocketPara(iSocketHnd int, cPara *byte, cValue *byte) (result int) {
	r0, _, _ := syscall.Syscall(procFETCP_GetSocketPara.Addr(), 3, uintptr(iSocketHnd), uintptr(unsafe.Pointer(cPara)), uintptr(unsafe.Pointer(cValue)))
	result = int(r0)
	return
}

func FETCP_SetSocketPara(iSocketHnd int, cPara *byte, cValue *byte) (result int) {
	r0, _, _ := syscall.Syscall(procFETCP_SetSocketPara.Addr(), 3, uintptr(iSocketHnd), uintptr(unsafe.Pointer(cPara)), uintptr(unsafe.Pointer(cValue)))
	result = int(r0)
	return
}

func FETCP_Transceive(iSocketHnd int, cSendProt *byte, iSendLen int, recvBuf *byte, recvBufLen int) (result int) {
	r0, _, _ := syscall.Syscall6(procFETCP_Transceive.Addr(), 5, uintptr(iSocketHnd), uintptr(unsafe.Pointer(cSendProt)), uintptr(iSendLen), uintptr(unsafe.Pointer(recvBuf)), uintptr(recvBufLen), 0)
	result = int(r0)
	return
}

func FETCP_Transmit(iSocketHnd int, cSendProt *byte, iSendLen int) (result int) {
	r0, _, _ := syscall.Syscall(procFETCP_Transmit.Addr(), 3, uintptr(iSocketHnd), uintptr(unsafe.Pointer(cSendProt)), uintptr(iSendLen))
	result = int(r0)
	return
}

func FETCP_Receive(iSocketHnd int, cRecProt *byte, cRecLec int) (result int) {
	r0, _, _ := syscall.Syscall(procFETCP_Receive.Addr(), 3, uintptr(iSocketHnd), uintptr(unsafe.Pointer(cRecProt)), uintptr(cRecLec))
	result = int(r0)
	return
}
