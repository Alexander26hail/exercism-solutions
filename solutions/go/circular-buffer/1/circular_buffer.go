package circularbuffer

// Implement a circular buffer of bytes supporting both overflow-checked writes
// and unconditional, possibly overwriting, writes.
//
// We chose the provided API so that Buffer implements io.ByteReader
// and io.ByteWriter and can be used (size permitting) as a drop in
// replacement for anything using that interface.

// Define the Buffer type here.
import "errors"
type Buffer struct {
    buffer []byte  
    write  int     
    read   int     
    size   int  
    count  int
}

func NewBuffer(size int) *Buffer {
	return &Buffer{ buffer: make([]byte, size),
                   write:0,
                   read:0,
                   size:size,
                   count:0,
                  }
}

func (b *Buffer) ReadByte() (byte, error) {
	if b.count == 0{
        return 0, errors.New("Error")
    }
    value:= b.buffer[b.read]
	b.read = (b.read  + 1 ) % b.size
    b.count--
    return value , nil
}

func (b *Buffer) WriteByte(c byte) error {
	if b.count == b.size {
        return errors.New("Error")
    }
    
    b.buffer[b.write]= c
    b.write = (b.write + 1) % b.size
    b.count++
    return nil
}

func (b *Buffer) Overwrite(c byte) {
	if b.count== b.size  {
    	b.read = (b.read + 1) % b.size
    }else{
        b.count++
    }
    b.buffer[b.write]= c
    b.write = (b.write + 1) % b.size

}

func (b *Buffer) Reset() {
	
   b.read = 0
   b.write = 0
   b.count = 0
}
