package ngpattern

// Company 会社
type Company interface {
	Selection()        // 選考
	FirstInterviewer() // 面談（1次面接官）
	LastInterviewer()  // 面談（最終面接官）
	Programming()      // 開発
	ManageEmployee()   // 社員管理
	CalcSalary()       // 給料計算
}

/*
会社という大きな概念で仕事を全部持ってる。
何かのフローが変る度に、会社全体に影響しちゃう。
*/
