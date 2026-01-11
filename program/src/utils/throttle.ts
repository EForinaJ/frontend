// utils/throttle.ts

interface ThrottleOptions {
  /** 间隔时间(毫秒) */
  wait?: number;
  /** 是否立即执行 */
  leading?: boolean;
  /** 是否在结束后执行 */
  trailing?: boolean;
}

interface ThrottledFunction<T extends (...args: any[]) => any> {
  (...args: Parameters<T>): ReturnType<T> | undefined;
  cancel: () => void;
}

class ThrottleUtils {
  /**
   * 创建节流函数
   */
  static create<T extends (...args: any[]) => any>(
    func: T,
    options: number | ThrottleOptions = {}
  ): ThrottledFunction<T> {
    const {
      wait = 1000,
      leading = true,
      trailing = false
    } = typeof options === 'number' ? { wait: options } : options;

    let timeoutId: ReturnType<typeof setTimeout> | null = null;
    let lastExecTime = 0;
    let result: ReturnType<T> | undefined;
    let lastArgs: Parameters<T> | null = null;

    const throttled = (...args: Parameters<T>): ReturnType<T> | undefined => {
      const currentTime = Date.now();
      const remainingWait = wait - (currentTime - lastExecTime);

      lastArgs = args;

      // 立即执行条件
      if (remainingWait <= 0 || remainingWait > wait) {
        if (timeoutId) {
          clearTimeout(timeoutId);
          timeoutId = null;
        }
        lastExecTime = currentTime;
        result = func.apply(this, args);
      } 
      // 设置延迟执行
      else if (!timeoutId && trailing) {
        timeoutId = setTimeout(() => {
          lastExecTime = Date.now();
          timeoutId = null;
          if (lastArgs) {
            result = func.apply(this, lastArgs);
          }
        }, remainingWait);
      }

      return result;
    };

    throttled.cancel = (): void => {
      if (timeoutId) {
        clearTimeout(timeoutId);
        timeoutId = null;
      }
      lastExecTime = 0;
      lastArgs = null;
    };

    return throttled as ThrottledFunction<T>;
  }
}

export default ThrottleUtils;