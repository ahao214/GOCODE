package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func main() {
	mySigningKey := []byte("aho214")
	c := MyClaims{
		Username: "枯藤不二",
		StandardClaims: jwt.StandardClaims{
			Audience:  time.Now().Unix() - 60,
			ExpiresAt: time.Now().Unix() - 60*60*2,
			Issuer:    "aho",
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	s, err := t.SignedString(mySigningKey)
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Println(s)
	time.Sleep(6 * time.Second)
	token, err := jwt.ParseWithClaims(s, &MyClaims{}, func(token, *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		fmt.Printf("%s", err)
	} else {
		fmt.Println(token.Claims.(*MyClaims))
		fmt.Println(token.Claims.(*MyClaims).Username)

	}

}
