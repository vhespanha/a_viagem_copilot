package ui

const (
	ScreenWidth  = 1920
	ScreenHeight = 1080
)

const (
	// dialogue box elements
	DialogueBoxID ElementID = iota
	ChoiceOneID
	ChoiceTwoID
	ChoiceDontKnowID
	SkipDialogueID

	// death screen elements
	DeathScreenID

	// hud elements
	FullScreenButtonID
)

type CommandID uint8

const (
	UpdateDialogueBox CommandID = iota
	HideDialogue
	ShowDeathScreen
	HideDeathScreen
	ToggleFullScreen
)
