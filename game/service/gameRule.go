package service

type GameRuleCheck interface {
	FinishCheck() bool
	ReturnResult() *Result
}

type Game struct {
	Winner   *Result
	IsFinish bool
}
