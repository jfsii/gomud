package main

const (
	PORT = 6060
)

var Log = NewSyslog()

func main() {
	Log.Std("go Mud() init")
	s := NewServer(PORT)
	Log.Std("Launching new server on port: ", s.Port)
	go s.Run()
    <-s.shutdown
}
