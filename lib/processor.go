package featureflag

import (
	"errors"
	"fmt"
	"strings"

	"github.com/ssongin/core"
)

func (f *Features) GetChoiceValue(path string) string {
	node := f.Get(path)

	cnode, ok := node.(*ChoiceNode)
	if !ok {
		core.CheckError("GetChoiceValue: type assertion failed", fmt.Errorf("node at %q is not a ChoiceNode", path))
		return ""
	}

	return cnode.Value
}

func (f *Features) GetChoiceOptions(path string) []string {
	node := f.Get(path)

	cnode, ok := node.(*ChoiceNode)
	if !ok {
		core.CheckError("GetChoiceOptions: type assertion failed", fmt.Errorf("node at %q is not a ChoiceNode", path))
		return nil
	}

	return cnode.Options
}

func (f *Features) GetPercentageValue(path string) int {
	node := f.Get(path)

	pnode, ok := node.(*PercentageNode)
	if !ok {
		core.CheckError("GetPercentageValue: type assertion failed", fmt.Errorf("node at %q is not a PercentageNode", path))
		return 0
	}

	return pnode.Value
}

func (f *Features) GetStringValue(path string) string {
	node := f.Get(path)

	snode, ok := node.(*StringNode)
	if !ok {
		core.CheckError("GetStringValue: type assertion failed", fmt.Errorf("node at %q is not a StringNode", path))
		return ""
	}

	return snode.Value
}

func (f *Features) GetBoolValue(path string) bool {
	node := f.Get(path)

	bnode, ok := node.(*BooleanNode)
	if !ok {
		core.CheckError("GetBoolValue: type assertion failed", fmt.Errorf("node at %q is not a BooleanNode", path))
		return false
	}

	return bnode.Value
}

func (f *Features) Get(path string) Node {
	parts := strings.Split(path, ".")
	if len(parts) == 0 {
		core.CheckError("Get: empty path", errors.New("empty path"))
	}

	for _, cluster := range f.Clusters {
		if cluster.Label == parts[0] {
			return cluster.getRecursive(parts[1:])
		}
	}

	core.CheckError("Features: get", errors.New("Cluster not found: "+parts[0]))
	return nil
}

func (c *Cluster) getRecursive(parts []string) Node {
	if len(parts) == 0 {
		core.CheckError("getRecursive:", errors.New("no node label specified"))
		return nil
	}
	current := parts[0]

	for _, sub := range c.Clusters {
		if sub.Label == current {
			return sub.getRecursive(parts[1:])
		}
	}

	if len(parts) == 1 {
		label := parts[0]

		for i := range c.BooleanNodes {
			if c.BooleanNodes[i].Label == label {
				return &c.BooleanNodes[i]
			}
		}
		for i := range c.StringNodes {
			if c.StringNodes[i].Label == label {
				return &c.StringNodes[i]
			}
		}
		for i := range c.PercentNodes {
			if c.PercentNodes[i].Label == label {
				return &c.PercentNodes[i]
			}
		}
		for i := range c.ChoiceNodes {
			if c.ChoiceNodes[i].Label == label {
				return &c.ChoiceNodes[i]
			}
		}

		core.CheckError("getRecursive:", errors.New("node not found: "+label))
		return nil
	}
	core.CheckError("getRecursive:", errors.New("invalid path: "+strings.Join(parts, ".")))
	return nil
}
