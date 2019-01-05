package tag_api

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func handleAuthTestpage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var testframe string = `<!DOCTYPE html>
<html>
<head>
  <script src="https://ajax.googleapis.com/ajax/libs/jquery/3.2.1/jquery.min.js"></script>
  <script>
    $(document).ready(function(){
      $("#authBasic").click(function(){
        $.getJSON({
            url: "/authenticate",
            type: "POST",
            beforeSend: setHeaderAuthBasic,
            success: function(result) { document.getElementById("response").innerHTML = JSON.stringify(result); },
            error: function() { alert("Request error"); }
        });
      });
      $("#authPremium").click(function(){
        $.getJSON({
            url: "/authenticate",
            type: "POST",
            beforeSend: setHeaderAuthPremium,
            success: function(result) { document.getElementById("response").innerHTML = JSON.stringify(result); },
            error: function() { alert("Request error"); }
        });
      });
    });

    function setHeaderAuthBasic(xhr) {
      xhr.setRequestHeader("Authorization", "Bearer eyJhbGciOiJBMTI4S1ciLCJlbmMiOiJBMTI4R0NNIn0.xlIzeZcOi2JKi3TTvetA4SIJElZk09xo.68Zod7rdo3353I75.Yom0Fny7lCWkATFHbESpfdmDiT2OM7JYNesP65lKXw0U-b2obWyT1Z_q2V0bRwaIoNrhdP4Y.e6XmMjFp-DfoswLV40kP_A");
    }
    function setHeaderAuthPremium(xhr) {
      xhr.setRequestHeader("Authorization", "Bearer eyJhbGciOiJBMTI4S1ciLCJlbmMiOiJBMTI4R0NNIn0.hog6SGNfUyzDpvG5QED1NwQkUCcTIM9Z.fK5rSSShlN7Qqg2_.-ZmqyI5B13iot9aW5GSxBrb1rMAWnQPA-nhdu_0M6WLNASkFRyRqSQAhXtJiUBfsfVPKqUj9cg.-_zi4UsMl7O-VudPYevxCg");
    }
  </script>
</head>
<body>
    <h2>Test Framework</h2>
    <ul>
        <li><button id="authBasic">POST to /authenticate</button> (Basic)</li>
        <li><button id="authPremium">POST to /authenticate</button> (Premium)</li>
    </ul>
    <h3>Response</h3>
    <div id="response"></div>
</body>
</html>`

	// Reply status
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, testframe)
}
