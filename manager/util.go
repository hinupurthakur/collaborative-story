package manager

import (
	"strings"

	"github.com/hinupurthakur/collaborative-story/model"
)

/*
completeStory function checks if the story has STORY_SIZE paragraphs and
returns true otherwise false
*/
func completeStory(paragraphCount int, lastParagraph model.ParagraphDTO) bool {
	if paragraphCount == STORY_SIZE {
		if len(lastParagraph.Sentences) == PARAGRAPH_SIZE {
			if len(strings.Split(lastParagraph.Sentences[len(lastParagraph.Sentences)-1], " ")) == SENTENCE_SIZE {
				return true
			}
		}
	}
	return false
}

/*
completeParagraph function checks if
the story has PARAGRAPH_SIZE sentences and
returns true otherwise false
*/
func completeParagraph(lastParagraph model.ParagraphDTO) bool {
	if len(lastParagraph.Sentences) == PARAGRAPH_SIZE {
		if len(strings.Split(lastParagraph.Sentences[len(lastParagraph.Sentences)-1], " ")) == SENTENCE_SIZE {
			return true
		}
	}
	return false
}
