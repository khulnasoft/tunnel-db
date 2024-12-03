// Code generated by mockery v1.0.0. DO NOT EDIT.

package github

import context "context"
import mock "github.com/stretchr/testify/mock"
import os "os"
import v38github "github.com/google/go-github/v38/github"

// MockRepositoryInterface is an autogenerated mock type for the RepositoryInterface type
type MockRepositoryInterface struct {
	mock.Mock
}

type CreateReleaseArgs struct {
	Ctx             context.Context
	CtxAnything     bool
	Release         *v38github.RepositoryRelease
	ReleaseAnything bool
}

type CreateReleaseReturns struct {
	Result   *v38github.RepositoryRelease
	Response *v38github.Response
	Err      error
}

type CreateReleaseExpectation struct {
	Args    CreateReleaseArgs
	Returns CreateReleaseReturns
}

func (_m *MockRepositoryInterface) ApplyCreateReleaseExpectation(e CreateReleaseExpectation) {
	var args []interface{}
	if e.Args.CtxAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Ctx)
	}
	if e.Args.ReleaseAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Release)
	}
	_m.On("CreateRelease", args...).Return(e.Returns.Result, e.Returns.Response, e.Returns.Err)
}

func (_m *MockRepositoryInterface) ApplyCreateReleaseExpectations(expectations []CreateReleaseExpectation) {
	for _, e := range expectations {
		_m.ApplyCreateReleaseExpectation(e)
	}
}

// CreateRelease provides a mock function with given fields: ctx, release
func (_m *MockRepositoryInterface) CreateRelease(ctx context.Context, release *v38github.RepositoryRelease) (*v38github.RepositoryRelease, *v38github.Response, error) {
	ret := _m.Called(ctx, release)

	var r0 *v38github.RepositoryRelease
	if rf, ok := ret.Get(0).(func(context.Context, *v38github.RepositoryRelease) *v38github.RepositoryRelease); ok {
		r0 = rf(ctx, release)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v38github.RepositoryRelease)
		}
	}

	var r1 *v38github.Response
	if rf, ok := ret.Get(1).(func(context.Context, *v38github.RepositoryRelease) *v38github.Response); ok {
		r1 = rf(ctx, release)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*v38github.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *v38github.RepositoryRelease) error); ok {
		r2 = rf(ctx, release)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

type DeleteRefArgs struct {
	Ctx         context.Context
	CtxAnything bool
	Ref         string
	RefAnything bool
}

type DeleteRefReturns struct {
	Res *v38github.Response
	Err error
}

type DeleteRefExpectation struct {
	Args    DeleteRefArgs
	Returns DeleteRefReturns
}

func (_m *MockRepositoryInterface) ApplyDeleteRefExpectation(e DeleteRefExpectation) {
	var args []interface{}
	if e.Args.CtxAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Ctx)
	}
	if e.Args.RefAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Ref)
	}
	_m.On("DeleteRef", args...).Return(e.Returns.Res, e.Returns.Err)
}

func (_m *MockRepositoryInterface) ApplyDeleteRefExpectations(expectations []DeleteRefExpectation) {
	for _, e := range expectations {
		_m.ApplyDeleteRefExpectation(e)
	}
}

// DeleteRef provides a mock function with given fields: ctx, ref
func (_m *MockRepositoryInterface) DeleteRef(ctx context.Context, ref string) (*v38github.Response, error) {
	ret := _m.Called(ctx, ref)

	var r0 *v38github.Response
	if rf, ok := ret.Get(0).(func(context.Context, string) *v38github.Response); ok {
		r0 = rf(ctx, ref)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v38github.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, ref)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type DeleteReleaseArgs struct {
	Ctx         context.Context
	CtxAnything bool
	Id          int64
	IdAnything  bool
}

type DeleteReleaseReturns struct {
	Res *v38github.Response
	Err error
}

type DeleteReleaseExpectation struct {
	Args    DeleteReleaseArgs
	Returns DeleteReleaseReturns
}

func (_m *MockRepositoryInterface) ApplyDeleteReleaseExpectation(e DeleteReleaseExpectation) {
	var args []interface{}
	if e.Args.CtxAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Ctx)
	}
	if e.Args.IdAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Id)
	}
	_m.On("DeleteRelease", args...).Return(e.Returns.Res, e.Returns.Err)
}

func (_m *MockRepositoryInterface) ApplyDeleteReleaseExpectations(expectations []DeleteReleaseExpectation) {
	for _, e := range expectations {
		_m.ApplyDeleteReleaseExpectation(e)
	}
}

// DeleteRelease provides a mock function with given fields: ctx, id
func (_m *MockRepositoryInterface) DeleteRelease(ctx context.Context, id int64) (*v38github.Response, error) {
	ret := _m.Called(ctx, id)

	var r0 *v38github.Response
	if rf, ok := ret.Get(0).(func(context.Context, int64) *v38github.Response); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v38github.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type GetReleaseByTagArgs struct {
	Ctx         context.Context
	CtxAnything bool
	Tag         string
	TagAnything bool
}

type GetReleaseByTagReturns struct {
	Release  *v38github.RepositoryRelease
	Response *v38github.Response
	Err      error
}

type GetReleaseByTagExpectation struct {
	Args    GetReleaseByTagArgs
	Returns GetReleaseByTagReturns
}

func (_m *MockRepositoryInterface) ApplyGetReleaseByTagExpectation(e GetReleaseByTagExpectation) {
	var args []interface{}
	if e.Args.CtxAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Ctx)
	}
	if e.Args.TagAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Tag)
	}
	_m.On("GetReleaseByTag", args...).Return(e.Returns.Release, e.Returns.Response, e.Returns.Err)
}

func (_m *MockRepositoryInterface) ApplyGetReleaseByTagExpectations(expectations []GetReleaseByTagExpectation) {
	for _, e := range expectations {
		_m.ApplyGetReleaseByTagExpectation(e)
	}
}

// GetReleaseByTag provides a mock function with given fields: ctx, tag
func (_m *MockRepositoryInterface) GetReleaseByTag(ctx context.Context, tag string) (*v38github.RepositoryRelease, *v38github.Response, error) {
	ret := _m.Called(ctx, tag)

	var r0 *v38github.RepositoryRelease
	if rf, ok := ret.Get(0).(func(context.Context, string) *v38github.RepositoryRelease); ok {
		r0 = rf(ctx, tag)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v38github.RepositoryRelease)
		}
	}

	var r1 *v38github.Response
	if rf, ok := ret.Get(1).(func(context.Context, string) *v38github.Response); ok {
		r1 = rf(ctx, tag)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*v38github.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, tag)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

type ListReleasesArgs struct {
	Ctx         context.Context
	CtxAnything bool
	Opt         *v38github.ListOptions
	OptAnything bool
}

type ListReleasesReturns struct {
	Releases []*v38github.RepositoryRelease
	Response *v38github.Response
	Err      error
}

type ListReleasesExpectation struct {
	Args    ListReleasesArgs
	Returns ListReleasesReturns
}

func (_m *MockRepositoryInterface) ApplyListReleasesExpectation(e ListReleasesExpectation) {
	var args []interface{}
	if e.Args.CtxAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Ctx)
	}
	if e.Args.OptAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Opt)
	}
	_m.On("ListReleases", args...).Return(e.Returns.Releases, e.Returns.Response, e.Returns.Err)
}

func (_m *MockRepositoryInterface) ApplyListReleasesExpectations(expectations []ListReleasesExpectation) {
	for _, e := range expectations {
		_m.ApplyListReleasesExpectation(e)
	}
}

// ListReleases provides a mock function with given fields: ctx, opt
func (_m *MockRepositoryInterface) ListReleases(ctx context.Context, opt *v38github.ListOptions) ([]*v38github.RepositoryRelease, *v38github.Response, error) {
	ret := _m.Called(ctx, opt)

	var r0 []*v38github.RepositoryRelease
	if rf, ok := ret.Get(0).(func(context.Context, *v38github.ListOptions) []*v38github.RepositoryRelease); ok {
		r0 = rf(ctx, opt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v38github.RepositoryRelease)
		}
	}

	var r1 *v38github.Response
	if rf, ok := ret.Get(1).(func(context.Context, *v38github.ListOptions) *v38github.Response); ok {
		r1 = rf(ctx, opt)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*v38github.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, *v38github.ListOptions) error); ok {
		r2 = rf(ctx, opt)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

type UploadReleaseAssetArgs struct {
	Ctx          context.Context
	CtxAnything  bool
	Id           int64
	IdAnything   bool
	Opt          *v38github.UploadOptions
	OptAnything  bool
	File         *os.File
	FileAnything bool
}

type UploadReleaseAssetReturns struct {
	Asset    *v38github.ReleaseAsset
	Response *v38github.Response
	Err      error
}

type UploadReleaseAssetExpectation struct {
	Args    UploadReleaseAssetArgs
	Returns UploadReleaseAssetReturns
}

func (_m *MockRepositoryInterface) ApplyUploadReleaseAssetExpectation(e UploadReleaseAssetExpectation) {
	var args []interface{}
	if e.Args.CtxAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Ctx)
	}
	if e.Args.IdAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Id)
	}
	if e.Args.OptAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.Opt)
	}
	if e.Args.FileAnything {
		args = append(args, mock.Anything)
	} else {
		args = append(args, e.Args.File)
	}
	_m.On("UploadReleaseAsset", args...).Return(e.Returns.Asset, e.Returns.Response, e.Returns.Err)
}

func (_m *MockRepositoryInterface) ApplyUploadReleaseAssetExpectations(expectations []UploadReleaseAssetExpectation) {
	for _, e := range expectations {
		_m.ApplyUploadReleaseAssetExpectation(e)
	}
}

// UploadReleaseAsset provides a mock function with given fields: ctx, id, opt, file
func (_m *MockRepositoryInterface) UploadReleaseAsset(ctx context.Context, id int64, opt *v38github.UploadOptions, file *os.File) (*v38github.ReleaseAsset, *v38github.Response, error) {
	ret := _m.Called(ctx, id, opt, file)

	var r0 *v38github.ReleaseAsset
	if rf, ok := ret.Get(0).(func(context.Context, int64, *v38github.UploadOptions, *os.File) *v38github.ReleaseAsset); ok {
		r0 = rf(ctx, id, opt, file)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v38github.ReleaseAsset)
		}
	}

	var r1 *v38github.Response
	if rf, ok := ret.Get(1).(func(context.Context, int64, *v38github.UploadOptions, *os.File) *v38github.Response); ok {
		r1 = rf(ctx, id, opt, file)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*v38github.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, int64, *v38github.UploadOptions, *os.File) error); ok {
		r2 = rf(ctx, id, opt, file)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
