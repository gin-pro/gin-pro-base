package uid

import (
	"crypto/md5"
	randc "crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"math/rand"
	"os"
	"sync/atomic"
	"time"
)

// RandomString 随机生成字符串
func RandomString(l int) string {
	str := "0123456789AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz"
	bts := []byte(str)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bts[r.Intn(len(bts))])
	}
	return string(result)
}

// NewXid 生成分布式唯一ID
func NewXid() string {
	return hex.EncodeToString(newXid())
}

// NewXid64 生成唯一ID-base64
func NewXid64() string {
	return base64.RawURLEncoding.EncodeToString(newXid())
}

func newXid() []byte {
	var b [12]byte
	// Timestamp, 4 bytes, big endian
	binary.BigEndian.PutUint32(b[:], uint32(time.Now().Unix()))
	// binary.BigEndian.PutUint32(b[:], uint32(time.Now().UnixNano()/1e6))
	// Machine, first 3 bytes of md5(hostname)
	b[4] = machineId[0]
	b[5] = machineId[1]
	b[6] = machineId[2]
	// Pid, 2 bytes, specs don't specify endianness, but we use big endian.
	pid := os.Getpid()
	b[7] = byte(pid >> 8)
	b[8] = byte(pid)
	// Increment, 3 bytes, big endian
	i := atomic.AddUint32(&objectIdCounter, 1)
	b[9] = byte(i >> 16)
	b[10] = byte(i >> 8)
	b[11] = byte(i)
	return b[:]
}

// objectIdCounter is atomically incremented when generating a new Xid
var objectIdCounter uint32 = 0

// machineId stores machine id generated once and used in subsequent calls
var machineId = readMachineId()

// initMachineId generates machine id and puts it into the machineId global
// variable. If this function fails to get the hostname, it will cause
// a runtime error.
func readMachineId() []byte {
	var sum [3]byte
	id := sum[:]
	hostname, err1 := os.Hostname()
	if err1 != nil {
		_, err2 := io.ReadFull(randc.Reader, id)
		if err2 != nil {
			panic(fmt.Errorf("cannot get hostname: %v; %v", err1, err2))
		}
		return id
	}
	hw := md5.New()
	hw.Write([]byte(hostname))
	copy(id, hw.Sum(nil))
	return id
}
