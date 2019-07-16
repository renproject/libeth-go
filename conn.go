package libeth

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

// RenNetwork identifies which set of contract to use.
type RenNetwork string

// Ren networks.
const (
	Mainnet  = RenNetwork("mainnet")
	Ropsten  = RenNetwork("ropsten")
	Testnet  = RenNetwork("testnet")
	Devnet   = RenNetwork("devnet")
	Localnet = RenNetwork("localnet")
)

// ErrCannotConvertToBigInt is returned when string cannot be parsed into a
// big.Int format.
var ErrCannotConvertToBigInt = errors.New("cannot convert hex string to int: invalid format")

// Client will have a connection to an ethereum client (specified by the url)
type Client struct {
	renNetwork  RenNetwork
	ethClient   *ethclient.Client
	ethWSClient *ethclient.Client
	addrBook    AddressBook
	contracts   Contracts
	url         string
	apiKey      string
}

// NewMercuryClient creates a new infura client
func NewMercuryClient(renNetwork RenNetwork, tag string) (Client, error) {
	if tag != "" {
		tag = fmt.Sprintf("?tag=%s", tag)
	}

	var network string
	switch renNetwork {
	case Mainnet:
		network = "eth"
	case Testnet:
		network = "eth-kovan"
	case Devnet:
		network = "eth-kovan"
	case Localnet:
		network = "eth-kovan"
	case Ropsten:
		network = "eth-ropsten"
	default:
		return Client{}, fmt.Errorf("unsupported network: %s", renNetwork)
	}

	return Connect(renNetwork, fmt.Sprintf("http://139.59.221.34/%s%s", network, tag))
}

// NewInfuraClient creates a new infura client
func NewInfuraClient(renNetwork RenNetwork, apiKey string) (Client, error) {
	var network string
	switch renNetwork {
	case Mainnet:
		network = "mainnet"
	case Testnet:
		network = "kovan"
	case Devnet:
		network = "kovan"
	case Localnet:
		network = "kovan"
	case Ropsten:
		network = "ropsten"
	default:
		return Client{}, fmt.Errorf("unsupported network: %s", renNetwork)
	}
	return Connect(renNetwork, fmt.Sprintf("https://%s.infura.io/v3/%s", network, apiKey))
}

// NewFullInfuraClient creates a new infura client
func NewFullInfuraClient(renNetwork RenNetwork, apiKey string) (Client, error) {
	var network string
	switch renNetwork {
	case Mainnet:
		network = "mainnet"
	case Testnet:
		network = "kovan"
	case Devnet:
		network = "kovan"
	case Localnet:
		network = "kovan"
	case Ropsten:
		network = "ropsten"
	default:
		return Client{}, fmt.Errorf("unsupported network: %s", renNetwork)
	}
	return NewClient(
		renNetwork,
		fmt.Sprintf("https://%s.infura.io/v3/%s", network, apiKey),
		fmt.Sprintf("wss://%s.infura.io/ws/v3/%s", network, apiKey),
	)
}

// NewClient creates a new client
func NewClient(renNetwork RenNetwork, URL, wsURL string) (Client, error) {
	client, err := Connect(renNetwork, URL)
	if err != nil {
		return Client{}, err
	}

	wsClient, err := ethclient.Dial(wsURL)
	if err != nil {
		return Client{}, err
	}

	client.ethWSClient = wsClient
	return client, nil
}

// Deprecated
// Connect to an infura network (Supported networks: mainnet and kovan).
func Connect(renNetwork RenNetwork, url string) (Client, error) {

	ethClient, err := ethclient.Dial(url)
	if err != nil {
		return Client{}, err
	}

	return Client{
		renNetwork: renNetwork,
		ethClient:  ethClient,
		addrBook:   NetworkAddressBook(renNetwork),
		url:        url,
		apiKey:     "R8F2CVXTVSCIDD2IQ2ZQP9P6VZADUWHDHN",
	}, nil
}

// WriteAddress to the address book, overwrite if already exists
func (client *Client) WriteAddress(key string, address common.Address) {
	client.addrBook[key] = address
}

// FormatTransactionView returns the formatted string with the URL at which the
// transaction can be viewed.
func (client *Client) FormatTransactionView(msg, txHash string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	netID, err := client.ethClient.NetworkID(ctx)
	if err != nil {
		return "", err
	}
	switch netID.Int64() {
	case 1:
		return fmt.Sprintf("%s, the transaction can be viewed at https://etherscan.io/tx/%s", msg, txHash), nil
	case 3:
		return fmt.Sprintf("%s, the transaction can be viewed at https://ropsten.etherscan.io/tx/%s", msg, txHash), nil
	case 42:
		return fmt.Sprintf("%s, the transaction can be viewed at https://kovan.etherscan.io/tx/%s", msg, txHash), nil
	default:
		return "", fmt.Errorf("unknown network id : %d", netID.Int64())
	}
}

// Network returns the network of the underlying client.
func (client *Client) Network() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	netID, err := client.ethClient.NetworkID(ctx)
	if err != nil {
		return "", err
	}
	switch netID.Int64() {
	case 1:
		return "mainnet", nil
	case 3:
		return "ropsten", nil
	case 42:
		return "kovan", nil
	default:
		return "", fmt.Errorf("unknown network id : %d", netID.Int64())
	}
}

// ReadAddress from the address book, return an error if the address does not
// exist
func (client *Client) ReadAddress(key string) (common.Address, error) {
	if address, ok := client.addrBook[key]; ok {
		return address, nil
	}
	return common.Address{}, ErrAddressNotFound
}

// WaitMined waits for tx to be mined on the blockchain.
// It stops waiting when the context is canceled.
func (client *Client) WaitMined(ctx context.Context, tx *types.Transaction) (*types.Receipt, error) {
	return bind.WaitMined(ctx, client.ethClient, tx)
}

// Get will perform a read-only transaction on the ethereum blockchain.
func (client *Client) Get(ctx context.Context, f func() error) (err error) {

	sleepDurationMs := time.Duration(1000)

	// Keep retrying until the read-only transaction succeeds or until context
	// times out
	for {
		select {
		case <-ctx.Done():
			if err == nil {
				return ctx.Err()
			}
			return
		default:
		}

		if err = f(); err == nil {
			return
		}

		// If transaction errors, wait for sometime before retrying
		select {
		case <-ctx.Done():
			if err == nil {
				return ctx.Err()
			}
			return
		case <-time.After(sleepDurationMs * time.Millisecond):
		}

		// Increase delay for next round but saturate at 30s
		sleepDurationMs = time.Duration(float64(sleepDurationMs) * 1.6)
		if sleepDurationMs > 30000 {
			sleepDurationMs = 30000
		}
	}
}

// Call a function on a contract with the given parameters
func (client Client) Call(ctx context.Context, address, fnName string, params ...interface{}) ([]interface{}, error) {
	net, err := client.ethClient.NetworkID(ctx)
	if err != nil {
		return nil, err
	}

	contractAbi, err := getABI(net.Int64(), address, client.apiKey)
	if err != nil {
		return nil, err
	}
	parsed, err := abi.JSON(strings.NewReader(contractAbi))
	if err != nil {
		return nil, err
	}

	data, err := parsed.Pack(fnName, params...)
	if err != nil {
		return nil, err
	}

	contractAddr := common.HexToAddress(address)
	sleepDurationMs := time.Duration(1000)
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(sleepDurationMs * time.Millisecond):
			resp, err := client.ethClient.CallContract(ctx, ethereum.CallMsg{To: &contractAddr, Data: data}, nil)
			if err != nil || len(resp) == 0 {
				break
			}
			return parsed.Methods[fnName].Outputs.UnpackValues(resp)
		}
		sleepDurationMs = time.Duration(float64(sleepDurationMs) * 1.6)
		if sleepDurationMs > 30000 {
			sleepDurationMs = 30000
		}
	}
}

// Query a function on a contract with the given parameters
func (client Client) Query(ctx context.Context, address, fnName string, params ...[]byte) ([]interface{}, error) {
	net, err := client.ethClient.NetworkID(ctx)
	if err != nil {
		return nil, err
	}
	contractAddr := common.HexToAddress(address)
	contractAbi, ok := ContractABIs[client.contracts[contractAddr]]
	if !ok {
		contractAbi, err = getABI(net.Int64(), address, client.apiKey)
		if err != nil {
			return nil, err
		}
	}

	parsed, err := abi.JSON(strings.NewReader(contractAbi))
	if err != nil {
		return nil, err
	}

	arguments := []byte{}
	for _, param := range params {
		arguments = append(arguments, padParam(param)...)
	}

	data := append(parsed.Methods[fnName].Id(), arguments...)
	sleepDurationMs := time.Duration(1000)
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(sleepDurationMs * time.Millisecond):
			resp, err := client.ethClient.CallContract(ctx, ethereum.CallMsg{To: &contractAddr, Data: data}, nil)
			if err != nil || len(resp) == 0 {
				break
			}
			return parsed.Methods[fnName].Outputs.UnpackValues(resp)
		}
		sleepDurationMs = time.Duration(float64(sleepDurationMs) * 1.6)
		if sleepDurationMs > 30000 {
			sleepDurationMs = 30000
		}
	}
}

// BalanceOf returns the ethereum balance of the addr passed.
func (client *Client) BalanceOf(ctx context.Context, addr common.Address) (val *big.Int, err error) {
	err = client.Get(ctx, func() (err error) {
		val, err = client.ethClient.BalanceAt(ctx, addr, nil)
		return
	})
	return
}

// EthClient returns the ethereum client connection.
func (client *Client) EthClient() *ethclient.Client {
	return client.ethClient
}

// Relay the following transaction.
func (client *Client) Relay(address, fnName string, params ...[]byte) (string, error) {
	data := make([]string, len(params))
	for i := range data {
		data[i] = hex.EncodeToString(params[i])
	}

	req := struct {
		Address string   `json:"address"`
		FnName  string   `json:"fnName"`
		Data    []string `json:"data"`
	}{address, fnName, data}

	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(req); err != nil {
		return "", err
	}

	resp, err := http.Post(fmt.Sprintf("%s/relay", client.url), "encoding/json", buf)
	if err != nil {
		return "", err
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		errObj := struct {
			Error string `json:"error"`
		}{}
		if err := json.Unmarshal(respBytes, &errObj); err != nil {
			return "", err
		}
		return "", fmt.Errorf(errObj.Error)
	}

	respObj := struct {
		TxHash string `json:"txHash"`
	}{}
	if err := json.Unmarshal(respBytes, &respObj); err != nil {
		return "", err
	}

	return respObj.TxHash, nil
}

// EthWSClient returns the ethereum ws client connection.
func (client *Client) EthWSClient() *ethclient.Client {
	return client.ethWSClient
}

// TxBlockNumber retrieves tx's block number using the tx hash.
func (client *Client) TxBlockNumber(ctx context.Context, hash string) (*big.Int, error) {

	type Result struct {
		BlockNumber string `json:"blockNumber,omitempty"`
	}
	type JSONResponse struct {
		Result Result `json:"result,omitempty"`
	}
	var data JSONResponse

	var jsonStr = `{"jsonrpc":"2.0","method":"eth_getTransactionByHash",` +
		`"params":["` + hash + `"],"id":1}`

	// Keep retrying until a block number is returned or until context times out
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		response, err := sendInfuraRequest(ctx, client.url, jsonStr)
		if err != nil {
			continue
		}
		err = json.Unmarshal(response, &data)

		if err != nil || data.Result == (Result{}) || data.Result.BlockNumber == "" {
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case <-time.After(5 * time.Millisecond):
			}
			continue
		}
		break
	}

	return hexToBigInt(data.Result.BlockNumber)
}

// CurrentBlockNumber will retrieve the current block that is confirmed by
// infura.
func (client *Client) CurrentBlockNumber(ctx context.Context) (*big.Int, error) {

	type Result struct {
		Number string `json:"number,omitempty"`
	}
	type JSONResponse struct {
		Result Result `json:"result,omitempty"`
	}
	var data JSONResponse

	var jsonStr = `{"jsonrpc":"2.0","method":"eth_getBlockByNumber",` +
		`"params":["latest", false],"id":1}`

	// Keep retrying until a block number is returned or until context times out
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		response, err := sendInfuraRequest(ctx, client.url, jsonStr)
		if err != nil {
			continue
		}
		err = json.Unmarshal(response, &data)

		if err != nil || data.Result == (Result{}) || data.Result.Number == "" {
			select {
			case <-ctx.Done():
				return nil, ctx.Err()
			case <-time.After(5 * time.Millisecond):
			}
			continue
		}
		break
	}

	return hexToBigInt(data.Result.Number)
}

// hexToBigInt will convert a hex value in string format to the corresponding
// big.Int value. For example : "0xFD6CE" will return big.Int(1038030).
func hexToBigInt(hex string) (*big.Int, error) {
	bigInt := big.NewInt(0)
	bigIntStr := hex[2:]
	bigInt, ok := bigInt.SetString(bigIntStr, 16)
	if !ok {
		return bigInt, ErrCannotConvertToBigInt
	}
	return bigInt, nil
}

// sendInfuraRequest will send a request to infura and return the unmarshalled data
// back to the caller. It will retry until a valid response is returned, or
// until the context times out.
func sendInfuraRequest(ctx context.Context, url string, request string) (body []byte, err error) {

	sleepDurationMs := time.Duration(1000)

	// Retry until a valid response is returned or until context times out
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		if body, err = func() ([]byte, error) {
			// Create a new http POST request
			req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(request)))
			if err != nil {
				return nil, err
			}

			// Send http POST request
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				return nil, err
			}

			// Decode response body
			return func() ([]byte, error) {
				defer resp.Body.Close()

				// Check status
				if resp.StatusCode != http.StatusOK {
					return nil, fmt.Errorf("unexpected status %v", resp.StatusCode)
				}
				// Check body
				if resp.Body != nil {
					return ioutil.ReadAll(resp.Body)
				}
				return nil, fmt.Errorf("response body is nil")
			}()
		}(); err == nil {
			break
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(sleepDurationMs * time.Millisecond):

		}

		// Increase delay for next round but saturate at 30s
		sleepDurationMs = time.Duration(float64(sleepDurationMs) * 1.6)
		if sleepDurationMs > 30000 {
			sleepDurationMs = 30000
		}
	}
	return
}

func getABI(net int64, address, apiKey string) (string, error) {
	network := ""
	switch net {
	case 1:
		network = "api"
	case 3:
		network = "api-ropsten"
	case 42:
		network = "api-kovan"
	default:
		return "", fmt.Errorf("unsupported network on etherscan")
	}

	value := struct {
		ABI string `json:"result"`
	}{}

	resp, err := http.Get(fmt.Sprintf("https://%s.etherscan.io/api?module=contract&action=getabi&address=%s&apikey=%s", network, address, apiKey))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	return value.ABI, json.NewDecoder(resp.Body).Decode(&value)
}

func padParam(param []byte) []byte {
	if len(param) > 32 {
		return param[:32]
	}
	padding := make([]byte, 32-len(param))
	return append(padding, param...)
}
