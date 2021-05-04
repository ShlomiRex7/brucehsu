package main

type CallFrame struct {
	parent    *CallFrame
	var_table map[string]Object
	stack     []Object
	me        Object
}

type GobiesVM struct {
	instList       []*Instruction
	callFrameStack []*CallFrame
	consts         map[string]Object
	symbols        map[string]int
}

func initVM() *GobiesVM {
	VM := &GobiesVM{}
	top := initCallFrame()
	VM.callFrameStack = append(VM.callFrameStack, top)
	top.me = initRKernel()
	return VM
}

func initCallFrame() *CallFrame {
	frame := &CallFrame{}
	frame.var_table = make(map[string]Object)
	return frame
}

func (VM *GobiesVM) executeBytecode() {
	for _, v := range VM.instList {
		currentCallFrame := VM.callFrameStack[len(VM.callFrameStack)-1]
		switch v.inst_type {
		case BC_PUTSELF:
			currentCallFrame.stack = append(currentCallFrame.stack, currentCallFrame.me)
		case BC_PUTNIL:
			currentCallFrame.stack = append(currentCallFrame.stack, nil)
		case BC_PUTOBJ:
			currentCallFrame.stack = append(currentCallFrame.stack, v.obj)
		case BC_PUTTRUE:
		case BC_PUTFALSE:
		case BC_SETLOCAL:
			top := currentCallFrame.stack[len(currentCallFrame.stack)-1]
			currentCallFrame.var_table[v.obj.getString()] = top
			currentCallFrame.stack = currentCallFrame.stack[0 : len(currentCallFrame.stack)-1]
		case BC_GETLOCAL:
			currentCallFrame.stack = append(currentCallFrame.stack, currentCallFrame.var_table[v.obj.getString()])
		case BC_SETGLOBAL:
		case BC_GETGLOBAL:
		case BC_SETSYMBOL:
		case BC_GETSYMBOL:
		case BC_SETCONST:
		case BC_GETCONST:
		case BC_SETIVAR:
		case BC_GETIVAR:
		case BC_SETCVAR:
		case BC_GETCVAR:
		case BC_SEND:
			argLists := currentCallFrame.stack[len(currentCallFrame.stack)-(v.argc+1):]
			currentCallFrame.stack = currentCallFrame.stack[:len(currentCallFrame.stack)-(v.argc+1)]
			recv := argLists[0]
			argLists = argLists[1:]
			recv.getMethods()[v.obj.getString()].gofunc(VM, recv, argLists)
		case BC_JUMP:
		}
	}
}