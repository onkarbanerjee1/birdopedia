package birds

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/gorilla/schema"
	cfg "github.com/onkarbanerjee1/birdopedia/config"
)

// GetAllBirds gets a list of all birds
func GetAllBirds(env *cfg.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		birds, err := All(env.DB)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		for _, each := range birds {
			fmt.Fprintf(w, "<br>%s<br>", each)
		}

		fmt.Fprintln(w, "<br><br><a href=\"/\">Home Page</a>")
		return
	})

}

// GetBirdsByName gets a list of all birds matching the common name
func GetBirdsByName(env *cfg.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		vars := mux.Vars(r)

		if len(vars) != 1 || vars["name"] == "" {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("No name passed"))
			return
		}

		bird, err := ByName(env.DB, vars["name"])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		if bird == nil {
			fmt.Fprint(w, "No such bird<br><br><a href=\"/\">Home Page</a>")
			return
		}

		fmt.Fprintf(w, "%s", bird)
		fmt.Fprintln(w, "<br><br><a href=\"/\">Home Page</a>")
		return
	})
}

// NewBirdForm handles bird adding form
func NewBirdForm() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("tpl/new.gtpl")
		if err != nil {
			log.Fatal("Got error in template execute")
		}

		err = t.Execute(w, nil)
	})
}

// InsertNewBird will insert a new bird
func InsertNewBird(env *cfg.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		r.ParseForm()

		decoder := schema.NewDecoder()
		bird := Bird{}
		// r.PostForm is a map of our POST form values
		err := decoder.Decode(&bird, r.PostForm)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		_, err = Insert(env.DB, &bird)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "bird added successfully!!")
		fmt.Fprintln(w, "<br><br><a href=\"/\">Home Page</a>")
		return
	})
}

// UpdateBirdForm handles bird updating form
func UpdateBirdForm() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("tpl/update.gtpl")
		if err != nil {
			log.Fatal("Got error in template execute")
		}

		t.Execute(w, nil)
	})
}

// UpdateBird updates a bird by name
func UpdateBird(env *cfg.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		r.ParseForm()

		decoder := schema.NewDecoder()
		bird := Bird{}
		// r.PostForm is a map of our POST form values
		err := decoder.Decode(&bird, r.PostForm)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		id, err := getID(env.DB, bird.CommonName)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		err = Update(env.DB, id, &bird)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "bird updated successfully!!")
		fmt.Fprintln(w, "<br><br><a href=\"/\">Home Page</a>")
		return
	})
}

// DeleteBirdForm handles bird deleting form
func DeleteBirdForm() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("tpl/delete.gtpl")
		if err != nil {
			log.Fatal("Got error in template execute")
		}

		t.Execute(w, nil)
	})
}

// DeleteBird deletes a bird by name
func DeleteBird(env *cfg.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		name := r.FormValue("CommonName")
		if name == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("No bird with %s", name)))
			return
		}

		id, err := getID(env.DB, name)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		err = Delete(env.DB, id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "bird deleted successfully!!")
		fmt.Fprintln(w, "<br><br><a href=\"/\">Home Page</a>")
		return
	})
}

// MainPage returns a handler for the main page
func MainPage() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("tpl/main.gtpl")
		if err != nil {
			log.Fatal("Got error in template execute")
		}

		t.Execute(w, nil)
	})
}

// QueryPage returns a handler for the main page
func QueryPage() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("tpl/query.gtpl")
		if err != nil {
			log.Fatal("Got error in template execute")
		}

		t.Execute(w, nil)
	})
}

// GetBirdByName gets a bird matching the common name
func GetBirdByName(env *cfg.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		vars := r.URL.Query()

		if len(vars) != 1 || len(vars["name"]) != 1 {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("No name passed"))
			return
		}

		bird, err := ByName(env.DB, vars["name"][0])
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.WriteHeader(http.StatusOK)
		if bird == nil {
			fmt.Fprint(w, "No such bird")
			fmt.Fprintln(w, "<br><br><a href=\"/\">Home Page</a>")
			return
		}

		fmt.Fprintf(w, "%s", bird)
		fmt.Fprintln(w, "<br><br><a href=\"/\">Home Page</a>")
		return
	})
}
