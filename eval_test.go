package goevel

import (
	"reflect"
	"testing"
)

func TestEval(t *testing.T) {
	type args struct {
		defineCode string
		code       string
		imports    []string
	}
	tests := []struct {
		name    string
		args    args
		wantRe  []byte
		wantErr bool
	}{
		{name: "1", args: args{defineCode: "", code: "fmt.Print(\"Hello World!\"", imports: []string{"fmt"}}, wantRe: []byte("Hello World!"), wantErr: true},	// hello world err case
		{name: "2", args: args{defineCode: "", code: "fmt.Print(\"Hello World!\")", imports: []string{"fmt"}}, wantRe: []byte("Hello World!")},	// hello world case
		{name: "3", args: args{defineCode: `const (
	SEX_MAN = true
	SEX_WOMEN = false
)

type Person struct {
	Name string ` + "`" + `json:"name"` + "`" + `
	Age  uint32 ` + "`" + `json:"age,omitempty"` + "`" + `
	Sex  bool   ` + "`" + `json:"sex,omitempty"` + "`" + `
}

var thePerson2 = Person{"李梅", 18, SEX_WOMEN}`, code: `var thePerson = Person{"张三", 18, SEX_MAN}
if re, err := json.Marshal(thePerson); nil == err {
	fmt.Print(string(re))
}
`, imports: []string{"fmt", "encoding/json"}}, wantRe: []byte("{\"name\":\"张三\",\"age\":18,\"sex\":true}")},	// golang json package case
{name: "4", args: args{defineCode: `type Product struct {
	gorm.Model
	Code string
	Price uint
}

func GormDemo() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// 创建
	db.Create(&Product{Code: "L1212", Price: 1000})

	// 读取
	var product Product
	db.First(&product, 1) // 查询id为1的product
	fmt.Printf("get id:1 row : %+v", product)
	db.First(&product, "code = ?", "L1212") // 查询code为l1212的product

	// 更新 - 更新product的price为2000
	db.Model(&product).Update("Price", 2000)

	// 删除 - 删除product
	db.Delete(&product)
}`, code: `GormDemo()`, imports: []string{"fmt", "github.com/jinzhu/gorm", "_ github.com/jinzhu/gorm/dialects/sqlite"}},
wantRe: []byte("get id:1 row : {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:<nil>} Code: Price:0}")}, // third part pcakge case
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRe, err := Eval(tt.args.defineCode, tt.args.code, tt.args.imports...)
			if err != nil {
				if !tt.wantErr {
					t.Errorf("Eval() error = %v, wantErr %v", err, tt.wantErr)
				}
				return
			}
			if !reflect.DeepEqual(gotRe, tt.wantRe) {
				t.Errorf("Eval() gotRe = %v, want %v", gotRe, tt.wantRe)
			}
		})
	}
}
