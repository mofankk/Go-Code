# 测试结论

使用gorm框架，如果用gorm的tag标识外建,但不再数据库表结构上实际定义外建，是可以使用gorm的preload方法的。

### 注意：
1. 数据库表结构要用自己写的SQL来生成
2. 在struct上用了gorm tag，即`gorm:"foreignKey: xxx; references: xxx"`,就不要对该struct 施行 AutoMigrate 方法
3. 然后在代码中就可以使用preload的这种嵌套效果，即满足了DBA不要用数据库外建的建议，又在代码中享受到了外建带来的便利。