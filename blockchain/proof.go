package blockchain

import "math/big"


const Difficulty=12

type ProofWork struct{
	
	block *Block
	Target big.Int
}

func Proof (*Block) *ProofofWork{
	target:= big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))

	pow:= &ProofofWork {b, target}

	return pow
}

func (pow *ProofofWork) InitData(nonce int) []byte{
	
	data:=bytes.Join(
		[][]byte{
		pow.Block.PrevHash,
		pow.Block.Data,
		ToHex(int64(nonce)),
		ToHex(int64(Difficulty))
     }, 
	 []byte{},
	)
	return data
}

func ToHex(num int64) [] byte {
	buff:= new(bytes.Buffer)
	err:=binary.Write(buff, binary.BigEndian, num)
	if err!=nil{
		log.Panic(err)
	}
    return buff.Bytes()
}

func (pow *ProofofWork) Run() (int, []byte){
	var intHash big.Int
	var hash [32]byte

	nonce:=0


	for nonce< math.MaxInt64{
		data:=pow.InitData(nonce)
		hash=sha256.Sum256(data)

		fmt.Printf("\r%x", hash)
		intHash.SetBytes(hash[:])   //cmvert hash from byte array to big.Int

		if intHash.Cmp(pow.Target)==-1 {
           break
		} else {
          nonce++
		}
	
        fmt.Println()
		return nonce,hash[:]
	}
}