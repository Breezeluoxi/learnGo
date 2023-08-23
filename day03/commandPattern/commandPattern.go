package commandPattern

import (
	"container/list"
	"fmt"
)

// Command 接口定义了执行操作的方法
type Command interface {
	Execute()
}

// Light 是接收者，负责实际执行操作
type Light struct {
	isOn    bool
	isDelay bool
}

func (l *Light) TurnOn() {
	l.isOn = true
	fmt.Println("Light is on")
}

func (l *Light) TurnOff() {
	l.isOn = false
	fmt.Println("Light is off")
}

func (l *Light) TurnDelay() {
	l.isDelay = true
	fmt.Println("Light is delay")
}

// LightOnCommand 开灯指定
type LightOnCommand struct {
	light *Light
}

// 等待指令 LightDelayCommand
type LightDelayCommand struct {
	light *Light
}

func NewLightDelayCommand(light *Light) *LightDelayCommand {
	return &LightDelayCommand{light: light}
}

func (c *LightDelayCommand) Execute() {
	c.light.TurnDelay()
}

func NewLightOnCommand(light *Light) *LightOnCommand {
	return &LightOnCommand{light: light}
}

func (c *LightOnCommand) Execute() {
	c.light.TurnOn()
}

// LightOffCommand 关灯指令
type LightOffCommand struct {
	light *Light
}

func NewLightOffCommand(light *Light) *LightOffCommand {
	return &LightOffCommand{light: light}
}

func (c *LightOffCommand) Execute() {
	c.light.TurnOff()
}

// RemoteControl 是调用者，负责调用命令执行操作
type RemoteControl struct {
	// 队列
	commandQueue list.List
}

func (r *RemoteControl) PublishCommand(command Command) {
	r.commandQueue.PushBack(command)
}

func (r *RemoteControl) PressButton() {
	// 遍历队列，执行命令
	for e := r.commandQueue.Front(); e != nil; e = e.Next() {
		command := e.Value.(Command)
		command.Execute()
	}
}
