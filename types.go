package main

type Result interface {
	OK() any
	Err() error
}

type result struct {
	Ok    any
	Errer error
}

func (r *result) OK() any {
	return r.Ok
}

func (r *result) Err() error {
	return r.Errer
}

func OK(t any) Result {
	return &result{Ok: t, Errer: nil}
}

func Err(err error) Result {
	return &result{Ok: struct{}{}, Errer: err}
}
