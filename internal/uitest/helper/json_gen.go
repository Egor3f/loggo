package helper

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

func JsonGenerator(writer io.Writer) {
	b, err := ioutil.ReadFile("internal/testdata/test1.json")
	if err != nil {
		panic(err)
	}
	jm := make(map[string]interface{})
	_ = json.Unmarshal(b, &jm)
	i := 0
	for {
		//if i != 0 && i%(rand.Intn(27)+1) == 0 {
		//	_, _ = fmt.Fprintln(writer, "bad json")
		//	i++
		//	continue
		//}
		i++
		uid1 := uuid.New().String()
		uid2 := uuid.New().String()
		id3 := rand.Intn(30)
		jm["severity"] = []string{"INFO", "ERROR", "DEBUG", "WARN"}[rand.Intn(4)]
		jm["insertId"] = uid1
		jm["trace"] = uid2
		jm["spanId"] = fmt.Sprintf("%d", id3)
		jm["timestamp"] = time.Now().Format("2006-01-02T15:04:05-0700")
		b, _ = json.Marshal(jm)
		_, _ = fmt.Fprintln(writer, string(b))
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(800)))
		if i%(rand.Intn(5)+1) == 0 {
			time.Sleep(time.Second * time.Duration(rand.Intn(5)))
		}

	}
}
