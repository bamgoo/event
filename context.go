package event

import (
	"github.com/infrago/infra"
	. "github.com/infrago/base"
)

type (
	Context struct {
		inst *Instance
		*infra.Meta

		index int
		nexts []ctxFunc

		Name    string
		Config  *Event
		Setting Map

		Value  Map
		Args   Map
		Locals Map
		Body   Any
	}

	ctxFunc func(*Context)
)

func (ctx *Context) clear() {
	ctx.index = 0
	ctx.nexts = make([]ctxFunc, 0)
}

func (ctx *Context) next(nexts ...ctxFunc) {
	ctx.nexts = append(ctx.nexts, nexts...)
}

func (ctx *Context) Next() {
	if len(ctx.nexts) > ctx.index {
		next := ctx.nexts[ctx.index]
		ctx.index++
		if next != nil {
			next(ctx)
		} else {
			ctx.Next()
		}
	}
}

func (ctx *Context) Found() {
	ctx.inst.found(ctx)
}

func (ctx *Context) Error(res Res) {
	ctx.Result(res)
	ctx.inst.error(ctx)
}

func (ctx *Context) Failed(res Res) {
	ctx.Result(res)
	ctx.inst.failed(ctx)
}

func (ctx *Context) Denied(res Res) {
	ctx.Result(res)
	ctx.inst.denied(ctx)
}
