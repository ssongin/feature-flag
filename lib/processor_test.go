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
	val, err := f.GetBoolValue("MainCluster.Enabled")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !val {
		t.Errorf("expected true, got %v", val)
	}
}

func TestGetChoiceValue_Success(t *testing.T) {
	f := mockFeatures()
	val, err := f.GetChoiceValue("MainCluster.Mode")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if val != "auto" {
		t.Errorf("expected 'auto', got %q", val)
	}
}

func TestGetPercentageValue_Success(t *testing.T) {
	f := mockFeatures()
	val, err := f.GetPercentageValue("MainCluster.Threshold")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if val != 75 {
		t.Errorf("expected 75, got %d", val)
	}
}

func TestGetStringValue_Success(t *testing.T) {
	f := mockFeatures()
	val, err := f.GetStringValue("MainCluster.Username")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if val != "admin" {
		t.Errorf("expected 'admin', got %q", val)
	}
}

func TestGetBoolValue_FromNestedCluster(t *testing.T) {
	f := mockFeatures()
	val, err := f.GetBoolValue("MainCluster.SubCluster.BetaFeature")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if val {
		t.Errorf("expected false, got %v", val)
	}
}

func TestGet_InvalidPath(t *testing.T) {
	f := mockFeatures()
	_, err := f.Get("NonExistentCluster.Node")
	if err == nil {
		t.Fatalf("expected error for invalid path, got nil")
	}
	if err.Error() != "cluster not found: NonExistentCluster" {
		t.Errorf("unexpected error message: %v", err)
	}
}

func TestGet_WrongNodeType(t *testing.T) {
	f := mockFeatures()
	_, err := f.GetBoolValue("MainCluster.Username")
	if err == nil {
		t.Fatalf("expected type error, got nil")
	}
	expected := `node at "MainCluster.Username" is not a BooleanNode`
	if err.Error() != expected {
		t.Errorf("unexpected error message: %v", err)
	}
}
