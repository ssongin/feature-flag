package featureflag

import (
	"errors"
	"fmt"
	"strings"
)

func (f *Features) GetChoiceValue(path string) (string, error) {
	node, err := f.Get(path)
	if err != nil {
		return "", err
	}

	cnode, ok := node.(*ChoiceNode)
	if !ok {
		return "", fmt.Errorf("node at %q is not a ChoiceNode", path)
	}

	return cnode.Value, nil
}

func (f *Features) GetPercentageValue(path string) (int, error) {
	node, err := f.Get(path)
	if err != nil {
		return 0, err
	}

	pnode, ok := node.(*PercentageNode)
	if !ok {
		return 0, fmt.Errorf("node at %q is not a PercentageNode", path)
	}

	return pnode.Value, nil
}

func (f *Features) GetStringValue(path string) (string, error) {
	node, err := f.Get(path)
	if err != nil {
		return "", err
	}

	snode, ok := node.(*StringNode)
	if !ok {
		return "", fmt.Errorf("node at %q is not a StringNode", path)
	}

	return snode.Value, nil
}

func (f *Features) GetBoolValue(path string) (bool, error) {
	node, err := f.Get(path)
	if err != nil {
		return false, err
	}

	bnode, ok := node.(*BooleanNode)
	if !ok {
		return false, fmt.Errorf("node at %q is not a BooleanNode", path)
	}

	return bnode.Value, nil
}

func (f *Features) Get(path string) (Node, error) {
	parts := strings.Split(path, ".")
	if len(parts) == 0 {
		return nil, errors.New("empty path")
	}

	for _, cluster := range f.Cluster {
		if cluster.Label == parts[0] {
			return cluster.getRecursive(parts[1:])
		}
	}

	return nil, errors.New("cluster not found: " + parts[0])
}

func (c *Cluster) getRecursive(parts []string) (Node, error) {
	if len(parts) == 0 {
		return nil, errors.New("no node label specified")
	}
	current := parts[0]

	// Traverse into sub-clusters
	for _, sub := range c.Clusters {
		if sub.Label == current {
			return sub.getRecursive(parts[1:])
		}
	}

	// If we're at the leaf, check all node types
	if len(parts) == 1 {
		label := parts[0]

		for i := range c.BooleanNodes {
			if c.BooleanNodes[i].Label == label {
				return &c.BooleanNodes[i], nil
			}
		}
		for i := range c.StringNodes {
			if c.StringNodes[i].Label == label {
				return &c.StringNodes[i], nil
			}
		}
		for i := range c.PercentNodes {
			if c.PercentNodes[i].Label == label {
				return &c.PercentNodes[i], nil
			}
		}
		for i := range c.ChoiceNodes {
			if c.ChoiceNodes[i].Label == label {
				return &c.ChoiceNodes[i], nil
			}
		}

		return nil, errors.New("node not found: " + label)
	}

	return nil, errors.New("invalid path: " + strings.Join(parts, "."))
}
