package featureflag

import "errors"

type Node interface {
	GetLabel() string
	NodeType() string
	GetValue() interface{}
	SetValue(interface{}) error
	GetDefault() interface{}
	Reset()
}

func (n *BooleanNode) GetLabel() string        { return n.Label }
func (n *BooleanNode) NodeType() string        { return "boolean" }
func (n *BooleanNode) GetValue() interface{}   { return n.Value }
func (n *BooleanNode) GetDefault() interface{} { return n.Default }
func (n *BooleanNode) SetValue(v interface{}) error {
	b, ok := v.(bool)
	if !ok {
		return errors.New("invalid type for BooleanNode, expected bool")
	}
	n.Value = b
	return nil
}
func (n *BooleanNode) Reset() {
	n.Value = n.Default
}

func (n *StringNode) GetLabel() string        { return n.Label }
func (n *StringNode) NodeType() string        { return "string" }
func (n *StringNode) GetValue() interface{}   { return n.Value }
func (n *StringNode) GetDefault() interface{} { return n.Default }
func (n *StringNode) SetValue(v interface{}) error {
	s, ok := v.(string)
	if !ok {
		return errors.New("invalid type for StringNode, expected string")
	}
	n.Value = s
	return nil
}
func (n *StringNode) Reset() {
	n.Value = n.Default
}

func (n *PercentageNode) GetLabel() string        { return n.Label }
func (n *PercentageNode) NodeType() string        { return "percentage" }
func (n *PercentageNode) GetValue() interface{}   { return n.Value }
func (n *PercentageNode) GetDefault() interface{} { return n.Default }
func (n *PercentageNode) SetValue(v interface{}) error {
	i, ok := v.(int)
	if !ok {
		return errors.New("invalid type for PercentageNode, expected int")
	}
	n.Value = i
	return nil
}
func (n *PercentageNode) Reset() {
	n.Value = n.Default
}

func (n *ChoiceNode) GetLabel() string           { return n.Label }
func (n *ChoiceNode) NodeType() string           { return "choice" }
func (n *ChoiceNode) GetValue() interface{}      { return n.Value }
func (n *ChoiceNode) GetDefault() interface{}    { return n.Default }
func (n *ChoiceNode) GetChoiceOptions() []string { return n.GetChoiceOptions() }

func (n *ChoiceNode) SetValue(v interface{}) error {
	s, ok := v.(string)
	if !ok {
		return errors.New("invalid type for ChoiceNode, expected string")
	}

	for _, opt := range n.Options {
		if opt == s {
			n.Value = s
			return nil
		}
	}
	return errors.New("invalid choice value: " + s)
}

func (n *ChoiceNode) Reset() {
	n.Value = n.Default
}
