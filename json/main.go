package main

import "encoding/json"
import "fmt"
import "os"

type response1 struct {
	Page   int
	Fruits []string
}
type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

var p = fmt.Println

func main() {
	bolB, _ := json.Marshal(true)
	p(string(bolB))
	intB, _ := json.Marshal(10)
	p(string(intB))
	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))
	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcD := []string{"apple", "grape", "peach"}
	slcB, _ := json.Marshal(slcD)
	p(string(slcB))
	mapD := map[string]int{"apple": 5, "lettuce": 3}
	mapB, _ := json.Marshal(mapD)
	p(string(mapB))

	res1D := &response1{
		Page:   2,
		Fruits: []string{"apple", "peach", "grape"}}
	resp1D, _ := json.Marshal(res1D)
	p(string(resp1D))

	res2D := &response2{
		Page:   3,
		Fruits: []string{"apple2", "peach2", "grape2"}}
	resp2D, _ := json.Marshal(res2D)
	p(string(resp2D))

	byt := []byte(`{"num":4.3, "strs":["a","b"]}`)
	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	p(dat)

	num := dat["num"].(float64)
	p(num)
	strs := dat["strs"].([]interface{})
	p(strs[0].(string))

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 3, "peach": 2, "tinda": 5}
	enc.Encode(d)

}
