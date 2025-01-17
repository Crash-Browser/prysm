package verification

import "github.com/prysmaticlabs/prysm/v4/consensus-types/blocks"

// BlobSidecarNoop is a FAKE verification function that simply launders a ROBlob->VerifiedROBlob.
// TODO: find all code that uses this method and replace it with full verification.
func BlobSidecarNoop(b blocks.ROBlob) (blocks.VerifiedROBlob, error) {
	return blocks.NewVerifiedROBlob(b), nil
}

// BlobSidecarSliceNoop is a FAKE verification function that simply launders a ROBlob->VerifiedROBlob.
// TODO: find all code that uses this method and replace it with full verification.
func BlobSidecarSliceNoop(b []blocks.ROBlob) ([]blocks.VerifiedROBlob, error) {
	vbs := make([]blocks.VerifiedROBlob, len(b))
	for i := range b {
		vbs[i] = blocks.NewVerifiedROBlob(b[i])
	}
	return vbs, nil
}
