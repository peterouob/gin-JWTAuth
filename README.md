### 引入包

go get go.mongodb.org/mongo-driver/mongo
go get github.com/gin-gonic/gin

github.com/spf13/viper

### 編寫model模塊

1. 使用*string是避免讓json unmarshall的時候無法判斷為空值還是0
2. validate為驗證方法的tag, eq則類似enum, 強制限定可接收的值
