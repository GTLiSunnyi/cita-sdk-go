pragma solidity >=0.6.3;
pragma experimental ABIEncoderV2;

contract citaSdkTest23 {
    struct Sdk {
        int num;
        string str;
    }

    int a;
    Sdk b = Sdk(1, "test");
    string[] c;

    event Add(int A, string B);
    event AddB(int A, string B);

    function addc() public {
        c = new string[](2);
        c.push("test");
    }

    function add() public {
        emit Add(a, "hello");
        emit AddB(a, "world");
        a++;
    }

    function getStruct() public view returns(Sdk memory) {
        return b;
    }

    function getArr() public view returns(string[] memory) {
        return c;
    }
}
