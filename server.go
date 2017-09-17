package main

import (
	"net/http"
	// "gopkg.in/mgo.v2"
)

// var DB *mgo.Database
// var Customer *mgo.Collection

// type lead struct {
// 	Name      string `json:"name" bson:"name"`
// 	Email     string `json:"email" bson:"email"`
// 	Ip        string `json:"ip" bson:"ip"`
// 	Timestamp string `json:"timestamp" bson:"timestamp"`
// }

func init() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	// http.HandleFunc("/form", func(w http.ResponseWriter, r *http.Request) {
	// 	ctx := appengine.NewContext(r)
	// 	if r.Method == http.MethodPost {
	// 		err := r.ParseForm()
	// 		if err != nil {
	// 			log.Errorf(ctx, "error", err)
	// 		}
	// 		t := time.Now()
	// 		utc, err := time.LoadLocation("America/Sao_Paulo")
	// 		if err != nil {
	// 			fmt.Println("timezone err: ", err.Error())
	// 		}
	// 		newLead := lead{
	// 			Name:      r.FormValue("name"),
	// 			Email:     r.FormValue("email"),
	// 			Ip:        r.RemoteAddr,
	// 			Timestamp: t.In(utc).Format("2006-01-02 15:04:05"),
	// 		}
	// 		//Mongo
	// 		s, err := mgo.Dial("mongodb://url")
	// 		if err != nil {
	// 			panic(err)
	// 		}
	// 		if err = s.Ping(); err != nil {
	// 			panic(err)
	// 		}
	// 		session := s.Copy()
	// 		defer session.Close()
	// 		entry := session.DB("nhoque").C("leads")
	// 		err = entry.Insert(newLead)
	// 		if err != nil {
	// 			if mgo.IsDup(err) {
	// 				fmt.Fprintf(w, "Already present", http.StatusBadRequest)
	// 				return
	// 			}
	// 			fmt.Fprintf(w, "Database error", http.StatusInternalServerError)
	// 			log.Errorf(ctx, "Failed insert: ", err)
	// 			return
	// 		}
	// 		//Mongo
	// 		out, _ := json.Marshal(newLead)
	// 		log.Errorf(ctx, string(out))
	// 		http.Redirect(w, r, "/thanks.html", 301)
	// 	} else {
	// 		http.Redirect(w, r, "/", 301)
	// 	}
	// })
}
