// Code generated by mockery v1.0.0. DO NOT EDIT.

package gitlab

import (
	mock "github.com/stretchr/testify/mock"
	"github.com/xanzy/go-gitlab"
)

// mockGoGitlabProxy is an autogenerated mock type for the goGitlabProxy type
type mockGoGitlabProxy struct {
	mock.Mock
}

// AddProjectHook provides a mock function with given fields: projectNameWithNamespace, options
func (_m *mockGoGitlabProxy) AddProjectHook(projectNameWithNamespace string, options *gitlab.AddProjectHookOptions) error {
	ret := _m.Called(projectNameWithNamespace, options)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *gitlab.AddProjectHookOptions) error); ok {
		r0 = rf(projectNameWithNamespace, options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListGroupProjects provides a mock function with given fields: owner, options
func (_m *mockGoGitlabProxy) ListGroupProjects(owner string, options *gitlab.ListGroupProjectsOptions) ([]*gitlab.Project, *gitlab.Response, error) {
	ret := _m.Called(owner, options)

	var r0 []*gitlab.Project
	if rf, ok := ret.Get(0).(func(string, *gitlab.ListGroupProjectsOptions) []*gitlab.Project); ok {
		r0 = rf(owner, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*gitlab.Project)
		}
	}

	var r1 *gitlab.Response
	if rf, ok := ret.Get(1).(func(string, *gitlab.ListGroupProjectsOptions) *gitlab.Response); ok {
		r1 = rf(owner, options)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, *gitlab.ListGroupProjectsOptions) error); ok {
		r2 = rf(owner, options)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListProjectHooks provides a mock function with given fields: projectNameWithNamespace, options
func (_m *mockGoGitlabProxy) ListProjectHooks(projectNameWithNamespace string, options *gitlab.ListProjectHooksOptions) ([]*gitlab.ProjectHook, *gitlab.Response, error) {
	ret := _m.Called(projectNameWithNamespace, options)

	var r0 []*gitlab.ProjectHook
	if rf, ok := ret.Get(0).(func(string, *gitlab.ListProjectHooksOptions) []*gitlab.ProjectHook); ok {
		r0 = rf(projectNameWithNamespace, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*gitlab.ProjectHook)
		}
	}

	var r1 *gitlab.Response
	if rf, ok := ret.Get(1).(func(string, *gitlab.ListProjectHooksOptions) *gitlab.Response); ok {
		r1 = rf(projectNameWithNamespace, options)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, *gitlab.ListProjectHooksOptions) error); ok {
		r2 = rf(projectNameWithNamespace, options)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
