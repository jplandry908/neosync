// Code generated by mockery. DO NOT EDIT.

package sqlmanager

import (
	context "context"

	connectionmanager "github.com/nucleuscloud/neosync/internal/connection-manager"

	mock "github.com/stretchr/testify/mock"

	slog "log/slog"
)

// MockSqlManagerClient is an autogenerated mock type for the SqlManagerClient type
type MockSqlManagerClient struct {
	mock.Mock
}

type MockSqlManagerClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSqlManagerClient) EXPECT() *MockSqlManagerClient_Expecter {
	return &MockSqlManagerClient_Expecter{mock: &_m.Mock}
}

// NewSqlConnection provides a mock function with given fields: ctx, session, connection, slogger
func (_m *MockSqlManagerClient) NewSqlConnection(ctx context.Context, session connectionmanager.SessionInterface, connection connectionmanager.ConnectionInput, slogger *slog.Logger) (*SqlConnection, error) {
	ret := _m.Called(ctx, session, connection, slogger)

	if len(ret) == 0 {
		panic("no return value specified for NewSqlConnection")
	}

	var r0 *SqlConnection
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, connectionmanager.SessionInterface, connectionmanager.ConnectionInput, *slog.Logger) (*SqlConnection, error)); ok {
		return rf(ctx, session, connection, slogger)
	}
	if rf, ok := ret.Get(0).(func(context.Context, connectionmanager.SessionInterface, connectionmanager.ConnectionInput, *slog.Logger) *SqlConnection); ok {
		r0 = rf(ctx, session, connection, slogger)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*SqlConnection)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, connectionmanager.SessionInterface, connectionmanager.ConnectionInput, *slog.Logger) error); ok {
		r1 = rf(ctx, session, connection, slogger)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockSqlManagerClient_NewSqlConnection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NewSqlConnection'
type MockSqlManagerClient_NewSqlConnection_Call struct {
	*mock.Call
}

// NewSqlConnection is a helper method to define mock.On call
//   - ctx context.Context
//   - session connectionmanager.SessionInterface
//   - connection connectionmanager.ConnectionInput
//   - slogger *slog.Logger
func (_e *MockSqlManagerClient_Expecter) NewSqlConnection(ctx interface{}, session interface{}, connection interface{}, slogger interface{}) *MockSqlManagerClient_NewSqlConnection_Call {
	return &MockSqlManagerClient_NewSqlConnection_Call{Call: _e.mock.On("NewSqlConnection", ctx, session, connection, slogger)}
}

func (_c *MockSqlManagerClient_NewSqlConnection_Call) Run(run func(ctx context.Context, session connectionmanager.SessionInterface, connection connectionmanager.ConnectionInput, slogger *slog.Logger)) *MockSqlManagerClient_NewSqlConnection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(connectionmanager.SessionInterface), args[2].(connectionmanager.ConnectionInput), args[3].(*slog.Logger))
	})
	return _c
}

func (_c *MockSqlManagerClient_NewSqlConnection_Call) Return(_a0 *SqlConnection, _a1 error) *MockSqlManagerClient_NewSqlConnection_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockSqlManagerClient_NewSqlConnection_Call) RunAndReturn(run func(context.Context, connectionmanager.SessionInterface, connectionmanager.ConnectionInput, *slog.Logger) (*SqlConnection, error)) *MockSqlManagerClient_NewSqlConnection_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSqlManagerClient creates a new instance of MockSqlManagerClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSqlManagerClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSqlManagerClient {
	mock := &MockSqlManagerClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
