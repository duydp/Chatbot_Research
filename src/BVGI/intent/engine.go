package intent

import "fptai-sdk-go"

// Engine represents an intent recognition engine, such as Wit.AI, FPT.AI
type Engine interface {
	Detect(text string) string // Detect recognites intent of input text
	AddUtterance(intent, utterance string) error
	AddIntentUtterances(ius []IntentUtterance) error
	DeleteAllIntents() error
	Train() error
	ExtractEntity(text string) (fptai.Meaning, error)
}