package diffiehellman

import "math/big"
import "math/rand"
import "time"
// Diffie-Hellman-Merkle key exchange
// Private keys should be generated randomly.

func PrivateKey(p *big.Int) *big.Int {

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
    for {
        newvalue:= new(big.Int)
        newvalue.Rand(r , p)
        if newvalue.Cmp(big.NewInt(1))==1 {
            return newvalue
        }
        
    }
}

func PublicKey(private, p *big.Int, g int64) *big.Int {
    newvalue:= new(big.Int)
	return newvalue.Exp(big.NewInt(g) , private, p )
}

func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	
    newPrivate:= PrivateKey(p)
    newPublic := PublicKey(newPrivate , p , g)

    return  newPrivate, newPublic
}

func SecretKey(private1, public2, p *big.Int) *big.Int {
	newValue:= new(big.Int)
    return newValue.Exp(public2, private1 , p )
}
