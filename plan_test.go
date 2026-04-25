package acp

import "testing"

func TestParsePlanEntryStatus(t *testing.T) {
	t.Parallel()
	cases := []struct {
		input string
		want  PlanEntryStatus
	}{
		{"pending", PlanEntryStatusPending},
		{"PENDING", PlanEntryStatusPending},
		{"in_progress", PlanEntryStatusInProgress},
		{"IN_PROGRESS", PlanEntryStatusInProgress},
		{"completed", PlanEntryStatusCompleted},
		{"COMPLETED", PlanEntryStatusCompleted},
		{"", PlanEntryStatusPending},
		{"unknown", PlanEntryStatusPending},
	}
	for _, tc := range cases {
		got := ParsePlanEntryStatus(tc.input)
		if got != tc.want {
			t.Errorf("ParsePlanEntryStatus(%q) = %q; want %q", tc.input, got, tc.want)
		}
	}
}

func TestParsePlanEntryPriority(t *testing.T) {
	t.Parallel()
	cases := []struct {
		input string
		want  PlanEntryPriority
	}{
		{"high", PlanEntryPriorityHigh},
		{"HIGH", PlanEntryPriorityHigh},
		{"medium", PlanEntryPriorityMedium},
		{"MEDIUM", PlanEntryPriorityMedium},
		{"low", PlanEntryPriorityLow},
		{"LOW", PlanEntryPriorityLow},
		{"", PlanEntryPriorityMedium},
		{"unknown", PlanEntryPriorityMedium},
	}
	for _, tc := range cases {
		got := ParsePlanEntryPriority(tc.input)
		if got != tc.want {
			t.Errorf("ParsePlanEntryPriority(%q) = %q; want %q", tc.input, got, tc.want)
		}
	}
}

func TestNewPlanEntry(t *testing.T) {
	t.Parallel()
	e := NewPlanEntry("fix the bug", PlanEntryStatusInProgress, PlanEntryPriorityHigh)
	if e.Content != "fix the bug" {
		t.Errorf("Content = %q; want %q", e.Content, "fix the bug")
	}
	if e.Status != PlanEntryStatusInProgress {
		t.Errorf("Status = %q; want %q", e.Status, PlanEntryStatusInProgress)
	}
	if e.Priority != PlanEntryPriorityHigh {
		t.Errorf("Priority = %q; want %q", e.Priority, PlanEntryPriorityHigh)
	}
}
