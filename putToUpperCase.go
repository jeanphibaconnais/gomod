package gomod

import (
	"strings"
	"github.com/robfig/cron"
	"fmt"
)

func PutToUppercase(s string) {
	c := cron.New()
	c.AddFunc("@every 1s", func() {
		fmt.Println("Every each second i display the world to upper case : " + strings.ToUpper(s))
	})
	c.Start()

	inspect(c.Entries())

	var Test string
	fmt.Scan(&Test)

	//return strings.ToUpper(s)
}

func inspect(entries []*cron.Entry) {

	for _, entry := range entries {
		entry.Job.Run();
	}
}
