package main

import (

	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"encoding/binary"
	"math/big"
	"fmt"
	"io"
	"os"
	"strconv"
	"gocv.io/x/gocv"

	quic "github.com/lucas-clemente/quic-go"
)

//The reciever function that recieves the frames from the sender
//input args - the directory to store the frames. Run the viewer function to show the video

const quicServerAddr = "0.0.0.0:4242"

func HandleError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}


func main() {

	videoDir := os.Args[1]
	fmt.Println("Saving Video in: ", videoDir)

	quicConfig := &quic.Config{
		CreatePaths: true,
	}

	// initializing mpquic server
	fmt.Println("Attaching to: ", quicServerAddr)
	listener, err := quic.ListenAddr(quicServerAddr,generateTLSConfig(), quicConfig)
	HandleError(err)

	fmt.Println("Server started! Waiting for streams from client...")

	sess, err := listener.Accept()//accepting connection from sender
	HandleError(err)

	fmt.Println("session created: ", sess.RemoteAddr())

	stream, err := sess.AcceptStream()
	HandleError(err)

	defer stream.Close()

	fmt.Println("stream created: ", stream.StreamID())

	frame_counter := 0

	for {
		siz := make([]byte, 60)// size is needed to make use of ReadFull(). ReadAll() needs EOF to stop accepting while ReadFull just needs the fixed size.

		_,err := io.ReadFull(stream,siz)//recieve the size
		data := binary.LittleEndian.Uint64(siz)//if the first few bytes contain the length; else use BigEndian or reverse the byte[] and use LittleEndian
		HandleError(err)
		
		if(data==0){
			defer stream.Close()
			return
		}

		buff := make([]byte, data)	
		len2,err := io.ReadFull(stream,buff)// recieve image

		HandleError(err)
		
		//if empty buffer
		if(len2==0){
			defer stream.Close()
			return
		}
		
		img,err := gocv.IMDecode(buff,1)//IMReadFlag 1 ensure that image is converted to 3 channel RGB
		
		HandleError(err)
		// if decoding fails

		if(img.Empty()){
			defer stream.Close()
			return
		}

		//everything good !!
		//save image and call viewer.py which shows the stream 
		
		
		file, err := os.Create(videoDir + "/img" + strconv.Itoa(frame_counter) + ".jpg")
	 	HandleError(err)
	 	fmt.Println(frame_counter)

	 	gocv.IMWrite(videoDir + "/img" + strconv.Itoa(frame_counter) + ".jpg",img)
	 	frame_counter += 1
	 	fmt.Println(videoDir + "/img" + strconv.Itoa(frame_counter) + ".jpg")
	 	
	 	file.Close()
	}

}

func generateTLSConfig() *tls.Config {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	template := x509.Certificate{SerialNumber: big.NewInt(1)}
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
	if err != nil {
		panic(err)
	}

	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})

	tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		panic(err)
	}
	return &tls.Config{Certificates: []tls.Certificate{tlsCert}}
}
