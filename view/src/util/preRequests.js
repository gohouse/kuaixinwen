import axios from 'axios'
import qs from 'qs'
import util from '@/util/util'

axios.defaults.timeout = 15000;
// axios.defaults.headers["Content-Type"] = 'application/x-www-form-urlencoded;charset=UTF-8';
axios.defaults.baseURL = 'http://localhost:8888';//api通用前缀设置
axios.defaults.baseURL = 'http://47.89.48.72:8888';//api通用前缀设置
// axios.defaults.baseURL = 'https://api.simboss.com';//api通用前缀设置
// axios.defaults.baseURL = 'http://mobile_card.cc';//api通用前缀设置

axios.interceptors.request.use(
    function (config) {
        if (config.method == 'post') {
            config.data = qs.stringify(config.data
                // {
            // ...config.data,
            // timestamp: timestamp,
            // appid: 123
        // }
        )
        }
        return config
    }, function (error) {
        return Promise.reject(error)
    }
);


export default axios