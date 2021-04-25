package okpattern

// HumanResources 人事
type HumanResources interface {
	Selection()        // 選考
	FirstInterviewer() // 面談（1次面接官）
}
