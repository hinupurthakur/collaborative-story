package db

import (
	"github.com/hinupurthakur/collaborative-story/model"
	"github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

func UpdatePragraph(sentences []string, paragraphID int) error {
	sqlQuery := `
		UPDATE paragraphs
		SET
			sentences=$1
		WHERE
			id=$2`
	_, err := db.Exec(sqlQuery, pq.StringArray(sentences), paragraphID)
	if err != nil {
		log.Errorln(err)
		return err
	}
	return nil
}

func GetLastParagraphByStoryID(storyID int) (model.ParagraphDTO, error) {
	var lastParagraph model.ParagraphDTO
	sqlQuery := `
			SELECT
				id,
				sentences
			FROM
				paragraphs
			WHERE story_id=$1
			ORDER by id desc
			LIMIT 1;`
	err := db.QueryRowx(sqlQuery, storyID).Scan(&lastParagraph.ID, (*pq.StringArray)(&lastParagraph.Sentences))
	if err != nil {
		return model.ParagraphDTO{}, err
	}
	return lastParagraph, nil
}

func GetNoOfParagraphByStory(storyID int) (int, error) {
	var paragraphCount int
	checkParaquery := `SELECT
							count(id)
						FROM
							paragraphs
						WHERE
							story_id = $1;`
	err := db.Get(&paragraphCount, checkParaquery, storyID)
	if err != nil {
		return 0, err
	}
	return paragraphCount, nil
}

func InsertParagraph(sentences []string, storyID int) error {
	var id int
	sqlQuery := `INSERT INTO
			paragraphs
			(
				sentences,
				story_id
			)
			VALUES
			(
				$1,
				$2
			)
			RETURNING id`
	err := db.QueryRowx(sqlQuery, pq.StringArray(sentences), storyID).Scan(&id)
	if err != nil {
		return err
	}
	return nil
}
