package main

import (
	"encoding/json"
	"fmt"
	"github.com/codegangsta/cli"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type jsonRespEmail struct {
	Title      string
	Name       string
	Domain     string
	BreachDate string
	PwnCount   int
	IsVerified bool
}

func main() {
	respcodes := map[int]string{
		400: "Bad request — the account does not comply with an acceptable format (i.e. it's an empty string)",
		403: "Forbidden — no user agent has been specified in the request",
		404: "Not found — the account could not be found and has therefore not been pwned",
	}

	app := cli.NewApp()
	app.Name = "iGOtpwned"
	app.Usage = "'Have I been pwned?' golang cli checker app"
	app.Version = "0.0.1"
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Ye Myat Kaung (Maverick)",
			Email: "mavjs01@gmail.com"},
	}
	app.Commands = []cli.Command{
		{
			Name:    "email",
			Aliases: []string{"m"},
			Usage:   "email address to look up all breaches associated with it",
			Action: func(c *cli.Context) {
				// create http client
				client := new(http.Client)

				// request http api
				req, err := http.NewRequest("Get", "https://haveibeenpwned.com/api/v2/breachedaccount/"+c.Args().First(), nil)
				if err != nil {
					log.Fatal(err)
					os.Exit(1)
				}

				// set haveibeenpwned content negotiation header
				req.Header.Add("Accept", "application/vnd.haveibeenpwned.v2+json")
				req.Header.Add("User-Agent", "iGotpwned golang cli app")

				// make the request
				res, err := client.Do(req)
				if err != nil {
					log.Fatal(err)
				}

				// return status codes and exit
				if res.StatusCode == 400 {
					fmt.Println(respcodes[res.StatusCode])
					os.Exit(1)
				} else if res.StatusCode == 403 {
					fmt.Println(respcodes[res.StatusCode])
					os.Exit(1)
				} else if res.StatusCode == 404 {
					fmt.Println(respcodes[res.StatusCode])
					os.Exit(1)
				}

				// read body
				body, err := ioutil.ReadAll(res.Body)
				defer res.Body.Close()
				if err != nil {
					log.Fatal(err)
				}

				var email []jsonRespEmail
				err = json.Unmarshal(body, &email)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Printf("Breaches for %s: %+v\n", c.Args().First(), email)
			},
		},
		{
			Name:    "site",
			Aliases: []string{"s"},
			Usage:   "info associated with a single breached site",
			Action: func(c *cli.Context) {
				fmt.Println("breached site: ", c.Args().First())
			},
		},
	}

	app.Run(os.Args)
}
