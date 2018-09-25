package milestone

import "fmt"

type Milestone struct {
	CurrentPoint int
	IsWhiteValid bool
}

var blackNum = []int{5, 8, 12, 17, 23, 30}

func NewMilestone() Milestone {
	return Milestone{
		CurrentPoint: 0,
		IsWhiteValid: false,
	}
}

func (m Milestone) GetCurrentPoint() int {
	return m.CurrentPoint
}

func (m *Milestone) ResetCurrentPoint() {
	m.CurrentPoint = 0
}

func (m *Milestone) SetCurrentPoint(reportedNumber int) error {
	if m.IsWhiteValid || contains(blackNum, reportedNumber) {
		m.CurrentPoint = reportedNumber
		return nil
	}
	return fmt.Errorf("Not valid number")
}

func (m *Milestone) SetWhiteValid() {
	m.IsWhiteValid = true
}

func (m *Milestone) RemoveWhiteValid() {
	m.IsWhiteValid = false
}

func contains(s []int, e int) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}
