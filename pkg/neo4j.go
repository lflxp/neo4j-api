package pkg

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/lflxp/neo4j-api/models"

	"github.com/neo4j/neo4j-go-driver/neo4j"
)

var (
	driver   neo4j.Driver
	session  neo4j.Session
	result   neo4j.Result
	username string
	password string
	uri      string
	err      error
)

func init() {
	username = beego.AppConfig.String("neo4j::user")
	password = beego.AppConfig.String("neo4j::pwd")
	uri = beego.AppConfig.String("neo4j::uri")
}

func initDriver() error {
	driver, err = neo4j.NewDriver(uri, neo4j.BasicAuth(username, password, ""))
	// driver, err = neo4j.NewDriver(uri, neo4j.BasicAuth("", "", ""))
	if err != nil {
		return err
	}
	return err
}

func initSession(mode neo4j.AccessMode) error {
	err = initDriver()
	if err != nil {
		return err
	}
	session, err = driver.Session(mode)
	if err != nil {
		return err
	}
	return nil
}

func ReadTran(cql string, arg map[string]interface{}) (*models.Neo4j, error) {
	var rs *models.Neo4j
	err = initSession(neo4j.AccessModeWrite)
	if err != nil {
		return rs, err
	}
	defer driver.Close()
	defer session.Close()
	tmp_rs, err := session.ReadTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		rs_tmp := &models.Neo4j{}
		result, err = transaction.Run(cql, arg)
		if err != nil {
			return rs, err
		}
		// node剔重
		node_map := map[string]interface{}{}
		node := []interface{}{}
		relation := []interface{}{}
		str := []interface{}{}
		ints := []interface{}{}

		for {
			if result.Next() {
				// fmt.Println(result.Record().Keys())
				// fmt.Println(result.Record().Keys(), result.Record().Values())
				// person := result.Record().GetByIndex(0).(neo4j.Node)
				// // fmt.Println(person.Id(), person.Labels(), person.Props())
				// group := result.Record().GetByIndex(1).(neo4j.Node)
				// // fmt.Println(group.Id(), group.Labels(), group.Props())
				// host := result.Record().GetByIndex(2).(neo4j.Node)
				// // fmt.Println(host.Id(), host.Labels(), host.Props())
				// relation := result.Record().GetByIndex(3).(neo4j.Relationship)
				// relation1 := result.Record().GetByIndex(4).(neo4j.Relationship)
				// fmt.Println(person.Props(), group.Props(), host.Props(), relation.Props(), relation.Type(), relation1.Props(), relation1.Type())

				// data, err := json.Marshal(host.Props())
				// if err != nil {
				// 	panic(err)
				// }
				// fmt.Println(string(data))

				for n, x := range result.Record().Keys() {
					var tmp_rs map[string]interface{}
					rs := result.Record().GetByIndex(n)
					switch v := rs.(type) {
					case neo4j.Node:
						// fmt.Println("node")
						if _, ok := node_map[fmt.Sprintf("%d", rs.(neo4j.Node).Id())]; !ok {
							fmt.Println(fmt.Sprintf("-%d-", rs.(neo4j.Node).Id()), node_map)
							tmp_rs = map[string]interface{}{}
							tmp_rs["group"] = x
							tmp_rs["props"] = rs.(neo4j.Node).Props()
							tmp_rs["id"] = rs.(neo4j.Node).Id()
							// tmp_rs["name"] = rs.(neo4j.Node).Id()
							tmp_rs["labels"] = rs.(neo4j.Node).Labels()
							tmp_rs["type"] = "node"
							// ss, _ := json.Marshal(tmp_rs)
							// fmt.Println(string(ss))
							// node = append(node, tmp_rs)
							node_map[fmt.Sprintf("%d", rs.(neo4j.Node).Id())] = tmp_rs
						}

					case neo4j.Relationship:
						tmp_rs = map[string]interface{}{}
						// fmt.Println("relationship")
						tmp_rs["name"] = x
						// tmp_rs["id"] = rs.(neo4j.Relationship).Id()
						tmp_rs["props"] = rs.(neo4j.Relationship).Props()
						tmp_rs["relation"] = rs.(neo4j.Relationship).Type()
						tmp_rs["source"] = rs.(neo4j.Relationship).StartId()
						tmp_rs["target"] = rs.(neo4j.Relationship).EndId()
						tmp_rs["value"] = 1
						// ss, _ := json.Marshal(tmp_rs)
						// fmt.Println(string(ss))
						relation = append(relation, tmp_rs)
					case string:
						tmp_rs = map[string]interface{}{}
						// fmt.Println("string")
						tmp_rs["group"] = x
						tmp_rs["value"] = rs
						tmp_rs["type"] = "string"
						// ss, _ := json.Marshal(tmp_rs)
						// fmt.Println(string(ss))
						str = append(str, tmp_rs)
					case int64:
						// 	tmp_rs["length"] = rs
						tmp_rs = map[string]interface{}{}
						tmp_rs["group"] = x
						tmp_rs["value"] = rs
						tmp_rs["type"] = "int64"
						ints = append(ints, tmp_rs)
					default:
						fmt.Println("unknow type", v)
					}
				}
				// fmt.Println(result.Record().Keys(), result.Record().Values())
				// fmt.Println(result.Record().GetByIndex(0))
				// return result.Record().GetByIndex(0), nil
				// return "ok", nil

			} else {
				break
			}
		}

		// 生成node
		for k, v := range node_map {
			fmt.Println("id " + k)
			node = append(node, v)
		}
		// 文本结果
		rs_tmp.Str = str
		rs_tmp.Nodes = node
		rs_tmp.Ints = ints

		node_string, _ := json.Marshal(node)

		fmt.Println(string(node_string))
		relation_string, _ := json.Marshal(relation)
		fmt.Println(string(relation_string))

		// 去除没有在links的node 不行 需由前端自行剔除
		// link获取node的index
		countNode := map[int64]int{}
		for n1, x := range node {
			countNode[x.(map[string]interface{})["id"].(int64)] = n1
		}
		// node_string, _ := json.Marshal(node)

		// fmt.Println(string(node_string))
		// fmt.Println("##################################################")
		for n2, y := range relation {
			y.(map[string]interface{})["source"] = countNode[y.(map[string]interface{})["source"].(int64)]
			y.(map[string]interface{})["target"] = countNode[y.(map[string]interface{})["target"].(int64)]
			relation[n2] = y
		}
		// relation_string, _ := json.Marshal(relation)
		// fmt.Println(string(relation_string))

		rs_tmp.Links = relation
		// return nil, result.Err()
		return rs_tmp, result.Err()
	})

	if err != nil {
		return rs, result.Err()
	}
	return tmp_rs.(*models.Neo4j), nil
}
