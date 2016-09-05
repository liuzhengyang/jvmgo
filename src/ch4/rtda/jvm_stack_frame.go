package rtda

type Frame struct {
	//cp classfile.ConstantPool  // 指向常量池的指针
	lower *Frame
	operandStack *OperandStack
	localVars    LocalVars
}

func NewFrame(maxStack uint, maxLocals uint) *Frame {
	return &Frame{operandStack:newOperandStack(maxStack),
	localVars:newLocalVar(maxLocals)}
}

func (this *Frame) LocalVars() LocalVars {
	return this.localVars
}

func (this *Frame) OperandStack() *OperandStack {
	return this.operandStack;
}