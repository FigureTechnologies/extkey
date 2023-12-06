package keys

type SomeKey struct {
	Seed      string        `json:"seed,omitempty" yaml:"seed,omitempty"`
	Mnemonic  string        `json:"mnemonic,omitempty" yaml:"mnemonic,omitempty"`
	Hrp       string        `json:"hrp,omitempty" yaml:"hrp,omitempty"`
	RootKey   *ExtKeyData   `json:"rootKey,omitempty" yaml:"rootKey,omitempty"`
	ChildKeys []*ExtKeyData `json:"childKey,omitempty" yaml:"childKey,omitempty"`
}

type ExtKeyData struct {
	Address     string        `json:"address" yaml:"address"`
	Path        string        `json:"path,omitempty" yaml:"path,omitempty"`
	PrivateKey  *InnerKeyData `json:"privateKey" yaml:"privateKey"`
	PublicKey   *InnerKeyData `json:"publicKey" yaml:"publicKey"`
	Depth       byte          `json:"depth" yaml:"depth"`
	DepthLoc    string        `json:"depthLoc" yaml:"depthLoc"`
	Chaincode   string        `json:"chaincode" yaml:"chaincode"`
	Fingerprint string        `json:"fingerprint" yaml:"fingerprint"`
}

type InnerKeyData struct {
	Base58    string `json:"base58" yaml:"base58"`
	Base64    string `json:"base64" yaml:"base64"`
	Bytes     string `json:"bytes" yaml:"bytes,flow"`
	BigInt    string `json:"bigInt" yaml:"bigInt"`
	Proto     string `json:"proto" yaml:"proto"`
	ProtoJSON string `json:"protoJson" yaml:"protoJson"`
}
