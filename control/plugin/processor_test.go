package plugin

import (
	"log"
	"testing"
	"time"

	"github.com/intelsdi-x/pulse/control/plugin/cpolicy"
	"github.com/intelsdi-x/pulse/core/ctypes"

	. "github.com/smartystreets/goconvey/convey"
)

type MockProcessor struct {
	Meta PluginMeta
}

func (f *MockProcessor) Process(contentType string, content []byte, config map[string]ctypes.ConfigValue) (string, []byte, error) {
	return "", nil, nil
}

func (f *MockProcessor) GetConfigPolicyNode() cpolicy.ConfigPolicyNode {
	return cpolicy.ConfigPolicyNode{}
}

type MockProcessorSessionState struct {
	PingTimeoutDuration time.Duration
	Daemon              bool
	listenAddress       string
	listenPort          string
	token               string
	logger              *log.Logger
	killChan            chan int
}

func (s *MockProcessorSessionState) Ping(arg PingArgs, b *bool) error {
	return nil
}

func (s *MockProcessorSessionState) Kill(arg KillArgs, b *bool) error {
	s.killChan <- 0
	return nil
}

func (s *MockProcessorSessionState) Logger() *log.Logger {
	return s.logger
}

func (s *MockProcessorSessionState) ListenAddress() string {
	return s.listenAddress
}

func (s *MockProcessorSessionState) ListenPort() string {
	return s.listenPort
}

func (s *MockProcessorSessionState) SetListenAddress(a string) {
	s.listenAddress = a
}

func (s *MockProcessorSessionState) Token() string {
	return s.token
}

func (m *MockProcessorSessionState) ResetHeartbeat() {

}

func (s *MockProcessorSessionState) KillChan() chan int {
	return s.killChan
}

func (s *MockProcessorSessionState) isDaemon() bool {
	return !s.Daemon
}

func (s *MockProcessorSessionState) generateResponse(r *Response) []byte {
	return []byte("mockResponse")
}

func (s *MockProcessorSessionState) heartbeatWatch(killChan chan int) {
	time.Sleep(time.Millisecond * 200)
	killChan <- 0
}

func TestStartProcessor(t *testing.T) {
	Convey("Processor", t, func() {
		Convey("start with dynamic port", func() {
			c := new(MockProcessor)
			m := &PluginMeta{
				RPCType: JSONRPC,
				Type:    ProcessorPluginType,
			}
			// we will panic since rpc.HandleHttp has already
			// been called during TestStartCollector
			Convey("RPC service already registered", func() {
				So(func() { Start(m, c, "{}") }, ShouldPanic)
			})

		})
	})
}