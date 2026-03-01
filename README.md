# event

`event` 是 infrago 的模块包。

## 安装

```bash
go get github.com/infrago/event@latest
```

## 最小接入

```go
package main

import (
    _ "github.com/infrago/event"
    "github.com/infrago/infra"
)

func main() {
    infra.Run()
}
```

## 配置示例

```toml
[event]
driver = "default"
```

## 公开 API（摘自源码）

- `func (Event) RegistryComponent() string`
- `func (Events) RegistryComponent() string`
- `func (ctx *Context) Next()`
- `func (ctx *Context) Found()`
- `func (ctx *Context) Error(res Res)`
- `func (ctx *Context) Failed(res Res)`
- `func (ctx *Context) Denied(res Res)`
- `func (d *defaultDriver) Connect(inst *Instance) (Connection, error)`
- `func (c *defaultConnection) Open() error  { return nil }`
- `func (c *defaultConnection) Close() error { return nil }`
- `func (c *defaultConnection) Register(name, _ string) error`
- `func (c *defaultConnection) Start() error`
- `func (c *defaultConnection) Stop() error`
- `func (c *defaultConnection) Publish(name string, data []byte) error`
- `func Publish(name string, values ...Map) error`
- `func PublishTo(conn, name string, values ...Map) error`
- `func Broadcast(name string, values ...Map) error`
- `func BroadcastTo(conn, name string, values ...Map) error`
- `func RegisterDriver(name string, driver Driver)`
- `func RegisterConfig(name string, cfg Config)`
- `func RegisterConfigs(cfgs Configs)`
- `func (m *Module) RegisterEvent(name string, cfg Event)`
- `func (m *Module) RegisterDeclare(name string, cfg Declare)`
- `func (m *Module) RegisterFilter(name string, cfg Filter)`
- `func (m *Module) RegisterHandler(name string, cfg Handler)`
- `func (m *Module) Register(name string, value Any)`
- `func (m *Module) RegisterEvents(prefix string, events Events)`
- `func (m *Module) RegisterDriver(name string, driver Driver)`
- `func (m *Module) RegisterConfig(name string, cfg Config)`
- `func (m *Module) RegisterConfigs(configs Configs)`
- `func (m *Module) Config(global Map)`
- `func (m *Module) Setup()`
- `func (m *Module) Open()`
- `func (m *Module) Start()`
- `func (m *Module) Stop()`
- `func (m *Module) Close()`
- `func (inst *Instance) Submit(next func())`
- `func (inst *Instance) Serve(name string, data []byte)`

## 排错

- 模块未运行：确认空导入已存在
- driver 无效：确认驱动包已引入
- 配置不生效：检查配置段名是否为 `[event]`
