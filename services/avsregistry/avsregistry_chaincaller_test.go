package avsregistry

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"math/big"
	"reflect"
	"testing"

	"github.com/Layr-Labs/eigensdk-go/crypto/bls"
	"github.com/Layr-Labs/eigensdk-go/internal/fakes"
	"github.com/Layr-Labs/eigensdk-go/testutils"
	"github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/common"
)

type fakeOperatorInfoService struct {
	operatorInfo types.OperatorInfo
}

func newFakeOperatorInfoService(operatorInfo types.OperatorInfo) *fakeOperatorInfoService {
	return &fakeOperatorInfoService{
		operatorInfo: operatorInfo,
	}
}

func (f *fakeOperatorInfoService) GetOperatorInfo(
	ctx context.Context,
	operator common.Address,
) (operatorInfo types.OperatorInfo, operatorFound bool) {
	return f.operatorInfo, true
}

func TestAvsRegistryServiceChainCaller_getOperatorPubkeys(t *testing.T) {
	logger := testutils.GetTestLogger()
	testOperator1 := fakes.TestOperator{
		OperatorAddr: common.HexToAddress("0x1"),
		OperatorId:   types.OperatorId{1},
		OperatorInfo: types.OperatorInfo{
			Pubkeys: types.OperatorPubkeys{
				G1Pubkey: bls.NewG1Point(big.NewInt(1), big.NewInt(1)),
				G2Pubkey: bls.NewG2Point(
					[2]*big.Int{big.NewInt(1), big.NewInt(1)},
					[2]*big.Int{big.NewInt(1), big.NewInt(1)},
				),
			},
			Socket: "localhost:8080",
		},
	}

	// TODO(samlaf): add error test cases
	var tests = []struct {
		name             string
		operator         *fakes.TestOperator
		queryOperatorId  types.OperatorId
		wantErr          error
		wantOperatorInfo types.OperatorInfo
	}{
		{
			name:             "should return operator info",
			operator:         &testOperator1,
			queryOperatorId:  testOperator1.OperatorId,
			wantErr:          nil,
			wantOperatorInfo: testOperator1.OperatorInfo,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mocks
			mockAvsRegistryReader := fakes.NewFakeAVSRegistryReader(tt.operator, nil)
			mockOperatorsInfoService := newFakeOperatorInfoService(tt.operator.OperatorInfo)

			// Create a new instance of the avsregistry service
			service := NewAvsRegistryServiceChainCaller(mockAvsRegistryReader, mockOperatorsInfoService, logger)

			// Call the GetOperatorPubkeys method with the test operator address
			gotOperatorInfo, gotErr := service.getOperatorInfo(context.Background(), tt.queryOperatorId)
			if !errors.Is(gotErr, tt.wantErr) {
				t.Fatalf("GetOperatorPubkeys returned wrong error. Got: %v, want: %v.", gotErr, tt.wantErr)
			}
			if tt.wantErr == nil && !reflect.DeepEqual(tt.wantOperatorInfo, gotOperatorInfo) {
				t.Fatalf(
					"GetOperatorPubkeys returned wrong operator pubkeys. Got: %v, want: %v.",
					gotOperatorInfo,
					tt.wantOperatorInfo,
				)
			}
		})
	}
}

type TestData struct {
	Input struct {
		QueryQuorumNumbers types.QuorumNums `json:"query_quorum_numbers"`
		QueryBlockNum      int32            `json:"query_block_num"`
	} `json:"input"`
}

// converts a hex string (starting with "0x") into a Bytes32.
func NewBytes32FromString(hexString string) [32]byte {
	var b32 [32]byte

	// Remove the "0x" prefix if it's present
	if len(hexString) >= 2 && hexString[:2] == "0x" {
		hexString = hexString[2:]
	}

	// Decode the hex string
	bytes, _ := hex.DecodeString(hexString)

	// Copy the bytes into the Bytes32 array
	copy(b32[:], bytes)
	return b32
}

func TestAvsRegistryServiceChainCaller_GetOperatorsAvsState(t *testing.T) {
	test_data_path := os.Getenv("TEST_DATA_PATH")
	data, err := ioutil.ReadFile(test_data_path)
	if err != nil {
		t.Fatalf("Failed to read JSON file for test %s: %v", t.Name(), err)
	}

	var testData TestData
	if err := json.Unmarshal(data, &testData); err != nil {
		t.Fatalf("Failed to unmarshal JSON data for test %s: %v", t.Name(), err)
	}

	logger := testutils.GetTestLogger()

	testOperator1 := fakes.TestOperator{
		OperatorAddr: common.HexToAddress("0x1"),
		OperatorId:   types.OperatorId{1},
		OperatorInfo: types.OperatorInfo{
			Pubkeys: types.OperatorPubkeys{
				G1Pubkey: bls.NewG1Point(big.NewInt(1), big.NewInt(1)),
				G2Pubkey: bls.NewG2Point(
					[2]*big.Int{big.NewInt(1), big.NewInt(1)},
					[2]*big.Int{big.NewInt(1), big.NewInt(1)},
				),
			},
			Socket: "localhost:8080",
		},
	}

	var tests = []struct {
		name                      string
		queryQuorumNumbers        types.QuorumNums
		queryBlockNum             types.BlockNum
		wantErr                   error
		wantOperatorsAvsStateDict map[types.OperatorId]types.OperatorAvsState
		operator                  *fakes.TestOperator
	}{
		{
			name:               "should return operatorsAvsState",
			queryQuorumNumbers: testData.Input.QueryQuorumNumbers,
			operator:           &testOperator1,
			queryBlockNum:      uint32(testData.Input.QueryBlockNum),
			wantErr:            nil,
			wantOperatorsAvsStateDict: map[types.OperatorId]types.OperatorAvsState{
				testOperator1.OperatorId: {
					OperatorId:     testOperator1.OperatorId,
					OperatorInfo:   testOperator1.OperatorInfo,
					StakePerQuorum: map[types.QuorumNum]types.StakeAmount{testData.Input.QueryQuorumNumbers[0]: big.NewInt(123)},
					BlockNumber:    uint32(testData.Input.QueryBlockNum),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mocks
			mockAvsRegistryReader := fakes.NewFakeAVSRegistryReader(tt.operator, nil)
			mockOperatorsInfoService := newFakeOperatorInfoService(tt.operator.OperatorInfo)

			// Create a new instance of the avsregistry service
			service := NewAvsRegistryServiceChainCaller(mockAvsRegistryReader, mockOperatorsInfoService, logger)

			// Call the GetOperatorPubkeys method with the test operator address
			gotOperatorsAvsStateDict, gotErr := service.GetOperatorsAvsStateAtBlock(
				context.Background(),
				tt.queryQuorumNumbers,
				tt.queryBlockNum,
			)
			if !errors.Is(gotErr, tt.wantErr) {
				t.Fatalf("GetOperatorsAvsState returned wrong error. Got: %v, want: %v.", gotErr, tt.wantErr)
			}
			if tt.wantErr == nil && !reflect.DeepEqual(tt.wantOperatorsAvsStateDict, gotOperatorsAvsStateDict) {
				t.Fatalf(
					"GetOperatorsAvsState returned wrong operatorsAvsStateDict. Got: %v, want: %v.",
					gotOperatorsAvsStateDict,
					tt.wantOperatorsAvsStateDict,
				)
			}
		})
	}
}

func TestAvsRegistryServiceChainCaller_GetQuorumsAvsState(t *testing.T) {
	logger := testutils.GetTestLogger()
	testOperator1 := fakes.TestOperator{
		OperatorAddr: common.HexToAddress("0x1"),
		OperatorId:   types.OperatorId{1},
		OperatorInfo: types.OperatorInfo{
			Pubkeys: types.OperatorPubkeys{
				G1Pubkey: bls.NewG1Point(big.NewInt(1), big.NewInt(1)),
				G2Pubkey: bls.NewG2Point(
					[2]*big.Int{big.NewInt(1), big.NewInt(1)},
					[2]*big.Int{big.NewInt(1), big.NewInt(1)},
				),
			},
			Socket: "localhost:8080",
		},
	}

	var tests = []struct {
		name                    string
		queryQuorumNumbers      types.QuorumNums
		queryBlockNum           types.BlockNum
		wantErr                 error
		wantQuorumsAvsStateDict map[types.QuorumNum]types.QuorumAvsState
		operator                *fakes.TestOperator
	}{
		{
			name:               "should return operatorsAvsState",
			queryQuorumNumbers: types.QuorumNums{1},
			operator:           &testOperator1,
			queryBlockNum:      1,
			wantErr:            nil,
			wantQuorumsAvsStateDict: map[types.QuorumNum]types.QuorumAvsState{
				1: {
					QuorumNumber: types.QuorumNum(1),
					TotalStake:   big.NewInt(123),
					AggPubkeyG1:  bls.NewG1Point(big.NewInt(1), big.NewInt(1)),
					BlockNumber:  1,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mocks
			mockAvsRegistryReader := fakes.NewFakeAVSRegistryReader(tt.operator, nil)
			mockOperatorsInfoService := newFakeOperatorInfoService(tt.operator.OperatorInfo)

			// Create a new instance of the avsregistry service
			service := NewAvsRegistryServiceChainCaller(mockAvsRegistryReader, mockOperatorsInfoService, logger)

			// Call the GetOperatorPubkeys method with the test operator address
			aggG1PubkeyPerQuorum, gotErr := service.GetQuorumsAvsStateAtBlock(
				context.Background(),
				tt.queryQuorumNumbers,
				tt.queryBlockNum,
			)
			if !errors.Is(gotErr, tt.wantErr) {
				t.Fatalf("GetOperatorsAvsState returned wrong error. Got: %v, want: %v.", gotErr, tt.wantErr)
			}
			if tt.wantErr == nil && !reflect.DeepEqual(tt.wantQuorumsAvsStateDict, aggG1PubkeyPerQuorum) {
				t.Fatalf(
					"GetOperatorsAvsState returned wrong aggG1PubkeyPerQuorum. Got: %v, want: %v.",
					aggG1PubkeyPerQuorum,
					tt.wantQuorumsAvsStateDict,
				)
			}
		})
	}
}
