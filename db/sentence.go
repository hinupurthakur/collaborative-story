package db

import "github.com/hinupurthakur/collaborative-story/model"

func GetLastSentenceByStoryID(storyID int) (model.SentenceDTO, error) {
	var lastSentence model.SentenceDTO
	selectQuery := `SELECT
						id,
						sentence
					FROM
						sentences
					WHERE 
						story_id=$1
					ORDER BY id desc
					LIMIT 1;`
	err := db.Get(&lastSentence, selectQuery, storyID)
	if err != nil {
		return model.SentenceDTO{}, err
	}
	return lastSentence, nil
}

func UpdateSentenceByID(sentence string, storyID, id int) (model.NewWordResponse, error) {
	var result model.NewWordResponse
	insertQuery := `UPDATE	sentences
					SET
						sentence=$1,
						story_id=$2
					WHERE id=$3`
	_, err := db.Exec(insertQuery, sentence, storyID, id)
	if err != nil {
		return model.NewWordResponse{}, err
	}
	result.ID = storyID
	result.CurrentSentence = sentence
	return result, nil
}

func InsertSentence(sentence string, storyID int) (model.NewWordResponse, error) {
	var id int
	var result model.NewWordResponse
	insertQuery := `INSERT INTO
					sentences
					(
						sentence,
						story_id
					) 
					VALUES
					(
						$1,
						$2
					)
					RETURNING id`
	err := db.QueryRowx(insertQuery, sentence, storyID).Scan(&id)
	result.ID = storyID
	result.CurrentSentence = sentence
	return result, err
}
