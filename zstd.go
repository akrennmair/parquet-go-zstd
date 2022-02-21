package zstd

import (
	goparquet "github.com/fraugster/parquet-go"
	"github.com/fraugster/parquet-go/parquet"
	"github.com/klauspost/compress/zstd"
)

type ZstdBlockCompressor struct {
	enc *zstd.Encoder
	dec *zstd.Decoder
}

func NewZstdBlockCompressor() *ZstdBlockCompressor {
	dec, _ := zstd.NewReader(nil)
	enc, _ := zstd.NewWriter(nil)
	return &ZstdBlockCompressor{
		enc: enc,
		dec: dec,
	}
}

func (c *ZstdBlockCompressor) CompressBlock(data []byte) ([]byte, error) {
	return c.enc.EncodeAll(data, nil), nil
}

func (c *ZstdBlockCompressor) DecompressBlock(data []byte) ([]byte, error) {
	return c.dec.DecodeAll(data, nil)
}

func init() {
	goparquet.RegisterBlockCompressor(parquet.CompressionCodec_ZSTD, NewZstdBlockCompressor())
}
