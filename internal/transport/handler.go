package transport

import (
	"LearnGoPersonGinPsql/internal/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

// Person предъявялет требования к сервису от handler
type Person interface {
	CreatePerson(person models.Person) error //create
	GetAllPersons() ([]models.Person, error) //read
	UpdatePerson(person models.Person) error //update
	DeletePerson(name string) error          //delete
}

type Handler struct {
	ServicePerson Person
}

// NewHandler создает Handler, кладем в него сервис
func NewHandler(person Person) *Handler {
	return &Handler{
		ServicePerson: person,
	}
}

func (h *Handler) InitRouter() *mux.Router {
	r := mux.NewRouter()

	person := r.PathPrefix("/person").Subrouter()
	{
		person.HandleFunc("", h.createPerson).Methods(http.MethodPost)
		person.HandleFunc("", h.getAllPersons).Methods(http.MethodGet)
		person.HandleFunc("", h.DeletePerson).Methods(http.MethodDelete)
		person.HandleFunc("", h.UpdatePerson).Methods(http.MethodPut)
	}
	return r
}

func (h *Handler) createPerson(w http.ResponseWriter, r *http.Request /*запрос*/) {
	//разгружаю (анмаршалю) JSON в структуру
	//вариант разгрузки джэсона в структуру: func (dec *Decoder) Decode(v any) error
	//здесь иду в бизнес логику (service) (если это все для сервиса, то бизнес-логика туть)

	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var person models.Person
	if err = json.Unmarshal(reqBytes, &person); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.ServicePerson.CreatePerson(person)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			return
		}
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) getAllPersons(w http.ResponseWriter, r *http.Request /*запрос*/) {
	persons, err := h.ServicePerson.GetAllPersons()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	response, err := json.Marshal(persons)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	_, err = w.Write(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (h *Handler) DeletePerson(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	name := string(reqBytes)
	err = h.ServicePerson.DeletePerson(name)
	if err != nil {
		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			return
		}
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) UpdatePerson(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var person models.Person
	if err = json.Unmarshal(reqBytes, &person); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.ServicePerson.UpdatePerson(person)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			return
		}
		return
	}
	w.WriteHeader(http.StatusCreated)
}
