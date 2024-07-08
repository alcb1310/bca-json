// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// Handler is an autogenerated mock type for the Handler type
type Handler struct {
	mock.Mock
}

type Handler_Expecter struct {
	mock *mock.Mock
}

func (_m *Handler) EXPECT() *Handler_Expecter {
	return &Handler_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: w, r
func (_m *Handler) Execute(w http.ResponseWriter, r *http.Request) error {
	ret := _m.Called(w, r)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(http.ResponseWriter, *http.Request) error); ok {
		r0 = rf(w, r)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Handler_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type Handler_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - w http.ResponseWriter
//   - r *http.Request
func (_e *Handler_Expecter) Execute(w interface{}, r interface{}) *Handler_Execute_Call {
	return &Handler_Execute_Call{Call: _e.mock.On("Execute", w, r)}
}

func (_c *Handler_Execute_Call) Run(run func(w http.ResponseWriter, r *http.Request)) *Handler_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(http.ResponseWriter), args[1].(*http.Request))
	})
	return _c
}

func (_c *Handler_Execute_Call) Return(_a0 error) *Handler_Execute_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Handler_Execute_Call) RunAndReturn(run func(http.ResponseWriter, *http.Request) error) *Handler_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewHandler creates a new instance of Handler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *Handler {
	mock := &Handler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
