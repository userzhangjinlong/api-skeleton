package MysqlCommand

import (
	"api-skeleton/app/Util"
	"bytes"
	"fmt"
	"os"
	"text/template"
)

const databaseTpl = `package {{.DatabaseName}}

import "api-skeleton/app/Model"

type {{.TableName | index}} struct {
 {{range.Column}} {{$length := len .Comment}} {{if gt $length 0}}
	//{{.Comment}}  {{else}} // {{.Name}} {{ end }}
	{{$typeLen := len .Type}} {{ if gt $typeLen 0}}{{.Name | index}} {{.Type}} {{.Tag}} {{else}} {{.Name}} {{end}}
{{end}}
	//继承父类model
	Model.Model
}

func (model *{{.TableName | index}}) TableName() string {
	return "{{.TableName}}"
}`

type DatabaseTemplate struct {
	databaseTpl string
}

type DatabaseColumn struct {
	Name    string
	Type    string
	Tag     string
	Comment string
}

type StructTemplateDB struct {
	TableName    string
	DatabaseName string
	Column       []*DatabaseColumn
}

//NewDatabaseTemplate 实例化表名称
func NewDatabaseTemplate() *DatabaseTemplate {
	return &DatabaseTemplate{
		databaseTpl: databaseTpl,
	}
}

//AssemblyColumns 格式化获取到的字段输出模板内容
func (db *DatabaseTemplate) AssemblyColumns(tbColumns []*TableColumn) []*DatabaseColumn {
	tplColumns := make([]*DatabaseColumn, 0, len(tbColumns))
	fmt.Println("表中字段")
	for _, column := range tbColumns {
		fmt.Println(Util.InArray(column.ColumnName, []string{"createTime", "updateTime"}))
		if !Util.InArray(column.ColumnName, []string{"createtime", "updatetime"}) {
			tag := fmt.Sprintf("`"+"gorm:"+"\"%s\""+" "+"json:"+"\"%s\""+"`", column.ColumnName, column.ColumnName)
			tplColumns = append(tplColumns, &DatabaseColumn{
				Name:    column.ColumnName,
				Type:    DBTypeToStructType[column.DataType],
				Tag:     tag,
				Comment: column.ColumnComment,
			})
		}

	}

	return tplColumns
}

//Generate 处理渲染模板内容
func (db *DatabaseTemplate) Generate(dataBaseName string, tableName string, tplColumns []*DatabaseColumn) error {
	tpl := template.Must(template.New("MysqlCommand").Funcs(template.FuncMap{
		"index": Util.UnderscoreToUpperCamelCase,
	}).Parse(db.databaseTpl))

	tplDB := StructTemplateDB{
		TableName:    tableName,
		DatabaseName: Util.Capitalize(dataBaseName),
		Column:       tplColumns,
	}

	buf := new(bytes.Buffer)
	err := tpl.Execute(buf, tplDB)

	if err != nil {
		return err
	}
	putStringInFile(dataBaseName, Util.UnderscoreToUpperCamelCase(tableName), buf.String())

	return nil
}

//putInFile 将写入buff池内的内容写入指定文件
func putStringInFile(dataBaseName string, fileName string, bufString string) error {
	Util.CreateIfNotExistDir("app/Model/" + Util.Capitalize(dataBaseName))
	file, err := os.OpenFile("app/Model/"+Util.Capitalize(dataBaseName)+"/"+fileName+".go", os.O_CREATE|os.O_WRONLY, 0664)

	if err != nil {
		fmt.Println("open file err", err)
		return err
	}

	fmt.Fprintf(file, bufString)
	file.Close()

	return nil
}
