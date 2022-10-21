package lang

import "go-jvm/ch11-2/instructions/base"
import "go-jvm/ch11-2/runtime_data_area"
import "go-jvm/ch11-2/runtime_data_area/heap"

/*
Field(Class<?> declaringClass,

	String name,
	Class<?> type,
	int modifiers,
	int slot,
	String signature,
	byte[] annotations)
*/
const _fieldConstructorDescriptor = "" +
	"(Ljava/lang/Class;" +
	"Ljava/lang/String;" +
	"Ljava/lang/Class;" +
	"II" +
	"Ljava/lang/String;" +
	"[B)V"

// private native Field[] getDeclaredFields0(boolean publicOnly);
// (Z)[Ljava/lang/reflect/Field;
func getDeclaredFields0(frame *runtime_data_area.Frame) {
	vars := frame.LocalVars()
	classObj := vars.GetThis()
	publicOnly := vars.GetBoolean(1)

	class := classObj.Extra().(*heap.Class)
	fields := class.GetFields(publicOnly)
	fieldCount := uint(len(fields))

	classLoader := frame.Method().Class().Loader()
	fieldClass := classLoader.LoadClass("java/lang/reflect/Field")
	fieldArr := fieldClass.ArrayClass().NewArray(fieldCount)

	stack := frame.OperandStack()
	stack.PushRef(fieldArr)

	if fieldCount > 0 {
		thread := frame.Thread()
		fieldObjs := fieldArr.Refs()
		fieldConstructor := fieldClass.GetConstructor(_fieldConstructorDescriptor)
		for i, goField := range fields {
			fieldObj := fieldClass.NewObject()
			fieldObj.SetExtra(goField)
			fieldObjs[i] = fieldObj

			ops := runtime_data_area.NewOperandStack(8)
			ops.PushRef(fieldObj)                                          // this
			ops.PushRef(classObj)                                          // declaringClass
			ops.PushRef(heap.JString(classLoader, goField.Name()))         // name
			ops.PushRef(goField.Type().JClass())                           // type
			ops.PushInt(int32(goField.AccessFlags()))                      // modifiers
			ops.PushInt(int32(goField.SlotId()))                           // slot
			ops.PushRef(getSignatureStr(classLoader, goField.Signature())) // signature
			ops.PushRef(toByteArr(classLoader, goField.AnnotationData()))  // annotations

			shimFrame := runtime_data_area.NewShimFrame(thread, ops)
			thread.PushFrame(shimFrame)

			// init fieldObj
			base.InvokeMethod(shimFrame, fieldConstructor)
		}
	}
}
