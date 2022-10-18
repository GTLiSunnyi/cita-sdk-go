pragma solidity >=0.6.3;
pragma experimental ABIEncoderV2;

contract citaSdkTest15 {
    struct Sdk {
        int num;
        string str;
    }

    int a;
    Sdk b = Sdk(1, "test");

    event Add(int A, string B);
    event AddB(int A, string B);

    function add() public {
        emit Add(a, "hello");
        emit AddB(a, "world");
        a++;
    }

    function get() public view returns(Sdk memory) {
        return b;
    }
}
