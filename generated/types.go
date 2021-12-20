// WARNING: This file has automatically been generated
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package generated

/*
#cgo LDFLAGS: -L${SRCDIR}/..
#cgo pkg-config: ${SRCDIR}/../filcrypto.pc
#include "../filcrypto.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

// FilBLSSignature as declared in filecoin-ffi/filcrypto.h:77
type FilBLSSignature struct {
	Inner          [96]byte
	refa2ac09ba    *C.fil_BLSSignature
	allocsa2ac09ba interface{}
}

// FilAggregateResponse as declared in filecoin-ffi/filcrypto.h:84
type FilAggregateResponse struct {
	Signature      FilBLSSignature
	refb3efa36d    *C.fil_AggregateResponse
	allocsb3efa36d interface{}
}

// FilAggregateProof as declared in filecoin-ffi/filcrypto.h:91
type FilAggregateProof struct {
	StatusCode     FCPResponseStatus
	ErrorMsg       string
	ProofLen       uint
	ProofPtr       []byte
	ref22b6c4f6    *C.fil_AggregateProof
	allocs22b6c4f6 interface{}
}

// Fil32ByteArray as declared in filecoin-ffi/filcrypto.h:95
type Fil32ByteArray struct {
	Inner          [32]byte
	ref373ec61a    *C.fil_32ByteArray
	allocs373ec61a interface{}
}

// FilAggregationInputs as declared in filecoin-ffi/filcrypto.h:103
type FilAggregationInputs struct {
	CommR          Fil32ByteArray
	CommD          Fil32ByteArray
	SectorId       uint64
	Ticket         Fil32ByteArray
	Seed           Fil32ByteArray
	ref90b967c9    *C.fil_AggregationInputs
	allocs90b967c9 interface{}
}

// FilSealCommitPhase2Response as declared in filecoin-ffi/filcrypto.h:112
type FilSealCommitPhase2Response struct {
	StatusCode      FCPResponseStatus
	ErrorMsg        string
	ProofPtr        []byte
	ProofLen        uint
	CommitInputsPtr []FilAggregationInputs
	CommitInputsLen uint
	ref5860b9a4     *C.fil_SealCommitPhase2Response
	allocs5860b9a4  interface{}
}

// FilClearCacheResponse as declared in filecoin-ffi/filcrypto.h:117
type FilClearCacheResponse struct {
	ErrorMsg       string
	StatusCode     FCPResponseStatus
	refa9a80400    *C.fil_ClearCacheResponse
	allocsa9a80400 interface{}
}

// FilCreateFvmMachineResponse as declared in filecoin-ffi/filcrypto.h:123
type FilCreateFvmMachineResponse struct {
	ErrorMsg       string
	StatusCode     FCPResponseStatus
	MachineId      uint64
	ref40465416    *C.fil_CreateFvmMachineResponse
	allocs40465416 interface{}
}

// FilZeroSignatureResponse as declared in filecoin-ffi/filcrypto.h:130
type FilZeroSignatureResponse struct {
	Signature      FilBLSSignature
	ref835a0405    *C.fil_ZeroSignatureResponse
	allocs835a0405 interface{}
}

// FilDropFvmMachineResponse as declared in filecoin-ffi/filcrypto.h:135
type FilDropFvmMachineResponse struct {
	ErrorMsg       string
	StatusCode     FCPResponseStatus
	ref94e2357b    *C.fil_DropFvmMachineResponse
	allocs94e2357b interface{}
}

// FilEmptySectorUpdateDecodeFromResponse as declared in filecoin-ffi/filcrypto.h:140
type FilEmptySectorUpdateDecodeFromResponse struct {
	StatusCode     FCPResponseStatus
	ErrorMsg       string
	reff02a01b8    *C.fil_EmptySectorUpdateDecodeFromResponse
	allocsf02a01b8 interface{}
}

// FilEmptySectorUpdateEncodeIntoResponse as declared in filecoin-ffi/filcrypto.h:148
type FilEmptySectorUpdateEncodeIntoResponse struct {
	ErrorMsg       string
	StatusCode     FCPResponseStatus
	CommRNew       [32]byte
	CommRLastNew   [32]byte
	CommDNew       [32]byte
	ref8d3238a7    *C.fil_EmptySectorUpdateEncodeIntoResponse
	allocs8d3238a7 interface{}
}

// FilEmptySectorUpdateProofResponse as declared in filecoin-ffi/filcrypto.h:155
type FilEmptySectorUpdateProofResponse struct {
	StatusCode    FCPResponseStatus
	ErrorMsg      string
	ProofLen      uint
	ProofPtr      []byte
	ref5c2faef    *C.fil_EmptySectorUpdateProofResponse
	allocs5c2faef interface{}
}

// FilEmptySectorUpdateRemoveEncodedDataResponse as declared in filecoin-ffi/filcrypto.h:160
type FilEmptySectorUpdateRemoveEncodedDataResponse struct {
	StatusCode     FCPResponseStatus
	ErrorMsg       string
	ref50783b83    *C.fil_EmptySectorUpdateRemoveEncodedDataResponse
	allocs50783b83 interface{}
}

// FilVerifyEmptySectorUpdateProofResponse as declared in filecoin-ffi/filcrypto.h:166
type FilVerifyEmptySectorUpdateProofResponse struct {
	StatusCode    FCPResponseStatus
	ErrorMsg      string
	IsValid       bool
	ref50b7b13    *C.fil_VerifyEmptySectorUpdateProofResponse
	allocs50b7b13 interface{}
}

// FilFauxRepResponse as declared in filecoin-ffi/filcrypto.h:172
type FilFauxRepResponse struct {
	ErrorMsg       string
	StatusCode     FCPResponseStatus
	Commitment     [32]byte
	refaa003f71    *C.fil_FauxRepResponse
	allocsaa003f71 interface{}
}

// FilFinalizeTicketResponse as declared in filecoin-ffi/filcrypto.h:178
type FilFinalizeTicketResponse struct {
	StatusCode     FCPResponseStatus
	ErrorMsg       string
	Ticket         [32]byte
	refb370fa86    *C.fil_FinalizeTicketResponse
	allocsb370fa86 interface{}
}

// FilFvmMachineExecuteResponse as declared in filecoin-ffi/filcrypto.h:183
type FilFvmMachineExecuteResponse struct {
	ErrorMsg       string
	StatusCode     FCPResponseStatus
	ref88f63595    *C.fil_FvmMachineExecuteResponse
	allocs88f63595 interface{}
}

// FilGenerateDataCommitmentResponse as declared in filecoin-ffi/filcrypto.h:189
type FilGenerateDataCommitmentResponse struct {
	StatusCode     FCPResponseStatus
	ErrorMsg       string
	CommD          [32]byte
	ref87da7dd9    *C.fil_GenerateDataCommitmentResponse
	allocs87da7dd9 interface{}
}

// FilPartitionProof as declared in filecoin-ffi/filcrypto.h:194
type FilPartitionProof struct {
	ProofLen       uint
	ProofPtr       []byte
	ref566a2be6    *C.fil_PartitionProof
	allocs566a2be6 interface{}
}

// FilPartitionProofResponse as declared in filecoin-ffi/filcrypto.h:201
type FilPartitionProofResponse struct {
	StatusCode     FCPResponseStatus
	ErrorMsg       string
	ProofsLen      uint
	ProofsPtr      []FilPartitionProof
	ref51343e7a    *C.fil_PartitionProofResponse
	allocs51343e7a interface{}
}

// FilGenerateFallbackSectorChallengesResponse as declared in filecoin-ffi/filcrypto.h:211
type FilGenerateFallbackSectorChallengesResponse struct {
	ErrorMsg         string
	StatusCode       FCPResponseStatus
	IdsPtr           []uint64
	IdsLen           uint
	ChallengesPtr    []uint64
	ChallengesLen    uint
	ChallengesStride uint
	ref7047a3fa      *C.fil_GenerateFallbackSectorChallengesResponse
	allocs7047a3fa   interface{}
}

// FilGeneratePieceCommitmentResponse as declared in filecoin-ffi/filcrypto.h:222
type FilGeneratePieceCommitmentResponse struct {
	StatusCode      FCPResponseStatus
	ErrorMsg        string
	CommP           [32]byte
	NumBytesAligned uint64
	ref4b00fda4     *C.fil_GeneratePieceCommitmentResponse
	allocs4b00fda4  interface{}
}

// FilVanillaProof as declared in filecoin-ffi/filcrypto.h:227
type FilVanillaProof struct {
	ProofLen       uint
	ProofPtr       []byte
	refb3e7638c    *C.fil_VanillaProof
	allocsb3e7638c interface{}
}

// FilGenerateSingleVanillaProofResponse as declared in filecoin-ffi/filcrypto.h:233
type FilGenerateSingleVanillaProofResponse struct {
	ErrorMsg       string
	VanillaProof   FilVanillaProof
	StatusCode     FCPResponseStatus
	reff9d21b04    *C.fil_GenerateSingleVanillaProofResponse
	allocsf9d21b04 interface{}
}

// FilPartitionSnarkProof as declared in filecoin-ffi/filcrypto.h:239
type FilPartitionSnarkProof struct {
	RegisteredProof FilRegisteredPoStProof
	ProofLen        uint
	ProofPtr        []byte
	ref4de03739     *C.fil_PartitionSnarkProof
	allocs4de03739  interface{}
}

// FilGenerateSingleWindowPoStWithVanillaResponse as declared in filecoin-ffi/filcrypto.h:247
type FilGenerateSingleWindowPoStWithVanillaResponse struct {
	ErrorMsg         string
	PartitionProof   FilPartitionSnarkProof
	FaultySectorsLen uint
	FaultySectorsPtr []uint64
	StatusCode       FCPResponseStatus
	ref96c012c3      *C.fil_GenerateSingleWindowPoStWithVanillaResponse
	allocs96c012c3   interface{}
}

// FilPoStProof as declared in filecoin-ffi/filcrypto.h:253
type FilPoStProof struct {
	RegisteredProof FilRegisteredPoStProof
	ProofLen        uint
	ProofPtr        []byte
	ref3451bfa      *C.fil_PoStProof
	allocs3451bfa   interface{}
}

// FilGenerateWindowPoStResponse as declared in filecoin-ffi/filcrypto.h:262
type FilGenerateWindowPoStResponse struct {
	ErrorMsg         string
	ProofsLen        uint
	ProofsPtr        []FilPoStProof
	FaultySectorsLen uint
	FaultySectorsPtr []uint64
	StatusCode       FCPResponseStatus
	ref2a5f3ba8      *C.fil_GenerateWindowPoStResponse
	allocs2a5f3ba8   interface{}
}

// FilGenerateWinningPoStResponse as declared in filecoin-ffi/filcrypto.h:269
type FilGenerateWinningPoStResponse struct {
	ErrorMsg       string
	ProofsLen      uint
	ProofsPtr      []FilPoStProof
	StatusCode     FCPResponseStatus
	ref1405b8ec    *C.fil_GenerateWinningPoStResponse
	allocs1405b8ec interface{}
}

// FilGenerateWinningPoStSectorChallenge as declared in filecoin-ffi/filcrypto.h:276
type FilGenerateWinningPoStSectorChallenge struct {
	ErrorMsg       string
	StatusCode     FCPResponseStatus
	IdsPtr         []uint64
	IdsLen         uint
	ref69d2a405    *C.fil_GenerateWinningPoStSectorChallenge
	allocs69d2a405 interface{}
}

// FilGetNumPartitionForFallbackPoStResponse as declared in filecoin-ffi/filcrypto.h:282
type FilGetNumPartitionForFallbackPoStResponse struct {
	ErrorMsg       string
	StatusCode     FCPResponseStatus
	NumPartition   uint
	refc0084478    *C.fil_GetNumPartitionForFallbackPoStResponse
	allocsc0084478 interface{}
}

// FilGpuDeviceResponse as declared in filecoin-ffi/filcrypto.h:289
type FilGpuDeviceResponse struct {
	StatusCode     FCPResponseStatus
	ErrorMsg       string
	DevicesLen     uint
	DevicesPtr     []string
	ref58f92915    *C.fil_GpuDeviceResponse
	allocs58f92915 interface{}
}

// FilBLSDigest as declared in filecoin-ffi/filcrypto.h:293
type FilBLSDigest struct {
	Inner          [96]byte
	ref215fc78c    *C.fil_BLSDigest
	allocs215fc78c interface{}
}

// FilHashResponse as declared in filecoin-ffi/filcrypto.h:300
type FilHashResponse struct {
	Digest         FilBLSDigest
	refc52a22ef    *C.fil_HashResponse
	allocsc52a22ef interface{}
}

// FilInitLogFdResponse as declared in filecoin-ffi/filcrypto.h:305
type FilInitLogFdResponse struct {
	StatusCode     FCPResponseStatus
	ErrorMsg       string
	ref3c1a0a08    *C.fil_InitLogFdResponse
	allocs3c1a0a08 interface{}
}

// FilMergeWindowPoStPartitionProofsResponse as declared in filecoin-ffi/filcrypto.h:311
type FilMergeWindowPoStPartitionProofsResponse struct {
	ErrorMsg       string
	Proof          FilPoStProof
	StatusCode     FCPResponseStatus
	ref3369154e    *C.fil_MergeWindowPoStPartitionProofsResponse
	allocs3369154e interface{}
}

// FilBLSPrivateKey as declared in filecoin-ffi/filcrypto.h:315
type FilBLSPrivateKey struct {
	Inner          [32]byte
	ref2f77fe3a    *C.fil_BLSPrivateKey
	allocs2f77fe3a interface{}
}

// FilPrivateKeyGenerateResponse as declared in filecoin-ffi/filcrypto.h:322
type FilPrivateKeyGenerateResponse struct {
	PrivateKey    FilBLSPrivateKey
	ref2dba09f    *C.fil_PrivateKeyGenerateResponse
	allocs2dba09f interface{}
}

// FilBLSPublicKey as declared in filecoin-ffi/filcrypto.h:326
type FilBLSPublicKey struct {
	Inner          [48]byte
	ref6d0cab13    *C.fil_BLSPublicKey
	allocs6d0cab13 interface{}
}

// FilPrivateKeyPublicKeyResponse as declared in filecoin-ffi/filcrypto.h:333
type FilPrivateKeyPublicKeyResponse struct {
	PublicKey      FilBLSPublicKey
	refee14e59d    *C.fil_PrivateKeyPublicKeyResponse
	allocsee14e59d interface{}
}

// FilPrivateKeySignResponse as declared in filecoin-ffi/filcrypto.h:340
type FilPrivateKeySignResponse struct {
	Signature      FilBLSSignature
	refcdf97b28    *C.fil_PrivateKeySignResponse
	allocscdf97b28 interface{}
}

// FilSealCommitPhase1Response as declared in filecoin-ffi/filcrypto.h:347
type FilSealCommitPhase1Response struct {
	StatusCode                FCPResponseStatus
	ErrorMsg                  string
	SealCommitPhase1OutputPtr []byte
	SealCommitPhase1OutputLen uint
	ref61ed8561               *C.fil_SealCommitPhase1Response
	allocs61ed8561            interface{}
}

// FilSealPreCommitPhase1Response as declared in filecoin-ffi/filcrypto.h:354
type FilSealPreCommitPhase1Response struct {
	ErrorMsg                     string
	StatusCode                   FCPResponseStatus
	SealPreCommitPhase1OutputPtr []byte
	SealPreCommitPhase1OutputLen uint
	ref132bbfd8                  *C.fil_SealPreCommitPhase1Response
	allocs132bbfd8               interface{}
}

// FilSealPreCommitPhase2Response as declared in filecoin-ffi/filcrypto.h:362
type FilSealPreCommitPhase2Response struct {
	ErrorMsg        string
	StatusCode      FCPResponseStatus
	RegisteredProof FilRegisteredSealProof
	CommD           [32]byte
	CommR           [32]byte
	ref2aa6831d     *C.fil_SealPreCommitPhase2Response
	allocs2aa6831d  interface{}
}

// FilStringResponse as declared in filecoin-ffi/filcrypto.h:371
type FilStringResponse struct {
	StatusCode     FCPResponseStatus
	ErrorMsg       string
	StringVal      string
	ref4f413043    *C.fil_StringResponse
	allocs4f413043 interface{}
}

// FilUnsealRangeResponse as declared in filecoin-ffi/filcrypto.h:376
type FilUnsealRangeResponse struct {
	StatusCode     FCPResponseStatus
	ErrorMsg       string
	ref61e219c9    *C.fil_UnsealRangeResponse
	allocs61e219c9 interface{}
}

// FilVerifyAggregateSealProofResponse as declared in filecoin-ffi/filcrypto.h:382
type FilVerifyAggregateSealProofResponse struct {
	StatusCode    FCPResponseStatus
	ErrorMsg      string
	IsValid       bool
	ref66180e0    *C.fil_VerifyAggregateSealProofResponse
	allocs66180e0 interface{}
}

// FilVerifyPartitionProofResponse as declared in filecoin-ffi/filcrypto.h:388
type FilVerifyPartitionProofResponse struct {
	StatusCode    FCPResponseStatus
	ErrorMsg      string
	IsValid       bool
	refaed1b67    *C.fil_VerifyPartitionProofResponse
	allocsaed1b67 interface{}
}

// FilVerifySealResponse as declared in filecoin-ffi/filcrypto.h:394
type FilVerifySealResponse struct {
	StatusCode     FCPResponseStatus
	ErrorMsg       string
	IsValid        bool
	refd4397079    *C.fil_VerifySealResponse
	allocsd4397079 interface{}
}

// FilVerifyWindowPoStResponse as declared in filecoin-ffi/filcrypto.h:400
type FilVerifyWindowPoStResponse struct {
	StatusCode     FCPResponseStatus
	ErrorMsg       string
	IsValid        bool
	ref34c4d49f    *C.fil_VerifyWindowPoStResponse
	allocs34c4d49f interface{}
}

// FilVerifyWinningPoStResponse as declared in filecoin-ffi/filcrypto.h:406
type FilVerifyWinningPoStResponse struct {
	StatusCode     FCPResponseStatus
	ErrorMsg       string
	IsValid        bool
	refaca6860c    *C.fil_VerifyWinningPoStResponse
	allocsaca6860c interface{}
}

// FilWriteWithAlignmentResponse as declared in filecoin-ffi/filcrypto.h:414
type FilWriteWithAlignmentResponse struct {
	CommP                 [32]byte
	ErrorMsg              string
	LeftAlignmentUnpadded uint64
	StatusCode            FCPResponseStatus
	TotalWriteUnpadded    uint64
	refa330e79            *C.fil_WriteWithAlignmentResponse
	allocsa330e79         interface{}
}

// FilWriteWithoutAlignmentResponse as declared in filecoin-ffi/filcrypto.h:421
type FilWriteWithoutAlignmentResponse struct {
	CommP              [32]byte
	ErrorMsg           string
	StatusCode         FCPResponseStatus
	TotalWriteUnpadded uint64
	refc8e1ed8         *C.fil_WriteWithoutAlignmentResponse
	allocsc8e1ed8      interface{}
}

// FilPublicPieceInfo as declared in filecoin-ffi/filcrypto.h:426
type FilPublicPieceInfo struct {
	NumBytes       uint64
	CommP          [32]byte
	refd00025ac    *C.fil_PublicPieceInfo
	allocsd00025ac interface{}
}

// FilPrivateReplicaInfo as declared in filecoin-ffi/filcrypto.h:434
type FilPrivateReplicaInfo struct {
	RegisteredProof FilRegisteredPoStProof
	CacheDirPath    string
	CommR           [32]byte
	ReplicaPath     string
	SectorId        uint64
	ref81a31e9b     *C.fil_PrivateReplicaInfo
	allocs81a31e9b  interface{}
}

// FilPublicReplicaInfo as declared in filecoin-ffi/filcrypto.h:440
type FilPublicReplicaInfo struct {
	RegisteredProof FilRegisteredPoStProof
	CommR           [32]byte
	SectorId        uint64
	ref81b617c2     *C.fil_PublicReplicaInfo
	allocs81b617c2  interface{}
}
