package custom_errors

import (
	"net/http"
)

// HttpSearchParameters http error message for when the parameters or its values are wrong.
func HttpSearchParameters(w http.ResponseWriter) {
	http.Error(w, "Search must contain the valid amount of search parameter(s) "+
		"with a valid value. See the documentation.", http.StatusBadRequest)
}

// HttpErrorFromBackendApi http error message for when the backend apis returns an error.
func HttpErrorFromBackendApi(w http.ResponseWriter) {
	http.Error(w, "Error from backend api", http.StatusBadGateway)
}

// HttpUnsupportedMethod http error message for when the rest method is invalid.
func HttpUnsupportedMethod(w http.ResponseWriter) {
	http.Error(w, "Method not supported.", http.StatusMethodNotAllowed)
}

// HttpUnknownServerError http error message for when the web_server has an undefined error or an error the user should not know.
func HttpUnknownServerError(w http.ResponseWriter) {
	http.Error(w, "Server side error, please try again later", http.StatusInternalServerError)
}

// HttpWrongJsonInfo http error for when the user provides the wrong json structure when posting.
func HttpWrongJsonInfo(w http.ResponseWriter) {
	http.Error(w, "Some information is missing from the json body, see the documentation.", http.StatusBadRequest)
}

// HttpNotFound gives an error message that the endpoint does not exist.
func HttpNotFound(w http.ResponseWriter) {
	http.Error(w, "The endpoint does not exist, please see the documentation: "+
		"https://git.gvk.idi.ntnu.no/course/prog2005/prog2005-2022-workspace/mathias_ws/assignment-2/-/blob/main/README.md",
		http.StatusNotFound)
}

// HttpNoPolicy gives an error message that the policy endpoint did not find any data matching the search criteria.
func HttpNoPolicy(w http.ResponseWriter) {
	http.Error(w, "The policy country and scope did not return any valid data.", http.StatusNotFound)
}
