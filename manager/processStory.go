package manager

import (
	"database/sql"
	"strings"

	"github.com/hinupurthakur/collaborative-story/db"
	"github.com/hinupurthakur/collaborative-story/model"
	log "github.com/sirupsen/logrus"
)

const (
	STORY_SIZE     = 7  // paragraphs
	PARAGRAPH_SIZE = 10 // sentences
	SENTENCE_SIZE  = 15 // words
)

/*
ProcessWord processes each word and
create a story with title size as 2
and story containing STORY_SIZE,
paragraphs with PARAGRAPH_SIZE and
sentences of SENTENCE_SIZE words.
*/
func ProcessWord(word string) (model.NewWordResponse, error) {
	log.Infoln("processing word: ", word)
	var lastStory model.StoryDTO
	var result model.NewWordResponse
	lastStory, err := db.GetLastStory()
	if err != nil {
		if err == sql.ErrNoRows {
			result, err = db.InsertTitle(word)
			if err != nil {
				log.Errorln("add new word: unable to insert word", err)
				return model.NewWordResponse{}, err
			} else {
				return result, nil
			}
		}
		log.Errorln("add new word: unable to insert title", err)
		return model.NewWordResponse{}, err
	}
	if !strings.Contains(lastStory.Title, " ") {
		result, err = db.UpdateTitleByID(lastStory.Title+" "+word, lastStory.ID)
		if err != nil {
			log.Errorln("add new word: unable to insert title with two words", err)
			return model.NewWordResponse{}, err
		}
	} else {
		result, err = CheckParagraph(word, lastStory)
		if err != nil {
			log.Errorln("add new word: unable to insert sentence", err)
			return model.NewWordResponse{}, err
		}
		result.Title = lastStory.Title
		return result, nil
	}

	return result, nil
}

/*
ProcessSentence processes each word and creates the sentences
of provided SENTENCE_SIZE
*/
func ProcessSentence(word string, storyID int) (model.NewWordResponse, error) {
	var result model.NewWordResponse
	lastSentence, err := db.GetLastSentenceByStoryID(storyID)
	if err != nil {
		if err == sql.ErrNoRows {
			result, err = db.InsertSentence(word, storyID)
			result.CurrentSentence = word
			if err != nil {
				return model.NewWordResponse{}, err
			} else {
				return result, nil
			}
		}
		return model.NewWordResponse{}, err
	}
	totalWords := strings.Split(lastSentence.Sentence, " ")
	if len(totalWords) < SENTENCE_SIZE {
		result, err = db.UpdateSentenceByID(lastSentence.Sentence+" "+word, storyID, lastSentence.ID)
		if err != nil {
			return model.NewWordResponse{}, err
		} else {
			return result, nil
		}

	} else {
		result, err = db.InsertSentence(word, storyID)
		if err != nil {
			return model.NewWordResponse{}, err
		} else {
			return result, nil
		}
	}
}

/*
CheckParagraph checks the last paragraph size and
creates new or update the last one if
the paragraph have sentences less than PARAGRAPH_SIZE
and if paragraphs are completed for a story,
then creates a new story
*/
func CheckParagraph(word string, lastStory model.StoryDTO) (model.NewWordResponse, error) {
	var result model.NewWordResponse
	var sentences []string
	paragraphCount, err := db.GetNoOfParagraphByStory(lastStory.ID)
	if err != nil {
		return model.NewWordResponse{}, err
	}
	lastParagraph, err := db.GetLastParagraphByStoryID(lastStory.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			result, err = ProcessSentence(word, lastStory.ID)
			if err != nil {
				log.Errorln("add new word: unable to insert sentence", err)
				return model.NewWordResponse{}, err
			}
			result.Title = lastStory.Title
			sentences = append(sentences, result.CurrentSentence)
			err = db.InsertParagraph(sentences, lastStory.ID)
			if err != nil {
				log.Errorln(err)
			}
			return result, nil
		}
		log.Errorln("add new word: unable to insert title", err)
		return model.NewWordResponse{}, err
	}
	if completeStory(paragraphCount, lastParagraph) {
		result, err = db.InsertTitle(word)
		if err != nil {
			log.Errorln("add new word: unable to insert word", err)
			return model.NewWordResponse{}, err
		} else {
			return result, nil
		}
	} else {
		result, err = ProcessSentence(word, lastStory.ID)
		if err != nil {
			log.Errorln("add new word: unable to insert sentence", err)
			return model.NewWordResponse{}, err
		}
		result, err := ProcessParagraph(lastStory, lastParagraph, result)
		if err != nil {
			return model.NewWordResponse{}, nil
		}
		return result, nil
	}
}

/*
ProcessParagraph processes the last paragraph
and takes actions if it is complete or not
*/
func ProcessParagraph(lastStory model.StoryDTO, lastParagraph model.ParagraphDTO, result model.NewWordResponse) (model.NewWordResponse, error) {
	var sentences []string

	if completeParagraph(lastParagraph) {
		sentences = append(sentences, result.CurrentSentence)
		err := db.InsertParagraph(sentences, lastStory.ID)
		if err != nil {
			return model.NewWordResponse{}, err
		}
		result.Title = lastStory.Title
		return result, nil
	} else {
		if len(strings.Split(lastParagraph.Sentences[len(lastParagraph.Sentences)-1], " ")) < SENTENCE_SIZE {
			lastParagraph.Sentences[len(lastParagraph.Sentences)-1] = result.CurrentSentence
		} else {
			lastParagraph.Sentences = append(lastParagraph.Sentences, result.CurrentSentence)
		}
		db.UpdatePragraph(lastParagraph.Sentences, lastParagraph.ID)
		result.Title = lastStory.Title
		return result, nil
	}
}
