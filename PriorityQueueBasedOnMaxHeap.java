/**
 * 题目名称：Dijkstra 最短路径算法实现
 *
 * 题目理解：本题要求实现 Dijkstra 算法，该算法用于解决带权重的有向图或无向图中单源最短路径问题。
 *
 * 解题思路：
 * 1. 初始化：创建一个优先队列（最小堆）用于存储待处理节点及其距离，并将起始点的距离设为 0，其余节点的距离设为无穷大；同时初始化一个数组记录每个节点的前驱节点。
 * 2. 循环处理：在优先队列非空的情况下，取出当前已知距离最小的节点 u，然后遍历所有与节点 u 相邻的节点 v。
 * 对于每个相邻节点 v，计算从起始点经由 u 到达 v 的距离，如果这个距离比目前已知的到达 v 的距离更小，则更新 v 的距离值和前驱节点。
 * 3. 继续循环步骤 2，直到优先队列为空，此时所有节点的最短路径已经计算完毕。
 *
 * 注意要点：
 * - 边的权重需为非负数，否则 Dijkstra 算法可能无法正确找到最短路径。
 * - 使用优先队列（如最小堆）来保证每次都能获取到当前已知距离最小的节点。
 * - 更新节点距离时需要判断新路径是否更短，避免无效更新。

 * 时间复杂度：
 * - 在邻接表表示下，Dijkstra 算法的时间复杂度主要取决于优先队列的操作，因此时间复杂度大致为 O(ElogV)，其中 E 是边的数量，V 是顶点的数量。
 */

 import java.util.*;

 public class DijkstraShortestPath {
     
     private static final int INFINITY = Integer.MAX_VALUE; // 用最大整数值表示无穷大
     
     public static void dijkstra(Graph graph, int source) {
         int[] dist = new int[graph.v]; // 存储每个节点到起点的最短距离
         int[] predecessor = new int[graph.v]; // 存储每个节点的前驱节点
         
         Arrays.fill(dist, INFINITY); // 初始化所有节点距离为无穷大
         dist[source] = 0; // 起始点距离自身为 0
         
         // lambda 表达式，表示比较逻辑， dist[u] - dist[v] < 0 是 u 的优先级比较高， 小的数值优先级更高
         PriorityQueue<Integer> pq = new PriorityQueue<>((u, v) -> dist[u] - dist[v]); // 创建优先队列
         pq.offer(source); // 将起始点加入优先队列
         
         while (!pq.isEmpty()) {
             int u = pq.poll(); // 获取当前已知距离最小的节点
             
             // 遍历与节点 u 相邻的所有节点
             for (int w : graph.adj(u)) {
                // dist[u] 就是 u 到起点的最短距离，dist[u] 和 邻接点挨个相加。记录下最短的那个顶点。

                 int alt = dist[u] + 1; // 假设边权重为 1，实际应用中应根据实际权重计算
                 
                 // 如果通过 u 到达 w 的距离比已知距离更短，则更新
                 if (alt < dist[w]) {
                     dist[w] = alt;
                     predecessor[w] = u;
                     pq.offer(w); // 将节点 w 加入优先队列以便后续处理
                 }
                 // 经过几轮遍历后， dist[x]中就是距离source 最小的值，自然会停止
             }
         }
         
         // 这里可以输出最短路径或其他相关操作
         // dist[w]就是最短，
         // 打印路径，打印 predecessor, 直到source
     }
     // 求举例 s 最短的值
     public static void edgeTo(int target){
        System.out.println("最短路径为：" + dist[target]);
     }
 
     public static void main(String[] args) {
         Graph graph = new Graph(9);
         graph.addEdge(0, 1);
         graph.addEdge(0, 4);
         // ... 添加其他边
 
         dijkstra(graph, 0);
     }
 }