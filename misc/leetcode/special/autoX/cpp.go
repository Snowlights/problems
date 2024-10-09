package autoX

//// 1
//class Solution {
//public:
//int getLengthOfWaterfallFlow(int num, vector<int>& block) {
//        vector<int> a(num,0);
//        for(auto v:block){
//            int chs=-1;
//            for(int i=0;i<num;++i){
//                if(chs==-1||a[i]<a[chs])
//                    chs=i;
//            }
//            a[chs]+=v;
//        }
//        int ans=0;
//        for(auto v:a)
//            ans=max(ans,v);
//        return ans;
//    }
//};
//
//// 2
//
//class Solution {
//public:
//vector<double> honeyQuotes(vector<vector<int>>& handle) {
//    vector<double> ans;
//    map<int, int> s;
//    double total = 0, count = 0;
//
//    for (auto x : handle) {
//        switch (x[0]) {
//            case 1:
//                s[x[1]]++;
//                total += x[1];
//                count++;
//                break;
//            case 2:
//                s[x[1]]--;
//                total -= x[1];
//                count--;
//                break;
//            case 3:
//                if (count > 0) {
//                    ans.push_back(total/count);
//                } else {
//                    ans.push_back(-1);
//                }
//                break;
//            case 4:
//                double avg = total / count;
//                if (count > 0) {
//                    double res = 0;
//                    for (int i = 0; i <= 100; ++i) {
//                        if (s[i] > 0) {
//                            res += (avg - i) * (avg - i) * s[i];
//                        }
//                    }
//                    ans.push_back(res/count);
//                } else {
//                    ans.push_back(-1);
//                }
//                break;
//        }
//    }
//    return ans;
//}
//};
//
//// 3
//typedef long long ll;
//class Solution {
//public:
//    long long minCostToTravelOnDays(vector<int>& days, vector<vector<int>>& tickets) {
//        vector<ll> dp(days.size()+1, LONG_LONG_MAX);
//        dp[0] = 0;
//        for (int i = 0; i < days.size(); ++i) {
//            for (auto& t : tickets) {
//                int len = t[0], cost = t[1];
//                int j = lower_bound(days.begin(), days.end(), days[i] + len) - days.begin();
//                dp[j] = min(dp[j], dp[i]+cost);
//            }
//        }
//        return dp[days.size()];
//    }
//};
//
//// 4
//#define int long long
//const int maxn=2e3+10;
//const double eps=1e-8;
//struct dsu{
//	int fa[maxn],siz[maxn];
//	inline void init(int k){for(int i=0;i<k;++i)fa[i]=i,siz[i]=1;}
//	int find(int k){return k==fa[k]?k:fa[k]=find(fa[k]);}
//	inline bool merge(int x,int y){
//		x=find(x),y=find(y);
//		if(x!=y){
//			if(siz[x]>siz[y])swap(x,y);
//			fa[x]=y,siz[y]+=siz[x];
//			return true;
//		}
//		return false;
//	}
//}d;
//
//int sgn(double x)//定义符号函数
//{
//    if(fabs(x)<eps)
//        return 0;
//    else if(x<0)
//        return -1;
//    else
//        return 1;
//}
//
//struct Point
//{
//    double x,y;
//    Point(){};//默认构造函数
//    Point(int _x,int _y)
//    {
//        x=_x;
//        y=_y;
//    }
//    Point operator-(const Point &b) const
//    {
//        return Point(x-b.x,y-b.y);
//    }
//    Point operator+(const Point &b) const
//    {
//        return Point(x+b.x,y+b.y);
//    }
//    double operator^(const Point &b) const//重载叉积
//    {
//        return x*b.y-y*b.x;
//    }
//    double operator*(const Point &b) const//点积
//    {
//        return x*b.x+y*b.y;
//    }
//    void transXY(double B)//绕原点旋转B弧度后，XY的值
//    {
//        double tx=x,ty=y;
//        x=tx*cos(B)-ty*sin(B);
//        y=tx*sin(B)+ty*cos(B);
//    }
//};
//
////圆
//struct Circle{
//	double r,x,y;
//	Circle(){};
//	Circle(int r,int x,int y):r(r),x(x),y(y){}
//};
//
////判断直线p1p2与圆c是否相交，相交返回true，否则返回false
//bool judge(Point p1,Point p2,Circle c)
//{
//	bool flag1=(p1.x-c.x)*(p1.x-c.x)+(p1.y-c.y)*(p1.y-c.y)<=c.r*c.r;
//	bool flag2=(p2.x-c.x)*(p2.x-c.x)+(p2.y-c.y)*(p2.y-c.y)<=c.r*c.r;
//	if(flag1&&flag2)	//情况一、两点都在圆内 :一定不相交
//	  return false;
//	else if(flag1||flag2) //情况二、一个点在圆内，一个点在圆外：一定相交
//	  return true;
//	else //情况三、两个点都在圆外
//	{
//
//		double A,B,C,dist1,dist2,angle1,angle2;
//		//将直线p1p2化为一般式：Ax+By+C=0的形式。先化为两点式，然后由两点式得出一般式
//		A=p1.y-p2.y;
//		B=p2.x-p1.x;
//		C=p1.x*p2.y-p2.x*p1.y;
//		//使用距离公式判断圆心到直线ax+by+c=0的距离是否大于半径
//		dist1=A*c.x+B*c.y+C;
//		dist1*=dist1;
//		dist2=(A*A+B*B)*c.r*c.r;
//		if(dist1>dist2)//圆心到直线p1p2的距离大于半径，不相交
//		  return false;
//		angle1=(c.x-p1.x)*(p2.x-p1.x)+(c.y-p1.y)*(p2.y-p1.y);
//		angle2=(c.x-p2.x)*(p1.x-p2.x)+(c.y-p2.y)*(p1.y-p2.y);
//		if(angle1>0&&angle2>0)//余弦为正，则是锐角，一定相交
//		  return true;
//		else
//		  return false;
//	}
//}
//
//
//struct Line
//{
//    Point s,e;
//    double k;
//    Line(){}
//    Line(Point _s,Point _e)
//    {
//        s = _s;e = _e;
//        k = atan2(e.y - s.y,e.x - s.x);//斜率
//    }
//    //两条直线求交点，
//    //第一个值为0表示直线重合，为1表示平行，为0表示相交,为2是相交
//    //只有第一个值为2时，交点才有意义
//    pair<int,Point> operator &(const Line &b)const
//    {
//        Point res = s;
//        if(sgn((s-e)^(b.s-b.e)) == 0)//如果两条线的叉积等于0，那么要么具有平行关系，要么重合
//        {
//            if(sgn((s-b.e)^(b.s-b.e)) == 0)//分别取两条边的任意两个点，进行判断即可
//                return make_pair(0,res);//重合
//            else return make_pair(1,res);//平行
//        }
//        double t = ((s-b.s)^(b.s-b.e))/((s-e)^(b.s-b.e));//利用正弦定理：a/sinA=b/sinB=c/sinC，分析可知，利用叉乘可以得出sin值
//        res.x += (e.x-s.x)*t;
//        res.y += (e.y-s.y)*t;
//        return make_pair(2,res);
//    }
//};
//double dist(Point a,Point b)
//{
//    return sqrt((a-b)*(a-b));
//}
//
//bool inter(Line l1,Line l2)
//{
//    return
//        max(l1.s.x,l1.e.x) >= min(l2.s.x,l2.e.x) &&     //前4项为矩形实验，后两项为两次跨立实验，跨立实验的思想是，任取一个点 都在另一条边的两侧，利用叉积值
//        max(l2.s.x,l2.e.x) >= min(l1.s.x,l1.e.x) &&     //..的正负性来判断
//        max(l1.s.y,l1.e.y) >= min(l2.s.y,l2.e.y) &&
//        max(l2.s.y,l2.e.y) >= min(l1.s.y,l1.e.y) &&
//        sgn((l2.s-l1.s)^(l1.e-l1.s))*sgn((l1.e-l1.s)^(l2.e-l1.s))>=0&&
//      sgn((l1.s-l2.s)^(l2.e-l2.s))*sgn((l2.e-l2.s)^(l1.e-l2.s))>=0;
//}
//
//bool cinter(int x1, int y1, int r1, int x2, int y2, int r2) {
//    double d_r = r1 + r2;
//
//    double d_x = x1 - x2;
//    double d_y = y1 - y2;
//    double d_distance = sqrt(d_x*d_x + d_y*d_y);
//
//    if(abs(r2-r1)>d_distance)return false;
//
//    if (d_distance > d_r) return false;
//    else return true;
//}
//#undef int
//class Solution {
//public:
//    vector<bool> antPass(vector<vector<int>>& g, vector<vector<int>>& p) {
//        int n=g.size();
//        d.init(n);
//        for(int i=0;i<n;++i)
//            for(int j=i+1;j<n;++j)
//                if(g[i].size()==4&&g[j].size()==4)
//                    if(inter(Line{Point{g[i][0],g[i][1]},Point{g[i][2],g[i][3]}},Line{Point{g[j][0],g[j][1]},Point{g[j][2],g[j][3]}}))
//                        d.merge(i,j);
//        for(int i=0;i<n;++i)
//            for(int j=i+1;j<n;++j)
//                if(g[i].size()==3&&g[j].size()==3)
//                    if(cinter(g[i][0],g[i][1],g[i][2],g[j][0],g[j][1],g[j][2]))
//                        d.merge(i,j);
//        for(int i=0;i<n;++i)
//            for(int j=0;j<n;++j)
//                if(i!=j&&g[i].size()==4&&g[j].size()==3)
//                    if(judge(Point{g[i][0],g[i][1]},Point{g[i][2],g[i][3]},Circle{g[j][2],g[j][0],g[j][1]}))
//                        d.merge(i,j);
//        vector<bool>ret;
//        for(const auto &i:p)ret.push_back(d.find(i[0])==d.find(i[1]));
//        return ret;
//    }
//};
