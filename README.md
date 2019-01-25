# studyNeo4j
study neo4j api

# 导入基础数据

LOAD CSV WITH HEADERS FROM "http://192.168.1.1:8888/group.csv" as row
create (n:Group)
set n = row

LOAD CSV WITH HEADERS FROM "http://192.168.1.1:8888/host.csv" as row
create (a:Host)
set a = row

LOAD CSV WITH HEADERS FROM "http://192.168.1.1:8888/person.csv" as row
create (b:Person)
set b = row

LOAD CSV WITH HEADERS FROM "http://192.168.1.1:8888/service.csv" as row
create (c:Service)
set c = row

# 创建关系

张三
李四
王五
周期
李想
冉伊
材质鹏
吕集

## 人员关系 普通员工  
match (n:Person),(m:Group)
where n.name in ['张三','王五','冉伊'] and m.group="云平台"
create (n)-[:BELONG {role:['普通员工']}]->(m)

## 人员关系 管理员
match (n:Person),(m:Group)
where n.name in ['材质鹏'] and m.group="云平台"
create (n)-[:BELONG {role:['管理员']}]->(m)

## 人员关系 管理员
match (n:Person),(m:Group)
where n.name in ['材质鹏','李想'] and m.group="运维"
create (n)-[:BELONG {role:['管理员']}]->(m)

match (n:Person),(m:Group)
where n.name in ['李四','周期'] and m.group="安防"
create (n)-[:BELONG {role:['员工']}]->(m)

match (n:Person),(m:Group)
where n.name in ['李四','吕集'] and m.group="dba"
create (n)-[:BELONG {role:['员工']}]->(m)

match (n:Person),(m:Group)
where n.name in ['李四','吕集','李想'] and m.group="基础平台"
create (n)-[:BELONG {role:['员工']}]->(m)

match (n:Person),(m:Group)
where n.name in ['冉伊'] and m.group="人事"
create (n)-[:BELONG {role:['员工']}]->(m)

match (n:Person),(m:Group)
where n.name in ['王五'] and m.group="创新"
create (n)-[:BELONG {role:['员工']}]->(m)

match (n:Person),(m:Group)
where n.name in ['周期'] and m.group="研发"
create (n)-[:BELONG {role:['员工']}]->(m)

<!-- create (G:Group{group:"前台",description:"美女"})
create (One1:Person{name:"测试",email:"test@cloudwalk.cn",age:"99",sex:"男"})
create (Two:Person{name:"测试1",email:"test@cloudwalk.cn",age:"99",sex:"男"})
create (One2:Person{name:"测试2",email:"test@cloudwalk.cn",age:"99",sex:"男"})
create (One3:Person{name:"测试3",email:"test@cloudwalk.cn",age:"99",sex:"男"})

create 
    (One1)-[:BELONG {role:['admin']}]->(G),
    (Two)-[:BELONG {role:['admin']}]->(G),
    (One2)-[:BELONG {role:['admin']}]->(G),
    (One3)-[:BELONG {role:['admin']}]->(G) -->

# 设置Service到host

match (a:Host),(b:Service)
where b.ip=a.ip
create (b)-[:Deploy {type:"install"}]->(a)
return a,b

# 设置host到group

<!-- "研发"
"运维"
"安防"
"创新"
"人事"
"基础平台"
"云平台"
"dba" -->

match (h:Host),(g:Group)
where h.hostname in ['hostname1',
    'hostname11',
    'hostname21',
    'hostname31',
    'hostname41',
    'hostname51',
    'hostname61',
    'hostname71',
    'hostname81',
    'hostname91'] and g.group='研发'
create (h)-[:分配]->(g)
return h,g

match (h:Host),(g:Group)
where h.hostname in ['hostname2',
    'hostname12',
    'hostname22',
    'hostname32',
    'hostname42',
    'hostname52',
    'hostname62',
    'hostname72',
    'hostname82',
    'hostname92'] and g.group='运维'
create (h)-[:分配]->(g)
return h,g

match (h:Host),(g:Group)
where h.hostname in ['hostname3',
    'hostname13',
    'hostname23',
    'hostname33',
    'hostname43',
    'hostname53',
    'hostname63',
    'hostname73',
    'hostname83',
    'hostname93'] and g.group='安防'
create (h)-[:分配]->(g)
return h,g

match (h:Host),(g:Group)
where h.hostname in ['hostname4',
    'hostname14',
    'hostname24',
    'hostname34',
    'hostname44',
    'hostname54',
    'hostname64',
    'hostname74',
    'hostname84',
    'hostname94'] and g.group='创新'
create (h)-[:分配]->(g)
return h,g

match (h:Host),(g:Group)
where h.hostname in ['hostname5',
    'hostname15',
    'hostname25',
    'hostname35',
    'hostname45',
    'hostname55',
    'hostname65',
    'hostname75',
    'hostname85',
    'hostname95'] and g.group='人事'
create (h)-[:分配]->(g)
return h,g

match (h:Host),(g:Group)
where h.hostname in ['hostname6',
    'hostname16',
    'hostname26',
    'hostname36',
    'hostname46',
    'hostname56',
    'hostname66',
    'hostname76',
    'hostname86',
    'hostname96'] and g.group='基础平台'
create (h)-[:分配]->(g)
return h,g

match (h:Host),(g:Group)
where h.hostname in ['hostname7',
    'hostname17',
    'hostname27',
    'hostname37',
    'hostname47',
    'hostname57',
    'hostname67',
    'hostname77',
    'hostname87',
    'hostname97'] and g.group='云平台'
create (h)-[:分配]->(g)
return h,g

match (h:Host),(g:Group)
where h.hostname in ['hostname8',
    'hostname18',
    'hostname28',
    'hostname38',
    'hostname48',
    'hostname58',
    'hostname68',
    'hostname78',
    'hostname88',
    'hostname98'] and g.group='dba'
create (h)-[:分配]->(g)
return h,g


# 修改关系 添加属性
match (n)-[r:Deploy]->(a)
set r.desc='部署'

# 批量关系 csv

<!-- LOAD CSV WITH HEADERS FROM "http://data.neo4j.com/northwind/order-details.csv" AS row
MATCH (p:Product), (o:Order)
WHERE p.productID = row.productID AND o.orderID = row.orderID
CREATE (o)-[details:ORDERS]->(p)
SET details = row,
  details.quantity = toInteger(row.quantity) -->

<!-- orderID	productID	unitPrice	quantity	discount
10248	11	14	12	0
10248	42	9.8	10	0
10248	72	34.8	5	0
10249	14	18.6	9	0
10249	51	42.4	40	0
10250	41	7.7	10	0
10250	51	42.4	35	0.15
10250	65	16.8	15	0.15
10251	22	16.8	6	0.05
10251	57	15.6	15	0.05
10251	65	16.8	20	0
10252	20	64.8	40	0.05
10252	33	2	25	0.05
10252	60	27.2	40	0
10253	31	10	20	0
10253	39	14.4	42	0
10253	49	16	40	0
10254	24	3.6	15	0.15 -->

# 权限表

create (n:Privileges{name:'管理员',desc:'admin'})
create (b:Privileges{name:'员工',desc:'worker'})
create (c:Privileges{name:'游客',desc:'guest'})

match (a:Person),(p:Privileges)
where a.name in ['张三'] and p.name='管理员'
create (a)-[:Privileges]->(p)
return a,p

match (a:Person),(p:Privileges)
where a.name in ['李四','王五','周期','李想','冉伊'] and p.name='员工'
create (a)-[:Privileges]->(p)
return a,p

match (a:Person),(p:Privileges)
where a.name in ['吕集','材质鹏'] and p.name='游客'
create (a)-[:Privileges]->(p)
return a,p

# 查询

match (c)-[:Privileges]->(a)-[:BELONG]->(m)<-[:分配]-(n) return a,m,n,c