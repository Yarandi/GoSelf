package main

import (
	"crypto/sha1"
	"fmt"
	"github.com/google/uuid"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {

	http.HandleFunc("/", index)
	http.Handle("/upload/", http.StripPrefix("/upload", http.FileServer(http.Dir("./upload"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	c := getCookie(w, r)

	//get the file name and hash them
	if r.Method == http.MethodPost {
		f, h, err := r.FormFile("file")
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()
		//hash filename
		hash := sha1.New()

		//get the file name and extension
		parts := strings.Split(h.Filename, ".")
		ext := parts[len(parts)-1]

		io.Copy(hash, f)
		fname := fmt.Sprintf("%x", hash.Sum(nil)) + "." + ext
		fmt.Println(fname)

		//getting working directory
		pwd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}
		// Create subdirectories if they don't exist
		err = os.MkdirAll(filepath.Join("upload", "users", "files"), 0755)
		if err != nil {
			fmt.Println(err)
			return
		}

		//create new file in upload/users path
		path := filepath.Join(pwd, "upload", "users", "files", fname)
		nf, err := os.Create(path)
		if err != nil {
			fmt.Println(err)
		}
		//defer nf.Close()

		//copy
		// Moves the reading cursor of the uploaded file back to the beginning. This is needed before copying the file content to the newly created file.
		f.Seek(0, 0)
		//Copies the content of the uploaded file to the newly created file. This effectively saves the uploaded file with the new hashed filename and extension.
		io.Copy(nf, f)

		//add file name to users cookie
		c = appendValue(w, c, fname)
	}

	xs := strings.Split(c.Value, "|")
	tpl.ExecuteTemplate(w, "index.gohtml", xs)
}

func getCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {

	c, err := r.Cookie("session")
	if err != nil {
		//create the cookie
		SID := uuid.NewString()
		c = &http.Cookie{
			Name:  "session",
			Value: SID,
		}
		//set the cookie
		http.SetCookie(w, c)
	}
	return c
}

func appendValue(w http.ResponseWriter, c *http.Cookie, fname string) *http.Cookie {

	s := c.Value
	if !strings.Contains(s, fname) {
		s += "|" + fname
	}
	c.Value = s
	http.SetCookie(w, c)
	return c
}
