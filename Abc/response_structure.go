package Abc

func responseStructure(isSuccess bool, data string) string {

	if isSuccess {
		return "Error: null" + "\n" + "Result: " + data
	}
	return "Error: " + data + "\n" + "Result: null"
}
