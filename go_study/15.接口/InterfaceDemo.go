package main

import "fmt"

type Database interface {
	Connect() error
	Query(sql string) ([]string, error)
	Close() error
}

type Mysql struct {
	connection string
}

func (m *Mysql) Connect() error {
	fmt.Printf("连接 MySQL: %s\n", m.connection)
	return nil
}

func (m *Mysql) Query(sql string) ([]string, error) {
	fmt.Printf("查询 SQL: %s\n", sql)
	return []string{"1", "2", "3"}, nil
}

func (m *Mysql) Close() error {
	fmt.Println("关闭 MySQL")
	return nil
}

type Postgresql struct {
	connection string
}

func (p *Postgresql) Connect() error {
	fmt.Printf("连接 PostgreSQL: %s\n", p.connection)
	return nil
}
func (p *Postgresql) Query(sql string) ([]string, error) {
	fmt.Printf("查询 SQL: %s\n", sql)
	return []string{"1", "2", "3"}, nil
}
func (p *Postgresql) Close() error {
	fmt.Println("关闭 PostgreSQL")
	return nil
}

func executeQuery(db Database, sql string) {
	if err := db.Connect(); err != nil {
		fmt.Println("连接数据库失败")
		return
	}

	defer db.Close()

	if result, err := db.Query(sql); err != nil {
		fmt.Println("查询数据库失败")
	} else {
		fmt.Println("查询结果:", result)
	}
}
func main() {

	mysql := &Mysql{connection: "mysql://root:123456@localhost:3306"}
	postgres := &Postgresql{connection: "postgres://localhost:5432"}

	// mock
	executeQuery(mysql, "SELECT * FROM users")
	executeQuery(postgres, "SELECT * FROM orders")
}
