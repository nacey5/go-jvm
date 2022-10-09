# go-Jvm-6
## 指令/解释器
## 制作一个建议的解释器让其能够运行指令


## 迪米特原则
## 对类进行堆解析
## 对方法进行堆解析
## 对字段进行堆解析
## 堆解析的数据的结构在此进行定义解析

loader字段存放类加载器指针，superClass和 interfaces字段存放类的超类和接口指针，这三个字段将在6.3节介绍。staticSlotCount和
instanceSlotCount字段分别存放类变量和实例变量占据的空间大小，staticVars字段存放静态变量