// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	types "github.com/alcb1310/bca-json/internals/types"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

type Service_Expecter struct {
	mock *mock.Mock
}

func (_m *Service) EXPECT() *Service_Expecter {
	return &Service_Expecter{mock: &_m.Mock}
}

// CreateCompany provides a mock function with given fields: company, user
func (_m *Service) CreateCompany(company *types.Company, user types.CreateUser) (types.User, error) {
	ret := _m.Called(company, user)

	if len(ret) == 0 {
		panic("no return value specified for CreateCompany")
	}

	var r0 types.User
	var r1 error
	if rf, ok := ret.Get(0).(func(*types.Company, types.CreateUser) (types.User, error)); ok {
		return rf(company, user)
	}
	if rf, ok := ret.Get(0).(func(*types.Company, types.CreateUser) types.User); ok {
		r0 = rf(company, user)
	} else {
		r0 = ret.Get(0).(types.User)
	}

	if rf, ok := ret.Get(1).(func(*types.Company, types.CreateUser) error); ok {
		r1 = rf(company, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Service_CreateCompany_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateCompany'
type Service_CreateCompany_Call struct {
	*mock.Call
}

// CreateCompany is a helper method to define mock.On call
//   - company *types.Company
//   - user types.CreateUser
func (_e *Service_Expecter) CreateCompany(company interface{}, user interface{}) *Service_CreateCompany_Call {
	return &Service_CreateCompany_Call{Call: _e.mock.On("CreateCompany", company, user)}
}

func (_c *Service_CreateCompany_Call) Run(run func(company *types.Company, user types.CreateUser)) *Service_CreateCompany_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*types.Company), args[1].(types.CreateUser))
	})
	return _c
}

func (_c *Service_CreateCompany_Call) Return(_a0 types.User, _a1 error) *Service_CreateCompany_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Service_CreateCompany_Call) RunAndReturn(run func(*types.Company, types.CreateUser) (types.User, error)) *Service_CreateCompany_Call {
	_c.Call.Return(run)
	return _c
}

// CreateProject provides a mock function with given fields: project
func (_m *Service) CreateProject(project types.Project) (types.Project, error) {
	ret := _m.Called(project)

	if len(ret) == 0 {
		panic("no return value specified for CreateProject")
	}

	var r0 types.Project
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Project) (types.Project, error)); ok {
		return rf(project)
	}
	if rf, ok := ret.Get(0).(func(types.Project) types.Project); ok {
		r0 = rf(project)
	} else {
		r0 = ret.Get(0).(types.Project)
	}

	if rf, ok := ret.Get(1).(func(types.Project) error); ok {
		r1 = rf(project)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Service_CreateProject_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateProject'
type Service_CreateProject_Call struct {
	*mock.Call
}

// CreateProject is a helper method to define mock.On call
//   - project types.Project
func (_e *Service_Expecter) CreateProject(project interface{}) *Service_CreateProject_Call {
	return &Service_CreateProject_Call{Call: _e.mock.On("CreateProject", project)}
}

func (_c *Service_CreateProject_Call) Run(run func(project types.Project)) *Service_CreateProject_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.Project))
	})
	return _c
}

func (_c *Service_CreateProject_Call) Return(_a0 types.Project, _a1 error) *Service_CreateProject_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Service_CreateProject_Call) RunAndReturn(run func(types.Project) (types.Project, error)) *Service_CreateProject_Call {
	_c.Call.Return(run)
	return _c
}

// CreateUser provides a mock function with given fields: user
func (_m *Service) CreateUser(user types.CreateUser) (types.User, error) {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 types.User
	var r1 error
	if rf, ok := ret.Get(0).(func(types.CreateUser) (types.User, error)); ok {
		return rf(user)
	}
	if rf, ok := ret.Get(0).(func(types.CreateUser) types.User); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Get(0).(types.User)
	}

	if rf, ok := ret.Get(1).(func(types.CreateUser) error); ok {
		r1 = rf(user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Service_CreateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateUser'
type Service_CreateUser_Call struct {
	*mock.Call
}

// CreateUser is a helper method to define mock.On call
//   - user types.CreateUser
func (_e *Service_Expecter) CreateUser(user interface{}) *Service_CreateUser_Call {
	return &Service_CreateUser_Call{Call: _e.mock.On("CreateUser", user)}
}

func (_c *Service_CreateUser_Call) Run(run func(user types.CreateUser)) *Service_CreateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.CreateUser))
	})
	return _c
}

func (_c *Service_CreateUser_Call) Return(_a0 types.User, _a1 error) *Service_CreateUser_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Service_CreateUser_Call) RunAndReturn(run func(types.CreateUser) (types.User, error)) *Service_CreateUser_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteUser provides a mock function with given fields: id, companyID
func (_m *Service) DeleteUser(id uuid.UUID, companyID uuid.UUID) error {
	ret := _m.Called(id, companyID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uuid.UUID, uuid.UUID) error); ok {
		r0 = rf(id, companyID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Service_DeleteUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteUser'
type Service_DeleteUser_Call struct {
	*mock.Call
}

// DeleteUser is a helper method to define mock.On call
//   - id uuid.UUID
//   - companyID uuid.UUID
func (_e *Service_Expecter) DeleteUser(id interface{}, companyID interface{}) *Service_DeleteUser_Call {
	return &Service_DeleteUser_Call{Call: _e.mock.On("DeleteUser", id, companyID)}
}

func (_c *Service_DeleteUser_Call) Run(run func(id uuid.UUID, companyID uuid.UUID)) *Service_DeleteUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uuid.UUID), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *Service_DeleteUser_Call) Return(_a0 error) *Service_DeleteUser_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Service_DeleteUser_Call) RunAndReturn(run func(uuid.UUID, uuid.UUID) error) *Service_DeleteUser_Call {
	_c.Call.Return(run)
	return _c
}

// GetRole provides a mock function with given fields: name
func (_m *Service) GetRole(name string) (types.Role, error) {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for GetRole")
	}

	var r0 types.Role
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (types.Role, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) types.Role); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(types.Role)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Service_GetRole_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRole'
type Service_GetRole_Call struct {
	*mock.Call
}

// GetRole is a helper method to define mock.On call
//   - name string
func (_e *Service_Expecter) GetRole(name interface{}) *Service_GetRole_Call {
	return &Service_GetRole_Call{Call: _e.mock.On("GetRole", name)}
}

func (_c *Service_GetRole_Call) Run(run func(name string)) *Service_GetRole_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Service_GetRole_Call) Return(_a0 types.Role, _a1 error) *Service_GetRole_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Service_GetRole_Call) RunAndReturn(run func(string) (types.Role, error)) *Service_GetRole_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserByID provides a mock function with given fields: id, companyID
func (_m *Service) GetUserByID(id uuid.UUID, companyID uuid.UUID) (types.User, error) {
	ret := _m.Called(id, companyID)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByID")
	}

	var r0 types.User
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID, uuid.UUID) (types.User, error)); ok {
		return rf(id, companyID)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID, uuid.UUID) types.User); ok {
		r0 = rf(id, companyID)
	} else {
		r0 = ret.Get(0).(types.User)
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID, uuid.UUID) error); ok {
		r1 = rf(id, companyID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Service_GetUserByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserByID'
type Service_GetUserByID_Call struct {
	*mock.Call
}

// GetUserByID is a helper method to define mock.On call
//   - id uuid.UUID
//   - companyID uuid.UUID
func (_e *Service_Expecter) GetUserByID(id interface{}, companyID interface{}) *Service_GetUserByID_Call {
	return &Service_GetUserByID_Call{Call: _e.mock.On("GetUserByID", id, companyID)}
}

func (_c *Service_GetUserByID_Call) Run(run func(id uuid.UUID, companyID uuid.UUID)) *Service_GetUserByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uuid.UUID), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *Service_GetUserByID_Call) Return(_a0 types.User, _a1 error) *Service_GetUserByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Service_GetUserByID_Call) RunAndReturn(run func(uuid.UUID, uuid.UUID) (types.User, error)) *Service_GetUserByID_Call {
	_c.Call.Return(run)
	return _c
}

// GetUsers provides a mock function with given fields: companyID
func (_m *Service) GetUsers(companyID uuid.UUID) ([]types.User, error) {
	ret := _m.Called(companyID)

	if len(ret) == 0 {
		panic("no return value specified for GetUsers")
	}

	var r0 []types.User
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) ([]types.User, error)); ok {
		return rf(companyID)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) []types.User); ok {
		r0 = rf(companyID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.User)
		}
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(companyID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Service_GetUsers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUsers'
type Service_GetUsers_Call struct {
	*mock.Call
}

// GetUsers is a helper method to define mock.On call
//   - companyID uuid.UUID
func (_e *Service_Expecter) GetUsers(companyID interface{}) *Service_GetUsers_Call {
	return &Service_GetUsers_Call{Call: _e.mock.On("GetUsers", companyID)}
}

func (_c *Service_GetUsers_Call) Run(run func(companyID uuid.UUID)) *Service_GetUsers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uuid.UUID))
	})
	return _c
}

func (_c *Service_GetUsers_Call) Return(_a0 []types.User, _a1 error) *Service_GetUsers_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Service_GetUsers_Call) RunAndReturn(run func(uuid.UUID) ([]types.User, error)) *Service_GetUsers_Call {
	_c.Call.Return(run)
	return _c
}

// LoadScript provides a mock function with given fields: fileName
func (_m *Service) LoadScript(fileName string) error {
	ret := _m.Called(fileName)

	if len(ret) == 0 {
		panic("no return value specified for LoadScript")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(fileName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Service_LoadScript_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LoadScript'
type Service_LoadScript_Call struct {
	*mock.Call
}

// LoadScript is a helper method to define mock.On call
//   - fileName string
func (_e *Service_Expecter) LoadScript(fileName interface{}) *Service_LoadScript_Call {
	return &Service_LoadScript_Call{Call: _e.mock.On("LoadScript", fileName)}
}

func (_c *Service_LoadScript_Call) Run(run func(fileName string)) *Service_LoadScript_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *Service_LoadScript_Call) Return(_a0 error) *Service_LoadScript_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Service_LoadScript_Call) RunAndReturn(run func(string) error) *Service_LoadScript_Call {
	_c.Call.Return(run)
	return _c
}

// Login provides a mock function with given fields: email, password
func (_m *Service) Login(email string, password string) (types.User, error) {
	ret := _m.Called(email, password)

	if len(ret) == 0 {
		panic("no return value specified for Login")
	}

	var r0 types.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (types.User, error)); ok {
		return rf(email, password)
	}
	if rf, ok := ret.Get(0).(func(string, string) types.User); ok {
		r0 = rf(email, password)
	} else {
		r0 = ret.Get(0).(types.User)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(email, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Service_Login_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Login'
type Service_Login_Call struct {
	*mock.Call
}

// Login is a helper method to define mock.On call
//   - email string
//   - password string
func (_e *Service_Expecter) Login(email interface{}, password interface{}) *Service_Login_Call {
	return &Service_Login_Call{Call: _e.mock.On("Login", email, password)}
}

func (_c *Service_Login_Call) Run(run func(email string, password string)) *Service_Login_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *Service_Login_Call) Return(_a0 types.User, _a1 error) *Service_Login_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Service_Login_Call) RunAndReturn(run func(string, string) (types.User, error)) *Service_Login_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateUser provides a mock function with given fields: user
func (_m *Service) UpdateUser(user types.CreateUser) error {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(types.CreateUser) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Service_UpdateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateUser'
type Service_UpdateUser_Call struct {
	*mock.Call
}

// UpdateUser is a helper method to define mock.On call
//   - user types.CreateUser
func (_e *Service_Expecter) UpdateUser(user interface{}) *Service_UpdateUser_Call {
	return &Service_UpdateUser_Call{Call: _e.mock.On("UpdateUser", user)}
}

func (_c *Service_UpdateUser_Call) Run(run func(user types.CreateUser)) *Service_UpdateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.CreateUser))
	})
	return _c
}

func (_c *Service_UpdateUser_Call) Return(_a0 error) *Service_UpdateUser_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *Service_UpdateUser_Call) RunAndReturn(run func(types.CreateUser) error) *Service_UpdateUser_Call {
	_c.Call.Return(run)
	return _c
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewService(t interface {
	mock.TestingT
	Cleanup(func())
}) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
