package consensus

import (
	"bytes"
	"encoding/hex"
	. "github.com/yoss22/bulletproofs"
	"testing"
)

func TestTransactionDeserialization(t *testing.T) {
	transactionMsg, _ := hex.DecodeString(
		"d29efa3aa282679fba8353ac370bc9994394d1b1964fb4830c58bf4a402d4f800000000000000001000000000000000200000000000000010009de9aceb09ff7cc422a3704dddf9373a7bdcc8805b2f81a9ee05786f49238a7660008bc127c31911faf56ed4b3bcc9819dc34b47e104de991ce32519ec333c6638e0600000000000002a34199a825fc69cd4030d11924f3011bba3322fecf866bfd44b05d84ddf4c5fada39284bb90e3594386bb8b116825f90ab2c9ea7e3bbec047a3c0b2586bb58dc520f9687a3f713f6921109ea6474f8219fdbd2aa2b37620b5f2b1da736f6ee43e9d8216ebbc921a8355db6952210d625a89155d60f32ed3deb8dabbc8a8f059bbfa1f1803a367a404121824db8111202311641d61b5c03e0c694b326e4124c3dc9c48a7d9b745bc0e1f1db6cf746d7d183300e3f212d6e4bc4c69b4275644993bd1766d732862ff77ebb1667a8ff9e318338f91cbb0b494e5b934721f24818acfaf27ab2cf5ca78aefa6510bb032666e198000f7b81499fc50fa4ba40b8e866721fa69e08940968769459ae60c22bef2c1ffb6f472243a04ffa049af2cb0ab206349c03cad3c3e40d46f0a5d1999825df1d5af75caa726ec78eb312a716468e8e455071c109a01086c5531fce7d1dad145ceba46a55a32096c23ab867a4842e650f8630027a5206192d55bfae463a340cd95143139af48fa175f158ee0be66715409c1d2db2630a2e85e414cc123bb4a67dfd0b9ea05e22dea7b2fc5f2c28462cb1a74f1e63a513826d9403f38067a68cf5e3172b4023c541d97480ed2421179d7e2abee1d5e11524971adc95682845bea7303427423ec84adf7c4bf99a1f03d6fc02ca6a6328d81c4b23c1c7230a9e0d42b000b885dbd0681ada4ddef22386c97ee1a9dad87c39234ceefc3b5ec41c744f546924a5de250c7f5e1bfe8ee631dc049ec748c0f3702d8b2f2e650af38c5287a64ebe51b43c527978a16226a0eb4b632a11ec51040627e92ce2529c30e3e1b34170d490d5f50abe57fed3e6c2fcde6c39115e002edb5eae0b9a006434ca8e8985e4bf93cb20ee52342b4698865d870258e8ac9a15c55853eabef06f32b12ba28bbec31ecbac8a8bfe12e5c0ccfa21c1a5120000832a3990cd8a497ad280394afeaa5fcbdf02d6e8c86eb7fc47ba6bb25cd8973fd00000000000002a38bfaaaeebbba7ff4b614c75390729666e2ce1cfa0d5fa9d33c7e8327dd468711b1cd2e6a72108735e183ab232114969adf1bda21b78d524ac8d76ab1f68b8c5700c09e2ec7874e6948d6367f8af295b54806dcfe46021ed115f74ba0679509e1b650def8083790f8e26745fa3141d69fd350c7c726bc9453d3e1598cae27f8c131f233de12d947bb4b0c0b0d12fb147d4780eca856b380f9a8a952031e201d12bfc7815082fa7710a0e57feb0a514b8dddcf98c277cc16b6c5347805afc095ad304dbada87330648c6ef0a18f21ad6dde0460c416811f1276e8ea335491b8e297be2f27e8e827f112ef66a3f86e978c39ea770c084d4dbaa96dcfab6966c822a3ef3af42435602689ed3c3b4578097e06a8f609b442ee309a5e5348dcbad74086c9ebb72fa6588db3316dd9e41262cd807b8565d3b1e1a71e7400e90aeffb841ace29355ebd07521672f8b8c1d32f055b8d794bc8ad46150efb595e171cad20f0d0700594fa5d850eb688f4b871c43c8ef039679b2b282ab968a91e4e78d5f223a4acac2484240496912b875c87c5fca8490a78fe78fc18c8f17b87ce8300d8a7360b3e2878302c57747ceb107556640c620c64b196b5a94079b188086d456554fae8cc33e0dab35618e3a4c7645fd16112443f649f2e0b679d63dea8d5da2ed783af2d068b1d830c821af27b954daab77405736258d25c005b9b98634a816d266264d1097824f479ab9addec644e744c579a8b4c0b814ee147a241098de6fb739c88e32828130b62435d7b7836cbd213f7b364e156ee42ed69aa8d83dd5d97af3e4ab2af17e14fab5276bbc0c1243996fed445648a882e5e1a53509ecd5225bbce19f82937ff680f476619242096cff5cf2e19712f62bdeac1488824db6945583bfea1762f26f76dccbd9b1c970e242632480b1abbdb709ce1fbfaf7a092051bdb3dc0000000000007a1200000000000001011908b3e6f62be2a299e1c96627822f1228aeb977a79a7074872e91cb6d0c2f239ab9752e5a56dd5d0d16bfa7c19defe154b1185b3d40acb34d75e73d3de288c5c9dc6326369b0d216ac21ef5e2240f3578e7503aa71a4405ce8f2ee6a0696ed98de7")

	r := bytes.NewReader(transactionMsg)

	tx := &Transaction{}
	if err := tx.Read(r); err != nil {
		t.Errorf("failed to parse transaction: %v", err)
	}
}

func TestTransactionOutputSerialize(t *testing.T) {
	serialized, _ := hex.DecodeString(
		"0832a3990cd8a497ad280394afeaa5fcbdf02d6e8c86eb7fc47ba6bb25cd8973fd")

	Commit := new(Point)
	if err := Commit.Read(bytes.NewReader(serialized)); err != nil {
		t.Errorf("failed to read point")
	}

	actual := Commit.Bytes()
	if !bytes.Equal(actual, serialized) {
		t.Errorf("TestTransactionOutputSerialize wrong, got %x, expected %x",
			actual, serialized)
	}
}
