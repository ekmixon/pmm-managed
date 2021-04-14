// Code generated by mockery v1.0.0. DO NOT EDIT.

package backup

import mock "github.com/stretchr/testify/mock"

// mockS3 is an autogenerated mock type for the s3 type
type mockS3 struct {
	mock.Mock
}

// BucketExists provides a mock function with given fields: host, secure, accessKey, secretKey, name
func (_m *mockS3) BucketExists(host string, secure bool, accessKey string, secretKey string, name string) (bool, error) {
	ret := _m.Called(host, secure, accessKey, secretKey, name)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, bool, string, string, string) bool); ok {
		r0 = rf(host, secure, accessKey, secretKey, name)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, bool, string, string, string) error); ok {
		r1 = rf(host, secure, accessKey, secretKey, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBucketLocation provides a mock function with given fields: host, secure, accessKey, secretKey, name
func (_m *mockS3) GetBucketLocation(host string, secure bool, accessKey string, secretKey string, name string) (string, error) {
	ret := _m.Called(host, secure, accessKey, secretKey, name)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, bool, string, string, string) string); ok {
		r0 = rf(host, secure, accessKey, secretKey, name)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, bool, string, string, string) error); ok {
		r1 = rf(host, secure, accessKey, secretKey, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
