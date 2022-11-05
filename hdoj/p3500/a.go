package p3500

// 3555
//typedef long long ll;
//
//// p3555
//string v3555;
//ll dp3555[30][3];
//
//// 0表示不存在 1 表示上一位是4， 2表示存在
//ll f3555(int i, int pre, bool limit) {
//    if (i == v3555.length()) {
//        if (pre == 2) {
//            return 1;
//        }
//        return 0;
//    }
//
//    if (!limit) {
//        if (dp3555[i][pre] >= 0) {
//            return dp3555[i][pre];
//        }
//    }
//
//    ll res = 0;
//    int up = 9;
//    if (limit) {
//        up = v3555[i] - '0';
//    }
//    for (int j = 0; j <= up; j++)
//        if (pre == 2 || (pre == 1 && j == 9))
//            res += f3555(i + 1, 2, limit && j == up);
//        else if (j == 4)
//            res += f3555(i + 1, 1, limit && j == up);
//        else
//            res += f3555(i + 1, 0, limit && j == up);
//
//    if (!limit) {
//        dp3555[i][pre] = res;
//    }
//    return res;
//}
//
//void P3555() {
//    int T;
//    cin >> T;
//    while (T--) {
//        cin >> v3555;
//        memset(dp3555, -1, sizeof(dp3555));
//        cout << f3555(0, 0, true) << endl;
//    }
//}
//
