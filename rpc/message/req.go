package message

import "encoding/binary"

func EncodeReq(req *Request) []byte {
	bs := make([]byte, req.BodyLength+req.HeadLength)

	binary.BigEndian.PutUint32(bs, req.HeadLength)

	binary.BigEndian.PutUint32(bs[4:8], req.HeadLength)

	binary.BigEndian.PutUint32(bs[8:12], req.RequestID)

	bs[12] = req.Version
	bs[13] = req.Compresser
	bs[14] = req.Serializer

	copy(bs[15:], req.ServiceName)

	return bs
}

func DecodeReq(data []byte) *Request {

	req := &Request{}

	req.HeadLength = binary.BigEndian.Uint32(data[:4])

	req.BodyLength = binary.BigEndian.Uint32(data[4:8])

	req.RequestID = binary.BigEndian.Uint32(data[8:12])

	req.Version = data[12]
	req.Compresser = data[13]
	req.Serializer = data[14]

	return nil
}
