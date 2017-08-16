package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/mysqlerr"
	_ "github.com/mysqlerr"
	"log"
	"reflect"
)

var (
	id        int
	last_name string
)

/*
	sql.DB不是一个连接，它是数据库的抽象接口。它可以根据driver打开关闭数据库连接，管理连接池。正在使用的连接被标记为繁忙，
	用完后回到连接池等待下次使用。所以，如果你没有把连接释放回连接池，会导致过多连接使系统资源耗尽。

	sql.DB的设计就是用来作为长连接使用的。不要频繁Open, Close。比较好的做法是，为每个不同的datastore建一个DB对象，
	保持这些对象Open。如果需要短连接，那么把DB作为参数传入function，而不要在function中Open, Close。
*/
/**
关于连接池
	避免错误操作，例如LOCK TABLE后用 INSERT会死锁，因为两个操作不是同一个连接，insert的连接没有table lock。
	当需要连接，且连接池中没有可用连接时，新的连接就会被创建。
	默认没有连接上限，你可以设置一个，但这可能会导致数据库产生错误“too many connections”
	db.SetMaxIdleConns(N)设置最大空闲连接数
	db.SetMaxOpenConns(N)设置最大打开连接数
	长时间保持空闲连接可能会导致db timeout
*/
func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/lilinlin")
	if err != nil {
		//log.Fatal(err)
		if driverErr, ok := err.(*mysql.MySQLError); ok {
			if driverErr.Number == mysqlerr.ER_ACCESS_DENIED_ERROR {
				log.Fatal("Handle the permission-denied error")
			} else if driverErr.Number == mysqlerr.ER_NO_DB_ERROR {
				log.Fatal("Handle the no this db error")
			} else if driverErr.Number == mysqlerr.ER_BAD_DB_ERROR {
				log.Fatal("Handle the Unknown database error")
			} else if driverErr.Number == 44702 {
				log.Fatal("Handle 44702 error")
			}
		}
	}

	//update(db)

	//updateInTransaction(db)

	//readData(db)

	readUnknowColumnSize(db)

	//test(db)

	//readSingleData(db)

	defer db.Close()
}

func test(db *sql.DB) {
	rows, err := db.Query("select id, last_name, address from author where id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}

	cols, err := rows.Columns() // Remember to check err afterwards
	values := make([]sql.RawBytes, len(cols))
	for i, col := range values {
		fmt.Println(cols[i], string(col))
	}

	fmt.Println("-----------")

	vals := make([]interface{}, len(cols))
	//vals := make([][]byte, len(cols))
	for i, _ := range cols {
		vals[i] = new(sql.RawBytes)
	}

	for rows.Next() {
		err = rows.Scan(vals...) //vals...表示把其作为不定参数传入

		if err != nil {
			log.Fatal(err)
		}

		for index, value := range vals {
			tmp := value.([]byte)
			fmt.Print("index:", index)
			//fmt.Println(" value:", value)

			if value == nil {
				value = "NULL"
			} else {
				fmt.Println(" type:", reflect.TypeOf(value))
				value = string(tmp)
			}
			//fmt.Println(" value:", value)

		}
	}
}

/**
read数据
*/
func readData(db *sql.DB) {
	/*如果id为1的不存在，err为sql.ErrNoRows
	    if err != nil {
			if err == sql.ErrNoRows {
				// there were no rows, but otherwise no error occurred
			} else {
				log.Fatal(err)
			}
		}
	*/
	rows, err := db.Query("select id, last_name, address from author where id < ?", 14)
	if err != nil {
		log.Fatal(err)
	}

	/*
		结果集(rows)未关闭前，底层的连接处于繁忙状态。当遍历读到最后一条记录时，会发生一个内部EOF错误，自动调用rows.Close()，
		但是如果提前退出循环，rows不会关闭，连接不会回到连接池中，连接也不会关闭。所以手动关闭非常重要。rows.Close()可以   多次   调用，是无害操作。
	*/
	defer rows.Close()

	// columns
	cln, clnError := rows.Columns()
	if clnError != nil {
		log.Fatal(clnError)
	}
	log.Println(cln)

	////columnType
	//clnType, clnTypeError := rows.ColumnTypes()
	//if clnTypeError != nil {
	//	log.Fatal(clnTypeError)
	//}
	////i := clnType[0].ScanType().NumField()
	////log.Println(clnType[0].ScanType().Field(i).Name)
	//log.Println(clnType)

	/*
		如果循环中发生错误会自动运行rows.Close()，用rows.Err()接收这个错误，Close方法可以多次调用。循环之后判断error是非常必要的。

	*/
	for rows.Next() {
		var address sql.NullString

		err := rows.Scan(&id, &last_name, &address)

		if err != nil {
			log.Fatal(err)
		}

		// Valid is true if String is not NULL
		if !address.Valid {
			fmt.Printf("address is null and id=%v last_name=%v\n", id, last_name)
		} else {
			fmt.Println(id, last_name)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

/**
读取单条数据
*/
func readSingleData(db *sql.DB) {
	var name string
	err := db.QueryRow("select last_name as name from author where id = ?", 1).Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("[readSingleData]Id = 1, Name:", name)

}

/**
读取未知列数的数据
*/
func readUnknowColumnSize(db *sql.DB) {
	rows, err := db.Query("select id, last_name, address from author where id < ?", 6)
	if err != nil {
		fmt.Println("Failed to run query", err)
		return
	}

	cols, err := rows.Columns()
	if err != nil {
		fmt.Println("Failed to get columns", err)
		return
	}

	columnNameResult := make([]string, len(cols))
	for i := 0; i < len(cols); i++ {
		columnNameResult[i] = cols[i]
	}
	fmt.Printf("%v\n", columnNameResult)

	// Result is your slice string.
	rawResult := make([][]byte, len(cols))
	dest := make([]interface{}, len(cols)) // A temporary interface{} slice
	for i, _ := range rawResult {
		dest[i] = &rawResult[i] // Put pointers to each string in the interface slice；因为rows.scan必须操作一个interface{}的对象
	}

	result := make([]string, len(cols))
	for rows.Next() {
		err = rows.Scan(dest...) //vals...表示把其作为不定参数传入
		if err != nil {
			fmt.Println("Failed to scan row", err)
			return
		}

		for i, raw := range rawResult {
			if raw == nil {
				result[i] = "Null"
			} else {
				result[i] = string(raw)
			}
		}

		fmt.Printf("%#v\n", result)
	}

}

/**
更新数据
*/
func update(db *sql.DB) {
	stmt, err := db.Prepare("INSERT INTO author(first_name,last_name) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec("Jim", "Jim")
	if err != nil {
		log.Fatal(err)
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Insert ID = %d, affected = %d\n", lastId, rowCnt)
}

/*
1.db.Begin()开始事务，Commit() 或 Rollback()关闭事务。Tx从连接池中取出一个连接，在关闭之前都是使用这个连接。Tx不能和DB层的BEGIN, COMMIT混合使用。
2.在数据库层面，Prepared Statements是和单个数据库连接绑定的。客户端发送一个有占位符的statement到服务端，服务器返回一个statement ID，然后客户端发送ID和参数来执行statement。
3.在GO中，连接不直接暴露，你不能为连接绑定statement，而是只能为DB或Tx绑定。database/sql包有自动重试等功能。当你生成一个Prepared Statement:
	- 自动在连接池中绑定到一个空闲连接
	- Stmt对象记住绑定了哪个连接
	- 执行Stmt时，尝试使用该连接。如果不可用，例如连接被关闭或繁忙中，会自动re-prepare，绑定到另一个连接。
   这就导致在高并发的场景，过度使用statement可能导致statement泄漏，statement持续重复prepare和re-prepare的过程，甚至会达到服务器端statement数量上限。
   某些操作使用了PS，例如db.Query(sql, param1, param2), 并在最后自动关闭statement。
4.PS(Prepared Statements)在Tx（transaction）中唯一绑定一个连接，不会re-prepare。
5.Tx和statement不能分离，在DB中创建的statement也不能在Tx中使用，因为他们必定不是使用同一个连接使用Tx必须十分小心
*/
func updateInTransaction(db *sql.DB) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO author(first_name,last_name) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close() // danger!

	for i := 0; i < 10; i++ {
		_, err = stmt.Exec(i, i)
		if err != nil {
			log.Fatal(err)
		}
	}

	//一下两句同时执行：Transaction has already been committed or rolled back
	//且以上insert语句不会反映到数据库中
	//tx.Rollback()
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
	//*sql.Tx一旦释放，连接就回到连接池中，这里stmt在关闭时就无法找到连接。所以必须在Tx commit或rollback之前关闭statement。
	// stmt.Close() runs here!
}
