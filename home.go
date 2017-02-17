package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"google.golang.org/appengine"
)
import _ "github.com/go-sql-driver/mysql"

const loginPageHtml = `<html><head><title>TrisKitBox</title></head><body> 
                        <form action=\"login\" method=\"post\"> <br><table> 
                        <tr><td>User Name</td><td><input type=\"text\" name=\"userid\" required/></td></tr> 
                        <tr><td>Password</td><td><input type=\"password\" name=\"password\" required/> </td></tr> </table>
                        <input type="submit" value="Login" /> <br><br>
                        <a href="http://triskitbox-158401.appspot.com/signup">Sign Up</a>
                        </form> </body> </html>`

const SignupPageHtml = `<html><head><title>TrisKitBox - Signup</title></head><body> 
                        <form action=\"login\" method=\"post\"> <br><table> 
                        <tr><td>User Name</td><td><input type=\"text\" name=\"userid\" required/></td></tr> 
                        <tr><td>Email Address</td><td><input type="email" name="email" required/></td></tr> 
                        <tr><td>Password</td><td><input type=\"password\" name=\"password\" required/> </td></tr> </table>
                        <input type="submit" value="Login" /> <br><br>
                        </form> </body> </html>`

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "TrisKitBox, the best place on the web to track your wagers.\n"+loginPageHtml)
}

func signupPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Fprint(w, SignupPageHtml)
	} else {
		username := r.FormValue("userid")
		email := r.FormValue("email")
		pwd := r.FormValue("password")
		connectionName := os.Getenv("CLOUDSQL_CONNECTION_NAME")
		user := os.Getenv("CLOUDSQL_USER")
		password := os.Getenv("CLOUDSQL_PASSWORD")
		db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@cloudsql(%s)/", user, password, connectionName))
		_, err = db.Exec("INSERT INTO triskitbox.Users (suserid, semailaddr, spasswd) VALUES (?, ?, ?)", username, email, pwd)
		http.Redirect(w, r, "/", 301)
		if err != nil {
			http.Error(w, fmt.Sprintf("Could not open db: %v", err), 500)
			return
		}
		rows, err := db.Query("SELECT COUNT(*) FROM triskitbox.Users;")
		if err != nil {
			http.Error(w, fmt.Sprintf("Could not query db: %v", err), 500)
			return
		}
		for rows.Next() {
			var count int
			err = rows.Scan(&count)
			fmt.Fprint(w, count)
		}
	}
}

func main() {
	http.HandleFunc("/signup", signupPage)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
	appengine.Main()
}
