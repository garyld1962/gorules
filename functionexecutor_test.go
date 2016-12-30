package gorules

// import (
// 	"bytes"
// 	"log"
// 	"os/exec"
// 	"testing"
// )

// func TestExternalFile(t *testing.T) {

// 	cmd := exec.Command("cmd", "/c", "date")
// 	stdout, err := cmd.StdoutPipe()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if err := cmd.Start(); err != nil {
// 		log.Fatal(err)
// 	}
// 	defer cmd.Wait()
// 	buf := new(bytes.Buffer)
// 	// buf.ReadFrom(response.Body)

// 	// // s := string(b)
// 	// fmt.Println(s)

// 	// layout := "2006-01-02 15:04:05"
// 	// ti, _ := time.Parse(layout, s)
// 	// fmt.Println(ti.Format("2006-01-02"))
// 	//assert.Equal(ti, out, "The current date is: Tue 12/27/2016", "values are not equal")

// }
