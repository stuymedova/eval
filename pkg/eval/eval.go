package eval

import (
    "errors"
    "fmt"
    "strconv"
    "unicode"
)

func Calc(expression string) (float64, error) {
    postfix, err := toRPN(expression)
    if err != nil {
        return 0, err
    }
    return evalRPN(postfix)
}

func toRPN(expression string) ([]string, error) {
    var stack []rune
    var postfix []string
    priority := map[rune]int{
        '+': 1,
        '-': 1,
        '*': 2,
        '/': 2,
        '(': 0,
    }
    var number string
    for _, ch := range expression {
        if unicode.IsDigit(ch) || ch == '.' {
            number += string(ch)
        } else {
            if number != "" {
                postfix = append(postfix, number)
                number = ""
            }
            switch ch {
            case ' ':
                continue
            case '(':
                stack = append(stack, ch)
            case ')':
                for len(stack) > 0 && stack[len(stack)-1] != '(' {
                    postfix = append(postfix, string(stack[len(stack)-1]))
                    stack = stack[:len(stack)-1]
                }
                if len(stack) == 0 {
                    return nil, errors.New("mismatched parentheses")
                }
                stack = stack[:len(stack)-1]
            case '+', '-', '*', '/':
                for len(stack) > 0 && priority[stack[len(stack)-1]] >= priority[ch] {
                    postfix = append(postfix, string(stack[len(stack)-1]))
                    stack = stack[:len(stack)-1]
                }
                stack = append(stack, ch)
            default:
                return nil, fmt.Errorf("invalid character: %c", ch)
            }
        }
    }
    if number != "" {
        postfix = append(postfix, number)
    }
    for len(stack) > 0 {
        if stack[len(stack)-1] == '(' {
            return nil, errors.New("mismatched parentheses")
        }
        postfix = append(postfix, string(stack[len(stack)-1]))
        stack = stack[:len(stack)-1]
    }
    return postfix, nil
}

func evalRPN(postfix []string) (float64, error) {
    var stack []float64
    for _, token := range postfix {
        if val, err := strconv.ParseFloat(token, 64); err == nil {
            stack = append(stack, val)
        } else if len(stack) >= 2 {
            b := stack[len(stack)-1]
            a := stack[len(stack)-2]
            stack = stack[:len(stack)-2]
            switch token {
            case "+":
                stack = append(stack, a+b)
            case "-":
                stack = append(stack, a-b)
            case "*":
                stack = append(stack, a*b)
            case "/":
                if b == 0 {
                    return 0, errors.New("division by zero")
                }
                stack = append(stack, a/b)
            default:
                return 0, fmt.Errorf("invalid operator: %s", token)
            }
        } else {
            return 0, errors.New("insufficient values in expression")
        }
    }
    if len(stack) != 1 {
        return 0, errors.New("error evaluating expression")
    }
    return stack[0], nil
}
