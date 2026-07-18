import {MyMathLib} from "./library/MyMathLib.sol";

contract MyContract {

    using MyMathLib for int;

    // 传统方式调用库函数
    function CalculateTradition(int x, int y) public pure returns (int) {
        return MyMathLib.add(x, y);
    }

    function CalculateUsingFor(int x, int y) public pure returns (int) {
        return x.add(y);
    }
}