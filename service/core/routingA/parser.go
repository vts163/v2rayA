package routingA

func Parse(program string) (rules []interface{}, err error) {
	S, err := generateSyntaxTree(program)
	if err != nil {
		return
	}
	As := parseS(S)
	for _, A := range As {
		switch {
		case symMatch(A.children, []rune("B")):
			rules = append(rules, *newDefine(A.children[0]))
		case symMatch(A.children, []rune("C")):
			rules = append(rules, *newRouting(A.children[0]))
		}
	}
	return
}
