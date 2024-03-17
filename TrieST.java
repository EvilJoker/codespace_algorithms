import java.util.ArrayDeque;
import java.util.Deque;

/**
 * 题目名称：实现 Trie 字典树（TrieST）
 * 
 * 题目理解：
 * TrieST 是一种基于字典树（也称前缀树）的数据结构，用于存储字符串集合，并支持高效地插入、查找和删除操作。本题要求实现一个 TrieST
 * 类，不使用递归。
 * 
 * 解题思路：
 * 1. 定义一个 TrieNode 节点类来表示 Trie 树中的节点，每个节点包含若干子节点以及一个布尔值标志表示是否为结束节点。
 * 2. TrieST 类中定义根节点，并实现一系列方法包括 insert、contains、startsWith 和 size 等。
 * 3. 使用循环遍历输入字符串的字符，在 Trie 树中逐层向下创建或找到对应的子节点。
 * 4. 在插入过程中，遇到新字符则创建新的子节点并更新指针；若已存在，则直接移动指针至对应子节点。
 * 5. 查找和 startsWith 方法采用类似的方法遍历 Trie 树，前者需要检查最后一个节点是否为结束节点，后者只需要遍历到指定长度即可。
 * 
 * 注意要点：
 * - 每个节点可以有多个孩子，因此需用数组或其他动态数据结构（如哈希表）存储子节点。
 * - 结束节点是指代表完整字符串的节点，通常会有一个布尔标志标明其状态。
 * - 插入、查找和 startsWith 的时间复杂度主要取决于字符串长度，对于所有操作，平均时间复杂度均为 O(m)，其中 m 为待处理字符串的长度。
 * 
 * 时间复杂度：
 * - 插入操作（insert）、查找操作（contains）和前缀判断（startsWith）的时间复杂度均为 O(m)，其中 m 为字符串长度。
 * - 删除操作（未在代码示例中给出）的时间复杂度为 O(m)，且可能更复杂，因为需要处理子树为空的情况。
 * - 获取大小（size）的时间复杂度为 O(1)。
 * 
 */
public class TrieST<Value> {
    private static final int R = 26; // 假设只处理小写字母，字符集大小为26

    private static class Node {
        public Node[] children = new Node[R]; // 子节点数组, 如果该字符后还存在字符，那么children[r], 对应就有字符
        public Object value; // 结束节点关联的值
        public boolean isEndOfWord; // 是否为结束节点

        Node(){}

        Node(Object value, boolean isEndOfWord) {
            this.value = value;
            this.isEndOfWord = isEndOfWord;
        }
    }

    private Node root; // Trie 树根节点

    public TrieST() {
        root = new Node();
    }

    // 插入字符串及其关联值，更新字典树
    public void put(String key, Value val) {
        Node curNode = root;
        // 遍历字符串的每个字符，构造节点
        for (int i = 0; i < key.length(); i++) {
            char ch = key.charAt(i);
            int index = ch - 'a';

            if (curNode.children[index] == null) {
                curNode.children[index] = new Node();
            }
            // 更新指针，指向下一个节点
            curNode = curNode.children[index];
        }

        curNode.value = val;
        curNode.isEndOfWord = true;
    }

    // 判断字符串是否存在：
    public boolean contains(String key) {
        Node curNode = root;
        for (int i = 0; i < key.length(); i++) {
            char ch = key.charAt(i);
            int index = ch - 'a';
            // 如果当前节点不存在，则返回 false
            if (curNode.children[index] == null) {
                return false;
            }
            curNode = curNode.children[index];
        }

        // 阶段存在，且是结束节点
        return curNode != null && curNode.isEndOfWord;
    }

     // 辅助查找函数，从当前节点开始搜索匹配字符串
     private boolean search(Node node, String key, int d) {
        if (d == key.length()) {
            return node.isEndOfWord;
        }

        char ch = key.charAt(d);
        
        if (ch == '*') {
            // 处理通配符，遍历所有可能的字符分支
            for (int i = 0; i < R; i++) {
                // 对应每一个子节点,都去匹配
                Node child = node.children[i];
                if (child != null && search(child, key, d + 1)) {
                    return true;
                }
            }
        } else {
            int index = ch - 'a';
            Node child = node.children[index];
            if (child != null) {
                return search(child, key, d + 1);
            }
        }
        // 当没有一个 return 时，就是没有找到
        return false;

    }
    // 判断给定前缀是否存在
    public boolean startsWith(String prefix) {
        Node curNode = root;
        for (int i = 0; i < prefix.length(); i++) {
            char ch = prefix.charAt(i);
            int index = ch - 'a';

            if (curNode.children[index] == null) {
                return false;
            }
            curNode = curNode.children[index];
        }
        // 不需要结束节点
        return curNode != null;
    }

    // 获取 Trie 树中键的数量
    public int size() {
        int count = 0;
        Deque<Node> stack = new ArrayDeque<Node>();
        stack.push(root);

        while (!stack.isEmpty()) {
            Node node = stack.pop();
            if (node.isEndOfWord) {
                count++;
            }
            for (Node child : node.children) {
                if (child != null) {
                    stack.push(child);
                }
            }
        }

        return count;
    }

    // 其他辅助方法（比如删除操作）可以根据需要在此添加...

    // 示例主函数
    public static void main(String[] args) {
        TrieST<String> trie = new TrieST<>();
        trie.put("apple", "水果");
        trie.put("app", "应用");
        System.out.println(trie.contains("apple")); // 输出: true
        System.out.println(trie.startsWith("app")); // 输出: true
        System.out.println(trie.size()); // 输出: 2
    }
}