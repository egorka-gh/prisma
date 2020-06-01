package main

import (
	"fmt"
	"time"

	"github.com/egorka-gh/sm/prisma"
)

func main() {
	client := prisma.NewClient("127.0.0.1")
	m := &prisma.Message{
		Prefix:     "TSD",
		Number:     11,
		Mode:       333,
		CassirItem: "staff",
		Cassir:     "Калоша Ирина Викторовна",
		CKNumber:   "0",
		Count:      1,
		BarCode:    "4813566002120",
		GoodsItem:  "0803180",
		GoodsName:  "Пакет-майка с логотипом",
		//GoodsName:  "paket-maika s logotipom",
		GoodsPrice: 77,
		GoodsQuant: 2.2,
		Year:       "20",
		Month:      "5",
		Day:        "22",
	}

	for {

		err := client.Send(m)
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(3 * time.Second)
	}
}
