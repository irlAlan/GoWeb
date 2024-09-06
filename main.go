package main

import (
	"fmt"
	//	"html/template"
	"log"
	"net/http"
  "strings"

	pages "GoWeb/pages"
  "GoWeb/pages/blog/stem"
  "GoWeb/pages/blog"

	"github.com/a-h/templ"
	// "os"
)


func postHandleSearch(w http.ResponseWriter, r *http.Request){
  srch := r.PostFormValue("search");
  fmt.Println("Search: " + srch);
  srch = strings.ToLower(srch);

  redirect := ""
  switch srch{
  case "Projects": redirect = "/projects"
  case "Home" : redirect = "/"
  case "Blog" : redirect = "/blog"
  default: redirect="/ErrorPage";
  }

  fmt.Println("Redirect: " + redirect);
  http.Redirect(w,r,redirect,http.StatusSeeOther);
   // component := pages.Page("from projects search");
   // component.Render(r.Context(), w);
}



func main(){
  fmt.Println("Hello there");
  //handle1 := func (writer http.ResponseWriter, receiver *http.Request){
    // component := hello("there");
    // component.Render(context.Background(), writer);

    // templ.Handler(component);
    // Index().Render(context.Background(), writer);
   // templ.Handler(Index());

    // templ := template.Must(template.ParseFiles("../index.html"));
    //templ.Execute(writer, nil);
  //}

  // for parsing html into a template that can be shown on the web 
//  handleProjects := func (writer http.ResponseWriter, receiver *http.Request){
//    templ := template.Must(template.ParseFiles("../index.html"));
//    templ.Execute(writer, nil);
//  }

  port := "8000";
  fileServer := http.FileServer(http.Dir("/home/ubuntu/GoWeb/wwwroot"));
  http.Handle("/wwwroot/*", http.StripPrefix("/wwwroot/", fileServer));
  // http.Handle("/alan", templ.Handler(pages.Index()));

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
    if r.Method == http.MethodPost{
      postHandleSearch(w,r);
      return
    }
    fmt.Println("is going to index");
    if r.Method == http.MethodGet{
      // http.Handle("/", templ.Handler(pages.Index()));
      component := pages.Index();
      component.Render(r.Context(), w);
      return
    }
  })

   http.Handle("/projects", templ.Handler(pages.Projects()));
   http.Handle("/blog", templ.Handler(blog.BlogPage()));
   http.Handle("/blog/stem/", templ.Handler(stem.StemPage()));
   http.Handle("/blog/stem/pdfs", templ.Handler(stem.StemPdfPage()))
   http.Handle("/ErrorPage", templ.Handler(pages.ErrorPage()))
   http.Handle("/contactme", templ.Handler(pages.ContactMe()))
  //http.HandleFunc("/Project", handleProjects);
  log.Println("Listening on port " + port);
  log.Fatal(http.ListenAndServe(":"+port, nil));
}
