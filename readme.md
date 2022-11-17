学习笔记

SQL.DB笔记

sql.open 

开始连接数据库，
第一个参数为所要连接的数据库类型。
第二个参数是dsn,可以理解为连接数据库的参数信息

conn.SetConnMaxLifetime

为设置连接的空闲时间，为什么要设置这个参数以及该设置多少是有讲究的。

conn.SetMaxOpenConns(10)
和conn.SetMaxOpenConns(10)

这两个方法会初始化一个连接池，保证池子里会有足够的连接，这个值要设置多少需要根据应用的并发情况。