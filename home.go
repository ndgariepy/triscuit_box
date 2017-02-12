package main

import (
	"fmt"
	"net/http"

	"google.golang.org/appengine"
)

const loginPageHtml = `<html><head><title>TrisKitBox</title></head><body> 
                        <form action=\"login\" method=\"post\"> <br><table> 
                        <tr><td>User Name</td><td><input type=\"text\" name=\"userid\" /></td></tr> 
                        <tr><td>Password</td><td><input type=\"password\" name=\"password\" /> </td></tr> </table>
                        <input type="submit" value="Login" /> <br><br>
                        <a href="http://http://triskitbox-158401.appspot.com/signup">Sign Up</a>
                        </form> </body> </html>`

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "TrisKitBox, the best place on the web to track your wagers.\n"+loginPageHtml)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
	appengine.Main()
}
