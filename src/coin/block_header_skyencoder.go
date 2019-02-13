// Code generated by github.com/skycoin/skyencoder. DO NOT EDIT.
package coin

import "github.com/skycoin/skycoin/src/cipher/encoder"

// encodeSizeBlockHeader computes the size of an encoded object of type BlockHeader
func encodeSizeBlockHeader(obj *BlockHeader) uint64 {
	i0 := uint64(0)

	// obj.Version
	i0 += 4

	// obj.Time
	i0 += 8

	// obj.BkSeq
	i0 += 8

	// obj.Fee
	i0 += 8

	// obj.PrevHash
	i0 += 32

	// obj.BodyHash
	i0 += 32

	// obj.UxHash
	i0 += 32

	return i0
}

// encodeBlockHeader encodes an object of type BlockHeader to the buffer in encoder.Encoder.
// The buffer must be large enough to encode the object, otherwise an error is returned.
func encodeBlockHeader(buf []byte, obj *BlockHeader) error {
	e := &encoder.Encoder{
		Buffer: buf[:],
	}

	// obj.Version
	e.Uint32(obj.Version)

	// obj.Time
	e.Uint64(obj.Time)

	// obj.BkSeq
	e.Uint64(obj.BkSeq)

	// obj.Fee
	e.Uint64(obj.Fee)

	// obj.PrevHash
	e.CopyBytes(obj.PrevHash[:])

	// obj.BodyHash
	e.CopyBytes(obj.BodyHash[:])

	// obj.UxHash
	e.CopyBytes(obj.UxHash[:])

	return nil
}

// decodeBlockHeader decodes an object of type BlockHeader from the buffer in encoder.Decoder.
// Returns the number of bytes used from the buffer to decode the object.
func decodeBlockHeader(buf []byte, obj *BlockHeader) (int, error) {
	d := &encoder.Decoder{
		Buffer: buf[:],
	}

	{
		// obj.Version
		i, err := d.Uint32()
		if err != nil {
			return len(buf) - len(d.Buffer), err
		}
		obj.Version = i
	}

	{
		// obj.Time
		i, err := d.Uint64()
		if err != nil {
			return len(buf) - len(d.Buffer), err
		}
		obj.Time = i
	}

	{
		// obj.BkSeq
		i, err := d.Uint64()
		if err != nil {
			return len(buf) - len(d.Buffer), err
		}
		obj.BkSeq = i
	}

	{
		// obj.Fee
		i, err := d.Uint64()
		if err != nil {
			return len(buf) - len(d.Buffer), err
		}
		obj.Fee = i
	}

	{
		// obj.PrevHash
		if len(d.Buffer) < len(obj.PrevHash) {
			return len(buf) - len(d.Buffer), encoder.ErrBufferUnderflow
		}
		copy(obj.PrevHash[:], d.Buffer[:len(obj.PrevHash)])
		d.Buffer = d.Buffer[len(obj.PrevHash):]
	}

	{
		// obj.BodyHash
		if len(d.Buffer) < len(obj.BodyHash) {
			return len(buf) - len(d.Buffer), encoder.ErrBufferUnderflow
		}
		copy(obj.BodyHash[:], d.Buffer[:len(obj.BodyHash)])
		d.Buffer = d.Buffer[len(obj.BodyHash):]
	}

	{
		// obj.UxHash
		if len(d.Buffer) < len(obj.UxHash) {
			return len(buf) - len(d.Buffer), encoder.ErrBufferUnderflow
		}
		copy(obj.UxHash[:], d.Buffer[:len(obj.UxHash)])
		d.Buffer = d.Buffer[len(obj.UxHash):]
	}

	return len(buf) - len(d.Buffer), nil
}
