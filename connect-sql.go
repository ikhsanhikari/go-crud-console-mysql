package main

import "fmt"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"

type student struct {
    id    string
    name  string
    age   int
    grade int
}

func connect() (*sql.DB, error) {
    db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/belajar_golang")
    if err != nil {
        return nil, err
    }

    return db, nil
}

func users(w http.ResponseWriter, r * http.Request) {
    w.Header().Set("Content-Type", "application/json")

    if r.Method == "GET" {
        var result, err = json.Marshal(data)

        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.Write(result)
        return
    }

    http.Error(w, "", http.StatusBadRequest)
}


func sqlQuery() {
    db, err := connect()
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    defer db.Close()

    //var age = 27
    //rows, err := db.Query("select id, name, grade from tb_student where age = ?", age)
	
	rows, err := db.Query("select * from tb_student ")
	
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    defer rows.Close()

    var result []student

    for rows.Next() {
        var each = student{}
        var err = rows.Scan(&each.id, &each.name, &each.grade, &each.age)

        if err != nil {
            fmt.Println(err.Error())
            return
        }

        result = append(result, each)
    }

    if err = rows.Err(); err != nil {
        fmt.Println(err.Error())
        return
    }

    for _, each := range result {
        fmt.Println(each.id, each.name,each.age,each.grade)
    }
}


func sqlQueryRow() {
    var db, err = connect()
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    defer db.Close()

    var result = student{}
    var id = "E001"
    err = db.
        QueryRow("select name, grade from tb_student where id = ?", id).
        Scan(&result.name, &result.grade)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    fmt.Printf("name: %s\ngrade: %d\n", result.name, result.grade)
}

func sqlPrepare() {
    db, err := connect()
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    defer db.Close()

    stmt, err := db.Prepare("select name, grade from tb_student where id = ?")
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    var result1 = student{}
    stmt.QueryRow("E001").Scan(&result1.name, &result1.grade)
    fmt.Printf("name: %s\ngrade: %d\n", result1.name, result1.grade)

    var result2 = student{}
    stmt.QueryRow("W001").Scan(&result2.name, &result2.grade)
    fmt.Printf("name: %s\ngrade: %d\n", result2.name, result2.grade)

    var result3 = student{}
    stmt.QueryRow("B001").Scan(&result3.name, &result3.grade)
    fmt.Printf("name: %s\ngrade: %d\n", result3.name, result3.grade)
}

func insert(){
	db, err := connect()
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    defer db.Close()
	var kode string
	var nama string
	var umur int
	var peringkat int
	
	fmt.Printf("Kode : ")
	fmt.Scan(&kode)
	
	fmt.Printf("Nama : ")
	fmt.Scan(&nama)
	
	fmt.Printf("Umur : ")
	fmt.Scan(&umur)
	
	fmt.Printf("Peringkat : ")
	fmt.Scan(&peringkat)
	
	_, err = db.Exec("insert into tb_student values (?, ?, ?, ?)", kode, nama, umur, peringkat)
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    fmt.Println("insert success!")
	
}
func update(){
	db, err := connect()
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    defer db.Close()
	var kode string
	var nama string
	var umur int
	var peringkat int
	
	fmt.Printf("Kode : ")
	fmt.Scan(&kode)
	
	fmt.Printf("Nama : ")
	fmt.Scan(&nama)
	
	fmt.Printf("Umur : ")
	fmt.Scan(&umur)
	
	fmt.Printf("Peringkat : ")
	fmt.Scan(&peringkat)
	
	_, err = db.Exec("update tb_student set age = ?, grade = ?, name = ? where id = ?", umur,peringkat,nama, kode)
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    fmt.Println("update success!")
	
}
func delete(){
	db, err := connect()
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    defer db.Close()
	var kode string
	
	fmt.Printf("Kode : ")
	fmt.Scan(&kode)
	
	
	_, err = db.Exec("delete from tb_student where id = ?", kode)
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    fmt.Println("delete success!")
	
}
func sqlExec() {
    db, err := connect()
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    defer db.Close()

    _, err = db.Exec("insert into tb_student values (?, ?, ?, ?)", "G001", "Galahad", 29, 2)
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    fmt.Println("insert success!")

    _, err = db.Exec("update tb_student set age = ? where id = ?", 28, "G001")
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    fmt.Println("update success!")

    _, err = db.Exec("delete from tb_student where id = ?", "G001")
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    fmt.Println("delete success!")
}
func menu(){
	fmt.Println("======================= Action ! =========================")
	fmt.Println("1. Insert ")
	fmt.Println("2. Update ")
	fmt.Println("3. Delete ")
	fmt.Println("4. Read ")
	fmt.Print("Choose : ")
	var pilih string 
	fmt.Scan(&pilih)
	if pilih == "1" {
		fmt.Println("================== Start Insert Data ================= ")
		insert()
		menu()
	}else if pilih == "2"{
		fmt.Println("================== Start Update Data ================= ")
		update()
		menu()
	}else if pilih == "3"{
		fmt.Println("================== Start Delete Data ================= ")
		delete()
		menu()
	}else if pilih == "4"{
		fmt.Println("================== Start Read Data ================= ")
		sqlQuery()
		menu()
	}
}

func main() {
    menu()
	
	//sqlQuery()
	//sqlQueryRow()
	//sqlPrepare()
	//sqlExec()
	
	//insert()
}