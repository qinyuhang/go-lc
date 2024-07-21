package runner

import "testing"

func exampleInt(a, b int) int {
    return a + b
}

func exampleString(a, b string) string {
    return a + b
}

func exampleSlice(a, b []int) []int {
    return append(a, b...)
}

func exampleSliceString(a, b []string) []string {
    return append(a, b...)
}

func TestRun(t *testing.T) {
    args := []int{1, 2}
    interfaceArgs := make([]interface{}, len(args))
    for i, v := range args {
        interfaceArgs[i] = v
    }
    Run(t, exampleInt, interfaceArgs, 3, IntVal)

}

func TestRun2(t *testing.T) {
    argsStr := []string{"a", "b"}
    interfaceArgsString := make([]interface{}, len(argsStr))
    for i, v := range argsStr {
        interfaceArgsString[i] = v
    }
    Run(t, exampleString, interfaceArgsString, "ab", StringVal)
}

func TestRun3(t *testing.T) {
    a := [][]int{{1, 2}, {3, 4}}
    interfaceArgs := make([]interface{}, 2)
    for i := 0; i < 2; i++ {
        interfaceArgs[i] = a[i]
    }
    Run(t, exampleSlice, interfaceArgs, []int{1, 2, 3, 4}, SliceInt)
}

func TestRun4(t *testing.T) {
    a := [][]string{{"1", "2"}, {"3", "4"}}
    interfaceArgs := make([]interface{}, 2)
    for i := 0; i < 2; i++ {
        interfaceArgs[i] = a[i]
    }
    Run(t, exampleSliceString, interfaceArgs, []string{"1", "2", "3", "4"}, SliceString)
}
