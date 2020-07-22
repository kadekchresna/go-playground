package main

import (
	"database/sql"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func insertRecord(b *testing.B, db *sql.DB) {
	stmn, errPrepare := db.Prepare("INSERT INTO not_hoax.blog (blog_title, blog_body, is_active) VALUES ('First', 'Heeyyy', 1);")
	if errPrepare != nil {
		b.Fatal(errPrepare)
	}
	defer stmn.Close()
	_, errInsert := stmn.Exec()
	if errInsert != nil {
		b.Fatal(errInsert)
	}
}

func BenchmarkMaxOpenConns1(b *testing.B) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/not_hoax?charset=utf8")
	if err != nil {
		b.Fatal(err)
	}
	db.SetMaxOpenConns(1)
	defer db.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			insertRecord(b, db)
		}
	})
}

func BenchmarkMaxOpenConns2(b *testing.B) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/not_hoax?charset=utf8")
	if err != nil {
		b.Fatal(err)
	}
	db.SetMaxOpenConns(2)
	defer db.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			insertRecord(b, db)
		}
	})
}

func BenchmarkMaxOpenConns5(b *testing.B) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/not_hoax?charset=utf8")
	if err != nil {
		b.Fatal(err)
	}
	db.SetMaxOpenConns(5)
	defer db.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			insertRecord(b, db)
		}
	})
}

func BenchmarkMaxOpenConns10(b *testing.B) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/not_hoax?charset=utf8")
	if err != nil {
		b.Fatal(err)
	}
	db.SetMaxOpenConns(10)
	defer db.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			insertRecord(b, db)
		}
	})
}

func BenchmarkMaxOpenConnsUnlimited(b *testing.B) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/not_hoax?charset=utf8")
	if err != nil {
		b.Fatal(err)
	}
	defer db.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			insertRecord(b, db)
		}
	})
}

func BenchmarkMaxIdleConnsNone(b *testing.B) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/not_hoax?charset=utf8")
	if err != nil {
		b.Fatal(err)
	}
	db.SetMaxIdleConns(0)
	defer db.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			insertRecord(b, db)
		}
	})
}

func BenchmarkMaxIdleConns1(b *testing.B) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/not_hoax?charset=utf8")
	if err != nil {
		b.Fatal(err)
	}
	db.SetMaxIdleConns(1)
	defer db.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			insertRecord(b, db)
		}
	})
}

func BenchmarkMaxIdleConns2(b *testing.B) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/not_hoax?charset=utf8")
	if err != nil {
		b.Fatal(err)
	}
	db.SetMaxIdleConns(2)
	defer db.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			insertRecord(b, db)
		}
	})
}

func BenchmarkMaxIdleConns5(b *testing.B) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/not_hoax?charset=utf8")
	if err != nil {
		b.Fatal(err)
	}
	db.SetMaxIdleConns(5)
	defer db.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			insertRecord(b, db)
		}
	})
}

func BenchmarkMaxIdleConns10(b *testing.B) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/not_hoax?charset=utf8")
	if err != nil {
		b.Fatal(err)
	}
	db.SetMaxIdleConns(10)
	defer db.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			insertRecord(b, db)
		}
	})
}

func BenchmarkConnMaxLifetimeUnlimited(b *testing.B) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/not_hoax?charset=utf8")
	if err != nil {
		b.Fatal(err)
	}
	db.SetConnMaxLifetime(0)
	defer db.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			insertRecord(b, db)
		}
	})
}

func BenchmarkConnMaxLifetime1000(b *testing.B) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/not_hoax?charset=utf8")
	if err != nil {
		b.Fatal(err)
	}
	db.SetConnMaxLifetime(1000 * time.Millisecond)
	defer db.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			insertRecord(b, db)
		}
	})
}

func BenchmarkConnMaxLifetime500(b *testing.B) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/not_hoax?charset=utf8")
	if err != nil {
		b.Fatal(err)
	}
	db.SetConnMaxLifetime(500 * time.Millisecond)
	defer db.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			insertRecord(b, db)
		}
	})
}

func BenchmarkConnMaxLifetime200(b *testing.B) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/not_hoax?charset=utf8")
	if err != nil {
		b.Fatal(err)
	}
	db.SetConnMaxLifetime(200 * time.Millisecond)
	defer db.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			insertRecord(b, db)
		}
	})
}

func BenchmarkConnMaxLifetime100(b *testing.B) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/not_hoax?charset=utf8")
	if err != nil {
		b.Fatal(err)
	}
	db.SetConnMaxLifetime(100 * time.Millisecond)
	defer db.Close()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			insertRecord(b, db)
		}
	})
}
