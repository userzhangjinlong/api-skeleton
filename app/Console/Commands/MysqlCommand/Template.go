package MysqlCommand

import (
	"api-skeleton/app/Util"
	"fmt"
	"os"
	"text/template"
)

const databaseTpl = `type {{.TableName | index}} struct {
 {{range.Column}} {{$length := len.Comment}} {{if gt $length 0}} 
//{{.comment}} {{else}} // {{.name}} {{ end }}
	{{$typeLen := len.Type}} {{ if gt $typeLen 0}}{{.name | index}}
	{{.Type}} {{.Tag}} {{else}} {{.Name}} {{end}}
{{end}}

func (model {{.TableName | index}}) TableName() string {
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
		tag := fmt.Sprintf("`"+"json:"+"\"%s\""+"`", column.ColumnName)
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
	tpl := template.Must(template.New("").Funcs(template.FuncMap{
		"index": Util.UnderscoreToUpperCamelCase,
	}).Parse(db.databaseTpl))
	//tpl, err := template.New("").Funcs(template.FuncMap{
	//	"ToCamelCase": Util.UnderscoreToUpperCamelCase,
	//}).ParseFiles("User.go")
	//fmt.Println("err:", err)
	tplDB := StructTemplateDB{
		TableName: tableName,
		Column:    tplColumns,
	}

	err := tpl.Execute(os.Stdout, tplDB)
	//err = tpl.ExecuteTemplate(os.Stdout, "User.go", tplDB)
	if err != nil {
		return err
	}

	return nil
}
