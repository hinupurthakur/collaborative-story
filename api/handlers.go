package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"strconv"
	"strings"
	"sync"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/hinupurthakur/collaborative-story/db"
	"github.com/hinupurthakur/collaborative-story/manager"
	"github.com/hinupurthakur/collaborative-story/model"
	log "github.com/sirupsen/logrus"
)

/*
HealthCheck is a GET method
with the endpoint as /health
to check the health of the server
*/
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Infoln("GET: healthcheck API")
	result := `{"status": "Connection successful"}`
	var output map[string]interface{}
	err := json.Unmarshal([]byte(result), &output)
	if err != nil {
		log.Errorln("healthCheck: unable to encode string to JSON", err)
		http.Error(w, "healthCheck: unable to encode string to JSON", http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(output); err != nil {
		log.Errorln("healthCheck: unable to encode output", err)
		http.Error(w, "healthCheck: unable to encode output", http.StatusBadRequest)
		return
	}
}

/*
AddNewWord is a POST method
with the endpoint /add
to add a new word
*/
func AddNewWord(w http.ResponseWriter, r *http.Request) {
	log.Infoln("add word api reached")
	cr := make(chan *http.Request, 1)
	cr <- r
	var pleasewait sync.WaitGroup
	pleasewait.Add(1)

	go func() {
		defer pleasewait.Done()
		processAddWordAPI(w, cr)
	}()

	pleasewait.Wait()
	log.Infoln("add word api completed")
}

/*
processAddWordAPI handles the request of AddNewWord handler
*/
func processAddWordAPI(w http.ResponseWriter, cr chan *http.Request) {
	r := <-cr
	word := model.WordDTO{}
	if err := json.NewDecoder(r.Body).Decode(&word); err != nil {
		log.Errorln("add new word: inappropriate input", err)
		http.Error(w, `{"error": "inappropriate input"}`, http.StatusBadRequest)
		return
	}
	v := validator.New()
	err := v.Struct(word)
	if err != nil {
		log.Errorln("add new word: inappropriate data for word json", err)
		http.Error(w, `{"error": "multiple words sent"}`, http.StatusBadRequest)
		return
	}

	result, err := manager.ProcessWord(word.Word)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": "%v"}`, err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		log.Errorln("add new word: unable to encode result", err)
		http.Error(w, `{"error": "unable to encode result"}`, http.StatusInternalServerError)
		return
	}
}

/*
GetAllStories is a GET method
with the endpoint /stories
and returns list of stories
*/
func GetAllStories(w http.ResponseWriter, r *http.Request) {
	log.Infoln("get all stories api reached")
	var result model.Stories
	params := r.URL.Query()
	offset := params.Get("offset")
	if offset != "" {
		offset, err := strconv.ParseUint(offset, 10, 32)
		if err != nil {
			log.Errorln("get all stories: unable to fread offset", err)
			http.Error(w, `{"error": "inappropriate offset"}`, http.StatusInternalServerError)
			return
		}
		result.Offset = uint32(offset)
	}

	limit := params.Get("limit")
	if limit != "" {
		limit, err := strconv.ParseUint(limit, 10, 32)
		if err != nil {
			log.Errorln("get all stories: unable to fread offset", err)
			http.Error(w, `{"error": "inappropriate offset"}`, http.StatusInternalServerError)
			return
		}
		result.Limit = uint32(limit)
	}

	sort := params.Get("sort")
	order := params.Get("order")

	orderByClause := genrateOrderByClause(strings.Split(sort, ","), strings.Split(order, ","))
	stories, err := db.GetAllStoriesObject(result.Offset, result.Limit, orderByClause)
	if err != nil {
		log.Errorln("get all stories: unable to fetch from database", err)
		http.Error(w, `{"error": "unable to fetch stories"}`, http.StatusInternalServerError)
		return
	}

	result.Count = uint32(len(stories))
	result.Results = stories
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		log.Errorln("get all stories: unable to encode result", err)
		http.Error(w, `{"error": "unable to encode stories"}`, http.StatusInternalServerError)
		return
	}
	log.Infoln("get all stories api completed")
}

/*
GetStory function is a GET Method
with url as /stories/:id and
will return details of the story
*/
func GetStory(w http.ResponseWriter, r *http.Request) {
	log.Infoln("get story api reached")
	pathParam := mux.Vars(r)
	id, err := strconv.Atoi(pathParam["id"])
	if err != nil {
		log.Errorln("get story: wrong id", err)
		http.Error(w, `{"error": "provide correct story id"}`, http.StatusNotFound)
		return
	}
	result, err := db.GetStoryByID(id)
	if err != nil {
		log.Errorln("get story: unable to fetch from database", err)
		http.Error(w, `{"error": "unable to fetch story"}`, http.StatusInternalServerError)
		return
	}
	if result.ID == 0 {
		log.Errorln("get story: unable to fetch from database for the provided id")
		http.Error(w, `{"error": "unable to fetch story for the provided ID"}`, http.StatusNotFound)
		return
	}
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		log.Errorln("get story: unable to encode result", err)
		http.Error(w, `{"error": "unable to encode story"}`, http.StatusInternalServerError)
		return
	}
	log.Infoln("get story api completed")
}
