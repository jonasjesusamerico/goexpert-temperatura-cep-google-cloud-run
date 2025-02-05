// Code generated by mockery v2.43.0. DO NOT EDIT.

package mock

import (
	domain "goexpert-temperature-system-by-cep/internal/domain"

	mock "github.com/stretchr/testify/mock"
)

// CepService is an autogenerated mock type for the CepService type
type CepService struct {
	mock.Mock
}

type ZipCodeService_Expecter struct {
	mock *mock.Mock
}

func (_m *CepService) EXPECT() *ZipCodeService_Expecter {
	return &ZipCodeService_Expecter{mock: &_m.Mock}
}

// BuscaLocalizacaoPorCep provides a mock function with given fields: zipCode
func (_m *CepService) BuscaLocalizacaoPorCep(zipCode string) (*domain.Localizacao, error) {
	ret := _m.Called(zipCode)

	if len(ret) == 0 {
		panic("no return value specified for BuscaLocalizacaoPorCep")
	}

	var r0 *domain.Localizacao
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.Localizacao, error)); ok {
		return rf(zipCode)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.Localizacao); ok {
		r0 = rf(zipCode)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Localizacao)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(zipCode)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ZipCodeService_GetLocationByZipCode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BuscaLocalizacaoPorCep'
type ZipCodeService_GetLocationByZipCode_Call struct {
	*mock.Call
}

// BuscaLocalizacaoPorCep is a helper method to define mock.On call
//   - zipCode string
func (_e *ZipCodeService_Expecter) BuscaLocalizacaoPorCep(zipCode interface{}) *ZipCodeService_GetLocationByZipCode_Call {
	return &ZipCodeService_GetLocationByZipCode_Call{Call: _e.mock.On("BuscaLocalizacaoPorCep", zipCode)}
}

func (_c *ZipCodeService_GetLocationByZipCode_Call) Run(run func(zipCode string)) *ZipCodeService_GetLocationByZipCode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *ZipCodeService_GetLocationByZipCode_Call) Return(_a0 *domain.Localizacao, _a1 error) *ZipCodeService_GetLocationByZipCode_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ZipCodeService_GetLocationByZipCode_Call) RunAndReturn(run func(string) (*domain.Localizacao, error)) *ZipCodeService_GetLocationByZipCode_Call {
	_c.Call.Return(run)
	return _c
}

// NewZipCodeService creates a new instance of CepService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewZipCodeService(t interface {
	mock.TestingT
	Cleanup(func())
}) *CepService {
	mock := &CepService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
