pragma solidity >=0.6.3;

contract citaSdkTest8 {
    int a;

    event Add(int A, string B);
    event AddB(int A, string B);

    function add() public {
        emit Add(a, "hello");
        emit AddB(a, "world");
        a++;
    }

    function get() public view returns(int) {
        return a;
    }
}
