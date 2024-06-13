package events

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TestEvent struct {
	Name    string
	PayLoad interface{}
}

func (e *TestEvent) GetName() string {
	return e.Name
}

func (e *TestEvent) GetPayLoad() interface{} {
	return e.PayLoad
}

func (e *TestEvent) GetDateTime() time.Time {
	return time.Now()
}

type TestEventHandler struct {
	ID int
}

func (h *TestEventHandler) Handle(event EventInterface, wg *sync.WaitGroup) {

}

type EventDispatcherTestSuite struct {
	suite.Suite
	event           TestEvent
	event2          TestEvent
	handler         TestEventHandler
	handler2        TestEventHandler
	handler3        TestEventHandler
	eventDispatcher *EventDispatcher
}

func (suite *EventDispatcherTestSuite) SetupTest() {
	suite.eventDispatcher = NewEventDispatcher()
	suite.handler = TestEventHandler{ID: 1}
	suite.handler2 = TestEventHandler{ID: 2}
	suite.handler3 = TestEventHandler{ID: 3}
	suite.event = TestEvent{Name: "Name 1", PayLoad: "PayLoad 1"}
	suite.event2 = TestEvent{Name: "Name 2", PayLoad: "PayLoad 2"}
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register() {
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	assert.Equal(suite.T(), &suite.handler, suite.eventDispatcher.handlers[suite.event.GetName()][0])
	assert.Equal(suite.T(), &suite.handler2, suite.eventDispatcher.handlers[suite.event.GetName()][1])
}

func (suite *EventDispatcherTestSuite) TestEventDispatchar_Register_WithSameHandler() {
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Equal(ErrHandleAlreadyRegister, err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Clear() {
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Register(suite.event2.GetName(), &suite.handler3)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event2.GetName()]))

	suite.eventDispatcher.Clear()
	suite.Equal(0, len(suite.eventDispatcher.handlers))

}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Test() {
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	suite.Equal(true, suite.eventDispatcher.Has(suite.event.GetName(), &suite.handler))

	// ou...
	assert.False(suite.T(), suite.eventDispatcher.Has(suite.event.GetName(), &suite.handler2))
}

type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) Handle(event EventInterface, wg *sync.WaitGroup) {
	m.Called(event)
	wg.Done()
}

func (suite *EventDispatcherTestSuite) TestEventDispacher_Dispatch() {
	eh := &MockHandler{}
	eh.On("Handle", &suite.event)

	eh2 := &MockHandler{}
	eh2.On("Handle", &suite.event)

	suite.eventDispatcher.Register(suite.event.GetName(), eh)
	suite.eventDispatcher.Register(suite.event.GetName(), eh2)

	suite.eventDispatcher.Dispatch(&suite.event)

	eh.AssertExpectations(suite.T())
	eh2.AssertExpectations(suite.T())

	eh.AssertNumberOfCalls(suite.T(), "Handle", 1)
	eh2.AssertNumberOfCalls(suite.T(), "Handle", 1)
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Remove() {
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Remove(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))
	suite.Equal(&suite.handler2, suite.eventDispatcher.handlers[suite.event.GetName()][0])

	err = suite.eventDispatcher.Remove(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(0, len(suite.eventDispatcher.handlers[suite.event.GetName()]))
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}
