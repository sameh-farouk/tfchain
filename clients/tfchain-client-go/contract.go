package substrate

import (
	"fmt"

	"github.com/centrifuge/go-substrate-rpc-client/v4/scale"
	"github.com/centrifuge/go-substrate-rpc-client/v4/types"
	"github.com/pkg/errors"
)

type DeletedState struct {
	IsCanceledByUser bool `json:"is_canceled_by_user"`
	IsOutOfFunds     bool `json:"is_out_of_funds"`
}

// Decode implementation for the enum type
func (r *DeletedState) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()
	if err != nil {
		return err
	}

	switch b {
	case 0:
		r.IsCanceledByUser = true
	case 1:
		r.IsOutOfFunds = true
	default:
		return fmt.Errorf("unknown deleted state value")
	}

	return nil
}

// Encode implementation
func (r DeletedState) Encode(encoder scale.Encoder) (err error) {
	if r.IsCanceledByUser {
		err = encoder.PushByte(0)
	} else if r.IsOutOfFunds {
		err = encoder.PushByte(1)
	}
	return
}

// ContractState enum
type ContractState struct {
	IsCreated                bool         `json:"is_created"`
	IsDeleted                bool         `json:"is_deleted"`
	AsDeleted                DeletedState `json:"as_deleted"`
	IsGracePeriod            bool         `json:"is_grace_period"`
	AsGracePeriodBlockNumber types.U64    `json:"as_grace_period_block_number"`
}

// Decode implementation for the enum type
func (r *ContractState) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()
	if err != nil {
		return err
	}

	switch b {
	case 0:
		r.IsCreated = true
	case 1:
		r.IsDeleted = true
		if err := decoder.Decode(&r.AsDeleted); err != nil {
			return errors.Wrap(err, "failed to get deleted state")
		}
	case 2:
		r.IsGracePeriod = true
		if err := decoder.Decode(&r.AsGracePeriodBlockNumber); err != nil {
			return errors.Wrap(err, "failed to get grace period")
		}
	default:
		return fmt.Errorf("unknown ContractState value")
	}

	return nil
}

// Encode implementation
func (r ContractState) Encode(encoder scale.Encoder) (err error) {
	if r.IsCreated {
		err = encoder.PushByte(0)
	} else if r.IsDeleted {
		if err = encoder.PushByte(1); err != nil {
			return err
		}
		err = encoder.Encode(r.AsDeleted)
	} else if r.IsGracePeriod {
		if err = encoder.PushByte(2); err != nil {
			return err
		}
		err = encoder.Encode(r.AsGracePeriodBlockNumber)
	}

	return
}

type HexHash [32]byte

func (h HexHash) String() string {
	return string(h[:])
}

// NewHexHash will create a new hash from a hex input (32 bytes)
func NewHexHash(hash string) (hexHash HexHash) {
	copy(hexHash[:], hash)
	return
}

type NodeContract struct {
	Node           types.U32  `json:"node_id"`
	DeploymentHash HexHash    `json:"deployment_hash"`
	DeploymentData string     `json:"deployment_data"`
	PublicIPsCount types.U32  `json:"public_ips_count"`
	PublicIPs      []PublicIP `json:"public_ips"`
}

type NameContract struct {
	Name string `json:"name"`
}

type RentContract struct {
	Node types.U32 `json:"node_id"`
}

type ContractType struct {
	IsNodeContract bool         `json:"is_node_contract"`
	NodeContract   NodeContract `json:"node_contract"`
	IsNameContract bool         `json:"is_name_contract"`
	NameContract   NameContract `json:"name_contract"`
	IsRentContract bool         `json:"is_rent_contract"`
	RentContract   RentContract `json:"rent_contract"`
}

// Decode implementation for the enum type
func (r *ContractType) Decode(decoder scale.Decoder) error {
	b, err := decoder.ReadOneByte()
	if err != nil {
		return err
	}

	switch b {
	case 0:
		r.IsNodeContract = true
		if err := decoder.Decode(&r.NodeContract); err != nil {
			return err
		}
	case 1:
		r.IsNameContract = true
		if err := decoder.Decode(&r.NameContract); err != nil {
			return err
		}
	case 2:
		r.IsRentContract = true
		if err := decoder.Decode(&r.RentContract); err != nil {
			return err
		}
	default:
		return fmt.Errorf("unknown contract type value")
	}

	return nil
}

// Encode implementation
func (r ContractType) Encode(encoder scale.Encoder) (err error) {
	if r.IsNodeContract {
		if err = encoder.PushByte(0); err != nil {
			return err
		}
		if err = encoder.Encode(r.NodeContract); err != nil {
			return err
		}
	} else if r.IsNameContract {
		if err = encoder.PushByte(1); err != nil {
			return err
		}

		if err = encoder.Encode(r.NameContract); err != nil {
			return err
		}
	} else if r.IsRentContract {
		if err = encoder.PushByte(2); err != nil {
			return err
		}

		if err = encoder.Encode(r.RentContract); err != nil {
			return err
		}
	}

	return
}

// Contract structure
type Contract struct {
	Versioned
	State              ContractState   `json:"state"`
	ContractID         types.U64       `json:"contract_id"`
	TwinID             types.U32       `json:"twin_id"`
	ContractType       ContractType    `json:"contract_type"`
	SolutionProviderID types.OptionU64 `json:"solution_provider_id"`
}

// CreateNodeContract creates a contract for deployment
func (s *Substrate) CreateNodeContract(identity Identity, node uint32, body string, hash string, publicIPs uint32, solutionProviderID *uint64) (uint64, error) {
	cl, meta, err := s.GetClient()
	if err != nil {
		return 0, err
	}

	var providerID types.OptionU64
	if solutionProviderID != nil {
		providerID = types.NewOptionU64(types.U64(*solutionProviderID))
	}

	h := NewHexHash(hash)
	c, err := types.NewCall(meta, "SmartContractModule.create_node_contract",
		node, h, body, publicIPs, providerID,
	)

	if err != nil {
		return 0, errors.Wrap(err, "failed to create call")
	}

	_, err = s.Call(cl, meta, identity, c)
	if err != nil {
		return 0, errors.Wrap(err, "failed to create contract")
	}

	return s.GetContractWithHash(node, h)
}

// BatchCreateContractData struct for batch create contract
type BatchCreateContractData struct {
	Node               uint32  `json:"node"`
	Body               string  `json:"body"`
	Hash               string  `json:"hash"`
	PublicIPs          uint32  `json:"public_ips"`
	SolutionProviderID *uint64 `json:"solution_provider_id"`

	// for name contracts. if set the contract is assumed to be a name contract
	// and other fields are ignored
	Name string `json:"name"`
}

// BatchAllCreateContract creates a batch of contracts for deployments atomically.
// transaction will rollback on error
func (s *Substrate) BatchAllCreateContract(identity Identity, contractData []BatchCreateContractData) ([]uint64, error) {
	contracts, _, err := s.batchCreateContract(identity, contractData, "Utility.batch_all")
	return contracts, err
}

// BatchCreateContract creates a batch of contracts for deployments non-atomically.
// on error returns the created contracts, the first failing contract index and the error message.
func (s *Substrate) BatchCreateContract(identity Identity, contractData []BatchCreateContractData) ([]uint64, *int, error) {
	return s.batchCreateContract(identity, contractData, "Utility.batch")
}

func (s *Substrate) batchCreateContract(identity Identity, contractData []BatchCreateContractData, batchType string) ([]uint64, *int, error) {
	cl, meta, err := s.GetClient()
	if err != nil {
		return nil, nil, err
	}

	calls := make([]types.Call, 0)
	hexHashes := make([]HexHash, len(contractData))
	for i, contract := range contractData {
		if contract.Name != "" {
			c, err := types.NewCall(meta, "SmartContractModule.create_name_contract",
				contract.Name,
			)
			if err != nil {
				return nil, nil, errors.Wrap(err, "failed to create call")
			}
			calls = append(calls, c)
			continue
		}

		var providerID types.OptionU64
		if contract.SolutionProviderID != nil {
			providerID = types.NewOptionU64(types.U64(*contract.SolutionProviderID))
		}

		h := NewHexHash(contract.Hash)
		c, err := types.NewCall(meta, "SmartContractModule.create_node_contract",
			contract.Node, h, contract.Body, contract.PublicIPs, providerID,
		)
		if err != nil {
			return nil, nil, errors.Wrap(err, "failed to create call")
		}

		calls = append(calls, c)
		hexHashes[i] = h
	}
	batchCall, err := types.NewCall(meta, batchType, calls)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to create batch call")
	}

	resp, err := s.Call(cl, meta, identity, batchCall)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to create contracts")
	}

	var failingContract *int
	if batchType == "Utility.batch" && resp.Events != nil && resp.Events.Utility_BatchInterrupted != nil {
		i := int(resp.Events.Utility_BatchInterrupted[0].Index)
		failingContract = &i
	}

	contracts := make([]uint64, 0)
	for i, contract := range contractData {
		if failingContract != nil && *failingContract == i {
			break
		}

		var contractID uint64
		var err error

		if contract.Name != "" {
			contractID, err = s.GetContractIDByNameRegistration(contract.Name)
		} else {
			contractID, err = s.GetContractWithHash(contract.Node, hexHashes[i])
		}

		if err != nil {
			failingContract = &i
			return nil, failingContract, err
		}
		contracts = append(contracts, contractID)
	}

	if failingContract != nil {
		err = errors.New("failed to create contracts")
	}
	return contracts, failingContract, err
}

// CreateNameContract creates a contract for deployment
func (s *Substrate) CreateNameContract(identity Identity, name string) (uint64, error) {
	cl, meta, err := s.GetClient()
	if err != nil {
		return 0, err
	}

	c, err := types.NewCall(meta, "SmartContractModule.create_name_contract",
		name,
	)

	if err != nil {
		return 0, errors.Wrap(err, "failed to create call")
	}

	_, err = s.Call(cl, meta, identity, c)
	if err != nil {
		return 0, errors.Wrap(err, "failed to create contract")
	}

	return s.GetContractIDByNameRegistration(name)
}

// CreateRentContract creates a rent contract on a node
func (s *Substrate) CreateRentContract(identity Identity, node uint32, solutionProviderID *uint64) (uint64, error) {
	cl, meta, err := s.GetClient()
	if err != nil {
		return 0, err
	}

	var providerID types.OptionU64
	if solutionProviderID != nil {
		providerID = types.NewOptionU64(types.U64(*solutionProviderID))
	}

	c, err := types.NewCall(meta, "SmartContractModule.create_rent_contract",
		node, providerID,
	)

	if err != nil {
		return 0, errors.Wrap(err, "failed to create call")
	}

	_, err = s.Call(cl, meta, identity, c)
	if err != nil {
		return 0, errors.Wrap(err, "failed to create rent contract")
	}

	return s.GetNodeRentContract(node)
}

// UpdateNodeContract updates existing contract
func (s *Substrate) UpdateNodeContract(identity Identity, contract uint64, body string, hash string) (uint64, error) {
	cl, meta, err := s.GetClient()
	if err != nil {
		return 0, err
	}

	h := NewHexHash(hash)
	c, err := types.NewCall(meta, "SmartContractModule.update_node_contract",
		contract, h, body,
	)

	if err != nil {
		return 0, errors.Wrap(err, "failed to create call")
	}

	_, err = s.Call(cl, meta, identity, c)
	if err != nil {
		return 0, errors.Wrap(err, "failed to update contract")
	}

	return contract, nil
}

// CancelContract creates a contract for deployment
func (s *Substrate) CancelContract(identity Identity, contract uint64) error {
	cl, meta, err := s.GetClient()
	if err != nil {
		return err
	}

	c, err := types.NewCall(meta, "SmartContractModule.cancel_contract", contract)
	if err != nil {
		return errors.Wrap(err, "failed to cancel call")
	}

	_, err = s.Call(cl, meta, identity, c)
	if err != nil {
		return errors.Wrap(err, "failed to cancel contract")
	}

	return nil
}

// BatchCancelContract cancels a batch of contracts
func (s *Substrate) BatchCancelContract(identity Identity, contracts []uint64) error {
	cl, meta, err := s.GetClient()
	if err != nil {
		return err
	}

	calls := make([]types.Call, 0)
	for _, id := range contracts {
		c, err := types.NewCall(meta, "SmartContractModule.cancel_contract", id)
		if err != nil {
			return err
		}
		calls = append(calls, c)
	}

	batchCall, err := types.NewCall(meta, "Utility.batch_all", calls)
	if err != nil {
		return errors.Wrap(err, "failed to create batch call")
	}

	_, err = s.Call(cl, meta, identity, batchCall)
	if err != nil {
		return errors.Wrap(err, "failed to cancel contracts")
	}

	return nil
}

// SetContractConsumption can only be called by the node that owns the contract to set the used
// resources associated with the node.
func (s *Substrate) SetContractConsumption(identity Identity, resources ...ContractResources) error {
	cl, meta, err := s.GetClient()
	if err != nil {
		return err
	}

	c, err := types.NewCall(meta, "SmartContractModule.report_contract_resources",
		resources,
	)

	if err != nil {
		return errors.Wrap(err, "failed to create call")
	}

	_, err = s.Call(cl, meta, identity, c)
	if err != nil {
		return errors.Wrap(err, "failed to set contract used resources")
	}

	return nil
}

// GetContract we should not have calls to create contract, instead only get
func (s *Substrate) GetContract(id uint64) (*Contract, error) {
	cl, meta, err := s.GetClient()
	if err != nil {
		return nil, err
	}

	bytes, err := Encode(id)
	if err != nil {
		return nil, errors.Wrap(err, "substrate: encoding error building query arguments")
	}

	key, err := types.CreateStorageKey(meta, "SmartContractModule", "Contracts", bytes, nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create substrate query key")
	}

	return s.getContract(cl, key)
}

// GetContractWithHash gets a contract given the node id and hash
func (s *Substrate) GetContractWithHash(node uint32, hash HexHash) (uint64, error) {
	cl, meta, err := s.GetClient()
	if err != nil {
		return 0, err
	}

	nodeBytes, err := Encode(node)
	if err != nil {
		return 0, err
	}
	hashBytes, err := Encode(hash)
	if err != nil {
		return 0, err
	}
	key, err := types.CreateStorageKey(meta, "SmartContractModule", "ContractIDByNodeIDAndHash", nodeBytes, hashBytes, nil)
	if err != nil {
		return 0, errors.Wrap(err, "failed to create substrate query key")
	}
	var contract types.U64
	_, err = cl.RPC.State.GetStorageLatest(key, &contract)
	if err != nil {
		return 0, errors.Wrap(err, "failed to lookup contracts")
	}

	if contract == 0 {
		return 0, errors.Wrap(ErrNotFound, "contract not found")
	}

	return uint64(contract), nil
}

// GetContractIDByNameRegistration gets a contract given the its name
func (s *Substrate) GetContractIDByNameRegistration(name string) (uint64, error) {
	cl, meta, err := s.GetClient()
	if err != nil {
		return 0, err
	}

	nameBytes, err := Encode(name)
	if err != nil {
		return 0, err
	}
	key, err := types.CreateStorageKey(meta, "SmartContractModule", "ContractIDByNameRegistration", nameBytes, nil)
	if err != nil {
		return 0, errors.Wrap(err, "failed to create substrate query key")
	}
	var contract types.U64
	_, err = cl.RPC.State.GetStorageLatest(key, &contract)
	if err != nil {
		return 0, errors.Wrap(err, "failed to lookup contracts")
	}

	if contract == 0 {
		return 0, errors.Wrap(ErrNotFound, "contract not found")
	}

	return uint64(contract), nil
}

// GetNodeContracts gets all contracts on a node (pk) in given state
func (s *Substrate) GetNodeContracts(node uint32) ([]types.U64, error) {
	cl, meta, err := s.GetClient()
	if err != nil {
		return nil, err
	}

	nodeBytes, err := Encode(node)
	if err != nil {
		return nil, err
	}

	key, err := types.CreateStorageKey(meta, "SmartContractModule", "ActiveNodeContracts", nodeBytes)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create substrate query key")
	}
	var contracts []types.U64
	_, err = cl.RPC.State.GetStorageLatest(key, &contracts)
	if err != nil {
		return nil, errors.Wrap(err, "failed to lookup contracts")
	}

	return contracts, nil
}

// GetNodeContracts gets all contracts on a node (pk) in given state
func (s *Substrate) GetNodeRentContract(node uint32) (uint64, error) {
	cl, meta, err := s.GetClient()
	if err != nil {
		return 0, err
	}

	nodeBytes, err := Encode(node)
	if err != nil {
		return 0, err
	}

	key, err := types.CreateStorageKey(meta, "SmartContractModule", "ActiveRentContractForNode", nodeBytes)
	if err != nil {
		return 0, errors.Wrap(err, "failed to create substrate query key")
	}

	raw, err := cl.RPC.State.GetStorageRawLatest(key)
	if err != nil {
		return 0, errors.Wrap(err, "failed to lookup contract")
	}

	if len(*raw) == 0 {
		return 0, errors.Wrap(ErrNotFound, "contract not found")
	}

	var contract uint64
	err = Decode(*raw, &contract)
	return contract, err
}

func (s *Substrate) getContract(cl Conn, key types.StorageKey) (*Contract, error) {
	raw, err := cl.RPC.State.GetStorageRawLatest(key)
	if err != nil {
		return nil, errors.Wrap(err, "failed to lookup contract")
	}

	if len(*raw) == 0 {
		return nil, errors.Wrap(ErrNotFound, "contract not found")
	}

	var contract Contract
	if err := Decode(*raw, &contract); err != nil {
		return nil, errors.Wrap(err, "failed to load object")
	}

	return &contract, nil
}

// Consumption structure
type NruConsumption struct {
	ContractID types.U64 `json:"contract_id"`
	Timestamp  types.U64 `json:"timestamp"`
	Window     types.U64 `json:"window"`
	NRU        types.U64 `json:"nru"`
}

// Consumption structure
type Consumption struct {
	ContractID types.U64 `json:"contract_id"`
	Timestamp  types.U64 `json:"timestamp"`
	CRU        types.U64 `json:"cru"`
	SRU        types.U64 `json:"sru"`
	HRU        types.U64 `json:"hru"`
	MRU        types.U64 `json:"mru"`
	NRU        types.U64 `json:"nru"`
}

// IsEmpty true if consumption is zero
func (s *NruConsumption) IsEmpty() bool {
	return s.NRU == 0
}

// Report send a capacity report to substrate
func (s *Substrate) Report(identity Identity, consumptions []NruConsumption) (types.Hash, error) {
	cl, meta, err := s.GetClient()
	if err != nil {
		return types.Hash{}, err
	}

	c, err := types.NewCall(meta, "SmartContractModule.add_nru_reports", consumptions)
	if err != nil {
		return types.Hash{}, errors.Wrap(err, "failed to create call")
	}

	callResponse, err := s.Call(cl, meta, identity, c)
	if err != nil {
		return types.Hash{}, errors.Wrap(err, "failed to create report")
	}

	return callResponse.Hash, nil
}

type ContractResources struct {
	ContractID types.U64 `json:"contract_id"`
	Used       Resources `json:"resources"`
}
