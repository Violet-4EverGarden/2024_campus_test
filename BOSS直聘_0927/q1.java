import java.util.Scanner;

// 多线程顺序打印数列
// 输出格式：1:1 2:2 3:3 1:4 2:5 3:6 ...
public class Main {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        while (in.hasNextInt()) {
            int k = in.nextInt();
            int n = in.nextInt();
            while (k != 0) {
                Thread t = new Thread(new CalculateTask(n));
                t.start();
                k--;
            }
        }
    }
}

class CalculateTask implements Runnable {
    private static volatile int count = 0;
    private static int threadTotal;
    private final int threadId;
    private int n;

    public CalculateTask(int n) {
        this.threadId = threadTotal;
        threadTotal++;
        this.n = n;
    }

    @Override
    public void run() {
        while (count < n) {
            synchronized (CalculateTask.class) {
                if (count % threadTotal == threadId) {
                    count++;
                    int id = threadId + 1;  // 加1是因为threadID是从0开始计数的，但题给要求从1开始
                    System.out.print(id + ":" + count + " ");
                    CalculateTask.class.notifyAll();
                } else {
                    try {
                        CalculateTask.class.wait();
                    } catch (InterruptedException e) {
                        throw new RuntimeException(e);
                    }
                }
            }
        }
    }
}