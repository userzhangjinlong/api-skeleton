package MysqlCommand

import (
	"api-skeleton/app/Util"
	"bytes"
	"fmt"
	"os"
	"text/template"
)

const databaseTpl = `package Model

type {{.TableName | index}} struct {
 {{range.Column}} {{$length := len .Comment}} {{if gt $length 0}}
	//{{.Comment}}  {{else}} // {{.Name}} {{ end }}
	{{$typeLen := len .Type}} {{ if gt $typeLen 0}}{{.Name | index}} {{.Type}} {{.Tag}} {{else}} {{.Name}} {{end}}
{{end}}
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
	TableName string
	Column    []*DatabaseColumn
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

	for _, column := range tbColumns {
		tag := fmt.Sprintf("`"+"gorm:"+"\"%s\""+" "+"json:"+"\"%s\""+"`", column.ColumnName, column.ColumnName)
		tplColumns = append(tplColumns, &DatabaseColumn{
			Name:    column.ColumnName,
			Type:    DBTypeToStructType[column.DataType],
			Tag:     tag,
			Comment: column.ColumnComment,
		})
	}

	return tplColumns
}

//Generate 处理渲染模板内容
func (db *DatabaseTemplate) Generate(tableName string, tplColumns []*DatabaseColumn) error {
	tpl := template.Must(template.New("MysqlCommand").Funcs(template.FuncMap{
		"index": Util.UnderscoreToUpperCamelCase,
	}).Parse(db.databaseTpl))

	tplDB := StructTemplateDB{
		TableName: tableName,
		Column:    tplColumns,
	}

	buf := new(bytes.Buffer)
	err := tpl.Execute(buf, tplDB)

	if err != nil {
		return err
	}
	putStringInFile(Util.UnderscoreToUpperCamelCase(tableName), buf.String())

	return nil
}

//putInFile 将写入buff池内的内容写入指定文件
func putStringInFile(fileName string, bufString string) error {
	file, err := os.OpenFile("app/Model/"+fileName+".go", os.O_CREATE|os.O_WRONLY, 0664)

	if err != nil {
		fmt.Println("open file err", err)
		return err
	}

	fmt.Fprintf(file, bufString)
	file.Close()

	return nil
}
