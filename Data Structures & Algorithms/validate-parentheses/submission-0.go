func isValid(s string) bool {
	pairs := make(map[rune]rune)
	pairs[')'] = '('
	pairs[']'] = '['
	pairs['}'] = '{'
	var flag bool
	stack := make([]rune, 0)
	for _, i := range s {
		if i == '(' || i == '[' || i == '{' {
			stack = append(stack, i) //если пришли открытые, добавляем в стек
		} else { //если пришла закрытая, то {() {(}
			if len(stack) == 0 {
				return false
			}
			
			if pairs[i] == stack[len(stack)-1] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	if len(stack) == 0 {
		flag = true
	}
	return flag
}