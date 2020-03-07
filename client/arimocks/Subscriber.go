// Code generated by mockery v1.0.0. DO NOT EDIT.

package arimocks

import (
	ari "github.com/CyCoreSystems/ari/v5"
	mock "github.com/stretchr/testify/mock"
)

// Subscriber is an autogenerated mock type for the Subscriber type
type Subscriber struct {
	mock.Mock
}

// Subscribe provides a mock function with given fields: key, n
func (_m *Subscriber) Subscribe(key *ari.Key, n ...string) ari.Subscription {
	_va := make([]interface{}, len(n))
	for _i := range n {
		_va[_i] = n[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 ari.Subscription
	if rf, ok := ret.Get(0).(func(*ari.Key, ...string) ari.Subscription); ok {
		r0 = rf(key, n...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ari.Subscription)
		}
	}

	return r0
}
