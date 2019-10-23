package starter

import (
	"net/http"
	"strings"

	"github.com/IMQS/nf"
	"github.com/IMQS/serviceauth"
	"github.com/julienschmidt/httprouter"
)

func (s *Service) ping(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	nf.SendPong(w)
}

func (s *Service) listFrogs(w http.ResponseWriter, r *http.Request, p httprouter.Params, auth *serviceauth.Token) {
	query := ""
	args := []interface{}{}

	qtype := r.FormValue("type")
	if qtype != "" {
		for _, t := range strings.Split(qtype, ",") {
			query += "OR frog_type_id = ?"
			args = append(args, t)
		}
	}

	if strings.HasPrefix(query, "OR") {
		query = query[3:]
	}

	frogs := []frog{}
	s.db.Where(query, args...).Find(&frogs)
	nf.SendJSON(w, frogs)
}

func (s *Service) addFrogs(w http.ResponseWriter, r *http.Request, p httprouter.Params, auth *serviceauth.Token) {
	frogs := []jsonInputFrog{}
	nf.ReadJSON(r, &frogs)
	for _, f := range frogs {
		mf := frog{
			FrogTypeID:  &f.FrogTypeID,
			Description: &f.Description,
		}
		s.db.Create(&mf)
		nf.Check(s.db.Error)
	}
	nf.SendOK(w)
}
