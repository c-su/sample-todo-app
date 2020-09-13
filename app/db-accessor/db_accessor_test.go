package dbaccessor_test

import (
	"database/sql/driver"
	"regexp"
	"testing"
	"time"
	dbaccessor "work/app/db-accessor"

	"github.com/jinzhu/gorm"

	"github.com/DATA-DOG/go-sqlmock"
)

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func createMockDb() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	gdb, err := gorm.Open("sqlite3", db)
	if err != nil {
		return nil, nil, err
	}
	gdb.LogMode(true)
	return gdb, mock, nil
}

func TestDbInsert(t *testing.T) {
	db, mock, err := createMockDb()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	text := "テストタスク"
	status := "未実施"

	mockedResult := sqlmock.NewResult(1, 1)

	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "todos"`)).WithArgs(AnyTime{}, AnyTime{}, nil, text, status).WillReturnResult(mockedResult)
	mock.ExpectCommit()

	dbaccessor.DbInsert(db, text, status)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDbUpdate(t *testing.T) {
	db, mock, err := createMockDb()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	text := "更新タスク"
	status := "実施中"

	mockedRows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "text", "status"}).AddRow(1, time.Now(), time.Now(), nil, "テストタスク", "未実施")
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT`)).WillReturnRows(mockedRows)

	mock.ExpectBegin()
	mockedResult := sqlmock.NewResult(1, 1)
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "todos"`)).WillReturnResult(mockedResult)
	mock.ExpectCommit()

	dbaccessor.DbUpdate(db, 1, text, status)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDbGetAll(t *testing.T) {
	db, mock, err := createMockDb()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	mockedRows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "text", "status"})
	mockedRows.AddRow(1, time.Now(), time.Now(), nil, "テストタスク1", "未実施")
	mockedRows.AddRow(2, time.Now(), time.Now(), nil, "テストタスク2", "未実施")
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "todos" WHERE "todos"."deleted_at" IS NULL ORDER BY created_at desc`)).WillReturnRows(mockedRows)

	dbaccessor.DbGetAll(db)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDbGetOne(t *testing.T) {
	db, mock, err := createMockDb()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	mockedRows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "text", "status"})
	mockedRows.AddRow(1, time.Now(), time.Now(), nil, "テストタスク1", "未実施")
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM "todos"  WHERE "todos"."deleted_at" IS NULL AND (("todos"."id" = 1)) ORDER BY "todos"."id" ASC LIMIT 1`)).WillReturnRows(mockedRows)

	dbaccessor.DbGetOne(db, 1)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDbDelete(t *testing.T) {
	db, mock, err := createMockDb()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	mockedRows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "text", "status"}).AddRow(1, time.Now(), time.Now(), nil, "テストタスク", "未実施")
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT`)).WillReturnRows(mockedRows)

	mock.ExpectBegin()
	mockedResult := sqlmock.NewResult(1, 1)
	mock.ExpectExec(regexp.QuoteMeta(`UPDATE "todos" SET "deleted_at"=? WHERE "todos"."deleted_at" IS NULL AND "todos"."id" = ?`)).WithArgs(AnyTime{}, 1).WillReturnResult(mockedResult)
	mock.ExpectCommit()

	dbaccessor.DbDelete(db, 1)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
