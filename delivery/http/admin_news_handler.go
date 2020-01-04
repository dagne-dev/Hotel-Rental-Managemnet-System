package http


import (
	"fmt"
	"html/template"
	"net/http"
	"github.com/getme04/Hotel-Rental-Managemnet-System/news"
)
// AdminNewsHandler handles news handler admin requests
type AdminNewsHandler struct {
	tmpl        *template.Template
	newsService news.NewsService
}

// NewAdminNewsHandler initializes and returns new AdminNewsHandler
func NewAdminNewsHandler(T *template.Template, NS news.NewsService) *AdminNewsHandler {
	return &AdminNewsHandler{tmpl: T, newsService: NS}
}

// AdminNews handle requests on route /admin/news
func (ach *AdminNewsHandler) AdminNews(w http.ResponseWriter, r *http.Request) {
	news, err := ach.newsService.News()
	fmt.Printf("dfgtg")
	if err != nil {
		fmt.Printf("errrr")
		fmt.Printf("hhhh1111")
		panic(err)
	}
	fmt.Printf("hhhh2222")
	ach.tmpl.ExecuteTemplate(w, "admin.news.layout", news)
	fmt.Printf("hhhh33332")
}

//
//// AdminNewsNew hanlde requests on route /admin/new
//func (ach *AdminNewsHandler) AdminNewsNew(w http.ResponseWriter, r *http.Request) {
//
//	if r.Method == http.MethodPost {
//		////////////////////////////
//		/////////////////////
//		////////////////////////////////////////form value
//		ctg := entity.News{}
//		ctg.Header = r.FormValue("name")
//		ctg.Description = r.FormValue("description")
//		mf, fh, err := r.FormFile("catimg")
//		if err != nil {
//			panic(err)
//		}
//		defer mf.Close()
//
//		ctg.Image = fh.Filename
//
//		WriteFile(&mf, fh.Filename)
//
//		err = ach.newsService.StoreNews(ctg)
//
//		if err != nil {
//			panic(err)
//		}
//
//		http.Redirect(w, r, "/admin/news", http.StatusSeeOther)
//
//	} else {
//
//		ach.tmpl.ExecuteTemplate(w, "admin.news.new.layout", nil)
//
//	}
//}
//
//// AdminNewsUpdate handle requests on /admin/news/update
//func (ach *AdminNewsHandler) AdminNewsUpdate(w http.ResponseWriter, r *http.Request) {
//
//	if r.Method == http.MethodGet {
//
//		idRaw := r.URL.Query().Get("id")
//		id, err := strconv.Atoi(idRaw)
//
//		if err != nil {
//			panic(err)
//		}
//
//		cat, err := ach.newsService.NewsById(id)
//
//		if err != nil {
//			panic(err)
//		}
//		fmt.Printf("\nhello update\n")
//		ach.tmpl.ExecuteTemplate(w, "admin.news.update.layout", cat)
//		fmt.Printf("update excuted\n")
//
//	} else if r.Method == http.MethodPost {
//
//		ctg := entity.News{}
//		ctg.Id, _ = strconv.Atoi(r.FormValue("id"))
//		ctg.Header = r.FormValue("name")
//		ctg.Description = r.FormValue("description")
//		ctg.Image = r.FormValue("image")
//
//		mf, _, err := r.FormFile("catimg")
//
//		if err != nil {
//			panic(err)
//		}
//		defer mf.Close()
//
//		WriteFile(&mf, ctg.Image)
//
//		err = ach.newsService.UpdateNews(ctg)
//
//		if err != nil {
//			panic(err)
//		}
//		fmt.Printf("\nupdated directed to home\n")
//		http.Redirect(w, r, "/admin/news", http.StatusSeeOther)
//
//	} else {
//		fmt.Printf("\n else updated directed to home\n")
//		http.Redirect(w, r, "/admin/news", http.StatusSeeOther)
//	}
//
//}
//
//// AdminNewsDelete handle requests on route /admin/categories/delete
//func (ach *AdminNewsHandler) AdminNewsDelete(w http.ResponseWriter, r *http.Request) {
//
//	if r.Method == http.MethodGet {
//
//		idRaw := r.URL.Query().Get("id")
//
//		id, err := strconv.Atoi(idRaw)
//
//		if err != nil {
//			panic(err)
//		}
//
//		err = ach.newsService.DeleteNews(id)
//
//		if err != nil {
//			panic(err)
//		}
//
//	}
//
//	http.Redirect(w, r, "/admin/news", http.StatusSeeOther)
//}
//
//func WriteFile(mf *multipart.File, fname string) {
//
//	wd, err := os.Getwd()
//
//	if err != nil {
//		panic(err)
//	}
//
//	path := filepath.Join(wd, "ui", "assets", "img", fname)
//	image, err := os.Create(path)
//
//	if err != nil {
//		panic(err)
//	}
//	defer image.Close()
//	io.Copy(image, *mf)
//}
//
