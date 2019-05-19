# JWT

```go
package main

import "fmt"

func main() {
	res, err := http.PostForm(
		fmt.Sprintf("http://localhost:%v/authenticate", serverPort),
		url.Values{
			"user": "test",
			"pass": "<secret>.^<password>",
		},
	)

	if err != nil  {
		fatal(err)
	}

	if res.StatusCode != 200 {
		fmt.Println("Unexpected status code", res.StatusCode)
	}

	// Read the token out of the response body
	buf := new(bytes.Buffer)
	io.Copy(buf, res.Body)
	res.Body.Close()
	tokenString := strings.TrimSpace(buf.String())

	// Parse the token
	token, err := jwt.ParseWithClaims(
		tokenString,
		&CustomerClaimsExample{},
		func(token *jwt.Token)(interface{}, error){
			return verifyKey, nil
	})
	fatal(err)

	claims := token.Claims.(*CustomerClaimsExample)
	fmt.Println(claims.CustomerInfo.Name)
}
```

- Use Token

```go
token, err := createToken("string-to-generate-token")
fatal(err)

req, err := http.NewRequest(
	"GET",
	fmt.Sprintf("http://localhost:%v/restricted", serverPort),
	nil,
)
fatal(err)
req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
res, err := http.DefaultClient.Do(req)
fatal(err)

// Read the response body
defer res.Body.Close()
buf := new(bytes.Buffer)
io.Copy(buf, res.Body)
fmt.Println(buf.String())
```
