package tps

var (
	Void       = NewPrimitive("V")
	Int        = NewPrimitive("I")
	Double     = NewPrimitive("D")
	Long       = NewPrimitive("J")
	String     = NewReference("java/lang/String")
	Object     = NewReference("java/lang/Object")
	MainMethod = NewMethod(Void, String)
)
