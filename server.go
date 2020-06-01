package prisma

import "net"

//MessageHandler func to hanle message
type MessageHandler func(m *Message)

//Server udp serevr
type Server struct {
	pr      *Prisma
	handler MessageHandler
	port    string
	cnn     net.PacketConn
}

//DefaultServer create default server
func DefaultServer(handler MessageHandler) *Server {
	return &Server{
		pr:      DefaultPrisma(),
		handler: handler,
		port:    "21845"}
}

//Run create and run udp server
func (s *Server) Run() error {
	pc, err := net.ListenPacket("udp", ":"+s.port)
	if err != nil {
		return err
	}
	s.cnn = pc
	defer pc.Close()

	for err == nil {
		buf := make([]byte, 1024)
		//n, addr, err := pc.ReadFrom(buf)
		n, _, err := pc.ReadFrom(buf)
		if err != nil {
			break
		}
		m, err := s.pr.Decode(buf[:n])
		if err != nil {
			//TODO some log
			continue
		}
		if s.handler != nil {
			go s.handler(m)
		}
	}
	return err
}

//Stop udp server
func (s *Server) Stop() error {
	return s.cnn.Close()
}
