package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/Hao1995/sql-performance/glob"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db   *sql.DB
	size int
)

type users struct {
	id   int
	name string
	age  int
}

func init() {

	size = glob.CfgData.Data.Size

	dbTmp, err := sql.Open("mysql", glob.CfgData.Mysql.User+":"+glob.CfgData.Mysql.Password+"@tcp("+glob.CfgData.Mysql.Host+":"+glob.CfgData.Mysql.Port+")/"+glob.CfgData.Mysql.Name)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Print(glob.CfgData)

	db = dbTmp
}

func main() {
	insert()
	query()
	update()
	query()
	delete()
}

func query() {
	fmt.Println("[Query]")

	// Method-1 : db.Query
	start := time.Now()
	rows, err := db.Query("SELECT `id`, `name`, `age` FROM users")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		item := users{}
		if err := rows.Scan(&item.id, &item.name, &item.age); err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("%+v\n", item)
	}
	end := time.Now()
	fmt.Println("Method-1 db.Query. total time = ", end.Sub(start).Seconds())

	// Method-2 db.Prepare -> Query
	start = time.Now()
	stm, err := db.Prepare("SELECT `id`, `name`, `age` FROM users")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer stm.Close()
	rows, err = stm.Query()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		item := users{}
		if err := rows.Scan(&item.id, &item.name, &item.age); err != nil {
			log.Fatal(err)
		}
	}
	end = time.Now()
	fmt.Println("Method-2 db.Prepare -> Query. total time = ", end.Sub(start).Seconds())

	// Method-3 db Transaction
	start = time.Now()
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer tx.Commit()
	rows, err = tx.Query("SELECT `id`, `name`, `age` FROM users")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		item := users{}
		if err := rows.Scan(&item.id, &item.name, &item.age); err != nil {
			log.Fatal(err)
		}
	}
	end = time.Now()
	fmt.Println("Method-3 db Transaction. total time = ", end.Sub(start).Seconds())
}

func insert() {
	fmt.Println("[Insert]")

	// Method-1 Exec within the loop
	start := time.Now()
	for i := size*1 + 1; i <= size*2; i++ {
		// New a connection every time. Worest performance.
		_, err := db.Exec("INSERT INTO `users`(`id`, `name`, `age`) values(?, ?, ?)", i, "user"+strconv.Itoa(i), i)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	end := time.Now()
	fmt.Println("Method-1 Exec within the loop total time = ", end.Sub(start).Seconds())

	// Method-2 Prepare and Exec within the loop.
	start = time.Now()
	for i := size*2 + 1; i <= size*3; i++ {
		// Prepare will new a connection every time. Worest performance.
		stm, err := db.Prepare("INSERT INTO `users`(`id`, `name`, `age`) values(?, ?, ?)")
		if err != nil {
			log.Fatal(err)
			return
		}
		_, err = stm.Exec(i, "user"+strconv.Itoa(i), i)
		if err != nil {
			log.Fatal(err)
			return
		}
		stm.Close()
	}
	end = time.Now()
	fmt.Println("Method-2 Prepare and Exec within the loop. total time = ", end.Sub(start).Seconds())

	// Method-3 Prepare. Then Exec within the loop.
	start = time.Now()
	stm, err := db.Prepare("INSERT INTO `users`(`id`, `name`, `age`) values(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
		return
	}
	for i := size*3 + 1; i <= size*4; i++ {
		// Why it's performance so bad even the Exec function do not new a connection.
		_, err := stm.Exec(i, "user"+strconv.Itoa(i), i)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	stm.Close()
	end = time.Now()
	fmt.Println("Method-3 Prepare. Then Exec within the loop. total time = ", end.Sub(start).Seconds())

	// Method-4 DB Transaction. Then Euec within the loop.
	start = time.Now()
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return
	}
	for i := size*4 + 1; i <= size*5; i++ {
		// No new a connection within 'tx'. High performance
		_, err := tx.Exec("INSERT INTO `users`(`id`, `name`, `age`) values(?, ?, ?)", i, "user"+strconv.Itoa(i), i)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	tx.Commit()
	end = time.Now()
	fmt.Println("Method-4 DB Transaction. Then Euec within the loop. total time: = ", end.Sub(start).Seconds())

	// Method-5 DB Transaction within the Loop.
	start = time.Now()
	for i := size*5 + 1; i <= size*6; i++ {
		// New a connection every db.Begin(). Worest Performance
		tx, err := db.Begin()
		if err != nil {
			log.Fatal(err)
			return
		}
		_, err = tx.Exec("INSERT INTO `users`(`id`, `name`, `age`) values(?, ?, ?)", i, "user"+strconv.Itoa(i), i)
		if err != nil {
			log.Fatal(err)
			return
		}
		err = tx.Commit()
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	end = time.Now()
	fmt.Println("Method-5 DB Transaction within the Loop. total time = ", end.Sub(start).Seconds())
}

func update() {
	fmt.Println("[Update]")

	// Method-1 Exec within the loop
	start := time.Now()
	for i := size*1 + 1; i <= size*2; i++ {
		_, err := db.Exec("UPDATE `users` SET `name`=?, `age`=? WHERE `id` = ?", "user"+strconv.Itoa(i), i, i)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	end := time.Now()
	fmt.Println("Method-1 Exec within the loop. total time = ", end.Sub(start).Seconds())

	// Method-2 Prepare and Exec within the loop.
	start = time.Now()
	for i := size*2 + 1; i <= size*3; i++ {
		stm, err := db.Prepare("UPDATE `users` SET `name`=?, `age`=? WHERE `id` = ?")
		if err != nil {
			log.Fatal(err)
			return
		}
		_, err = stm.Exec("user"+strconv.Itoa(i), i, i)
		if err != nil {
			log.Fatal(err)
			return
		}
		err = stm.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	end = time.Now()
	fmt.Println("Method-2 Prepare and Exec within the loop.. total time = ", end.Sub(start).Seconds())

	// Method-3 Prepare. Then Exec within the loop.
	start = time.Now()
	stm, err := db.Prepare("UPDATE `users` SET `name`=?, `age`=? WHERE `id` = ?")
	if err != nil {
		log.Fatal(err)
		return
	}
	for i := size*3 + 1; i <= size*4; i++ {
		_, err := stm.Exec("user"+strconv.Itoa(i), i, i)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	stm.Close()
	end = time.Now()
	fmt.Println("Method-3 Prepare. Then Exec within the loop. total time = ", end.Sub(start).Seconds())

	// Method-4 DB Transaction. Then Euec within the loop.
	start = time.Now()
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return
	}
	for i := size*4 + 1; i <= size*5; i++ {
		_, err := tx.Exec("UPDATE `users` SET `name`=?, `age`=? WHERE `id` = ?", "user"+strconv.Itoa(i), i, i)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
		return
	}

	end = time.Now()
	fmt.Println("Method-4 DB Transaction. Then Euec within the loop. total time = ", end.Sub(start).Seconds())

	// Method-5 DB Transaction within the loop.
	start = time.Now()
	for i := size*5 + 1; i <= size*6; i++ {
		tx, err := db.Begin()
		if err != nil {
			log.Fatal(err)
			return
		}
		_, err = tx.Exec("UPDATE `users` SET `name`=?, `age`=? WHERE `id` = ?", "user"+strconv.Itoa(i), i, i)
		if err != nil {
			log.Fatal(err)
			return
		}
		err = tx.Commit()
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	end = time.Now()
	fmt.Println("Method-5 DB Transaction within the loop. total time = ", end.Sub(start).Seconds())

}

func delete() {
	fmt.Println("[Delete]")

	// Method-1 Exec Delete
	start := time.Now()
	for i := size*1 + 1; i <= size*2; i++ {
		_, err := db.Exec("DELETE FROM users WHERE `id` = ?", i)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	end := time.Now()
	fmt.Println("Method-1 Exec Delete. total time = ", end.Sub(start).Seconds())

	// Method-2 Prepare and Exec within the loop.
	start = time.Now()
	for i := size*2 + 1; i <= size*3; i++ {
		stm, err := db.Prepare("DELETE FROM users WHERE `id` = ?")
		if err != nil {
			log.Fatal(err)
			return
		}
		_, err = stm.Exec(i)
		if err != nil {
			log.Fatal(err)
			return
		}
		err = stm.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	end = time.Now()
	fmt.Println("Method-2 Prepare and Exec within the loop.. total time:", end.Sub(start).Seconds())

	// Method-3 Prepare. Then Exec within the loop.
	start = time.Now()
	stm, err := db.Prepare("DELETE FROM users WHERE `id` = ?")
	if err != nil {
		log.Fatal(err)
		return
	}
	for i := size*3 + 1; i <= size*4; i++ {
		_, err := stm.Exec(i)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	err = stm.Close()
	if err != nil {
		log.Fatal(err)
		return
	}
	end = time.Now()
	fmt.Println("Method-3 Prepare. Then Exec within the loop. total time = ", end.Sub(start).Seconds())

	// Method-4 db Transaction. Exec Delete within the loop.
	start = time.Now()
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return
	}
	for i := size*4 + 1; i <= size*5; i++ {
		_, err := tx.Exec("DELETE FROM users WHERE `id` = ?", i)
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
		return
	}
	end = time.Now()
	fmt.Println("Method-4 db Transaction. Exec Delete within the loop. total time = ", end.Sub(start).Seconds())

	// Method-5 DB Transaction within the loop.
	start = time.Now()
	for i := size*5 + 1; i <= size*6; i++ {
		tx, err := db.Begin()
		if err != nil {
			log.Fatal(err)
			return
		}
		_, err = tx.Exec("DELETE FROM users WHERE `id` = ?", i)
		if err != nil {
			log.Fatal(err)
			return
		}
		err = tx.Commit()
		if err != nil {
			log.Fatal(err)
			return
		}
	}
	end = time.Now()
	fmt.Println("Method-5 DB Transaction within the loop. total time = ", end.Sub(start).Seconds())
}
