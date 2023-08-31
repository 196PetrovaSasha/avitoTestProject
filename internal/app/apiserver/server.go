package apiserver

import (
	"avitoTest/internal/app/model"
	"avitoTest/internal/app/store"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
)

var (
	incorrectSegmentName = errors.New("incorrect segment name")
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
	store  store.Store
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/segments", s.handleSegmentsCreate()).Methods("POST")
}

func newServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logrus.New(),
		store:  store,
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) handleSegmentsCreate() http.HandlerFunc {
	type request struct {
		Name string `json:"segment_name"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		u := &model.Segment{
			Name: req.Name,
		}

		if err := s.store.Segment().Create(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		s.respond(w, r, http.StatusCreated, u)
	}
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
