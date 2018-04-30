package symbol

import (
	. "github.com/kocircuit/kocircuit/lang/circuit/eval"
	. "github.com/kocircuit/kocircuit/lang/circuit/model"
)

func (vty *VarietySymbol) Evoke(span *Span, knot Fields) (Symbol, Effect, error) {
	if augmented, _, err := vty.Augment(span, knot); err != nil {
		return nil, nil, err
	} else if returns, _, err := augmented.Invoke(span); err != nil {
		return nil, nil, err
	} else {
		return returns.(Symbol), nil, nil
	}
}
