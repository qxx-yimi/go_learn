#### 效果演示

1.注册账号

![image-20230813112317247](https://github.com/qxx-yimi/go_learn/blob/main/images_md/image-20230813112317247.png)

注册之后，昵称默认为创建的时间戳，生日和个人简介是空白

![image-20230813112428475](https://github.com/qxx-yimi/go_learn/blob/main/images_md/image-20230813112428475.png)

2.登录账号

只有登录账号之后才能修改信息

![image-20230813113149888](C:\Users\isxia\AppData\Roaming\Typora\typora-user-images\image-20230813113149888.png)

3.编辑个人信息

需要传输的信息是nickname，birthday，personalProfile。校验时会校验birthday的格式，nickname是否为空，以及nickname和personalProfile的长度是否过长。

信息校验没有问题之后，将会利用用户登录时设置在session中的userId，进行数据库查询，然后更新相应的字段。

![image-20230813113406081](C:\Users\isxia\AppData\Roaming\Typora\typora-user-images\image-20230813113406081.png)

4.查看个人信息

![image-20230813113433487](C:\Users\isxia\AppData\Roaming\Typora\typora-user-images\image-20230813113433487.png)

在数据库中也可以查看

![image-20230813113529464](C:\Users\isxia\AppData\Roaming\Typora\typora-user-images\image-20230813113529464.png)
