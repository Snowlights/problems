import os
import sys
from functools import lru_cache

if __name__ == '__main__':
    n, m = map(int, input().split())
    mod = 10 ** 9 + 7


    @lru_cache(None)
    def dfs(wine, store, flower: int) -> int:
        if wine > flower:
            return 0
        if store == 0 or wine == 0 or flower == 0:
            return wine == flower and wine > 0
        return dfs(wine * 2, store - 1, flower) + dfs(wine - 1, store, flower - 1)


    print(dfs(2, n, m) % mod)
