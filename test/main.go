package main

// 抽象类型工厂
type TypeFactory interface {
	NewCtrl() Ctrl
	NewSvc() Svc
	NewDao() Dao
}

type Ctrl interface{ Add() }
type Svc interface{ Add() }
type Dao interface{ Add() }

type UserCtrl struct{}

func (u *UserCtrl) Add() {}

type UserSvc struct{}

func (u *UserSvc) Add() {}

type UserDao struct{}

func (u *UserDao) Add() {}

type UserFactory struct{}

func (f *UserFactory) NewCtrl() Ctrl { return &UserCtrl{} }
func (f *UserFactory) NewSvc() Svc   { return &UserSvc{} }
func (f *UserFactory) NewDao() Dao   { return &UserDao{} }
