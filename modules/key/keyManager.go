package key

// type keyManager struct {
// 	DB store.DB
// }

// func NewKeyManager(keyDAO store.KeyDAO, algo string) store.KeyManager {
// 	return keyManager{
// 		keyDAO: keyDAO,
// 		algo:   algo,
// 	}
// }

// func (k keyManager) Sign(name, password string, data []byte) ([]byte, *sm2.PublicKey, error) {
// 	info, err := k.keyDAO.Read(name, password)
// 	if err != nil {
// 		return nil, nil, fmt.Errorf("name %s not exist", name)
// 	}

// 	km, err := crypto.NewPrivateKeyManager([]byte(info.PrivKeyArmor), string(info.Algo))
// 	if err != nil {
// 		return nil, nil, fmt.Errorf("name %s not exist", name)
// 	}

// 	signByte, err := km.Sign(data)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	return signByte, FromTmPubKey(info.Algo, km.ExportPubKey()), nil
// }

// func (k keyManager) Recover(name, password, mnemonic string) (string, error) {
// 	if k.keyDAO.Has(name) {
// 		return "", fmt.Errorf("name %s has existed", name)
// 	}

// 	km, err := crypto.NewMnemonicKeyManager(mnemonic, k.algo)
// 	if err != nil {
// 		return "", err
// 	}

// 	_, priv := km.Generate()

// 	pubKey := km.ExportPubKey()
// 	address := types.AccAddress(pubKey.Address().Bytes()).String()

// 	info := store.KeyInfo{
// 		Name:         name,
// 		PubKey:       cryptoamino.MarshalPubkey(pubKey),
// 		PrivKeyArmor: string(cryptoamino.MarshalPrivKey(priv)),
// 		Algo:         k.algo,
// 	}

// 	err = k.keyDAO.Write(name, password, info)
// 	if err != nil {
// 		return "", err
// 	}

// 	return address, nil
// }

// func (k keyManager) Export(name, password string) (armor string, err error) {
// 	info, err := k.keyDAO.Read(name, password)
// 	if err != nil {
// 		return armor, fmt.Errorf("name %s not exist", name)
// 	}

// 	km, err := crypto.NewPrivateKeyManager([]byte(info.PrivKeyArmor), info.Algo)
// 	if err != nil {
// 		return "", err
// 	}

// 	return km.ExportPrivKey(password)
// }

// func (k keyManager) Delete(name, password string) error {
// 	return k.keyDAO.Delete(name, password)
// }

// func FromTmPubKey(algo string, pubKey tmcrypto.PubKey) codectypes.PubKey {
// 	var pubkey codectypes.PubKey
// 	pubkeyBytes := pubKey.Bytes()
// 	switch algo {
// 	case "sm2":
// 		pubkey = &sm2.PubKey{Key: pubkeyBytes}
// 	case "secp256k1":
// 		pubkey = &secp256k1.PubKey{Key: pubkeyBytes}
// 	}
// 	return pubkey
// }
