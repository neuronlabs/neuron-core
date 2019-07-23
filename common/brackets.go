package common

import (
	"github.com/neuronlabs/neuron-core/errors"
	"github.com/neuronlabs/neuron-core/errors/class"
)

// SplitBracketParameter splits the parameters within the '[' and ']' brackets.
func SplitBracketParameter(bracketed string) (values []string, err error) {
	doubleOpen := func() error {
		err := errors.New(class.CommonParseBrackets, "double open square brackets")
		err = err.SetDetailf("open square bracket '[' found, without closing ']' in: '%s'", bracketed)
		return err
	}

	// set initial indexes
	startIndex := -1
	endIndex := -1

	for i := 0; i < len(bracketed); i++ {
		c := bracketed[i]
		switch c {
		case AnnotationOpenedBracket:
			if startIndex > endIndex {
				err = doubleOpen()
				return nil, err
			}
			startIndex = i
		case AnnotationClosedBracket:
			// if opening bracket not set or in case of more than one brackets
			// if start was not set before this endIndex
			if startIndex == -1 || startIndex < endIndex {
				err := errors.New(class.CommonParseBrackets, "no opening bracket found")
				err = err.SetDetailf("close square bracket ']' found, without opening '[' in '%s'", bracketed)
				return nil, err
			}
			endIndex = i
			values = append(values, bracketed[startIndex+1:endIndex])
		}
	}
	if (startIndex != -1 && endIndex == -1) || startIndex > endIndex {
		err = doubleOpen()
		return nil, err
	}
	return values, nil
}
