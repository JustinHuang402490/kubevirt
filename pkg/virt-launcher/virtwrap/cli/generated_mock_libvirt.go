// Automatically generated by MockGen. DO NOT EDIT!
// Source: libvirt.go

package cli

import (
	"fmt"

	gomock "github.com/golang/mock/gomock"
	libvirt_go "libvirt.org/libvirt-go"

	stats "kubevirt.io/kubevirt/pkg/virt-launcher/virtwrap/stats"
)

// Mock of Connection interface
type MockConnection struct {
	ctrl     *gomock.Controller
	recorder *_MockConnectionRecorder
}

// Recorder for MockConnection (not exported)
type _MockConnectionRecorder struct {
	mock *MockConnection
}

func NewMockConnection(ctrl *gomock.Controller) *MockConnection {
	mock := &MockConnection{ctrl: ctrl}
	mock.recorder = &_MockConnectionRecorder{mock}
	return mock
}

func (_m *MockConnection) EXPECT() *_MockConnectionRecorder {
	return _m.recorder
}

func (_m *MockConnection) LookupDomainByName(name string) (VirDomain, error) {
	ret := _m.ctrl.Call(_m, "LookupDomainByName", name)
	ret0, _ := ret[0].(VirDomain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockConnectionRecorder) LookupDomainByName(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "LookupDomainByName", arg0)
}

func (_m *MockConnection) DomainDefineXML(xml string) (VirDomain, error) {
	fmt.Println("\t----------xml in MockConnection DomainDefineXML = \n", xml)
	ret := _m.ctrl.Call(_m, "DomainDefineXML", xml)
	ret0, _ := ret[0].(VirDomain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockConnectionRecorder) DomainDefineXML(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DomainDefineXML", arg0)
}

func (_m *MockConnection) Close() (int, error) {
	ret := _m.ctrl.Call(_m, "Close")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockConnectionRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

func (_m *MockConnection) DomainEventLifecycleRegister(callback libvirt_go.DomainEventLifecycleCallback) error {
	ret := _m.ctrl.Call(_m, "DomainEventLifecycleRegister", callback)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockConnectionRecorder) DomainEventLifecycleRegister(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DomainEventLifecycleRegister", arg0)
}

func (_m *MockConnection) AgentEventLifecycleRegister(callback libvirt_go.DomainEventAgentLifecycleCallback) error {
	ret := _m.ctrl.Call(_m, "AgentEventLifecycleRegister", callback)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockConnectionRecorder) AgentEventLifecycleRegister(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AgentEventLifecycleRegister", arg0)
}

func (_m *MockConnection) ListAllDomains(flags libvirt_go.ConnectListAllDomainsFlags) ([]VirDomain, error) {
	ret := _m.ctrl.Call(_m, "ListAllDomains", flags)
	ret0, _ := ret[0].([]VirDomain)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockConnectionRecorder) ListAllDomains(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListAllDomains", arg0)
}

func (_m *MockConnection) NewStream(flags libvirt_go.StreamFlags) (Stream, error) {
	ret := _m.ctrl.Call(_m, "NewStream", flags)
	ret0, _ := ret[0].(Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockConnectionRecorder) NewStream(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NewStream", arg0)
}

func (_m *MockConnection) SetReconnectChan(reconnect chan bool) {
	_m.ctrl.Call(_m, "SetReconnectChan", reconnect)
}

func (_mr *_MockConnectionRecorder) SetReconnectChan(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetReconnectChan", arg0)
}

func (_m *MockConnection) QemuAgentCommand(command string, domainName string) (string, error) {
	ret := _m.ctrl.Call(_m, "QemuAgentCommand", command, domainName)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockConnectionRecorder) QemuAgentCommand(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "QemuAgentCommand", arg0, arg1)
}

func (_m *MockConnection) GetAllDomainStats(statsTypes libvirt_go.DomainStatsTypes, flags libvirt_go.ConnectGetAllDomainStatsFlags) ([]libvirt_go.DomainStats, error) {
	ret := _m.ctrl.Call(_m, "GetAllDomainStats", statsTypes, flags)
	ret0, _ := ret[0].([]libvirt_go.DomainStats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockConnectionRecorder) GetAllDomainStats(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetAllDomainStats", arg0, arg1)
}

func (_m *MockConnection) GetDomainStats(statsTypes libvirt_go.DomainStatsTypes, flags libvirt_go.ConnectGetAllDomainStatsFlags) ([]*stats.DomainStats, error) {
	ret := _m.ctrl.Call(_m, "GetDomainStats", statsTypes, flags)
	ret0, _ := ret[0].([]*stats.DomainStats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockConnectionRecorder) GetDomainStats(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetDomainStats", arg0, arg1)
}

// Mock of Stream interface
type MockStream struct {
	ctrl     *gomock.Controller
	recorder *_MockStreamRecorder
}

// Recorder for MockStream (not exported)
type _MockStreamRecorder struct {
	mock *MockStream
}

func NewMockStream(ctrl *gomock.Controller) *MockStream {
	mock := &MockStream{ctrl: ctrl}
	mock.recorder = &_MockStreamRecorder{mock}
	return mock
}

func (_m *MockStream) EXPECT() *_MockStreamRecorder {
	return _m.recorder
}

func (_m *MockStream) Read(p []byte) (int, error) {
	ret := _m.ctrl.Call(_m, "Read", p)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockStreamRecorder) Read(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Read", arg0)
}

func (_m *MockStream) Write(p []byte) (int, error) {
	ret := _m.ctrl.Call(_m, "Write", p)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockStreamRecorder) Write(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Write", arg0)
}

func (_m *MockStream) Close() error {
	ret := _m.ctrl.Call(_m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockStreamRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

func (_m *MockStream) UnderlyingStream() *libvirt_go.Stream {
	ret := _m.ctrl.Call(_m, "UnderlyingStream")
	ret0, _ := ret[0].(*libvirt_go.Stream)
	return ret0
}

func (_mr *_MockStreamRecorder) UnderlyingStream() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UnderlyingStream")
}

// Mock of VirDomain interface
type MockVirDomain struct {
	ctrl     *gomock.Controller
	recorder *_MockVirDomainRecorder
}

// Recorder for MockVirDomain (not exported)
type _MockVirDomainRecorder struct {
	mock *MockVirDomain
}

func NewMockVirDomain(ctrl *gomock.Controller) *MockVirDomain {
	mock := &MockVirDomain{ctrl: ctrl}
	mock.recorder = &_MockVirDomainRecorder{mock}
	return mock
}

func (_m *MockVirDomain) EXPECT() *_MockVirDomainRecorder {
	return _m.recorder
}

func (_m *MockVirDomain) GetState() (libvirt_go.DomainState, int, error) {
	ret := _m.ctrl.Call(_m, "GetState")
	ret0, _ := ret[0].(libvirt_go.DomainState)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockVirDomainRecorder) GetState() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetState")
}

func (_m *MockVirDomain) Create() error {
	ret := _m.ctrl.Call(_m, "Create")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVirDomainRecorder) Create() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Create")
}

func (_m *MockVirDomain) Suspend() error {
	ret := _m.ctrl.Call(_m, "Suspend")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVirDomainRecorder) Suspend() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Suspend")
}

func (_m *MockVirDomain) Resume() error {
	ret := _m.ctrl.Call(_m, "Resume")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVirDomainRecorder) Resume() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Resume")
}

func (_m *MockVirDomain) DestroyFlags(flags libvirt_go.DomainDestroyFlags) error {
	ret := _m.ctrl.Call(_m, "DestroyFlags", flags)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVirDomainRecorder) DestroyFlags(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DestroyFlags", arg0)
}

func (_m *MockVirDomain) ShutdownFlags(flags libvirt_go.DomainShutdownFlags) error {
	ret := _m.ctrl.Call(_m, "ShutdownFlags", flags)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVirDomainRecorder) ShutdownFlags(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ShutdownFlags", arg0)
}

func (_m *MockVirDomain) UndefineFlags(flags libvirt_go.DomainUndefineFlagsValues) error {
	ret := _m.ctrl.Call(_m, "UndefineFlags", flags)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVirDomainRecorder) UndefineFlags(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UndefineFlags", arg0)
}

func (_m *MockVirDomain) GetName() (string, error) {
	ret := _m.ctrl.Call(_m, "GetName")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVirDomainRecorder) GetName() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetName")
}

func (_m *MockVirDomain) GetUUIDString() (string, error) {
	ret := _m.ctrl.Call(_m, "GetUUIDString")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVirDomainRecorder) GetUUIDString() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetUUIDString")
}

func (_m *MockVirDomain) GetXMLDesc(flags libvirt_go.DomainXMLFlags) (string, error) {
	ret := _m.ctrl.Call(_m, "GetXMLDesc", flags)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVirDomainRecorder) GetXMLDesc(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetXMLDesc", arg0)
}

func (_m *MockVirDomain) GetMetadata(tipus libvirt_go.DomainMetadataType, uri string, flags libvirt_go.DomainModificationImpact) (string, error) {
	ret := _m.ctrl.Call(_m, "GetMetadata", tipus, uri, flags)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVirDomainRecorder) GetMetadata(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetMetadata", arg0, arg1, arg2)
}

func (_m *MockVirDomain) OpenConsole(devname string, stream *libvirt_go.Stream, flags libvirt_go.DomainConsoleFlags) error {
	ret := _m.ctrl.Call(_m, "OpenConsole", devname, stream, flags)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVirDomainRecorder) OpenConsole(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "OpenConsole", arg0, arg1, arg2)
}

func (_m *MockVirDomain) MigrateToURI3(_param0 string, _param1 *libvirt_go.DomainMigrateParameters, _param2 libvirt_go.DomainMigrateFlags) error {
	ret := _m.ctrl.Call(_m, "MigrateToURI3", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVirDomainRecorder) MigrateToURI3(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MigrateToURI3", arg0, arg1, arg2)
}

func (_m *MockVirDomain) MigrateStartPostCopy(flags uint32) error {
	ret := _m.ctrl.Call(_m, "MigrateStartPostCopy", flags)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVirDomainRecorder) MigrateStartPostCopy(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MigrateStartPostCopy", arg0)
}

func (_m *MockVirDomain) MemoryStats(nrStats uint32, flags uint32) ([]libvirt_go.DomainMemoryStat, error) {
	ret := _m.ctrl.Call(_m, "MemoryStats", nrStats, flags)
	ret0, _ := ret[0].([]libvirt_go.DomainMemoryStat)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVirDomainRecorder) MemoryStats(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MemoryStats", arg0, arg1)
}

func (_m *MockVirDomain) GetJobStats(flags libvirt_go.DomainGetJobStatsFlags) (*libvirt_go.DomainJobInfo, error) {
	ret := _m.ctrl.Call(_m, "GetJobStats", flags)
	ret0, _ := ret[0].(*libvirt_go.DomainJobInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVirDomainRecorder) GetJobStats(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetJobStats", arg0)
}

func (_m *MockVirDomain) GetJobInfo() (*libvirt_go.DomainJobInfo, error) {
	ret := _m.ctrl.Call(_m, "GetJobInfo")
	ret0, _ := ret[0].(*libvirt_go.DomainJobInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockVirDomainRecorder) GetJobInfo() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetJobInfo")
}

func (_m *MockVirDomain) SetTime(secs int64, nsecs uint, flags libvirt_go.DomainSetTimeFlags) error {
	ret := _m.ctrl.Call(_m, "SetTime", secs, nsecs, flags)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVirDomainRecorder) SetTime(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SetTime", arg0, arg1, arg2)
}

func (_m *MockVirDomain) AbortJob() error {
	ret := _m.ctrl.Call(_m, "AbortJob")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVirDomainRecorder) AbortJob() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AbortJob")
}

func (_m *MockVirDomain) Free() error {
	ret := _m.ctrl.Call(_m, "Free")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockVirDomainRecorder) Free() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Free")
}
