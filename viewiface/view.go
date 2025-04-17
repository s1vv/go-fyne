package viewiface

type ViewInterface interface {
	SetStatus(text string)
	ShowProgress()
	HideProgress()
}
