package events

import (
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type TestEvent struct {
	Name    string
	Payload interface{}
}

func (e *TestEvent) GetName() string {
	return e.Name
}
func (e *TestEvent) GetPayload() interface{} {
	return e.Payload
}
func (e *TestEvent) GetDate() time.Time {
	return time.Now()

}

type TestEventHandler struct {
	ID uint
}

func (h *TestEventHandler) Handle(event EventInterface) {
}

type EventDispatcherTestSuit struct {
	suite.Suite
	event1     TestEvent
	event2     TestEvent
	handle1    TestEventHandler
	handle2    TestEventHandler
	handle3    TestEventHandler
	dispatcher *EventDispatcher
}

func (suite *EventDispatcherTestSuit) SetupTest() {
	suite.dispatcher = NewEventDispatcher()
	suite.handle1 = TestEventHandler{1}
	suite.handle2 = TestEventHandler{2}
	suite.handle3 = TestEventHandler{3}
	suite.event1 = TestEvent{"Gabriel", "Matias"}
	suite.event2 = TestEvent{"Debora", "Santos"}
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuit))
}

func (suite *EventDispatcherTestSuit) TestEventDispatcher_Register() {
	err := suite.dispatcher.Register(suite.event1.GetName(), &suite.handle1)
	suite.Nil(err)
	suite.Equal(1, len(suite.dispatcher.Handlers[suite.event1.Name]))
}

func (suite *EventDispatcherTestSuit) TestEventDispatcher_Register_When_Event_aredy_exists() {
	err := suite.dispatcher.Register(suite.event1.GetName(), &suite.handle1)
	suite.Nil(err)
	err = suite.dispatcher.Register(suite.event1.GetName(), &suite.handle1)
	suite.Equal(ErrorHandlerAlredyExists, err)
}

func (suite *EventDispatcherTestSuit) TestEventDispatcher_Clear() {
	err := suite.dispatcher.Register(suite.event1.GetName(), &suite.handle1)
	suite.Nil(err)
	suite.Equal(1, len(suite.dispatcher.Handlers[suite.event1.Name]))

	suite.dispatcher.Clear()

	suite.Equal(0, len(suite.dispatcher.Handlers[suite.event1.Name]))
}

func (suite *EventDispatcherTestSuit) TestEventDispatcher_Has() {
	err := suite.dispatcher.Register(suite.event1.GetName(), &suite.handle1)
	err = suite.dispatcher.Register(suite.event1.GetName(), &suite.handle2)
	suite.Nil(err)
	suite.Equal(2, len(suite.dispatcher.Handlers[suite.event1.Name]))

	suite.Equal(true, suite.dispatcher.Has(suite.event1.GetName(), &suite.handle1))
	suite.Equal(true, suite.dispatcher.Has(suite.event1.GetName(), &suite.handle2))
	suite.Equal(false, suite.dispatcher.Has(suite.event1.GetName(), &suite.handle3))
}

type mockHandler struct {
	mock.Mock
}

func (m *mockHandler) Handle(event EventInterface) {
	m.Called(event)
}

func (suite *EventDispatcherTestSuit) TestEventDispatcher_Dispatcher() {
	eh := &mockHandler{}
	eh.On("Handle", &suite.event1)
	suite.dispatcher.Register(suite.event1.GetName(), eh)
	suite.dispatcher.Dispatcher(&suite.event1)
	eh.AssertNumberOfCalls(suite.T(), "Handle", 1)
}
