#### 加载数据
* 从hive表加载数据：
    * 添加hive-site.xml,添加配置项hive.metastore.uris指向hive-matestore服务
    * ``` spark = SparkSession.builder.appName("reg_etl").enableHiveSupport().getOrCreate() # 创建spark与hive的会话```
    * 直接通过spark.sql()执行sql从hive中拉取数据。
* 从phoenix中拉取数据：
    * 将phoenix的client-jar包放入spark.driver.extraClassPath的目录中，如果没有配置加上
 ``` 
df = spark.read.format("org.apache.phoenix.spark")
.option("table", "").option("zkUrl", "").load() 
# table为phoenix的表名，zkUrl为hbase的zk集群地址 
```
 ```
 note: 不要对df进行直接的操作，这将导致spark拉取phoenix全表的数据，
 对df进行filter和select选择部分数据进行工作
 ```
* 加载数据文件
    * 使用spark的rdd进行数据加载 
```
conf = SparkConf().setAppName(appName).setMaster(master)  
sc = SparkContext(conf=conf)
distFile = sc.textFile("data.txt") 
# you can use textFile("/my/directory"), textFile("/my/directory/*.txt"), and textFile("/my/directory/*.gz").
# his method takes an URI for the file (either a local path on the machine, or a hdfs://, s3a://, etc URI) and reads it as a collection of lines. 
```

* 使用spark的datefream加载数据
```
spark dateframe支持从多种数据源加载数据
scv,orc,parquent,json,text,jdbc,hive table.
使用spark.read.load("path", format='')加载数据，或使用对应的方法，如spark.read.json()
## spark.read.format().option().load(),可以自定义拉取任意数据，如phoenix
```

#### 对数据进行处理
* 筛选
```
t1 = df.filter("a > 1 and c < 9") # 直接写条件a,c为数据列

t2 = df.filter( (df['b']<5) & (df['c']<8)) #  可以使用&或|对两个bool列进行逻辑运算，但必须要用圆括号括起，限定运算顺序。

## df['xxx']与df.xxx效果相同，在pyspark中没有区别，推荐使用df['xxx']. where与filter相同

```
    
* 赋值，加列，删除列
```
t1 = df.withColumn("c", df['a']+1) # 将df的a列的值+1并赋值给c列，如果c列不存在则添加一个新的c列
t2 = df.drop("c")删除c列
```
* 选取列
```
select(*cols):cols为列名或列对象。
赋值和删除操作，每次只能改加减一列数据，如果想要批量地改变，尤其是调整列顺序的时候，就非常有用了。在ETL中，当需要计算的列很多时，通常就是逐个计算出不同的列对象，最后用select把它们排好顺序。
选取单一列
a1 = df['a']+1).alias("a1")  # 新增一个列对象，取名为a1

t = d1.select("a", a1) #生成一个新的df
```
* 常数列
```
lit(value):value数必须是必须为pyspark.sql.types支持的类型，比如int,double,string,datetime等
t = df.withColumn("new", lit(10)) # 添加一个new列，值为常数10
```
* 条件分支
```
when(cond,value):符合cond就取value，value可以是常数也可以是一个列对象，连续可以接when构成多分支
otherwise(value):接在when后使用，所有不满足when的行都会取value,若不接这一项，则取Null。
t = df.withColumn("cc", when(df[a]==1,1).when(df[a]==2,2).otherwise(3)) # 列cc的值取决于条件的不同
```
* 自定义规则
```
from pyspark.sql.functions import udf
from pyspark.sql.types import StringType
def f(a):
    if a:
      rerutn a
    return "XX"

udfff = udf(f, StringType())

t = df.withColumn("new", udfff(df['a'])) #对df的a列使用f函数将值添加到new列

```

#### 写
* 参考数据加载，将read改为write，用法基本一致

#### 提交spark任务
* 代码打包
```
root
    -- src
    -- lib
main.py
requirentments.txt

pip install -r requirentments.txt -t ./lib
cd ./lib
zip -r ../lib.zip .
cd ../
zip -r src.zip ./src 

打包结束你将拥有一个入口main.py 代码包src.zip 依赖包lib.zip
```
* 提交参数
```
spark-submit
--class: The entry point for your application (e.g. org.apache.spark.examples.SparkPi) # 提交jar文件指定入口类
--master yarn,loacl，spark://HOST:PORT     # 运行方式yarn，本地，spark集群
--deploy-mode client,cluster  # driver的部署位置，集群或客户端
--executor-memory 20G  # 运行内存
--executor-cores # 使用cpu核数
--num-executors 50  #容器数量
--queue # 任务使用的yarn队列
--files # 
```
* 
* 