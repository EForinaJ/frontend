// utils/imageUrl.ts

// 定义平台类型
type PlatformType = 'h5' | 'mp-weixin' | 'mp-alipay' | 'mp-baidu' | 'mp-toutiao' | 'mp-qq' | 'mp-kuaishou' | 'app-plus' | 'app';

// 定义环境类型
type EnvType = 'development' | 'production' | 'test';

// 定义平台配置接口
interface PlatformConfig {
  baseUrl: string;
  publicPath: string;
}

// 定义环境配置接口
interface EnvConfig {
  h5: PlatformConfig;
  mp: PlatformConfig;
  app: PlatformConfig;
}

// 定义完整的配置结构
type ConfigStructure = Record<EnvType, EnvConfig>;

// 环境配置
const envConfig: ConfigStructure = {
  // 开发环境
  development: {
    h5: {
      baseUrl: '',
      publicPath: '/public'
    },
    mp: {
      baseUrl: 'http://127.0.0.1:8003',
      publicPath: '/public'
    },
    app: {
      baseUrl: 'http://192.168.1.100:8080',
      publicPath: '/public'
    }
  },
  // 生产环境
  production: {
    h5: {
      baseUrl: 'https://yourdomain.com',
      publicPath: '/public'
    },
    mp: {
      baseUrl: 'https://yourdomain.com',
      publicPath: '/public'
    },
    app: {
      baseUrl: 'https://yourdomain.com',
      publicPath: '/public'
    }
  },
  // 测试环境
  test: {
    h5: {
      baseUrl: 'https://test.yourdomain.com',
      publicPath: '/public'
    },
    mp: {
      baseUrl: 'https://test.yourdomain.com',
      publicPath: '/public'
    },
    app: {
      baseUrl: 'https://test.yourdomain.com',
      publicPath: '/public'
    }
  }
};

/**
 * 获取当前环境配置
 * @returns {PlatformConfig} - 当前环境的配置
 */
function getEnvConfig(): PlatformConfig {
  // 从环境变量获取环境类型，默认为开发环境
  const env: EnvType = (process.env.NODE_ENV as EnvType) || 'development';
  
  // 从环境变量获取平台类型，默认为 h5
  // UniApp 中可以通过 process.env.UNI_PLATFORM 获取当前平台
  const platform: string = process.env.UNI_PLATFORM || 'h5';
  
  // 判断平台类型
  let platformKey: keyof EnvConfig = 'h5'; // 默认 h5
  
  if (platform.startsWith('mp-')) {
    platformKey = 'mp';
  } else if (platform.startsWith('app')) {
    platformKey = 'app';
  }
  
  // 返回对应环境的配置，如果没有则返回开发环境 h5 配置
  return envConfig[env]?.[platformKey] || envConfig.development.h5;
}

/**
 * 检查是否为完整URL
 * @param {string} path - 路径字符串
 * @returns {boolean} - 是否为完整URL
 */
function isFullUrl(path: string): boolean {
  return path.startsWith('http://') || 
         path.startsWith('https://') || 
         path.startsWith('data:') ||
         path.startsWith('//');
}

/**
 * 获取完整的图片URL
 * @param {string} path - 图片路径，如：/public/uploads/xxx.png
 * @returns {string} - 完整的URL或相对路径
 */
export function getImageUrl(path: string): string {
  if (!path) return '';
  
  // 如果已经是完整URL，直接返回
  if (isFullUrl(path)) {
    return path;
  }
  
  const config = getEnvConfig();
  const platform: string = process.env.UNI_PLATFORM || 'h5';
  const env: EnvType = (process.env.NODE_ENV as EnvType) || 'development';
  
  // 开发环境下，不同平台处理方式不同
  if (env === 'development') {
    if (platform === 'h5') {
      // H5开发环境：使用代理，直接返回相对路径
      return path;
    } else {
      // 小程序/App开发环境：需要完整URL
      // 如果path已经以publicPath开头，直接拼接baseUrl
      if (path.startsWith(config.publicPath)) {
        return config.baseUrl + path;
      }
      // 否则先添加publicPath前缀
      return config.baseUrl + config.publicPath + (path.startsWith('/') ? '' : '/') + path;
    }
  }
  
  // 生产环境：使用完整URL
  // 如果path已经以publicPath开头，直接拼接baseUrl
  if (path.startsWith(config.publicPath)) {
    return config.baseUrl + path;
  }
  // 否则先添加publicPath前缀
  return config.baseUrl + config.publicPath + (path.startsWith('/') ? '' : '/') + path;
}

/**
 * 获取上传目录的URL
 * @param {string} filename - 文件名
 * @param {string} [dateStr] - 日期字符串，默认今天
 * @returns {string} - 完整的URL
 */
export function getUploadUrl(filename: string, dateStr?: string): string {
  // 如果没有提供日期字符串，使用今天
  if (!dateStr) {
    const date = new Date();
    dateStr = date.toISOString().split('T')[0]; // YYYY-MM-DD格式
  }
  
  const path = `${getEnvConfig().publicPath}/uploads/${dateStr}/${filename}`;
  return getImageUrl(path);
}

/**
 * 处理数据中的图片URL字段
 */
export function processImageUrls<T>(
  data: T, 
  fields: string[] = ['image', 'avatar', 'cover', 'logo', 'img', 'photo']
): T {
  if (!data) return data;
  
  const processItem = <U>(item: U): U => {
    if (!item || typeof item !== 'object') return item;
    
    if (Array.isArray(item)) {
      // 处理数组
      return item.map(processItem) as U;
    } else {
      // 处理对象
      const processed: Record<string, any> = {};
      
      for (const key in item) {
        if (item.hasOwnProperty(key)) {
          const value = (item as Record<string, any>)[key];
          
          if (fields.includes(key) && typeof value === 'string') {
            // 处理图片字段
            processed[key] = getImageUrl(value);
          } else if (typeof value === 'object' && value !== null) {
            // 递归处理嵌套对象
            processed[key] = processItem(value);
          } else {
            // 其他字段保持不变
            processed[key] = value;
          }
        }
      }
      
      return processed as U;
    }
  };
  
  return processItem(data);
}

/**
 * 批量获取图片URL
 * @param {string[]} paths - 图片路径数组
 * @returns {string[]} - 处理后的URL数组
 */
export function getImageUrls(paths: string[]): string[] {
  return paths.map(path => getImageUrl(path));
}

/**
 * 从完整URL中提取相对路径
 * @param {string} fullUrl - 完整URL
 * @returns {string} - 相对路径
 */
export function extractRelativePath(fullUrl: string): string {
  if (!fullUrl) return '';
  
  const config = getEnvConfig();
  
  // 如果URL以baseUrl开头，移除baseUrl
  if (fullUrl.startsWith(config.baseUrl)) {
    return fullUrl.substring(config.baseUrl.length);
  }
  
  return fullUrl;
}

/**
 * 判断图片URL是否需要处理
 * @param {string} url - 图片URL
 * @returns {boolean} - 是否需要处理
 */
export function needProcess(url: string): boolean {
  if (!url) return false;
  
  // 如果是完整URL且不以baseUrl开头，或者是以/public开头的相对路径，需要处理
  const config = getEnvConfig();
  const env: EnvType = (process.env.NODE_ENV as EnvType) || 'development';
  const platform: string = process.env.UNI_PLATFORM || 'h5';
  
  if (env === 'development' && platform === 'h5') {
    // H5开发环境：只处理非完整URL
    return !isFullUrl(url);
  } else {
    // 其他环境：处理不以baseUrl开头的URL
    return !url.startsWith(config.baseUrl) && !isFullUrl(url);
  }
}

// 类型定义导出
export type {
  PlatformType,
  EnvType,
  PlatformConfig,
  EnvConfig
};

// 默认导出
export default {
  getImageUrl,
  getUploadUrl,
  processImageUrls,
  getImageUrls,
  extractRelativePath,
  needProcess
};