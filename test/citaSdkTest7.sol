pragma solidity >=0.6.3;

contract citaSdkTest7 {
    int a;

    event Add(int indexed a, string indexed b);
    event AddB(int indexed a, string indexed b);

    function add() public {
        emit Add(a, "hello");
        emit AddB(a, "world");
        a++;
    }

    function get() public view returns(int) {
        return a;
    }
}
