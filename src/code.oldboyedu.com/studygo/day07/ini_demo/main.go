package main


import(
	"fmt"
	"reflect"
	"errors"
	"strings"
	"strconv"

)

//ini配置文件解析器

//Mysqlconfig mysql配置结构体
type MysqlConfig struct{
	Address string `ini:"address"`
	Port int `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`

}

type RedisConfig struct{
	Host string `ini:"host"`
	Port int `ini:"port"`	
	Password string `ini:"password"`
	Database string `ini:"database"`
}

type Config struct{
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:redis`	
}

func loadIni(fileName string,data interface{})(err error){
	//参数的校验
	t:=reflect.TypeOf(data)
	fmt.Println(t,t.Kind())
	if t.Kind()!=reflect.Ptr{
		err=errors.New("data param should be pointer")
		return 
	}

	if t.Elem().Kind()!=reflect.Struct{
		err=errors.New("data param should be a struct")
		return
	}

	//1.读文件得到字节类型数据
	b,err:=ioutil.ReadFile(fileName)
	if err!=nil{
		return
	}
	
	lineSlice:=strings.Split(string(b),"\r\n")
	var structName string
	//2.一行一行得读数据
	for idx,line:=range lineSlice{
		//去掉空格
		line=strings.TrimSpace(line)
		//如果是空行就跳过
		if len(line)==0{
			continue
		}
		//如果是以;和#开头
		//2.1如果是注释就跳过
		if strings.HasPrefix(line,";") || strings.HasPrefix(line,"#"){
			continue
		}
		//2.2如果是[开头就表示是节(section)
		if strings.HasPrefix(line,"["){
			if line[0]!='['||line[len(line)-1]!=']'{
				err=fmt.Errorf("line:%d syntax error",idx+1)
				return
			}
			//如果[]里面的内容是空
			sectionName:=strings.TrimSpace(line[1:len(line)-1])
			if len(sectionName)==0{
				err=fmt.Errorf("line:%d syntax error",idx+1)
				return				
			}
			//根据字符串sectionName去data里面根据反射找到对应的结构体			
			for i:=0;i<t.Elem().NumField();i++{
				field:=t.Elem().Field(i)
				if sectionName==field.Tag.Get("ini"){
					//说明找到了对应的嵌套结构体，把字段名记下来
					structName=field.Name				
				}
			}
		}else{
			//2.3如果不是[开头就是=分割的键值对
			//1、以等号分割这一行，等号左边是key，等号右边是value
			if strings.Index(line,"=")==-1 ||strings.HasPrefix(line,"="){
				err=fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
			index:=strings.Index(line,"=")
			key:=strings.TrimSpace(line[:index])
			value:=strings.TrimSpace(line[index+1:])

			//2、根据structName去data里面把对应的嵌套结构体给取出来
			v:=reflect.ValueOf(data)
			sValue:=v.Elem().FieldByName(structName)
			sType:=sValue.Type()

			if sType.Kind()!=reflect.Struct{
				err=fmt.Errorf("data中的%s字段应该是一个结构体",structName)
				return
			}
			var fieldName string
			var fieldType reflect.StructField
			//3、遍历嵌套结构体的每一个字段，判断tag是不是等于key
			for i:=0;i<sValue.NumField();i++{
				field:=sType.Field(i)
				if field.Tag.Get("ini")==key{
					//找到对应的字段
					fieldName=field.Name
					break
				}
			}

			//4、如果key=tag，给这个字段赋值
			//4.1 根据fieldName去取出这个字段
			if len(fieldName)==0{
				//在结构体中找不到对应的字段
				continue
			}
			fileObj:=sValue.FieldByName(fieldName)
			//4.2 对其赋值
			//fmt.Println(fieldName,fileType.Type.Kind())
			switch fileType.Type.Kind(){
			case reflect.String:
				fileObj.SetString(value)
			case reflect.Int,reflect.Int8,reflect.Int16,reflect.Int32,reflect.Int64:
				var valueInt int64
				valueInt,err=strconv.ParseInt(value,10,64)
				if err!=nil{
					fmt.Errorf("line:%d value type error",idx+1)
					return
				}
				fileObj.SetInt(valueInt)
			case reflect.Bool:
				var valueBool bool
				valueBool,err=strconv.ParseBool(value)
				if err!=nil{
					fmt.Errorf("line:%d value type error",idx+1)
					return
				}
				fileObj.SetBool(valueBool)
			case reflect.Float32,reflect.Float64:
				var valueFloat float64
				valueFloat,err=strconv.ParseFloat(value,64)
				if err!=nil{
					fmt.Errorf("line:%d value type error",idx+1)
					return
				}
				fileObj.SetFloat(valueFloat)

			}
		}		
	}
	return
}

func main(){
	var cfg Config
	err:=loadIni("./conf.ini",cfg)
	if err!=nil{
		fmt.Printf("load ini failed,err:%v\n",err)
		return
	}
	fmt.Println(cfg)

}