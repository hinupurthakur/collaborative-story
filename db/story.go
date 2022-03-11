package db

import (
	"encoding/json"
	"fmt"

	"github.com/hinupurthakur/collaborative-story/model"
	log "github.com/sirupsen/logrus"
)

func GetStoryByID(id int) (model.StoryDTO, error) {
	var story model.StoryDTO
	sqlQuery := `
		SELECT
			s.id,
			title,
			s.created_at,
			s.updated_at,
			json_agg(sentences)::jsonb as paragraphs
		FROM
			stories s,paragraphs p
		WHERE
			s.id=$1 AND s.id=p.story_id
		GROUP BY s.id;`
	rows, err := db.Queryx(sqlQuery, id)
	if err != nil {
		return model.StoryDTO{}, err
	}
	defer rows.Close()
	var sentences [][]string
	for rows.Next() {
		var jsonb string

		err = rows.Scan(&story.ID, &story.Title, &story.CreatedAt, &story.UpdatedAt, &jsonb)
		if err != nil {
			log.Errorln("get competitors monthly report: unable to scan the row", err)
			return model.StoryDTO{}, err
		}
		err := json.Unmarshal([]byte(jsonb), &sentences)
		if err != nil {
			log.Infoln("get story: unable to read paragraphs", err)
			return model.StoryDTO{}, err
		}
	}
	for _, value := range sentences {
		var sent model.Sentence
		sent.Sentences = value
		story.Paragraphs = append(story.Paragraphs, sent)
	}
	return story, nil
}

func InsertTitle(title string) (model.NewWordResponse, error) {
	var id int
	var result model.NewWordResponse
	insertQuery := `INSERT INTO
						stories
						(
							title
						) 
						VALUES
						(
							$1
						)
						RETURNING id`
	err := db.QueryRowx(insertQuery, title).Scan(&id)
	if err != nil {
		return model.NewWordResponse{}, err
	}
	result.ID = id
	result.Title = title
	result.CurrentSentence = ""
	return result, nil
}

func UpdateTitleByID(title string, id int) (model.NewWordResponse, error) {
	var result model.NewWordResponse
	insertQuery := `UPDATE
						stories
					SET
						title=$1
					WHERE id=$2`
	_, err := db.Exec(insertQuery, title, id)
	if err != nil {
		return model.NewWordResponse{}, err
	}
	result.ID = id
	result.Title = title
	result.CurrentSentence = ""
	return result, nil
}

func GetAllStoriesObject(offset, limit uint32, orderByClause string) ([]model.StoryDTO, error) {
	output := []model.StoryDTO{}
	sqlQuery := fmt.Sprintf(`SELECT 
					id,
					title,
					created_at,
					updated_at
				FROM
					stories
				%s
				OFFSET $1
				LIMIT $2;`, orderByClause)
	err := db.Select(&output, sqlQuery, offset, limit)
	return output, err
}

func GetLastStory() (model.StoryDTO, error) {
	var lastStory model.StoryDTO
	selectQuery := `SELECT
						id,
						title
					FROM
						stories
					ORDER BY id desc
					LIMIT 1;`
	err := db.Get(&lastStory, selectQuery)
	if err != nil {
		return model.StoryDTO{}, err
	}
	return lastStory, nil
}
