import java.util.ArrayList;
import java.util.List;

public class Graph {
    private final int v; // 顶点数目
    private int E; // 边数目
    private List<List<Integer>> adj;// 邻接表

    public Graph(int v) {

        this.v = v;
        // 初始化邻接表
        adj = new ArrayList<>(v);
        for (int i = 0; i < v; i++) {
            adj.add(new ArrayList<>());
        }
    }

    public void addEdge(int v, int w) {
        adj.get(v).add(w);
        adj.get(w).add(v); // 有向图， 需要删除这个
    }

    public Iterable<Integer> adj(int v) {
        return adj.get(v);
    }
}
