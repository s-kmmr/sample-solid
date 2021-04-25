package okpattern

// President 社長
type President interface {
	ManageEmployee()  // 社員管理
	LastInterviewer() // 面談（最終面接官）
}
