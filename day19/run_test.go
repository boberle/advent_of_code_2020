package day19

import "testing"

func TestTerminal_doesMatch(t *testing.T) {
	tests := []struct {
		name        string
		message     Message
		index       int
		wantIsMatch bool
		wantNext    int
	}{
		{"should match", "a", 0, true, 0},
		{"should match", "aa", 0, true, 1},
		{"should match", "aa", 1, true, 0},
		{"should match", "aaa", 1, true, 2},
		{"should not match", "b", 0, false, 0},
		{"should not match", "ab", 1, false, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			terminal := Terminal{'a'}
			gotIsMatch, gotNext := terminal.doesMatch(tt.message, tt.index)
			if gotIsMatch != tt.wantIsMatch {
				t.Errorf("doesMatch(), isMatch: got = %v, want %v", gotIsMatch, tt.wantIsMatch)
			}
			if gotNext != tt.wantNext {
				t.Errorf("doesMatch(): next: got = %v, want %v", gotNext, tt.wantNext)
			}
		})
	}
}

func TestRuleSequence_doesMatch(t *testing.T) {
	var r0, r1, r2, r3, r4 Rule
	r0 = Terminal{'a'}
	r1 = Terminal{'b'}
	r2 = RuleSequence{[]*Rule{&r0, &r1}}
	r3 = RuleSequence{[]*Rule{&r0, &r2}}
	r4 = RuleSequence{[]*Rule{&r2, &r3}}

	tests := []struct {
		name        string
		seq         *Rule
		message     Message
		index       int
		wantIsMatch bool
		wantNext    int
	}{
		{"test", &r2, "a", 0, false, 0},
		{"test", &r2, "ab", 0, true, 0},
		{"test", &r2, "aaba", 1, true, 3},
		{"test", &r2, "abba", 1, false, 0},
		{"test", &r3, "aab", 0, true, 0},
		{"test", &r3, "aabb", 0, true, 3},
		{"test", &r3, "babb", 0, false, 0},
		{"test", &r3, "baabb", 1, true, 4},
		{"test", &r4, "abaab", 0, true, 0},
		{"test", &r4, "babaab", 1, true, 0},
		{"test", &r4, "babaaba", 1, true, 6},
		{"test", &r4, "bbbaaaa", 1, false, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIsMatch, gotNext := (*tt.seq).doesMatch(tt.message, tt.index)
			if gotIsMatch != tt.wantIsMatch {
				t.Errorf("doesMatch(): isMatch: got = %v, want %v", gotIsMatch, tt.wantIsMatch)
			}
			if gotNext != tt.wantNext {
				t.Errorf("doesMatch(): next: got = %v, want %v", gotNext, tt.wantNext)
			}
		})
	}
}

func TestRuleChoices_doesMatch(t *testing.T) {
	var r0, r1, r2, r3 Rule
	r0 = Terminal{'a'}
	r1 = Terminal{'b'}
	r2 = RuleChoices{[]RuleSequence{
		{[]*Rule{&r0, &r1}},
		{[]*Rule{&r1, &r0}},
	}}
	r3 = RuleSequence{[]*Rule{&r0, &r2}}

	tests := []struct {
		name        string
		rule        *Rule
		message     Message
		index       int
		wantIsMatch bool
		wantNext    int
	}{
		{"test", &r2, "a", 0, false, 0},
		{"test", &r2, "b", 0, false, 0},
		{"test", &r2, "ab", 0, true, 0},
		{"test", &r2, "ba", 0, true, 0},
		{"test", &r2, "aaba", 1, true, 3},
		{"test", &r2, "abaa", 1, true, 3},
		{"test", &r3, "aab", 0, true, 0},
		{"test", &r3, "aba", 0, true, 0},
		{"test", &r3, "aabb", 0, true, 3},
		{"test", &r3, "abab", 0, true, 3},
		{"test", &r3, "babb", 0, false, 0},
		{"test", &r3, "baabb", 1, true, 4},
		{"test", &r3, "babab", 1, true, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIsMatch, gotNext := (*tt.rule).doesMatch(tt.message, tt.index)
			if gotIsMatch != tt.wantIsMatch {
				t.Errorf("doesMatch(): isMatch: got = %v, want %v", gotIsMatch, tt.wantIsMatch)
			}
			if gotNext != tt.wantNext {
				t.Errorf("doesMatch(): next: got = %v, want %v", gotNext, tt.wantNext)
			}
		})
	}
}

func Test_doesMatch(t *testing.T) {
	var r0, r1, r2, r3 Rule
	r0 = Terminal{'a'}
	r1 = Terminal{'b'}
	r2 = RuleChoices{[]RuleSequence{
		{[]*Rule{&r0, &r1}},
		{[]*Rule{&r1, &r0}},
	}}
	r3 = RuleSequence{[]*Rule{&r0, &r2}}

	tests := []struct {
		name        string
		rule        *Rule
		message     Message
		wantIsMatch bool
	}{
		{"test", &r0, "a", true},
		{"test", &r0, "aa", false},
		{"test", &r2, "a", false},
		{"test", &r2, "ab", true},
		{"test", &r2, "ba", true},
		{"test", &r2, "aaa", false},
		{"test", &r2, "aab", false},
		{"test", &r3, "aab", true},
		{"test", &r3, "aba", true},
		{"test", &r3, "aabb", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIsMatch := doesMatch(*tt.rule, tt.message)
			if gotIsMatch != tt.wantIsMatch {
				t.Errorf("doesMatch(): isMatch: got = %v, want %v", gotIsMatch, tt.wantIsMatch)
			}
		})
	}
}
