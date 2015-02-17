package mbapi

import (
	"fmt"
	"io"
	"log"
	"os"

	"bazil.org/fuse"
	"bazil.org/fuse/fs"
	"bitbucket.org/konair_konairius/mediafs/eventfs"
)

const MFILE = 0444

type StreamFile struct {
	Inode  uint64
	Mode   os.FileMode
	Size   uint64
	reader io.ReaderAt
	name   string
	// oc     int
}

func (b *Backend) NewStreamFile(i *BaseItemDto) (*StreamFile, error) {
	if len(i.MediaSources) != 1 {
		return nil, fmt.Errorf("We can only handle Files with 1 MediaSource (got %d)", len(i.MediaSources))
	}

	inode := eventfs.GetInode()
	net, err := b.GetStreamReader(i)
	if err != nil {
		return nil, fmt.Errorf("Failed to get stream: %v\n", err)
	}
	// b, err := filesystem.cache.NewMemoryFile(net, inode)
	// if err != nil {
	// 	return nil, fmt.Errorf("Failed to create Buffer:%v", err)
	// }
	f := &StreamFile{
		Inode:  inode,
		Mode:   MFILE,
		Size:   uint64(i.MediaSources[0].Size),
		reader: net,
		name:   i.Name,
	}
	return f, nil
}

func (s *StreamFile) Name() string {
	return s.name
}

func (s *StreamFile) Attr() fuse.Attr {
	return fuse.Attr{Inode: s.Inode, Mode: s.Mode, Size: s.Size}
}

func (s *StreamFile) Read(req *fuse.ReadRequest, resp *fuse.ReadResponse, intr fs.Intr) fuse.Error {
	// go log.Printf("Trying to Read: %v byte offset %v of %v ", req.Size, req.Offset, s.Size)
	p := make([]byte, req.Size)
	_, err := s.reader.ReadAt(p, req.Offset)
	if err != nil && err != io.EOF {
		log.Printf("IO Error\n")
		return fuse.EIO
	}
	resp.Data = p
	// log.Printf("Creating new Reader: (ofMediaFSet=%v,blocksize=%v)\n", buf.ofMediaFSet, buf.blocksize)
	return nil
}
