package dayninteen

import (
	"fmt"
	"strconv"
	"strings"
)

type part map[byte]int

type rule struct {
	Predicate *predicate
	Next      string
}

type predicate struct {
	CompareWith    byte
	Comparator     byte
	ValueToCompare int
}

func SumOfRatingsOfAcceptedParts(lines []string) int {
	workflows := map[string][]rule{}
	parts := []part{}

	parseParts := false
	for _, line := range lines {
		if line == "" {
			parseParts = true
			continue
		}

		if parseParts {
			parts = append(parts, parsePart(line))
			continue
		}

		name, rules := parseWorkflow(line)
		workflows[name] = rules
	}

	acceptedParts := runWorkFlow(parts, workflows)

	sum := 0
	for _, p := range acceptedParts {
		for _, rating := range p {
			sum += rating
		}
	}
	return sum
}

func parsePart(line string) part {
	// Remove '{' and '}'
	line = line[1 : len(line)-1]
	split := strings.Split(line, ",")

	p := part{}
	for _, s := range split {
		ratingSplit := strings.Split(s, "=")
		component := ratingSplit[0][0]
		value, err := strconv.Atoi(ratingSplit[1])
		if err != nil {
			panic(err)
		}

		p[component] = value
	}

	return p
}

func parseWorkflow(line string) (string, []rule) {
	split := strings.Split(line, "{")
	name := split[0]
	rawRules := split[1]
	rawRules = strings.TrimRight(rawRules, "}")

	rawRuleSplit := strings.Split(rawRules, ",")

	rules := []rule{}
	for _, rawRule := range rawRuleSplit {
		s := strings.Split(rawRule, ":")
		if len(s) == 1 {
			rules = append(rules, rule{Next: s[0]})
			continue
		}

		predicateRaw := s[0]
		next := s[1]

		compareWith := predicateRaw[0]
		comparator := predicateRaw[1]
		val, err := strconv.Atoi(predicateRaw[2:])
		if err != nil {
			panic(err)
		}

		p := &predicate{CompareWith: compareWith, Comparator: comparator, ValueToCompare: val}

		rules = append(rules, rule{Predicate: p, Next: next})
	}

	return name, rules
}

func runWorkFlow(parts []part, workflows map[string][]rule) []part {
	start := "in"
	accepted := "A"
	rejected := "R"

	acceptedParts := []part{}

	for _, part := range parts {
		// Always start in "in"
		workflow := start
		// Run until not accepted or rejected
		for workflow != accepted && workflow != rejected {
			rules := workflows[workflow]

			for _, rul := range rules {
				if rul.Predicate == nil {
					workflow = rul.Next
					break
				}

				pred := *rul.Predicate
				if pred.Comparator == byte('<') {
					if part[pred.CompareWith] < pred.ValueToCompare {
						workflow = rul.Next
						break
					}
				} else {
					if part[pred.CompareWith] > pred.ValueToCompare {
						workflow = rul.Next
						break
					}
				}
			}
		}

		if workflow == accepted {
			acceptedParts = append(acceptedParts, part)
		}
	}

	return acceptedParts
}

func printWorkflows(workflows map[string][]rule) {
	for k, v := range workflows {
		fmt.Printf("%v: [ ", k)
		for _, r := range v {
			fmt.Printf("%v ", ruleToString(r))
		}
		fmt.Printf("]\n")
	}
}

func ruleToString(r rule) string {
	predicateStr := "nil"
	if r.Predicate != nil {
		p := *r.Predicate
		predicateStr = fmt.Sprintf("%v%v%v", string(p.CompareWith), string(p.Comparator), p.ValueToCompare)
	}
	return predicateStr + ":" + r.Next
}

func printParts(l []part) {
	for _, m := range l {
		fmt.Println(partToString(m))
	}
}

func partToString(p part) string {
	pStr := ""
	for k, v := range p {
		pStr += string(k) + ":" + strconv.Itoa(v) + " "
	}

	return pStr
}
