package models

type Neo4jPost struct {
	Cql string `json:"cql"` // neo4j cql 查询语句
}

type Neo4j struct {
	Nodes []interface{} `json:"nodes"`
	Links []interface{} `json:"links"`
	Str   []interface{} `json:"str"`
}
