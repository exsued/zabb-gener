package zabbix
import (
	"log"
)
var (
	zabbixRef = "http://zabbix2023.deozal"
)
func main() {
	//Авторизация
	session, err := NewSession("http://zabbix/api_jsonrpc.php", "Admin", "zabbix")
	if err != nil {
			log.Fatal(err)
	}
	//Получение данных коммутаторов с GlobalMapDump
	//Сверка данных коммутаторов с данными хостов

	//В случае несоотвествия - добавление недостающих хостов в Zabbix
}
