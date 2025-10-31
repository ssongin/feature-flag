package featureflag

import (
	"testing"
)

// helper: creates a sample feature tree for testing
func mockFeatures() *Features {
	return &Features{
		Clusters: []Cluster{
			{
				Label: "MainCluster",
				BooleanNodes: []BooleanNode{
					{Label: "Enabled", Value: true, Default: false},
				},
				PercentNodes: []PercentageNode{
					{Label: "Threshold", Value: 75, Default: 50},
				},
				StringNodes: []StringNode{
					{Label: "Username", Value: "admin", Default: "guest"},
				},
				ChoiceNodes: []ChoiceNode{
					{Label: "Mode", Value: "auto", Default: "manual"},
				},
				Clusters: []Cluster{
					{
						Label: "SubCluster",
						BooleanNodes: []BooleanNode{
							{Label: "BetaFeature", Value: false, Default: true},
						},
					},
				},
			},
		},
	}
}

func TestGetBoolValue_Success(t *testing.T) {
	f := mockFeatures()
	val := f.GetBoolValue("MainCluster.Enabled")
	if !val {
		t.Errorf("expected true, got %v", val)
	}
}

func TestGetChoiceValue_Success(t *testing.T) {
	f := mockFeatures()
	val := f.GetChoiceValue("MainCluster.Mode")
	if val != "auto" {
		t.Errorf("expected 'auto', got %q", val)
	}
}

func TestGetPercentageValue_Success(t *testing.T) {
	f := mockFeatures()
	val := f.GetPercentageValue("MainCluster.Threshold")
	if val != 75 {
		t.Errorf("expected 75, got %d", val)
	}
}

func TestGetStringValue_Success(t *testing.T) {
	f := mockFeatures()
	val := f.GetStringValue("MainCluster.Username")
	if val != "admin" {
		t.Errorf("expected 'admin', got %q", val)
	}
}

func TestGetBoolValue_FromNestedCluster(t *testing.T) {
	f := mockFeatures()
	val := f.GetBoolValue("MainCluster.SubCluster.BetaFeature")
	if val {
		t.Errorf("expected false, got %v", val)
	}
}

func TestGet_InvalidPath(t *testing.T) {
	f := mockFeatures()
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("expected panic for invalid path, got none")
		} else {
			if err, ok := r.(error); ok {
				if err.Error() != "Cluster not found: NonExistentCluster" {
					t.Errorf("unexpected panic message: %v", err)
				}
			} else if s, ok := r.(string); ok {
				if s != "Cluster not found: NonExistentCluster" {
					t.Errorf("unexpected panic message: %v", s)
				}
			}
		}
	}()
	_ = f.Get("NonExistentCluster.Node")
}

func TestGet_WrongNodeType(t *testing.T) {
	f := mockFeatures()
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("expected panic for type error, got none")
		} else {
			expected := `node at "MainCluster.Username" is not a BooleanNode`
			if err, ok := r.(error); ok {
				if err.Error() != expected {
					t.Errorf("unexpected panic message: %v", err)
				}
			} else if s, ok := r.(string); ok {
				if s != expected {
					t.Errorf("unexpected panic message: %v", s)
				}
			}
		}
	}()
	_ = f.GetBoolValue("MainCluster.Username")
}
