package 03rpc



import (
	"net"

)

//测试网络中读写数据的情况

//会话链接的结构体
type Session struct{
	conn net.Conn

}
//构造方法
func NewSession(conn net.Conn) *Session{
	return &Session{conn:conn}
}

//向连接中写数据
func (s *Session)Write(data []byte) error {
	//定义写数据的格式,4字节的头部+可变体的长度
	buf := make([]byte,4+len(data))
	//写入头部，记录数据长度
	
}