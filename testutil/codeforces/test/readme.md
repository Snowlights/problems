# Codeforces测试用力模板

### AssertEqualCase

| 参数 | 类型              |是否必须| 描述                             |
|---|-----------------|---|--------------------------------|
| t | testing.T point |是| 测试用例的参数                        |
| f | interface{}     |是| 测试的函数                          |
| rawText  | string          |是| 测试用例数据                         |
| targetCaseNum  | int             |是| 指定第几个测试用例开始, 可用{-num}标识倒数第几个用例 |

#### rawExamples
每个测试用例使用字符串表示

```
	示例: inputCopy后为输入,outputCopy后为输出,直接从web端粘贴,需要换行
		rawText := `
inputCopy
5
3
1 2 -3
5
1 0 0 -1 -1
6
2 -4 3 -5 4 1
5
1 -1 1 -1 1
7
0 0 0 0 0 0 0
outputCopy
3
2
6
1
0`
```

### AssertEqualFileCaseWithName

| 参数            | 类型              |是否必须| 描述                             |
|---------------|-----------------|---|--------------------------------|
| t             | testing.T point |是| 测试用例的参数                        |
| f             | interface{}     |是| 测试的函数                          |
| dir           | string          |是| 测试用例所在目录                       |
| in            | string          |是| 测试用例数据编号,正则方式:in*.txt          |
| ans           | string          |是| 测试用例数据结果编号,正则方式:ans*.txt       |
| targetCaseNum | int             |是| 指定第几个测试用例开始, 可用{-num}标识倒数第几个用例 |

```
	codeforces.AssertEqualFileCaseWithName(t, CF1775E, dir, "in*.txt", "ans*.txt", 0)
```